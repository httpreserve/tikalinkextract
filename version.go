package main

import "github.com/httpreserve/linkscanner"

var version = "tikalinkextract-0.0.3"

func getVersion() string {
	return version + "\n" + linkscanner.GetVersion()
}
