import './../../App.css';

import { useState, useEffect } from 'react'; 

function Rewards({currentPage, setCurrentPage}) {
    useEffect(() => {
        setCurrentPage('Rewards');
    }, [setCurrentPage]);

    return (
        <div className='h-auto bg-red-200'>
            this is the rewards page
        </div>
    );
}

export default Rewards;