"use client";

import React, { useState, useEffect } from 'react'

import {
    Button,
    Dialog,
    DialogHeader,
    DialogBody,
    DialogFooter,
    Checkbox,
    Card,
    List,
    ListItem,
    ListItemPrefix,
    Typography,
    Spinner,
    Input,
  } from "@material-tailwind/react";
import axios from 'axios';
import { GetURL } from '../utils/http-store';
import Cookies from 'js-cookie'


export default function CreateChat({ open, handleOpen  }) {
  const [ users, setUsers ] = useState([])
  const [ loading, setLoading ] = useState(true)
  const [selectedUsers, setSelectedUsers] = useState([])
  const [chatName, setChatName ] = useState("")
  const [ chatNameEntered, setChatNameEntered ] = useState(false)
  const [ invalidForm, setInValidForm ] = useState(false)

  useEffect(() => {
      const jwtToken = Cookies.get("auth")
      axios.get(GetURL('users/all-users'), {
          headers: {
              'Authorization': jwtToken
          }
      }).then(res => {
          if (res.status !== 200) {
              console.log(res.statusText)
              return
          }

          console.log(res)
          if (res.status === 200 && res.data) {
              setLoading((prev) => !prev)
              setUsers(res.data)
          }
      })
  }, [])

  useEffect(() => {
    if (users.length > 0 ) {
      setLoading(false)
    }
  }, [users])

  const handleSelect = (user) => {
    setSelectedUsers((prevSelected) => {
        if (prevSelected.some((u) => u === user.uuid)) {
            return prevSelected.filter((u) => u !== user.uuid);
        } else {
            return [...prevSelected, user.uuid];
        }
    });

    handleChatName(user.name, false)
  };

  const handleChatName = (name, fromInpField) => {
    if (fromInpField) {
      setChatNameEntered(true)
      setChatName(name)
      return
    }

    if (!chatNameEntered) {      
      setChatName((prev) => {
        if(chatName !== "") {
          return `${prev}, ${name}`
        }

        return name
      })
    }
  }

  const submitChat = () => {
    if (chatName.length > 0 && selectedUsers.length >= 1) {
      const jwtToken = Cookies.get("auth")
      const chatObj = {
        "participants": selectedUsers,
        "chat_name": chatName,
      }
  
      // ToDo: Lav post method + set up route:
      axios.post(GetURL('messages/create-chat'), JSON.stringify(chatObj) ,{
        headers: {
            'Authorization': jwtToken
        }
      }).then(res => {
          if (res.status !== 200) {
              console.log(res.statusText)
              return
          }
  
          console.log(res)
          if (res.status === 200 && res.data) {
            handleOpen()      
          }
      })
    } else {
      setInValidForm(true)
      setTimeout(() => setInValidForm(false), 15000)
    }
  }

  return (
    <>
      <Dialog open={open} handler={handleOpen}>
        <DialogHeader className='text-black'>Start a chat</DialogHeader>
          <DialogBody>
              { (loading ) ? (
                  <Spinner className='text-black' />
              ) : (
                <>
                  { invalidForm && (
                    <span className='bg-rose-500 font-bold mx-auto block mb-5 border rounded p-2 '>You need to fill out this form to create a chat.</span>
                  )}
                  <Input variant='static' label='Chat name' value={chatName} className='text-black' onInput={(e) => {
                    handleChatName(e.currentTarget.value, true)
                  }}/>
                  <List>
                    {users.map((user) => {
                      const isSelected = selectedUsers.some((u) => u === user.uuid);
        
                      return (
                          <ListItem key={user.uuid} className="p-0">
                              <label htmlFor={`user${user.name}-id`} className="flex w-full cursor-pointer items-center px-3 py-2">
                                  <ListItemPrefix className="mr-3">
                                      <Checkbox
                                          id={`user${user.name}-id`}
                                          checked={isSelected}
                                          onChange={() => handleSelect(user)}
                                          ripple={false}
                                          className="hover:before:opacity-0"
                                          containerProps={{
                                              className: "p-0",
                                          }}
                                      />
                                  </ListItemPrefix>
                                  <Typography color="blue-gray" className="font-medium text-black">
                                      {user.name}
                                  </Typography>
                              </label>
                          </ListItem>
                      );
                    })}
                  </List>
                </>
              )}
          </DialogBody>
        <DialogFooter>
          <Button
            variant="text"
            color="red"
            onClick={handleOpen}
            className="mr-1"
          >
            <span>Cancel</span>
          </Button>
          <Button variant="gradient" color="dark-grey" onClick={submitChat}>
            <span>Confirm</span>
          </Button>
        </DialogFooter>
      </Dialog>
    </>
  )
}
