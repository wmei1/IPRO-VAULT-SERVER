package controllers

import (
    "encoding/json"
    "io/ioutil"
    "fmt"
    "log"
    "github.com/revel/revel"
    "github.com/skrzepto/IPRO-VAULT-SERVER/app/models"
    // "github.com/skrzepto/IPRO-VAULT-SERVER/app/routes"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Test() revel.Result {
    return c.RenderText("Test")
}

// Catch the json and store
func (c App) NewData() revel.Result {
    var s models.SensorData
    content, _ := ioutil.ReadAll(c.Request.Body)
    err := json.Unmarshal([]byte(content), &s)
    // fmt.Println(s.name) //sensor_data

    // var content []string
    // err := json.NewDecoder(c.Request.Body).Decode(&content)
    if err != nil {
        log.Fatal("JSON decode error: ", err)
    }
    defer c.Request.Body.Close()
    fmt.Println("%+v", s)
    return c.RenderText("%+v", s)
}
