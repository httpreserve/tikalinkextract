package main

import (
  	"os"
	"fmt"
   "strings"
	"encoding/json"
)

func getTikaRecursive (fname string, fp *os.File, accepttype string) ([]string, map[string]interface{}, error) {

   var fl_recursive_md_keys []string
   var fl_recursive_keys_values map[string]interface{}

   fp.Seek(0,0)
   resp := makeMultipartConnection(POST, tika_path_meta_recursive, fp, fname, accepttype) 
   trimmed := strings.Trim(resp, "[ ]")
   err := readTikaMetadataJson(trimmed, "", &fl_recursive_keys_values, &fl_recursive_md_keys)
   return fl_recursive_md_keys, fl_recursive_keys_values, err
}

func readTikaMetadataJson (output string, key string, kv *map[string]interface{}, mdkeys *[]string) error {
   if output != "" {
      //we can get multiple JSON sets from TIKA
      json_strings := strings.Split(output, "},")
      for k, v := range json_strings {
         last := v[len(v)-1:]
         if last != "}" {
            json_strings[k] = v + "}"
         }
      }
	   var tikamap map[string]interface{}
      for _, v := range json_strings {
	      if err := json.Unmarshal([]byte(v), &tikamap); err != nil {
		      fmt.Fprintln(os.Stderr, "ERROR: Handling TIKA JSON,", err)
	      }
	      *kv = tikamap
	      getTikaKeys(tikamap, mdkeys) 
      }
      return nil
   } else {
      return fmt.Errorf("Response data is a nil string.")
   }
} 

func getTikaKeys (tikamap map[string]interface{}, mdkeys *[]string) {	
	keys := make([]string, len(tikamap))
	i := 0
	for k := range tikamap {
		keys[i] = k
		i++
	}
	*mdkeys = keys    //alt: /meta/{field} TIKA URL
}

