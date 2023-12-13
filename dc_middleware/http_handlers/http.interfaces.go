package http_handlers

import "dc_middleware/ipfs_connector"

type HandlerData struct {
	connector    ipfs_connector.RpcConnector
	FileHashList []string
}
