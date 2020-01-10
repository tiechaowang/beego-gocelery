package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	 "firstproject/models"
	 //"strconv"
	 "reflect"
	 //"unsafe"
	 "encoding/json"
	 //"net/http"
	 //"github.com/astaxie/beego/toolbox"
	 "time"
	 "github.com/gocelery/gocelery"
	 "github.com/gomodule/redigo/redis"
)

type TestController struct {
	beego.Controller
}



func (c *TestController) Get() {
	beego.Info("aaaaaaaaaa")
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
		1,
	)

	// prepare arguments
	taskName := "worker.add"
	var argA int = 6
	var argB int= 7

	// run task
	cli.Delay(taskName, argA, argB)


	//tk1 := toolbox.NewTask("mytask", "* * * * * *", tadd)
    //go tk1.Run()

	c.SetSession("tests", string("444"))
	id := c.GetString("id")
	beego.Info(id)
	o := orm.NewOrm()
	o.Using("default")

	var list []orm.Params
	user := new(models.User)
	ids := []int{}
	ids  = append(ids, 1)
	ids  = append(ids, 2)
	o.QueryTable(user).Filter("id__in", ids).Values(&list)
	mystruct := map[string]interface{}{"success": 0, "message": &list}
	mystruct["success"] = 1
	mystruct["su"] = 1
	mystruct["kk"] =  c.XSRFToken()

    var uids []models.User
	for key, value := range list {
		beego.Info(key)
		beego.Info(value["Id"])
		beego.Info(reflect.TypeOf(value["Id"]))
		var y int
		if d,ok:=(value["Id"].(int64));ok{        // 转换回原始类型 
	       y = int(d)   //*(*int)(unsafe.Pointer(&d))
	    } 
		u := models.User{Id: y, Name: "y"}
        
		uids  = append(uids, u)
		//strconv.Itoa(32)
		// v, ok := value["Id"].(int)
		// beego.Info(v)
		// if ok {
		// 	beego.Info(v)
  //           uids = append(uids, v)
  //       } else{
  //       	uids = append(uids, v)
  //       }
        
    }
    mystruct["uid"] = &uids
    
	// qs = qs.Filter("name", "slene").Limit(10)
 //    for w range qs:
 //        beego.Info(w.id)
	// user := models.User{Name: "jojo"}
	// models.AddUser(&user)
	// var u models.User
	// u.Name = "slene"
	// id, err := o.Insert(&u)
 //    c.Data["d"] = id
 //    c.Data["dd"] = err

	// u := models.User{Name:"dd"}
	// o.Insert(u)
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "test.html"
	c.Data["json"] = &mystruct
	//c.Data["json"] = &list
    c.ServeJSON()
}

type da struct {
    Id          string
    Name        string
}


func (c *TestController) Post() {
    beego.Info("bbbbbbbbb")
    testsession := c.GetSession("tests")
    beego.Info(testsession)
    //iii := c.GetString("id")
    //beego.Info(iii)
    rqu := c.Ctx.Request.Header["Username"]
    beego.Info(rqu[0])
    ii := c.Ctx.Input.RequestBody
    var dadata da
    err := json.Unmarshal(ii, &dadata)
	if err != nil {
		beego.Info(err.Error())
    }

    beego.Info(dadata.Name)
    mystruct := map[string]interface{}{"success": 0, "message": "msg"}
    c.Data["json"] = &mystruct
    c.ServeJSON()

}



