// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

import "./LilypadEventsUpgradeable.sol";
import "./LilypadCallerInterface.sol";

contract Lilypad is LilypadCallerInterface {
    address public bridgeAddress; // Variable for interacting with the deployed LilypadEvents contract
    LilypadEventsUpgradeable bridge;
    uint256 public lilypadFee; //=30000000000000000;
    mapping(uint=>string) public prompts;
    mapping(uint=>string) public results;
    event JobResult(string result);
    event ResultForUser(string wallet,string result);

    constructor(address _bridgeContractAddress) {
        bridgeAddress = _bridgeContractAddress;
        bridge = LilypadEventsUpgradeable(_bridgeContractAddress);
        uint fee = bridge.getLilypadFee(); // you can fetch the fee amount required for the contract to run also
        lilypadFee = fee;
    }

    /** Call the runLilypadJob() to generate a stable diffusion image from a text prompt*/
    function request(string calldata _prompt) external payable {
        require(msg.value >= lilypadFee, "Not enough to run Lilypad job");
        string memory spec = string(abi.encodePacked(
                "{",
                '"Engine": "Docker",',
                '"Verifier": "Noop",',
                '"Publisher": "Estuary",',
                '"PublisherSpec": {"Type": "Estuary"},',
                '"Docker": {"Image": "dogukangun/donationforuniversity:go_latest6", "EnvironmentVariables": ["WALLET_ADDRESS=',
                _prompt,
                '"]},',
                '"Language": {"JobContext": {}},',
                '"Wasm": {"EntryModule": {}},',
                '"Resources": {"GPU": ""},',
                '"Network": {"Type": "HTTP", "Domains": ["eth.public-rpc.com"]},',
                '"Timeout": 1800,',
                '"Outputs": [{"StorageSource": "IPFS", "Name": "outputs", "Path": "/outputs"}],',
                '"Deal": {"Concurrency": 1}',
                "}"
            ));
        uint id = bridge.runLilypadJob{value: lilypadFee}(address(this), spec, uint8(LilypadResultType.StdOut));
        require(id > 0, "job didn't return a value");
        prompts[id] = _prompt;
        emit JobResult(_prompt);
    }

    function lilypadFulfilled(address _from, uint _jobId,
        LilypadResultType _resultType, string calldata _result)
    external override {
        require(_from == address(bridge));

        // Save the CID against the caller's address
        results[_jobId] = _result;
        emit ResultForUser(prompts[_jobId],_result);
    }

    function lilypadCancelled(address _from, uint _jobId, string
        calldata _errorMsg) external override {

        // Do something if there's an error returned by the
        // LilypadEvents contract
        results[_jobId] = _errorMsg;
        emit JobResult(_errorMsg);
    }
}