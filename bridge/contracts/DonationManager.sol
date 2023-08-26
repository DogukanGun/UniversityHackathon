// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;
import "./DonationFactory.sol";

contract DonationManager{

    struct DonationInfo {
        uint256 ID;
        string title;
        string description;
        address owner;
        address contractAddress;
        string coverImageCID;
        bool isApproved;
    }
    event ApplyForDonation(address userWalletAddress);

    DonationInfo[] donations;
    address public donationFactoryAddress;
    address bridgeAddress;

    constructor(address _bridgeAddress){
        donationFactoryAddress = address(new DonationFactory());
        bridgeAddress = _bridgeAddress;
    }

    modifier onlyBridge() {
        require(msg.sender == bridgeAddress,"Only bridge can run this function");
        _;
    }

    function applyForDonation() public {
        emit ApplyForDonation(msg.sender);
    }

    function getAllDonations() public view  returns(DonationInfo[] memory){
        return donations;
    }

    function getDonation(uint256 donationID) public view returns(DonationInfo memory){
        return donations[donationID];
    }

    function startDonation(
        address destionationWallet,
        string calldata title,
        string calldata description,
        address owner,
        string calldata coverImageCID
    ) external onlyBridge{
        bytes memory payload = abi.encodeWithSignature("startDonation(address,address)",destionationWallet,owner);
        (bool success,bytes memory returnData) = donationFactoryAddress.call(payload);
        require(success);
        (address returnValue) = abi.decode(returnData, (address));
        donations.push(DonationInfo(donations.length,title,description,owner,returnValue,coverImageCID,false));
    }

    function approveDonation(address approvedAddress) external {
        for (uint256 index; index<donations.length; index++)
        {
            if(donations[index].owner == approvedAddress){
                donations[index].isApproved = true;
            }
        }
    }
}