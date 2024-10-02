import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { useState } from 'react';

import './App.css';
import BottomBar from './Components/BottomBar';
import TopBar from './Components/TopBar';
import Home from './Pages/(Home)/Home';
import Rewards from './Pages/Rewards/Rewards';
import Saved from './Pages/Saved/Saved';
import Profile from './Pages/Profile/Profile';
import Login from './Pages/Auth/Login';
import Signup from './Pages/Auth/Signup';

function App() {
  const [currentPage, setCurrentPage] = useState('Home');

  return (
    <div className='min-h-screen'>
      <Router>
        {currentPage !== 'Login' && currentPage !== 'Signup' && (
          <TopBar currentPage={currentPage} setCurrentPage={setCurrentPage} />
        )}
        <Routes>
            <Route path="/" element={<Home currentPage={currentPage} setCurrentPage={setCurrentPage} />} />
            <Route path="/rewards" element={<Rewards currentPage={currentPage} setCurrentPage={setCurrentPage} />} />
            <Route path="/saved" element={<Saved currentPage={currentPage} setCurrentPage={setCurrentPage} />}/>
            <Route path="/profile" element={<Profile currentPage={currentPage} setCurrentPage={setCurrentPage} />} />
            <Route path="/login" element={<Login currentPage={currentPage} setCurrentPage={setCurrentPage} />} />
            <Route path="/signup" element={<Signup currentPage={currentPage} setCurrentPage={setCurrentPage} />} />
        </Routes>
        {currentPage !== 'Login' && currentPage !== 'Signup' && (
          <BottomBar currentPage={currentPage} setCurrentPage={setCurrentPage} />
        )}
      </Router>
    </div>
  );
}

export default App;
