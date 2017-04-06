package main

import (
	"fmt"
	"github.com/httpreserve/linkscanner"
	"os"
)

var linklist []string

func httpScanner(fname string, content string) {
	extracted, errs := linkscanner.HTTPScanner(content)
	if len(extracted) > 0 {
		for _, link := range extracted {
			addedValue := fname + ", " + link
			linklist = append(linklist, addedValue)
		}
	}

	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Fprintf(os.Stderr, "%s", e.Error())
		}
	}
}