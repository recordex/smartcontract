import {ethers} from "hardhat";
import {expect} from "chai";
import utils from "web3-utils";

describe('Record Contract', function () {

  it('デプロイ時にロールが正しく設定されているか', async function () {
    const [owner, user, anotherUser] = await ethers.getSigners();
    const hardhatRecord = await ethers.deployContract("Record", [owner.address]);

    // expect(await hardhatRecord.hasRole(utils.keccak256("DEFAULT_ADMIN_ROLE"), owner.address)).to.be.true;
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

  // it('should allow admin to grant hash addition role to others', async function () {
  //   const [owner, user, anotherUser] = await ethers.getSigners();
  //   const hardhatRecord = await ethers.deployContract("Record", [owner.address]);
  //
  //   await hardhatRecord.grantHasherRole(user.address);
  //   expect(await hardhatRecord.hasRole(utils.keccak256("HASHER_ROLE"), user.address)).to.be.true;
  // });
  //
  // it('should not allow duplicates hashes', async function () {
  //   await this.record.addHash(fileHash, { from: admin });
  //   await expectRevert(
  //     this.record.addHash(fileHash, { from: admin }),
  //     'This hash is already recorded.'
  //   );
  // });
  //
  // it('should allow admin to revoke hash addition role from others', async function () {
  //   await this.record.grantHasherRole(user, { from: admin });
  //   await this.record.revokeHasherRole(user, { from: admin });
  //   expect(await this.record.hasRole(Record.HASHER_ROLE, user)).to.be.false;
  // });
});
