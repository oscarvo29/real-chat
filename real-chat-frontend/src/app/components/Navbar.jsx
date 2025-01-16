'use client';

import React from 'react'
import Cookies from 'js-cookie'
import { redirect } from 'next/dist/server/api-utils';

export default function Navbar() {
  const jwtToken = Cookies.get("auth")

  const hanldeLogOut = (e) => {
    e.preventDefault();
    if (jwtToken !== "") {
      Cookies.remove("auth")
      window.location.href = '/login'
    }

  }

  return (
        <div className="col-span-5 flex justify-between">
          <h2 className=' py-2 mt-2  text-xl'>Index Page</h2>
          <button onClick={hanldeLogOut} className='inline-block mt-2 rounded border-2 border-primary-100 px-6 py-2 text-xs font-medium uppercase leading-normal text-primary-700 transition duration-150 ease-in-out hover:border-primary-accent-200 hover:bg-secondary-50/50 focus:border-primary-accent-200 focus:bg-secondary-50/50 focus:outline-none focus:ring-0 active:border-primary-accent-200 motion-reduce:transition-none dark:border-primary-400 dark:text-primary-300 dark:hover:bg-blue-950 dark:focus:bg-blue-950'>Log out</button>
        </div>
  )
}
