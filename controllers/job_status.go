package controllers

import (
	"app-service/status-service/models"
	"app-service/status-service/service"
	"encoding/json"
	"fmt"
	"model"

	"github.com/astaxie/beego"
)

// Operations about JobStatus
type JobStatusController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all user's jobs status
// @Param	userid		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.Response
// @router /job/:userId [get]
func (this *JobStatusController) GetAll() {
	var err error
	var response models.Response

	var userId int64
	userId, err = this.GetInt64(":userId")
	//beego.Debug("GetAll", userId)
	if userId > 0 && err == nil {
		var svc service.JobStatusService
		var allStatus []*model.JobStatus
		var result []byte
		allStatus, err = svc.GetAll(userId)
		if err == nil {
			result, err = json.Marshal(allStatus)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "user id is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}
