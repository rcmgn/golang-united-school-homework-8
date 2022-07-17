package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Arguments map[string]string

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

//go run .\main.go -operation="add" -item={"id": "1", "email": "email@test.com", "age": 23} -fileName="users.json"
// HW80.exe -operation add  -fileName test.json -item {\"id\":\"1\",\"email\":\"test@test.com\",\"age\":34}
// HW80.exe -operation add  -fileName test.json -item {\"id\":\"1\",\"email\":\"test@test.com\",\"age\":34}

func parseArgs() (args Arguments) {
	args = make(Arguments)
	operFlag := flag.String("operation", "", "-operation flag has to be specified")
	fileFlag := flag.String("fileName", "", "-fileName flag has to be specified")
	flagItem := flag.String("item", "", "item usage")
	flagId := flag.String("id", "", "id usage")
	flag.Parse()
	args["id"] = *flagId
	args["item"] = *flagItem
	args["operation"] = *operFlag
	args["fileName"] = *fileFlag
	return args
}

func findByID(data []User, id string) (res int) {
	res = -1
	for i, val := range data {
		if val.Id == id {
			return i
		}
	}
	return
}

func Perform(args Arguments, writer io.Writer) error {
	var data []User
	var item User

	if args["operation"] == "" {
		return errors.New("-operation flag has to be specified")
	}
	if args["fileName"] == "" {
		return errors.New("-fileName flag has to be specified")
	}
	b, _ := ioutil.ReadFile(args["fileName"])
	json.Unmarshal(b, &data)
	switch args["operation"] {
	case "list":
		writer.Write(b)
	case "add":
		if args["item"] == "" {
			return errors.New("-item flag has to be specified")
		}
		json.Unmarshal([]byte(args["item"]), &item)
		if findByID(data, item.Id) >= 0 {
			writer.Write([]byte("Item with id " + item.Id + " already exists"))
		} else {
			data = append(data, item)
		}
	case "findById":
		if args["id"] == "" {
			return errors.New("-id flag has to be specified")
		}
		i := findByID(data, args["id"])
		if i != -1 {
			b, _ := json.Marshal(data[i])
			writer.Write(b)
		}
	case "remove":
		if args["id"] == "" {
			return errors.New("-id flag has to be specified")
		}
		i := findByID(data, args["id"])
		if i == -1 {
			writer.Write([]byte("Item with id " + args["id"] + " not found"))
		} else {
			data = append(data[:i], data[i+1:]...)
		}
	default:
		return errors.New("Operation abcd not allowed!")
	}
	b, _ = json.Marshal(data)
	ioutil.WriteFile(args["fileName"], b, 0644)
	return nil
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}

//=========================
//} else {
// //defer os.Remove(args["fileName"])
//		f.Seek(0, io.SeekStart)
//		f.Truncate(0)
//		var s0 []byte
//		f.Write([]byte("["))
//		for i, v := range s {
//			fmt.Println("v = ", v)
//			s0, err = json.Marshal(s[i])
//			if err != nil {
//				fmt.Println("Operation not allowed! %w", err)
//				return err
//			}
//		_, err := f.Write([]byte(s0))
//		if err != nil {
//			//panic(err)
//			return err
//		}
//	}
//	f.Write([]byte("]"))
//==============================
