import {ethers} from "hardhat";
import {expect} from "chai";
import utils from "web3-utils";

describe('Record Contract', function () {

  it('デプロイ時にロールが正しく設定されているか', async function () {
    const [owner, user, anotherUser] = await ethers.getSigners();
    const hardhatRecord = await ethers.deployContract("Record", [owner.address]);

    expect(await hardhatRecord.hasRole(utils.keccak256("HASHER_ROLE"), owner.address)).to.be.true;
    expect(await hardhatRecord.hasRole(utils.keccak256("HASHER_ROLE"), user.address)).to.be.false;
    expect(await hardhatRecord.hasRole(utils.keccak256("HASHER_ROLE"), anotherUser.address)).to.be.false;
  });

  it('owner アカウントによる addHash メソッドの実行が許可されているか', async function () {
    const [owner, user, anotherUser] = await ethers.getSigners();
    const hardhatRecord = await ethers.deployContract("Record", [owner.address]);

    const fileHash = '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef';

    await hardhatRecord.addHash(fileHash);
    expect(await hardhatRecord.isHashRecorded("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")).to.be.true;
    expect(await hardhatRecord.isHashRecorded("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdee")).to.be.false;
  });

  it('owner アカウント以外からの addHash メソッドの実行が許可されていないか', async function () {
    const [owner, user, anotherUser] = await ethers.getSigners();
    const hardhatRecord = await ethers.deployContract("Record", [owner.address]);

    const fileHash = '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef';


    await expect(hardhatRecord.connect(user).addHash(fileHash)).to.be.revertedWithCustomError(
      hardhatRecord,
      "AccessControlUnauthorizedAccount",
    );
    await expect(hardhatRecord.connect(anotherUser).addHash(fileHash)).to.be.revertedWithCustomError(
      hardhatRecord,
      "AccessControlUnauthorizedAccount",
    );
  });

  it('owner にロールの付与権限があるかどうか', async function () {
    const [owner, user, anotherUser] = await ethers.getSigners();
    const hardhatRecord = await ethers.deployContract("Record", [owner.address]);

    await hardhatRecord.grantHasherRole(user.address);
    expect(await hardhatRecord.hasRole(utils.keccak256("HASHER_ROLE"), user.address)).to.be.true;
  });

  it('ハッシュが重複して保存されないことの確認', async function () {
    const [owner, user, anotherUser] = await ethers.getSigners();
    const hardhatRecord = await ethers.deployContract("Record", [owner.address]);
    const fileHash = '0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef';

    await hardhatRecord.addHash(fileHash);
    await expect(hardhatRecord.addHash(fileHash)).to.be.revertedWith("This hash is already recorded.");
  });

  it('HASHER_ROLE の revoke 権限が owner にあるか', async function () {
  // it('should allow admin to revoke hash addition role from others', async function () {
    const [owner, user, anotherUser] = await ethers.getSigners();
    const hardhatRecord = await ethers.deployContract("Record", [owner.address]);

    await hardhatRecord.grantHasherRole(user.address);
    await hardhatRecord.revokeHasherRole(user.address);
    expect(await hardhatRecord.hasRole(utils.keccak256("HASHER_ROLE"), user)).to.be.false;
  });
});
