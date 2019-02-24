package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/model"
	"intelliq/app/service"
)

//AddNewSchool adds new school
func AddNewSchool(ctx *gin.Context) {
	var school model.School
	err := ctx.BindJSON(&school)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.AddNewSchool(&school)
	ctx.JSON(http.StatusOK, res)
}

//UpdateSchoolProfile updates existing school profile
func UpdateSchoolProfile(ctx *gin.Context) {
	var school model.School
	err := ctx.BindJSON(&school)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.UpdateSchool(&school)
	ctx.JSON(http.StatusOK, res)
}

//ListAllSchools fetches all schools under a group using groupCode or groupID
func ListAllSchools(ctx *gin.Context) {
	key := ctx.Param("key")
	val := ctx.Param("val")
	res := service.FetchAllSchools(key, val)
	ctx.JSON(http.StatusOK, res)
}

//ListSchoolByCodeOrID list school by id or code
func ListSchoolByCodeOrID(ctx *gin.Context) {
	key := ctx.Param("key")
	val := ctx.Param("val")
	res := service.FetchSchoolByCodeOrID(key, val)
	ctx.JSON(http.StatusOK, res)
}
