class ChatHistory {
    messageUuid
    senderUuid
    receiverUuid
    messageValue
    sendTime
    read
    readAt

    constructor(messageUuid, senderUuid, receiverUuid, messageValue, sendTime, read, readAt) {
        this.messageUuid = messageUuid
        this.senderUuid = senderUuid
        this.receiverUuid = receiverUuid
        this.messageValue = messageValue
        this.sendTime = sendTime 
        this.read = read
        this.readAt
    }
}

export default ChatHistory