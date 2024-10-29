package huaweicloud
import (
	"fmt"
)

func SolveRecordValue(rType string,value string) []string {
	switch rType {
	case "TXT":
		value = fmt.Sprintf("\"%s\"", value)

	}
	
	return []string{value}
}