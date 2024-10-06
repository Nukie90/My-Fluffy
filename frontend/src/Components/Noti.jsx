import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './../App.css';

export default function Noti({ noti, setIsNotiVisible }) {
    const [notifications, setNotifications] = useState(noti);

    async function handleDeleteButton(id) {
        try {
            await axios.delete(`http://localhost:3000/api/v1/notifications/${id}`, {
                withCredentials: true,
            });

            setNotifications(prevNoti => prevNoti.filter(notification => notification.id !== id));

        } catch (error) {
            console.error('Error deleting notification:', error);
        }
    }

    const handleExitButton = () => {
        setIsNotiVisible(false);
    };

    const handleRewardButton = (id) => {
        // Placeholder for reward button logic
    };

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
                {notifications.map((notification) => (
                    <div key={notification.id} 
                        className='h-32 w-11/12 m-4 mb-1 p-4 flex flex-col items-center justify-center rounded-xl'
                        style={{ backgroundColor: 'var(--color-light-purple)' }} // Use the same color for all notifications
                    >
                        <div className='h-auto w-full flex justify-between items-center'>
                            <div className='h-auto w-1/2 flex justify-between items-center'>
                                <div className='flex items-center'>
                                    {/* Placeholder for OwnerPicture */}
                                    <img src={notification.UserPicture || '/Profiles/default_pfp.jpg'} alt={notification.UserPicture} className='w-16 h-16 rounded-full' />
                                </div>
                            </div>
                            <div className='h-full w-full w-1/2 flex flex-col justify-center text-left'>
                                <div className='w-full'>
                                    <p className='text-md colorct-dark-purple'>{notification.message}</p>
                                </div>
                            </div>
                        </div>
                        <div className='h-auto w-full flex justify-end items-center'>
                            <div className='w-full h-1/4 flex justify-end'>
                                <button onClick={() => handleDeleteButton(notification.id)}>
                                    <p className='text-xs mx-1 py-1 px-2 colorct-dark-purple rounded-lg'
                                        style={{ border: '1px solid var(--color-dark-purple)' }}
                                    >
                                        Delete
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
