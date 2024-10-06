import { Link } from 'react-router-dom';
import './../../App.css';
import Posts from '../../Components/Posts';
import axios from 'axios';

import { useState, useEffect } from 'react'; 

function Saved({currentPage, setCurrentPage}) {
    const [isLogged, setIsLogged] = useState(false);
    const [posts, setPosts] = useState([]);
    const [loading, setLoading] = useState(true);

    const fetchPosts = async () => {

        setLoading(true);
        try {
          const response = await axios.get(`http://localhost:3000/api/v1/saved_posts/`,
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
        setCurrentPage('Saved');
        const cookies = document.cookie.split('; ');
        const sessionCookie = cookies.find(cookie => cookie.startsWith('session='));
        
        if (sessionCookie) {
            setIsLogged(true);
            fetchPosts();
        }
    }, [setCurrentPage]);
    
    return (
        <div className='h-screen w-full bg-red-200 flex flex-col justify-center items-center'>
            {isLogged ? (
                <div>
                    <div className="flex sm:flex-col md:flex-row h-auto w-full pt-20 sm:px-4 md:px-6 lg:px-8">
                        <div className='h-auto pt-4 pb-20'>
                            <div className='h-auto'>
                            <h1 className='sm:text-xl md:text-2xl font-bold'>Saved Posts</h1>
                            </div>

                            {loading && <p>Loading posts...</p>}

                            <Posts data={posts} />
                        </div>
                    </div>
                </div>
            ) : (
                <div className='w-full h-full flex flex-col justify-center items-center bgct-green'>
                    <p>Please log in to saved posts</p>
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

export default Saved;