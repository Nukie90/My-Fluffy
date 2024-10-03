import React from 'react';
import { useState, useEffect } from 'react';

import './../../App';
import TopBar from '../../Components/TopBar';
import BottomBar from '../../Components/BottomBar';
import Posts from '../../Components/Posts';

function Home({CurrentPage, setCurrentPage}) {
  useEffect(() => {
      setCurrentPage('Home');
  }, [setCurrentPage]);

  const data = [
    {
      id: 1,
      user: 'User1',
      userImage: '/Profiles/default_pfp.jpg',
      pet: 'Pet1',
      location: 'Location1',
      date: 'Date1',
      time: 'Time1',
      description: 'Description1',
      tags: ['Tag1', 'Tag2', 'Tag3'],
      reward: '1000',
      image: '/Images/pet1.png',
    },
  ];

  return (
    <div 
      className="
        flex sm:flex-col md:flex-row
        h-auto w-full
        pt-20 
        sm:px-4 md:px-6 lg:px-8
    ">
        <div className='h-auto py-4'>
          <div className='h-auto'>
            <h1 className='sm:text-xl md:text-2xl font-bold'>Near You</h1>
          </div>
          <Posts data={data} />
        </div>
    </div>
  );
}

export default Home;