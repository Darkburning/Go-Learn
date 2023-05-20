package main

import (
	"Go_Learn/ProtoBuf/protos"
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func main() {

	p := &protos.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*protos.PersonPhoneNumber{
			{Number: "555-4321", Type: protos.Person_HOME},
		},
	}
	// Marshal and WriteFile
	var FileName = "ProtoTestFile"
	out, err := proto.Marshal(p)
	if err != nil {
		log.Println("Unmarshal Failed!")
	}
	if err := os.WriteFile(FileName, out, 0644); err != nil {
		log.Println("WriteFile Failed!")
	}
	// ReadFile and Unmarshal
	in, err := os.ReadFile(FileName)
	if err != nil {
		log.Println("ReadFile Failed！")
	}
	fmt.Printf("Read Binary Data: %v\n", in)
	var rp protos.Person
	err = proto.Unmarshal(in, &rp)
	if err != nil {
		log.Println("Unmarshal Failed!")
	}
	fmt.Printf("ReadData: %v\n", rp)

}
