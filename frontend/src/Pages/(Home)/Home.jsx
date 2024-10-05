import React, { useState, useEffect } from 'react';
import './../../App';
import Posts from '../../Components/Posts';
import axios from 'axios';

function Home({ currentPage, setCurrentPage }) {
  useEffect(() => {
    setCurrentPage('Home');
  }, [setCurrentPage]);

  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(false);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(true);

  const fetchPosts = async () => {
    if (loading || !hasMore) return; // Prevent fetching if already loading or no more posts

    setLoading(true);
    try {
      const response = await axios.get(`http://localhost:3000/api/v1/posts/feed?page=${page}`);
      const newPosts = response.data;
      console.log(newPosts);

      if (newPosts.length === 0) {
        setHasMore(false); 
      } else {
        const formattedPosts = newPosts.map(post => ({
          id: post.id,
          title: post.title,
          content: post.content,
          status: post.status,
          picture: post.picture
        }));

        setPosts(prevPosts => [...prevPosts, ...formattedPosts]);
        setPage(prevPage => prevPage + 1);
      }
    } catch (error) {
      console.error('Error fetching posts:', error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchPosts();
  }, []);

  const handleScroll = () => {
    if (window.innerHeight + document.documentElement.scrollTop !== document.documentElement.offsetHeight || loading) return;
    fetchPosts();
  };

  useEffect(() => {
    window.addEventListener('scroll', handleScroll);
    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  }, [loading, hasMore]);

  return (
    <div className="flex sm:flex-col md:flex-row h-auto w-full pt-20 sm:px-4 md:px-6 lg:px-8">
      <div className='h-auto pt-4 pb-20'>
        <div className='h-auto'>
          <h1 className='sm:text-xl md:text-2xl font-bold'>Posts</h1>
        </div>
        <Posts data={posts} />
        {loading && <p>Loading more posts...</p>}
        {!hasMore && <p>No more posts to load.</p>}
      </div>
    </div>
  );
}

export default Home;
