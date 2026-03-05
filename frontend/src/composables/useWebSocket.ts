import { ref } from "vue";


export function useWebSocket(url = "ws://localhost:5000/ws") {
    const messages = ref<string[]>([]);
    const socket = ref<WebSocket | null>(null);
    const isConnected = ref(false);

    const connect = () => {
        socket.value = new WebSocket(url);

        socket.value.onopen = () => {
            isConnected.value = true;
            console.log("Connected to WebSocket Server")
        };

        socket.value.onmessage = (event) => {
            messages.value.push("Recieved: " + event.data);
        };

        socket.value.onclose = () => {
            isConnected.value = false;
            messages.value.push("Disconnected from WebSocket server");
        }

        socket.value.onerror = (error) => {
            messages.value.push("WebSocket error: " + error);
        }
    }

    const sendMessage = (message: string) => {
        if (socket.value && isConnected.value) {
            socket.value.send(message);
            messages.value.push("Sent: " + message);
        } else {
            messages.value.push("Cannot send, socket is not connected");
        }
    }


    return {
        messages,
        socket,
        isConnected,

        connect,
        sendMessage,
    };
}