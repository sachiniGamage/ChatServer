package cassandra

import (
	"time"

	"github.com/gocql/gocql"
)

func DBConnection() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Consistency = gocql.Any
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10

}
