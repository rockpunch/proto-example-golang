package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/rockpunch/proto-example-golang/protos"
	"io/ioutil"
	"log"
)

func main() {
	sm := doSimple()

	readAndWriteDemo(sm)
	jsonDemo(sm)
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)

	fmt.Println(smAsString)

	sm = &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm)

	fmt.Println(sm)
}

func toJSON(message proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	str, err := marshaler.MarshalToString(message)
	if err != nil {
		log.Fatalln("Failed convert to JSON", err)
		return ""
	}
	return str
}

func fromJSON(s string, pb proto.Message) {
	err := jsonpb.UnmarshalString(s, pb)
	HandleError(err)
}

func readAndWriteDemo(sm proto.Message) {

	err := writeToFile("simple.bin", sm)
	HandleError(err)

	sm2 := &simplepb.SimpleMessage{}

	err = readFromFile("simple.bin", sm2)
	HandleError(err)

	fmt.Println(sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	bs, err := proto.Marshal(pb)
	HandleError(err)

	err = ioutil.WriteFile(fname, bs, 0644)
	HandleError(err)

	fmt.Println("Data has been written!")
	return err
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	HandleError(err)

	err = proto.Unmarshal(in, pb)
	HandleError(err)

	return err
}

func doSimple() *simplepb.SimpleMessage {

	sm := simplepb.SimpleMessage{
		Id:         0,
		IsSimple:   true,
		Name:       "James Bond",
		SampleList: []int32{1, 2, 3, 4, 5},
	}

	fmt.Printf("Value= %v\n", sm)

	sm.Name = "I renamed you"
	fmt.Printf("Renamed value= %v\n", sm)

	return &sm
}

// Handle error
func HandleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
