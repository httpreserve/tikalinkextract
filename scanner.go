package main

import (
	"fmt"
	"github.com/httpreserve/linkscanner"
	"os"
	"strings"
)

var linklist []string
var seeds map[string]bool

func init() {
	seeds = make(map[string]bool)
}

func httpScanner(fname string, content string) {
	extracted, errs := linkscanner.HTTPScanner(content)
	if len(extracted) > 0 {
		for _, link := range extracted {
			var addedValue string
			if seedList {
				if !seeds[strings.ToLower(link)] {
					seeds[strings.ToLower(link)] = true
					addedValue = link
				}
			} else {
				if quoteCells {
					addedValue = "\"" + fname + "\", \"" + link + "\""
				} else {
					addedValue = fname + ", " + link
				}
			}
			if addedValue != "" {
				linklist = append(linklist, addedValue)
			}
		}
	}
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Fprintf(os.Stderr, "%s", e.Error())
		}
	}
}
