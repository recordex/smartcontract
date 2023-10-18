// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Policy {

    // ポリシーのデータ構造を定義
    struct PolicyDetail {
        string resource;
        string action;
        bool isAllowed;  // trueなら許可、falseなら拒否
    }

    // 各ポリシーIDに対応するポリシーの詳細を保存
    mapping(uint256 => PolicyDetail) public policies;

    // ポリシーのカウンター（新しいポリシーのIDとしても使用）
    uint256 private policyCounter;

    // ポリシーの追加イベント
    event PolicyAdded(uint256 policyId, string resource, string action, bool isAllowed);

    // ポリシーを追加する関数
    function addPolicy(string memory _resource, string memory _action, bool _isAllowed) external returns (uint256) {
        policyCounter++;
        // policies は policyCounter から引いてくるのではなくて、自分でつけた policyName から引いて来れるようにしたほうがよさそう
        policies[policyCounter] = PolicyDetail(_resource, _action, _isAllowed);
        emit PolicyAdded(policyCounter, _resource, _action, _isAllowed);
        return policyCounter;
    }

    // 特定のポリシーIDに関連するポリシーの詳細を取得する関数
    function getPolicy(uint256 _policyId) external view returns (string memory, string memory, bool) {
        PolicyDetail memory policy = policies[_policyId];
        return (policy.resource, policy.action, policy.isAllowed);
    }
}
