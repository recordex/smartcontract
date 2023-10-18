// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IPolicy {
    function getPolicy(uint256 _policyId) external view returns (string memory, string memory, bool);
}

interface IUser {
    function getUser(address _userAddress) external view returns (string memory, uint256[] memory);
}

contract IAM {

    IPolicy public policyContract;
    IUser public userContract;

    constructor(address _policyContractAddress, address _userContractAddress) {
        policyContract = IPolicy(_policyContractAddress);
        userContract = IUser(_userContractAddress);
    }

    // 特定のユーザーが特定のリソース、アクションにアクセスできるかを確認する関数
    function canAccess(address _userAddress, string memory _resource, string memory _action) external view returns (bool) {
        (, uint256[] memory policyIds) = userContract.getUser(_userAddress);

        for (uint i = 0; i < policyIds.length; i++) {
            (string memory resource, string memory action, bool isAllowed) = policyContract.getPolicy(policyIds[i]);

            if (keccak256(bytes(resource)) == keccak256(bytes(_resource)) && keccak256(bytes(action)) == keccak256(bytes(_action))) {
                return isAllowed;
            }
        }

        return false; // ポリシーが一致しない場合、アクセスは拒否されます
    }
}
