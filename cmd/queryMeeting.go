package cmd

import (
	"Go-Agenda/entity"
	"fmt"

	"github.com/spf13/cobra"
)

// queryMeetingCmd represents the queryMeeting command
var queryMeetingCmd = &cobra.Command{
	Use:   "queryMeeting",
	Short: "query meeting by time",
	Long:  "query meeting between start time and end time",
	Run: func(cmd *cobra.Command, args []string) {
		// 检测是否已经登录
		if entity.IsLogin() {
			startTime, _ := cmd.Flags().GetString("starttime")
			endTime, _ := cmd.Flags().GetString("endtime")

			loginUser := entity.Loginuser.Name

			//检查输入的时间的合法性
			st := entity.StringToDate(startTime)
			et := entity.StringToDate(endTime)
			if !entity.IsDateValid(st) || !entity.IsDateValid(et) {
				fmt.Println("The startTime or endTime is not valid\nQuery Meeting Failed!")
				return
			}

			//获取符合条件的会议
			var queryMeetings []entity.Meeting
			for _, meet := range entity.GetMeetings() {
				if !(entity.IsEndBigthanStart(et, meet.Start) || entity.IsEndBigthanStart(meet.End, st)) {
					queryMeetings = append(queryMeetings, meet)
				}
			}

			// 输出会议的发起者，参与者，开始时间，结束时间
			fmt.Println("The list of meetings " + loginUser + " participated or sponsored are :")
			for _, meet := range queryMeetings {
				if meet.IsParticipator(loginUser) || meet.Sponsor == loginUser {
					fmt.Print("Title : " + meet.Title)
					fmt.Print("\tSponsor : " + meet.Sponsor)
					fmt.Print("\tParticipators : ")
					fmt.Print(meet.Participator)
					fmt.Print("\tStart time : " + meet.Start.DateToString())
					fmt.Println("\tEnd time :" + meet.End.DateToString())
				}
			}
		} else {
			fmt.Println("You must login first!")
		}
	},
}

func init() {
	rootCmd.AddCommand(queryMeetingCmd)
	queryMeetingCmd.Flags().StringP("starttime", "s", "", "the starting time(year-month-day/hour:min)")
	queryMeetingCmd.Flags().StringP("endtime", "e", "", "the end time(year-month-day/hour:mim)")
}
