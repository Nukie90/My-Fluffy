import { Link, useNavigate } from 'react-router-dom';
import './../../App.css';
import { useState, useEffect } from 'react';
import axios from 'axios';
import ProfileClick from './../../Components/Icons/profile_afterclick.svg';
import ProfileUnClick from './../../Components/Icons/profile_beforeclick.svg';
import SmallProfile from './../../Components/Icons/small_profile.svg';

import TopBar from './../../Components/TopBar';
import ShowPost from './../../Components/ShowPost'; // Importing the modal for showing posts

function Profile({ setCurrentPage }) {
    const [isLogged, setIsLogged] = useState(false);
    const [posts, setPosts] = useState([]);
    const [loading, setLoading] = useState(true);
    const [username, setUsername] = useState('Guest');
    const [petsCount, setPetsCount] = useState(0);
    const [lostCount, setLostCount] = useState(0);
    const [foundCount, setFoundCount] = useState(0);
    const [tab, setTab] = useState('finding');
    const [isModalVisible, setIsModalVisible] = useState(false);
    const [selectedPost, setSelectedPost] = useState(null);
    const navigate = useNavigate(); // Initialize useNavigate

    const fetchUser = async () => {
        try {
            const cookies = document.cookie.split('; ');
            const sessionCookie = cookies.find(cookie => cookie.startsWith('session='));
            if (!sessionCookie) {
                console.error('Session cookie not found');
                setIsLogged(false);
                return;
            }

            const response = await axios.get(`http://localhost:3000/api/v1/users/${sessionCookie.split('=')[1]}`, {
                withCredentials: true,
            });

            setUsername(response.data.username || 'Guest');
            setIsLogged(true);
        } catch (error) {
            console.error('Error fetching user:', error);
            setIsLogged(false);
        }
    };

    const fetchPosts = async () => {
        setLoading(true);
        try {
            const response = await axios.get('http://localhost:3000/api/v1/posts/user', {
                withCredentials: true,
            });
            const newPosts = response.data;

            const formattedPosts = newPosts.map((post) => ({
                id: post.id,
                title: post.title,
                content: post.content,
                status: post.status,
                picture: post.picture,
                reward: post.reward,
                username: post.username, 
                OwnerPicture: post.OwnerPicture, 
                isSaved: post.isSaved, 
                owner_id: post.owner_id, 
            }));

            setPosts(formattedPosts);
            // Updating counts
            setPetsCount(newPosts.length);
            setLostCount(newPosts.filter(post => post.status === 'Missing' || post.status === 'Pending').length);
            setFoundCount(newPosts.filter(post => post.status === 'Found').length);
        } catch (error) {
            console.error('Error fetching posts:', error);
        } finally {
            setLoading(false);
        }
    };

    const handleLogout = () => {
        // Clear the session cookie
        document.cookie = 'session=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';
        setIsLogged(false);
        navigate('/'); // Redirect to home
    };

    useEffect(() => {
        setCurrentPage('Profile');
        const cookies = document.cookie.split('; ');
        const sessionCookie = cookies.find((cookie) => cookie.startsWith('session='));

        if (sessionCookie) {
            fetchUser();
            fetchPosts();
        }
    }, [setCurrentPage]);

    const handlePostClick = (post) => {
        setSelectedPost(post);
        setIsModalVisible(true);
    };

    return (
        <div className='relative h-screen w-full bg-[#F9F9F9] flex flex-col'>
            <TopBar CurrentPage="Profile" setCurrentPage={setCurrentPage} />

            <div className='flex flex-col px-8 py-4 mt-20'>
                {isLogged ? (
                    <div className='flex flex-col items-start w-full mb-auto'>
                        <h1 className='text-3xl font-bold text-[#43337B] mb-4 mt-2'>Welcome, {username}!</h1>
                        <div className='flex items-center w-full mb-6'>
                            <div className='w-24 h-24 rounded-full bg-[#43337B] flex items-center justify-center relative ml-3'>
                                <span className='absolute bottom-0 right-[-10px] w-10 h-10 bg-[#C4C3E3] rounded-full flex items-center justify-center'>
                                    <img src={SmallProfile} alt='Profile Icon' className='w-6 h-6' />
                                </span>
                            </div>
                            <div className='flex space-x-8'>
                                <div className='flex flex-col items-center ml-12'>
                                    <span className='text-2xl font-semibold text-[#43337B]'>{petsCount}</span>
                                    <span className='text-[#43337B]'>Pets</span>
                                </div>
                                <div className='flex flex-col items-center'>
                                    <span className='text-2xl font-semibold text-[#43337B] ml-2'>{lostCount}</span>
                                    <span className='text-[#43337B] ml-2'>Lost</span>
                                </div>
                                <div className='flex flex-col items-center'>
                                    <span className='text-2xl font-semibold text-[#43337B] ml-1.5'>{foundCount}</span>
                                    <span className='text-[#43337B] ml-1.5'>Found</span>
                                </div>
                            </div>
                        </div>
                        <p className='text-sm mb-6 text-[#43337B] ml-1'>A cool kitten dad of 2<br />If you find any evidence of Lucy, contact me asap</p>
                        <div className='flex w-full justify-evenly mb-6'>
                            <button className='px-20 py-1 bg-[#F9F9F9] text-[#CF4B4B] rounded-lg border-2 border-[#CF4B4B] font-bold' onClick={handleLogout}>Log out</button>
                        </div>
                        <div className='w-full flex items-center justify-around mb-6 relative'>
                            <button
                                className={`flex items-center relative px-4 ${tab === 'finding' ? 'text-[#F1642E]' : 'text-[#A3B565]'}`}
                                onClick={() => setTab('finding')}
                            >
                                <img
                                    src={tab === 'finding' ? ProfileClick : ProfileUnClick}
                                    alt='Finding pets'
                                    className='w-6 h-6 mr-2'
                                />
                                Finding pets
                                <div className={`absolute bottom-[-2px] left-[-10px] right-[-10px] h-1 ${tab === 'finding' ? 'bg-[#FF7643]' : 'bg-transparent'}`}></div>
                            </button>
                            <button
                                className={`flex items-center relative px-4 ${tab === 'found' ? 'text-[#F1642E]' : 'text-[#A3B565]'}`}
                                onClick={() => setTab('found')}
                            >
                                <img
                                    src={tab === 'found' ? ProfileClick : ProfileUnClick}
                                    alt='Pets found'
                                    className='w-6 h-6 mr-2'
                                />
                                Pets found
                                <div className={`absolute bottom-[-2px] left-[-10px] right-[-10px] h-1 ${tab === 'found' ? 'bg-[#FF7643]' : 'bg-transparent'}`}></div>
                            </button>
                        </div>
                        <div className='grid grid-cols-2 gap-6 w-full'>
                            {loading ? (
                                <p>Loading posts...</p>
                            ) : (
                                posts
                                    .filter(post => (tab === 'finding' && (post.status === 'Missing' || post.status === 'Pending')) || (tab === 'found' && post.status === 'Found'))
                                    .map(post => (
                                        <div key={post.id} className='w-full h-auto rounded-lg overflow-hidden shadow-md cursor-pointer' onClick={() => handlePostClick(post)}>
                                            <img src={`data:image/jpeg;base64,${post.picture}`} alt={post.title} className='w-full h-40 object-cover' />
                                        </div>
                                    ))
                            )}
                        </div>
                    </div>
                ) : (
                    <div className='w-full h-full flex flex-col justify-center items-center bg-[#8DBA56]'>
                        <p className='text-white'>Please log in to view your profile</p>
                        <Link to='/login'>
                            <button
                                className='h-10 w-24 my-4 flex justify-center items-center bg-[#FF7643] rounded-lg text-white'
                                onClick={() => setCurrentPage('Login')}
                            >
                                Log In
                            </button>
                        </Link>
                    </div>
                )}
            </div>

            {isModalVisible && selectedPost && (
               <ShowPost post={selectedPost} setIsModalVisible={setIsModalVisible} username={username}/>
            )}
        </div>
    );
}

export default Profile;
