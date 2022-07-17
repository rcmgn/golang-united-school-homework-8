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
	id    string
	email string
	age   int
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

func Perform(args Arguments, writer io.Writer) error {
	var data []User
	if args["operation"] == "" {
		return errors.New("-operation flag has to be specified")
	}
	if args["fileName"] == "" {
		return errors.New("-fileName flag has to be specified")
	}
	//data = append(data, User{"1", "email@test.com", 23}, User{"2", "test2@test.com", 41})
	//fmt.Println(data)
	//b, err := json.Marshal(data)
	//fmt.Println(b, err)
	//ioutil.WriteFile("my.json", b, 0644)

	b, _ := ioutil.ReadFile(args["fileName"])
	//fmt.Println(b, err)
	json.Unmarshal(b, &data)
	//fmt.Println(err2)
	//fmt.Println(data)

	//plan = []byte("[{\"id\":\"1\",\"email\":\"test@test.com\",\"age\":34}]")
	//s, _ := strconv.Unquote(string(plan))

	switch args["operation"] {
	case "list":
		writer.Write(b)
	case "add":
		if args["item"] == "" {
			return errors.New("-item flag has to be specified")
		}
		fmt.Println(args["item"])
	case "findById":
		if args["id"] == "" {
			return errors.New("-id flag has to be specified")
		}
	case "remove":
		if args["id"] == "" {
			return errors.New("-id flag has to be specified")
		}
		writer.Write([]byte("Item with id " + args["id"] + " not found"))
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
