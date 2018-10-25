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

type USER struct {
	UserFile string
	userByte []byte
	user User
	users []User
}	
type MEETING struct {
	MeetingFile string
	meetingByte []byte
	meeting Meeting
	meetings []Meeting
}

func UJsonDecode(_USER *USER) {
	err := json.Unmarshal((*_USER).userByte, &(*_USER).user)
	var check Check = mystring{"JsonDecode"}; check.CheckError(err)
}
func MJsonDecode(_MEETING *MEETING)  {
	err := json.Unmarshal((*_MEETING).meetingByte, &(*_MEETING).meeting)
	var check Check = mystring{"JsonDecode"}; check.CheckError(err)
}

func UJsonEncode(_USER *USER) {
	bytes, err := json.Marshal((*_USER).user)
	(*_USER).userByte = bytes;
	var check Check = mystring{"JsonEncode"}; check.CheckError(err)
}
func MJsonEncode(_MEETING *MEETING)  {
	bytes, err := json.Marshal((*_MEETING).meeting)
	(*_MEETING).meetingByte = bytes;
	var check Check = mystring{"JsonEncode"}; check.CheckError(err)
}

func (_USER *USER) ReadFromJsonFile() {
	inputFile, err := os.Open((*_USER).UserFile)
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
		UJsonDecode(_USER)
		(*_USER).users = append((*_USER).users, (*_USER).user)
	}
}
func (_MEETING *MEETING) ReadFromJsonFile() {
	inputFile, err := os.Open((*_MEETING).MeetingFile)
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
		MJsonDecode(_MEETING)
		(*_MEETING).meetings = append((*_MEETING).meetings, (*_MEETING).meeting)
	}
}

func (_USER *USER) WriteToJsonFile() {
	outputFile, err := os.OpenFile((*_USER).UserFile, os.O_RDWR | os.O_CREATE, 0)
	var check Check = mystring{"WriteToJsonFile"}; check.CheckError(err)
	defer outputFile.Close();

	for num := range (*_USER).users {
		UJsonEncode(_USER)
		var bytes []byte = append((*_USER).userByte, '\n')
		outputFile.Write(bytes)
		if num == 1 {

		}
	}
}
func (_MEETING *MEETING) WriteToJsonFile() {
	outputFile, err := os.OpenFile((*_MEETING).MeetingFile, os.O_RDWR | os.O_CREATE, 0)
	var check Check = mystring{"WriteToJsonFile"}; check.CheckError(err)
	defer outputFile.Close();

	for num := range (*_MEETING).meetings {
		MJsonEncode(_MEETING)
		var bytes []byte = append((*_MEETING).meetingByte, '\n')
		outputFile.Write(bytes)
		if num == 1 {
			
		}
	}
}