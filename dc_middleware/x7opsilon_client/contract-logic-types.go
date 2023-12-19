package x7opsilon_client

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractLogicData struct {
	abi    abi.ABI
	client *ethclient.Client
}

type FileUploaded struct {
	FileID      string
	DirectoryId string
	FileHash    string
}
