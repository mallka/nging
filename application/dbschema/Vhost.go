// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type Vhost struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*Vhost
	namer   func(string) string
	connID  int
	
	Id      	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name    	string  	`db:"name" bson:"name" comment:"网站名称" json:"name" xml:"name"`
	GroupId 	uint    	`db:"group_id" bson:"group_id" comment:"组" json:"group_id" xml:"group_id"`
	Domain  	string  	`db:"domain" bson:"domain" comment:"域名" json:"domain" xml:"domain"`
	Root    	string  	`db:"root" bson:"root" comment:"网站物理路径" json:"root" xml:"root"`
	Created 	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated 	uint    	`db:"updated" bson:"updated" comment:"更新时间" json:"updated" xml:"updated"`
	Setting 	string  	`db:"setting" bson:"setting" comment:"设置" json:"setting" xml:"setting"`
	Disabled	string  	`db:"disabled" bson:"disabled" comment:"是否停用" json:"disabled" xml:"disabled"`
}

func (this *Vhost) Trans() *factory.Transaction {
	return this.trans
}

func (this *Vhost) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *Vhost) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *Vhost) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *Vhost) Objects() []*Vhost {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *Vhost) NewObjects() *[]*Vhost {
	this.objects = []*Vhost{}
	return &this.objects
}

func (this *Vhost) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *Vhost) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *Vhost) Name_() string {
	if this.namer != nil {
		return this.namer("vhost")
	}
	return factory.TableNamerGet("vhost")(this)
}

func (this *Vhost) FullName_(connID ...int) string {
	if len(connID) > 0 {
		return factory.DefaultFactory.Cluster(connID[0]).Table(this.Name_())
	}
	return factory.DefaultFactory.Cluster(this.connID).Table(this.Name_())
}

func (this *Vhost) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *Vhost) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *Vhost) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *Vhost) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *Vhost) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *Vhost) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return
}

func (this *Vhost) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	this.Updated = uint(time.Now().Unix())
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *Vhost) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *Vhost) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *Vhost) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *Vhost) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		this.Updated = uint(time.Now().Unix())
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return 
}

func (this *Vhost) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *Vhost) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *Vhost) Reset() *Vhost {
	this.Id = 0
	this.Name = ``
	this.GroupId = 0
	this.Domain = ``
	this.Root = ``
	this.Created = 0
	this.Updated = 0
	this.Setting = ``
	this.Disabled = ``
	return this
}

func (this *Vhost) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Name"] = this.Name
	r["GroupId"] = this.GroupId
	r["Domain"] = this.Domain
	r["Root"] = this.Root
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	r["Setting"] = this.Setting
	r["Disabled"] = this.Disabled
	return r
}

