import './../../App.css';
import { useState, useEffect } from 'react'; 
import afterClick from './../../Components/Icons/saved_afterclick.svg'; 
import beforeClick from './../../Components/Icons/saved_beforeclick.svg';
import findPets from './../../Components/Icons/find_pet_pic.svg'; 
import foundPets from './../../Components/Icons/found_pet_pic.svg'; 

function Saved({ currentPage, setCurrentPage }) {
    const [activeTab, setActiveTab] = useState('findingPets');

    useEffect(() => {
        setCurrentPage('Saved');
    }, [setCurrentPage]);

    return (
        <div 
            className="
                flex flex-col md:flex-col
                h-auto w-full
                pt-20
                sm:px-4 md:px-6 lg:px-8
            ">
           
            <div className="h-auto pt-4 w-full">
                <h1 className="text-xl sm:text-xl md:text-2xl font-bold text-[#504E76] text-left md:text-left">Saved</h1>
            </div>


            {/* Tabs */}
            <div className="flex justify-center w-full mb-4 mt-4">
                <button
                    onClick={() => setActiveTab('findingPets')}
                    className={`flex items-center justify-center flex-1 text-center py-2 font-semibold 
                    ${activeTab === 'findingPets' ? 'text-orange-500 border-b-4 border-orange-500' : 'text-gray-500'} 
                    hover:text-orange-500`}
                >
                    <img 
                        src={activeTab === 'findingPets' ? afterClick : beforeClick}
                        alt="Finding Icon" 
                        className="w-6 h-6 mr-2"
                    />
                    Finding pets
                </button>
                <button
                    onClick={() => setActiveTab('petsFound')}
                    className={`flex items-center justify-center flex-1 text-center py-2 font-semibold 
                    ${activeTab === 'petsFound' ? 'text-orange-500 border-b-4 border-orange-500' : 'text-gray-500'} 
                    hover:text-orange-500`}
                >
                    <img 
                        src={activeTab === 'petsFound' ? afterClick : beforeClick}
                        alt="Pets Found Icon" 
                        className="w-6 h-6 mr-2"
                    />
                    Pets found
                </button>
            </div>

            {/* Saved Items */}
            <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-4">
                {/* Example images */}
                {activeTab === 'findingPets' ? (
                    [...Array(30)].map((_, i) => (
                        <img
                            key={i}
                            src={findPets}
                            alt="Saved pet"
                            className="w-full h-auto object-cover rounded"
                        />
                    ))
                ) : (
                    [...Array(30)].map((_, i) => (
                        <img
                            key={i}
                            src={foundPets}
                            alt="Found pet"
                            className="w-full h-auto object-cover rounded"
                        />
                    ))
                )}
            </div>
        </div>
    );
}

export default Saved;
