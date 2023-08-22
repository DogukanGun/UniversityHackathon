// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

import "https://github.com/bacalhau-project/lilypad-v0/blob/main/hardhat/contracts/LilypadEventsUpgradeable.sol";
import "https://github.com/bacalhau-project/lilypad-v0/blob/main/hardhat/contracts/LilypadCallerInterface.sol";

contract Lilypad is LilypadCallerInterface {
    address public bridgeAddress; // Variable for interacting with the deployed LilypadEvents contract
    LilypadEventsUpgradeable bridge;
    uint256 public lilypadFee; //=30000000000000000;
    mapping(uint=>string) public prompts;
    mapping(uint=>string) public results;

    event Completed(uint result);
    event Success(string result);
    event ErrorRes(string result);

    constructor(address _bridgeContractAddress) {
        bridgeAddress = _bridgeContractAddress;
        bridge = LilypadEventsUpgradeable(_bridgeContractAddress);
        uint fee = bridge.getLilypadFee(); // you can fetch the fee amount required for the contract to run also
        lilypadFee = fee;
    }

    string constant specStart = '{'
        '"Engine": "Docker",'
        '"Verifier": "Noop",'
        '"Publisher": "Estuary",'
        '"PublisherSpec": {"Type": "Estuary"},'
        '"Docker": {'
        '"Image": "dogukangun/donationforuniversity:python",'
        '"Entrypoint": ["python", "worker.py"';

    string constant specEnd =
      '"]},'
      '"Resources": {"GPU": ""},'
      '"Outputs": [{"Name": "outputs", "StorageSource":"IPFS","Path": "/outputs"}],'
      '"Deal": {"Concurrency": 1}'
      '}';

      /** Call the runLilypadJob() to generate a stable diffusion image from a text prompt*/
    function request(string calldata _prompt) external payable {
        require(msg.value >= lilypadFee, "Not enough to run Lilypad job");
        string memory spec = string(abi.encodePacked(       
            "{",
            '"Engine": "Docker",',
            '"Verifier": "Noop",',
            '"Publisher": "Estuary",',
            '"PublisherSpec": {"Type": "Estuary"},',
            '"Docker": {"Image": "dogukangun/donationforuniversity:python3", "EnvironmentVariables": ["WALLET_ADDRESS=',
            _prompt,
            '"]},',
            '"Language": {"JobContext": {}},',
            '"Wasm": {"EntryModule": {}},',
            '"Resources": {"GPU": ""},',
            '"Timeout": 1800,',
            '"Outputs": [{"StorageSource": "IPFS", "Name": "outputs", "Path": "/outputs"}],',
            '"Deal": {"Concurrency": 1}',
            "}"
        ));      
        uint id = bridge.runLilypadJob{value: lilypadFee}(address(this), spec, uint8(LilypadResultType.StdOut));
        require(id > 0, "job didn't return a value");
        prompts[id] = _prompt;
        emit Completed(id);
    }

    function lilypadFulfilled(address _from, uint _jobId,
        LilypadResultType _resultType, string calldata _result)
        external override {
        // Do something when the LilypadEvents contract returns
        // results successfully

        results[_jobId] = _result;
        emit Success(_result);
    }

    function lilypadCancelled(address _from, uint _jobId, string
          calldata _errorMsg) external override {
          // Do something if there's an error returned by the
          // LilypadEvents contract
        results[_jobId] = _errorMsg;
        emit ErrorRes(_errorMsg);
    }
}