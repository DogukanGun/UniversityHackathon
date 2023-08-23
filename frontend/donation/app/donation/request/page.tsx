"use client"
import { useState, useEffect } from "react"
import { SSX } from '@spruceid/ssx';
import { useEnsAddress, useEnsAvatar, useEnsName } from "wagmi";

const DonationRequest = () => {
    const [step,setStep] = useState(1)
    const [ethDomainText,setEthDomainText] = useState("")
    const [isETHDomainButtonDisabled,setIsETHDomainButtonDisabled] = useState(true)
    const { data, isError, isLoading } = useEnsAddress({
        name: ethDomainText,
    })
    const ensAvatar = useEnsAvatar({ name: ethDomainText, chainId: (localStorage.getItem("Chain") ?? 1) as number})
    useEffect(()=>{
        setIsETHDomainButtonDisabled(isError || isLoading || data == null || data == undefined)
    },[isError,isLoading,data])
    
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
        setStep((prevValue)=>prevValue+1);
    };
    const nextStepToDonation = () =>{
        setStep((prevValue)=>prevValue+1);
    }
    const startDonation = () =>{

    }
    return(
        <div className="justify-center items-center h-1/2 my-32 grid grid-rows-2 gap-12">
            {step == 0 &&
                <>
                    <h1 className='text-black'>Verify Yourself Before Start Donation</h1>
                    <button type="button" onClick={signInButtonHandler} className="text-white w-full mx-auto bg-green-700 hover:bg-green-800 focus:outline-none focus:ring-4 focus:ring-green-300 font-medium rounded-full text-sm px-5 py-2.5 text-center mr-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">Verify</button>
                </> 
            }
            {step == 1 &&
                <>
                    <div>
                        <h1 className='text-black font-medium text-center'>Please enter your ETH domain</h1>
                        <img
                            src={ensAvatar.data || 'https://i.imgur.com/UhV7H97.jpeg'}
                            style={{ width: '2rem', height: '2rem', objectFit: 'cover' }}
                        />
                    </div>
                    <div>
                        <label htmlFor="first_name" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Domain Name</label>
                        <input type="text" id="first_name" value={ethDomainText} onChange={(event)=>setEthDomainText(event.target.value)} className="bg-gray-50 border w-96 border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="John" required/>
                    </div>
                    <button disabled={isETHDomainButtonDisabled} type="button" onClick={nextStepToDonation} className={`text-white w-full mx-auto ${isETHDomainButtonDisabled ? "bg-gray-500" : "bg-green-700 hover:bg-green-800 focus:outline-none focus:ring-4 focus:ring-green-300"} font-medium rounded-full text-sm px-5 py-2.5 text-center mr-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800`}>Next Step</button>
                </>
                
            }
            {step == 2 &&
                <>
                    <h1 className='text-black font-medium text-center'>Fill Out The Donation Form</h1>
                    <div>
                        <label htmlFor="first_name" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Destination Wallet Address</label>
                        <input type="text" id="first_name" className="bg-gray-50 border w-96 border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="John" required/>
                    </div>
                    <div>
                        <label htmlFor="first_name" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Title</label>
                        <input type="text" id="first_name" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="John" required/>
                    </div>
                    <div>
                        <label htmlFor="large-input" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Description</label>
                        <input type="text" id="large-input" className="block w-full p-4 text-gray-900 border border-gray-300 rounded-lg bg-gray-50 sm:text-md focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"/>
                    </div>
                    <div>
                        <label className="block mb-2 text-sm font-medium text-gray-900 dark:text-white" htmlFor="file_input">Upload Cover Image</label>
                        <input className="block w-full text-sm text-gray-900 border border-gray-300 rounded-lg cursor-pointer dark:text-gray-400 focus:outline-none dark:bg-gray-700 border-gray-600 placeholder-gray-400" id="file_input" type="file"/>
                    </div>
                    <div>
                        <button type="button" onClick={startDonation} className="text-white w-full mx-auto bg-green-700 hover:bg-green-800 focus:outline-none focus:ring-4 focus:ring-green-300 font-medium rounded-full text-sm px-5 py-2.5 text-center mr-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">Fire Up</button>
                    </div>
                </>
            }
        </div>
    )
}

export default DonationRequest;