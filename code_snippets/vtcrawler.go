package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/williballenthin/govt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

var apikey string
var apiurl string
var domain string
var ip string

//Initialization Function
func init() {
	flag.StringVar(&apikey, "apikey", os.Getenv("VT_API_KEY"), "VirusTotal APIK")
	flag.StringVar(&apiurl, "apiurl", "https://www.virustotal.com/vtapi/v2/", "URL of the VirusTotal API to be used.")
}

/*
*
 */
func main() {
	var scan_url string
	var verbose bool
	var silent bool
	//	var counter int
	var url_array [100]string
	var cli_str = flag.String("u", "", "Individual URL to be scanned")
	var cli_verb = flag.String("m", "v,s", "Verbose Mode (v) or Silent Mode (s)")
	var cli_file = flag.String("i", "urls.txt", "File of URLs to be scanned")
	var ii int
	flag.Parse()

	fmt.Println("VT URl Scanner")

	if *cli_str != "" && len(*cli_str) > 0 {
		scan_url = *cli_str
		if strings.EqualFold(*cli_verb, "v") {
			fmt.Println("Verbose Mode Selected")
			verbose = true
			fetch(scan_url, verbose)
		} else if strings.EqualFold(*cli_verb, "s") {
			fmt.Println("Silent Mode Running")
			silent = true
			fetch(scan_url, silent)
		}
	} else if *cli_file != "" && len(*cli_file) > 0 {

		file, err := os.Open("urls.txt")
		if err != nil {
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		ii = 0
		for scanner.Scan() {
			url_array[ii] = scanner.Text()
			ii++
		}
		var counter int
		counter = ii
		if strings.EqualFold(*cli_verb, "v") {
			fmt.Println("Verbose Mode Selected")
			fmt.Println("Reading File")
			verbose = true
			multi_fetch(url_array, "verbose", counter)
			/*	for counter = 0; counter < ii; counter++ {
				fmt.Println("urls", url_array[counter])
			}*/
			multi_fetch(url_array, "verbose", counter)
		} else if strings.EqualFold(*cli_verb, "s") {
			fmt.Println("Silent Mode Running")
			silent = true
			multi_fetch(url_array, "silent", counter)
		}
	}

}

func fetch(url string, mode bool) {
	fmt.Println("inside of fetch")
	var ip string
	file := "report.txt"
	if mode {
		fmt.Print("Analyzing URL(s):\n")
	}
	ip_addr, err := net.LookupIP(url)
	if mode {
		fmt.Print("Finished Domain Lookup\n")
	}
	if err != nil {
		fmt.Sprintf("ip lookup failed %s %v", ip_addr, err)
	}
	for i := 0; i < len(ip_addr); i++ {
		ip = ip_addr[i].String()
	}
	if mode {
		fmt.Print("Sending to VirusTotal: Awaiting Results\n")
	}
	if ip == "" {
		fmt.Println("-ip=<ip> fehlt!")
		os.Exit(1)
	}
	c := govt.Client{Apikey: apikey, Url: apiurl}

	// get a file report
	r, err := c.GetIpReport(ip)
	check(err)
	j, err := json.MarshalIndent(r, "", "    ")
	if _, err := os.Stat(file); err == nil {
		fmt.Println("File Exists")
		ioutil.WriteFile("test.txt", j, 0664)
	}
	ioutil.WriteFile("report.txt", j, 0664)
	check(err)

	if mode {
		fmt.Print("Report Generated\n")
		fmt.Println("IP Report:")
	}
}

func multi_fetch(url [100]string, mode_str string, size int) {
	fmt.Println("inside of multi_fetch")
	var aa int
	var mode bool
	mode = false

	if strings.EqualFold(mode_str, "verbose") {
		mode = true
	}
	for aa = 0; aa < size; aa++ {
		if url[aa] != "" {
			go fetch(url[aa], mode)
		}
	}

}

// check - an error checking function
func check(e error) {
	if e != nil {
		panic(e)
	}
}
