package huaweicloud
import (
	"fmt"
	"github.com/libdns/libdns"
    "time"
)

func parseTTL(record *libdns.Record)  {
	//ttl := int32(record.TTL.Seconds())
	record.TTL = time.Duration(300)* time.Second
	 
}

func SolveRecordValue(rType string,value string) []string {
	switch rType {
	case "TXT":
		value = fmt.Sprintf("\"%s\"", value)

	}
	
	return []string{value}
}