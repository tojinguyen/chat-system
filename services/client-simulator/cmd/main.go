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

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if *testDuration > 0 {
		log.Printf("⏱️ Chế độ: Chạy trong %s rồi tự động dừng.", *testDuration)
		var cancelTimer context.CancelFunc
		ctx, cancelTimer = context.WithTimeout(ctx, *testDuration)
		defer cancelTimer()
	} else {
		log.Println("♾️ Chế độ: Chạy liên tục vô hạn (nhấn Ctrl+C để dừng và xem báo cáo tổng kết).")
	}

	authClient := auth.NewAuthClient(*apiURL)
	convoManager := conversation.NewManager(*apiURL)
	tracker := metrics.NewTracker()

	// 1. GIAI ĐOẠN 1: Authenticate / Register Bots
	log.Printf("[Phase 1] Authenticating %d bots via API Service...", *numBots)
	botSessions := make([]*auth.BotSession, 0, *numBots)

	for i := 1; i <= *numBots; i++ {
		username := fmt.Sprintf("sim_bot_%04d", i)
		session, err := authClient.EnsureBotAuth(username, *botPassword)
		if err != nil {
			log.Fatalf("Failed to authenticate bot %s: %v", username, err)
		}
		botSessions = append(botSessions, session)
		if i%10 == 0 || i == *numBots {
			log.Printf("  - Authenticated: %d/%d bots", i, *numBots)
		}
	}

	// 2. GIAI ĐOẠN 2: Setup Conversation Network
	log.Printf("[Phase 2] Setting up direct conversations between bots...")
	botTargets := convoManager.EnsureDirectConversations(botSessions, *convsPerBot)

	// 3. GIAI ĐOẠN 3: Connect WebSocket Clients
	log.Printf("[Phase 3] Establishing WebSocket connections for %d bots...", len(botSessions))
	bots := make([]*bot.Bot, 0, len(botSessions))

	var activeBotsCount sync.WaitGroup
	var activeCount int64

	for _, session := range botSessions {
		targets := botTargets[session.UserID]
		b := bot.NewBot(session, *wsURL, targets, tracker)

		if err := b.Connect(ctx); err != nil {
			log.Printf("WARN: Bot %s failed to connect WS: %v", session.Username, err)
			continue
		}

		b.Start(ctx, *sendInterval)
		bots = append(bots, b)
		activeCount++
	}

	log.Printf("Successfully connected %d/%d bots to WebSocket Gateway!", activeCount, *numBots)
	time.Sleep(1 * time.Second)

	// 4. GIAI ĐOẠN 4: Live Dashboard Display Loop
	startTime := time.Now()
	dashboardTicker := time.NewTicker(1 * time.Second)
	defer dashboardTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("\nStopping simulator and closing connections...")
			for _, b := range bots {
				b.Stop()
			}
			activeBotsCount.Wait()

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
