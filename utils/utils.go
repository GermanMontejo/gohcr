package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func ExtractBodyFromOSArgs(osArgs []string) map[string]string {
	var body string
	for _, value := range osArgs[4:] {
		body += value + " "
	}
	fmt.Println(body)
	BodyMap := StringToMap(body)
	switch osArgs[1] {
	case "GET":
		return nil
	default:
		return BodyMap
	}
}

func StringToMap(str string) map[string]string {
	mapToReturn := make(map[string]string)

	if str == "" {
		log.Print("Empty string - aborting conversion of string to map...")
		return nil
	}

	newStr := strings.Replace(str, "[", "", -1)
	newStr = strings.Replace(newStr, "]", "", -1)
	newStr = strings.Replace(newStr, `"`, "", -1)
	newStr = strings.Replace(newStr, ":", " ", -1)
	newSlice := strings.Split(newStr, " ")

	for i := 0; i < len(newSlice); i += 2 {
		if i+1 < len(newSlice) {
			mapToReturn[newSlice[i]] = newSlice[i+1]
		}
	}
	return mapToReturn
}

func CreateNewRequest(method, address string, body io.Reader) *http.Request {
	var req *http.Request
	var err error
	if strings.EqualFold(method, "GET") {
		req, err = http.NewRequest(method, address, nil)
	} else {
		req, err = http.NewRequest(method, address, body)
	}

	if err != nil {
		log.Fatal("Error creating request object:", err)
		return nil
	}
	return req
}

func DisplayResponseDetails(resp *http.Response) {
	// var prettyJson bytes.Buffer
	var responseBody interface{}

	json.NewDecoder(resp.Body).Decode(&responseBody)
	j, err := json.MarshalIndent(responseBody, "", "\t")

	if err != nil {
		log.Println("Error marshaling response body:", err)
		return
	}

	if err != nil {
		log.Println("Error formatting response body:", err)
		return
	}

	fmt.Println("Response:", string(j))
	fmt.Println("Headers:", resp.Header)
	fmt.Println("Status:", resp.Status)
}

func CreateRequestBody(method string, osArgs []string) *strings.Reader {
	var body *strings.Reader
	bodyMap := make(map[string]string)

	if method != "GET" {
		bodyMap = ExtractBodyFromOSArgs(os.Args)
		j, err := json.Marshal(bodyMap)
		if err != nil {
			log.Fatal("Error marshaling object:", err)
			return nil
		}
		body = strings.NewReader(string(j))
	}
	return body
}
