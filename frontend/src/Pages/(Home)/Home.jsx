import React, { useState, useEffect, setError } from 'react';
import './../../App';
import MyFluffy from './../../Components/Icons/MyFluffy_logo.svg';
import Posts from '../../Components/Posts';
import axios from 'axios';
import BottomBar from '../../Components/BottomBar';

function Home({ currentPage, setCurrentPage }) {
  useEffect(() => {
    setCurrentPage('Home');
  }, [setCurrentPage]);

  const [posts, setPosts] = useState([]);
  const [savedPosts, setSavedPosts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [currentFetchPage, setCurrentFetchPage] = useState(1); 
  const [hasMore, setHasMore] = useState(true); 

  const fetchPosts = async (page) => {
    if (!hasMore) return;

    setLoading(true);
    try {
      const response = await axios.get(`http://localhost:3000/api/v1/posts/feed?page=${page}`);
      const newPosts = response.data; // Adjust this line based on your actual API response
      console.log('Fetched posts:', newPosts);

      if (newPosts.length === 0) {
        setHasMore(false); // No more posts to fetch
      } else {
        const formattedPosts = newPosts.map(post => ({
          id: post.id,
          username: post.username,
          owner_id: post.owner_id,
          found_id: post.found_id,
          title: post.title,
          content: post.content,
          status: post.status,
          picture: post.picture,
          reward: post.reward
        }));

        setPosts(prevPosts => [...prevPosts, ...formattedPosts]); // Append new posts
        setCurrentFetchPage(prevPage => prevPage + 1); // Increment current fetch page
      }
    } catch (error) {
      console.error('Error fetching posts:', error);
      setError('Error fetching posts. Please try again later.');
    } finally {
      setLoading(false);
    }
  };

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
        content: post.content,
        owner_id: post.owner_id,
        found_id: post.found_id,
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
    fetchPosts(currentFetchPage);
    fetchSavedPosts();
  }, []); 

  const handleScroll = () => {
    if (window.innerHeight + document.documentElement.scrollTop + 1 >= document.documentElement.offsetHeight) {
      fetchPosts(currentFetchPage);
    }
  };

  useEffect(() => {
    window.addEventListener('scroll', handleScroll);
    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  }, [currentFetchPage]); // Re-run when currentFetchPage changes

  return (
    <div className="flex sm:flex-col md:flex-row h-auto sm:w-full pt-20 sm:px-4 md:px-6 md:mx-auto lg:px-8">
      <div className='sm:hidden md:block h-auto pt-4 sm:w-full md:w-1/3'>
        <div className='h-screen w-full flex'>
          <img src={MyFluffy} alt='My Fluffy Logo' className='w-8 h-8 mr-4' />
          <h1 className='sm:text-xl md:text-2xl font-bold'>My Fluffy</h1>
        </div>
      </div>
      <div className="h-auto sm:w-full md:w-1/3 pt-4 pb-20 flex flex-col">
        <div className="h-auto md:ml-16 md:pl-3 text-left">
          <h1 className="sm:text-xl md:text-2xl font-bold">Posts</h1>
          <BottomBar currentPage={currentPage} />
        </div>

        <div className="flex flex-col items-center">
          {loading && <p>Loading posts...</p>}
          <Posts data={posts} savedPosts={savedPosts} />
        </div>


      </div>

    </div>
  );
}

export default Home;
