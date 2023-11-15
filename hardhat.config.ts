import { HardhatUserConfig } from 'hardhat/config';
import '@nomicfoundation/hardhat-toolbox';

const INFURA_API_KEY = process.env.INFURA_API_KEY;
if (!INFURA_API_KEY) {
  throw new Error(
    '環境変数 INFURA_API_KEY が設定されていません。INFURA_API_KEY -> ' +
      INFURA_API_KEY,
  );
}

const SEPOLIA_PRIVATE_KEY = process.env.SEPOLIA_PRIVATE_KEY;
if (!SEPOLIA_PRIVATE_KEY) {
  throw new Error(
    '環境変数 SEPOLIA_PRIVATE_KEY が設定されていません。SEPOLIA_PRIVATE_KEY -> ' +
      SEPOLIA_PRIVATE_KEY,
  );
}

// 参考: https://hardhat.org/tutorial/deploying-to-a-live-network
const config: HardhatUserConfig = {
  solidity: '0.8.20',
  networks: {
    sepolia: {
      url: `https://sepolia.infura.io/v3/${INFURA_API_KEY}`,
      accounts: [SEPOLIA_PRIVATE_KEY],
    },
  },
};

export default config;
