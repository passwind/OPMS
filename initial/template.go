package initial

import (
	//"fmt"
	"github.com/lock-upme/opms/models/businesstrips"
	"github.com/lock-upme/opms/models/expenses"
	"github.com/lock-upme/opms/models/goouts"
	"github.com/lock-upme/opms/models/groups"
	"github.com/lock-upme/opms/models/leaves"
	"github.com/lock-upme/opms/models/oagoods"
	"github.com/lock-upme/opms/models/overtimes"
	"github.com/lock-upme/opms/models/projects"
	"github.com/lock-upme/opms/models/users"
	"github.com/lock-upme/opms/utils"
	//"time"

	"github.com/astaxie/beego"
)

func InitTemplate() {
	beego.AddFuncMap("getRealname", users.GetRealname)
	beego.AddFuncMap("getNeedsname", projects.GetProjectNeedsName)
	beego.AddFuncMap("getProjectname", projects.GetProjectName)
	beego.AddFuncMap("getPermissionname", groups.GetPermissiontName)
	beego.AddFuncMap("getLeaveProcess", leaves.ListLeaveApproverProcessHtml)
	beego.AddFuncMap("getExpenseProcess", expenses.ListExpenseApproverProcessHtml)
	beego.AddFuncMap("getBusinesstripProcess", businesstrips.ListBusinesstripApproverProcessHtml)
	beego.AddFuncMap("getGooutProcess", goouts.ListGooutApproverProcessHtml)
	beego.AddFuncMap("getOagoodProcess", oagoods.ListOagoodApproverProcessHtml)
	beego.AddFuncMap("getOvertimeProcess", overtimes.ListOvertimeApproverProcessHtml)

	beego.AddFuncMap("getDate", utils.GetDate)
	beego.AddFuncMap("getDateMH", utils.GetDateMH)
	beego.AddFuncMap("getNeedsStatus", utils.GetNeedsStatus)
	beego.AddFuncMap("getNeedsSource", utils.GetNeedsSource)
	beego.AddFuncMap("getNeedsStage", utils.GetNeedsStage)
	beego.AddFuncMap("getTaskStatus", utils.GetTaskStatus)
	beego.AddFuncMap("getTaskType", utils.GetTaskType)
	beego.AddFuncMap("getTestStatus", utils.GetTestStatus)
	beego.AddFuncMap("getLeaveType", utils.GetLeaveType)

	beego.AddFuncMap("getOs", utils.GetOs)
	beego.AddFuncMap("getBrowser", utils.GetBrowser)
	beego.AddFuncMap("getAvatarSource", utils.GetAvatarSource)
	beego.AddFuncMap("getAvatar", utils.GetAvatar)
	beego.AddFuncMap("getAvatarUserid", users.GetAvatarUserid)
	beego.AddFuncMap("getPositionsName", users.GetPositionsNameForUserid)
	beego.AddFuncMap("getDepartmentsName", users.GetDepartmentsNameForUserid)

	beego.AddFuncMap("getEdu", utils.GetEdu)
	beego.AddFuncMap("getWorkYear", utils.GetWorkYear)
	beego.AddFuncMap("getResumeStatus", utils.GetResumeStatus)

	beego.AddFuncMap("getCheckworkType", utils.GetCheckworkType)

	beego.AddFuncMap("getMessageType", utils.GetMessageType)
	beego.AddFuncMap("getMessageSubtype", utils.GetMessageSubtype)

}
