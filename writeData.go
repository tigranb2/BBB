package main

import (
	"fmt"
	"os/exec"
	"github.com/gocql/gocql"
)

var Session *gocql.Session

type testData struct {
	id int
	msg string
}
func main() {
	msg1 := testData{1, "Hello World"}
	cassandraInit("127.0.0.1")
	cassandraWrite(msg1)
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

func cassandraWrite(data testData) {
	//create new row in test_table
	if err := Session.Query("INSERT INTO test_table(id, message) VALUES(?, ?)", data.id, data.msg).Exec(); err != nil {
		fmt.Println(err)
	}
}

//func cassandraRead() {}

func gethWrite(connect string, msg string){
	tx := fmt.Sprintf("eth.sendTransaction({from:eth.accounts[0],to:eth.accounts[0],value:1,data:web3.toHex('%v')})", msg)
	output, err := exec.Command("geth", "attach", connect, "--exec", tx).CombinedOutput() 

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
}