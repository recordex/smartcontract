// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract User {

    struct UserDetails {
        string username;
        uint256[] policyIds;
    }

    mapping(address => UserDetails) private users;

    event UserRegistered(address userAddress, string username);
    event PolicyAssigned(address userAddress, uint256 policyId);

    // ユーザーを登録する関数
    function registerUser(string memory _username) external {
        require(bytes(users[msg.sender].username).length == 0, "User already registered");

        users[msg.sender].username = _username;

        emit UserRegistered(msg.sender, _username);
    }

    // ユーザーにポリシーIDを関連付ける関数
    function assignPolicy(uint256 _policyId) external {
        require(bytes(users[msg.sender].username).length != 0, "User not registered");

        users[msg.sender].policyIds.push(_policyId);

        emit PolicyAssigned(msg.sender, _policyId);
    }

    // 特定のアドレスのユーザー情報を取得する関数
    function getUser(address _userAddress) external view returns (string memory, uint256[] memory) {
        UserDetails memory user = users[_userAddress];
        return (user.username, user.policyIds);
    }
}
