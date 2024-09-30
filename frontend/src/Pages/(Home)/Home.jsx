import React from 'react';

import './../../App';
import TopBar from '../../Components/TopBar';

function Home() {
  return (
    <div className="bg-blue-500 bg-auto h-screen">
      {/* <h1 className="text-4xl font-bold text-green-500">Hello, Tailwind CSS!</h1> */}
        <TopBar />
        <div className='h-auto bg-red-200'>

        </div>
    </div>
  );
}

export default Home;