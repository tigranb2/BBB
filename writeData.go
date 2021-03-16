package main

import (
	"fmt"
	"os/exec"
	"github.com/gocql/gocql"
)

var Session *gocql.Session

type Data struct {
	id int
	message string
}

func main() {
	msg1 := Data{1, "Hello World"}
	msg2 := Data{2, "Test message"}
	msg3 := Data{3, "message 3 !"}
	cassandraInit("127.0.0.1")
	cassandraWrite(msg1)
	cassandraWrite(msg2)
	cassandraWrite(msg3)
	fmt.Println(cassandraRead())
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

func cassandraWrite(data Data) {
	//create new row in test_table
	if err := Session.Query("INSERT INTO test_table(id, message) VALUES(?, ?)", data.id, data.message).Exec(); err != nil {
		fmt.Println(err)
	}
}

func cassandraRead() []Data {
	var data []Data
	m := map[string]interface{}{}

	//read all rows in test_table
	iterable := Session.Query("SELECT * FROM test_table").Iter()
	for iterable.MapScan(m) {
		data = append(data, Data{
			id: m["id"].(int),
			message: m["message"].(string),
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