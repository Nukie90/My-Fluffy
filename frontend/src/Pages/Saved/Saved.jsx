import { Link } from 'react-router-dom';
import './../../App.css';
import Posts from '../../Components/Posts';
import axios from 'axios';

import { useState, useEffect } from 'react'; 

function Saved({setCurrentPage}) {
    const [isLogged, setIsLogged] = useState(false);
    const [savedPosts, setSavedPosts] = useState([]);
    const [loading, setLoading] = useState(true);

    const fetchSavedPosts = async () => {

        setLoading(true);
        try {
          const response = await axios.get(`http://localhost:3000/api/v1/saved_posts/`,
            {
                withCredentials: true,
            }
          );
          const savedPosts = response.data; 
          console.log('Fetched savedPosts:', savedPosts);
    
        const formattedPosts = savedPosts.map(post => ({
            id: post.id,
            title: post.title,
            username: post.username,
            owner_id: post.owner_id,
            found_id: post.found_id,
            content: post.content,
            status: post.status,
            picture: post.picture,
            reward: post.reward
        }));

        setSavedPosts(prevPosts => [...prevPosts, ...formattedPosts]);
          
        } catch (error) {
          console.error('Error fetching savedPosts:', error);
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
            fetchSavedPosts();
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

                            <Posts data={savedPosts} savedPosts={savedPosts} />
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