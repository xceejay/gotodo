// Package main provides ...
package main

import (
	//"fmt"
	//"log"
	"time"
)

func start(filename string) {
	todolists := parse(filename)
	//todolists.printTodolist()
	for {

		for i := 0; i < len(todolists.Todos); i++ {
			//fmt.Printf("\033[31mdate: %v\033[0m\n", todolists.Todos[i].Date)
			//if time.Now().Day() == todolists.Todos[i].Date {

			//}
			for timing, task := range todolists.Todos[i].Time {
				//t, err := time.Parse(timelayoutISO, timing)
				//fmt.Println("time:", t.Format(timelayoutISO), "-> task:", task)

				if time.Now().Format(datelayoutISO) == todolists.Todos[i].Date && timing == time.Now().Format(timelayoutISO) {

					notify(task)
				}

				//if err != nil {
				//log.Fatal(err)
				//}
			}

			//fmt.Printf("repeat: %t \n\n", todolists.Todos[i].Repeat)

		}

		time.Sleep(time.Millisecond * 500)
	}
}
