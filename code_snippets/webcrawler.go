package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/williballenthin/govt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
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
func main() {
	fmt.Println("VT URl Scanner")
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // starts a go-routine
	}
	for range os.Args[1:] {

		fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	}
}

// check - an error checking function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	fmt.Println("At Begining of Fetch:\n")
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()

	fmt.Println("%.2fs %7d %s", secs, nbytes, url)
	fmt.Println("Analyzing URL(s):\n")
	ip_addr, err := net.LookupIP(url)
	if err != nil {
		fmt.Sprintf("ip lookup failed %s %v", url, err)
	}
	flag.StringVar(&ip, "ip", ip_addr[0].String(), "ip sum of a file to as VT about.")

	fmt.Println("Sending to VirusTotal:\n")
	flag.Parse()
	if ip == "" {
		fmt.Println("-ip=<ip> fehlt!")
		os.Exit(1)
	}
	c := govt.Client{Apikey: apikey, Url: apiurl}

	// get a file report
	r, err := c.GetIpReport(ip)
	fmt.Println("Report Generated\n")
	check(err)
	j, err := json.MarshalIndent(r, "", "    ")
	fmt.Printf("IP Report: ")
	os.Stdout.Write(j)
}
