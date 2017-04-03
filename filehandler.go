package main

import "os"

var allfiles []filedata

type filedata struct {
   fpath string
   fname string
}

//callback for walk needs to match the following:
//type WalkFunc func(path string, info os.FileInfo, err error) error
func readFile (path string, fi os.FileInfo, _ error) error {   
   switch mode := fi.Mode(); {
   case mode.IsRegular():
      allfiles = ExtendFileDataSlice(allfiles, filedata{path, fi.Name()})
   }
   return nil
}

//ExtendEntitySlice extends a slice of type EntityData
func ExtendFileDataSlice(slice []filedata, fd filedata) []filedata {
   n := len(slice)
   if n == cap(slice) {
      // Slice is full; must grow.kb
      // We double its size and add 1, so if the size is zero we still grow.
      newSlice := make([]filedata, len(slice), 2*len(slice)+1)
      copy(newSlice, slice)
      slice = newSlice
   }
   slice = slice[0 : n+1]
   slice[n] = fd
   return slice
}


