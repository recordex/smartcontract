import { ethers } from "hardhat";
import { expect } from "chai";

// 参考: https://hardhat.org/tutorial/testing-contracts
describe("Policy Contract", function () {

  it("Should add a new policy correctly", async function () {
    const hardhatPolicy = await ethers.deployContract("Policy");

    await hardhatPolicy.addPolicy("resource1", "action1", true);

    const policyCount = await hardhatPolicy.getPolicy(1);
    expect(policyCount[0]).to.equal("resource1");
    expect(policyCount[1]).to.equal("action1");
    expect(policyCount[2]).to.equal(true);
  })

  it("Should emit PolicyAdded event when a policy is added", async function () {
    const hardhatPolicy = await ethers.deployContract("Policy");

    await expect(hardhatPolicy.addPolicy("resource2", "action2", false))
      .to.emit(hardhatPolicy, "PolicyAdded")
      .withArgs(1, "resource2", "action2", false);
  })
});
