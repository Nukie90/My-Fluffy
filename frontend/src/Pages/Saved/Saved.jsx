import './../../App.css';

import { useState, useEffect } from 'react'; 

function Saved({currentPage, setCurrentPage}) {
    useEffect(() => {
        setCurrentPage('Saved');
    }, [setCurrentPage]);
    
    return (
        <div className='h-auto bg-red-200'>
            this is the Saved page
        </div>
    );
}

export default Saved;