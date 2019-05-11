// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type SendingLog struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*SendingLog
	namer   func(string) string
	connID  int
	
	Id              	uint64  	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Created         	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	SentAt          	uint    	`db:"sent_at" bson:"sent_at" comment:"发送时间" json:"sent_at" xml:"sent_at"`
	SourceId        	uint64  	`db:"source_id" bson:"source_id" comment:"来源ID" json:"source_id" xml:"source_id"`
	SourceType      	string  	`db:"source_type" bson:"source_type" comment:"来源类型" json:"source_type" xml:"source_type"`
	Disabled        	string  	`db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	Method          	string  	`db:"method" bson:"method" comment:"发送方式(mobile-手机;email-邮箱)" json:"method" xml:"method"`
	To              	string  	`db:"to" bson:"to" comment:"发送目标" json:"to" xml:"to"`
	Provider        	string  	`db:"provider" bson:"provider" comment:"发送平台" json:"provider" xml:"provider"`
	Result          	string  	`db:"result" bson:"result" comment:"发送结果描述" json:"result" xml:"result"`
	Status          	string  	`db:"status" bson:"status" comment:"发送状态(none-无需发送)" json:"status" xml:"status"`
	Retries         	uint    	`db:"retries" bson:"retries" comment:"重试次数" json:"retries" xml:"retries"`
	Content         	string  	`db:"content" bson:"content" comment:"发送消息内容" json:"content" xml:"content"`
	Params          	string  	`db:"params" bson:"params" comment:"发送消息参数(JSON)" json:"params" xml:"params"`
	AppointmentTime 	uint    	`db:"appointment_time" bson:"appointment_time" comment:"预约发送时间" json:"appointment_time" xml:"appointment_time"`
}

func (this *SendingLog) Trans() *factory.Transaction {
	return this.trans
}

func (this *SendingLog) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *SendingLog) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *SendingLog) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *SendingLog) Objects() []*SendingLog {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *SendingLog) NewObjects() *[]*SendingLog {
	this.objects = []*SendingLog{}
	return &this.objects
}

func (this *SendingLog) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *SendingLog) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *SendingLog) Name_() string {
	if this.namer != nil {
		return this.namer("sending_log")
	}
	return factory.TableNamerGet("sending_log")(this)
}

func (this *SendingLog) FullName_(connID ...int) string {
	if len(connID) > 0 {
		return factory.DefaultFactory.Cluster(connID[0]).Table(this.Name_())
	}
	return factory.DefaultFactory.Cluster(this.connID).Table(this.Name_())
}

func (this *SendingLog) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *SendingLog) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *SendingLog) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *SendingLog) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *SendingLog) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *SendingLog) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Method) == 0 { this.Method = "mobile" }
	if len(this.SourceType) == 0 { this.SourceType = "user" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Status) == 0 { this.Status = "waiting" }
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	return
}

func (this *SendingLog) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	
	if len(this.Method) == 0 { this.Method = "mobile" }
	if len(this.SourceType) == 0 { this.SourceType = "user" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Status) == 0 { this.Status = "waiting" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *SendingLog) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *SendingLog) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *SendingLog) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["method"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["method"] = "mobile" } }
	if val, ok := kvset["source_type"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["source_type"] = "user" } }
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	if val, ok := kvset["status"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["status"] = "waiting" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *SendingLog) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		
	if len(this.Method) == 0 { this.Method = "mobile" }
	if len(this.SourceType) == 0 { this.SourceType = "user" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Status) == 0 { this.Status = "waiting" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Method) == 0 { this.Method = "mobile" }
	if len(this.SourceType) == 0 { this.SourceType = "user" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	if len(this.Status) == 0 { this.Status = "waiting" }
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	return 
}

func (this *SendingLog) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *SendingLog) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *SendingLog) Reset() *SendingLog {
	this.Id = 0
	this.Created = 0
	this.SentAt = 0
	this.SourceId = 0
	this.SourceType = ``
	this.Disabled = ``
	this.Method = ``
	this.To = ``
	this.Provider = ``
	this.Result = ``
	this.Status = ``
	this.Retries = 0
	this.Content = ``
	this.Params = ``
	this.AppointmentTime = 0
	return this
}

func (this *SendingLog) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Created"] = this.Created
	r["SentAt"] = this.SentAt
	r["SourceId"] = this.SourceId
	r["SourceType"] = this.SourceType
	r["Disabled"] = this.Disabled
	r["Method"] = this.Method
	r["To"] = this.To
	r["Provider"] = this.Provider
	r["Result"] = this.Result
	r["Status"] = this.Status
	r["Retries"] = this.Retries
	r["Content"] = this.Content
	r["Params"] = this.Params
	r["AppointmentTime"] = this.AppointmentTime
	return r
}

