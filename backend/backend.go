package backend

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails"
)

// Local constants
//const brokerURI  string = "tcp://localhost:1883"
//const clientId string = "xxxx"
//const user     string = "xxxx"
//const password string = "xxxx"
const topic string = "sensorv02"

type SensorData struct {
	SensorName  string  `json:"sensor_name"`
	SensorValue float32 `json:"sensor_value"`
}

type Sensor_Data struct {
	SensorName  string
	SensorValue float32
}

type Sensor_DataArray []Sensor_Data

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()

	// AddBroker adds a broker URI to the list of brokers to be used.
	// The format should be  "scheme://host:port"
	opts.AddBroker(brokerURI)

	//	opts.SetUsername(user)
	//	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func connect(brokerURI string, clientId string) mqtt.Client {
	fmt.Println("Trying to connect (" + brokerURI + ", " + clientId + ")...")
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()

	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}

	return client
}

func listen(topic string, c chan Sensor_DataArray) {
	//client := connect("sub", uri)
	client := connect("tcp://localhost:1883", "go-sub-client")
	fmt.Println("Subrcribe on topic '" + topic + ")...")
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {

		var sdata SensorData
		json.Unmarshal([]byte(msg.Payload()), &sdata)
		var sname = sdata.SensorName
		var svalue = sdata.SensorValue
		//fmt.Println(sname, svalue)
		// m := make(map[string]float32)
		// m[sname] = svalue
		wCounts := Sensor_DataArray{
			Sensor_Data{sname, svalue},
		}
		//wCounts := sdata
		c <- wCounts

	})
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func sensorDataProcess() *Sensor_readings {
	// Create Database
	database, _ := sql.Open("sqlite3", "./sensordata.db")
	statement, _ := database.Prepare(" CREATE TABLE IF NOT EXISTS sensor_data (id INTEGER PRIMARY KEY, SensorName TEXT, SensorValue INTEGER, Time REAL)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO sensor_data (SensorName, SensorValue, Time) VALUES (?, ?, ?)")

	done := make(chan bool, 1)
	// don := make(chan bool, 1)
	c := make(chan Sensor_DataArray, 10)
	//c2 := make(chan Sensor_DataArray, 60)
	//go connectbroker()
	time.Sleep(2 * time.Second)

	go listen(topic, c)
	//go minuteaggreagate(c2)

	w, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	freq := make(map[string]float32)
	count := make(map[string]int)
	timer := 1
	freq2 := make(map[string]float32)
	count2 := make(map[string]int)

	for {
		val, ok := <-c
		//c2 <- val

		if !ok {
			fmt.Println(val)
			break
		} else {
			//go minuteaggreagate(c2)
			for _, wCount := range val {
				freq[wCount.SensorName] += wCount.SensorValue
				count[wCount.SensorName] += 1

				fmt.Println(wCount.SensorName, ":", wCount.SensorValue)
				fmt.Println(wCount.SensorName, "Total Consumed :", freq[wCount.SensorName])
				fmt.Println(wCount.SensorName, "count :", count[wCount.SensorName])
				fmt.Println(wCount.SensorName, "count :", freq[wCount.SensorName]/float32(count[wCount.SensorName]))
				fmt.Println()

				// ====Emit

				//// minute counter
				if count2[wCount.SensorName] <= 29 {
					// sum value
					freq2[wCount.SensorName] += wCount.SensorValue
					// add timer
					timer += 1
					count2[wCount.SensorName] += 1

				} else {
					// print(save) output
					fmt.Println("======= 5 Seconds Accumulator ===========")
					fmt.Println("Sensor Name : ", wCount.SensorName)
					fmt.Println("5 Secs Total: ", freq2[wCount.SensorName])
					fmt.Println("5 Secs Time: ", count2[wCount.SensorName])
					fmt.Println("=========================================")

					// Save

					_, err = fmt.Fprintf(w, "======================================================= \n")
					check(err)
					_, err = fmt.Fprintf(w, "Sensor Name: %v\n", wCount.SensorName)
					check(err)
					_, err = fmt.Fprintf(w, "Total Consumed in 10s %v\n", freq2[wCount.SensorName])
					check(err)
					_, err = fmt.Fprintf(w, "Count in last 10s %v\n", count2[wCount.SensorName])
					check(err)
					_, err = fmt.Fprintf(w, "======================================================\n")
					check(err)

					now := time.Now()
					tm := now.Unix()
					statement.Exec(wCount.SensorName, freq2[wCount.SensorName], tm)

					// reset timer and  accumulator
					timer = 0
					freq2[wCount.SensorName] = 0
					count2[wCount.SensorName] = 0
					// restart timer & accumulator
					timer = 1
					count2[wCount.SensorName] = 1
					freq2[wCount.SensorName] += wCount.SensorValue
					// reset values

				}
				return &Sensor_readings{
					// 			Value: svalue.SensorValue,
					SensorName:   wCount.SensorName,
					SensorValue:  wCount.SensorValue,
					TotalValue:   freq[wCount.SensorName],
					AverageValue: freq[wCount.SensorName] / float32(count[wCount.SensorName]),
				}
			}
		}
	}

	<-done
	return nil
}

// =========== Connect with Wails =====================

// Stats .
type Stats struct {
	log *wails.CustomLogger
}

// Sensor readings .
type Sensor_readings struct {
	//Value int16 `json:"value"`
	SensorName   string
	SensorValue  float32
	TotalValue   float32
	AverageValue float32
}

// WailsInit .
func (s *Stats) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Stats")

	go func() {
		for {

			//runtime.Events.Emit("Evernt ")
			runtime.Events.Emit("sensor_reading", sensorDataProcess())
			time.Sleep(1 * time.Second)
		}
	}()

	return nil
}
