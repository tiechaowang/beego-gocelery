package main

import (
	_ "firstproject/routers"
	"github.com/astaxie/beego"
	 "firstproject/models"
	 "github.com/astaxie/beego/logs"
	 //"github.com/astaxie/beego/toolbox"
	 	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
	"time"
	"fmt"
)


func tadd() error  {
    user := models.User{Name: "wtcc"}
    models.AddUser(&user)
    return nil
}

func add(a, b int) int {
	fmt.Println("eeeeeee")
	time.Sleep(10 * time.Second)

    user := models.User{Name: "wtcc"}
    models.AddUser(&user)

	fmt.Println("dddddddd")
	return a + b
}

func main() {
	redisPool := &redis.Pool{
		MaxIdle:     3,                 // maximum number of idle connections in the pool
		MaxActive:   0,                 // maximum number of connections allocated by the pool at a given time
		IdleTimeout: 240 * time.Second, // close connections after remaining idle for this duration
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL("redis://:test@172.17.32.218:6378/0")
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	// initialize celery client
	cli, _ := gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		&gocelery.RedisCeleryBackend{Pool: redisPool},
		5, // number of workers
	)

    


	// register task
	cli.Register("worker.add", add)

	// start workers (non-blocking call)
	cli.StartWorker()

	// tk1 := toolbox.NewTask("mytask", "*/1 * * * * *", tadd)
	// toolbox.AddTask("myTask", tk1)
 //    toolbox.StartTask()
	//log := logs.NewLogger()
	//beego.SetLogger("file", `{"filename":"/opt/ops/src/test.log"}`)
	logs.SetLogger("file", `{"filename":"/opt/ops/src/test.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info("start")
	beego.Run()
}

