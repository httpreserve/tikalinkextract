package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

const methodPUT = http.MethodPut
const methodGET = http.MethodGet
const methodPOST = http.MethodPost

const connOKAY int8 = 0
const connBAD int8 = 1

const acceptMIME = "text/csv"
const acceptJSON = "application/json"
const acceptXMP = "application/rdf+xml"

func testConnection(request string) int8 {

	conn := connOKAY
	stream, err := http.NewRequest(methodGET, request, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: error creating request,", err)
		os.Exit(1)
	}

	client := &http.Client{}
	_, err = client.Do(stream)
	if err != nil {
		conn = connBAD
	}

	return conn
}

func handleConnection(stream *http.Request) string {

	client := &http.Client{}
	response, err := client.Do(stream)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: Doing request,", err)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: Reading response body,", err)
		os.Exit(1)
	}

	dataString := string(data)
	trimmedResponse := strings.TrimSpace(dataString) //byte [] becomes string
	return trimmedResponse
}

func makeMultipartConnection(VERB string, request string, fp *os.File, fname string, accepttype string) string {

	//https://gist.github.com/mattetti/5914158/f4d1393d83ebedc682a3c8e7bdc6b49670083b84
	fileContents, err := ioutil.ReadAll(fp)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", fname)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	part.Write(fileContents)

	err = writer.Close() //writer must be closed to avoid: [UNEXPECTED EOF]
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	stream, err := http.NewRequest(VERB, request, body)
	stream.Header.Add("Content-Type", writer.FormDataContentType())

	if accepttype != "" {
		stream.Header.Add("Accept", accepttype)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: error creating request,", err)
		os.Exit(1)
	}

	return handleConnection(stream)
}

func makeConnection(VERB string, request string, fp *os.File, accepttype string) string {

	var stream *http.Request
	var err error

	if fp != nil {
		stream, err = http.NewRequest(VERB, request, fp)
	} else {
		stream, err = http.NewRequest(VERB, request, nil)
	}

	if accepttype != "" {
		stream.Header.Add("Accept", accepttype)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: error creating request,", err)
		os.Exit(1)
	}

	return handleConnection(stream)
}
