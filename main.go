package main

import (
	"encoding/base64"
	"fmt"
	"github.com/howeyc/gopass"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

func main() {
	var password string

	log.Println("Get password")

	fmt.Print("Enter router password: ")
	password = base64.StdEncoding.EncodeToString(gopass.GetPasswd())

	log.Println("Started auth..")

	_, err := http.PostForm("http://192.168.1.1/login.cgi", url.Values{
		"page":       {""},
		"logout":     {""},
		"action":     {"submit"},
		"itsbutton1": {"Submit"},
		"h_language": {"en"},
		"pws":        {password},
	})

	if err != nil {
		fmt.Errorf("Happend fucking error %s", err.Error())
		return
	}

	log.Println("Now trying to restart...")
	restartPage, err := http.PostForm("http://192.168.1.1/ut_reset.cgi", url.Values{
		"page":    {""},
		"action":  {"Reboot"},
		"logout":  {""},
		"webpage": {"ut_reset.html"},
		"reboot":  {"Restart Router"},
	})

	if err != nil {
		fmt.Errorf("Happend error.. %s", err.Error())
		return
	}

	body, err := ioutil.ReadAll(restartPage.Body)

	re := regexp.MustCompile("type=\"password\".*name=\"pws\"")
	if re.FindIndex(body) != nil {
		log.Println("Password was incorrect")
		return
	}

	log.Println("restarted...")

}
