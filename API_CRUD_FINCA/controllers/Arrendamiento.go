package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/models"

	"github.com/astaxie/beego"
)

// ArrendamientoController operations for Arrendamiento
type ArrendamientoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ArrendamientoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Arrendamiento
// @Param	body		body 	models.Arrendamiento	true		"body for Arrendamiento content"
// @Success 201 {int} models.Arrendamiento
// @Failure 403 body is empty
// @router / [post]
func (c *ArrendamientoController) Post() {
	var v models.Arrendamiento
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddArrendamiento(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Arrendamiento by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Arrendamiento
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ArrendamientoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetArrendamientoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Arrendamiento
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Arrendamiento
// @Failure 403
// @router / [get]
func (c *ArrendamientoController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllArrendamiento(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Arrendamiento
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Arrendamiento	true		"body for Arrendamiento content"
// @Success 200 {object} models.Arrendamiento
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ArrendamientoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Arrendamiento{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateArrendamientoById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Arrendamiento
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ArrendamientoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteArrendamiento(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Patch ...
// @Title Patch
// @Description update partially the Arrendamiento
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	map[string]interface{}	true		"Partial data to update Arrendamiento"
// @Success 200 {object} models.Arrendamiento
// @Failure 400 the request contains incorrect syntax
// @Failure 404 the arrendamiento is not found
// @router /:id [patch]
func (c *ArrendamientoController) Patch() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["message"] = "Error service Patch: Invalid ID format"
		c.Abort("400")
		return
	}

	// Obtener el arrendamiento existente
	arrendamiento, err := models.GetArrendamientoById(id)
	if err != nil {
		c.Data["message"] = "Error service Patch: Arrendamiento not found"
		c.Abort("404")
		return
	}

	// Leer los datos enviados en la solicitud
	var updateData map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &updateData); err != nil {
		c.Data["message"] = "Error service Patch: Invalid JSON format"
		c.Abort("400")
		return
	}

	// Actualizar solo los campos presentes en la solicitud
	for key, value := range updateData {
		switch key {
		case "Activo":
			if activo, ok := value.(bool); ok {
				arrendamiento.Activo = &activo
			}
		case "FkArrendamientoFinca":
			if fincaMap, ok := value.(map[string]interface{}); ok {
				if idFinca, ok := fincaMap["Id"].(float64); ok {
					arrendamiento.FkArrendamientoFinca = &models.Finca{Id: int(idFinca)}
				}
			}
		case "IdUserUserArrendatario":
			if userMap, ok := value.(map[string]interface{}); ok {
				if idUser, ok := userMap["Id"].(float64); ok {
					arrendamiento.IdUserUserArrendatario = &models.UserArrendatario{Id: int(idUser)}
				}
			}
		case "FkArrendamientoAnterior":
			if arrAntMap, ok := value.(map[string]interface{}); ok {
				if idAnterior, ok := arrAntMap["Id"].(float64); ok {
					arrendamiento.FkArrendamientoAnterior = &models.Arrendamiento{Id: int(idAnterior)}
				}
			}
		}
	}

	// Guardar los cambios en la base de datos
	if err := models.UpdateArrendamientoById(arrendamiento); err == nil {
		c.Data["json"] = map[string]interface{}{
			"Success": true,
			"Status":  "200",
			"Message": "Patch update successful",
			"Data":    arrendamiento,
		}
	} else {
		c.Data["message"] = "Error service Patch: Update failed"
		c.Abort("400")
	}

	c.ServeJSON()
}
