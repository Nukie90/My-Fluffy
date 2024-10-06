import React, { useEffect, useState } from 'react';

import './../App.css';

export default function Noti({ noti, setIsNotiVisible, setUnReadNoti }) {

    function handleReadButton(id) {
        noti = noti.map((notification) => {
            if (notification.id === id) {
                notification.isRead = true;
            }
            return notification;
        }
    );}

    const handleExitButton = () => {
        setIsNotiVisible(false);
    };

    const handleRewardButton = () => {
    };

    useEffect(() => {
        setUnReadNoti(noti.filter(noti => !noti.isRead).length);
    }, [noti, setUnReadNoti]);

    return (
        <div className='fixed top-0 right-0 h-full w-full flex flex-col justify-center items-center z-11'>
            {/* Background overlay with a lower z-index */}
            <div className='fixed inset-0 bg-black bg-opacity-50 z-11'>
                <button className='h-full w-full' onClick={handleExitButton} />
            </div>
            
            {/* Main content area with a higher z-index */}
            <div 
                className='relative h-4/5 w-11/12 flex flex-col items-center z-100 bgct-light-yellow
                rounded-xl shadow-lg transition-all duration-300 ease-out'
            >  
                {noti.map((notification) => (
                    <div key={notification.id} 
                        className='h-32 w-11/12 m-4 mb-1 p-4 flex flex-col items-center justify-center rounded-xl'
                        style={{ backgroundColor: notification.isRead ? 'var(--color-light-yellow)' : 'var(--color-light-purple)' }}
                    >
                        <div className='h-auto w-full flex justify-between items-center'>
                            <div className='h-auto w-1/2 flex justify-between items-center'>
                                <div className='flex items-center'>
                                    {/* Placeholder for OwnerPicture; you might want to adjust this based on the response */}
                                    <img src={notification.UserPicture || '/Profiles/default_pfp.jpg'} alt={notification.UserPicture} className='w-16 h-16 rounded-full' />
                                </div>
                            </div>
                            <div className='h-full w-full w-1/2 flex flex-col justify-center text-left'>
                                <div className='w-full mb-1'>
                                    <h1 className='text-lg font-semibold colorct-dark-purple'>{notification.username || 'Unknown'}</h1>
                                </div>
                                <div className='w-full'>
                                    <p className='text-md colorct-dark-purple'>{notification.noti}</p>
                                </div>
                            </div>
                        </div>
                        <div className='h-auto w-full flex justify-end items-center'>
                            <div className='w-full h-1/4 flex justify-end '>
                                <button onClick={() => handleReadButton(notification.id)}>
                                    <p className='text-xs mx-1 py-1 px-2 colorct-dark-purple rounded-lg'
                                        style={{ border: '1px solid var(--color-dark-purple)' }}
                                    >
                                        Mark as Read
                                    </p>
                                </button>
                                <button onClick={() => handleRewardButton(notification.id)}>
                                    <p className='text-xs mx-1 py-1 px-2 text-white bgct-orange rounded-lg'
                                        style={{ border: '1px solid var(--color-orange)' }}
                                    >
                                        Reward
                                    </p>
                                </button>
                            </div>
                        </div>
                    </div>  
                ))}

            </div>
        </div>
    );
}
