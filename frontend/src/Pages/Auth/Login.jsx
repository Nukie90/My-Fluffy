import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import axios from 'axios';

import './../../App.css';
import MyFluffyLogo from './../../Components/Icons/MyFluffy_logo.svg';

function Login({currentPage, setCurrentPage}) {
    const [error, setError] = useState('');

    useEffect(() => {
        setCurrentPage('Login');
    }, [setCurrentPage]);

    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();  // Prevent default form submission

        const formData = new FormData(e.target);
        const username = formData.get('username');
        const password = formData.get('password');

        if (!username || !password) {
            setError('Both fields are required');
            return;
        }

        setError('');
        
        try {
            const response = await axios.post(`http://localhost:3000/api/v1/users/login`, {
                username,
                password
            }, {
                withCredentials: true
            });

            if (response.data) {
                setCurrentPage('Home');
                navigate('/');
            } else {
                setError(response.data.message || 'Login failed');
            }
        } catch (err) {
            console.error(err);
            setError('An error occurred. Please try again.');
        }
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
                            <Link to='/signup'>
                                <p className='text-sm md:text-md text-right colorct-dark-purple hover:colorct-orange'>Sign Up</p>
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
                                Log In
                            </button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    );
}

export default Login;