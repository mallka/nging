// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type CodeInvitation struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*CodeInvitation
	namer   func(string) string
	connID  int
	
	Id      	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Uid     	uint    	`db:"uid" bson:"uid" comment:"创建者" json:"uid" xml:"uid"`
	RecvUid 	uint    	`db:"recv_uid" bson:"recv_uid" comment:"使用者" json:"recv_uid" xml:"recv_uid"`
	Code    	string  	`db:"code" bson:"code" comment:"邀请码" json:"code" xml:"code"`
	Created 	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Used    	uint    	`db:"used" bson:"used" comment:"使用时间" json:"used" xml:"used"`
	Start   	uint    	`db:"start" bson:"start" comment:"有效时间" json:"start" xml:"start"`
	End     	uint    	`db:"end" bson:"end" comment:"失效时间" json:"end" xml:"end"`
	Disabled	string  	`db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	RoleIds 	string  	`db:"role_ids" bson:"role_ids" comment:"注册为角色(多个用“,”分隔开)" json:"role_ids" xml:"role_ids"`
}

func (this *CodeInvitation) Trans() *factory.Transaction {
	return this.trans
}

func (this *CodeInvitation) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *CodeInvitation) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *CodeInvitation) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *CodeInvitation) Objects() []*CodeInvitation {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *CodeInvitation) NewObjects() *[]*CodeInvitation {
	this.objects = []*CodeInvitation{}
	return &this.objects
}

func (this *CodeInvitation) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *CodeInvitation) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *CodeInvitation) Name_() string {
	if this.namer != nil {
		return this.namer("code_invitation")
	}
	return factory.TableNamerGet("code_invitation")(this)
}

func (this *CodeInvitation) FullName_(connID ...int) string {
	if len(connID) > 0 {
		return factory.DefaultFactory.Cluster(connID[0]).Table(this.Name_())
	}
	return factory.DefaultFactory.Cluster(this.connID).Table(this.Name_())
}

func (this *CodeInvitation) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *CodeInvitation) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *CodeInvitation) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *CodeInvitation) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CodeInvitation) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CodeInvitation) Add() (pk interface{}, err error) {
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

func (this *CodeInvitation) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *CodeInvitation) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *CodeInvitation) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *CodeInvitation) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *CodeInvitation) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		
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

func (this *CodeInvitation) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *CodeInvitation) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *CodeInvitation) Reset() *CodeInvitation {
	this.Id = 0
	this.Uid = 0
	this.RecvUid = 0
	this.Code = ``
	this.Created = 0
	this.Used = 0
	this.Start = 0
	this.End = 0
	this.Disabled = ``
	this.RoleIds = ``
	return this
}

func (this *CodeInvitation) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Uid"] = this.Uid
	r["RecvUid"] = this.RecvUid
	r["Code"] = this.Code
	r["Created"] = this.Created
	r["Used"] = this.Used
	r["Start"] = this.Start
	r["End"] = this.End
	r["Disabled"] = this.Disabled
	r["RoleIds"] = this.RoleIds
	return r
}

