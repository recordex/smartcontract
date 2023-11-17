package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/recordex/smartcontract/pkg/record"
)

const RecordContractAddress = "0x474Fc6f8BAF5B4E26E238efbD1cA88FCd462F330"
const PrivateKeyAddress = "8784fb5064d9d32a752ac822b2cc397d40300240efe2cb19acdb8f6f662b54e3"
const NetworkID = 1337

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545/")
	if err != nil {
		log.Fatalf("イーサリアムノードへの接続に失敗しました。: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(PrivateKeyAddress)
	if err != nil {
		log.Fatalf("プライベートキーのロードに失敗しました。: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("パブリックキーの ECDSA へのキャストに失敗しました。")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(NetworkID))
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(RecordContractAddress)
	instance, err := record.NewRecord(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fileName := "test.txt"
	fileHash := "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"
	// 文字列をバイトスライスにデコード
	encodedFileHash, err := hex.DecodeString(fileHash[:64]) // 最初の32バイトのみをデコードする
	if err != nil {
		log.Fatal(err)
	}
	tx, err := instance.AddFile(auth, fileName, [32]byte(encodedFileHash))
	if err != nil {
		log.Fatalf("AddFile メソッドの実行に失敗しました。: %v", err)
	}

	log.Printf("トランザクションが送信されました。transactionID -> %s", tx.Hash().Hex())

	opts := &bind.CallOpts{
		From: fromAddress,
	}
	// Recordコントラクトのfilesマップから内容を取得します。
	fileMetadata, err := instance.GetFileMetadataHistory(opts, fileName)
	if err != nil {
		log.Fatalf("GetFileMetadataHistory の実行に失敗しました。: %v", err)
	}

	// ファイルのメタデータを表示します。
	for _, metadata := range fileMetadata {
		fmt.Printf("Hash: %x\n", metadata.Hash)
		fmt.Printf("Created by: %s\n", metadata.CreatedBy.Hex())
		fmt.Printf("Created at: %s\n", time.Unix(metadata.CreatedAt.Int64(), 0).UTC())
	}
}
