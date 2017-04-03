package main

import (
	"bufio"
	"strings"
)

// strings to look for that indicate a web resource
var (
	protoHTTPS = "https://"
	protoHTTP  = "http://"
	protoWww   = "www." // technically not a protocol
	protoFtp   = "ftp://"
)

//common line endings that shouldn't be in URL
var common = []string{"�", "\"", "'", ":", ";", ".", "`", ",", "*"}

func cleanLink(link string, www bool) string {
	if www && !noprotocol {
		link = protoHTTP + link
	}

	//utf-8 replacement code character
	//https://codingrigour.wordpress.com/2011/02/17/the-case-of-the-mysterious-characters/
	link = strings.Replace(link, "\xEF\xBF\xBD", "", 1)

	// replace common invalid line-endings
	for _, x := range common {
		if x == link[len(link)-1:] {
			substring := link[0 : len(link)-1]
			return cleanLink(substring, false)
		}
	}
	return link
}

func retrieveLink(literal string) string {
	if strings.Contains(literal, protoHTTPS) {
		literal = literal[strings.Index(literal, protoHTTPS):]
		return cleanLink(literal, false)
	}
	if strings.Contains(literal, protoHTTP) {
		literal = literal[strings.Index(literal, protoHTTP):]
		return cleanLink(literal, false)
	}
	if strings.Contains(literal, protoFtp) {
		literal = literal[strings.Index(literal, protoFtp):]
		return cleanLink(literal, false)
	}
	if strings.Contains(literal, protoWww) {
		literal = literal[strings.Index(literal, protoWww):]
		return cleanLink(literal, true)
	}
	return ""
}

var linklist []string

func httpScanner(fname string, content string) {

	reader := bufio.NewReader(strings.NewReader(content))
	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		link := retrieveLink(scanner.Text())
		if link != "" {
			addedValue := fname + ", " + link
			seen := false
			for _, x := range linklist {
				if x == addedValue {
					seen = true
					break
				}
			}
			if !seen {
				linklist = append(linklist, addedValue)
			}
		}
	}
}
