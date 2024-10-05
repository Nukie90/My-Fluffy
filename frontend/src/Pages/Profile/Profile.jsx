import { Link } from 'react-router-dom';
import './../../App.css';

import { useState, useEffect } from 'react'; 

function Profile({currentPage, setCurrentPage}) {
    useEffect(() => {
        setCurrentPage('Profile');
    }, [setCurrentPage]);

    let isLogged = false;

    return (
        <div className='h-screen w-full bg-red-200 flex flex-column justify-center items-center'>
            {isLogged ? (
                <div>
                    <h1>Profile</h1>
                    <p>Username: </p>
                    <p>Email: </p>
                    <p>Phone Number: </p>
                    <p>Address: </p>
                </div>
            ) : (
                <div className='w-full h-full flex flex-col justify-center items-center bgct-green'>
                    <p>Please log in to view your profile</p>
                    <Link to='/login'>
                        <button 
                            className='h-10 w-24 my-4 flex justify-center items-center bgct-yellow rounded-lg'
                            onClick={() => setCurrentPage('Login')}
                        >
                            Log In
                        </button>
                    </Link>
                </div>
            )}
        </div>
    );
}

export default Profile;