package jobs

import (
	"fmt"
	"git.qufenqi.com/golang/go_weekly_practise/homework/guoxiaofeng/app/models/usermodel"
	"github.com/qit-team/work"
	"time"
)

func test(task work.Task) work.TaskResult {
	time.Sleep(time.Millisecond * 5)
	//fmt.Println(task)
	s, err := work.JsonEncode(task)
	if err != nil {
		//work.StateFailed 不会进行ack确认
		//work.StateFailedWithAck 会进行actk确认
		//return work.TaskResult{Id: task.Id, State: work.StateFailed}
		return work.TaskResult{Id: task.Id, State: work.StateFailedWithAck}
	} else {
		fmt.Println(task.Id)
		/*user, err := usermodel.GetInstance().GetListByName(task.Id)
		fmt.Println(user, err)
		if err != nil {
			return work.TaskResult{Id: task.Id, State: work.StateFailedWithAck}
		}*/
		user := usermodel.User{
			Status: 2,
		}
		affectedRows, err := usermodel.GetInstance().Update(task.Id, user)
		fmt.Println(affectedRows, err)
		/*if err != nil {
			return work.TaskResult{Id: task.Id, State: work.StateFailedWithAck}
		}*/
		fmt.Println("do task", s)
		//fmt.Println("do task", affectedRows)
		return work.TaskResult{Id: task.Id, State: work.StateSucceed}
	}

}
