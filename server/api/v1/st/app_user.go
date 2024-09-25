package st

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/st"
	stReq "github.com/flipped-aurora/gin-vue-admin/server/model/st/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

type AppUserApi struct{}

// CreateAppUser 创建appUser表
// @Tags AppUser
// @Summary 创建appUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.AppUser true "创建appUser表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /appUser/createAppUser [post]
func (appUserApi *AppUserApi) CreateAppUser(c *gin.Context) {
	var appUser st.AppUser
	err := c.ShouldBindJSON(&appUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//创建uuid
	appUser.Uuid = uuid.Must(uuid.NewV4())
	err = appUserService.CreateAppUser(&appUser)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAppUser 删除appUser表
// @Tags AppUser
// @Summary 删除appUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.AppUser true "删除appUser表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /appUser/deleteAppUser [delete]
func (appUserApi *AppUserApi) DeleteAppUser(c *gin.Context) {
	ID := c.Query("ID")
	err := appUserService.DeleteAppUser(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAppUserByIds 批量删除appUser表
// @Tags AppUser
// @Summary 批量删除appUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /appUser/deleteAppUserByIds [delete]
func (appUserApi *AppUserApi) DeleteAppUserByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := appUserService.DeleteAppUserByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAppUser 更新appUser表
// @Tags AppUser
// @Summary 更新appUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.AppUser true "更新appUser表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /appUser/updateAppUser [put]
func (appUserApi *AppUserApi) UpdateAppUser(c *gin.Context) {
	var appUser st.AppUser
	err := c.ShouldBindJSON(&appUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = appUserService.UpdateAppUser(appUser)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAppUser 用id查询appUser表
// @Tags AppUser
// @Summary 用id查询appUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query st.AppUser true "用id查询appUser表"
// @Success 200 {object} response.Response{data=st.AppUser,msg=string} "查询成功"
// @Router /appUser/findAppUser [get]
func (appUserApi *AppUserApi) FindAppUser(c *gin.Context) {
	ID := c.Query("ID")
	reappUser, err := appUserService.GetAppUser(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reappUser, c)
}

// GetAppUserList 分页获取appUser表列表
// @Tags AppUser
// @Summary 分页获取appUser表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query stReq.AppUserSearch true "分页获取appUser表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /appUser/getAppUserList [get]
func (appUserApi *AppUserApi) GetAppUserList(c *gin.Context) {
	var pageInfo stReq.AppUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := appUserService.GetAppUserInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetAppUserPublic 不需要鉴权的appUser表接口
// @Tags AppUser
// @Summary 不需要鉴权的appUser表接口
// @accept application/json
// @Produce application/json
// @Param data query stReq.AppUserSearch true "分页获取appUser表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /appUser/getAppUserPublic [get]
func (appUserApi *AppUserApi) GetAppUserPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	appUserService.GetAppUserPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的appUser表接口信息",
	}, "获取成功", c)
}

// Login 登入接口
// @Tags AppUser
// @Summary 方法介绍
// @accept application/json
// @Produce application/json
// @Param data query stReq.AppUserSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /appUser/Login [GET]
func (appUserApi *AppUserApi) Login(c *gin.Context) {
	//获取openpid
	openpidcode := c.Query("openpid")
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=wx26bc089b38045c03&secret=a33a7cc27e1abe69c20caddf2fdb1a92&js_code=" + openpidcode + "&grant_type=authorization_code"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// 检查HTTP响应状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: HTTP status code %d received\n", resp.StatusCode)
		return
	}
	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// 解析JSON到SessionResponse结构体
	var session st.SessionResponse

	err = json.Unmarshal([]byte(body), &session)
	if err != nil {
		panic(err)
	}
	// 打印提取的openid
	fmt.Println("OpenID:", session.OpenID)

	res, err := appUserService.Login(session.OpenID)

	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	token, claims, err := utils.LoginToken(res)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	maxAge := int(claims.RegisteredClaims.ExpiresAt.Unix() - time.Now().Unix())
	utils.SetToken(c, token, maxAge)
	response.OkWithData(gin.H{
		"token":  token,
		"claims": claims,
		"user":   res,
		"info":   "登入成功",
	}, c)
}

// GetUserinfo 用户数据表单
// @Tags AppUser
// @Summary 用户数据表单
// @accept application/json
// @Produce application/json
// @Param data query stReq.AppUserSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /appUser/getUserinfo [GET]
func (appUserApi *AppUserApi) GetUserinfo(c *gin.Context) {
	// 请添加自己的业务逻辑
	id := utils.GetUserID(c)
	if id == 0 {
		response.FailWithMessage("获取失败", c)
		return
	}
	//获取小程序提交图片数据

	user, err := appUserService.GetUserinfo(id)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData(user, c)
}

// UpdateTheImage 更新头像
// @Tags AppUser
// @Summary 更新头像
// @accept application/json
// @Produce application/json
// @Param data query stReq.AppUserSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /appUser/updateTheImage [GET]
func (appUserApi *AppUserApi) UpdateTheImage(c *gin.Context) {
	// 请添加自己的业务逻辑
	id := utils.GetUserID(c)
	if id == 0 {
		response.FailWithMessage("获取失败", c)
		return
	}
	_, err := appUserService.UpdateTheImage(id, c.Query("image"))
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("更新成功", c)
}

// Upnickname 方法介绍
// @Tags AppUser
// @Summary 方法介绍
// @accept application/json
// @Produce application/json
// @Param data query stReq.AppUserSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /appUser/upnickname [GET]
func (appUserApi *AppUserApi) Upnickname(c *gin.Context) {
	// 请添加自己的业务逻辑
	id := utils.GetUserID(c)
	if id == 0 {
		response.FailWithMessage("获取失败", c)
		return
	}

	_, err := appUserService.Upnickname(id, c.Query("nickname"))
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("修改昵称成功", c)
}
