package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

// App struct
type App struct {
	ctx context.Context
}

type Item struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

var myList []Item

func init() {
	myList = make([]Item, 0)
	fp, err := os.OpenFile("./data.json", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	bytes, err := ioutil.ReadAll(fp)
	if err != nil {
		panic(err)
	}
	if len(bytes) == 0 {
		return
	}
	err = json.Unmarshal(bytes, &myList)
	if err != nil {
		panic(err)
	}
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetList() []Item {
	return myList
}

func (a *App) Add(item Item) string {
	for _, v := range myList {
		if v.Label == item.Label {
			return "快捷方式已存在"
		}
	}
	myList = append(myList, item)
	err := save()
	if err != nil {
		return err.Error()
	}
	return ""
}

func (a *App) Del(name string) string {
	i := 0
	for _, v := range myList {
		if v.Label != name {
			myList[i] = v
			i++
		}
	}
	myList = myList[:i]
	err := save()
	if err != nil {
		return err.Error()
	}
	return ""
}

func (a *App) Edit(item Item) string {
	for i := 0; i < len(myList); i++ {
		if myList[i].Label == item.Label {
			myList[i].Value = item.Value
			break
		}
	}
	err := save()
	if err != nil {
		return err.Error()
	}
	return ""
}

func (a *App) Open(name string) string {
	for _, v := range myList {
		if v.Label == name {
			defer func() {
				_ = recover()
			}()
			cmd := exec.Command("cmd", "/C", "start", v.Value)
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			err := cmd.Run()
			if err != nil {
				return err.Error()
			}
		}
	}
	return ""
}

func save() error {
	data, err := json.Marshal(myList)
	if err != nil {
		return err
	}
	err = os.WriteFile("./data.json", data, 0777)
	if err != nil {
		return err
	}
	return nil
}
