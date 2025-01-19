'use client'

import React, { useState, useEffect } from 'react'
import UserButtons from './UserButtons'
import Cookies from 'js-cookie'

import {FetchChatHistory, SendChatRequest} from '../utils/chat-utils'
import ChatBubble from './ChatBubble'
import { GetSocketConnection } from '../utils/socket'
import CreateChat from './CreateChat'
import { Button } from '@material-tailwind/react'



export default function ChatBox({ chatRooms }) {
    const [chatChosen, setChatChosen] = useState(false) // TODO: lav check på om activeChat er sat i stedet.
    const [ openModal, setOpenModal] = useState(false)
    const [activeChat, setActiveChat] = useState({})
    const [chatLog, setChatLog] = useState([])
    const [socket, setSocket] = useState(null)


    useEffect(() => {
        console.log(chatRooms)
        const jwtToken = Cookies.get("auth")

        const ws = GetSocketConnection(jwtToken)
        setSocket(ws)

        ws.onmessage = (event) => {
            if (event.data) {
                console.log(event)
                let msg = JSON.parse(event.data)
                setChatLog((prev) => [...prev, msg])
            }
        }
    
        return () => {
            ws.close(); // Cleanup connection when the component unmounts
        };
    }, [])

    const handleUserClick = async (chatIndex) => {
        setActiveChat(chatRooms[chatIndex]) 
    }

    useEffect(() => {
        // checks wether or not, an active chat have been chosen. 
        if (Object.keys(activeChat).length !== 0) {
            setChatChosen(true)
            FetchChatHistory(activeChat.chat_uuid).then((data) => {
                console.log("data:")
                console.log( data)
                if (data) {
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
            "event": "message",
            "jwt": jwtToken,
            "data": {
                "jwt": jwtToken,
                "chat_id": activeChat.chat_uuid,
                "message": msgValue,
            }
        }

        if (socket && socket.readyState === WebSocket.OPEN) {
            console.log("msg send.")
            socket.send(JSON.stringify(msg))
        }

        msgField.value = ""
    }

    const closeModalBox = () => setOpenModal((prevState) => !prevState)

    return (
        <>
            <CreateChat open={openModal} handleOpen={closeModalBox} />
            <div className="col-span-5 flex justify-end">
                    <Button className="btn rounded " onClick={closeModalBox} >Create new Chat!</Button>
            </div>
            <ul>
                {chatRooms.map((room, index) => <UserButtons key={room.chat_uuid} name={room.chat_name} index={index} handler={handleUserClick} />)}
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

                                if(chatItem.is_sender) {
                                    return <ChatBubble key={`${index}:${chatItem.message_value}`} msg={chatItem.message_value} name="" sendAt={chatItem.send_time} receivedMsg={false}/>
                                } else {
                                    return <ChatBubble key={`${index}:${chatItem.message_value}`} msg={chatItem.message_value} name="" sendAt={chatItem.send_time} receivedMsg={true}/>
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
