import React from 'react'
import { redirect } from 'next/navigation'
import { cookies } from 'next/headers'
import axios from 'axios'
import ChatBox from './components/ChatBox'
import Navbar from './components/Navbar'




export default async function page() {
  const cookieStore = await cookies()
  const auth = cookieStore.get('auth')

  if (auth === undefined) {
    redirect('/login')
  }

  let users = []


  const res = await axios.get('http://127.0.0.1:80/users/all-users', {
      headers: {
        'Authorization': auth.value
      }
    })

  if (res.status === 200) {
    users = res.data
  }

  const updateActiveUser = (userIdx) => {
    activeChat = users[userIdx]
  }

  return (
    <div className='mx-auto w-1/2 grid grid-cols-5 gap-4'>
      
      <Navbar />

      <ChatBox users={users}/>

      
    </div>
  )
}
