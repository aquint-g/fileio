package fileio

import (
	"bytes"
    "mime/multipart"
    "os"
    "fmt"
    "io"
    "log"
    "net/http"
    "io/ioutil"
    //"encoding/json"



)

func createMultipartFormData(fieldName, fileName string) (bytes.Buffer, *multipart.Writer) {
    var b bytes.Buffer
    var err error
    w := multipart.NewWriter(&b)
    var fw io.Writer
    file := mustOpen(fileName)
    if fw, err = w.CreateFormFile(fieldName, file.Name()); err != nil {
		fmt.Println("Error building form file")
    }
    if _, err = io.Copy(fw, file); err != nil {
		fmt.Println("Error with io.copy")    
	}
    w.Close()
    return b, w
}

func SendFile(filePath string) (string, int){

	b,w := createMultipartFormData("file",filePath)
	req, err := http.NewRequest("POST", "http://file.io", &b)
	if err != nil {
	    return "none",0
    }
	client := &http.Client{}
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, _ := client.Do(req)
	if resp.StatusCode != 200{
		return "",resp.StatusCode
	}else{
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
		    log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		return bodyString,resp.StatusCode;
	}
   	

}

func mustOpen(f string) *os.File {
    r, err := os.Open(f)
    if err != nil {
        pwd, _ := os.Getwd()
        fmt.Println("PWD: ", pwd)
        fmt.Println("The file specified can't be found.")
    }
    return r
}
