#!/bin/bash

# artifacts/contracts ディレクトリ内の.solファイルのリストを取得
contracts_path="artifacts/contracts/*.sol"
contract_files=( $contracts_path )

# コントラクトファイルが存在するかどうかをチェック
if [ ${#contract_files[@]} -eq 0 ]; then
  echo "Error: コントラクトファイルがディレクトリに存在しません。"
  exit 1
fi

# 各コントラクトファイルに対して処理を実行
for file_path in "${contract_files[@]}"; do
  # ファイル名から拡張子を取り除き、コントラクト名を取得
  contract_name=$(basename "$file_path" .sol)

  # 入力 JSON ファイルのパスを定義
  json_file_path="artifacts/contracts/${contract_name}.sol/${contract_name}.json"

  # ファイルが存在するかチェック
  if [ ! -f "$json_file_path" ]; then
      echo "Error: ファイル ${json_file_path} が存在しません。"
      continue
  fi

  # abi ファイルのパスを定義
  abi_file_path="artifacts/contracts/${contract_name}.sol/${contract_name}.abi"

  # jqを使用してJSONファイルからabiセクションを抽出し、出力ファイルに保存
  jq '.abi' "$json_file_path" > "$abi_file_path"

  # 結果のメッセージを表示
  echo "abi が ${abi_file_path} に抽出されました。"

  pkg_name=$(echo "$contract_name" | tr '[:upper:]' '[:lower:]')

  # 参考: https://geth.ethereum.org/docs/tools/abigen
  mkdir -p pkg/"${pkg_name}"
  abigen --abi="${abi_file_path}" --pkg="${pkg_name}" --out=pkg/"${pkg_name}"/"${pkg_name}".go

  echo "pkg/${pkg_name}/${pkg_name}.go が生成されました。"
done
