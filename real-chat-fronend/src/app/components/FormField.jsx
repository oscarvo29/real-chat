'use client';
import React from 'react'

export default function FormField(props) {

  let page = props.page

  let loginpage = false
  let signuppage = false

  switch (page) {
    case "login":


  }
  


  return (
    <form className='grid grid-cols-5 gap-4 mx-auto mt-5'  onSubmit={props.submition}>
    
      {props.children}


      { page == "login" ? (
        <div className='col-span-5 grid grid-cols-5 gap-4'>
          <button className='col-span-2' type="submit">Log In</button>
          <a className='col-span-2 col-start-3' href="/signup">Don't have an account? Sign up!</a>
         </div>
      ) : (
        <div className='col-span-5 grid grid-cols-5 gap-4'>
          <button className='col-span-2' type="submit">Sign Up</button>
          <a className='col-span-2 col-start-3' href="/login">All ready have an account? Log in!</a>
         </div>
      )}

      
     
   
  
       
</form>

  )
}
