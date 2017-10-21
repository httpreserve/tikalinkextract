package main

import "github.com/httpreserve/linkscanner"

var version = "tikalinkextract-0.0.2"

func getVersion() string {
	return version + "\n" + linkscanner.GetVersion()
}
