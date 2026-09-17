package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"client-simulator/internal/auth"
	"client-simulator/internal/bot"
	"client-simulator/internal/conversation"
	"client-simulator/internal/metrics"
)

func main() {
	numBots := flag.Int("bots", 10, "Số lượng bot muốn tạo và chạy mô phỏng")
	convsPerBot := flag.Int("convs", 5, "Số lượng cuộc hội thoại 1-1 tối đa trên mỗi bot")
	sendInterval := flag.Duration("interval", 1*time.Second, "Tần suất gửi tin nhắn của mỗi bot (vd: 500ms, 1s, 2s)")
	apiURL := flag.String("api", "http://localhost", "URL gốc của API Service (hoặc NGINX Ingress)")
	wsURL := flag.String("ws", "ws://localhost/ws", "URL endpoint WebSocket của WS Gateway (hoặc NGINX Ingress)")
	testDuration := flag.Duration("duration", 0, "Thời gian chạy test (0 = chạy liên tục đến khi Ctrl+C)")
	botPassword := flag.String("password", "Pass@123456", "Mật khẩu chung cho các bot tài khoản")
	flag.Parse()

	log.Println("=================================================================")
	log.Println("           🚀 KHỞI ĐỘNG CHAT CLIENT SIMULATOR                   ")
	log.Println("=================================================================")
	log.Printf("Cấu hình: %d bots | ~%d convs/bot | Interval: %s", *numBots, *convsPerBot, *sendInterval)
	log.Printf("API Service: %s | WS Gateway: %s", *apiURL, *wsURL)

	rootCtx, rootCancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer rootCancel()

	authClient := auth.NewAuthClient(*apiURL)
	convoManager := conversation.NewManager(*apiURL)
	tracker := metrics.NewTracker()

	// 1. GIAI ĐOẠN 1: Authenticate / Register Bots song song (Concurrency = 30)
	log.Printf("[Phase 1] Authenticating %d bots concurrently via API Service...", *numBots)
	botSessions := make([]*auth.BotSession, *numBots)
	var authWg sync.WaitGroup
	authSem := make(chan struct{}, 30)
	var authErrMu sync.Mutex
	var firstAuthErr error

	for i := 1; i <= *numBots; i++ {
		authWg.Add(1)
		go func(idx int) {
			defer authWg.Done()
			authSem <- struct{}{}
			defer func() { <-authSem }()

			username := fmt.Sprintf("sim_bot_%04d", idx)
			session, err := authClient.EnsureBotAuth(username, *botPassword)
			if err != nil {
				authErrMu.Lock()
				if firstAuthErr == nil {
					firstAuthErr = fmt.Errorf("failed to authenticate bot %s: %w", username, err)
				}
				authErrMu.Unlock()
				return
			}
			botSessions[idx-1] = session
		}(i)
	}
	authWg.Wait()

	if firstAuthErr != nil {
		log.Fatalf("Authentication failed: %v", firstAuthErr)
	}
	log.Printf("  - Successfully authenticated %d/%d bots in parallel!", len(botSessions), *numBots)

	// 2. GIAI ĐOẠN 2: Setup Conversation Network song song
	log.Printf("[Phase 2] Setting up direct conversations between bots...")
	botTargets := convoManager.EnsureDirectConversations(botSessions, *convsPerBot)

	// 3. GIAI ĐOẠN 3: Connect WebSocket Clients (Smooth Ramp-Up Concurrency = 20)
	log.Printf("[Phase 3] Establishing WebSocket connections for %d bots...", len(botSessions))
	bots := make([]*bot.Bot, len(botSessions))

	var wsWg sync.WaitGroup
	wsSem := make(chan struct{}, 20)
	var activeCount int64
	var countMu sync.Mutex

	for idx, session := range botSessions {
		if session == nil {
			continue
		}
		wsWg.Add(1)
		time.Sleep(5 * time.Millisecond) // Smooth TCP syn pacing for Windows
		go func(i int, s *auth.BotSession) {
			defer wsWg.Done()
			wsSem <- struct{}{}
			defer func() { <-wsSem }()

			targets := botTargets[s.UserID]
			b := bot.NewBot(s, *wsURL, targets, tracker)

			connectCtx, connectCancel := context.WithTimeout(rootCtx, 20*time.Second)
			defer connectCancel()

			if err := b.Connect(connectCtx); err != nil {
				log.Printf("WARN: Bot %s failed to connect WS: %v", s.Username, err)
				return
			}

			bots[i] = b

			countMu.Lock()
			activeCount++
			countMu.Unlock()
		}(idx, session)
	}
	wsWg.Wait()

	log.Printf("Successfully connected %d/%d bots to WebSocket Gateway!", activeCount, *numBots)
	log.Println("🚀 Kích hoạt đồng loạt 150 bots bắt đầu phát sinh tin nhắn...")
	for _, b := range bots {
		if b != nil {
			b.Start(rootCtx, *sendInterval)
		}
	}
	time.Sleep(500 * time.Millisecond)

	// 4. GIAI ĐOẠN 4: Live Dashboard Display Loop & Test Duration
	testCtx := rootCtx
	if *testDuration > 0 {
		log.Printf("⏱️ Bắt đầu đo tải: Chạy trong %s rồi tự động dừng...", *testDuration)
		var cancelTimer context.CancelFunc
		testCtx, cancelTimer = context.WithTimeout(rootCtx, *testDuration)
		defer cancelTimer()
	} else {
		log.Println("♾️ Bắt đầu đo tải liên tục (nhấn Ctrl+C để dừng)...")
	}

	startTime := time.Now()
	dashboardTicker := time.NewTicker(1 * time.Second)
	defer dashboardTicker.Stop()

	for {
		select {
		case <-testCtx.Done():
			log.Println("\nStopping simulator and closing connections...")
			for _, b := range bots {
				if b != nil {
					b.Stop()
				}
			}

			// Print final summary
			tracker.PrintDashboard(0, *numBots, time.Since(startTime))
			log.Println("\n✅ Simulator finished cleanly.")
			return

		case <-dashboardTicker.C:
			currentActive := 0
			for _, b := range bots {
				if b.IsConnected() {
					currentActive++
				}
			}
			tracker.PrintDashboard(currentActive, *numBots, time.Since(startTime))
		}
	}
}
