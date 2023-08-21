// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

contract AuthenticationManager{

    mapping(address => bool) _isAuthenticated;
    address owner;

    constructor(){
        owner = msg.sender;
    }

    modifier onlyOwner(){
        require(msg.sender == owner,"Only owner can call this function");
        _;
    }

    function authenticate() public payable onlyOwner{
        require(!_isAuthenticated[msg.sender],"User Authenticated Before");
        _isAuthenticated[msg.sender] = true;
    }

    function isAuthenticated() public view returns(bool){
        return _isAuthenticated[msg.sender];
    }
}