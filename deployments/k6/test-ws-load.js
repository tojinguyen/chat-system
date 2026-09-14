import ws from 'k6/ws';
import { check, sleep } from 'k6';
import { Trend, Counter } from 'k6/metrics';

// Custom Prometheus/VictoriaMetrics compatible Trends & Counters
const wsRoundTripTime = new Trend('chat_ws_round_trip_duration_ms', true);
const wsMessagesSent = new Counter('chat_ws_messages_sent_total');
const wsMessagesReceived = new Counter('chat_ws_messages_received_total');
const wsConnectionErrors = new Counter('chat_ws_connection_errors_total');

export const options = {
  scenarios: {
    stress_test: {
      executor: 'ramping-vus',
      startVUs: 100,
      stages: [
        { duration: '1m', target: 1000 },   // Warm-up 1k VUs
        { duration: '3m', target: 10000 },  // Scale to 10k VUs
        { duration: '5m', target: 50000 },  // Extreme High-Load 50k VUs
        { duration: '5m', target: 100000 }, // Breaking point 100k VUs
        { duration: '2m', target: 0 },      // Ramp down
      ],
      gracefulStop: '30s',
    },
  },
  thresholds: {
    'chat_ws_round_trip_duration_ms': ['p(95)<150', 'p(99)<300'], // SOTA SLA p99 < 300ms
    'chat_ws_connection_errors_total': ['count<100'],
  },
};

export default function () {
  const vuId = __VU;
  const iterId = __ITER;
  const userId = `user_${vuId}`;
  const deviceId = `device_${vuId}_${iterId}`;
  const receiverId = `user_${(vuId % 1000) + 1}`; // Gửi vòng tròn giữa các active VUs

  // Gateway URL (Load Balanced across ws-gateway nodes)
  const gatewayUrl = __ENV.WS_GATEWAY_URL || 'ws://nginx-gateway.chat-system.svc.cluster.local:80/ws';
  const url = `${gatewayUrl}?user_id=${userId}&device_id=${deviceId}`;

  const params = {
    headers: {
      'User-Agent': 'k6-load-generator/1.0',
    },
  };

  const res = ws.connect(url, params, function (socket) {
    socket.on('open', function () {
      // 1. Send Heartbeat immediately
      socket.send(JSON.stringify({
        type: 'HEARTBEAT',
        client_msg_id: `hb_${vuId}_${Date.now()}`,
        payload: {}
      }));

      // 2. Định kỳ bắn tin nhắn mỗi 2-3 giây
      socket.setInterval(function () {
        const clientMsgId = `k6_${vuId}_${Date.now()}`;
        const sendTimestamp = Date.now();

        const messagePayload = {
          type: 'SEND_MESSAGE',
          client_msg_id: clientMsgId,
          payload: {
            conversation_id: `conv_${Math.min(vuId, 100)}`,
            receiver_id: receiverId,
            content: `Stress payload from VU ${vuId} at ${sendTimestamp}`,
          }
        };

        socket.send(JSON.stringify(messagePayload));
        wsMessagesSent.add(1);
      }, 2500);
    });

    socket.on('message', function (data) {
      try {
        const msg = JSON.parse(data);
        if (msg.type === 'MESSAGE_DELIVERED' || msg.type === 'MESSAGE_SUBMITTED') {
          wsMessagesReceived.add(1);
          if (msg.timestamp) {
            const rtt = Date.now() - msg.timestamp;
            if (rtt > 0 && rtt < 60000) {
              wsRoundTripTime.add(rtt);
            }
          }
        }
      } catch (err) {
        // ignore malformed frame
      }
    });

    socket.on('error', function (e) {
      wsConnectionErrors.add(1);
    });

    socket.on('close', function () {
      // socket closed
    });

    // Giữ kết nối trong 60 giây trước khi VU bắt đầu vòng lặp tiếp theo
    socket.setTimeout(function () {
      socket.close();
    }, 60000);
  });

  check(res, { 'WebSocket connected successfully': (r) => r && r.status === 101 });
}
