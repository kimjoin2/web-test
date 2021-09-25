package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"web-test/configFormat"
)

func main() {
	configFilePath := flag.String("f", "", "configFormat file path")
	flag.Parse()

	// config file path check
	if *configFilePath == "" {
		fmt.Println("need to set configFormat file with -f option")
		os.Exit(1)
	}

	// config file open check
	f, e := os.Open(*configFilePath)
	if e != nil {
		fmt.Printf("cannot open file : %s\n", *configFilePath)
		os.Exit(1)
	}
	defer f.Close()
	byteValue, e := ioutil.ReadAll(f)
	if e != nil {
		fmt.Printf("cannot open file : %s\n", *configFilePath)
		os.Exit(1)
	}

	// json format check
	var config configFormat.TestConfig
	e = json.Unmarshal(byteValue, &config)
	if e != nil {
		fmt.Printf("configFormat configFormat error : %s\n%s", *configFilePath, e.Error())
		os.Exit(1)
	}

	// run test
	totalCount := 0
	failCount := 0
	successCount := 0
	errorCount := 0

	// for avoiding unused error
	_ = totalCount
	_ = failCount
	_ = successCount
	_ = errorCount

	baseUrl := config.BaseUrl
	for _, testCase := range config.TestCases{

		requestData := testCase.RequestData

		targetUrl := baseUrl + requestData.Path
		method := strings.ToUpper(requestData.Method)
		headers := requestData.Headers
		body := requestData.Body
		contentType := requestData.ContentType
		var res *http.Response
		var e error

		// for avoiding unused error
		_ = headers
		_ = body
		_ = contentType
		_ = res
		_ = e

		switch method {
		case http.MethodGet:
			res, e = http.Get(targetUrl)
		case http.MethodPost:
			res, e = http.Post(targetUrl, contentType, strings.NewReader(body))
		case http.MethodHead:
			res, e = http.Head(targetUrl)
		default:
			e = errors.New(fmt.Sprintf("%s is not supported.\n", requestData.Method))
		}
		totalCount++

		expectedResponse := testCase.ExpectedResponse
		if e != nil {
			errorCount++
			fmt.Println("error occurred : ", testCase, e.Error())
		} else if res.StatusCode != expectedResponse.Code {
			// TODO : check response data all(not only response code)
			failCount++
			fmt.Println("test fail : ", testCase)
		} else {
			successCount++
		}
	}
	if totalCount > 0{
		fmt.Println("==============================")
		fmt.Println("summary :")
		fmt.Printf("total: %d, success: %d(%0.1f%%), fail: %d(%0.1f%%), error: %d(%0.1f%%)",
			totalCount,
			successCount, 100 * float32(successCount) / float32(totalCount),
			failCount, 100 * float32(failCount) / float32(totalCount),
			errorCount, 100 * float32(errorCount) / float32(totalCount))
	} else if totalCount == 0 {
		fmt.Println("no test case")
		os.Exit(1)
	} else {
		fmt.Println("internal error")
		os.Exit(1)
	}
}
