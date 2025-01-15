"use client";

import React from 'react'

export default function UserButtons({ user, index, handler }) {
    const handeUserClick = (e) => {
        e.preventDefault()
        handler(index)
    }

    return (
        <li className='my-5 col-span-2' key={index}>
          <button
            type="button"
            onClick={handeUserClick}
            className="inline-block rounded border-2 border-primary-100 px-6 pb-[6px] w-full pt-2 text-xs font-medium uppercase leading-normal text-primary-700 transition duration-150 ease-in-out hover:border-primary-accent-200 hover:bg-secondary-50/50 focus:border-primary-accent-200 focus:bg-secondary-50/50 focus:outline-none focus:ring-0 active:border-primary-accent-200 motion-reduce:transition-none dark:border-primary-400 dark:text-primary-300 dark:hover:bg-blue-950 dark:focus:bg-blue-950"
            data-twe-ripple-init>
              {user.name}
          </button>
        </li>
      )
}
