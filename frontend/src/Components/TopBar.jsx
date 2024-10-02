import './../App.css';
import SearchIcon from './Icons/search_icon.svg';
import NotiIcon from './Icons/noti_icon.svg';
import DMIcon from './Icons/dm_icon.svg';
import DefaultPFP from './Profiles/default_pfp.jpg';

import { useState } from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';

function TopBar({CurrentPage, setCurrentPage}) {

    let pfp = DefaultPFP;
    let dmNoti = '';

    dmNoti = '2';
    const toprightIcons = [SearchIcon, NotiIcon, DMIcon];

    return (
        <div className="fixed h-20 w-full flex text-center bg-white shadow-md items-center justify-between sm:pt-6 md:pt-0 sm:px-4 md:px-6 lg:px-8"
            style={{ color: 'var(--color-dark-purple)' }}
        >
            <div className='flex items-center justify-center rounded-full'>
                <Link to={'/profile'}>
                    <button 
                        onClick={() => setCurrentPage('Profile')}
                        className='
                            sm:w-9 h-9 
                            sm:hover:w-10 h-10 
                            lg:w-10 h-10 
                            lg:hover:w-12 h-12 
                            xl:w-12 h-12
                            xl:hover:w-14 h-14
                            rounded-full transition-all duration-300'
                    >
                        <img src={pfp} alt="Profile Pic" key={"Profile Pic"} className='rounded-full' />
                    </button>
                </Link>
            </div>
            <div className='flex justify-end py-5'>
            {toprightIcons.map((icon, index) => (
                <button
                    className='
                        relative flex items-center justify-center
                        rounded-full transition-all duration-300
                        sm:ml-4 w-6 h-6
                        lg:ml-6 w-8 h-8
                        xl:w-9 h-9
                    '
                    key={index}
                >
                    <img src={icon} alt={icon} key={icon}
                        className="w-full h-full" 
                    />
                    {icon === DMIcon && dmNoti !== '' && (
                        <div className='
                            absolute top-0 right-0
                            flex items-center justify-center
                            sm:w-4 sm:h-4
                            xl:w-5 xl:h-5
                            bg-orange-500 text-white
                            rounded-full text-xs
                            transform translate-x-1 translate-y-5
                        '>
                            2
                        </div>
                    )}
                </button>
            ))}

            </div>
        </div>
    )
}

export default TopBar;