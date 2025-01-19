import React from 'react'
import { redirect } from 'next/navigation'
import { cookies } from 'next/headers'
import axios from 'axios'
import ChatBox from './components/ChatBox'
import Navbar from './components/Navbar'
import { GetURL, SetUpUrlStore } from './utils/http-store'



export default async function page() {
  const cookieStore = await cookies()
  const auth = cookieStore.get('auth')



  if (auth === undefined) {
    redirect('/login')
  }

  let chatRooms = []


  const res = await axios.get(GetURL('messages/get-chats'), {
      headers: {
        'Authorization': auth.value
      }
  })
  console.log("res status code: ", res.status)
  if (res.status === 200) {
    chatRooms = res.data
    console.log("Chatrooms", chatRooms)
  }


  return (
    <div className='mx-auto w-1/2 grid grid-cols-5 gap-4'>      
      <h1>HEJ MATE</h1>
      <Navbar />

      <ChatBox chatRooms={chatRooms}/>

      
    </div>
  )
}
