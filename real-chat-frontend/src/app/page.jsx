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

  let users = []


  const res = await axios.get(GetURL('users/all-users'), {
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
      <h1>HEJ MATE</h1>
      <Navbar />

      <ChatBox users={users}/>

      
    </div>
  )
}
