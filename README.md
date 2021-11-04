# WIFI Password

> Get current wifi password


## Install

```shell
$ go get github.com/g-lib/wifi-password
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/g-lib/wifi-password"
)


func main(){
    ssid,_ :=wifipw.WIFISSID()
    password,_ := wifipw.WIFIPassword(ssid)
    fmt.Println(ssid,password)
}
```

OR

```go
package main

import (
    "fmt"
    "github.com/g-lib/wifi-password"
)


func main(){
    ssid,password,_ := wifipw.GetWIFIPassword()
    fmt.Println(ssid,password)
}
```