package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct{}

type Header struct {
	api_key string
}

type Persons struct {
	Name   string `json:"name" binding:"required,alpha,min=2,max=10"`
	Family string `json:"family" binding:"required,alpha,min=3,max=15"`
	Mobile string `json:"mobile" binding:"required,mobile,min=11,max=11"`
	Age    int    `json:"age" binding:"numeric,gte=0,lte=120"`
	Gender string `json:"gender" binding:"required,oneof=male female"`
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (t *TestHandler) TestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}

func (t *TestHandler) UserById(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"message": "test",
		"id":      id,
	})
}

// read from header
func (t *TestHandler) HeaderBinderMethod1(c *gin.Context) {
	api_key := c.GetHeader("api_key")

	c.JSON(http.StatusOK, gin.H{
		"message": "HeaderBinderMethod1",
		"api_key": api_key,
	})
}

func (t *TestHandler) HeaderBinderMethod2(c *gin.Context) {
	api_key := c.Request.Header.Get("api_key")

	c.JSON(http.StatusOK, gin.H{
		"message": "HeaderBinderMethod2",
		"api_key": api_key,
	})
}

func (t *TestHandler) HeaderBinderMethod3(c *gin.Context) {
	header := Header{}
	c.BindHeader(&header)

	c.JSON(http.StatusOK, gin.H{
		"message": "HeaderBinderMethod3",
		"header":  header,
	})
}

func (t *TestHandler) HeaderBinderMethod4(c *gin.Context) {
	header := Header{}
	c.ShouldBindHeader(&header)

	c.JSON(http.StatusOK, gin.H{
		"message": "HeaderBinderMethod4",
		"header":  header,
	})
}

// read from query
func (t *TestHandler) ReadQuery(c *gin.Context) {
	name := c.Query("name")   // 1 query param
	ids := c.QueryArray("id") // all query params of name "id"

	c.JSON(http.StatusOK, gin.H{
		"message": "ReadQuery",
		"name":    name,
		"ids":     ids,
	})
}

// read from root uri
func (t *TestHandler) UriBinder(c *gin.Context) {
	code := c.Param("code")
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"message": "RootUri",
		"code":    code,
		"id":      id,
	})
}

// read from body
func (t *TestHandler) BodyBinder(c *gin.Context) {
	person := Persons{}
	// c.Bind(&person)           // if err returns bad request
	// c.BindJSON(&person)       // if err returns error and should handle it
	err := c.ShouldBindJSON(&person) // if err returns error and should handle it

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "RootUri",
		"person":  person,
	})
}

// read from form
func (t *TestHandler) FormBinder(c *gin.Context) {
	person := Persons{}
	// c.Bind(&person)           // if err returns bad request
	c.ShouldBind(&person) // if err returns error and should handle it

	c.JSON(http.StatusOK, gin.H{
		"message": "RootUri",
		"person":  person,
	})
}

// read from file
func (t *TestHandler) FileBinder(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "FileBinder",
		"person":  file.Filename,
	})
}
