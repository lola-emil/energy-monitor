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
        // Don't recreate an already active connection
        if (
            this.eventSource &&
            this.eventSource.readyState !== EventSource.CLOSED
        ) {
            return;
        }

        const API_URL = import.meta.env.VITE_API_URL;

        const eventSource = new EventSource(
            `${API_URL}/stream`
        );

        this.eventSource = eventSource;

        eventSource.onopen = () => {
            console.log("SSE connected");
        };

        eventSource.onmessage = (event) => {
            try {
                const payload: ReadingUpdatePayload =
                    JSON.parse(event.data);

                if (payload.type === "reading_update") {
                    onMessage(payload);
                }
            } catch (error) {
                console.error("SSE parse error:", error);
            }
        };

        eventSource.onerror = () => {
            if (eventSource.readyState === EventSource.CONNECTING) {
                console.warn("SSE disconnected, reconnecting...");
                return;
            }

            if (eventSource.readyState === EventSource.CLOSED) {
                console.warn("SSE connection closed");
            }
        };
    }

    disconnect() {
        this.eventSource?.close();
        this.eventSource = null;
    }
}

export const sseService = new SSEService();