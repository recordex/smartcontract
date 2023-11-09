import { ethers } from 'hardhat';

// 参考: https://hardhat.org/hardhat-runner/docs/guides/deploying
async function main() {
  // Record コントラクトのデプロイ
  const record = await ethers.deployContract('Record');
  await record.waitForDeployment();
  console.log('Record contract deployed to:', record.target);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
