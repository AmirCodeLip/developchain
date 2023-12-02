package main

import (
	"bytes"
	"dc_middleware/ipfs_connector"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func createForm(form map[string]string) (string, io.Reader, error) {
	body := new(bytes.Buffer)
	mp := multipart.NewWriter(body)
	defer mp.Close()
	for key, val := range form {
		if strings.HasPrefix(val, "@") {
			val = val[1:]
			file, err := os.Open(val)
			if err != nil {
				return "", nil, err
			}
			defer file.Close()
			part, err := mp.CreateFormFile(key, val)
			if err != nil {
				return "", nil, err
			}
			io.Copy(part, file)
		} else {
			mp.WriteField(key, val)
		}
	}
	return mp.FormDataContentType(), body, nil
}

func addFile(action string) {

}

func main() {
	// "Connect" to local node
	connector := ipfs_connector.NewRpc("http://168.119.58.0:5001")
	response := connector.Add("C:\\Users\\amir\\Pictures\\FFFF.JPG", false)

	//form := map[string]string{"path": "@C:\\Users\\amir\\Pictures\\mili.jpg"}
	//contentType, body, err := createForm(form)
	fmt.Println(response.Hash)
	fmt.Println(response.Name)
}
