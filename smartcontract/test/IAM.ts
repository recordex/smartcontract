// import { ethers } from "hardhat";
// import { expect } from "chai";
//
// // 参考: https://hardhat.org/tutorial/testing-contracts
// describe("User Contract", function () {
//   async function deployIAMContract() {
//     const [owner] = await ethers.getSigners();
//     const hardhatPolicy = await ethers.deployContract("Policy");
//     const hardhatUser = await ethers.deployContract("User");
//     const hardhatIAM = await ethers.deployContract("IAM", [hardhatPolicy.target, hardhatUser.target]);
//
//     return { owner, hardhatPolicy, hardhatUser, hardhatIAM };
//   }
// })
