package models

import (
	"github.com/beego/beego/orm"
	"github.com/beego/beego/v2/core/logs"
)

type Collect struct {
	Id      uint64 `orm:"pk"`
	Userid  uint   `orm:"index"`
	Videoid uint64 `orm:"index"`
}

func init() {
	orm.RegisterModel(new(Collect))
}
func (c *Collect) SelectByUseridAndVideoid() string {
	o := orm.NewOrm()
	qs := o.QueryTable(c)
	err := qs.Filter("Userid", c.Userid).Filter("Videoid", c.Videoid).Limit(1).One(c)
	if err != nil {
		return err.Error()
		// Not exists in table
	} else if c.Id == 0 {
		return "0"
	}
	return ""
}
func (c *Collect) Insert() string {
	o := orm.NewOrm()
	_, err := o.Insert(c)
	if err != nil {
		return err.Error()
	}
	return ""
}

// select * from collect where userid=1 and id>11 limit 1;
func (c *Collect) SelectNextByUserIdAndId() string {
	o := orm.NewOrm()
	qs := o.QueryTable(c)
	err := qs.Filter("Userid", c.Userid).Filter("Id__gt", c.Id).Limit(1).One(c)
	if err != nil {
		logs.Info(err.Error())
		return err.Error()
	}
	return ""
}
func (c *Collect) DeleteByUserId() string {
	o := orm.NewOrm()
	_, err := o.Delete(c, "Userid")
	if err != nil {
		return err.Error()
	}
	return ""
}
func (c *Collect) DeleteByVideoid() string {
	o := orm.NewOrm()
	_, err := o.Delete(c, "Videoid")
	if err != nil {
		return err.Error()
	}
	return ""
}
func (c *Collect) DeleteByUserIdAndVideoid() string {
	o := orm.NewOrm()
	_, err := o.Delete(c, "Videoid", "Userid")
	if err != nil {
		return err.Error()
	}
	return ""
}
func (c *Collect) DeleteById() string {
	o := orm.NewOrm()
	_, err := o.Delete(c)
	if err != nil {
		return err.Error()
	}
	return ""
}
func (c *Collect) DeleteByUseridAndId() string {
	o := orm.NewOrm()
	_, err := o.Delete(c, "Userid", "Id")
	if err != nil {
		return err.Error()
	}
	return ""
}
