package main

import (
   "os"
   "fmt"
)

func findOpenConnections() {
   var tika string = "http://127.0.0.1:9998/"
   var resp int8

   resp = testConnection(tika)
   if resp == CONN_BAD {
      fmt.Fprintln(os.Stdout, "INFO: Tika connection not available to connect to. Check localhost:9998.")
      os.Exit(1)
   }
}

