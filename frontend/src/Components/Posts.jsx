import React, { useState } from 'react';

import FoundIcon from './Icons/found_post_icon.svg';
import RewardIcon from './Icons/reward_post_icon.svg';
import UnsaveIcon from './Icons/unsave_post_icon.svg';
import SaveIcon from './Icons/save_post_icon.svg';

export default function Posts({ data }) {
  const [savedPosts, setSavedPosts] = useState({});

  const handleSave = (id) => {
    setSavedPosts((prev) => ({
      ...prev,
      [id]: !prev[id],
    }));
  };

  return (
    <div className='h-auto pb-4 flex flex-col justify-center items-center'>
      {data.map((post) => (
        <div key={post.id} className='h-auto w-full bg-white p-4 my-2 rounded-lg shadow-md'>
          <div className='h-auto w-full flex flex-col'>
            <div className='w-full flex justify-between items-center'>
              {/* Left side: User image and info */}
              <div className='flex items-center'>
                {/* Placeholder for OwnerPicture; you might want to adjust this based on the response */}
                <img src={post.OwnerPicture || '/Profiles/default_pfp.jpg'} alt={post.OwnerPicture} className='w-10 h-10 rounded-full' />
                <div className='w-1/2 flex flex-col ml-2'>
                  <h1 className='text-sm font-semibold colorct-dark-purple'>{post.Owner || 'Unknown'}</h1>
                </div>
              </div>
            </div>

            <div className='h-auto w-full my-4'>
              {/* Updated to use the picture field from the response */}
              <img src={`data:image/png;base64,${post.picture}`} alt={post.title} className='w-full h-full object-cover' />
            </div>
          </div>

          <div className='h-auto w-full flex flex-col items-center'>
            {/* Title + Description */}
            <div className='h-auto w-full flex'>
              <h2 className='text-xl colorct-dark-purple font-bold'>{post.title}</h2>
              {/* Right side: Save button */}
              <div className='h-auto w-full flex justify-end items-center'>
                <button onClick={() => handleSave(post.id)}>
                  <img src={savedPosts[post.id] ? SaveIcon : UnsaveIcon} alt='save_icon' className='w-6 h-6' />
                </button>
              </div>
            </div>
            <div className='h-auto w-full'>
              <p className='text-md colorct-dark-purple'>{post.content}</p>
            </div>

            {/* Reward - Update according to the new structure if applicable */}
            <div className='w-full flex items-center mt-2'>
              <div className='w-1/2 flex items-center'>
                <img src={RewardIcon} alt='reward_icon' className='w-4 h-4' />
                <h2 className='text-md colorct-dark-purple font-semibold mx-2'>Reward:</h2>
                <h1>{post.reward || 'N/A'}</h1> {/* Adjust accordingly if Reward is part of the new structure */}
              </div>
              <div className='w-1/2 flex items-center justify-end'>
                <button className='flex w-auto h-8 py-1 px-2 bgct-orange rounded-lg'>
                  <h2 className='text-md text-white font-medium mx-2'>Found</h2>
                  <img src={FoundIcon} alt='found_icon' className='w-6 h-6' />
                </button>
              </div>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}
