import React from 'react';
import RewardPic from './../../Components/Icons/reward_pic.svg';

export default function RewardPage() {
    const rewards = [
        {
            id: 1,
            name: 'Alice Johnson',
            petDescription: 'Bella, Golden Retriever, 1 year',
            details: 'Last seen at Central Park. Please contact if found.',
            reward: '2,000 B',
            image: 'https://example.com/golden-retriever.jpg'
        },
        {
            id: 2,
            name: 'Michael Smith',
            petDescription: 'Charlie, Beagle, 3 years',
            details: 'Lost near Elm Street. He is very friendly.',
            reward: '1,800 B',
            image: 'https://example.com/beagle.jpg'
        },
        {
            id: 3,
            name: 'Emily Davis',
            petDescription: 'Daisy, Persian Cat, 6 months',
            details: 'Missing from Green Avenue. She is shy and scared.',
            reward: '2,500 B',
            image: 'https://example.com/persian-cat.jpg'
        },
        {
            id: 4,
            name: 'James Wilson',
            petDescription: 'Max, Labrador, 2 years',
            details: 'Lost near the riverbank. Loves playing fetch.',
            reward: '2,200 B',
            image: 'https://example.com/labrador.jpg'
        },
        {
            id: 5,
            name: 'Sophia Brown',
            petDescription: 'Luna, Siamese Cat, 1 year',
            details: 'Disappeared from my backyard on Oak Street.',
            reward: '1,900 B',
            image: 'https://example.com/siamese-cat.jpg'
        },
        {
            id: 6,
            name: 'Daniel Miller',
            petDescription: 'Rocky, Bulldog, 4 years',
            details: 'Went missing near the downtown area. Very gentle.',
            reward: '2,300 B',
            image: 'https://example.com/bulldog.jpg'
        }
    ];

    return (
        <div className="
                flex flex-col md:flex-row flex-wrap
                h-auto w-full
                pt-20
                px-4 md:px-6 lg:px-8
            ">

            <div className="h-auto pt-4 md:w-full">
                <h1 className="text-xl md:text-2xl font-bold mb-6 text-[#504E76] text-center md:text-left mt-3">Your Rewards</h1>
            </div>

            <div className="flex justify-around mb-6 md:w-full">
                <div className="text-center">
                    <h2 className="text-3xl font-bold text-[#504E76]">3</h2>
                    <p className="text-gray-600">Found</p>
                </div>
                <div className="text-center">
                    <h2 className="text-3xl font-bold text-[#504E76]">6,000</h2>
                    <p className="text-gray-600">Reward</p>
                </div>
            </div>

            <div className="w-full grid gap-4 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
                {rewards.map((reward) => (
                    <div key={reward.id} className="bg-[#C4C3E3] rounded-lg p-4 flex flex-col space-y-4">
                        <div className="flex items-center space-x-2">
                            <div className="w-8 h-8 bg-purple-500 rounded-full"></div>
                            <h3 className="text-md font-bold text-gray-900">{reward.name}</h3>
                        </div>

                        <div className="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-4">
                            <div className="w-full sm:w-24 h-24">
                                <img
                                    src={reward.image || RewardPic}
                                    alt="Pet"
                                    className="w-full h-full object-cover rounded-lg"
                                />
                            </div>

                            <div className="flex-1">
                                <p className="text-lg font-bold text-gray-900">{reward.petDescription}</p>
                                <p className="text-sm text-gray-500 mt-1">{reward.details}</p>
                            </div>
                        </div>

                        <div className="text-right">
                            <p className="font-bold text-gray-900">Reward: {reward.reward}</p>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}