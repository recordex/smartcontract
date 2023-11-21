# smartcontract

## Description

Data will be stored in a JSON file. The JSON file will contain the following:

```json
{
  "users": [
    { "id": "user1", "name": "Alice", "password": "hash1" },
    { "id": "user2", "name": "Bob", "password": "hash2" }
  ],
  "files": [
    { "id": "file1", "name": "report.docx", "owner": "user1", "size": "2MB" },
    { "id": "file2", "name": "image.png", "owner": "user2", "size": "3MB" }
  ],
  "permissions": [
    { "id": "read", "description": "Read file contents" },
    { "id": "write", "description": "Write file contents" },
    { "id": "delete", "description": "Delete file" }
  ],
  "access_control_list": [
    {
      "user_id": "user1",
      "file_id": "file1",
      "permissions": ["read", "write"]
    },
    { "user_id": "user2", "file_id": "file1", "permissions": ["read"] },
    { "user_id": "user1", "file_id": "file2", "permissions": ["read"] },
    {
      "user_id": "user2",
      "file_id": "file2",
      "permissions": ["read", "write", "delete"]
    }
  ]
}
```

## About Hardhat

This project demonstrates a basic Hardhat use case. It comes with a sample contract, a test for that contract, and a script that deploys that contract.

Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat run scripts/deploy.ts
```
