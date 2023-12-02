package ipfs_connector

type RpcConnector interface {
	Add(filePath string, pin bool) *AddResult
}
