import React from 'react';
import { useState } from 'react';

import './../../App';
import TopBar from '../../Components/TopBar';
import BottomBar from '../../Components/BottomBar';

function Home() {
  return (
    <div className="bg-white bg-auto h-screen">
        <div className='h-auto bg-red-200'>
          this is the home page
        </div>
        <BottomBar />
    </div>
  );
}

export default Home;