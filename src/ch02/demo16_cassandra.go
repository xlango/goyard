package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("192.168.10.28", "192.168.10.29", "192.168.10.31")
	cluster.Keyspace = "gotest"
	cluster.Consistency = gocql.Quorum
	//设置连接池的数量,默认是2个（针对每一个host,都建立起NumConns个连接）
	cluster.NumConns = 3

	session, _ := cluster.CreateSession()
	time.Sleep(1 * time.Second) //Sleep so the fillPool can complete.

	defer session.Close()

	//unlogged batch, 进行批量插入，最好是partition key 一致的情况
	batch := session.NewBatch(gocql.UnloggedBatch)

	batch.Query(`INSERT INTO user(id,name) VALUES (?,?)`, 4, "aa")

	if err := session.ExecuteBatch(batch); err != nil {
		fmt.Println("execute batch:", err)
	}

	var id int
	var name string

	if err := session.Query(`SELECT * FROM user WHERE id = ? LIMIT 1`,
		1).Consistency(gocql.One).Scan(&id, &name); err != nil {
		log.Fatal(err)
	}
	fmt.Println("user:", id, name)

	iter := session.Query(`SELECT * FROM user`).Iter()

	for iter.Scan(&id, &name) {
		fmt.Println(id, name)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

}
