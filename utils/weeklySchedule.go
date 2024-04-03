package utils

import (
	"log"
	"strings"
	"time"
)


type WeeklySchedule struct {
	Day			string
	StartTime	time.Time
	EndTime		time.Time
}

func parseTime(timeArray []string) (time.Time) {
	timeString := timeArray[0] + timeArray[1]

	if len(timeString) > 4{
		parsedTime, err := time.Parse("2006-01-02 3:04pm", "1970-01-01 " + timeString)
		if err != nil {
			log.Fatal("Error Parsing Time: ", err)
		}

		return parsedTime;
		
	} else {
		parsedTime, err := time.Parse("2006-01-02 3pm", "1970-01-01 " + timeString)
		if err != nil {
			log.Fatal("Error Parsing Time: ", err)
		}

		return parsedTime;
	}
}

var weeksFullName = map[string]string{
	"Sat":		"Saturday",
	"Sun":		"Sunday",
	"Mon":		"Monday",
	"Tues":		"Tuesday",
	"Weds":		"Wednesday",
	"Thurs":	"Thursday",
	"Fri":		"Friday",
}

var weeksArray = []string{"Sat", "Sun", "Mon", "Tues", "Weds", "Thurs", "Fri", "Sat", "Sun", "Mon", "Tues", "Weds", "Thurs"}

// Mon - Weds, Sat
func parseWeeks(weekArray []string) ([]string) {
	refinedWeeks := []string{} 

	var lastWeek = ""
	for i := 0; i < len(weekArray); i++ {
		val, ok := weeksFullName[weekArray[i]]
		if ok {
			refinedWeeks = append(refinedWeeks, val)
		}else if weekArray[i] == "-" {
			var j = 0
			for ; j < len(weeksArray); j++ {
				if weeksArray[j] == lastWeek {break}
			}

			for j = j + 1; j < len(weeksArray); j++ {
				val, _:= weeksFullName[weeksArray[j]]
				refinedWeeks = append(refinedWeeks, val)

				if weekArray[i+1] == weeksArray[j] {break}
			}
			i++
		}

		lastWeek = weekArray[i]
	}

	return refinedWeeks
}


// "Mon - Weds, Sat 11 am - 10:15 pm / Thurs 8:30 am - 2:30 am / Fri 6 am - 3:15 am / Sun 3:45 pm - 4:15 pm"
// 1. Split by " / "
// 2. ParseTime(startTime, endTime)
// 3. ParseWeeks
func WeeklyScheduleDecoder(scheduleString string) ([]WeeklySchedule) {
	// Remove the ',' character
	schedulesFormating := strings.Replace(scheduleString, ",", "", -1)
	schedulesFormating = strings.Replace(schedulesFormating, " - ", "-", -1)
	schedulesFormating = strings.Replace(schedulesFormating, "-", " - ", -1)
	schedulesFormating = strings.Replace(schedulesFormating, "  /", " /", -1)
	schedules := strings.Split(schedulesFormating, " / ")

	var weeklySchedules []WeeklySchedule

	for i := 0; i < len(schedules); i++ {
		dates := strings.Split(schedules[i], " ")

		startTime := parseTime(dates[(len(dates) - 5):(len(dates) - 3)])
		endTime := parseTime(dates[(len(dates) - 2):(len(dates))])

		days := parseWeeks(dates[0:(len(dates) - 5)])

		for _, val := range days {
			weeklySchedules = append(weeklySchedules, WeeklySchedule{
				Day: val,
				StartTime: startTime,
				EndTime: endTime,
			})
		}
	}

	return weeklySchedules
}






