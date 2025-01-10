'use client';

import React from 'react'
import FormField from '../components/FormField'
import axios from 'axios';

export default function Signup() {

    async function signUp(e) {
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
        await axios.post('http://127.0.0.1:80/signup', user, config)

    }

    return (
        <div className='mx-auto mt-5 w-96'>
            <h2>Sign up</h2>
            <FormField submition={signUp} page="signup">
                
                <label className='col-span-2' htmlFor="name">Name  </label>
                <input className='col-span-3' type="text" name="name" id='nameInp'/>

            
                <label className='col-span-2' htmlFor="password">Password  </label>
                <input className='col-span-3' type="password" name="password" id='passwordInp'/>
            
            </FormField>
        </div>
    )
}
