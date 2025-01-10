'use client';

import React from 'react'
import FormField from '../components/FormField'
import axios from 'axios';
import { cookies } from 'next/headers';

export default function Login() {

    async function login(e) {
        e.preventDefault()
        
        const name = document.querySelector('#nameInp').value
        const password = document.querySelector('#passwordInp').value

        const user = {
            name: name,
            password: password
        }

        const config = {
            headers: {
                'Content-type': 'application/json'
            }
        }
        let res = await axios.post('http://127.0.0.1:80/login', user, config)
        if (res.status === 200) {
            await cookies().set('accessToken', res.data, {
                httpOnly: true,
                maxAge: 24 * 60 * 60,
                sameSite: "strict"
            })
        }

    }

    return (
        <div className='mx-auto mt-5 w-96'>
            <h2>Login Page</h2>
            <FormField submition={login} page="login">
                
                <label className='col-span-2' htmlFor="name">Name  </label>
                <input className='col-span-3' type="text" name="name" id='nameInp'/>

            
                <label className='col-span-2' htmlFor="password">Password  </label>
                <input className='col-span-3' type="password" name="password" id='passwordInp'/>
            
            </FormField>
        </div>
    )
}
