package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
)

const checkInterval = time.Minute * 5

func main() {
	var (
		vpnIP = flag.String("ip", "", "the VPN public IP address, see https://ifconfig.co")
		limit = flag.Duration("limit", time.Minute*30, "time limit before notifications start")
	)
	flag.Parse()
	if *vpnIP == "" {
		fmt.Println("missing -ip flag, you must specify VPN IP, see --help")
		os.Exit(2)
	}
	for {
		time.Sleep(checkInterval)
		if err := checkIP(*limit, *vpnIP); err != nil {
			log.Fatal(err)
		}
	}
}

var (
	connectionTime time.Duration
	lastIP         string
)

func checkIP(timeLimit time.Duration, vpnIP string) error {
	req, err := http.NewRequest("GET", "https://ifconfig.co", nil)
	if err != nil {
		return err
	}
	cl := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := cl.Do(req)
	if err != nil {
		return err
	}
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return err
	}
	resp.Body.Close()

	lastIP = strings.TrimSpace(string(dat))
	if lastIP != vpnIP {
		return nil
	}
	connectionTime += checkInterval
	if connectionTime >= timeLimit {
		alert := fmt.Sprintf("You've been connected to the VPN for over %s", connectionTime)
		log.Printf(alert)
		beeep.Alert("VPN Warning", alert, "")
	}
	return nil
}
