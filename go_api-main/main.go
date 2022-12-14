package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Votes *Votes `json:"votes"`
	Count string `json:"count,omitempty"`
}

type Votes struct {
	OptionA string `json:"option_A"`
}

func mains() {
	s := `{ "votes": { "option_A": "3" } }`
	data := &Data{
		Votes: &Votes{},
	}
	err := json.Unmarshal([]byte(s), data)
	fmt.Println(err)
	fmt.Println(data.Votes)
	s2, _ := json.Marshal(data)
	fmt.Println(string(s2))
	data.Count = "2"
	s3, _ := json.Marshal(data.Count)
	fmt.Println("s3: ", string(s3))
}
