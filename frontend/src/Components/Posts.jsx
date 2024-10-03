import React from 'react';

export default function Posts({ data }) {

    return (
        <div className='h-auto py-4 flex justify-center items-center'>
            {data.map((post, index) => (
              <div key={index} className='h-auto w-full bg-white p-4 rounded-lg shadow-md'>
                <div className='h-auto w-full flex justify-between items-center'>
                  <div className='flex items-center'>
                    <img src={post.userImage} alt={post.userImage} className='w-10 h-10 rounded-full' />
                    <div className='flex flex-col ml-2'>
                        <h1 className='font-bold'>{post.user}</h1>
                        <h1>{post.location}</h1>
                        <h2>{post.pet}</h2>
                    </div>
                  </div>
                </div>
                <div className='h-auto w-full'>
                    <img src={post.image} alt={post.image} className='w-full h-64 object-cover' />
                    <h2>{post.date} {post.time}</h2>
                    <p>{post.description}</p>
                </div>
                <div className='h-auto w-full flex'>
                  {post.tags.map((tag, index) => (
                    <p key={index} className='text-sm pr-2 colorct-dark-purple'>#{tag}</p>
                  ))}
                </div>
                <h1>{post.reward}</h1>
              </div>
            ))}
          </div>
    );
}