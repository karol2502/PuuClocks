package actions

import (
	"fmt"
	"strconv"
	"strings"
)

type ReportTime struct {
	action
}

func (r ReportTime) Validate(data string) *action {
	d := strings.Split(data, "/")
	if len(d) != 2 {
		fmt.Println("Couldn't split")
		return nil
	}

	if d[0] != string(ActionTypeReportTime) {
		fmt.Println("Incorrect type")
		return nil
	}

	parsedTime, err := strconv.ParseFloat(d[1], 64)
	if err != nil {
		fmt.Println("Wrong format: ", d[1])
		return nil
	}

	if parsedTime > 12 || parsedTime < 0 {
		fmt.Println("Incorrect time")
		return nil
	}

	return &action{
		Type: ActionTypeReportTime,
		Data: ActionData{
			ReportedTime: &parsedTime,
		},
	}
}
