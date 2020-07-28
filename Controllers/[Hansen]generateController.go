package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	os.Mkdir(pwd+`\controller\`, os.ModePerm)

	f, err := os.Create(pwd + `\controller\test.go`)
	if err != nil {
		fmt.Println(err)
		return
	}

	var input string
	fmt.Scanln(&input)

	input = strings.Title(input)

	d := []string{
		// Package
		"package controller\n\n",

		// Import
		"import (",
		"\n\t", `models "crud-api/models"`,
		"\n\t", `"fmt"`,
		"\n\t", `"net/http"`,
		"\n\n\t", `"github.com/gin-gonic/gin"`,
		"\n)\n\n",

		//Functions
		"func Get" + input + "(c *gin.Context) {",
		"\n\tvar " + input + " []models." + input + "",
		"\n\terr := models.GetAll" + input + "s(&" + input + ")",
		"\n\tif err != nil {",
		"\n\t\tc.AbortWithStatus(http.StatusNotFound)",
		"\n\t} else {",
		"\n\t\tc.JSON(http.StatusOK, " + input + ")",
		"\n\t}",
		"\n}\n\n",

		"func Create" + input + "(c *gin.Context) {",
		"\n\tvar " + input + " models." + input + "",
		"\n\tc.BindJSON(&" + input + ")",
		"\n\terr := models.Create" + input + "(&" + input + ")",
		"\n\tif err != nil {",
		"\n\t\tfmt.Println(err.Error())",
		"\n\t\tc.AbortWithStatus(http.StatusNotFound)",
		"\n\t} else {",
		"\n\t\tc.JSON(http.StatusOK, " + input + ")",
		"\n\t}",
		"\n}\n\n",

		"func Get" + input + "ByUsername(c *gin.Context) {",
		"\n\t", `username := c.Params.ByName("username")`,
		"\n\tvar " + input + " models." + input + "",
		"\n\terr := models.Get" + input + "ByUsername(&" + input + ", username)",
		"\n\tif err != nil {",
		"\n\t\tc.AbortWithStatus(http.StatusNotFound)",
		"\n\t} else {",
		"\n\t\tc.JSON(http.StatusOK, " + input + ")",
		"\n\t}",
		"\n}\n\n",

		"func Update" + input + "(c *gin.Context) {",
		"\n\tvar " + input + " models." + input + "",
		"\n\t", `username := c.Params.ByName("username")`,
		"\n\terr := models.Get" + input + "ByUsername(&" + input + ", username)",
		"\n\tif err != nil {",
		"\n\tc.JSON(http.StatusNotFound, " + input + ")",
		"\n\t}",
		"\n\tc.BindJSON(&" + input + ")",
		"\n\terr = models.Update" + input + "(&" + input + ", username)",
		"\n\tif err != nil {",
		"\n\tc.AbortWithStatus(http.StatusNotFound)",
		"\n\t} else {",
		"\n\tc.JSON(http.StatusOK, " + input + ")",
		"\n\t}",
		"\n}\n\n",

		"func Delete" + input + "(c *gin.Context) {",
		"\n\tvar " + input + " models." + input + "",
		"\n\t", `username := c.Params.ByName("username")`,
		"\n\terr := models.Delete" + input + "(&" + input + ", username)",
		"\n\tif err != nil {",
		"\n\t\tc.AbortWithStatus(http.StatusNotFound)",
		"\n\t} else {",
		"\n\t\t", `c.JSON(http.StatusOK, gin.H{"username" + username: "is deleted"}`,
		"\n\t}",
		"\n}\n\n",
	}

	for _, v := range d {
		fmt.Fprint(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}
