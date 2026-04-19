// stores/sse.js
import { defineStore } from 'pinia';
import { ref } from 'vue';


export const useSSEStore = defineStore('sse', () => {
    const eventSource = ref<EventSource | null>(null);
    const listeners = ref<{ [key: string]: (event: MessageEvent<string>) => void; }>({});

    const connect = (userId: string) => {
        if (eventSource.value) return;

        const URL = import.meta.env.VITE_API_URL ?? "http://localhost:5000"
        eventSource.value = new EventSource(`${URL}/events?userid=${userId}`);
    };

    const subscribe = (eventName: string, callback: (event: MessageEvent<string>) => void) => {
        if (!eventSource.value) return;

        console.log("Subscribed");

        eventSource.value.addEventListener(eventName, callback);

        listeners.value[eventName] = callback;
    };

    const unsubscribe = (eventName: string) => {
        if (!eventSource.value || !listeners.value[eventName]) return;
        console.log("Unsubscribed");

        eventSource.value.removeEventListener(
            eventName,
            listeners.value[eventName]
        );

        delete listeners.value[eventName];
    };


    const disconnect = () => {
        if (eventSource.value) {
            eventSource.value.close();
            eventSource.value = null;
            listeners.value = {};
        }
    };

    return {
        eventSource,
        listeners,

        connect,
        subscribe,
        unsubscribe,
        disconnect
    }
});