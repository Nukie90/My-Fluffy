import React from 'react';
import { useState, useEffect } from 'react';

import './../../App';

function Signup({currentPage, setCurrentPage}) {
    useEffect(() => {
        setCurrentPage('Signup');
    }, [setCurrentPage]);

    return (
        <div 
        className="
            bgct-yellow h-screen w-full
        ">
            <div className='h-auto bg-red-200'>
            this is the Signup page
            </div>
        </div>
    );
}

export default Signup;