const { default: axios } = require("axios");
import Cookies from 'js-cookie';
import { GetURL } from './http-store';

export async function FetchChatHistory(chatId) {
    const jwtToken = Cookies.get("auth")
    if (jwtToken === "") {
        return
    }

    const config = {
        headers: {
            'Content-type': 'application/json',
            'Authorization': jwtToken
        }
    }
    
    let res = await axios.get(GetURL(`messages/get-chat-history/${chatId}`), config)
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

    
    let res = await axios.post(GetURL('messages/send-message'), msgPayload, config)
    return res.status
}
