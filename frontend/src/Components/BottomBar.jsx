import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { useState, useEffect } from 'react';

import HomeIcon from './Icons/home_icon.svg';
import AddIcon from './Icons/add_icon.svg';
import RewardsIcon from './Icons/rewards_icon.svg';
import RewardsIconActive from './Icons/rewards_icon_active.svg';
import SavedIcon from './Icons/saved_icon.svg';
import SavedIconActive from './Icons/saved_icon_active.svg';

import AddPost from './AddPost';

function BottomBar({currentPage, setCurrentPage}) {
    const [isAddPostVisible, setIsAddPostVisible] = useState(false);
    
    const pages = [
        { page: 'Rewards', icon: RewardsIcon, activeIcon: RewardsIconActive, link: '/rewards' },
        { page: 'Home', icon: HomeIcon, activeIcon: AddIcon, link: '/' },
        { page: 'Saved', icon: SavedIcon, activeIcon: SavedIconActive, link: '/saved' },
    ];

    const handlePageClick = (page) => {
        if (currentPage === 'Home' && page === 'Home') {
            setIsAddPostVisible(true);
        } else {
            setIsAddPostVisible(false);
            setCurrentPage(page);
        }
    };

    return (
        <div 
            className='
                z-10
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
                            onClick={() => handlePageClick(page)}
                            className='
                                sm:mx-8
                                lg:mx-12
                                xl:mx-14
                                sm:w-12 sm:h-12
                                lg:w-16 lg:h-16
                                xl:w-20 xl:h-20
                                flex items-center justify-center
                                rounded-full'
                        >
                            <img 
                                src={currentPage === page ? activeIcon : icon} 
                                alt={page} 
                                className={`
                                    ${currentPage === page ? 'sm:w-9 sm:h-9 lg:w-12 lg:h-12 xl:w-14 xl:h-14' : 'sm:w-8 sm:h-8 lg:w-10 lg:h-10 xl:w-12 xl:h-12'}
                                    sm:hover:w-10 sm:hover:h-10 lg:hover:w-14 lg:hover:h-14 xl:hover:w-16 xl:hover:h-16
                                    transition-all duration-200 ease-out
                                `}
                            />
                        </button>
                    </Link>
                ))}
            </div>
            {isAddPostVisible && <AddPost setIsAddPostVisible={setIsAddPostVisible} />}
        </div>
    );
}

export default BottomBar;
