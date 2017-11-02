package main

import (
	"fmt"
)

func main() {
	var dates []string
	var dates_size int
	fmt.Scanf("%d", &dates_size)

	for i := 0; i < dates_size; i++ {
		var dates_item string
		fmt.Scanf("%s", &dates_item)
		dates = append(dates, dates_item)

	}

	// 아래와 같은 형식이 되어야 함
	//for i := 0; i < dates_size; i++ {
	//	var d1, d2, d3 string
	//	fmt.Scanf("%s %s %s", &d1, &d2, &d3)
	//	dates = append(dates, fmt.Sprintf("%s %s %s", d1, d2, d3))
	//}


	fmt.Println(dates)
	fmt.Println()

	var res []string = reformatDate(dates)
	for res_i := range res {
		fmt.Println(res[res_i])
	}
}

var MONTH = map[string]string {
	"Jan": "01",
	"Feb": "02",
	"Mar": "03",
	"Apr": "04",
	"May": "05",
	"Jun": "06",
	"Jul": "07",
	"Aug": "08",
	"Sep": "09" ,
	"Oct": "10",
	"Nov": "11",
	"Dec": "12",
}

var DAY = map[string] string {
	"1st":"01",
	"2nd":"02",
	"3rd":"03",
	"4th":"04",
	"5th":"05",
	"6th":"06",
	"7th":"07",
	"8th":"08",
	"9th":"09",
	"10th":"10",
	"11st":"11",
	"12nd":"12",
	"13rd":"13",
	"14th":"14",
	"15th":"15",
	"16th":"16",
	"17th":"17",
	"18th":"18",
	"19th":"19",
	"20th":"20",
	"21st":"21",
	"22nd":"22",
	"23rd":"23",
	"24th":"24",
	"25th":"25",
	"26th":"26",
	"27th": "27",
	"28th": "28",
	"29th": "29",
	"30th": "30",
	"31st": "31",
}

func reformatDate(dates []string) []string {

	formattedDates := make([]string, len(dates))

	for i := 0; i < len(dates); i++ {
		formattedDates[i] = transform(split(dates[i]))
	}

	return formattedDates
}

// strings 패키지를 사용할 수 없는 제약이 있음
func split(date string) []string {
	dateRunes := []rune(date)
	splitIndex := 0
	for index, rune := range dateRunes {
		// whitespace: 32
		if rune == 32 {
			splitIndex = index
			break
		}
	}

	dayRunes := dateRunes[:splitIndex]
	monthRunes := dateRunes[splitIndex+1:splitIndex+4]
	yearRunes := dateRunes[splitIndex+5:]

	return []string{string(dayRunes), string(monthRunes), string(yearRunes)}
}

func transform(date []string) string {
	day := DAY[date[0]]
	month := MONTH[date[1]]
	year := date[2]
	return  fmt.Sprintf("%s-%s-%s", year,month,day)
}
