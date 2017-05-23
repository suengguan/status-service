package service

import (
	"strconv"

	"model"
	"utility/fileoperator"

	daoApi "api/dao_service"
	"encoding/json"

	"github.com/astaxie/beego"
)

type JobStatusService struct {
}

func (this *JobStatusService) GetAll(userId int64) ([]*model.JobStatus, error) {
	var err error
	var allStatus []*model.JobStatus

	// get user
	beego.Debug("->get user")
	var user *model.User
	user, err = daoApi.UserDaoApi.GetById(userId)
	if err != nil {
		beego.Debug(err)
		return nil, err
	}

	// get projects
	beego.Debug("->get all projects")
	var projects []*model.Project
	projects, err = daoApi.BussinessDaoApi.GetAllProjects(userId)
	if err != nil {
		beego.Debug("get projects failed")
		beego.Debug(err)
		return nil, err
	}

	// get status
	beego.Debug("->get job status")
	for _, p := range projects {
		for _, j := range p.Jobs {
			var status *model.JobStatus
			status, err = this.getJobStatus(j, p, user)
			if err != nil {
				beego.Debug("get job status failed")
				beego.Debug(err)
				return nil, err
			}

			allStatus = append(allStatus, status)
		}
	}

	beego.Debug("result:", allStatus)

	return allStatus, err
}

func (this *JobStatusService) getJobStatus(job *model.Job, project *model.Project, user *model.User) (*model.JobStatus, error) {
	var err error
	var status model.JobStatus
	status.UserName = user.Name
	status.ProjectId = project.Id
	status.ProjectName = project.Name
	status.JobId = job.Id
	status.JobName = job.Name
	status.JobDescription = job.Description

	// todo
	var sumProgress int
	var podCnt int
	for _, m := range job.Modules {
		podCnt = len(m.Pods)
		for _, p := range m.Pods {
			// get value
			var fn string
			var cfg = beego.AppConfig
			// if config.DEBUG_ONLY {
			// 	fn = "C:/Works/PME2017/solution/PMEServer/workspace/debug.st"
			// } else {
			// 	fn = config.WORKSPACE_PATH
			// 	// pod path
			// 	fn += "/" + user.Name
			// 	fn += "/" + project.Name + "-" + strconv.FormatInt(project.Id, 36)
			// 	fn += "/" + job.Name + "-" + strconv.FormatInt(job.Id, 36)
			// 	fn += "/" + m.Name + "-" + strconv.FormatInt(m.Id, 36)
			// 	fn += "/" + p.Name + "-" + strconv.FormatInt(p.Id, 36) + ".st"
			// 	beego.Debug("status file:", fn)
			// }
			if cfg.String("runmode") == "dev-windows" {
				fn = "C:/Works/PME2017/solution/PMEServer/workspace/debug.st"
			} else {
				fn = cfg.String("workspace")
				// pod path
				fn += "/" + user.Name
				fn += "/" + project.Name + "-" + strconv.FormatInt(project.Id, 36)
				fn += "/" + job.Name + "-" + strconv.FormatInt(job.Id, 36)
				fn += "/" + m.Name + "-" + strconv.FormatInt(m.Id, 36)
				fn += "/" + p.Name + "-" + strconv.FormatInt(p.Id, 36) + ".st"
				beego.Debug("status file:", fn)
			}

			var statusStr string
			statusStr, err = fileoperator.Read(fn)
			if err == nil {
				var podStatus model.PodStatus
				err = json.Unmarshal(([]byte)(statusStr), &podStatus)

				status.Status = podStatus.Status
				sumProgress += podStatus.Progress
			} else {
				return nil, err
			}

		}
	}
	status.Progress = sumProgress / podCnt

	return &status, err
}
