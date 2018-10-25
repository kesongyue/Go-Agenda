package entity

import (
	"encoding/json"
	"fmt"
	"bufio"
	"io"
	"os"
)

type Check interface {
	CheckError(err error)
}

type mystring struct {
	name string
}

func (funcName mystring) CheckError(err error) {
	if err != nil {
		fmt.Println("Can not run %s", funcName.name)
	    panic(err)
	}
}

type JSON interface {
	ReadFromJsonFile()
	WriteToJsonFile()
}
type JSONCode interface {
	JsonDecode()
	JsonEncode()
}

type USER struct {
	userByte []byte
	user User
	users []User
}	
type MEETING struct {
	meetingByte []byte
	meeting Meeting
	meetings []Meeting
}	

func (_USER *USER) JsonDecode() {
	err := json.Unmarshal((*_USER).userByte, &(*_USER).user)
	var check Check = mystring{"JsonDecode"}; check.CheckError(err)
}
func (_MEETING *MEETING) JsonDecode()  {
	err := json.Unmarshal((*_MEETING).meetingByte, &(*_MEETING).meeting)
	var check Check = mystring{"JsonDecode"}; check.CheckError(err)
}

func (_USER *USER) JsonEncode() {
	bytes, err := json.Marshal((*_USER).user)
	(*_USER).userByte = bytes;
	var check Check = mystring{"JsonEncode"}; check.CheckError(err)
}
func (_MEETING *MEETING) JsonEncode() {
	bytes, err := json.Marshal((*_MEETING).meeting)
	(*_MEETING).meetingByte = bytes;
	var check Check = mystring{"JsonEncode"}; check.CheckError(err)
}

func (_USER *USER) ReadFromJsonFile() {
	inputFile, err := os.Open("/home/wjh/gowork/src/Agenda1/entity/Data/Users")
	var check Check = mystring{"ReadFromJsonFile"}; check.CheckError(err)
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	lineCounter := 0
	for {
		line, err := inputReader.ReadString('\n')
		(*_USER).userByte = []byte(line)
		if err != nil || err == io.EOF {
			break
		}		
		lineCounter++;
		var jsoncode JSONCode = (*_USER)
		jsoncode.JsonDecode()
		(*_USER).users = append((*_USER).users, (*_USER).user)
	}
}
func (_MEETING *MEETING) ReadFromJsonFile() {
	inputFile, err := os.Open("/home/wjh/gowork/src/Agenda1/entity/Data/Meetings")
	var check Check = mystring{"ReadFromJsonFile"}; check.CheckError(err)
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	lineCounter := 0
	for {
		line, err := inputReader.ReadString('\n')
		(*_MEETING).meetingByte = []byte(line)
		if err != nil || err == io.EOF {
			break
		}		
		lineCounter++;
		var jsoncode JSONCode = *_MEETING
		jsoncode.JsonDecode()
		(*_MEETING).meetings = append((*_MEETING).meetings, (*_MEETING).meeting)
	}
}

func (_USER *USER) WriteToJsonFile() {
	outputFile, err := os.OpenFile("/home/wjh/gowork/src/Agenda1/entity/Data/UsersOutput", os.O_RDWR | os.O_CREATE, 0)
	var check Check = mystring{"WriteToJsonFile"}; check.CheckError(err)
	defer outputFile.Close();

	for num := range (*_USER).users {
		var jsoncode JSONCode = *_USER
		jsoncode.JsonEncode()
		var bytes []byte = append((*_USER).userByte, '\n')
		outputFile.Write(bytes)
	}
}
func (_MEETING *MEETING) WriteToJsonFile() {
	outputFile, err := os.OpenFile("/home/wjh/gowork/src/Agenda1/entity/Data/MeetingsOutput", os.O_RDWR | os.O_CREATE, 0)
	var check Check = mystring{"WriteToJsonFile"}; check.CheckError(err)
	defer outputFile.Close();

	for num := range (*_MEETING).meetings {
		var jsoncode JSONCode = *_MEETING
		jsoncode.JsonEncode()
		var bytes [] bytes = append((*_MEETING).meetingByte, '\n')
		outputFile.Write(bytes)
	}
}