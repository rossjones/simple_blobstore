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
    //    "log"
    "net/http"
)

var metadataFuncs = map[string]http.HandlerFunc{
    "GET":  getMetaData,
    "POST": postMetaData,
}
var dataFuncs = map[string]http.HandlerFunc{
    "GET":  getData,
    "POST": postData,
}

func unsupportedMethod(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, req.Method+" is unsupported\n")
}

func getMetaData(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "GET Metadata\n")
}

func postMetaData(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "POST Metadata\n")
}

func getData(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "GET Data\n")
}

func postData(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "POST Data\n")
}

func MetadataServer(w http.ResponseWriter, req *http.Request) {
    _, present := metadataFuncs[req.Method]
    if present {
        metadataFuncs[req.Method](w, req)
    } else {
        unsupportedMethod(w, req)
    }
}

func DataServer(w http.ResponseWriter, req *http.Request) {
    _, present := dataFuncs[req.Method]
    if present {
        dataFuncs[req.Method](w, req)
    } else {
        unsupportedMethod(w, req)
    }
}
