package ipfs_connector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
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

func (rpcConnectorData *RpcConnectorData) Pin(id string) {
	req, error := http.NewRequest("post", rpcConnectorData.url+"pin/add?arg="+id, nil)
	if error != nil {
		panic(error)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//_, bytesBodyErr := io.ReadAll(response.Body)
	//if bytesBodyErr != nil {
	//	panic(bytesBodyErr)
	//}
	if resp != nil {
		fmt.Println("file is pinned with hash %s", id)
	}
	//print(string(bytesBody))
}

func (rpcConnectorData *RpcConnectorData) Add(fileHeader *multipart.FileHeader, pin bool) *AddResult {
	body := new(bytes.Buffer)
	mp := multipart.NewWriter(body)
	defer mp.Close()
	file, fileError := fileHeader.Open()
	if fileError != nil {
		panic(fileError)
	}
	defer file.Close()
	fileName := filepath.Base(fileHeader.Filename)
	//fmt.Println(strconv.FormatBool(pin))
	formFile, formFileError := mp.CreateFormFile("path", fileName)
	if formFileError != nil {
		panic(formFileError)
	}
	io.Copy(formFile, file)
	response, responseError := http.Post(rpcConnectorData.url+"add?pin="+strconv.FormatBool(pin), mp.FormDataContentType(), body)
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
