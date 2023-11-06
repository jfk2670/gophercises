package main

import "fmt"

func HumanReadableTime(seconds int) string {
	m,seconds := seconds / 60, seconds % 60
	h, m := m / 60, m % 60

	return fmt.Sprintf("%02d:%02d:%02d",h,m,seconds)
}

func main () {
	fmt.Println(HumanReadableTime(0)) // (Equal("00:00:00"))
	fmt.Println(HumanReadableTime(59)) // (Equal("00:00:59"))
	fmt.Println(HumanReadableTime(60)) // (Equal("00:01:00"))
	fmt.Println(HumanReadableTime(90)) // (Equal("00:01:30"))
	fmt.Println(HumanReadableTime(3599)) // (Equal("00:59:59"))
	fmt.Println(HumanReadableTime(3600)) // (Equal("01:00:00"))
	fmt.Println(HumanReadableTime(3700)) // (Equal("01:01:40"))
	fmt.Println(HumanReadableTime(45296)) // (Equal("12:34:56"))
	fmt.Println(HumanReadableTime(86399)) // (Equal("23:59:59"))
	fmt.Println(HumanReadableTime(86400)) // (Equal("24:00:00"))
	fmt.Println(HumanReadableTime(359999)) // (Equal("99:59:59"))
}