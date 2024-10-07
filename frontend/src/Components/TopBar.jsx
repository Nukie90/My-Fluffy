import './../App.css';
import Noti from './Noti';
import SearchIcon from './Icons/search_icon.svg';
import NotiIcon from './Icons/noti_icon.svg';
import DMIcon from './Icons/dm_icon.svg';
import { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

function TopBar({ CurrentPage, setCurrentPage }) {
    const [isNotiVisible, setIsNotiVisible] = useState(false);
    const [unReadNoti, setUnReadNoti] = useState(0);
    const [noti, setNoti] = useState([]);

    let pfp = "/Profiles/default_pfp.jpg";
    const toprightIcons = [NotiIcon];

    useEffect(() => {
        const fetchNotifications = async () => {
            try {
                const response = await axios.get('http://localhost:3000/api/v1/notifications', {
                    withCredentials: true,
                });
                const notifications = response.data || [];

                setNoti(notifications);
                
                if (notifications.length > 0) {
                    setUnReadNoti(notifications.filter(noti => !noti.isRead).length);
                }
            } catch (error) {
                console.error('Error fetching notifications:', error);
            }
        };

        fetchNotifications();
    }, []);

    function handleNotiClick() {
        setIsNotiVisible(!isNotiVisible);
    }

    return (
        <>
            <div className="fixed top-0 left-0 h-20 w-full flex text-center bg-white shadow-md items-center justify-between sm:pt-6 md:pt-0 sm:px-4 md:px-6 lg:px-8 z-50"
                style={{ color: 'var(--color-dark-purple)' }}
            >
                <div className='flex items-center justify-center rounded-full'>
                    <Link to={'/profile'}>
                        <button 
                            onClick={() => setCurrentPage('Profile')}
                            className='
                                sm:w-9 h-9 
                                sm:hover:w-10 h-10 
                                lg:w-10 h-10 
                                lg:hover:w-12 h-12 
                                xl:w-12 h-12
                                xl:hover:w-14 h-14
                                rounded-full transition-all duration-300'
                        >
                            <img src={pfp} alt="Profile Pic" key={"Profile Pic"} className='rounded-full' />
                        </button>
                    </Link>
                </div>
                <div className='flex justify-end py-5'>
                    {toprightIcons.map((icon, index) => (
                        <button
                            className='
                                relative sm:flex items-center justify-center
                                rounded-full transition-all duration-300
                                sm:ml-4 w-6 h-6
                                lg:ml-6 w-8 h-8
                                xl:w-9 h-9
                            '
                            key={index}
                            onClick={icon === NotiIcon ? handleNotiClick : null}
                        >
                            <img src={icon} alt={icon} key={icon}
                                className="w-full h-full" 
                            />
                            {icon === NotiIcon && unReadNoti > 0 && (
                                <div className='
                                    absolute top-0 right-0
                                    flex items-center justify-center
                                    sm:w-4 sm:h-4
                                    xl:w-5 xl:h-5
                                    bg-orange-500 text-white
                                    rounded-full text-xs
                                    transform translate-x-1 translate-y-5'
                                >
                                    {unReadNoti}
                                </div>
                            )}
                        </button>
                    ))}
                </div>
            </div>

            {isNotiVisible && (
                <>
                    <div
                        className="fixed inset-0 bg-black bg-opacity-50 z-40"
                        onClick={() => setIsNotiVisible(false)}
                    ></div>
                    <div className="fixed top-0 right-0 h-full w-full flex items-center justify-center z-50">
                        <Noti noti={noti} setIsNotiVisible={setIsNotiVisible} setUnReadNoti={setUnReadNoti} />
                    </div>
                </>
            )}
        </>
    )
}

export default TopBar;
