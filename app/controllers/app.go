package controllers

import (
    "encoding/json"
    "io/ioutil"
    "fmt"
    // "log"
    "github.com/revel/revel"
    // "github.com/skrzepto/IPRO-VAULT-SERVER/app/models"
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
    content, _ := ioutil.ReadAll(c.Request.Body)
    var s map[string]interface{}
    if err := json.Unmarshal(content, &s); err != nil {
        panic(err)
    }
    fmt.Printf("decoded to %#v\n", s)

    defer c.Request.Body.Close()

    // result := &models.SensorData{}
    // err := result.FillStruct(s)

    // if err != nil {
    //     fmt.Println(err)
    // }

    // fmt.Println(result)
    // fmt.Println("%+v", s)
    return c.RenderText("%+v", s)
}
