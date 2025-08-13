package blockchain

import (
	"bytes"
	"chapter8/wallet"
	"encoding/gob"
)

type TxOutput struct {
	Value      int    // 金額
	PubKeyHash []byte // 被鎖定的公鑰哈希
}

type TxOutputs struct {
	Outputs []TxOutput
}

type TxInput struct {
	ID        []byte // 要花費的交易ID
	Out       int    // 交易中的第幾個輸出
	Signature []byte // 花這筆錢的簽章
	PubKey    []byte // 用來驗簽的公鑰
}

func NewTXOutput(value int, address string) *TxOutput {
	txo := &TxOutput{value, nil}
	txo.Lock([]byte(address))
	return txo
}

// 篩選 哪一些UTXO是我的
func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

// 把Output限制給接收位址
func (out *TxOutput) Lock(address []byte) {
	pubKeyHash := wallet.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]

	out.PubKeyHash = pubKeyHash
}

// 檢查接收人是否擁有該公鑰
func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

func (outs TxOutputs) Serialize() []byte {
	var buffer bytes.Buffer
	encode := gob.NewEncoder(&buffer)
	err := encode.Encode(outs)
	Handle(err)
	return buffer.Bytes()
}

func DeserializeOutputs(data []byte) TxOutputs {
	var outputs TxOutputs
	decode := gob.NewDecoder(bytes.NewReader(data))
	err := decode.Decode(&outputs)
	Handle(err)
	return outputs
}
