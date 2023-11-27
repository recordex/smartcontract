// SPDX-License-Identifier: apache-2.0
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract Record is AccessControl {
    // 権限を示す役割のバイト列を定義
    bytes32 public constant READ_ROLE = keccak256("READ_ROLE");
    bytes32 public constant WRITE_ROLE = keccak256("WRITE_ROLE");
    bytes32 public constant DELETE_ROLE = keccak256("DELETE_ROLE");

    // ファイルのメタデータを格納するための構造体
    struct FileMetadata {
        bytes32 hash;
        address created_by;
        uint256 created_at;
    }
    // ファイル名をキーとして、ファイルのメタデータを格納するマッピング
    mapping(string => FileMetadata[]) public files;

    // イベント定義
    event FileAdded(string fileName, bytes32 fileHash, address createdBy, uint256 createdAt);
    event FileDeleted(string fileName, address deletedBy);

    // コントラクトをデプロイするときの初期設定
    constructor() {
        // メッセージ送信者（デプロイヤー）に管理者役割を付与
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);

        // 初期管理者に全役割を付与
        _grantRole(READ_ROLE, msg.sender);
        _grantRole(WRITE_ROLE, msg.sender);
        _grantRole(DELETE_ROLE, msg.sender);
    }

    receive() external payable {}

    fallback() external payable {}

    // ファイルを追加するための関数（書き込み権限が必要）
    function addFile(string memory fileName, bytes32 fileHash, bytes32 previousFileHash) public {
        require(hasRole(WRITE_ROLE, msg.sender), "Caller does not have write permission");

        // ファイルが存在し、最後のファイルのハッシュが引数で指定されたハッシュと一致する場合にのみ追加
        uint256 metadataListLength = files[fileName].length;
        require(metadataListLength == 0 || files[fileName][metadataListLength - 1].hash == previousFileHash, "Previous file hash does not match");
        require(fileHash != previousFileHash, "New file hash must be different from previous file hash");

        files[fileName].push(FileMetadata(fileHash, msg.sender, block.timestamp));
        emit FileAdded(fileName, fileHash, msg.sender, block.timestamp);
    }

    // ファイルを削除するための関数（削除権限が必要）
    function deleteFile(string memory fileName) public {
        require(hasRole(DELETE_ROLE, msg.sender), "Caller does not have delete permission");
        uint256 metadataListLength = files[fileName].length;
        require(metadataListLength > 0, "File does not exist");
        delete files[fileName];
        emit FileDeleted(fileName, msg.sender);
    }

    // ファイルのメタデータの一覧を取得するための関数
    function getFileMetadataHistory(string memory fileName) public view returns (FileMetadata[] memory) {
        return files[fileName];
    }

    // 役割を付与するための関数
    function grantRole(bytes32 role, address account) public override {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender), "Caller is not an admin");
        super.grantRole(role, account);
    }

    // 役割を取り消すための関数
    function revokeRole(bytes32 role, address account) public override {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender), "Caller is not an admin");
        super.revokeRole(role, account);
    }
}
