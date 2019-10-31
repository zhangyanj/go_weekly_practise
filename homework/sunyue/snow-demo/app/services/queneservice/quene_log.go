package queneservice

import (
	"fmt"
	"git.qufenqi.com/sunyue/go-test/week_5/sunyue/snow-demo/app/models/quenemodel"
)

func InsertQueneLog(queneData string) (Id int64, err error)  {
	queneLogInsert := quenemodel.QueneLog{
		Data : queneData,
	}
	fmt.Printf("%s", queneLogInsert)
	Id, err = quenemodel.GetInstance().GetDb().Table(queneLogInsert.TableName()).Insert(queneLogInsert)
	return
}