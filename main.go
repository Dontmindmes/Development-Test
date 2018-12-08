package main

/*
	-
	- Gradually increase volume to set amount / DONE
	- Play Quran on fridays
	-

*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/evalphobia/google-home-client-go/googlehome"
)

type Athan struct {
	Data YtData `json:"data"`
}

type YtData struct {
	Timings YtTime `json:"timings"`
}

type YtTime struct {
	F string `json:"Fajr"`
	S string `json:"Dhuhr"`
	A string `json:"Asr"`
	M string `json:"Maghrib"`
	I string `json:"Isha"`
}

type Config struct {
	Settings struct {
		IP       string  `json:"IP"`
		Language string  `json:"Language"`
		Accent   string  `json:"Accent"`
		City     string  `json:"City"`
		Country  string  `json:"Country"`
		Athan    string  `json:"Athan"`
		Volume   float64 `json:"Volume"`
	}
}

var Y Athan

var (
	API1 string = "http://api.aladhan.com/v1/timingsByCity?city="
	API2 string = "&country="
	API3 string = "&method=8"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fmt.Fprintln(w, "Google Home, Providence")
}
func main() {
	http.HandleFunc("/", indexHandler)

	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal("Error importing config.json file", err)
	}

	var Api_Connect = API1 + config.Settings.City + API2 + config.Settings.Country + API3

	ret := fmt.Sprintf(Api_Connect)

	cli, err := googlehome.NewClientWithConfig(googlehome.Config{
		Hostname: config.Settings.IP,
		Lang:     config.Settings.Language,
		Accent:   config.Settings.Accent,
	})
	if err != nil {
		panic(err)
	} else {
		cli.SetVolume(config.Settings.Volume)
		cli.Notify("Successfully Connected.")
	}

	fmt.Println(ret)
	resp, err := http.Get(ret)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &Y)
	if err != nil {
		log.Fatal(err)
	}

	for range time.Tick(time.Second * 15) {
		Ts := time.Now()
		Atimes := Ts.Format("03:04:05PM")

		t, _ := time.Parse("03:04:05PM", Atimes)
		test := t.Format("15:04")

		fmt.Println("Current Isha time", Y.Data.Timings.I)
		fmt.Println(test)

		//cli.Play(config.Settings.Athan)
		time.Sleep(15 * time.Second)
		//T := time.Now()
		//Atime := T.Format("15:04")

		if Y.Data.Timings.F == test {
			fmt.Println("Its time for Fajr")
			for i := 0.00; i < 1.00; i = +0.01 {
				cli.SetVolume(i)
				time.Sleep(10 * time.Second)
			}
			cli.Play(config.Settings.Athan)
			time.Sleep(4 * time.Minute)
			continue
		} else if Y.Data.Timings.S == test {
			fmt.Println("Its time for Dzuhur")
			for i := 0.00; i < 1.00; i = +0.01 {
				cli.SetVolume(i)
				time.Sleep(10 * time.Second)
			}
			cli.Play(config.Settings.Athan)
			time.Sleep(4 * time.Minute)
			continue
		} else if Y.Data.Timings.A == test {
			fmt.Println("Its time for Aseir")
			for i := 0.00; i < 1.00; i = +0.01 {
				cli.SetVolume(i)
				time.Sleep(10 * time.Second)
			}
			cli.Play(config.Settings.Athan)
			time.Sleep(4 * time.Minute)
			continue
		} else if Y.Data.Timings.M == test {
			fmt.Println("Its time for Maghrib")
			for i := 0.00; i < 1.00; i = +0.01 {
				cli.SetVolume(i)
				time.Sleep(10 * time.Second)
			}
			cli.Play(config.Settings.Athan)
			time.Sleep(4 * time.Minute)
			continue
		} else if Y.Data.Timings.I == test {
			fmt.Println("Its time for Isha")
			for i := 0.00; i < 1.00; i = +0.01 {
				cli.SetVolume(i)
				time.Sleep(10 * time.Second)
			}
			cli.Play(config.Settings.Athan)
			time.Sleep(4 * time.Minute)
			continue
		}

	}

	for {

	}

}

//LoadConfig file
func LoadConfig(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)

	defer configFile.Close()
	if err != nil {
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}
