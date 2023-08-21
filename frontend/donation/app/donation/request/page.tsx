"use client"

import { SSX } from '@spruceid/ssx';

const DonationRequest = () => {
    const signInButtonHandler = async () => {
        const ssx = new SSX({
            modules: {
              storage: true
            },
          });
        const session = await ssx.userAuthorization.signIn();
        await ssx.storage.put('wallet', session["address"]);
        await ssx.storage.put('user', session["sessionKey"]);
        console.log(session);
        
    };
    return(
        <div className="justify-center items-center h-1/2 my-32 grid grid-rows-2 gap-12">
            <h1 className='text-black'>Verify Yourself Before Start Donation</h1>
            <button type="button" onClick={signInButtonHandler} className="text-white w-full mx-auto bg-green-700 hover:bg-green-800 focus:outline-none focus:ring-4 focus:ring-green-300 font-medium rounded-full text-sm px-5 py-2.5 text-center mr-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">Verify</button>
        </div>
    )
}

export default DonationRequest;