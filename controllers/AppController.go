package controllers

import (
	// "log"

	"GO-CRUD/models"
	"encoding/json"
	"strconv"

	// "fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type AppController struct {
	web.Controller
}

func (c *AppController) GetAll() {
// 	log.Print("Hi from Get All")
	var demo = []models.Demo{}
	o := orm.NewOrm()
	sql := "select * from demo"

	_,err := o.Raw(sql).QueryRows(&demo) //number of rows , error

	if err != nil {
		if err == orm.ErrNoRows {
			c.CustomAbort(404, "No Student Data Found") //customAbort - support status Code and msg
			return
		}
		c.CustomAbort(500, "Database Error : "+ err.Error())
		return
	}

	c.Data["json"] = demo
	c.ServeJSON()
}

func (c *AppController) GetAllStudent() {
	var students = []models.Student{}
	o := orm.NewOrm()
	sql := "select * from student"

	rowsCount,err := o.Raw(sql).QueryRows(&students) //QueryRows - for multiple rows

	if err != nil {
		if err == orm.ErrNoRows || rowsCount == 0 {
			c.CustomAbort(404, "No Student Data Found")
			return
		}
		c.CustomAbort(500, "Database Error : "+ err.Error())
		return
	}

	c.Data["json"] = students
	c.ServeJSON()
}

func (c *AppController) GetStudent() {
	studentId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	var student = models.Student{}
	o := orm.NewOrm()
	sql := "select * from student where id = ?"

	err := o.Raw(sql,studentId).QueryRow(&student)
	
	if err != nil {
		if err == orm.ErrNoRows {
			c.CustomAbort(404, "No Student Data Found")
			return
		}
		c.CustomAbort(500, "Database Error : "+ err.Error())
		return
	}

	c.Data["json"] = student
	c.ServeJSON()
}

func (c *AppController) CreateStudent() {
	var student models.Student

	requestBody := c.Ctx.Input.RequestBody

	//standard function to decode json data to a GO variable
	err := json.Unmarshal(requestBody , &student) 

	if err != nil {
		c.CustomAbort(400, "Invalid JSON")
		return
	}

	o := orm.NewOrm()

	//raw sql
	sql := "INSERT INTO student( name, address, class, school) values (?, ?, ?, ?);"

	res , err := o.Raw(sql, student.Name , student.Address , student.Class , student.StudentSchool).Exec()

	if err != nil {
		c.CustomAbort(500 , "Insert Failed")
		return
	}

	id,_ := res.LastInsertId()
	c.Data["json"] = id
	c.ServeJSON()
}

func (c *AppController) UpdateStudent() {
	studentId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	var student models.Student

	requestBody := c.Ctx.Input.RequestBody
	err := json.Unmarshal(requestBody , &student) 

	if err != nil {
		c.CustomAbort(400, "Invalid JSON")
		return
	}

	o := orm.NewOrm()

	sql := `UPDATE student 
		   SET name = ?, address = ?, class = ?, school = ?
		   WHERE id = ?`
	
	res , err := o.Raw(sql, student.Name , student.Address , student.Class , student.StudentSchool, studentId).Exec()

	if err != nil {
		c.CustomAbort(500, "Database Error : "+ err.Error())
		return
	}

	//confirm the update
	_, rowsErr := res.RowsAffected()

	if rowsErr != nil {
		c.CustomAbort(404, "Student not found in DB")
		return
	}

	c.Data["json"] = map[string]interface{}{
		"message" : "Student Updated Successfully",
		"id" : studentId,
	}

	c.ServeJSON()
}


