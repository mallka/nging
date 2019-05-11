// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type SshUser struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*SshUser
	namer   func(string) string
	connID  int
	
	Id         	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Uid        	uint    	`db:"uid" bson:"uid" comment:"UID" json:"uid" xml:"uid"`
	Host       	string  	`db:"host" bson:"host" comment:"主机名" json:"host" xml:"host"`
	Port       	int     	`db:"port" bson:"port" comment:"端口" json:"port" xml:"port"`
	Username   	string  	`db:"username" bson:"username" comment:"用户名" json:"username" xml:"username"`
	Password   	string  	`db:"password" bson:"password" comment:"密码" json:"password" xml:"password"`
	Name       	string  	`db:"name" bson:"name" comment:"账号名称" json:"name" xml:"name"`
	Options    	string  	`db:"options" bson:"options" comment:"其它选项(JSON)" json:"options" xml:"options"`
	PrivateKey 	string  	`db:"private_key" bson:"private_key" comment:"私钥内容" json:"private_key" xml:"private_key"`
	Passphrase 	string  	`db:"passphrase" bson:"passphrase" comment:"私钥口令" json:"passphrase" xml:"passphrase"`
	Protocol   	string  	`db:"protocol" bson:"protocol" comment:"连接协议" json:"protocol" xml:"protocol"`
	Description	string  	`db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	GroupId    	uint    	`db:"group_id" bson:"group_id" comment:"组" json:"group_id" xml:"group_id"`
	Created    	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated    	uint    	`db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
}

func (this *SshUser) Trans() *factory.Transaction {
	return this.trans
}

func (this *SshUser) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *SshUser) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *SshUser) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *SshUser) Objects() []*SshUser {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *SshUser) NewObjects() *[]*SshUser {
	this.objects = []*SshUser{}
	return &this.objects
}

func (this *SshUser) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *SshUser) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *SshUser) Name_() string {
	if this.namer != nil {
		return this.namer("ssh_user")
	}
	return factory.TableNamerGet("ssh_user")(this)
}

func (this *SshUser) FullName_(connID ...int) string {
	if len(connID) > 0 {
		return factory.DefaultFactory.Cluster(connID[0]).Table(this.Name_())
	}
	return factory.DefaultFactory.Cluster(this.connID).Table(this.Name_())
}

func (this *SshUser) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *SshUser) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *SshUser) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *SshUser) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *SshUser) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *SshUser) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Host) == 0 { this.Host = "localhost" }
	if len(this.Username) == 0 { this.Username = "root" }
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

func (this *SshUser) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	this.Updated = uint(time.Now().Unix())
	if len(this.Host) == 0 { this.Host = "localhost" }
	if len(this.Username) == 0 { this.Username = "root" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *SshUser) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *SshUser) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *SshUser) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["host"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["host"] = "localhost" } }
	if val, ok := kvset["username"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["username"] = "root" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *SshUser) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		this.Updated = uint(time.Now().Unix())
	if len(this.Host) == 0 { this.Host = "localhost" }
	if len(this.Username) == 0 { this.Username = "root" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Host) == 0 { this.Host = "localhost" }
	if len(this.Username) == 0 { this.Username = "root" }
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

func (this *SshUser) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *SshUser) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *SshUser) Reset() *SshUser {
	this.Id = 0
	this.Uid = 0
	this.Host = ``
	this.Port = 0
	this.Username = ``
	this.Password = ``
	this.Name = ``
	this.Options = ``
	this.PrivateKey = ``
	this.Passphrase = ``
	this.Protocol = ``
	this.Description = ``
	this.GroupId = 0
	this.Created = 0
	this.Updated = 0
	return this
}

func (this *SshUser) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Uid"] = this.Uid
	r["Host"] = this.Host
	r["Port"] = this.Port
	r["Username"] = this.Username
	r["Password"] = this.Password
	r["Name"] = this.Name
	r["Options"] = this.Options
	r["PrivateKey"] = this.PrivateKey
	r["Passphrase"] = this.Passphrase
	r["Protocol"] = this.Protocol
	r["Description"] = this.Description
	r["GroupId"] = this.GroupId
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	return r
}

