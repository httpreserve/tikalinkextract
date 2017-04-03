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

const PUT = http.MethodPut
const GET = http.MethodGet
const POST = http.MethodPost

const CONN_OKAY int8 = 0
const CONN_BAD int8 = 1

const ACCEPT_MIME_CSV = "text/csv"
const ACCEPT_MIME_JSON = "application/json"
const ACCEPT_MIME_XMP = "application/rdf+xml"

func testConnection(request string) int8 {

	conn := CONN_OKAY
	stream, err := http.NewRequest(GET, request, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: error creating request,", err)
		os.Exit(1)
	}

	client := &http.Client{}
	_, err = client.Do(stream)
	if err != nil {
		conn = CONN_BAD
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

	data_string := string(data)
	trimmed_response := strings.TrimSpace(data_string) //byte [] becomes string
	return trimmed_response
}

func makeMultipartConnection(VERB string, request string, fp *os.File, fname string, accepttype string) string {

	//https://gist.github.com/mattetti/5914158/f4d1393d83ebedc682a3c8e7bdc6b49670083b84
	fileContents, err := ioutil.ReadAll(fp)
	if err != nil {
		fmt.Println(err)
	}

	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", fname)
	if err != nil {
		fmt.Println(err)
	}
	part.Write(fileContents)

	err = writer.Close() //writer must be closed to avoid: [UNEXPECTED EOF]
	if err != nil {
		fmt.Println(err)
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
