package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"

	"github.com/williballenthin/govt"
)

// KEA -Go uses camelCase instead of spacing_like_this

/*
  KEA - A good design would be to populate a []string of URL's in your main function. Then just loop and call fetch. It would clean up your "if cli_str/cli_file" logic, and allow you to get rid of the if/else. You could simply append to the []string if the flag is given. Nice, clean and concise.

  Pseudocode:

  urls := make([]string, 0)

  concurrent := flag.Bool("j", false, "Concurrently fetch url reports")

  if *cli_str != "" {
      // Do whatever validity checks you want
      urls = append(urls, *cli_str)
  }

  if *cli_file != "" {
      // Open file
      // For line in file - urls = append(urls, *cli_str)
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

/*
KEA - Init should only be used for initializing packages (init is called on import). It should not be used for your main functions. (the reason is mainly idiomatic, but also semantic)
*/
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

	/*
		KEA - In Go, prefer Slices to Arrays. Do this instead

		urls := make([]string, 0)

		This creates a slice of strings of length 0 and capacity 0. Later in your code instead of doing url_array[i] = <some string>, do urls = append(urls, <some string>)

		The append function will handle growing and memcopying the slice for you as needed. You'd normally want to do urls := make([]string, 0, <how many you know you'll need>) to avoid the alloc and memcopy, but in this case you don't know so this is ok.
	*/
	var url_array [100]string

	/*
		KEA - It's idiotmatic to leave the var out and do
		cli_str := flag.String("u", "", "Individual URL to be scanned")
	*/
	var cli_str = flag.String("u", "", "Individual URL to be scanned")
	var cli_verb = flag.String("m", "v,s", "Verbose Mode (v) or Silent Mode (s)")
	var cli_file = flag.String("i", "urls.txt", "File of URLs to be scanned")
	var ii int
	flag.Parse()

	fmt.Println("VT URl Scanner")

	/*
		KEA - The check on the length is not neccessary. If no flag is provided, it is set to the default you gave the flag call , which is the empty string.
	*/
	if *cli_str != "" && len(*cli_str) > 0 {
		scan_url = *cli_str

		/*
			KEA - Better here to define two separate flags instead of combining it under cli_verb.

			silent := flag.Bool("s", false, "Enable silent mode")
			verbose := flag.Bool("v", false, "Enable verbose mode")
		*/
		if strings.EqualFold(*cli_verb, "v") {
			fmt.Println("Verbose Mode Selected")
			verbose = true
			fetch(scan_url, verbose)
		} else if strings.EqualFold(*cli_verb, "s") {
			fmt.Println("Silent Mode Running")
			silent = true
			fetch(scan_url, silent)
		}
	} else if *cli_file != "" && len(*cli_file) > 0 { // KEA: Same here of course with the length check

		file, err := os.Open("urls.txt")
		if err != nil {
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file) // KEA: Love it, nice use of the fact that os.File is an io.Reader. Well done.
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

	// KEA - Careful, you're swalloing the json.MarshalIndent error here. Always always always error check
	j, err := json.MarshalIndent(r, "", "    ")

	/*
		KEA

		Idiomatic Go prefers handling errors in the indented blocks and keeping the normal code flow as unindented as possible. Really helps readability and flow.

		rc, err := SomeFunc()
		if err != nil {
		  // Error Handle
		}

		// Normal Code Path
	*/
	if _, err := os.Stat(file); err == nil {
		fmt.Println("File Exists")
		ioutil.WriteFile("test.txt", j, 0664)
	}
	ioutil.WriteFile("report.txt", j, 0664) // KEA - Don't swallow errors.  err := ioutil.WriteFile
	check(err)                              // KEA - This is the json.MarshalIndent err currently

	if mode {
		fmt.Print("Report Generated\n")
		fmt.Println("IP Report:")
	}
}

// KEA - You'll instead be able to pass a slice here. This is massively more efficient as when you pass an array go must copy the array on the runtime stack for the function call. With a slice it simply puts the reference on the stack.
func multi_fetch(url [100]string, mode_str string, size int) {
	fmt.Println("inside of multi_fetch")
	var aa int
	var mode bool
	mode = false

	// KEA - Check this explicitly with bool vars, string comparison is error prone and easy to break.
	if strings.EqualFold(mode_str, "verbose") {
		mode = true
	}
	for aa = 0; aa < size; aa++ {
		if url[aa] != "" {
			// KEA - Aha, the problem you are having with the goroutines is that each one is going to try to do something with "report.txt". If you make it so that the filename is different every time somehow (perhaps a random number?) then your concurrent code will be ok.
			go fetch(url[aa], mode)
		}
	}

}

// check - an error checking function

/*
KEA - Never panic unless there's a really really good reason to. Better to log and then os.Exit, or log.Fatal. Go isn't an exception driven language.
*/
func check(e error) {
	if e != nil {
		panic(e)
	}
}
