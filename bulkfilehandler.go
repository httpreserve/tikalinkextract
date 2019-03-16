package main

import (
	"fmt"
	"os"
)

type contenterror struct {
	content string
	err     error
	fname   string
}

func extractAndAnalyse(filepool []filedata) (bool, error) {
	//make channel run goroutine...
	ch := make(chan contenterror)
	for _, fi := range filepool {
		go getFileContent(fi, ch)
	}
	for range filepool {
		ce := <-ch
		if ce.err != nil {
			logFileMessage("INFO: '%s' cannot be handled by Tika.", ce.fname)
		} else {
			httpScanner(ce.fname, ce.content)
		}
	}
	return false, nil
}

func openFile(path string) (*os.File, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return fp, nil
}

func getFileContent(fi filedata, ch chan contenterror) {
	//create empty struct to return...
	var ce contenterror
	ce.fname = fi.fname
	//what are we doing..?
	logFileMessage("INFO: '%s' being processed.", fi.fname)
	//process...
	fp, err := openFile(fi.fpath)
	defer fp.Close()
	if err != nil {
		ce.err = err
		ch <- ce
		return
	}
	_, flRecursiveKeysValues, err := getTikaRecursive(fi.fname, fp, acceptJSON)
	if err != nil {
		ce.err = err
		ch <- ce
		return
	}
	if val, ok := flRecursiveKeysValues[tikaPlainText]; ok {
		ce.content = val.(string)
		ch <- ce
		return
	}
	ce.err = fmt.Errorf("No plain text data to analyse.")
	ch <- ce
	return
}
