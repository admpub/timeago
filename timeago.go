package timeago

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// Take coverts given datetime into `x time ago` format.
// For displaying `Online` word if date interval within
// 60 seconds, add `|online` flag to the datetime string.
// Format must be [year-month-day hours:minutes:seconds}
func Take(datetime string, langs ...string) string {
	lang, format, option, hasOption := getOption(&datetime)
	seconds := getSeconds(datetime, format)
	if len(langs) > 0 {
		lang = langs[0]
	}
	switch {
	case seconds < 0 && option == "online":
		return trans("online", lang)
	case seconds < 0:
		return getWords("seconds", 0, lang)
	}

	return calculateTheResult(seconds, hasOption, option, lang)
}

func getSeconds(datetime, format string) (seconds int) {
	if len(format) == 0 {
		format = "2006-01-02 15:04:05"
	}
	if loc != nil {
		parsedTime, _ := time.ParseInLocation(format, datetime, loc)
		seconds = int(time.Now().In(loc).Sub(parsedTime).Seconds())
	} else {
		parsedTime, _ := time.ParseInLocation(format, datetime, time.Local)
		seconds = int(time.Since(parsedTime).Seconds())
	}
	return
}

func calculateTheResult(seconds int, hasOption bool, option string, lang string) string {
	minutes, hours, days, weeks, months, years := getTimeCalculations(float64(seconds))

	switch {
	case hasOption && option == "online" && seconds < 60:
		return trans("online", lang)
	case seconds < 60:
		return getWords("seconds", seconds, lang)
	case minutes < 60:
		return getWords("minutes", minutes, lang)
	case hours < 24:
		return getWords("hours", hours, lang)
	case days < 7:
		return getWords("days", days, lang)
	case weeks < 4:
		return getWords("weeks", weeks, lang)
	case months < 12:
		if months == 0 {
			months = 1
		}

		return getWords("months", months, lang)
	}

	return getWords("years", years, lang)
}

func getTimeCalculations(seconds float64) (int, int, int, int, int, int) {
	minutes := math.Round(seconds / 60)
	hours := math.Round(seconds / 3600)
	days := math.Round(seconds / 86400)
	weeks := math.Round(seconds / 604800)
	months := math.Round(seconds / 2629440)
	years := math.Round(seconds / 31553280)

	return int(minutes), int(hours), int(days), int(weeks), int(months), int(years)
}

// get the last number of a given integer
func getLastNumber(num int) int {
	numStr := strconv.Itoa(num)
	result, _ := strconv.Atoi(numStr[len(numStr)-1:])

	return result
}

// getWords decides rather the word must be singular or plural,
// and depending on the result it adds the correct word after
// the time number
func getWords(timeKind string, num int, lang string) string {
	lastNum := getLastNumber(num)
	index := 2

	switch {
	case lastNum == 1 && num == 11:
		index = 2
	case lastNum == 1 && language == "ru" || num == 1 && language == "en":
		index = 0
	case lastNum > 1 && lastNum < 5:
		index = 1
	}

	timeTrans := getTimeTranslations(lang)
	format := trans("format", lang)
	if len(format) > 0 && format != `format` {
		return fmt.Sprintf(format, strconv.Itoa(num), timeTrans[timeKind][index], trans("ago", lang))
	}
	return strconv.Itoa(num) + " " + timeTrans[timeKind][index] + " " + trans("ago", lang)
}

// getOption check if datetime has option with time,
// if yes, it will return this option and remove it
// from datetime
func getOption(datetime *string) (string, string, string, bool) {
	date := *datetime
	spittedDateString := strings.Split(date, "|")

	var (
		option    string
		format    string
		lang      string
		hasOption bool
	)
	size := len(spittedDateString)
	if size > 1 {
		*datetime = spittedDateString[0]
		if len(spittedDateString[1]) > 0 {
			option = spittedDateString[1]
			hasOption = true
		}
		if size > 2 {
			format = spittedDateString[2]
		}
		if size > 3 {
			lang = spittedDateString[3]
		}
	}

	return lang, format, option, hasOption
}
