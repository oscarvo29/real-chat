'use client'

import React, { useState, useEffect } from 'react'
import UserButtons from './UserButtons'
import axios from 'axios'
import Cookies from 'js-cookie'

export default function ChatBox({ users }) {
    const [chatChosen, setChatChosen] = useState(false)
    const [activeChat, setActiveChat] = useState({})

    const handleUserClick = (userIndex) => {
        setActiveChat(users[userIndex])
    }

    useEffect(() => {
        setChatChosen(true)
    }, [activeChat])

    const sendMessage = async (e) => {
        e.preventDefault()
        let msg = document.querySelector('#msgField').value
        const jwtToken = Cookies.get("auth")

        console.log(activeChat)
        const msgPayload = {
            "receiver_uuid": activeChat.uuid,
            "message_value": msg 
        }

        const config = {
            headers: {
                'Content-type': 'application/json',
                'Authorization': jwtToken
            }
        }

        
        let res = await axios.post('http://127.0.0.1:80/messages/send-message', msgPayload, config)
        if (res.status == 200) {
            console.log("send succesfully.")
        }
        
    }

    return (
        <>
            <ul>
                {users.map((user, index) => <UserButtons key={user.uuid} user={user} index={index} handler={handleUserClick} />)}
            </ul>
            <div className="col-span-3 rounded border-2">
                { chatChosen === false ? (
                    <p className='display-block my-auto text-center'>No chat is active!</p>
                ) : (
                <>
                    <div className="">
                    <p>Chat message</p>
                    </div>
                    <div className='grid grid-cols-5 gap-5'>
                        <input type="text" id='msgField'  className='cols-span-4 text-black' />
                        <input type="button" value="Send" onClick={sendMessage} className='cols-span-1 rounded border-2'/>
                    </div>
                </>
                )}
            </div>
        </>
    )
}
