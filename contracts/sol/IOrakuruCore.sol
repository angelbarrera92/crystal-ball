// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

interface IOrakuruCore {
    enum Type {
        MostFrequent,
        Median
    }

    event Requested(
        bytes32 indexed requestId,
        string dataSource,
        string selector,
        address indexed callbackAddr,
        uint256 executionTimestamp,
        uint256 fulfillmentTimestamp,
        Type aggrType
    );

    event Canceled(bytes32 indexed requestId);

    function submitResult(bytes32 _requestId, string memory _result) external;
}
