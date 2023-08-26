// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;
import "./Donation.sol";

contract DonationFactory{

    mapping(address=>Donation) donations;
    address owner; //DonationManagerAddress

    constructor(){
        owner = msg.sender;
    }

    modifier onlyOwner(){
        require(msg.sender == owner,"Only admin can call this function");
        _;
    }

    function startDonation(address destinationWallet) public onlyOwner{
        Donation donation = new Donation(destinationWallet,owner);
        address newDonationAddress = address(donation);
        donations[newDonationAddress] = donation;
    }

}