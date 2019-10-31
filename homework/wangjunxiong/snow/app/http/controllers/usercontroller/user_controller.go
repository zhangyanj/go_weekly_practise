package usercontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go_weekly_practise/homework/wangjunxiong/snow/app/http/controllers"
	"github.com/go_weekly_practise/homework/wangjunxiong/snow/app/http/entities"
	"github.com/go_weekly_practise/homework/wangjunxiong/snow/app/jobs/basejob"
	"github.com/go_weekly_practise/homework/wangjunxiong/snow/app/jobs/userjob"
	"github.com/qit-team/work"
	"math/rand"
	"strconv"
	"time"
)

func HandlePostUserLogin(c *gin.Context) {
	request := new(entities.UserLoginRequest)
	request.LoginTime = time.Now()
	err := controllers.GenRequest(c, request)
	if err != nil {
		fmt.Println(err)
		controllers.Error(c, 500)
		return
	}

	jsonBytes, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err);
		controllers.Error(c, 500)
		return
	}

	task := work.Task{
		Id:      strconv.Itoa(rand.Intn(100)),
		Message: string(jsonBytes),
	}
	ok, err := basejob.EnqueueWithTask(context.TODO(), userjob.TOPIC, task)

	if ok {
		controllers.Success(c, 200)

	} else {
		fmt.Println(err);
		controllers.Error(c, 500)
	}
	return

}
