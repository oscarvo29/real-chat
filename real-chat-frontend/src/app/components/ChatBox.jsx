'use client'

import React, { useState, useEffect } from 'react'
import UserButtons from './UserButtons'

export default function ChatBox({ users }) {
    const [chatChosen, setChatChosen] = useState(false)
    const [activeChat, setActiveChat] = useState({})

    const handleUserClick = (userIndex) => {
        setActiveChat(users[userIndex])
    }

    useEffect(() => {
        setChatChosen(true)
    }, [activeChat])


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
                        <input type="text" className='cols-span-4 text-black' />
                        <input type="button" value="Send" className='cols-span-1 rounded border-2'/>
                    </div>
                </>
                )}
            </div>
        </>
    )
}
