package models

import "github.com/beego/beego/orm"

type User struct {
	Id           uint `orm:"pk"`
	Name         string
	Password     string
	Viewposition uint64
}

func init() {
	orm.RegisterModel(new(User))
}
func (u *User) Insert() string {
	o := orm.NewOrm()
	_, err := o.Insert(u)
	if err != nil {
		return err.Error()
	}
	return ""

}

func (u *User) Delete() string {
	o := orm.NewOrm()
	_, err := o.Delete(u)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (u *User) Select() string {
	o := orm.NewOrm()
	err := o.Read(u, "Name")
	if err != nil {
		return err.Error()
	}
	if u.Id == 0 {
		return "0"
	}
	return ""
}

func (u *User) Update() string {
	o := orm.NewOrm()
	_, err := o.Update(u)
	if err != nil {
		return err.Error()
	}
	return ""
}
func (u *User) UpdateViewposition() string {
	o := orm.NewOrm()
	_, err := o.Update(u, "Viewposition")
	if err != nil {
		return err.Error()
	}
	return ""
}
