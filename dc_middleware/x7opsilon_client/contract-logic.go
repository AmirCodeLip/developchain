package x7opsilon_client

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
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
	contractLogic = data
	//query := ethereum.FilterQuery{
	//	FromBlock: nil,
	//	ToBlock:   nil,
	//	Addresses: []common.Address{common.HexToAddress("0x4ADd1cf81038D78c8b06c93e39B139a326465E48")},
	//}
	//l, _ := client.FilterLogs(context.Background(), query)
	//for _, element := range l {
	//	data, err := contractAbi.Unpack("FileUploaded", element.Data)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(data)
	//}
	return contractLogic
}

func getUploadedFiles() {
	
}
