package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"

)

var result []string

func main() {
	fmt.Printf("\r\n------------------------------------------------------------------------------\r\n");
	fmt.Printf(" Welcome to the real world!                                     qq:1141056911\r\n");
	fmt.Printf("                                                                       By Lcy \r\n");
	fmt.Printf("                                                             http:/phpinfo.me \r\n");
	fmt.Printf("------------------------------------------------------------------------------\r\n");
	if len(os.Args) != 4 {
		fmt.Printf("Example: %s http://phpinfo.me/ 20 PHP\r\n", os.Args[0])
		fmt.Printf("Example: %s http://phpinfo.me/ 20 ASP\r\n", os.Args[0])
		fmt.Printf("Example: %s http://phpinfo.me/ 20 ASPX\r\n", os.Args[0])
		fmt.Printf("Example: %s http://phpinfo.me/ 20 DIR\r\n", os.Args[0])
		fmt.Printf("Example: %s http://phpinfo.me/ 20 MDB\r\n", os.Args[0])
		fmt.Printf("Example: %s http://phpinfo.me/ 20 JSP\r\n", os.Args[0])
		os.Exit(1)
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
	threa := os.Args[2]
	thread, _ := strconv.Atoi(threa)
	url := os.Args[1]
	scan_type :=strings.ToUpper( os.Args[3])
	var file string
	switch scan_type{
		case "PHP":
			file ="PHP.txt"
		case "ASP":
			file = "ASP.txt"
		case "JSP":
			file = "JSP.txt"
		case "ASPX":
			file="ASPX.txt"
		case "DIR":
			file = "DIR.txt"
		case "MDB":
			file = "MDB.txt"
		default:
			fmt.Println("\r\n你逗我呢\r\n")
			os.Exit(1)
	}
	fs, e := readfile(file)
	if e != nil {
		e.Error()
	}
	arr := strings.Split(string(fs), "\n")
	//字典长度
	lens := len(arr)
	//每个线程分配任务数

	task := lens / thread

	ch := make(chan int)

	for i := 0; i < thread; i++ {
		go run(url, arr, i, task, ch)
	}
	<-ch

}

func run(urls string, dir []string, tnum int, task int, ch chan int) {
	for i := tnum*task + 1; i < (tnum*task)+task; i++ {
		dir[i-1] = strings.TrimSpace(dir[i-1])
		url := urls + dir[i-1]
		code, err := scandir(url)
		if err != nil {
			continue
		}
		if code == 403 || code == 404 {
			fmt.Printf("                                                                              \r")
			fmt.Printf("Checking: %s ...\r",dir[i-1])
		} else {
			fmt.Printf("Found: %s [%d]!!!\r\n",dir[i-1],code)
			result = append(result,dir[i-1])
		}

	}
	ch <- 1
}

func readfile(dir string) ([]byte, error) {
	f, err := os.Open(dir)
	if err != nil {
		err.Error()
	}
	return ioutil.ReadAll(f)
}
func scandir(url string) (int, error) {
	resp, err := http.Get(url)
	var status int
	if err != nil {
		status = 404
	} else {
		status = resp.StatusCode
	}

	return status, err
}
