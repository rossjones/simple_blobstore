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

All writes are going to just write to a folder named after the user supplied GUID. POSTs to existing GUIDs will overwrite the data.

There are no permissions or checks.

* **POST** some JSON to /metadata/GUID

* **POST** a blob to /data/GUID
* If GUID matches metadata then it is associated, either way the blob will be stored.


### Reading

Still no permissions.

* **GET** to /data/GUID will return the data written.  If there was no metadata written then the content-type will be application/octet-stream.
* **GET** to /metadata/GUID will return the appropriate metadata, if it exists.


### TODO

Everything.