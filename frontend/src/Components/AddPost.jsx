import React, {useState} from 'react';

import './../App.css';

export default function AddPost({setIsAddPostVisible}) {


    return (
        <div 
            className='absolute 
                min-h-screen h-auto w-full 
                flex flex-col justify-center items-center 
                bg-white
                shadow-md
                transition-all duration-300 ease-out
        '>
            <div className='absolute top-0 right-0'>
                <button 
                    onClick={() => setIsAddPostVisible(false)}
                    className='
                        w-8 h-8
                        flex items-center justify-center
                        rounded-full
                        bgct-dark-purple
                        text-white
                        transition-all duration-300 ease-out
                    '
                >
                    <h1 className='text-lg font-bold'>X</h1>
                </button>
            </div>

            <h1 className='text-2xl font-bold'>Add a Post</h1>
        </div>
    )
}