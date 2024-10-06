import { Link } from 'react-router-dom';
import './../../App.css';
import { useState, useEffect } from 'react';
import Posts from '../../Components/Posts';
import axios from 'axios';


function Profile({ setCurrentPage }) {
    const [isLogged, setIsLogged] = useState(false);
    const [posts, setPosts] = useState([]);
    const [loading, setLoading] = useState(true);

    const fetchPosts = async () => {

        setLoading(true);
        try {
          const response = await axios.get(`http://localhost:3000/api/v1/posts/user`,
            {
                withCredentials: true,
            }
          );
          const newPosts = response.data; 
          console.log('Fetched posts:', newPosts);
    
        const formattedPosts = newPosts.map(post => ({
            id: post.id,
            title: post.title,
            content: post.content,
            status: post.status,
            picture: post.picture,
            reward: post.reward
        }));

        setPosts(prevPosts => [...prevPosts, ...formattedPosts]); // Append new posts
          
        } catch (error) {
          console.error('Error fetching posts:', error);
        } finally {
          setLoading(false);
        }
    };

    useEffect(() => {
        setCurrentPage('Profile');
        const cookies = document.cookie.split('; ');
        const sessionCookie = cookies.find(cookie => cookie.startsWith('session='));
        
        if (sessionCookie) {
            setIsLogged(true);
            fetchPosts();
        }
    }, [setCurrentPage]);

    return (
        <div className='h-screen w-full bg-red-200 flex flex-col'>
            {isLogged ? (
                <div>
                    <h1>Profile</h1>
                    <p>Username: </p>
                    <p>Email: </p>
                    <p>Phone Number: </p>
                    <p>Address: </p>
                    <div className="flex sm:flex-col md:flex-row h-auto w-full pt-20 sm:px-4 md:px-6 lg:px-8">
                        <div className='h-auto pt-4 pb-20'>
                            <div className='h-auto'>
                            <h1 className='sm:text-xl md:text-2xl font-bold'>Your Posts</h1>
                            </div>

                            {loading && <p>Loading posts...</p>}

                            <Posts data={posts} />
                        </div>
                    </div>
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
