package main
import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	var time string
	fmt.Scanf("%v", &time)
	if len(time) == 10 {
		hours_val,err := strconv.Atoi(time[0:2])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		if hours_val > 0 {
			if am_pm := time[8:10]; am_pm == "PM" {
				if hours_val < 12 {
					hours_val+=12
					fmt.Println(strconv.Itoa(hours_val) + time[2:len(time)-2])	
				} else {
					fmt.Println("12" + time[2:len(time)-2])	
				}
			} else if am_pm == "AM" {
				if hours_val == 12 {
					fmt.Println("00" + time[2:len(time)-2])
				} else {
					fmt.Println(time[0:len(time)-2])
				}
			}	
		}
	}
}