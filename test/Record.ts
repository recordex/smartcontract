import { ethers } from 'hardhat';
import { expect } from 'chai';
import utils from 'web3-utils';
import { anyValue } from '@nomicfoundation/hardhat-chai-matchers/withArgs';
import { Record } from '../typechain-types';

// 参考: https://hardhat.org/tutorial/testing-contracts

describe('Record Contract', function () {
  it('正常系: デプロイ時にロールが正しく設定されているか', async function () {
    const [deployer, user, anotherUser] = await ethers.getSigners();
    const hardhatRecord = await ethers.deployContract('Record');

    expect(
      await hardhatRecord.hasRole(
        utils.keccak256('READ_ROLE'),
        deployer.address,
      ),
    ).to.be.true;
    expect(
      await hardhatRecord.hasRole(
        utils.keccak256('WRITE_ROLE'),
        deployer.address,
      ),
    ).to.be.true;
    expect(
      await hardhatRecord.hasRole(
        utils.keccak256('DELETE_ROLE'),
        deployer.address,
      ),
    ).to.be.true;
    expect(
      await hardhatRecord.hasRole(utils.keccak256('READ_ROLE'), user.address),
    ).to.be.false;
    expect(
      await hardhatRecord.hasRole(
        utils.keccak256('READ_ROLE'),
        anotherUser.address,
      ),
    ).to.be.false;
  });

  it('正常系: deployer アカウントが role の付与を行えるか', async function () {
    const [, user] = await ethers.getSigners();
    const hardhatRecord = await ethers.deployContract('Record');

    await hardhatRecord.grantRole(utils.keccak256('READ_ROLE'), user.address);
    expect(
      await hardhatRecord.hasRole(utils.keccak256('READ_ROLE'), user.address),
    ).to.be.true;
  });

  it('正常系: deployer アカウントが role の剥奪を行えるか', async function () {
    const [, user] = await ethers.getSigners();
    const hardhatRecord = await ethers.deployContract('Record');

    await hardhatRecord.grantRole(utils.keccak256('READ_ROLE'), user.address);
    await hardhatRecord.revokeRole(utils.keccak256('READ_ROLE'), user.address);
    expect(
      await hardhatRecord.hasRole(utils.keccak256('READ_ROLE'), user.address),
    ).to.be.false;
  });

  it('正常系: 書き込み権限があるアカウントがファイルの追加を行えるか', async function () {
    const [deployer] = await ethers.getSigners();
    const hardhatRecord = await ethers.deployContract('Record');

    const fileName = 'test.txt';
    const fileHash =
      '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef';

    // anyValue は、withArgs の引数の値が何であっても良いことを表す。
    // 参照: https://hardhat.org/hardhat-chai-matchers/docs/overview
    // anyValue にしているのは timestamp がテスト実行時によって異なった値になってしまうため
    await expect(
      await hardhatRecord.addFile(
        fileName,
        fileHash,
        ethers.encodeBytes32String('0'),
      ),
    )
      .to.emit(hardhatRecord, 'FileAdded')
      .withArgs(fileName, fileHash, deployer.address, anyValue);

    const newFileHash =
      '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdee';
    await expect(await hardhatRecord.addFile(fileName, newFileHash, fileHash))
      .to.emit(hardhatRecord, 'FileAdded')
      .withArgs(fileName, newFileHash, deployer.address, anyValue);
  });

  it('正常系: 削除権限があるアカウントがファイルの削除を行えるか', async function () {
    const [deployer] = await ethers.getSigners();
    const hardhatRecord = await ethers.deployContract('Record');

    const fileName = 'test.txt';
    const fileHash =
      '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef';

    await hardhatRecord.addFile(
      fileName,
      fileHash,
      ethers.encodeBytes32String('0'),
    );
    await expect(await hardhatRecord.deleteFile(fileName))
      .to.emit(hardhatRecord, 'FileDeleted')
      .withArgs(fileName, deployer.address);
  });

  it('正常系: ファイルのハッシュ値の履歴を取得できるか', async function () {
    const hardhatRecord = await ethers.deployContract('Record');

    const fileName = 'test.txt';
    const fileHash =
      '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef';

    await hardhatRecord.addFile(
      fileName,
      fileHash,
      ethers.encodeBytes32String('0'),
    );
    let fileHashHistory: Record.FileMetadataStructOutput[];

    fileHashHistory = await hardhatRecord.getFileMetadataHistory(fileName);
    expect(fileHashHistory.length).to.be.equal(1);
    expect(fileHashHistory[0].hash).to.be.equal(fileHash);

    await hardhatRecord.deleteFile(fileName);
    fileHashHistory = await hardhatRecord.getFileMetadataHistory(fileName);
    expect(fileHashHistory.length).to.be.equal(0);
  });

  it('異常系: 正しい1つ前のファイルのハッシュ値を指定できていない場合にファイルの追加を行えないか', async function () {
    const hardhatRecord = await ethers.deployContract('Record');

    const fileName = 'test.txt';
    const fileHash =
      '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef';
    const newFileHash =
      '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdee';

    await hardhatRecord.addFile(
      fileName,
      fileHash,
      ethers.encodeBytes32String('0'),
    );
    await expect(
      hardhatRecord.addFile(
        fileName,
        newFileHash,
        ethers.encodeBytes32String('0'),
      ),
    ).to.be.revertedWith('Previous file hash does not match');
  });

  it('異常系: addFile 関数の fileHash 引数と previousFileHash 引数が同じ場合にファイルの追加を行えないか', async function () {
    const hardhatRecord = await ethers.deployContract('Record');

    const fileName = 'test.txt';
    const fileHash =
      '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef';

    await hardhatRecord.addFile(
      fileName,
      fileHash,
      ethers.encodeBytes32String('0'),
    ),
      await expect(
        hardhatRecord.addFile(fileName, fileHash, fileHash),
      ).to.be.revertedWith(
        'New file hash must be different from previous file hash',
      );
  });

  it('異常系: 書き込み権限がないアカウントがファイルの追加を行えないか', async function () {
    const [, user] = await ethers.getSigners();
    const hardhatRecord = await ethers.deployContract('Record');

    const fileName = 'test.txt';
    const fileHash =
      '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef';

    await expect(
      hardhatRecord
        .connect(user)
        .addFile(fileName, fileHash, ethers.encodeBytes32String('0')),
    ).to.be.revertedWith('Caller does not have write permission');
  });

  it('異常系: 削除権限がないアカウントがファイルの削除を行えないか', async function () {
    const [, user] = await ethers.getSigners();
    const hardhatRecord = await ethers.deployContract('Record');

    const fileName = 'test.txt';
    const fileHash =
      '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef';

    await hardhatRecord.addFile(
      fileName,
      fileHash,
      ethers.encodeBytes32String('0'),
    );
    await expect(
      hardhatRecord.connect(user).deleteFile(fileName),
    ).to.be.revertedWith('Caller does not have delete permission');
  });
});
