package etl

import (
	"fmt"

	"gitlab.com/IstiyakRiyad/technical-assessment-pathao/utils"
)


func TransformData(restaurants []Restaurant, users []User) {
	weeklySchedule := utils.WeeklyScheduleDecoder("Mon 6 am - 8:45 am / Tues 2 pm - 3 pm / Sun, Weds - Fri 5 am - 9 am / Sat 11:30 am - 3:15 am")

	fmt.Println(weeklySchedule)
}
