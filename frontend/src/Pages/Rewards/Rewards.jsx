import React, { useEffect, useState } from 'react';
import axios from 'axios';
import RewardPic from './../../Components/Icons/reward_pic.svg';

export default function RewardPage() {
    const [rewards, setRewards] = useState([]);
    const [usernames, setUsernames] = useState({}); // State to hold usernames
    const [totalFound, setTotalFound] = useState(0); // State to hold total found
    const [totalAmount, setTotalAmount] = useState(0); // State to hold total reward amount

    useEffect(() => {
        const fetchRewards = async () => {
            try {
                const response = await axios.get('http://localhost:3000/api/v1/payments/user', {
                    withCredentials: true,
                });
                setRewards(response.data);
                console.log(response.data);

                // Calculate total found and total reward amount
                setTotalFound(response.data.length);
                const totalRewardAmount = response.data.reduce((sum, reward) => sum + parseFloat(reward.amount), 0);
                setTotalAmount(totalRewardAmount);

                // Fetch usernames for each owner_id
                const usernamesFetched = await Promise.all(response.data.map(async (reward) => {
                    const userResponse = await axios.get(`http://localhost:3000/api/v1/users/${reward.owner_id}`, {
                        withCredentials: true,
                    });
                    return { id: reward.id, owner_id: reward.owner_id, username: userResponse.data.username };
                }));

                // Convert array to an object for easy access
                const usernamesObject = usernamesFetched.reduce((acc, curr) => {
                    acc[curr.owner_id] = curr.username;
                    return acc;
                }, {});
                
                setUsernames(usernamesObject);

            } catch (err) {
                console.log(err.message || 'Failed to fetch rewards');
            }
        };

        fetchRewards();
    }, []);

    function convertDate(dateString) {
        const options = {
            year: 'numeric',
            month: 'long',
            day: 'numeric',
            hour: '2-digit',
            minute: '2-digit',
            hour12: false, // Use 24-hour format
        };
    
        const date = new Date(dateString);
        return date.toLocaleString('en-US', options);
    }

    return (
        <div className="
                flex flex-col md:flex-row flex-wrap
                h-auto w-full
                pt-20 pb-24
                px-4 md:px-6 lg:px-8
            ">

            <div className="h-auto pt-4 md:w-full">
                <h1 className="text-xl md:text-2xl font-bold mb-6 text-[#504E76] text-center md:text-left mt-3">Your Rewards</h1>
            </div>

            <div className="flex justify-around mb-6 md:w-full">
                <div className="text-center">
                    <h2 className="text-3xl font-bold text-[#504E76]">{totalFound}</h2> {/* Display total found */}
                    <p className="text-gray-600">Found</p>
                </div>
                <div className="text-center">
                    <h2 className="text-3xl font-bold text-[#504E76]">{totalAmount.toLocaleString()}</h2> {/* Display total amount */}
                    <p className="text-gray-600">Reward</p>
                </div>
            </div>

            <div className="w-full grid gap-4 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
                {rewards.map((reward) => (
                    <div key={reward.id} className="bg-[#C4C3E3] rounded-lg p-4 flex flex-col space-y-4">
                        <div className="flex items-center space-x-2">
                            <div className="w-8 h-8 bg-purple-500 rounded-full">
                                <img
                                    src={`/Profiles/default_pfp.jpg`}
                                    alt="User"
                                    className="w-full h-full object-cover rounded-full"
                                />
                            </div>
                            <h3 className="text-md font-bold text-gray-900">{reward.id || 'Loading...'}</h3> {/* Display username */}
                        </div>

                        <div className="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4">
                            <div className="w-28 h-28 rounded-lg bg-green-200">
                                <img
                                    src={reward.image || RewardPic}
                                    alt="Pet"
                                    className="w-full h-full object-cover"
                                />
                            </div>

                            <div className="flex-1">
                                <p className="text-lg font-semibold text-gray-900">{usernames[reward.owner_id]} gave you a reward!</p>
                                <p className="text-sm text-gray-500 mt-1">{convertDate(reward.created_at)}</p>
                            </div>
                        </div>

                        <div className="text-right">
                            <p className="font-bold text-gray-900">Reward: {reward.amount}</p>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}
