// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract Record is AccessControl {
    bytes32 public constant ADMIN_ROLE = keccak256("DEFAULT_ADMIN_ROLE");
    // 定義: ハッシュ値を追加するための役割
    bytes32 public constant HASHER_ROLE = keccak256("HASHER_ROLE");

    // ハッシュ値の記録
    mapping(bytes32 => bool) private _hashes;

    // コントラクトをデプロイするときの初期設定
    constructor(address admin) {
        require(admin != address(0), "Invalid admin address");

        _grantRole(ADMIN_ROLE, admin); // 初期管理者に管理者役割を付与
        _grantRole(HASHER_ROLE, admin);        // ハッシュ追加役割も初期管理者に付与
    }

    // ハッシュ値を追加する関数 (特定の役割を持つアカウントのみ実行可能)
    function addHash(bytes32 fileHash) external onlyRole(HASHER_ROLE) {
        require(!_hashes[fileHash], "This hash is already recorded.");
        _hashes[fileHash] = true;
    }

    // ハッシュ値が存在するか確認する関数
    function isHashRecorded(bytes32 fileHash) external view returns (bool) {
        return _hashes[fileHash];
    }

    // ハッシュ値追加の権限を新しいアカウントに付与する関数 (管理者のみ実行可能)
    function grantHasherRole(address account) external onlyRole(DEFAULT_ADMIN_ROLE) {
        grantRole(HASHER_ROLE, account);
    }

    // ハッシュ値追加の権限をアカウントから剥奪する関数 (管理者のみ実行可能)
    function revokeHasherRole(address account) external onlyRole(DEFAULT_ADMIN_ROLE) {
        revokeRole(HASHER_ROLE, account);
    }
}
