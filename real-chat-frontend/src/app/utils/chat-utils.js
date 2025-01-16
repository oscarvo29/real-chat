const { default: axios } = require("axios");
import Cookies from 'js-cookie';

export async function FetchChatHistory(receiverUuid) {
    const jwtToken = Cookies.get("auth")
    if (jwtToken === "") {
        return
    }
    console.log("Receiver uiid: ",receiverUuid)

    const config = {
        headers: {
            'Content-type': 'application/json',
            'Authorization': jwtToken
        }
    }
    

    const body = {
        'receiver_uuid': receiverUuid,
    }


    
    let res = await axios.post('http://127.0.0.1:80/messages/get-chat-history', body, config)
    if (res.status !== 200 ) {
        return []
    }

    return res.data
} 

export async function SendChatRequest(msg, activeChatUuid, jwtToken) {

    const msgPayload = {
        "receiver_uuid": activeChatUuid,
        "message_value": msg 
    }

    const config = {
        headers: {
            'Content-type': 'application/json',
            'Authorization': jwtToken
        }
    }

    
    let res = await axios.post('http://127.0.0.1:80/messages/send-message', msgPayload, config)
    return res.status
}
