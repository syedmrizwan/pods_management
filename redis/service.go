package redis

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/syedmrizwan/pods_management/model"
	"log"
)

func InsertVcenterDetailToRedis() {
	conn := GetPool().Get()
	defer conn.Close()

	if err := insertToRedis("vcenters", vcneter); err != nil {
		log.Fatal(err)
	}

	for _, ds := range datastores {
		_, err := conn.Do("RPUSH", fmt.Sprintf("%s.datastore", vcneter.IP), ds)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, cl := range clusters {
		_, err := conn.Do("RPUSH", fmt.Sprintf("%s.cluster", vcneter.IP), cl)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func RoundRobin() {
	conn := GetPool().Get()
	defer conn.Close()

	vcneter := &model.VcenterSummary{}
	getRedisElement("vcenters", vcneter)
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

	_, err = conn.Do("RPUSH", "vcenters", string(b))
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

func RemoveConfigFromRedis(venter string) {
	conn := GetPool().Get()
	defer conn.Close()

	b, _ := json.Marshal(&vcneter)
	_, err := conn.Do("LREM", "vcenters", 0, string(b))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Removed vCenter")

	_, err = conn.Do("DEL", fmt.Sprintf("%s.datastore", vcneter.IP))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted datastore of %s vCenter\n", vcneter.IP)

	_, err = conn.Do("DEL", fmt.Sprintf("%s.cluster", vcneter.IP))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted cluster of %s vCenter\n", vcneter.IP)

}
