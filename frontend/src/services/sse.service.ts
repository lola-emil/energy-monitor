type ReadingUpdatePayload = {
    type: string;
    device_code: string;

    power: number;
    voltage: number;
    current: number;
    energy_kwh: number;

    timestamp: string;
};

class SSEService {
    private eventSource: EventSource | null = null;

    connect(
        onMessage: (payload: ReadingUpdatePayload) => void
    ) {
        if (this.eventSource) {
            this.eventSource.close();
        }

        this.eventSource = new EventSource(
            "http://localhost:5000/api/stream"
        );

        this.eventSource.onmessage = (event) => {
            try {
                const payload = JSON.parse(event.data);

                if (payload.type === "reading_update") {
                    onMessage(payload);
                }
            } catch (error) {
                console.error("SSE parse error:", error);
            }
        };

        this.eventSource.onerror = (error) => {
            console.error("SSE connection error:", error);
        };
    }

    disconnect() {
        if (this.eventSource) {
            this.eventSource.close();
            this.eventSource = null;
        }
    }
}

export const sseService = new SSEService();