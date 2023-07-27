package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"course_name"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "xyz123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "abc123", []string{"fullstack", "js"}},
		{"DevOps Bootcamp", 299, "LearnCodeOnline.in", "qwe123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, " ", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(` {
		"Name": "ReactJS Bootcamp",
		"Price": 299,
		"Platform": "LearnCodeOnline.in",
		"Password": "xyz123",
		"Tags": ["web-dev","js"]
	}
	`)

	var lcoCourses course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON data is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v", lcoCourses)
	} else {
		fmt.Println("JSON data is not valid")
	}
	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is :%T\n", k, v, v)
	}
}
