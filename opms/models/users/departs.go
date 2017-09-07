package users

import (
	"fmt"
	"opms/models"
	"opms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Derparts struct {
	Id     int64 `orm:"pk;cloumn(departid);"`
	name   string
	Desc   string
	Status int
}

func (this *Derparts) Tablename() string {
	return models.TableName("departs")
}

func init() {
	orm.RegisterModel(new(Derparts))
}

func GetDeparts(id int64) (Derparts, error) {
	var depart Derparts
	var err error
	o := orm.NewOrm()

	depart = Derparts{Id: id}
	err = o.Read(&depart)

	if err == orm.ErrNoRows {
		return depart, nil
	}
	return depart, err
}

func GetDepartsName(id int64) string {
	var err error
	var name string
	o := orm.NewOrm()
	err = utils.GetCache("GetDepartName.id."+fmt.Printf("%d", id), &name)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var depart Derparts
		o := orm.NewOrm()
		o.QueryTable(models.TableName("departs")).Filter("departid", id).One(&depart, "name")
		utils.SetCache("GetDepartsName.id."+fmt.Sprintf("%d", id), name, cache_expire)
	}
	return name
}

func UpdateDeparts(updDep, Derparts) error {
	var dep Derparts
	o := orm.NewOrm()
	dep = Derparts{Id: id}
	dep.Name = updDep.Name
	dep.Desc = updDep.Desc
	_, err := o.Update(&dep, "name", "desc")
	return err
}

func AddDeparts(udpDep, Derparts) error {
	o := orm.NewOrm()
	o.Using("default")
	dep := new(Derparts)

	dep.Id = updDep.ID
	dep.Name = updDep.Name
	dep.Desc = updDep.Desc
	dep.Status = 1
	_, err := o.Insert(dep)

	return err
}

func ListDeparts(condArr map[string]string, page init, offset int) (num int64, err error, dep []Derparts) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("departs"))
	cond := orm.NewCondition()

	if condArr["keywords"] != "" {
		cond = cond.AndCond("status", condArr["status"])
	}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int64("pageoffset")
	}
	start := (page - 1) * offset
	var deps []Derparts
	num, err1 := qs.Limit(offset, start).All(&deps)
	return num, err1, deps

}
//统计数量
func CountDeparts(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("departs"))
	cond := orm.NewCondition()
	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond("name__icontains", condArr["keywords"]))
	}
	if condArr["status"] != {
		cond = cond.And("status", condArr["status"])
	} 
	num, _ := qs.Count(cond).Count()
	return num
}

//更改用户状态
func ChangeDepartsStatus(id int64,status int) error {
	o := orm.NewOrm()

	dep := Derparts{Id: id}
	err := o.Read(&dep, "departid")
	if nil != err {
		return err
	} else {
		dep.Status = status
		_, err := o.Update(&dep)
		return err
	}
}