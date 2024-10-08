import React, { useState, useEffect } from 'react';
import './../App.css';
import XFindpetIcon from './Icons/x_findpet_icon.svg';
import axios from 'axios';

export default function AddPost({ setIsAddPostVisible }) {
    const [isLogged, setIsLogged] = useState(false);
    const [userId, setUserId] = useState(null);
    const [formData, setFormData] = useState({
        Title: '',
        Reward: '',  // Start as an empty string to allow input changes
        Content: '',
        File: null,
        OwnerID: userId,
    });
    const [error, setError] = useState('');

    const handleInputChange = (e) => {
        const { name, value } = e.target;

        setFormData((prevData) => ({
            ...prevData,
            [name]: value,  // Store Reward as a string
        }));
    };

    const handleFileChange = (e) => {
        setFormData((prevData) => ({
            ...prevData,
            File: e.target.files[0],
        }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
    
        // Validate form inputs
        if (!formData.Title || !formData.Content || !formData.File) {
            setError('Please fill in all fields correctly.');
            return;
        }
    
        // Log the reward value and formData before sending it to the backend
        console.log('Reward value:', formData.Reward); // Log reward value
    
        const postData = new FormData();
        postData.append('title', formData.Title);
        postData.append('content', formData.Content);
        postData.append('file', formData.File);
        postData.append('reward', formData.Reward); // Ensure it's sent as a string
        postData.append('OwnerID', userId);
    
        // Log entire FormData for debugging
        postData.forEach((value, key) => {
            console.log(`${key}: ${value}, ${typeof(value)}`);
        });
    
        try {
            const response = await axios.post('http://localhost:3000/api/v1/posts', postData, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                },
            });
            // console.log('Post created:', response.data);
            window.location.reload();
            setIsAddPostVisible(false); // Close the modal after submission
        } catch (error) {
            console.error('Error creating post:', error.response?.data || error.message);
            setError(error.response?.data?.error || 'An error occurred while creating the post.');
        }
    };


    const handleExitButton = () => {
        setIsAddPostVisible(false);
    };

    useEffect(() => {
        const cookies = document.cookie.split('; ');
        const sessionCookie = cookies.find(cookie => cookie.startsWith('session='));

        if (sessionCookie) {
            const sessionToken = sessionCookie.split('=')[1];
            if (sessionToken) {
                setIsLogged(true);
                setUserId(sessionToken);
            }
        } else {
            console.log('No session cookie found.');
        }
    }, []);

    return (
        <div className='fixed h-screen w-full flex flex-col justify-center items-center'>
            <div className='fixed inset-0 bg-black bg-opacity-50 z-0' />
            <div className='relative h-screen w-full flex flex-col items-center rounded-xl shadow-md transition-all duration-300 ease-out bg-white'>  
                <div className='h-12 w-full z-10 mx-6 mt-6'>
                    <div className='absolute left-6 p-0 flex items-center'>
                        <img src={"/Profiles/default_pfp.jpg"} alt="Profile Pic" className='w-12 h-12 rounded-full mr-4' />
                        <h3 className='text-xl font-semibold colorct-dark-purple'>Username</h3>
                    </div>
                    <div className='absolute right-6'>
                        <button onClick={handleExitButton} className='w-12 h-12 flex items-center justify-center rounded-full text-white transition-all duration-300 ease-out'>
                            <img src={XFindpetIcon} alt='exit_icon' className='w-10 h-10' />
                        </button>
                    </div>
                </div>
                <div className='h-full w-full flex flex-col items-center mt-4'>
                    <form className='flex w-full flex-col items-center px-8' onSubmit={handleSubmit}>
                        <div className='w-full flex flex-col items-center'>
                            <input
                                name='Title'
                                type='text'
                                placeholder='Title'
                                className='h-10 w-full font-semibold text-2xl bg-transparent colorct-dark-purple placeholder-green'
                                onChange={handleInputChange}
                            />
                        </div>
                        <div className='w-full flex flex-col items-center'>
                            <input
                                name='Reward'
                                type='text'
                                placeholder='Reward'
                                className='h-10 w-full font-medium text-lg bg-transparent colorct-dark-purple placeholder-green'
                                onChange={handleInputChange}
                            />
                        </div>
                        <div className='w-full flex flex-col items-center'>
                            <textarea
                                name='Content'
                                placeholder='Describe your fluffy...'
                                className='sm:min-h-24 h-44 md:min-h-28 w-full text-lg bg-transparent colorct-dark-purple placeholder-green'
                                onChange={handleInputChange}
                            />
                        </div>
                        <div className='w-full flex items-center justify-center mt-4'>
                            <input
                                name='File'
                                type='file'
                                accept='image/*'
                                className='h-10 w-full font-semibold text-md colorct-dark-purple placeholder-green'
                                onChange={handleFileChange}
                            />
                        </div>
                        <div className='w-full flex flex-col items-center mt-4'>
                            <button
                                type='submit'
                                className='h-12 w-4/5 flex justify-center items-center rounded-full font-semibold text-lg colorct-light-yellow hover:text-white transition-all duration-300 ease-out bgct-orange'
                            >
                                Post
                            </button>
                        </div>
                        {error && <p className='text-red-500'>{error}</p>} {/* Display error messages */}
                    </form>
                </div>
            </div>
        </div>
    );
}
