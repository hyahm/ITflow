package bug

import (
	"bug/handle"
	"fmt"
	"gaconfig"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func RunHttp(c chan os.Signal) {

	router := mux.NewRouter()

	//

	router.HandleFunc("/login/login", handle.Login)
	router.HandleFunc("/login/logout", handle.Loginout)
	//
	router.HandleFunc("/dashboard/usercount", handle.UserCount)
	router.HandleFunc("/dashboard/projectcount", handle.ProjectCount)

	//router.HandleFunc("/task/list", handle.TaskList)

	router.HandleFunc("/search/allbugs", handle.SearchAllBugs)
	router.HandleFunc("/search/mybugs", handle.SearchMyBugs)
	router.HandleFunc("/search/mytasks", handle.SearchMyTasks)
	router.HandleFunc("/search/bugmanager", handle.SearchBugManager)
	router.HandleFunc("/search/log", handle.SearchLog)
	router.HandleFunc("/log/classify", handle.LogClassify)

	router.HandleFunc("/get/user", handle.GetUser)
	router.HandleFunc("/get/project", handle.GetProject)
	router.HandleFunc("/get/version", handle.GetVersion)
	router.HandleFunc("/get/env", handle.GetEnv)
	router.HandleFunc("/get/status", handle.GetStatus)

	router.HandleFunc("/get/permstatus", handle.GetPermStatus)
	router.HandleFunc("/get/info", handle.GetInfo)

	router.HandleFunc("/get/thisrole", handle.GetThisRoles)
	router.HandleFunc("/get/group", handle.GetGroup)

	router.HandleFunc("/bug/pass", handle.PassBug)
	router.HandleFunc("/bug/create", handle.BugCreate)
	router.HandleFunc("/bug/edit", handle.BugEdit)
	router.HandleFunc("/bug/mybugs", handle.GetMyBugs)
	router.HandleFunc("/bug/close", handle.CloseBug)
	//router.HandleFunc("/bug/", handle.GetMyBugs)

	//router.HandleFunc("/bug/getallbugs", handle.GetAllBugs)
	router.HandleFunc("/bug/changestatus", handle.ChangeBugStatus)
	router.HandleFunc("/status/filter", handle.ChangeFilterStatus)
	router.HandleFunc("/status/show", handle.ShowStatus)
	router.HandleFunc("/bug/show", handle.BugShow)
	//router.HandleFunc("/article/list", handle.AList)

	router.HandleFunc("/uploadimg", handle.UploadImgs)
	router.HandleFunc("/upload/headimg", handle.UploadHeadImg)
	router.HandleFunc("/showimg/{imgname}", handle.ShowImg)
	//

	router.HandleFunc("/user/info", handle.UserInfo)
	router.HandleFunc("/user/list", handle.UserList)
	router.HandleFunc("/user/update", handle.UserUpdate)
	router.HandleFunc("/user/create", handle.CreateUser)
	router.HandleFunc("/user/remove", handle.RemoveUser)
	router.HandleFunc("/user/disable", handle.DisableUser)
	router.HandleFunc("/password/update", handle.ChangePassword)
	router.HandleFunc("/password/reset", handle.ResetPwd)
	//
	router.HandleFunc("/version/add", handle.AddVersion)
	router.HandleFunc("/version/list", handle.VersionList)
	router.HandleFunc("/version/remove", handle.VersionRemove)
	router.HandleFunc("/version/update", handle.VersionUpdate)
	//
	router.HandleFunc("/project/list", handle.ProjectList)
	router.HandleFunc("/project/add", handle.AddProject)
	router.HandleFunc("/project/update", handle.UpdateProject)
	router.HandleFunc("/project/delete", handle.DeleteProject)
	//
	router.HandleFunc("/env/list", handle.EnvList)
	router.HandleFunc("/env/add", handle.AddEnv)
	router.HandleFunc("/env/update", handle.UpdateEnv)
	router.HandleFunc("/env/delete", handle.DeleteEnv)
	//
	router.HandleFunc("/department/add", handle.AddBugGroup)
	router.HandleFunc("/department/edit", handle.EditBugGroup)
	router.HandleFunc("/department/list", handle.BugGroupList)
	router.HandleFunc("/department/remove", handle.BugGroupDel)

	router.HandleFunc("/role/add", handle.AddRole)
	router.HandleFunc("/role/edit", handle.EditRole)
	router.HandleFunc("/role/list", handle.RoleList)
	router.HandleFunc("/role/remove", handle.RoleDel)
	router.HandleFunc("/role/get", handle.GetRoles)
	router.HandleFunc("/role/groupname", handle.RoleGroupName)

	router.HandleFunc("/email/test", handle.TestEmail)
	router.HandleFunc("/email/save", handle.SaveEmail)
	router.HandleFunc("/email/get", handle.GetEmail)

	router.HandleFunc("/status/list", handle.StatusList)
	router.HandleFunc("/status/add", handle.StatusAdd)
	router.HandleFunc("/status/remove", handle.StatusRemove)
	router.HandleFunc("/status/update", handle.StatusUpdate)
	router.HandleFunc("/status/groupname", handle.StatusGroupName)

	router.HandleFunc("/admin/reset", handle.Reset)
	router.HandleFunc("/info/update", handle.UpdateInfo)
	//router.HandleFunc("/role/update", handle.UpdateRoles)
	router.HandleFunc("/log/list", handle.LogList)

	//-------------------------------------------
	router.HandleFunc("/share/list", handle.ShareList)
	router.HandleFunc("/share/upload", handle.ShareUpload)
	router.HandleFunc("/share/mkdir", handle.ShareMkdir)
	router.HandleFunc("/share/remove", handle.ShareRemove)
	router.HandleFunc("/share/rename", handle.ShareRename)
	//router.HandleFunc("/share/down", handle.ShareDownload)
	router.HandleFunc("/share/down", handle.ShareShow)

	//----------------------------------------------
	router.HandleFunc("/rest/list", handle.RestList)
	router.HandleFunc("/rest/update", handle.RestUpdate)
	router.HandleFunc("/rest/add", handle.RestAdd)
	router.HandleFunc("/rest/delete", handle.RestDel)

	router.HandleFunc("/api/list", handle.ApiList)
	router.HandleFunc("/api/update", handle.ApiUpdate)
	router.HandleFunc("/api/add", handle.ApiAdd)
	router.HandleFunc("/api/delete", handle.ApiDel)
	router.HandleFunc("/api/one", handle.ApiOne)
	router.HandleFunc("/edit/one", handle.EditOne)
	router.HandleFunc("/api/resp", handle.ApiResp)

	router.HandleFunc("/type/list", handle.TypeList)
	router.HandleFunc("/type/update", handle.TypeUpdate)
	router.HandleFunc("/type/add", handle.TypeAdd)
	router.HandleFunc("/type/delete", handle.TypeDel)
	router.HandleFunc("/type/get", handle.GetType)

	router.HandleFunc("/group/get", handle.GroupGet)
	router.HandleFunc("/group/add", handle.GroupAdd)
	router.HandleFunc("/group/del", handle.GroupDel)
	router.HandleFunc("/group/update", handle.GroupUpdate)

	router.HandleFunc("/header/list", handle.HeaderList)
	router.HandleFunc("/header/add", handle.HeaderAdd)
	router.HandleFunc("/header/del", handle.HeaderDel)
	router.HandleFunc("/header/update", handle.HeaderUpdate)
	router.HandleFunc("/header/get", handle.HeaderGet)

	router.HandleFunc("/important/get", handle.ImportantGet)
	router.HandleFunc("/important/add", handle.ImportantAdd)
	router.HandleFunc("/important/del", handle.ImportantDel)
	router.HandleFunc("/important/update", handle.ImportantUpdate)
	router.HandleFunc("/get/importants", handle.GetImportants)

	router.HandleFunc("/level/get", handle.LevelGet)
	router.HandleFunc("/level/add", handle.LevelAdd)
	router.HandleFunc("/level/del", handle.LevelDel)
	router.HandleFunc("/level/update", handle.LevelUpdate)
	router.HandleFunc("/get/levels", handle.GetLevels)

	router.HandleFunc("/position/list", handle.PositionGet)
	router.HandleFunc("/position/add", handle.PositionAdd)
	router.HandleFunc("/position/del", handle.PositionDel)
	router.HandleFunc("/position/update", handle.PositionUpdate)
	router.HandleFunc("/get/hypos", handle.GetHypos)
	router.HandleFunc("/get/positions", handle.GetPositions)
	//---------------------------------------------------------
	router.HandleFunc("/default/status", handle.DefaultStatus)
	router.HandleFunc("/default/save", handle.DefaultSave)
	router.HandleFunc("/default/important", handle.DefaultImportant)
	router.HandleFunc("/default/level", handle.DefaultLevel)

	listenaddr := gaconfig.ReadString("listenaddr")
	fmt.Println("listen on " + listenaddr)
	if gaconfig.ReadBool("ssl") {
		if err := http.ListenAndServeTLS(listenaddr,gaconfig.ReadString("certfile"), gaconfig.ReadString("keyfile"), router) ; err != nil {
			log.Fatal(err)
		}
	} else {
		if err := http.ListenAndServe(listenaddr, router); err != nil {
			log.Fatal(err)
		}
	}


	signal.Notify(c, os.Interrupt, os.Kill)
}
