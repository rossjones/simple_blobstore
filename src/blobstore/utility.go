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
    "fmt"
    "os"
)

func makeGuid() string {
    f, _ := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
    b := make([]byte, 16)
    f.Read(b)
    f.Close()
    return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
