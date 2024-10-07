import React, { useEffect, useState } from 'react';
import FoundIcon from './Icons/found_post_icon.svg';
import RewardIcon from './Icons/reward_post_icon.svg';
import UnsaveIcon from './Icons/unsave_post_icon.svg';
import SaveIcon from './Icons/save_post_icon.svg';
import axios from 'axios';

export default function ShowPost({ post, setIsModalVisible, username}) {
  const [isSaved, setIsSaved] = useState(false);
  const [postStatus, setPostStatus] = useState(post.status);
  const [sessionCookie, setSessionCookie] = useState('');

  useEffect(() => {
    setIsSaved(post.isSaved);

    // Safely read the session cookie
    const cookies = document.cookie.split('; ');
    const sessionCookieValue = cookies.find(cookie => cookie.startsWith('session='));
    if (sessionCookieValue) {
      setSessionCookie(sessionCookieValue.split('=')[1]);
    } else {
      setSessionCookie(''); // or handle it however you'd like, e.g., set to null
    }
  }, [post]);

  const handleSave = async () => {
    try {
      if (isSaved) {
        await axios.delete(`http://localhost:3000/api/v1/saved_posts`, {
          data: { post_id: post.id },
          withCredentials: true,
        });
      } else {
        await axios.post(`http://localhost:3000/api/v1/saved_posts`, {
          post_id: post.id,
        }, { withCredentials: true });
      }
      setIsSaved(!isSaved);
    } catch (error) {
      console.error('Error saving/unsaving post:', error);
    }
  };

  const handleStatusClick = async () => {
    try {
      if (postStatus === 'Missing') {
        await axios.put(`http://localhost:3000/api/v1/posts/found`, {
          found_id: sessionCookie,
          id: post.id,
        }, {
          withCredentials: true,
        });

        setPostStatus('Pending');
      } else if (postStatus === 'Pending' && sessionCookie === post.owner_id) {
        await axios.put(`http://localhost:3000/api/v1/posts/confirmation`, {
          id: post.id,
        }, {
          withCredentials: true,
        });

        setPostStatus('Found');

        console.log({
          amount: post.reward,
          receiver_id: post.found_id,
        });

        await axios.post(`http://localhost:3000/api/v1/payments/`, {
          amount: post.reward,
          receiver_id: post.found_id,
        }, {
          withCredentials: true,
        });
      }
    } catch (error) {
      console.error('Error updating post status:', error);
    }
  };

        useEffect(() => {
            console.log(post); // Check the value of `post` including `username`
            setIsSaved(post.isSaved);

            // Safely read the session cookie
            const cookies = document.cookie.split('; ');
            const sessionCookieValue = cookies.find(cookie => cookie.startsWith('session='));
            if (sessionCookieValue) {
                setSessionCookie(sessionCookieValue.split('=')[1]);
            } else {
                setSessionCookie(''); // or handle it however you'd like, e.g., set to null
            }
        }, [post]);

  return (
    <>
      <div
        className="fixed inset-0 bg-black bg-opacity-50 z-40"
        onClick={() => setIsModalVisible(false)}
      ></div>
      <div className="fixed top-0 left-0 w-full h-full flex items-center justify-center z-50">
        <div className="bg-white w-11/12 md:w-1/2 lg:w-1/3 p-6 rounded-lg shadow-lg relative">
          <button
            className="absolute top-2 right-2 text-[#43337B] hover:text-red-500 text-3xl font-bold"
            onClick={() => setIsModalVisible(false)}
          >
            &times;
          </button>
          <div className="flex flex-col items-start">
            <div className="w-full flex justify-between items-center mb-4">
              <div className="flex items-center">
                <img src={post.OwnerPicture || '/Profiles/default_pfp.jpg'} alt="Owner" className="w-10 h-10 rounded-full" />
                <div className="flex flex-col ml-2">
                  <h1 className="text-sm font-semibold text-[#43337B]">{username || 'Unknown'}</h1>
                </div>
              </div>
            </div>

            <img
              src={`data:image/jpeg;base64,${post.picture}`}
              alt={post.title}
              className="w-full h-56 object-cover rounded-lg mb-4 shadow-md"
            />

            <div className="w-full flex justify-between items-center mb-2">
              <h2 className="text-xl text-[#43337B] font-bold">{post.title}</h2>
              <button onClick={handleSave}>
                <img src={isSaved ? SaveIcon : UnsaveIcon} alt="save_icon" className="w-6 h-6" />
              </button>
            </div>

            <p className="text-sm text-[#43337B] mb-5 text-left">
              <span className="font-bold">Detail:</span> {post.content}
            </p>

            <div className="w-full flex items-center mt-2">
              <div className="w-1/2 flex items-center">
                <img src={RewardIcon} alt="reward_icon" className="w-4 h-4" />
                <h2 className="text-md text-[#43337B] font-semibold mx-2">Reward:</h2>
                <h1>{post.reward || 'N/A'}</h1>
              </div>

              <div className="w-1/2 flex items-center justify-end">
                <button 
                  className="flex w-auto h-8 py-1 px-2 bg-[#F1642E] rounded-lg"
                  onClick={handleStatusClick}
                >
                  <h2 className="text-md text-white font-medium mx-2">
                    {postStatus === 'Pending' && sessionCookie === post.owner_id ? 'Confirm?' : postStatus}
                  </h2>
                  <img src={FoundIcon} alt="found_icon" className="w-6 h-6" />
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
