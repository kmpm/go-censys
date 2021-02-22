package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/kmpm/go-censys"
)

var (
	appid     string
	appsecret string
)

func main() {
	appid = os.Getenv("CENSYS_ID")
	appsecret = os.Getenv("CENSYS_SECRET")
	flag.StringVar(&appid, "id", appid, "censys appid")
	flag.StringVar(&appsecret, "secret", appsecret, "censys appsecret")
	flag.Parse()

	client := censys.NewClient(nil, appid, appsecret)

	query := censys.ReportQuery{
		Query:   "80.http.get.headers.server: nginx",
		Field:   "location.country",
		Buckets: 10,
	}
	a, err := client.GetReport(context.Background(), censys.IPV4REPORT, query)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v+", a)
}
