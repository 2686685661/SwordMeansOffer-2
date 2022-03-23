/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/3/15 11:28 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package service_registry

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Registry struct {
	apps map[string]*Application
	lock sync.RWMutex
}

type Application struct {
	appid string
	instances map[string]*Instance
	latestTimestamp int64
	lock sync.RWMutex
}


type RequestRegister Instance

type Instance struct {
	Env      string   `json:"env"`
	AppId    string   `json:"appid"`
	Hostname string   `json:"hostname"`
	Addrs    []string `json:"addrs"`
	Version  string   `json:"version"`
	Status   uint32   `json:"status"`

	RegTimestamp    int64 `json:"reg_timestamp"`
	UpTimestamp     int64 `json:"up_timestamp"`
	RenewTimestamp  int64 `json:"renew_timestamp"`
	DirtyTimestamp  int64 `json:"dirty_timestamp"`
	LatestTimestamp int64 `json:"latest_timestamp"`
}


func NewRegistry() *Registry {
	registry := &Registry{
		apps: make(map[string]*Application),
	}
	return registry
}

func NewInstance(req *RequestRegister) *Instance {
	now := time.Now().UnixNano()
	instance := &Instance{
		Env:             req.Env,
		AppId:           req.AppId,
		Hostname:        req.Hostname,
		Addrs:           req.Addrs,
		Version:         req.Version,
		Status:          req.Status,
		RegTimestamp:    now,
		UpTimestamp:     now,
		RenewTimestamp:  now,
		DirtyTimestamp:  now,
		LatestTimestamp: now,
	}
	return instance
}


func (r *Registry) Register(instance *Instance, latestTimestamp int64) (*Application, error) {
	key := fmt.Sprintf("%s_%s", instance.AppId, instance.Env)
	r.lock.RLock()
	app, ok := r.apps[key]
	r.lock.RUnlock()
	if !ok { //new app
		app = NewApplication(instance.AppId)
	}
	//add instance
	_, isNew := app.AddInstance(instance, latestTimestamp)
	if isNew {
		//add into registry apps
		r.lock.Lock()
		r.apps[key] = app
		r.lock.Unlock()
		return app, nil
	}
	return app, nil
}

func NewApplication(appid string) *Application {
	return &Application{
		appid:     appid,
		instances: make(map[string]*Instance),
	}
}


func (app *Application) AddInstance(in *Instance, latestTimestamp int64) (*Instance, bool) {
	app.lock.Lock()
	defer app.lock.Unlock()
	appIns, ok := app.instances[in.Hostname]
	if ok { //exist
		in.UpTimestamp = appIns.UpTimestamp
		//dirtytimestamp
		if in.DirtyTimestamp < appIns.DirtyTimestamp {
			log.Println("register exist dirty timestamp")
			in = appIns
		}
	}
	//add or update instances
	app.instances[in.Hostname] = in
	app.latestTimestamp = latestTimestamp
	returnIns := new(Instance)
	*returnIns = *in
	return returnIns, !ok
}