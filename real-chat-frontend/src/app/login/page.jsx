'use client';

import React from 'react'
import FormField from '../components/FormField'
import axios from 'axios';
import setCookie from '../utils/cookies';
import { redirect } from 'next/navigation'
import { GetURL } from '../utils/http-store';
import { Button } from "@material-tailwind/react";

export default function Login() {

    async function login(e) {
        e.preventDefault()
        console.log("from login", GetURL('auth/login'))
        
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
        // let res = await axios.post('http://127.0.0.1:80/auth/login', user, config)
        let res = await axios.post(GetURL('auth/login'), user, config)
        if (res.status === 200) {
            setCookie('auth', res.data, 1)
            redirect('/')
        }

    }

    return (
        <div className='mx-auto mt-5 w-96'>
            <h2>Login Page</h2>
            <FormField submition={login} page="login">
                
                <label className='col-span-2' htmlFor="name">Name  </label>
                <input className='col-span-3 text-black' type="text" name="name" id='nameInp'/>

            
                <label className='col-span-2' htmlFor="password">Password  </label>
                <input className='col-span-3 text-black' type="password" name="password" id='passwordInp'/>
            
            </FormField>
        </div>
    )
}
