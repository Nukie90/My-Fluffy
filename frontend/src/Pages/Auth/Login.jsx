import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';

import './../../App.css';


function Login({ currentPage, setCurrentPage }) {
    const [error, setError] = useState('');

    useEffect(() => {
        setCurrentPage('Login');
    }, [setCurrentPage]);

    const navigate = useNavigate();

    const handleLogin = (e) => {
        e.preventDefault();  // Prevent default form submission

        const formData = new FormData(e.target);
        const username = formData.get('username');
        const password = formData.get('password');

        // Simple validation for empty fields
        if (!username || !password) {
            setError('Both fields are required');
            return;
        }

        setError('');

        // Navigate to Home after login
        setCurrentPage('Home');
        navigate('/');
    };

    return (
        <div className="bg-white h-screen w-full flex justify-center">
            <div className='h-auto w-full bgct-light-purple flex sm:flex-col md:flex-row'>
                <div className='sm:w-full md:w-1/2 md:pr-10 bgct-light-purple colorct-dark-purple'>
                    <div className='flex items-center pt-8 sm:px-6 md:px-8 lg:px-12'>
                        <div className='flex justify-start w-1/2'>
                            <h1 className='text-2xl text-left'>{'<'}</h1>
                        </div>
                        <div className='flex justify-end w-1/2'>
                            <Link to='/signup'>
                                <p className='text-sm text-right'>Register</p>
                            </Link>
                        </div>
                    </div>
                    <div className='mt-10 flex flex-col sm:px-6 md:px-8 lg:px-12'>
                        <h1 className='font-bold text-4xl text-black md:text-center mb-4'>My Fluffy</h1>
                        <p className='text-md md:text-center'>Welcome back to My Fluffy!</p>
                    </div>
                </div>
                <div className='sm:absolute sm:bottom-0 md:flex md:top-0 md:right-0 flex-col justify-center items-center w-full mt-8 md:m-8 bg-white md:w-1/2 sm:rounded-t-3xl md:rounded-3xl py-8 sm:px-6 md:px-10 lg:px-12'>
                    <h1 className='font-bold text-2xl sm:text-left md:text-center md:mb-6 xl:mb-8'>Log In</h1>
                    <form className='flex flex-col w-full' onSubmit={handleLogin}>
                        <div className='flex flex-col mt-8'>
                            <label className='text-sm text-left'>Username</label>
                            <input
                                name="username"
                                type="username"
                                className='h-16 w-full rounded-full px-6 mt-8 bgct-light-yellow'
                            />
                        </div>
                        <div className='flex flex-col mt-8'>
                            <label className='text-sm text-left'>Password</label>
                            <input
                                name="password"
                                type="password"
                                className='h-16 w-full rounded-full px-6 mt-8 bgct-light-yellow'
                            />
                        </div>
                        {error && <p className="text-red-500 mt-4">{error}</p>}
                        <div className="flex justify-end my-10 mt-12">
                            <button
                                type="submit"
                                className="h-16 w-full flex justify-center items-center rounded-full 
                                    bgct-yellow colorct-dark-purple hover:text-white 
                                    transition-all duration-300 ease-out"
                                style={{ backgroundColor: 'var(--color-yellow)' }}
                                onMouseEnter={(e) => e.target.style.backgroundColor = 'var(--color-orange)'}
                                onMouseLeave={(e) => e.target.style.backgroundColor = 'var(--color-yellow)'}
                            >
                                <h2 className="text-lg font-bold">Log In</h2>
                            </button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    );
}

export default Login;
