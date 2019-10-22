package main

import (
	"fmt"
	"time"

	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client/v2"
)

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
		return
	}
	defer c.Close()
	//ping
	fmt.Println(c.Ping(time.Second))

	//创建数据库
	q := client.NewQuery("CREATE DATABASE db", "", "")
	r, err := c.Query(q)
	if err != nil {
		fmt.Println("Error creating InfluxDB Database: ", err.Error())
		return
	}
	fmt.Println(r)
	//write
	bps, err := client.NewBatchPoints(client.BatchPointsConfig{
		Precision:        "s",  //时间精度，纳秒
		Database:         "db", //数据库名称
		RetentionPolicy:  "",   //保留策略
		WriteConsistency: "",   //写入一致性，集群的时候
	})
	//influxdb的数据是point 里面包含time+fields+tags
	tags := map[string]string{"sex": "1"} //类似于mysql的索引

	fields := map[string]interface{}{
		"name": "alpha",
	} //类似于mysql的值
	p, _ := client.NewPoint("user", tags, fields, time.Now())
	bps.AddPoint(p)
	err = c.Write(bps)
	if err != nil {
		fmt.Println("Error creating InfluxDB Database: ", err.Error())
		return
	}
	fmt.Println("创建成功")
	//query
	q1 := client.NewQuery(`SELECT * FROM "users"`, "db", "")
	response, err := c.Query(q1)
	if err == nil {
		fmt.Println(response)
	} else {
		fmt.Println(err)
	}

}
