'use client'

import React, { useState, useEffect } from 'react'
import UserButtons from './UserButtons'
import axios from 'axios'
import Cookies from 'js-cookie'

import {FetchChatHistory, SendChatRequest} from '../utils/chat-utils'
import ChatBubble from './ChatBubble'
import { GetSocketConnection } from '../utils/socket'



export default function ChatBox({ users }) {
    const [chatChosen, setChatChosen] = useState(false)
    const [activeChat, setActiveChat] = useState({})
    const [chatLog, setChatLog] = useState([])
    const [input, setInput] = useState("")
    const [socket, setSocket] = useState(null)

    
    

    useEffect(() => {
        const ws = GetSocketConnection()
        setSocket(ws)

        ws.onmessage = (event) => {
            if (event.data) {
                let msg = JSON.parse(event.data)
                setChatLog((prev) => [...prev, msg])
            }
        }
    
        return () => {
            ws.close(); // Cleanup connection when the component unmounts
        };
    }, [])

    const handleUserClick = async (userIndex) => {
        setActiveChat(users[userIndex]) 
    }

    useEffect(() => {
        // checks wether or not, an active chat have been chosen. 
        if (Object.keys(activeChat).length !== 0) {
            setChatChosen(true)
            FetchChatHistory(activeChat.uuid).then((data) => {
                if (data) {
                    console.log(data    )
                    setChatLog((prev) => [...prev, ...data])
                }
            })
        }


    }, [activeChat])

    const sendMessage = async (e) => {
        e.preventDefault()
        let msgField = document.querySelector('#msgField') 
        let msgValue = msgField.value
        const jwtToken = Cookies.get("auth")

        const msg = {
            "jwt": jwtToken,
            "receiver_uuid": activeChat.uuid,
            "message": msgValue,
        }

        if (socket && socket.readyState === WebSocket.OPEN) {
            socket.send(JSON.stringify(msg))
        }

        msgField.value = ""
    }

    return (
        <>
            <ul>
                {users.map((user, index) => <UserButtons key={user.uuid} user={user} index={index} handler={handleUserClick} />)}
            </ul>
            <div className="col-span-3 rounded border-2 grid grid-cols-5">
                { chatChosen === false ? (
                    <p className='display-block my-auto text-center col-span-5'>No chat is active!</p>
                ) : (
                <>
                    <div className="p-5 flex flex-col overflow-scroll h-96 col-span-5 ">
                        { (!chatLog || chatLog.length === 0)  ? (
                            <p className='display-block my-auto text-center col-span-5'>Send a message and say hey!</p>
                             
                        ) : (
                            chatLog.map((chatItem, index) => {

                                if(chatItem.sender_uuid === activeChat.uuid) {
                                    return <ChatBubble key={`${index}:${chatItem.message_value}`} msg={chatItem.message_value} name={activeChat.name} sendAt={chatItem.send_time} receivedMsg={true}/>
                                } else {
                                    return <ChatBubble key={`${index}:${chatItem.message_value}`} msg={chatItem.message_value} name="Ozzz" sendAt={chatItem.send_time} receivedMsg={false}/>
                                }
                            })
                        )}
                    </div>
                    <div className='grid grid-cols-5 gap-5 col-span-5 p-5'>
                        
                        <input type="text" id='msgField' className='col-span-4 text-black bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500' />
                        <input type="button" value="Send" onClick={sendMessage} className='col-span-1 rounded border-2'/>
                    </div>
                </>
                )}
            </div>
        </>
    )
}


  /*
    
    message_id: "72aa9ade-9bf1-4339-939c-b4db2018ceb7"
​​
message_value: "hey mate!"
​​
read: false
​​
read_at: Object { Time: "0001-01-01T00:00:00Z", Valid: false }
​​
receiver_uuid: "f9d75903-9179-4865-adb1-13c22503417a"
​​
send_time: "2025-01-15T16:37:47.708718Z"
​​
sender_uuid: "c1d9a936-275b-4cb4-88b4-3452ff76de51"

    */