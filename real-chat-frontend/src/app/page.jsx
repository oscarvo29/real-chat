import React from 'react'
import { redirect } from 'next/navigation'
import { cookies } from 'next/headers'
import axios from 'axios'

export default async function page() {
  const cookieStore = await cookies()
  const auth = cookieStore.get('auth')
  let users = []

  if (auth === undefined) {
    redirect('/login')
  }

  const res = await axios.get('http://127.0.0.1:80/users/all-users', {
    headers: {
      'Authorization': auth.value
    }
  })

  if (res.status === 200) {
    users = res.data
  }

  return (
    <div>
      
      <h2>Index Page:</h2>

      {users.map((user, index) => {
          return (
            <div key={index} className="">
              <h2>{user.name}</h2>
            </div> 
          )
      })}
    </div>
  )
}
