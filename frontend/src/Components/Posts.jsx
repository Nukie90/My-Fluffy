import React, { useState, useEffect } from 'react';
import FoundIcon from './Icons/found_post_icon.svg';
import RewardIcon from './Icons/reward_post_icon.svg';
import UnsaveIcon from './Icons/unsave_post_icon.svg';
import SaveIcon from './Icons/save_post_icon.svg';
import axios from 'axios';

export default function Posts({ data, savedPosts }) {
  const [localSavedPosts, setLocalSavedPosts] = useState({});
  const [postStatuses, setPostStatuses] = useState({});
  const [sessionCookie, setSessionCookie] = useState(''); // Track the fetched username

  useEffect(() => {
    const initialSavedPosts = {};
    savedPosts.forEach(post => {
      initialSavedPosts[post.id] = true; 
    });
    setLocalSavedPosts(initialSavedPosts);
    
    const initialPostStatuses = {};
    data.forEach(post => {
      initialPostStatuses[post.id] = post.status; 
    });
    setPostStatuses(initialPostStatuses);
  
    // Safely read the session cookie
    const cookies = document.cookie.split('; ');
    const sessionCookieValue = cookies.find(cookie => cookie.startsWith('session='));
    if (sessionCookieValue) {
      setSessionCookie(sessionCookieValue.split('=')[1]);
    } else {
      setSessionCookie(''); // or handle it however you'd like, e.g., set to null
    }
  
  }, [savedPosts, data]); 
  

  const handleSave = async (id) => {
    const isSaved = localSavedPosts[id];

    try {
      if (isSaved) {
        await axios.delete(`http://localhost:3000/api/v1/saved_posts`, {
          data: {
            post_id: id,
          },
          withCredentials: true,
        });        
      } else {
        await axios.post(`http://localhost:3000/api/v1/saved_posts`, {
          post_id: id,
        }, {
          withCredentials: true,
        });
      }

      setLocalSavedPosts((prev) => ({
        ...prev,
        [id]: !isSaved,
      }));
    } catch (error) {
      console.error('Error saving/unsaving post:', error);
    }
  };

  const handleStatusClick = async (post) => {
    try {
      if (post.status === 'Missing') {
        await axios.put(`http://localhost:3000/api/v1/posts/found`, {
          found_id: sessionCookie,
          id: post.id,
        }, {
          withCredentials: true,
        });
        
        // Update the local status
        setPostStatuses(prev => ({
          ...prev,
          [post.id]: 'Pending', // Update status to 'Pending'
        }));
      } else if (post.status === 'Pending' && sessionCookie === post.owner_id) {
        await axios.put(`http://localhost:3000/api/v1/posts/confirmation`, {
          id: post.id,
        }, {
          withCredentials: true,
        });
        
        // Update the local status
        setPostStatuses(prev => ({
          ...prev,
          [post.id]: 'Found', // Update status to 'Confirmed'
        }));

        console.log({
          amount: post.reward,
          receiver_id: post.found_id
        });
        
        await axios.post(`http://localhost:3000/api/v1/payments/`, {
          amount: post.reward,
          receiver_id: post.found_id
        }, {
          withCredentials: true,
        });
      }
    } catch (error) {
      console.error('Error updating post status:', error);
    }
  };  
  
  return (
    <div className='md:w-3/4 h-auto pb-4 flex flex-col justify-center items-center'>
      {data.map((post) => {
        const isSaved = localSavedPosts[post.id]; // Check saved status in local state
        const postStatus = postStatuses[post.id]; // Get the current status from local state

        return (
          <div key={post.id} className='h-auto sm:w-full bg-white p-4 my-2 rounded-lg shadow-md'>
            <div className='h-auto w-full flex flex-col'>
              <div className='w-full flex justify-between items-center'>
                <div className='flex items-center'>
                  <img src={post.OwnerPicture || '/Profiles/default_pfp.jpg'} alt={post.OwnerPicture} className='w-10 h-10 rounded-full' />
                  <div className='w-1/2 flex flex-col ml-2'>
                    <h1 className='text-sm font-semibold colorct-dark-purple'>{post.username || 'Unknown'}</h1>
                  </div>
                </div>
              </div>

              <div className='h-auto w-full my-4'>
                <img src={`data:image/png;base64,${post.picture}`} alt={post.title} className='w-full h-full object-cover' />
              </div>
            </div>

            <div className='h-auto w-full flex flex-col items-center'>
              <div className='h-auto w-full flex'>
                <h2 className='text-xl colorct-dark-purple font-bold'>{post.title}</h2>
                <div className='h-auto w-full flex justify-end items-center'>
                  <button onClick={() => handleSave(post.id)}>
                    <img src={isSaved ? SaveIcon : UnsaveIcon} alt='save_icon' className='w-6 h-6' />
                  </button>
                </div>
              </div>
              <div className='h-auto w-full'>
                <p className='text-md colorct-dark-purple'>{post.content}</p>
              </div>

              <div className='w-full flex items-center mt-2'>
                <div className='w-1/2 flex items-center'>
                  <img src={RewardIcon} alt='reward_icon' className='w-4 h-4' />
                  <h2 className='text-md colorct-dark-purple font-semibold mx-2'>Reward:</h2>
                  <h1>{post.reward || 'N/A'}</h1>
                </div>
                <div className='w-1/2 flex items-center justify-end'>
                  <button 
                    className='flex w-auto h-8 py-1 px-2 bgct-orange rounded-lg' 
                    onClick={() => handleStatusClick(post)}
                  >
                    <h2 className='text-md text-white font-medium mx-2'>{postStatus === 'Pending' && sessionCookie === post.owner_id ? 'Confirm?' : postStatus}</h2>
                    <img src={FoundIcon} alt='found_icon' className='w-6 h-6' />
                  </button>
                </div>
              </div>
            </div>
          </div>
        );
      })}
    </div>
  );
}
