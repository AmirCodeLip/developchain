package x7opsilon_client

import "math/big"

type ContractLogic interface {
	GetUploadedFiles(lastBlock *big.Int) (uint64, []FileUploaded)
}
