import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';

import './../../App.css';
import MyFluffyLogo from './../../Components/Icons/MyFluffy_logo.svg';

function Signup({currentPage, setCurrentPage}) {
    const [error, setError] = useState('');

    useEffect(() => {
        setCurrentPage('Signup');
    }, [setCurrentPage]);

    const navigate = useNavigate();

    const handleSignup = (e) => {
        e.preventDefault();  // Prevent default form submission

        const formData = new FormData(e.target);
        const username = formData.get('username');
        const password = formData.get('password');

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
                
                <div className='sm:w-full md:w-1/2 md:pr-10 flex flex-col md:justify-between md:h-screen'>
                    <div className='flex items-center pt-8 sm:px-6 md:px-8 lg:px-12'>
                        <div className='flex justify-start w-1/2'>
                            <Link to='/'>
                                <h1 className='text-3xl text-left colorct-dark-purple hover:colorct-orange'>{'<'}</h1>
                            </Link>
                        </div>
                        <div className='flex justify-end w-1/2'>
                            <Link to='/login'>
                                <p className='text-sm md:text-md text-right colorct-dark-purple hover:colorct-orange'>Log in</p>
                            </Link>
                        </div>
                    </div>

                    <div className='flex flex-col sm:px-6 md:px-8 lg:px-12 flex-grow justify-center'>
                        <img src={MyFluffyLogo} alt='My Fluffy Logo' className='sm:w-12 md:w-1/4 mx-auto sm:mb-2 md:my-10' />
                        <h1 className='font-bold sm:text-3xl md:text-4xl text-black text-center mb-4'>My Fluffy</h1>
                        <p className='text-md text-center'>Welcome to My Fluffy!</p>
                    </div>
                </div>

                <div className='sm:absolute sm:bottom-0 md:flex md:top-0 md:right-0 flex-col justify-center items-center w-full mt-8 md:m-8 bg-white md:w-1/2 sm:rounded-t-3xl md:rounded-3xl sm:pt-8 md:py-8 sm:px-6 md:px-10 lg:px-24'>
                    <h1 className='font-bold text-2xl sm:text-left md:text-center md:mb-6 xl:mb-8'>Sign up</h1>
                    <form className='flex flex-col w-full' onSubmit={handleSignup}>
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
                        {error ? (
                            <p className="text-red-500 mt-4 h-4">{error}</p>
                            ) : (
                            <div className="mt-4 h-4"></div>
                            )}
                        <div className="flex justify-end sm:my-8 md:my-10 mt-12">
                            <button
                                type="submit"
                                className="h-16 w-full flex justify-center items-center rounded-full font-bold text-lg
                                    bgct-yellow colorct-dark-purple hover:text-white 
                                    transition-all duration-300 ease-out"
                                style={{ backgroundColor: 'var(--color-yellow)' }}
                                onMouseEnter={(e) => e.target.style.backgroundColor = 'var(--color-orange)'}
                                onMouseLeave={(e) => e.target.style.backgroundColor = 'var(--color-yellow)'}
                            >
                                Sign Up
                            </button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    );
}

export default Signup;