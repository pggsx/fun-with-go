package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/williballenthin/govt"
	"io/ioutil"
	"net"
	"os"
)

// KEA -Go uses camelCase instead of spacing_like_this

/*
KEA - A good design would be to populate a []string of URL's in your main function. Then just loop and call fetch. It would clean up your "if cliStr/cliFile" logic, and allow you to get rid of the if/else. You could simply append to the []string if the flag is given. Nice, clean and concise.

Pseudocode:

urls := make([]string, 0)

concurrent := flag.Bool("j", false, "Concurrently fetch url reports")

if *cliStr != "" {
	// Do whatever validity checks you want
	urls = append(urls, *cliStr)
}

if *cliFile != "" {
	// Open file
	// For line in file - urls = append(urls, *cliStr)
}

for _, url := range urls {
	if *concurrent {
		go fetch()
	} else {
		fetch()
	}
}
*/

var apikey string
var apiurl string
var domain string
var ip string

//Initialization Function
func init_var() {
	flag.StringVar(&apikey, "apikey", os.Getenv("VT_API_KEY"), "VirusTotal APIK")
	flag.StringVar(&apiurl, "apiurl", "https://www.virustotal.com/vtapi/v2/", "URL of the VirusTotal API to be used.")
}

/*
*
 */
func main() {
	init_var()

	urls := make([]string, 0, 100)
	var verb bool
	cliStr := flag.String("u", "", "Individual URL to be scanned")
	cliFile := flag.String("i", "urls.txt", "File of URLs to be scanned")
	concurrent := flag.Bool("j", false, "Concurrently fetch url reports")
	silent := flag.Bool("s", false, "Enable silent mode")
	verbose := flag.Bool("v", false, "Enable verbose mode")

	flag.Parse()

	if *silent == false && *verbose == false {
		fmt.Println("Unable to Determine Mode Selection Terminating Now")
		os.Exit(-2)
	}

	if *silent == true && *verbose == true {
		fmt.Println("Verbose and Silent Mode Detected!")
		fmt.Println("User must either choose Verbose or Silent Mode")
		fmt.Println("Terminating Now")
		os.Exit(-2)
	}
	if *concurrent == false && *silent == false && *verbose == false {
		fmt.Println("Invalid Number of Arguments Specified Terminating Now...")
		os.Exit(-2)
	}

	fmt.Println("VT URl Scanner")

	if *cliStr != "" {
		urls = append(urls, *cliStr)

		if *verbose {
			fmt.Println("Verbose Mode Selected")
		} else if *silent {
			fmt.Println("Silent Mode Running")
		}

	} else if *cliFile != "" {

		file, err := os.Open("urls.txt")
		if err != nil {
			fmt.Println("FNF Error Exiting..")
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		ii := 0
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
			ii++
		}
		if *verbose {
			fmt.Println("Verbose Mode Selected")
			fmt.Println("Reading File")
			verb = true
			/*	for counter = 0; counter < ii; counter++ {
				fmt.Println("urls", urls[counter])
			}*/
		} else if *silent {
			fmt.Println("Silent Mode Running Ouput will be txt format")
			verb = false
		}
	}

	for _, url := range urls {
		if *concurrent {
			go fetch(url, verb)
		} else {
			fetch(url, verb)
		}

	}
}
func fetch(url string, mode bool) {
	var ip string
	file := "report.txt"

	if mode {
		fmt.Print("Analyzing URL(s):\n")
		fmt.Print("Resolving URL:")
	}
	if govalidator.IsURL(url) {

		ipAddr, err := net.LookupIP(url)
		if mode {
			fmt.Print("Finished Domain Lookup\n")
		}
		if err != nil {
			fmt.Sprintf("ip lookup failed %s %v", ipAddr, err)
		}
		for i := 0; i < len(ipAddr); i++ {
			ip = ipAddr[i].String()
		}
		if mode {
			fmt.Print("Sending to VirusTotal: Awaiting Results\n")
		}
		if ip == "" {
			fmt.Println("-ip=<ip> fehlt!")
			os.Exit(0)
		}
		c := govt.Client{Apikey: apikey, Url: apiurl}

		// get a file report
		r, err := c.GetIpReport(ip)
		check(err)

		j, err := json.MarshalIndent(r, "", "    ")
		if err != nil {
			fmt.Println("Formatting Error")
			return
		}
		//		currDir, err := os.Getwd()
		if _, err := os.Stat(file); err == nil {
			fmt.Println("File Exists Moving to Reports Directory")
			//		os.Mkdir("report", 0664)
			//		reportDir, err := filepath.Abs("report/")
			//		filepath.Join(currDir, reportDir)
			ioutil.WriteFile(url+"-report", j, 0664)
		} else {
			ioutil.WriteFile("report.txt", j, 0664)
		}
		check(err)

		if mode {
			fmt.Print("Report Generated\n")
			fmt.Println("IP Report:")
		}
	} else {
		fmt.Println("Invalid URL")
		os.Exit(-1)

	}

}

// check - an error checking function

func check(e error) {
	if e != nil {
		os.Exit(-1)
	}
}
