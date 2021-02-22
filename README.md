# go-censys

This project started of as a clone of https://github.com/oucema001/censys-go
but with added support for go modules and other stuff.

To start working with censys you have to create an account and get an application Key and application secret First. You can do this at https://censys.io.

### Installation

Download the package:

```bash
go get github.com/kmpm/go-censys
```

You can now use the library in your own projects :-)

### Usage

```go
package main

import (
    "log"
    
    "github.com/kmpm/go-censys"
)

func main() {
    client := censys.NewClient(nil, "MY_APPID", "MY_APPSECRET")
    dns, err := client.GetDNSResolve([]string{"google.com", "ya.ru"})
    
    a, err := client.Search(context.Background(), "www.google.com", censys.WEBSITES)
	if err != nil {
		log.Panic(err)
    } else {
        fmt.Println("alexa rank : %d",a.Results[0].AlexaRank)
    }
}
```
This will output : 

```bash
alexa rank : 1
```

###Implemented Censys APIs : 

#### Account
- [x] /api/v1/account
#### Search
- [x] /api/v1/search/ipv4
- [x] /api/v1/search/websites
- [x] /api/v1/search/certificates
#### View
- /api/v1/view/:index/:id
- [x] /api/v1/view/ipv4/:id
- [x] /api/v1/view/websites/:id
- [x] /api/v1/view/certificates/:id
#### Data
- [ ] /api/v1/data
#### Report
- /api/v1/report/:index
- [x] /api/v1/report/ipv4
- [ ] /api/v1/report/websites
- [ ] /api/v1/report/certificates

### Links
* [Censys API documentation](https://censys.io/api/v1/docs/report)
