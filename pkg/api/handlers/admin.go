package handlers

import (
	"net/http"
	"strconv"

	"github.com/ECOMMERCE_PROJECT/pkg/api/handlerutils"
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUseCase services.AdminUsecase
}

func NewAdminHandler(adminUseCae services.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: adminUseCae,
	}
}

func (cr *AdminHandler) CreateAdmin(c *gin.Context) {
	var adminData helperstruct.CreateAdmin
	err := c.Bind(&adminData)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	createrId, err := handlerutils.GetAdminIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find AdminId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	admin, err := cr.adminUseCase.CreateAdmin(c.Request.Context(), adminData, createrId)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Create Admin",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "Admin created",
		Data:       admin,
		Errors:     nil,
	})
}

func (cr *AdminHandler) AdminLoging(c *gin.Context) {
	var admin helperstruct.LoginReq
	err := c.Bind(&admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	ss, err := cr.adminUseCase.AdminLogin(admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AdminAuth", ss, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "logined success fully",
		Data:       nil,
		Errors:     nil,
	})
}

func (cr *AdminHandler) AdminLogout(c *gin.Context) {
	c.SetCookie("AdminAuth", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "admin logouted",
		Data:       nil,
		Errors:     nil,
	})

}

func (cr *AdminHandler) BlockUser(c *gin.Context) {
	var body helperstruct.BlockData
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	adminId, err := handlerutils.GetAdminIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find AdminId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.adminUseCase.BlockUser(body, adminId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Block",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "User Blocked",
		Data:       nil,
		Errors:     nil,
	})
}

func (cr *AdminHandler) UnblockUser(c *gin.Context) {
	paramsId := c.Param("user_id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.adminUseCase.UnblockUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant unblock user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "user unblocked",
		Data:       nil,
		Errors:     nil,
	})
}
func (cr *AdminHandler) FindUser(c *gin.Context) {
	paramsID := c.Param("user_id")
	id, err := strconv.Atoi(paramsID)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Bind failed",
		})
	}
	user, err := cr.adminUseCase.FindUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find User",
			Data:       nil,
			Errors:     err,
		})
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "User found",
		Data:       user,
		Errors:     nil,
	})
}
func (cr *AdminHandler) ListAllUsers(c *gin.Context) {

	userList, err := cr.adminUseCase.ListAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Error listing users",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "List of all users",
		Data:       userList,
		Errors:     nil,
	})
}
func (cr *AdminHandler) GetDashBoard(c*gin.Context){
	DashBoard, err := cr.adminUseCase.GetDashBoard()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: 500,
			Message:    "Error in getting Admin Dashboard",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Admin DashBoard",
		Data:       DashBoard,
		Errors:     nil,
	})

}
func(cr*AdminHandler)ViewDailySalesReport(c*gin.Context){
	sales ,err:=cr.adminUseCase.ViewDailySalesReport()
	if err !=nil {c.JSON(http.StatusBadRequest,response.Response{
		StatusCode:400 ,
		Message: "Can't get daily sales report",
		Data: nil,
		Errors: err.Error(),
	})
     return
	}
	c.JSON(http.StatusOK,response.Response{
		StatusCode: 200,
		Message: "daily sales report :",
		Data: sales,
		Errors: nil,
	})
}
func(cr*AdminHandler)ViewWeelySalesReport(c*gin.Context){
	sales ,err:=cr.adminUseCase.ViewWeeklySalesReport()
	if err !=nil {c.JSON(http.StatusBadRequest,response.Response{
		StatusCode:400 ,
		Message: "Can't get daily sales report",
		Data: nil,
		Errors: err.Error(),
	})
     return
	}
	c.JSON(http.StatusOK,response.Response{
		StatusCode: 200,
		Message: "weekly sales report :",
		Data: sales,
		Errors: nil,
	})
}
func(cr*AdminHandler)ViewMonthlySalesReport(c*gin.Context){
	sales ,err:=cr.adminUseCase.ViewMonthlySalesReport()
	if err !=nil {c.JSON(http.StatusBadRequest,response.Response{
		StatusCode:400 ,
		Message: "Can't get daily sales report",
		Data: nil,
		Errors: err.Error(),
	})
     return
	}
	c.JSON(http.StatusOK,response.Response{
		StatusCode: 200,
		Message: "monthly sales report :",
		Data: sales,
		Errors: nil,
	})
}
func(cr*AdminHandler)ViewYearlySalesReport(c*gin.Context){
	sales ,err:=cr.adminUseCase.ViewYearlySalesReport()
	if err !=nil {c.JSON(http.StatusBadRequest,response.Response{
		StatusCode:400 ,
		Message: "Can't get daily sales report",
		Data: nil,
		Errors: err.Error(),
	})
     return
	}
	c.JSON(http.StatusOK,response.Response{
		StatusCode: 200,
		Message: "yearly sales report :",
		Data: sales,
		Errors: nil,
	})
}