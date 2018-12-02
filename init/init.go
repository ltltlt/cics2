package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"

	"cytus2.rocks/cics2/dao/redis"
	"cytus2.rocks/cics2/models/user"
)

func main() {
	filename := os.Getenv("CICS_USERS")
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(f)
	client := redis.MustGetClient("user")
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		u := user.User{
			ID:   record[0],
			Name: record[1],
			Pass: record[2],
		}
		us, err := json.Marshal(u)
		if err != nil {
			panic(err)
		}
		redis.MustCmdSuccess(client.Set(u.ID, us, 0))
	}
}
