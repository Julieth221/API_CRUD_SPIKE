package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/models"

	"github.com/astaxie/beego"
)

// GeolocalizacionController operations for Geolocalizacion
type GeolocalizacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *GeolocalizacionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Geolocalizacion
// @Param	body		body 	models.Geolocalizacion	true		"body for Geolocalizacion content"
// @Success 201 {int} models.Geolocalizacion
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *GeolocalizacionController) Post() {
	var v models.Geolocalizacion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddGeolocalizacion(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Registration successful", "Data": v}
		} else {
			c.Data["mesaage"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
			c.Abort("400")
		}
	} else {
		c.Data["mesaage"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Geolocalizacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Geolocalizacion
// @Failure 404 not found resource
// @router /:id [get]
func (c *GeolocalizacionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetGeolocalizacionById(id)
	if err != nil {
		c.Data["mesaage"] = "Error service GetOne: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	} else {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Request successful", "Data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Geolocalizacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Geolocalizacion
// @Failure 404 not found resource
// @router / [get]
func (c *GeolocalizacionController) GetAll() {
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

	l, err := models.GetAllGeolocalizacion(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["mesaage"] = "Error service GetAll: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Request successful", "Data": l}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Geolocalizacion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Geolocalizacion	true		"body for Geolocalizacion content"
// @Success 200 {object} models.Geolocalizacion
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *GeolocalizacionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Geolocalizacion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateGeolocalizacionById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Update successful", "Data": v}
		} else {
			c.Data["mesaage"] = "Error service Put: The request contains an incorrect data type or an invalid parameter"
			c.Abort("400")
		}
	} else {
		c.Data["mesaage"] = "Error service Put: The request contains an incorrect data type or an invalid parameter"
		c.Abort("400")
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Geolocalizacion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *GeolocalizacionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteGeolocalizacion(id); err == nil {
		d := map[string]interface{}{"Id": id}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Delete successful", "Data": d}
	} else {
		c.Data["mesaage"] = "Error service Delete: Request contains incorrect parameter"
		c.Abort("404")
	}
	c.ServeJSON()
}

// Patch ...
// @Title Patch
// @Description update partially the Geolocalizacion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	map[string]interface{}	true		"Partial data to update Geolocalizacion"
// @Success 200 {object} models.Geolocalizacion
// @Failure 400 the request contains incorrect syntax
// @Failure 404 the geolocalizacion is not found
// @router /:id [patch]
func (c *GeolocalizacionController) Patch() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["message"] = "Error service Patch: Invalid ID format"
		c.Abort("400")
		return
	}

	// Obtener la geolocalización existente
	geolocalizacion, err := models.GetGeolocalizacionById(id)
	if err != nil {
		c.Data["message"] = "Error service Patch: Geolocalizacion not found"
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
		case "LatitudInicial":
			if latInicial, ok := value.(string); ok {
				geolocalizacion.LatitudInicial = latInicial
			}
		case "LongitudInicial":
			if lonInicial, ok := value.(string); ok {
				geolocalizacion.LongitudInicial = lonInicial
			}
		case "LatitudFinal":
			if latFinal, ok := value.(string); ok {
				geolocalizacion.LatitudFinal = latFinal
			}
		case "LongitudFinal":
			if lonFinal, ok := value.(string); ok {
				geolocalizacion.LongitudFinal = lonFinal
			}
		}
	}

	// Guardar los cambios en la base de datos
	if err := models.UpdateGeolocalizacionById(geolocalizacion); err == nil {
		c.Data["json"] = map[string]interface{}{
			"Success": true,
			"Status":  "200",
			"Message": "Patch update successful",
			"Data":    geolocalizacion,
		}
	} else {
		c.Data["message"] = "Error service Patch: Update failed"
		c.Abort("400")
	}

	c.ServeJSON()
}
