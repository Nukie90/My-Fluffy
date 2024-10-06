import React, { useState, useEffect, setError } from 'react';
import './../../App';
import Posts from '../../Components/Posts';
import axios from 'axios';

function Home({ currentPage, setCurrentPage }) {
  useEffect(() => {
    setCurrentPage('Home');
  }, [setCurrentPage]);

  const [posts, setPosts] = useState([]);
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

  useEffect(() => {
    fetchPosts(currentFetchPage); // Fetch posts on initial load
  }, []); // Empty dependency array means this effect runs only once on mount

  const handleScroll = () => {
    if (window.innerHeight + document.documentElement.scrollTop + 1 >= document.documentElement.offsetHeight) {
      fetchPosts(currentFetchPage); // Fetch next page when scrolled to the bottom
    }
  };

  useEffect(() => {
    window.addEventListener('scroll', handleScroll);
    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  }, [currentFetchPage]); // Re-run when currentFetchPage changes

  return (
    <div className="flex sm:flex-col md:flex-row h-auto w-full pt-20 sm:px-4 md:px-6 lg:px-8">
      <div className='h-auto pt-4 pb-20'>
        <div className='h-auto'>
          <h1 className='sm:text-xl md:text-2xl font-bold'>Posts</h1>
        </div>

        {loading && <p>Loading posts...</p>}

        <Posts data={posts} />
      </div>
    </div>
  );
}

export default Home;
