package controllers

import (
	"fmt"
	"git.qufenqi.com/golang/go_weekly_practise/homework/guoxiaofeng/app/constants/errorcode"
	"git.qufenqi.com/golang/go_weekly_practise/homework/guoxiaofeng/app/http/entities"
	"git.qufenqi.com/golang/go_weekly_practise/homework/guoxiaofeng/app/jobs/basejob"
	"git.qufenqi.com/golang/go_weekly_practise/homework/guoxiaofeng/app/models/usermodel"
	"github.com/gin-gonic/gin"
	"github.com/qit-team/snow-core/utils"
	"github.com/qit-team/work"
)

//队列test
// HandleTestValidator godoc
// @Summary HandleQueneProcess的示例
// @Description HandleQueneProcess的示例
// @Tags snow
// @Accept  json
// @Produce json
// @Param queneProcess body entities.QueneProcessRequest true "example of validator"
// @Success 200 {array} entities.TestQueueResponse
// @Failure 400 {object} controllers.HTTPError
// @Failure 404 {object} controllers.HTTPError
// @Failure 500 {object} controllers.HTTPError
// @Router /quene/test [post]
func QueueTest(c *gin.Context) {
	request := new(entities.TestQueueRequest)
	err := GenRequest(c, request)
	if err != nil {
		Error(c, errorcode.ParamError)
		return
	}
	user := usermodel.User{
		Name:   request.Name,
		Mobile: request.Mobile,
		Email:  request.Email,
		Status: 0,
	}
	affectedRows, e := usermodel.GetInstance().SaveUser(user)
	if e != nil {
		Error(c, errorcode.ParamError)
		return
	}
	jsonMessage, err := utils.JsonEncode(request)
	fmt.Println(jsonMessage)
	task := work.Task{
		Id:      request.Name,
		Message: jsonMessage,
	}
	ok, err := basejob.EnqueueWithTask(c, "test-queue", task)
	fmt.Println(err)
	if ok {
		response := new(entities.TestQueueResponse)
		response.Status = affectedRows
		Success(c, response)
	} else {
		Error(c, errorcode.SystemError)
	}

	return
}
