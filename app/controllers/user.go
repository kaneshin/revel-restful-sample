package controllers

import (
	"revel-boilerplate/app/models"

	"github.com/revel/revel"
)

type User struct {
	Controller
}

func result(err error) map[string]interface{} {
	result := map[string]interface{}{}
	status := (err == nil)
	result["status"] = status
	if !status {
		result["error"] = map[string]interface{}{
			"reason": err.Error(),
		}
	}
	return result
}

func (c User) Create() revel.Result {
	params := c.Params.Form
	userModel := models.GetUserInstance()
	err := userModel.Create(params)
	result := result(err)
	return c.RenderJson(result)
}

func (c User) Read(id int64) revel.Result {
	userModel := models.GetUserInstance()
	user, err := userModel.Get(id)
	result := result(err)
	if user != nil {
		result["user"] = user
	}
	return c.RenderJson(result)
}

func (c User) Update(id int64) revel.Result {
	params := c.Params.Form
	userModel := models.GetUserInstance()
	err := userModel.Update(id, params)
	result := result(err)
	return c.RenderJson(result)
}

func (c User) Delete(id int64) revel.Result {
	userModel := models.GetUserInstance()
	err := userModel.Delete(id)
	result := result(err)
	return c.RenderJson(result)
}
