package redis

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

type Vcenter struct {
	IP       string
	Username string
	Password string
}

func InsertVcenterDetailToRedis() {
	conn := GetPool().Get()
	defer conn.Close()

	vcneter := Vcenter{
		IP:       "10.0.0.11",
		Username: "rizwan",
		Password: "rizwan",
	}

	if err := insertToRedis("vcenters12", vcneter); err != nil {
		log.Fatal(err)
	}

	datastores := []string{
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

	for _, ds := range datastores {
		_, err := conn.Do("RPUSH", fmt.Sprintf("%s.datastore", vcneter.IP), ds)
		if err != nil {
			log.Fatal(err)
		}
	}

	clusters := []string{
		"OLD LAB HOST",
		"Hyperflex-c240-1",
	}

	for _, cl := range clusters {
		_, err := conn.Do("RPUSH", fmt.Sprintf("%s.cluster", vcneter.IP), cl)
		if err != nil {
			log.Fatal(err)
		}
	}
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

func RoundRobin() {
	conn := GetPool().Get()
	defer conn.Close()

	vcneter := &Vcenter{}
	getRedisElement("vcenters12", vcneter)
	//reply, _ := redis.ByteSlices(conn.Do("BLPOP", "vcenters11", 5))
	//
	//vcneter := &Vcenter{}
	//err := json.Unmarshal(reply[1], vcneter)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	fmt.Println(vcneter)

	dsReply, _ := redis.ByteSlices(conn.Do("BLPOP", fmt.Sprintf("%s.datastore", vcneter.IP), 5))
	ds := string(dsReply[1])
	fmt.Println(ds)

	clReply, _ := redis.ByteSlices(conn.Do("BLPOP", fmt.Sprintf("%s.cluster", vcneter.IP), 5))
	cl := string(clReply[1])
	fmt.Println(cl)

	//insert back
	b, err := json.Marshal(&vcneter)
	if err != nil {
		return
	}

	_, err = conn.Do("RPUSH", "vcenters12", string(b))
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Do("RPUSH", fmt.Sprintf("%s.datastore", vcneter.IP), ds)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Do("RPUSH", fmt.Sprintf("%s.cluster", vcneter.IP), cl)
	if err != nil {
		log.Fatal(err)
	}
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
