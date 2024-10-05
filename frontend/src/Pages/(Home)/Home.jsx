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
      Type: 'Lost',
      Owner: 'Owner1',
      OwnerPicture: '/Profiles/default_pfp.jpg',
      Location: 'Location1',
      Title: 'Title1',
      Description: 'Description1',
      Reward: '1000',
      Picture: '/Images/pet1.png',
    },
    {
      id: 2,
      Type: 'Found',
      Owner: 'Owner2',
      OwnerPicture: '/Profiles/default_pfp.jpg',
      Location: 'Location2',
      Title: 'Title2',
      Description: 'Description2',
      Reward: '2000',
      Picture: '/Images/pet1.png',
    }
  ];

  return (
    <div 
      className="
        flex sm:flex-col md:flex-row
        h-auto w-full
        pt-20
        sm:px-4 md:px-6 lg:px-8
    ">
        <div className='h-auto pt-4 pb-20'>
          <div className='h-auto'>
            <h1 className='sm:text-xl md:text-2xl font-bold'>Posts</h1>
          </div>
          <Posts data={data} />
        </div>
    </div>
  );
}

export default Home;