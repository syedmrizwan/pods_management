package redis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/syedmrizwan/pods_management/model"
)

var vcneter = model.VcenterSummary{
	IP:       "10.0.0.11",
	Username: "rizwan",
	Password: "rizwan",
}

var datastores = []string{
	"SpringpathDS-WZP23201B14",
	"SpringpathDS-WZP23201B1Y",
	"SpringpathDS-WZP23201B2U",
	"NN-Datastore-01",
	"SpringpathDS-WZP23090WYM",
	"SpringpathDS-WZP23130UA4",
	"SpringpathDS-WZP23090WXQ",
	"Datastore3",
	"SpringpathDS-WZP232011C7",
	"SpringpathDS-WZP2329045K",
	"datastore1 (2)",
	"SpringpathDS-WZP23090WY1",
	"Datastore2",
	"SpringpathDS-WZP23090WYJ",
	"Datastore",
	"datastore1 (3)",
	"datastore1",
	"datastore1 (1)",
	"SpringpathDS-WZP23201B1Z",
	"c220-07-Disk-1",
	"SpringpathDS-WZP232903X0",
	"c220-06-Disk-1",
}

var clusters = []string{
	"OLD LAB HOST",
	"Hyperflex-c240-1",
}

func insertToRedis(listName string, objectToPush interface{}) error {
	conn := GetPool().Get()
	defer conn.Close()

	b, err := json.Marshal(&objectToPush)
	if err != nil {
		return err
	}

	if _, err = conn.Do("RPUSH", listName, string(b)); err != nil {
		return err
	}

	return nil

}

func getRedisElement(listName string, objectToPop interface{}) error {
	conn := GetPool().Get()
	defer conn.Close()

	reply, _ := redis.ByteSlices(conn.Do("BLPOP", listName, 5))

	if err := json.Unmarshal(reply[1], objectToPop); err != nil {
		return err
	}
	return nil
}
