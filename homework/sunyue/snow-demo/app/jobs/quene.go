package jobs

import (
	"context"
	"fmt"
	"git.qufenqi.com/sunyue/go-test/week_5/sunyue/snow-demo/app/services/queneservice"
	"github.com/qit-team/snow-core/log/logger"
	"github.com/qit-team/work"
	"time"
)

func quene(task work.Task) (work.TaskResult) {
	time.Sleep(time.Millisecond * 5)
	s, err := work.JsonEncode(task)
	if err != nil {
		//work.StateFailed 不会进行ack确认
		//work.StateFailedWithAck 会进行actk确认
		//return work.TaskResult{Id: task.Id, State: work.StateFailed}
		return work.TaskResult{Id: task.Id, State: work.StateFailedWithAck}
	} else {
		//work.StateSucceed 会进行ack确认
		logger.Info(context.TODO(), "quene-test", s, "exit")
		logId, errInsert := queneservice.InsertQueneLog(s)
		if errInsert != nil {
			logger.Error(context.TODO(), "quene-test", "insert err")
			fmt.Println(errInsert)
		}
		return work.TaskResult{Id: string(logId), State: work.StateSucceed}
	}

}