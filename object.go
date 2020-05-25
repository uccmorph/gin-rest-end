package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	API_GET_1 = iota
	API_GET_2
	API_POST_1
)

type Resource struct {
	tags map[string]string
}

func NewResource() *Resource {
	p := &Resource{
		map[string]string{},
	}
	return p
}

func (p *Resource) BuildHandler(api int) (string, gin.HandlerFunc) {
	switch api {
	case API_GET_1:
		return "/object/:name", p.apiGet1
	case API_GET_2:
		return "/api/test", nil
	case API_POST_1:
		return "object/:name", p.apiPost1
	}
	return "", nil
}

func (p *Resource) apiGet1(c *gin.Context) {
	name := c.Param("name")
	if tag, ok := p.tags[name]; ok {
		c.JSON(200, gin.H{
			"message": tag,
		})
	}
	c.JSON(404, nil)
}

type PostData struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

func (p *Resource) apiPost1(c *gin.Context) {
	name := c.Param("name")
	if c.ContentType() != "json" && c.ContentType() != "application/json" {
		log.Printf("error content type: %s", c.ContentType())
		c.String(400, "incorrect Content-Type")
		return
	}
	data := PostData{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.String(500, fmt.Sprintf("parse post failed: %s", err.Error()))
		return
	}
	if name != data.Name {
		c.String(400, "invalid url path")
		return
	}
	p.tags[data.Name] = data.Tag
	c.Status(204)

}
