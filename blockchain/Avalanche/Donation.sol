// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

contract Donation{

    address destinationWallet;
    address owner;

    modifier onlyOwner(){
        require(msg.sender == owner,"Only admin can call this function");
        _;
    }

    constructor(address _destinationWallet,address _owner){
        destinationWallet = _destinationWallet;
        owner = _owner;
    }

    function donate() payable  public {
        (bool success,) = payable(address(this)).call{value: msg.value}("");
        require(success,"Donation is not succesfull");
    }

    function getMoney() payable public onlyOwner{
        (bool success,) = payable(destinationWallet).call{value: balance()}("");
        require(success,"Payment is not succesfull");
    }

    function balance() public view returns (uint256) {
        return address(this).balance;
    }
}