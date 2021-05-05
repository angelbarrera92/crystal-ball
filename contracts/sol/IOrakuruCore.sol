// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

interface IOrakuruCore {
    enum Type {
        MostFrequent,
        Median
    }

    struct Response {
        bytes32 id;
        bytes32 requestId;
        bytes result;
        address submittedBy;
        uint256 submittedAt;
    }

    struct Request {
        bytes32 id;
        string dataSource;
        string selector;
        address callbackAddr;
        uint256 executionTimestamp;
        bool isFulfilled;
        Response[] responses;
    }

    function requests(bytes32) external view returns (Request memory);

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
