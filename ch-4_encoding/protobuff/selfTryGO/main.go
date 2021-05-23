package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/goyalayush57/DataIntensiveApplication/ch-4_encoding/protobuff/selfTryGO/protogenerated"
)

func main() {

	elliot := &protogenerated.Person{
		Name: "Elliot",
		Age:  24,
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// printing out our raw protobuf object
	fmt.Println(data)

	// let's go the other way and unmarshal
	// our byte array into an object we can modify
	// and use
	newElliot := &protogenerated.Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// print out our `newElliot` object
	// for good measure
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())

}
