package x7opsilon_client

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewContractLogic(rpcUrl string) ContractLogic {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	entries, _ := os.Open("./static_data/ContractLogic.json")
	contractAbi, _ := abi.JSON(entries)
	var contractLogic ContractLogic
	data := ContractLogicData{contractAbi, client}
	contractLogic = &data
	return contractLogic
}

func (data *ContractLogicData) GetUploadedFiles(lastBlock *big.Int) (uint64, []FileUploaded) {
	var eventResults []FileUploaded
	var blockNumber uint64 = 0
	query := ethereum.FilterQuery{
		FromBlock: lastBlock,
		ToBlock:   nil,
		Addresses: []common.Address{common.HexToAddress("0xC3B9caAf849fF0764Ec704eE2C3743D2b8B0a0D1")},
	}
	l, _ := data.client.FilterLogs(context.Background(), query)
	for _, element := range l {
		blockNumber = element.BlockNumber
		data, err := data.abi.Unpack("FileUploaded", element.Data)
		//if err != nil {
		//	log.Fatal(err)
		//}
		if err == nil {
			fileId, _ := data[0].(string)
			directoryId, _ := data[1].(string)
			fileHash, _ := data[2].(string)
			//fmt.Println(data)
			eventResults = append(eventResults, FileUploaded{fileId, directoryId, fileHash})
		}
	}
	return blockNumber, eventResults
}
