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
    "flag"
    "fmt"
    "github.com/kless/goconfig/config"
    "log"
    "net/http"
    "os"
)

var StorageLocation = ""
var configFile = flag.String("c", "config.ini", "Location of the config file")

func main() {
    flag.Parse()

    c, err := config.ReadDefault(*configFile)
    if err != nil {
        log.Fatal("Couldn't find the config file: ", err)
    }

    host, _ := c.String("network", "listen-ip")
    port, _ := c.Int("network", "listen-port")
    listen := fmt.Sprintf("%s:%d", host, port)

    StorageLocation, err = c.String("storage", "location")
    if err != nil {
        StorageLocation = "./"
    }
    log.Println("Storing/Retrieving data from ", StorageLocation)

    _, err = os.Stat(StorageLocation)
    if err != nil {
        log.Fatal("Folder not exist: ", err)
    }

    http.HandleFunc("/", DataServer)

    log.Println("Server listening: ", listen)
    err = http.ListenAndServe(listen, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
