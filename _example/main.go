package main

import (
	"context"
	"fmt"
    "time"
	"github.com/jicjoy/huaweicloud"
	"github.com/libdns/libdns"
)

func main() {
	p := huaweicloud.Provider{
		AccessKeyId:     "YOUR_Secret_ID",
		SecretAccessKey: "YOUR_Secret_Key",
	}

	ret, err := p.SetRecords(context.TODO(), "iitmall.com",[]libdns.Record{
		 {
			Type: "TXT",
			Name: "es",
			Value: "ssssfsdfsdf",
			TTL: time.Duration(10)* time.Second,
		 },
	})

	fmt.Println("Result:", ret)
	fmt.Println("Error:", err)
}
