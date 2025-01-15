import React from 'react'

export default function ChatBubble({ msg, name, sendAt, receivedMsg }) {
    

    if (receivedMsg) {
        return (
                <>
                    <div className="flex items-start gap-2.5 my-5 ">
                        <div className="flex flex-col gap-1 w-full max-w-[320px]">
                            <div className="flex items-center space-x-2 rtl:space-x-reverse">
                                <span className="text-sm font-semibold text-gray-900 dark:text-white">{ name }</span>
                                <span className="text-sm font-normal text-gray-500 dark:text-gray-400">{ sendAt }</span>
                            </div>
                            <div className="flex flex-col leading-1.5 p-4 border-gray-200 bg-gray-100 rounded-xl dark:bg-gray-700">
                                <p className="text-sm font-normal text-gray-900 dark:text-white"> { msg }</p>
                            </div>
                        </div>
                    </div>
                </>
        )
    } else {
        return (
            <>
                <div className="flex items-start gap-2.5 my-5 justify-end">
                    <div className="flex flex-col gap-1 w-full max-w-[320px]">
                        <div className="flex items-center space-x-2 rtl:space-x-reverse">
                            <span className="text-sm font-semibold text-gray-900 dark:text-white">{ name }</span>
                            <span className="text-sm font-normal text-gray-500 dark:text-gray-400">{ sendAt }</span>
                        </div>
                        <div className="flex flex-col leading-1.5 p-4 border-gray-200 bg-gray-100 rounded-xl dark:bg-gray-700">
                            <p className="text-sm font-normal text-gray-900 dark:text-white"> { msg }</p>
                        </div>
                    </div>
                </div>
            </>
        )
    }
}
