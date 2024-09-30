import './../App';
import SearchIcon from './Icons/search_icon.svg';
import NotiIcon from './Icons/noti_icon.svg';
import DMIcon from './Icons/dm_icon.svg';
import DefaultPFP from './Profiles/default_pfp.jpg';

function TopBar() {

    let pfp = DefaultPFP;
    const toprightIcons = [SearchIcon, NotiIcon, DMIcon];

    return (
        <div className="h-auto flex text-center bg-white shadow-md items-center justify-between sm:px-4 md:px-6 lg:px-8"
            style={{ color: 'var(--color-dark-purple)' }}
        >
            <div className='flex items-center justify-center rounded-full'>
                <button className='sm:w-8 h-8 sm:hover:w-10 h-10 lg:w-10 h-10 lg:hover:w-12 h-12 rounded-full  transition-all duration-300'>
                    <img src={pfp} alt="Profile Pic" key={"Profile Pic"} className='rounded-full' />
                </button>
            </div>
            <div className='flex justify-end py-5'>
                {toprightIcons.map((icon, index) => (
                    <button
                        className='
                            flex items-center justify-center
                            rounded-full transition-all duration-300
                            sm:ml-4 w-6 h-6
                            lg:ml-6 w-8 h-8
                            xl:w-9 h-9
                        '
                        key={index}
                    >
                        <img src={icon} alt={icon} key={icon}
                            className="w-full h-full" 
                        />
                    </button>
                ))}
            </div>
        </div>
    )
}

export default TopBar;