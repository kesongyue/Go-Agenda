package main

import (
	// "fmt"
	"./entity"
)

func main() {

	// use Interface
	var u entity.USER
	u.UserFile = "/home/wjh/gowork/src/Agenda1/entity/Data/Users"
	var m entity.MEETING
	m.MeetingFile = "/home/wjh/gowork/src/Agenda1/entity/Data/Meetings"
	var json entity.JSON 

	json = &u
	json.ReadFromJsonFile()
	json.WriteToJsonFile()

	json = &m
	json.ReadFromJsonFile()
	json.WriteToJsonFile()

	// not use Interface
	// const userFile string = "/home/wjh/gowork/src/Agenda1/entity/Data/Users"
	// const meetingFile string = "/home/wjh/gowork/src/Agenda1/entity/Data/Meetings"
	
	// var users []entity.User = entity.UReadFromJsonFile(userFile);
	// entity.UWriteToJsonFile(users, userFile)

	// var meetings []entity.Meeting = entity.MReadFromJsonFile(meetingFile);
	// entity.MWriteToJsonFile(meetings, meetingFile)
}