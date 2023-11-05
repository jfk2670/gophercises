//Your task in order to complete this Kata is to write a function which formats a duration, given as a number of seconds, in a human-friendly way.
// The function must accept a non-negative integer. If it is zero, it just returns "now". Otherwise, the duration is expressed as a combination of years, days, hours, minutes and seconds.
//
// It is much easier to understand with an example:
//
// * For seconds = 62, your function should return 
//     "1 minute and 2 seconds"
// * For seconds = 3662, your function should return
//     "1 hour, 1 minute and 2 seconds"

package main

import "fmt"
import "strings"

func FormatDuration(seconds int64) string {
	const MINUTE int64 = 60
	const HOUR int64 = 3600
	const DAY int64 = 86400
	const YEAR int64 = 31536000

	var minutes, hours, days, years int64 = 0, 0, 0, 0

	if seconds == 0 {
		return "now"
	} else if seconds > 0 && seconds < MINUTE {
		return formatResponse(seconds, minutes, hours, days, years)
	} else if seconds >= MINUTE && seconds < HOUR {
		minutes := seconds / MINUTE
		seconds -= minutes * MINUTE
		
		return formatResponse(seconds, minutes, hours, days, years)
	} else if seconds >= HOUR && seconds < DAY {
		hours := seconds / HOUR
		seconds -= hours * HOUR
		
		minutes := seconds / MINUTE
		seconds -= minutes * MINUTE
		
		return formatResponse(seconds, minutes, hours, days, years)
	} else if seconds >= DAY && seconds < YEAR {
		days := seconds / DAY
		seconds -= days * DAY

		hours := seconds / HOUR
		seconds -= hours * HOUR
		
		minutes := seconds / MINUTE
		seconds -= minutes * MINUTE
		
		return formatResponse(seconds, minutes, hours, days, years)
	} else if seconds >= YEAR {
		years := seconds / YEAR
		seconds -= years * YEAR

		days := seconds / DAY
		seconds -= days * DAY

		hours := seconds / HOUR
		seconds -= hours * HOUR
		
		minutes := seconds / MINUTE
		seconds -= minutes * MINUTE
		
		return formatResponse(seconds, minutes, hours, days, years)
	}

	return ""
}

func formatResponse(seconds, minutes, hours, days, years int64) string {
	var response string
	var firstUnit bool = false

	if years > 0 {
		firstUnit = true
		response += fmt.Sprintf("%d year", years)
		if years > 1 {
			response += "s"
		}
	}

	if firstUnit && days > 0{
		response += ", "
	}

	if days > 0 {
		firstUnit = true
		response += fmt.Sprintf("%d day", days)
		if days > 1 {
			response += "s"
		}
	}

	if firstUnit && hours > 0{
		response += ", "
	}

	if hours > 0 {
		firstUnit = true
		response += fmt.Sprintf("%d hour", hours)
		if hours > 1 {
			response += "s"
		}
	}

	if firstUnit && minutes > 0{
		response += ", "
	}

	if minutes > 0 {
		firstUnit = true
		response += fmt.Sprintf("%d minute", minutes)
		if minutes > 1 {
			response += "s"
		}
	}

	if firstUnit && seconds > 0{
		response += ", "
	}

	if seconds > 0 {
		response += fmt.Sprintf("%d second", seconds)
		if seconds > 1 {
			response += "s"
		}
	}

	commas := strings.Count(response, ",")

	if commas == 1 {
		response = strings.Replace(response, ",", " and", 1)
	} else if commas >= 2 {
		target := strings.LastIndex(response, ",")
		response = response[:target] + " and" + response[target+1:]
	}


	return response
}

func main () {
	fmt.Println(FormatDuration(0))
	fmt.Println(FormatDuration(62))
	fmt.Println(FormatDuration(820)) //.To(Equal("2 minutes"))
	fmt.Println(FormatDuration(3600)) //.To(Equal("1 hour"))
	fmt.Println(FormatDuration(3662)) //.To(Equal("1 hour, 1 minute and 2 seconds"))
	fmt.Println(FormatDuration(8662))
	fmt.Println(FormatDuration(86102))
	fmt.Println(FormatDuration(86502))
	fmt.Println(FormatDuration(288702))
	fmt.Println(FormatDuration(873333098)) //.To(Equal("2 minutes"))
}