package ipfs_connector

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func NewRpc(url string) RpcConnector {
	rpcConnectorData := RpcConnectorData{url + "/api/v0/"}
	var rpcConnector RpcConnector
	rpcConnector = &rpcConnectorData
	return rpcConnector
}

func (rpcConnectorData *RpcConnectorData) Add(filePath string, pin bool) *AddResult {
	body := new(bytes.Buffer)
	mp := multipart.NewWriter(body)
	defer mp.Close()
	file, fileError := os.Open(filePath)
	if fileError != nil {
		panic(fileError)
	}
	defer file.Close()
	fileName := filepath.Base(filePath)
	formFile, formFileError := mp.CreateFormFile("path", fileName)
	mp.WriteField("pin", strconv.FormatBool(pin))
	if formFileError != nil {
		panic(formFileError)
	}
	io.Copy(formFile, file)
	response, responseError := http.Post(rpcConnectorData.url+"add", mp.FormDataContentType(), body)
	if responseError != nil {
		panic(responseError)
	}
	bytesBody, bytesBodyErr := io.ReadAll(response.Body)
	if bytesBodyErr != nil {
		panic(bytesBodyErr)
	}
	dataList := strings.Split(string(bytesBody), "}")
	var stringBody []byte
	if len(dataList) > 0 {
		stringBody = []byte(dataList[0] + "}")
	} else {
		return nil
	}
	var addResult AddResult
	json.Unmarshal(stringBody, &addResult)
	return &addResult
}
