import React from 'react';
import { useState, useEffect } from 'react';

import './../../App';
import TopBar from '../../Components/TopBar';
import BottomBar from '../../Components/BottomBar';

function Home({CurrentPage, setCurrentPage}) {
  useEffect(() => {
      setCurrentPage('Home');
  }, [setCurrentPage]);

  return (
    <div 
      className="
        bgct-yellow h-screen w-full
        pt-20
    ">
        <div className='h-auto bg-red-200'>
          this is the home page
        </div>
        <BottomBar />
    </div>
  );
}

export default Home;