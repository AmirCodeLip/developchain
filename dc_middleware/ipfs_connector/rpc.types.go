package ipfs_connector

type RpcConnectorData struct {
	url string
}

type AddResult struct {
	Name string `json:"Name"`
	Hash string `json:"Hash"`
	Size string `json:"Size"`
}
