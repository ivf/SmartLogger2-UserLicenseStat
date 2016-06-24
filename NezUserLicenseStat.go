/*

Parsing the log file SmartLogger2 (DataService)
Search for minimum and maximum values are admin and lite license by each hour.
Display search progress. Output in tabular form to the console.
c: / programdata / speech technology center / smartlogger2 / dataservice / log / dataservice.log *

File Format:
ASCII text, with CRLF line terminators

String:
06/07/2016 08: 31: 08,813 [SessionsManager.Cleaner] DEBUG - Protection.PrintInfo (): Free Client licenses: 6 admins, 24 lites.

Options start:
1) Search the current date
./NezUserLicenseStat
2) Search the specific date
./NezUserLicenseStat -d=2016-06-14

*/

package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"path"
	"flag"
	"io/ioutil"
)

type lineLog struct {
	date  time.Time
	admin int
	lite  int
}

type processedLineLog struct {
	isFounded bool

	date time.Time
	hour int

	adminMin int
	adminMax int

	liteMax int
	liteMin int
}

var dirPath string = path.Dir("")	//Search in the local directory.
var statistics [24]processedLineLog        //Each hour
var searchDayFlag = flag.String("d", "", "Set the date in YYYY-MM-DD format.")
var searchDay time.Time = time.Now()
var programName string

//var t time.Time = time.Date(2016, 6, 5, 0, 0, 0, 0, time.UTC)

func init() {
	for i, _ := range statistics {
		statistics[i].hour = i
		statistics[i].adminMin = math.MaxInt32
		statistics[i].liteMin = math.MaxInt32
	}
}

func main() {
	flag.Parse()
	_, programName = path.Split(os.Args[0])

	if *searchDayFlag != "" {
		searchDay, _ = time.Parse("2006-01-02", *searchDayFlag)
	}

	scanDir()
	printStatistics()
}

func scanDir() {
	dir, _ := os.Open(dirPath)
	defer dir.Close()

	files, _ := dir.Readdir(-1)

	//var dayNext time.Time = searchDay.Add(24 * time.Hour)
	//fmt.Println(dayNext)
	//dayLastCount := 1
	for _, file := range files {
		/*if file.ModTime().Format("02-Jan-2006") == SEARCH_DAY.Format("02-Jan-2006") {
			//openFile(file)
			fmt.Println("#0")
			fmt.Println(dayNext)
		} else if file.ModTime().Format("02-Jan-2006") == (dayNext.Format("02-Jan-2006")) {
			//openFile(file)
			fmt.Println("#1")
			fmt.Println(dayNext)
		} else {
			dayNext = dayNext.Add(24 * time.Hour)
			fmt.Println("#2 ")
			fmt.Println(dayNext)
		}*/

		openFile(file)
	}
}

func openFile(file os.FileInfo) {
	if file.Name() == programName {
		return
	}

	fullPath := path.Join(dirPath, file.Name())
	data, _ := ioutil.ReadFile(fullPath)

	fmt.Printf("%s %s\n", file.Name(), file.ModTime().Format("02-Jan-2006"))
	fmt.Println("Loading...")

	processingFile(data)
}

func processingFile(data []byte) {
	start := time.Now()
	dataStr := string(data)
	dataStrLen := strings.Count(dataStr, "\n")

	for num, line := range strings.Split(dataStr, "\n") {
		showProgress(dataStrLen, num)

		if isLineContainLicInfo(line) {
			l := parseLine(line)

			if l.date.Day() != searchDay.Day() {
				continue
			}

			statistics[l.date.Hour()].isFounded = true
			statistics[l.date.Hour()].date = l.date

			if l.admin > statistics[l.date.Hour()].adminMax {
				statistics[l.date.Hour()].adminMax = l.admin
			}

			if l.admin < statistics[l.date.Hour()].adminMin {
				statistics[l.date.Hour()].adminMin = l.admin
			}

			if l.lite > statistics[l.date.Hour()].liteMax {
				statistics[l.date.Hour()].liteMax = l.lite
			}

			if l.lite < statistics[l.date.Hour()].liteMin {
				statistics[l.date.Hour()].liteMin = l.lite
			}
		}
	}

	fmt.Printf(" %.2fsec\n", time.Since(start).Seconds())
}

func showProgress(len int, num int) {
	procent := (num * 100) / len

	fmt.Print("\r")
	fmt.Printf("%v/%v %v%v", num, len, procent, "%")

}

func isLineContainLicInfo(str string) bool {
	isMatched, _ := regexp.MatchString("Free Client licenses", str)
	return isMatched
}

func parseLine(str string) lineLog {
	date := getTime(str)
	admin, lite := getAdminLite(str)
	l := lineLog{date, admin, lite}

	return l
}

func getTime(str string) time.Time {
	year, _ := strconv.Atoi(str[0:4])
	monthInt, _ := strconv.Atoi(str[5:7])
	month := time.Month(monthInt)
	day, _ := strconv.Atoi(str[8:10])
	hour, _ := strconv.Atoi(str[11:13])
	min, _ := strconv.Atoi(str[14:16])
	sec, _ := strconv.Atoi(str[17:19])
	nsec, _ := strconv.Atoi(str[20:23])
	date := time.Date(year, month, day, hour, min, sec, nsec, time.UTC)

	return date
}

func getAdminLite(str string) (int, int) {
	adminBegin := strings.Index(str, "licenses: ") + 10
	adminEnd := strings.Index(str, " admins")
	adminStr := str[adminBegin:adminEnd]
	admin, _ := strconv.Atoi(adminStr)

	liteBegin := strings.Index(str, "admins, ") + 8
	liteEnd := strings.Index(str, " lites")
	liteStr := str[liteBegin:liteEnd]
	lite, _ := strconv.Atoi(liteStr)

	return admin, lite
}

func printStatistics() {
	fmt.Printf("%11s%5s%9s%9s%8s%8s\n", "Date", "Hour", "AdminMax", "AdminMin", "LiteMax", "LiteMin")

	for i := 0; i <= 23; i++ {
		s := statistics[i]

		if !s.isFounded {
			fmt.Printf("%v%5v\n", searchDay.Format("02-Jan-2006"), i)
			continue
		}

		formatStr := "%v%5v%9v%9v%8v%8v\n"
		fmt.Printf(formatStr, s.date.Format("02-Jan-2006"), s.hour,
			s.adminMax, s.adminMin, s.liteMax, s.liteMin)
	}
}

/*
func printStatistics() {
	for i := 0; i <= 23; i++ {
		s := statistics[i]

		if (!s.isFounded) {
			fmt.Printf("Day=%v and hour=%v are not found.\n", SEARCH_DAY, i)
			continue
		}

		formatStr := "%v.%v.%v %v %v %v %v %v\n"
		fmt.Printf(formatStr, s.date.Day(), s.date.Month(), s.date.Year(), s.hour,
			s.adminMax, s.adminMin, s.liteMax, s.liteMin)
	}
}*/
