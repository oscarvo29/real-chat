

export function GetSocketConnection() {
    const ws = new WebSocket("ws://127.0.0.1:80/chat-ws")

    ws.onopen = () => {
        console.log("Connected to the server!")
    }

    ws.onerror = (error) => {
        console.error("WebSocket error:", error);
    };

    ws.onclose = () => {
        console.log("WebSocket connection closed");
    };

    return ws
}