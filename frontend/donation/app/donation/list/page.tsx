"use client"
import DonationItem from "@/components/donation_item/DonationItem";
import CustomModal from "@/components/modal/CustomModal";
import { useState } from "react";

const DonationList = () =>{
    const [isModalVisible,setIsModalVisible] = useState(false)

    const itemOnClick = () =>{
        setIsModalVisible(true)
    }

    const modelOnCloseClicked = () => {
        setIsModalVisible(false)
    }

    const modelOnAcceptClicked = () => {
        setIsModalVisible(false)
    }

    return(
        <div className="bg-white py-6 sm:py-8 lg:py-12">
            <div className="mx-auto max-w-screen-2xl px-4 md:px-8">
                <div className="mb-6 flex items-end justify-between gap-4">
                    <div className="grid gap-x-4 gap-y-8 sm:grid-cols-2 md:gap-x-6 lg:grid-cols-3 xl:grid-cols-4">
                        <DonationItem onClick={itemOnClick}/>
                        <DonationItem onClick={itemOnClick}/>
                        <DonationItem onClick={itemOnClick}/>
                        <DonationItem onClick={itemOnClick}/>
                        <DonationItem onClick={itemOnClick}/>

                        {isModalVisible &&
                            <CustomModal onAccepted={modelOnAcceptClicked} onClose={modelOnCloseClicked}/>
                        }
                    </div>
                </div>
            </div>
        </div>
    )
}

export default DonationList;