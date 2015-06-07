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
	routerIP := "192.168.1.1" //static for now

	log.Println("Get password")

	fmt.Print("Enter router password: ")
	password := base64.StdEncoding.EncodeToString(gopass.GetPasswd())

	log.Println("Started auth..")

	_, err := http.PostForm("http://"+routerIP+"/login.cgi", url.Values{
		"page":       {""},
		"logout":     {""},
		"action":     {"submit"},
		"itsbutton1": {"Submit"},
		"h_language": {"en"},
		"pws":        {password},
	})

	if err != nil {
		fmt.Println("Happend fucking error: ", err)
		return
	}

	log.Println("Now trying to restart...")
	restartPage, err := http.PostForm("http://"+routerIP+"/ut_reset.cgi", url.Values{
		"page":    {""},
		"action":  {"Reboot"},
		"logout":  {""},
		"webpage": {"ut_reset.html"},
		"reboot":  {"Restart Router"},
	})

	if err != nil {
		fmt.Println("Happend error.. %s", err)
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
