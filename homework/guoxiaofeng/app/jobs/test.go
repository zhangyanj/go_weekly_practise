package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"git.qufenqi.com/golang/go_weekly_practise/homework/guoxiaofeng/app/http/entities"
	"git.qufenqi.com/golang/go_weekly_practise/homework/guoxiaofeng/app/models/usermodel"
	"github.com/qit-team/snow-core/log/logger"
	"github.com/qit-team/work"
	"time"
)

func test(task work.Task) work.TaskResult {
	time.Sleep(time.Millisecond * 5)
	//fmt.Println(task)

	fmt.Println(task.Id)
	fmt.Println(2)
	params := new(entities.TestQueueRequest)
	err := json.Unmarshal([]byte(task.Message), params)
	fmt.Println(params.Name)
	fmt.Println(params.Mobile)
	fmt.Println(params.Email)
	if err != nil {
		logger.Error(context.Background(), "test-queue", err.Error(), logger.NewWithField("task.Message", task.Message), logger.NewWithField("err", err))
	}
	user := usermodel.User{
		Status: 2,
	}
	affectedRows, err := usermodel.GetInstance().Update(task.Id, user)
	fmt.Println(affectedRows, err)
	//fmt.Println("do task", affectedRows)
	return work.TaskResult{Id: task.Id, State: work.StateSucceed}
}
