package user

import (
	"net/http"

	"../../constants"
	"github.com/EDDYCJY/edit-video-go/pkg/app"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary Get Users
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/users [get]
func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	userService := tag_service.Tag{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	tags, err := userService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_GET_TAGS_FAIL, nil)
		return
	}

	count, err := userService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_COUNT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, constants.SUCCESS, map[string]interface{}{
		"lists": tags,
		"total": count,
	})
}

type AddUserForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Add User
// @Produce  json
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/users [post]
func AddUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddUserForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != constants.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userService := tag_service.Tag{
		Name:      form.Name,
		CreatedBy: form.CreatedBy,
		State:     form.State,
	}
	exists, err := userService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, constants.ERROR_EXIST_TAG, nil)
		return
	}

	err = userService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, constants.SUCCESS, nil)
}

type EditUserForm struct {
	ID         int    `form:"id" valid:"Required;Min(1)"`
	Name       string `form:"name" valid:"Required;MaxSize(100)"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Update User
// @Produce  json
// @Param id path int true "ID"
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param modified_by body string true "ModifiedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/users/{id} [put]
func EditUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = EditTagForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != constants.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	userService := tag_service.Tag{
		ID:         form.ID,
		Name:       form.Name,
		ModifiedBy: form.ModifiedBy,
		State:      form.State,
	}

	exists, err := userService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, constants.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = userService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_EDIT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, constants.SUCCESS, nil)
}

// @Summary Delete User
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/users/{id} [delete]
func DeleteUser(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, constants.INVALID_PARAMS, nil)
	}

	userService := tag_service.Tag{ID: id}
	exists, err := userService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, constants.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	if err := userService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_DELETE_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, constants.SUCCESS, nil)
}
