package main

import (
	"context"
	"fmt"

	"github.com/libdns/huaweicloud"
)

func main() {
	p := huaweicloud.Provider{
		AccessKeyId:     "YOUR_Secret_ID",
		SecretAccessKey: "YOUR_Secret_Key",
	}

	ret, err := p.GetRecords(context.TODO(), "your-domain")

	fmt.Println("Result:", ret)
	fmt.Println("Error:", err)
}
