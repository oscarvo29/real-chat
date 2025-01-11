import React from 'react'
import { redirect } from 'next/navigation'
import { cookies } from 'next/headers'

export default async function page() {
  const cookieStore = await cookies()
  const auth = cookieStore.get('auth')

  console.log("auth: ", auth)

  if (auth === undefined) {
    redirect('/login')
  }

  return (
    <div>Index page</div>
  )
}
