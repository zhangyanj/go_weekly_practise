package userjob

import (
	"encoding/json"
	"fmt"
	"github.com/go_weekly_practise/homework/wangjunxiong/snow/app/http/entities"
	"github.com/go_weekly_practise/homework/wangjunxiong/snow/app/services/userservice"
	"github.com/qit-team/work"
	"log"
)

const (
	TOPIC = "topic_user"
)

func Login(task work.Task) work.TaskResult {
	s, err := work.JsonEncode(task)
	if err != nil {
		//work.StateFailed 不会进行ack确认
		//work.StateFailedWithAck 会进行actk确认
		//return work.TaskResult{Id: task.Id, State: work.StateFailed}
		return work.TaskResult{Id: task.Id, State: work.StateFailedWithAck}
	} else {
		//work.StateSucceed 会进行ack确认
		//id,err := userservice.InsertLoginInfo()
		fmt.Println("do task", s)
		fmt.Println("do task", task.Message)
		var t entities.UserLoginRequest
		err := json.Unmarshal([]byte(task.Message), &t)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		fmt.Println("do task", t)
		id, err := userservice.InsertLoginInfo(t.UserId, t.Ip, t.LoginTime)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		fmt.Println("succces log", id)
		return work.TaskResult{Id: task.Id, State: work.StateSucceed}
	}
}
