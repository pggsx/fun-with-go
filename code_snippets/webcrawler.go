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
	for _, url := range os.Args[1:] {
		fetch(url) // fetches url and scans it
	}
}

func fetch(url string) {
	fmt.Println("url %s", url)
	fmt.Println("At Begining of Fetch:\n")
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	fmt.Println("%d %s", nbytes, url)
	fmt.Println("Analyzing URL(s):\n")
	ip_addr, err := net.LookupIP(url)
	fmt.Println("Finished Domain Lookup\n")
	if err != nil {
		fmt.Sprintf("ip lookup failed %s %v", url, err)
	}
	for i := 0; i < len(ip_addr); i++ {
		fmt.Println("ip_addr", ip_addr[i])
		flag.StringVar(&ip, "ip", ip_addr[i].String(), "ip sum of a file to as VT about.")
	}
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

// check - an error checking function
func check(e error) {
	if e != nil {
		panic(e)
	}
}
