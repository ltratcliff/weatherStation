package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type AmbientWeather []struct {
	Dateutc        int64     `json:"dateutc"`
	Tempinf        float64   `json:"tempinf"`
	Humidityin     int       `json:"humidityin"`
	Baromrelin     float64   `json:"baromrelin"`
	Baromabsin     float64   `json:"baromabsin"`
	Tempf          float64   `json:"tempf"`
	Battout        int       `json:"battout"`
	Humidity       int       `json:"humidity"`
	Winddir        int       `json:"winddir"`
	Windspeedmph   float64   `json:"windspeedmph"`
	Windgustmph    float64   `json:"windgustmph"`
	Maxdailygust   float64   `json:"maxdailygust"`
	Hourlyrainin   float64   `json:"hourlyrainin"`
	Eventrainin    float64   `json:"eventrainin"`
	Dailyrainin    float64   `json:"dailyrainin"`
	Weeklyrainin   float64    `json:"weeklyrainin"`
	Monthlyrainin  float64   `json:"monthlyrainin"`
	Totalrainin    float64   `json:"totalrainin"`
	Solarradiation float64   `json:"solarradiation"`
	Uv             int       `json:"uv"`
	BattCo2        string    `json:"batt_co2"`
	FeelsLike      float64   `json:"feelsLike"`
	DewPoint       float64   `json:"dewPoint"`
	FeelsLikein    float64   `json:"feelsLikein"`
	DewPointin     float64   `json:"dewPointin"`
	Loc            string    `json:"loc"`
	Date           time.Time `json:"date"`
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	sqliteDatabase, _ := sql.Open("sqlite3", "./weather.db")
	for {
		weatherData := fetchWeather()
		updateDB(sqliteDatabase, weatherData)
		time.Sleep(5 * time.Minute)
	}
}

func fetchWeather() AmbientWeather {
	mac := os.Getenv("MAC")
	appKey := os.Getenv("APPKEY")
	apiKey := os.Getenv("APIKEY")
	endDate := ""
	limit := "1"



	params := "apiKey=" + url.QueryEscape(apiKey) + "&" +
		"applicationKey=" + url.QueryEscape(appKey) + "&" +
		"endDate=" + url.QueryEscape(endDate) + "&" +
		"limit=" + url.QueryEscape(limit)

	path := fmt.Sprintf("https://api.ambientweather.net/v1/devices/%s?%s", mac, params)

	log.Println("Fetching data")
	resp, err := http.Get(path)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var weather AmbientWeather
	err = decoder.Decode(&weather)
	if err != nil {
		log.Fatal(err)
	}


	log.Println(weather)
	return weather
	//fmt.Println(weather[0].Tempf)
}

func updateDB(db *sql.DB, weatherData AmbientWeather) {
	log.Println("Inserting record ...")
	insertSQL := `INSERT INTO weather(Dateutc, Tempinf, Humidityin, Baromrelin, Baromabsin, Tempf, Battout, Humidity, 
                 Winddir, Windspeedmph, Windgustmph, Maxdailygust, Hourlyrainin, Eventrainin, Dailyrainin, Weeklyrainin,
                    Monthlyrainin, Totalrainin, Solarradiation, Uv, BattCo2, FeelsLike, DewPoint, FeelsLikein, DewPointin,
                    Loc, Date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertSQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	w := weatherData[0]
	_, err = statement.Exec(w.Dateutc, w.Tempinf, w.Humidityin, w.Baromrelin, w.Baromabsin, w.Tempf, w.Battout, w.Humidity,
		w.Winddir, w.Windspeedmph, w.Windgustmph, w.Maxdailygust, w.Hourlyrainin, w.Eventrainin, w.Dailyrainin, w.Weeklyrainin,
		w.Monthlyrainin, w.Totalrainin, w.Solarradiation, w.Uv, w.BattCo2, w.FeelsLike, w.DewPoint, w.FeelsLikein, w.DewPointin,
		w.Loc, w.Date)
	if err != nil {
		log.Fatalln(err.Error())
	}
}