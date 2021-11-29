package main

import {
	"APISERVER/log"
}

func main() {
	//str := time.Now().String()
	//t, _ := time.Parse(time.RFC822Z, str)
	l := new(log)
	l.Add("record 1")
	l.Add("record 2")
	l.Add("record 3")
	l.Add("record 4")
	l.Add("record 5")
	l.Add("record 6")
	l.Add("record 7")
	l.Add("record 8")
	l.Add("record 9")
	l.Print()

}

