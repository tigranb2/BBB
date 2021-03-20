package main

import (
	"fmt"
	"os/exec"
	"time"
	"github.com/gocql/gocql"
)

/*
CREATE TABLE car_stats (
	sensor_id int,
	collect_time timestamp,
	temperature text,
	speed text,
	PRIMARY KEY (sensor_id,collect_time)
)
*/

var Session *gocql.Session

type Car_stats struct {
	sensor_id int
	collect_time time.Time
	temperature string
	speed string
}

func main() {
	cassandraInit("127.0.0.1")
	msg1 := Car_stats{10101, time.Now(), "32C", "61km"}
	time.Sleep(1000 * time.Millisecond)
	msg2 := Car_stats{10132, time.Now(), "31C", "60km"}
	time.Sleep(1000 * time.Millisecond)
	msg3 := Car_stats{10101, time.Now(), "24C", "37km"}
	time.Sleep(1000 * time.Millisecond)
	msg4 := Car_stats{10101, time.Now(), "19C", "25km"}
	cassandraWrite(msg1)
	cassandraWrite(msg2)
	cassandraWrite(msg3)
	cassandraWrite(msg4)

	fmt.Println(cassandraRead(10101, time.Now().Add(-4 * time.Minute), time.Now()))
}

func cassandraInit(CONNECT string){
	var err error
	cluster := gocql.NewCluster("127.0.0.1") //connect to cassandra database
	cluster.Keyspace = "test_keyspace"
	Session, err = cluster.CreateSession() 
	if err != nil {
		fmt.Print(err)
	}
}

func cassandraWrite(data Car_stats) {
	//create new row in test_table
	if err := Session.Query("INSERT INTO car_stats(sensor_id,collect_time,temperature,speed) VALUES(?, ?, ?, ?)", data.sensor_id, data.collect_time, data.temperature, data.speed).Exec(); err != nil {
		fmt.Println(err)
	}
}

func cassandraRead(sensor_id int, time_lower time.Time, time_upper time.Time) []Car_stats {
	var data []Car_stats
	m := map[string]interface{}{}

	//read from specifed range in car_stats
	iterable := Session.Query("SELECT * FROM car_stats WHERE sensor_id=? AND collect_time>='?' AND collect_time<='?'", sensor_id, time_lower.Format("2006-01-02T15:04:05.000+0000"), time_upper.Format("2006-01-02T15:04:05.000+0000")).Iter()
	for iterable.MapScan(m) {
		data = append(data, Car_stats{
			sensor_id: m["sensor_id"].(int),
			collect_time: m["colect_time"].(time.Time),
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