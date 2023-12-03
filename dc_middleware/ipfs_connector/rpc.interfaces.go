package ipfs_connector

import "mime/multipart"

type RpcConnector interface {
	Add(fileHeader *multipart.FileHeader, pin bool) *AddResult
	Pin(id string)
}
