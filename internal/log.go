package log

import (
	"fmt"
	"time"
)

type log struct {
	Count    int
	logsList []string
}

func (_l *log) Add(record string) {
	t := time.Now().Format("2006-01-02T15:04:05")
	_l.logsList = append(_l.logsList, t+" "+record+"\n")
}

func (_l log) Print() {
	for _, v := range _l.logsList {
		fmt.Printf("%s", v)
	}

}
