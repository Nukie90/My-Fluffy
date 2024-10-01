import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { useState } from 'react';

import HomeIcon from './Icons/home_icon.svg';
import AddIcon from './Icons/add_icon.svg';
import RewardsIcon from './Icons/rewards_icon.svg';
import RewardsIconActive from './Icons/rewards_icon_active.svg';
import SavedIcon from './Icons/saved_icon.svg';
import SavedIconActive from './Icons/saved_icon_active.svg';

function BottomBar({currentPage, setCurrentPage}) {
    
    const pages = [
        { page: 'Rewards', icon: RewardsIcon, activeIcon: RewardsIconActive, link: '/rewards' },
        { page: 'Home', icon: HomeIcon, activeIcon: AddIcon, link: '/' },
        { page: 'Saved', icon: SavedIcon, activeIcon: SavedIconActive, link: '/saved' },
    ];

    return (
        <div 
            className='
                z-100
                fixed bottom-10 left-0 
                w-full h-auto
                flex items-center justify-center'
        >
            <div
                className='
                w-auto h-auto 
                bgct-dark-purple 
                rounded-full
                flex items-center justify-center
                mx-20'
            >
                {pages.map(({ page, icon, activeIcon, link }, index) => (
                    <Link to={link} key={index}> 
                        <button
                            onClick={() => setCurrentPage(page)}
                            className='
                                sm:mx-8
                                lg:mx-20
                                xl:mx-32
                                sm:w-12 sm:h-12
                                lg:w-16 lg:h-16
                                xl:w-20 xl:h-20
                                flex items-center justify-center
                                rounded-full'
                        >
                            <img 
                                src={currentPage === page ? activeIcon : icon} 
                                alt={page} 
                                className='sm:w-8 sm:h-8 lg:w-18 lg:h-18' 
                            />
                        </button>
                    </Link>
                ))}
            </div>
        </div>
    );
}

export default BottomBar;
