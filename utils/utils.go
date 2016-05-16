package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func ExtractBodyFromOSArgs(osArgs []string) map[string]string {
	var body string
	osHeaderMaxIndex := getOsHeaderMaxIndex(osArgs)
	for _, value := range osArgs[osHeaderMaxIndex+1:] {
		if value != "" {
			body += value + " "
		}
	}

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

	// We replace all instances of ":" with "==", and then we'll use it as our separator
	// for when we are now trying to compose a slice of string.
	// We are doing this because we can't rely on " " (spaces) as our separator since
	// on our request headers we could have a value (for a key) that has a space in between them
	// and if we split the string, the other string that was supposed to be a value will be treated
	// as a key for the headers.
	str = strings.Replace(str, ":", "==", -1)
	newSlice := strings.Split(str, "==")

	mapToReturn = SliceToMap(newSlice)
	return mapToReturn
}

func SliceToMap(slc []string) map[string]string {
	mapToReturn := make(map[string]string)
	for i := 0; i < len(slc); i += 2 {
		if i+1 < len(slc) {
			mapToReturn[slc[i]] = slc[i+1]
		}
	}
	return mapToReturn
}

func CreateNewRequest(method, address string, body io.Reader, headers map[string]string) *http.Request {
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

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return req
}

func DisplayResponseDetails(resp *http.Response) {
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

	log.Println("Response:", string(j))
	log.Println("Headers:", resp.Header)
	log.Println("Status:", resp.Status)
}

func CreateRequestBody(method string, osArgs []string) *strings.Reader {
	var body *strings.Reader
	bodyMap := make(map[string]string)

	if method != "GET" {
		bodyMap = ExtractBodyFromOSArgs(osArgs)
		j, err := json.Marshal(bodyMap)
		if err != nil {
			log.Fatal("Error marshaling object:", err)
			return nil
		}
		body = strings.NewReader(string(j))
	}
	return body
}

func CreateRequestHeaders(method string, osArgs []string) map[string]string {
	headersMap := make(map[string]string)
	osHeaderIndex := getOsHeaderMaxIndex(osArgs)
	var headers string

	for _, value := range osArgs[3:osHeaderIndex+1] {
		headers += value + " "
	}

	// we need to replace all extra unneeded characters set in our headers when we
	// inputted it in the terminal.
	headers = strings.Replace(headers, "[", "", -1)
	headers = strings.Replace(headers, "]", "", -1)
	headers = strings.Replace(headers, ",", "==", -1)

	headersMap = StringToMap(headers)
	return headersMap
}

func getOsHeaderMaxIndex(osArgs []string) int {
	var osHeaderIndex int
	for i, value := range os.Args {
		if strings.Contains(value, "]") {
			osHeaderIndex = i
		}
	}
	return osHeaderIndex
}