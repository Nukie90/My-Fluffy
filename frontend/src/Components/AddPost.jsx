import React, { useState } from 'react';

import './../App.css';
import XFindpetIcon from './Icons/x_findpet_icon.svg';

export default function AddPost({ setIsAddPostVisible }) {

    const handleExitButton = () => {
        setIsAddPostVisible(false);
    };

    return (
        <div className='fixed h-screen w-full flex flex-col justify-center items-center'>
            {/* Background overlay with a lower z-index */}
            <div className='fixed inset-0 bg-black bg-opacity-50 z-0' />
            
            {/* Main content area with a higher z-index */}
            <div 
                className='relative h-screen w-full flex flex-col items-center 
                rounded-xl shadow-md transition-all duration-300 ease-out bg-white'
            >  
                <div className='h-12 w-full z-10 mx-6 mt-6'>
                    <div className='absolute left-6 p-0 flex items-center'>
                        <img src={"/Profiles/default_pfp.jpg"} alt="Profile Pic" className='w-12 h-12 rounded-full mr-4' />
                        <h3 className='text-xl font-semibold colorct-dark-purple'>
                            Username
                        </h3>
                    </div>
                    <div className='absolute right-6'>
                        <button 
                            onClick={handleExitButton}
                            className='
                                w-12 h-12
                                flex items-center justify-center
                                rounded-full
                                text-white
                                transition-all duration-300 ease-out
                            '
                        >
                            <img src={XFindpetIcon} alt='exit_icon' className='w-10 h-10' />
                        </button>
                    </div>
                </div>
                <div className='h-full w-full flex flex-col items-center mt-4'>
                    <form className='flex w-full flex-col items-center px-8'>
                        <div className='w-full flex flex-col items-center'>
                            <input
                                name='title'
                                type='text'
                                placeholder='Title'
                                className={'h-10 w-full font-semibold text-2xl bg-transparent colorct-dark-purple placeholder-green'}
                            />
                        </div>
                        <div className='w-full flex flex-col items-center'>
                            <input
                                name='reward'
                                type='text'
                                placeholder='Reward'
                                className={'h-10 w-full font-medium text-lg bg-transparent colroct-dark-purple placeholder-green'}
                            />
                        </div>
                        <div className='w-full flex flex-col items-center'>
                            <textarea
                                name='content'
                                type='text'
                                placeholder='Describe your fluffy...'
                                className={'sm:min-h-24 h-44 md:min-h-28 w-full text-lg bg-transparent colroct-dark-purple placeholder-green'}
                            />
                        </div>
                        <div className='w-full flex items-center justify-center mt-4'>
                            <input
                                name='picture'
                                type='file'
                                accept='image/*'
                                className={'h-10 w-full font-semibold text-md colroct-dark-purple placeholder-green'}
                            />
                        </div>
                        <div className='w-full flex flex-col items-center mt-4'>
                            <button
                                type='submit'
                                className='h-12 w-4/5 flex justify-center items-center rounded-full font-semibold text-lg
                                    colorct-light-yellow hover:text-white 
                                    transition-all duration-300 ease-out bgct-orange'
                                onMouseEnter={(e) => e.target.style.backgroundColor = 'var(--color-dark-purple)'}
                                onMouseLeave={(e) => e.target.style.backgroundColor = 'var(--color-orange)'}
                            >
                                Post
                            </button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    );
}
