import { ethers } from "hardhat";
import { expect } from "chai";

// 参考: https://hardhat.org/tutorial/testing-contracts
describe("User Contract", function () {

  it("Should register a new user correctly", async () => {
    const [owner] = await ethers.getSigners();
    const hardhatUser = await ethers.deployContract("User");

    await hardhatUser.registerUser("alice");

    const [username, policies] = await hardhatUser.getUser(owner.address);
    expect(username).to.equal("alice");
    expect(policies.length).to.equal(0); // 新規ユーザーなのでポリシーはない
  });

  it("Should emit UserRegistered event when a user is registered", async () => {
    const [owner] = await ethers.getSigners();
    const hardhatUser = await ethers.deployContract("User");

    await expect(hardhatUser.registerUser("bob"))
      .to.emit(hardhatUser, "UserRegistered")
      .withArgs(owner.address, "bob");
  });

  it("Should allow assignment of policy ID to a user", async () => {
    const [owner] = await ethers.getSigners();
    const hardhatUser = await ethers.deployContract("User");

    await hardhatUser.registerUser("charlie");
    await hardhatUser.assignPolicy(1);

    const [, policies] = await hardhatUser.getUser(owner.address);
    expect(policies.length).to.equal(1);
    expect(policies[0]).to.equal(1);
  });

  it("Should emit PolicyAssigned event when a policy ID is assigned", async () => {
    const [owner] = await ethers.getSigners();
    const hardhatUser = await ethers.deployContract("User");

    await hardhatUser.registerUser("dave");
    await expect(hardhatUser.assignPolicy(2))
      .to.emit(hardhatUser, "PolicyAssigned")
      .withArgs(owner.address, 2);
  });
})
