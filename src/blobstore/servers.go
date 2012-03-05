/******************************************************************************
* Simple_blobstore - stores blobs of data as simply as possible
* Copyright (C) 2012 Ross Jones
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU Affero General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU Affero General Public License for more details.
*
* You should have received a copy of the GNU Affero General Public License
* along with this program.  If not, see <http://www.gnu.org/licenses/>.
******************************************************************************/
package main

import (
    "io"
    "os"
    "log"
    "net/http"
    "path"
)

var dataFuncs = map[string]http.HandlerFunc{
    "GET":  getData,
    "POST": postData,
}

func unsupportedMethod(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, req.Method+" is unsupported\n")
}

func getData(w http.ResponseWriter, req *http.Request) {
    guid := req.FormValue("id")
    if guid == "" {
        http.NotFound(w,req)
        return;
    }

    target := path.Join(StorageLocation, guid, "blob")
    _, err := os.Stat(target)
    if err != nil {
        log.Println("Data not exist: ", err)
        http.NotFound(w,req)
        return
    }

    output, err := os.Open( target )
    if err != nil {
        log.Println("Can't open target file ", target, " ", err)
        return
    }
    defer output.Close()


    _, err = io.Copy(w, output)
    if err != nil {
        return
    }
}

// Reads a form variable called 'file' which it is hoped is
// an actual file.
//
// Currently loads file into memory, but obviously this won't work for
// anything other than tiny files.
func postData(w http.ResponseWriter, req *http.Request) {

    file, handler, err := req.FormFile("file")
    if err != nil {
        log.Fatal(err)
        return
    }
    defer file.Close()

    name := handler.Filename
    log.Println(name)

    guid := makeGuid()
    target_folder := path.Join(StorageLocation, guid)
    err = os.Mkdir(target_folder, 0777)

    target_file := path.Join(target_folder, "blob")
    output, err := os.Create( target_file )
    if err != nil {
        log.Fatal("Can't create target file ", target_file, " ", err)
        return
    }
    defer output.Close()

    _, err = io.Copy(output, file)
    if err != nil {
        log.Fatal("Can't copy to target file", target_file)
        return
    }

    io.WriteString( w, guid )
}


func DataServer(w http.ResponseWriter, req *http.Request) {
    _, present := dataFuncs[req.Method]
    if present {
        dataFuncs[req.Method](w, req)
    } else {
        unsupportedMethod(w, req)
    }
}
