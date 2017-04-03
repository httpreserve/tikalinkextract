package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func getTikaRecursive(fname string, fp *os.File, accepttype string) ([]string, map[string]interface{}, error) {

	var flRecursiveMdKeys []string
	var flRecursiveKeysValues map[string]interface{}

	fp.Seek(0, 0)
	resp := makeMultipartConnection(methodPOST, tikaPathMetaRecursive, fp, fname, accepttype)
	trimmed := strings.Trim(resp, "[ ]")
	err := readTikaMetadataJSON(trimmed, "", &flRecursiveKeysValues, &flRecursiveMdKeys)
	return flRecursiveMdKeys, flRecursiveKeysValues, err
}

func readTikaMetadataJSON(output string, key string, kv *map[string]interface{}, mdkeys *[]string) error {
	if output != "" {
		//we can get multiple JSON sets from TIKA
		jsonStrings := strings.Split(output, "},")
		for k, v := range jsonStrings {
			last := v[len(v)-1:]
			if last != "}" {
				jsonStrings[k] = v + "}"
			}
		}
		var tikamap map[string]interface{}
		for _, v := range jsonStrings {
			if err := json.Unmarshal([]byte(v), &tikamap); err != nil {
				fmt.Fprintln(os.Stderr, "ERROR: Handling TIKA JSON,", err)
			}
			*kv = tikamap
			getTikaKeys(tikamap, mdkeys)
		}
		return nil
	}
	return fmt.Errorf("Response data is a nil string.")
}

func getTikaKeys(tikamap map[string]interface{}, mdkeys *[]string) {
	keys := make([]string, len(tikamap))
	i := 0
	for k := range tikamap {
		keys[i] = k
		i++
	}
	*mdkeys = keys //alt: /meta/{field} TIKA URL
}
