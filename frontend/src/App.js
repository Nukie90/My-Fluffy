import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { useState } from 'react';

import './App.css';
import BottomBar from './Components/BottomBar';
import TopBar from './Components/TopBar';
import Home from './Pages/(Home)/Home';
import Rewards from './Pages/Rewards/Rewards';
import Saved from './Pages/Saved/Saved';
import Profile from './Pages/Profile/Profile';

import HomeIcon from './Components/Icons/home_icon.svg';
import AddIcon from './Components/Icons/add_icon.svg';
import RewardsIcon from './Components/Icons/rewards_icon.svg';
import RewardsIconActive from './Components/Icons/rewards_icon_active.svg';
import SavedIcon from './Components/Icons/saved_icon.svg';
import SavedIconActive from './Components/Icons/saved_icon_active.svg';

function App() {
  const [currentPage, setCurrentPage] = useState('home');

  return (
    <Router>
      <TopBar />
      <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/rewards" element={<Rewards />} />
          <Route path="/saved" element={<Saved />} />
          <Route path="/profile" element={<Profile />} />
      </Routes>
      <BottomBar currentPage={currentPage} setCurrentPage={setCurrentPage} />
    </Router>
  );
}

export default App;
