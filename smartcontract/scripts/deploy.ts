import { ethers } from "hardhat";

// 参考: https://hardhat.org/hardhat-runner/docs/guides/deploying
async function main() {
  // Policyコントラクトのデプロイ
  const policy = await ethers.deployContract("Policy");
  await policy.waitForDeployment();
  console.log("Policy contract deployed to:", policy.target);

  // Userコントラクトのデプロイ
  const user = await ethers.deployContract("User");
  await user.waitForDeployment();
  console.log("User contract deployed to:", user.target);

  // IAMコントラクトのデプロイ
  // PolicyコントラクトとUserコントラクトのアドレスを引数として渡します
  const iam = await ethers.deployContract("IAM", [policy.target, user.target]);
  await iam.waitForDeployment();
  console.log("IAM contract deployed to:", iam.target);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
