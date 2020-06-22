package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const timelayoutISO = "15:04:05"
const datelayoutISO = "2006-01-02"

type Todos struct {
	Todos []Todo `json:"dates"`
}

//Time is being exported
type Todo struct {
	Date   string            `json:"date"`
	Time   map[string]string `json:"time"`
	Repeat bool              `json:"repeat"`
}

func parse(filename string) Todos {

	var todo Todos

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bytes, &todo)
	if err != nil {
		log.Fatal(err)
	}
	return todo
}

func (t *Todos) printTodolist() {
	//fmt.Println(t)
	fmt.Println()
	for i := 0; i < len(t.Todos); i++ {
		fmt.Printf("\033[31mdate: %v\033[0m\n", t.Todos[i].Date)
		for timing, task := range t.Todos[i].Time {
			t, err := time.Parse(timelayoutISO, timing)
			fmt.Println("time:", t.Format(timelayoutISO), "-> task:", task)

			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Printf("repeat: %t \n\n", t.Todos[i].Repeat)

	}

}
