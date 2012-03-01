# simple_blobstore

## Overview

**simple_blobstore**, an almost pointless blobstore with far too little useful functionality.  Built only to make quickly storing blobs that are HTTP accessible for another project.


## Requirements

goconfig - code at <https://github.com/kless/goconfig>


## Installation

```
go get github.com/kless/goconfig/config


```


## Notes

When this is done the following should be true:

### Writing

All writes are going to just write to a folder named after the user supplied GUID. 

There are no permissions or checks.

* **POST** a blob to /data/
* GUID returns as entirety of the body response.


### Reading

Still no permissions.

* **GET** to /data/?id=GUID will return the data written.  If there was no metadata written then the content-type will be application/octet-stream.


### TODO

Everything.