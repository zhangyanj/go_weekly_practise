package controllers

import (
	"fmt"
	"git.qufenqi.com/sunyue/go-test/week_5/sunyue/snow-demo/app/constants/errorcode"
	"git.qufenqi.com/sunyue/go-test/week_5/sunyue/snow-demo/app/http/entities"
	"git.qufenqi.com/sunyue/go-test/week_5/sunyue/snow-demo/app/jobs/basejob"
	"github.com/gin-gonic/gin"
	"github.com/qit-team/snow-core/log/logger"
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
// @Success 200 {array} entities.QueneProcessRequest
// @Failure 400 {object} controllers.HTTPError
// @Failure 404 {object} controllers.HTTPError
// @Failure 500 {object} controllers.HTTPError
// @Router /quene/process [post]
func HandleQueneProcess(c *gin.Context)  {
	request := new(entities.QueneProcessRequest)
	err := GenRequest(c, request)
	fmt.Println(err)
	if err != nil {
		Error(c, errorcode.ParamError)
		return
	}

	logger.Info(c, "quene-test", "enter")

	jsonMessage, err := utils.JsonEncode(request)
	task := work.Task{
		Id: string(request.Id),
		Message: jsonMessage,
	}
	ok, err := basejob.EnqueueWithTask(c, "quene-test", task)
	if ok {
		Success(c, "success")
	} else {
		Error(c, errorcode.SystemError)
	}
}
