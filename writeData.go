package main

import (
	"fmt"
	"os/exec"
	"time"
	"math/rand"
	"strconv"
	"github.com/gocql/gocql"
)

/*
CREATE TABLE car_stats (
     sensor_id int,
     collect_time timestamp,
     temperature text,
     speed text,
     PRIMARY KEY ((sensor_id), collect_time)
   ) WITH CLUSTERING ORDER BY (collect_time DESC);
*/

var Session *gocql.Session

type Car_stats struct {
	sensor_id int
	Collect_time time.Time
	temperature string
	speed string
}

func main() {
	cassandraInit("127.0.0.1")
	simulateWrites()
}

func cassandraInit(CONNECT string){
	var err error
	cluster := gocql.NewCluster(CONNECT) //connect to cassandra database
	cluster.Keyspace = "test_keyspace"
	Session, err = cluster.CreateSession() 
	if err != nil {
		fmt.Print(err)
	}
}

func simulateWrites() {
	sensors := []int{10101, 10102, 10103}
	for {
		for _, sensor := range sensors {
			randI := rand.Intn(8 - 3) + 8 //generates a value within [3, 8]
			str := strconv.Itoa(randI*11) //returns string of randomly generated int
			data := Car_stats{sensor, time.Now(), (str+"C"), (str+"km")}
			cassandraWrite(data)
		}
		time.Sleep(5 * time.Second)
	}
}

func cassandraWrite(data Car_stats) {
	//create new row in test_table
	if err := Session.Query("INSERT INTO car_stats(sensor_id,collect_time,temperature,speed) VALUES(?, ?, ?, ?)", data.sensor_id, data.Collect_time, data.temperature, data.speed).Exec(); err != nil {
		fmt.Println(err)
	}
}

func cassandraRead(sensor_id int, time_lower time.Time, time_upper time.Time) []Car_stats {
	var data []Car_stats
	m := map[string]interface{}{}
	query := fmt.Sprintf("SELECT * FROM car_stats WHERE sensor_id=%v AND collect_time>='%s' AND collect_time<='%s'", sensor_id, time_lower.Format("2006-01-02 15:04:05.000"), time_upper.Format("2006-01-02 15:04:05.000"))
	
	//read from specifed range in car_stats
	iterable := Session.Query(query).Iter()
	for iterable.MapScan(m) {
		data = append(data, Car_stats{
			sensor_id: m["sensor_id"].(int),
			Collect_time: m["collect_time"].(time.Time),
			temperature: m["temperature"].(string),
			speed: m["speed"].(string),
		})
		m = map[string]interface{}{}
	}
	return data
}

func gethWrite(connect string, msg string){
	tx := fmt.Sprintf("eth.sendTransaction({from:eth.accounts[0],to:eth.accounts[0],value:1,data:web3.toHex('%v')})", msg)
	output, err := exec.Command("geth", "attach", connect, "--exec", tx).CombinedOutput() 

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
}