package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

/*
	example

假設 Bob 要給 Alice 5 BTC

 1. 交易輸入Input
    功能:  證明Bob有權利花費某個UTXO
    內容:
    * Bob 選擇一個或多個自己擁有的UTXO為輸入。
    * 為了花費這個UTXO，Bob 必須提供一個簽章(用Bob的私鑰對整個交易內容進行簽屬產生的)。 => 像是 我，Bob，同意並授權我給Alice這筆交易的內容，並且有權力花費其中的輸入
    * 驗證者(礦工)會用Bob的公鑰來驗證簽章。

 2. 交易輸出Output
    功能: 定義新產的UTXO，並將其所定給接收方
    內容:
    * 此時會有兩個輸出(Payment 跟 Recharge)
    * Payment 的 pubKey 為 Alice 的公鑰(只有擁有Alice的私鑰才可解開)
    * 而Recharge 的 pubKey 為 Bob的公鑰
*/
type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

type TxOutput struct {
	Value  int
	PubKey string
}

type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}
	input := TxInput{[]byte{}, -1, data}
	output := TxOutput{100, to}
	tx := Transaction{[]byte{}, []TxInput{input}, []TxOutput{output}}
	tx.SetID()
	return &tx
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte
	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	Handle(err)
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func NewTransaction(from, to string, amount int, chain *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOutpus := chain.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("Error: not enough funds")
	}
	// 把找到的 Output 包成 Input
	for txid, outs := range validOutpus {
		txID, err := hex.DecodeString(txid)
		Handle(err)

		for _, out := range outs {
			input := TxInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TxOutput{amount, to})

	if acc > amount {
		outputs = append(outputs, TxOutput{acc - amount, from})
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
