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

func UserJsonDecode(userByte []byte) User {
	var user User
	err := json.Unmarshal(userByte, &user)
	var check Check = mystring{"UserJsonDecode"}; check.CheckError(err)
	return user
}

func UReadFromJsonFile(userFile string) []User {
	var users []User
	inputFile, err := os.Open(userFile)
	var check Check = mystring{"ReadFromJsonFile"}; check.CheckError(err)
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	lineCounter := 0
	for {
		line, err := inputReader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}		
		lineCounter++;
		users = append(users, UserJsonDecode([]byte(line)))
	}
	return users
}

func UserJsonEncode(user User) []byte {
	var userByte []byte
	userByte, err := json.Marshal(user)
	var check Check = mystring{"UserJsonEncode"}; check.CheckError(err)
	return userByte
}

func UWriteToJsonFile(users []User, userFile string) {
	outputFile, err := os.OpenFile(userFile, os.O_RDWR | os.O_CREATE, 0)
	var check Check = mystring{"ReadFromJsonFile"}; check.CheckError(err)
	for num := range users {
		var userByte []byte = UserJsonEncode(users[num])
		userByte = append(userByte, '\n')
		outputFile.Write(userByte)
	}
	outputFile.Close();
}

func MeetingJsonDecode(meetingByte []byte) Meeting {
	var meeting Meeting
	err := json.Unmarshal(meetingByte, &meeting)
	var check Check = mystring{"MeetingJsonDecode"}; check.CheckError(err)
	return meeting
}

func MReadFromJsonFile(meetingFile string) []Meeting {
	var meetings []Meeting
	inputFile, err := os.Open(meetingFile)
	var check Check = mystring{"ReadFromJsonFile"}; check.CheckError(err)
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	lineCounter := 0
	for {
		line, err := inputReader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}		
		lineCounter++;
		meetings = append(meetings, MeetingJsonDecode([]byte(line)))
	}
	return meetings
}

func MeetingJsonEncode(meeting Meeting) []byte {
	var meetingByte []byte
	meetingByte, err := json.Marshal(meeting)
	var check Check = mystring{"UserJsonEncode"}; check.CheckError(err)
	return meetingByte
}

func MWriteToJsonFile(meetings []Meeting, meetingFile string) {
	outputFile, err := os.OpenFile(meetingFile, os.O_RDWR | os.O_CREATE, 0)
	var check Check = mystring{"ReadFromJsonFile"}; check.CheckError(err)
	for num := range meetings {
		var meetingByte []byte = MeetingJsonEncode(meetings[num])
		meetingByte = append(meetingByte, '\n')
		outputFile.Write(meetingByte)
	}
	outputFile.Close();
}