package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/goyalayush57/DataIntensiveApplication/ch-4_encoding/protobuff/selfTryGO/nestedExample/protogenerated2"
)

func main() {
	fmt.Println("Hello World")
	getData("secretData")
	if 1 == 1 {
		return
	}
	ashish := &protogenerated2.Person{
		Name: "Ashwariya",
		Age:  100,
		SocialFollowers: []*protogenerated2.SocialFollowers{
			&protogenerated2.SocialFollowers{
				Youtube: 1,
				Twitter: 2,
			},
			&protogenerated2.SocialFollowers{
				Youtube: 3,
				Twitter: 4,
			},
		},
	}
	data, err := proto.Marshal(ashish)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	f, err := os.Create("secretData")
	n2, err := f.Write(data)
	fmt.Println("Number of bytes written", n2)
	defer f.Close()
	//json
	// f, err = os.Create("jsonData")
	// n2, err = f.Write(createJsonFile(*ashish))
	// fmt.Println("Number of bytes written", n2)
	// getData("secretData")
	// getData("jsonData")
}

func createJsonFile(p protogenerated2.Person) []byte {
	dat, err := json.Marshal(p)
	if err != nil {
		log.Fatal("json marshalling: ", err)
	}
	return dat
}
func getData(filePath string) protogenerated2.Person {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	person := &protogenerated2.Person{}
	err = proto.Unmarshal(dat, person)
	fmt.Printf("%+v", person)
	return *person
}
