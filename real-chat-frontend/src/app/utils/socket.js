

export function GetSocketConnection(jwtToken) {

    const ws = new WebSocket("ws://127.0.0.1:80/chat-ws")

    ws.onopen = () => {
        const data = {
            "event": "connection_open",
            "jwt": jwtToken,
            "data": {}
        }

        ws.send(JSON.stringify(data))
        console.log("Connected to the server!")

    }

    ws.onerror = (error) => {
        console.error("WebSocket error:", error);
    };

    ws.onclose = () => {
        const data = {
            "event": "conn_close",
            "jwt": jwtToken,
            "data": {}
        }

        ws.send(JSON.stringify(data))

        console.log("WebSocket connection closed");
    };

    return ws
}