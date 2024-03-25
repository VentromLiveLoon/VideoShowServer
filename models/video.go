package models

import "github.com/beego/beego/orm"

type VideoPath struct {
	Id   uint64 `orm:"pk"`
	Name string
}

func init() {
	orm.RegisterModel(new(VideoPath))
}

func (vp *VideoPath) Insert() string {
	o := orm.NewOrm()
	_, err := o.Insert(vp)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (vp *VideoPath) SelectNext() string {
	// ip=select * from  VideoPath where id>1 limit 1;
	o := orm.NewOrm()
	qs := o.QueryTable(vp)
	//
	err := qs.Filter("Id__gt", vp.Id).Limit(1).One(vp)
	if err != nil {
		return err.Error()
	}
	return ""
}
func (vp *VideoPath) Select() string {
	o := orm.NewOrm()
	err := o.Read(vp)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (vp *VideoPath) SelectByName() string {
	o := orm.NewOrm()
	err := o.Read(vp, "Name")
	if err != nil {
		return err.Error()
	}
	return ""
}
func (vp *VideoPath) Delete() string {
	o := orm.NewOrm()
	_, err := o.Delete(vp, "Id")
	if err != nil {
		s := err.Error()
		return s
	}
	return ""
}
