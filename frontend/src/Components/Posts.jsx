import React from 'react';

import FoundIcon from './Icons/found_post_icon.svg';
import RewardIcon from './Icons/reward_post_icon.svg';
import UnsaveIcon from './Icons/unsave_post_icon.svg';
import SaveIcon from './Icons/save_post_icon.svg';

export default function Posts({ data }) {

    return (
        <div className='h-auto py-4 flex flex-col justify-center items-center'>
            {data.map((post, index) => (
              <div key={index} className='h-auto w-full bg-white p-4 my-2 rounded-lg shadow-md'>
                <div className='h-auto w-full flex flex-col'>
                  
                  <div className='w-full flex justify-between items-center'>
                    {/* Left side: User image and info */}
                    <div className='flex items-center'>
                      <img src={post.OwnerPicture} alt={post.OwnerPicture} className='w-10 h-10 rounded-full' />
                      <div className='w-1/2 flex flex-col ml-2'>
                        <h1 className='text-sm font-semibold colorct-dark-purple'>{post.Owner}</h1>
                      </div>
                    </div>

                  </div>
                  
                  <div className='h-auto w-full my-4'>
                    <img src={post.Picture} alt={post.Picture} className='w-full h-full object-cover' />
                  </div>

                </div>

                <div className='h-auto w-full flex flex-col items-center'>

                  {/* Title + Description */}
                  <div className='h-auto w-full flex'>
                    <h2 className='text-xl colorct-dark-purple font-bold'>{post.Title}</h2>
                    {/* Right side: Save button */}
                    <div className='h-auto w-full flex justify-end items-center'>
                      <button >
                        <img src={UnsaveIcon} alt='save_icon' className='w-6 h-6' />
                      </button>
                    </div>
                  </div>
                  <div className='h-auto w-full'>
                    <p className='text-md colorct-dark-purple'>{post.Description}</p>
                  </div>
                  
                  {/* Reward */}
                  <div className='w-full flex items-center mt-2'>
                    <div className='w-1/2 flex items-center'>
                      <img src={RewardIcon} alt='reward_icon' className='w-4 h-4' />
                      <h2 className='text-md colorct-dark-purple font-semibold mx-2'>Reward:</h2>
                      <h1>{post.Reward}</h1>
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