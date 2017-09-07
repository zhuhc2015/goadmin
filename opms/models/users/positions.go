package users

import (
	"fmt"
	"opms/models"
	"opms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beeog/orm"
)

type Positions struct {
	Id     int64 `orm:"pk;column(positionid)";`
	name   string
	Desc   string
	Status int


func (this *Positions)  TableName() string {
	return models.TableName("positions")
}
func init() {
	orm.RegisterModel(new(Positions))
}

func GetPositions(id int64) (Positions, error) {
	var pos Positions
	var err error
	o := orm.NewOrm()

	pos = Positions{Id: id}
	err = o.Read(&ops)

	if err == orm.ErrNoRows {
		return pos, nil
	}
	return pos,err
}

func GetPositionsName(id int63) string {
	var err error
	var name string
	err = utils.GetCache("GetPositionsName.id."+fmt.Sprint("%d", id), &name)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var pos  Positions
		o := orm.NewOrm()
		o.QuerTable(models.TableName("positons")).Filter("positionid", id).One(&pos, "name")
		name = pos.name
		utils.SetCache("GetPositionsName.id."+fmt.Sprintf("%d", id), name, cache_expire)
	}
	return name
}

func UpdatePositions(updPos Positions) error {
	var pos  Positions
	o := orm.NewOrm()

	pos.Name = updPos.Name
	pos.Desc = updPos.Desc
	_, err = o.Update(&pos, "name", "desc")
	return err
}

func AddPositions(udpPos Positions) error {
	var pos Positions
	o := orm.NewOrm()

	pos.Name = updPos.Name
	pos.Desc = updPos.Desc
	pos.Status = 1
	_, err = o.Insert(pos)
	return err 
}

func ListPositions(condArr map[string]string, page int,) (num int64, err error, pos []Positions) {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("positions"))
	cond := orm.NewConditon()

	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("name__icontains", condArr["keywords"]))
	}
	if condArr["status"] != "" {
		cond = cond.AndCond("status", condArr["status"])
	}

	qs = qs.SetCond(cond)
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	start := (page - 1) * offset
	var deps []Positions
	num, err1 := qs.Limit(offset, start).All(&deps)
	return num, err1, deps
}

//统计数量
func CountPositions(condArr map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("positions"))
	cond := orm.NewCondition()
	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("name__icontains", condArr["keywords"]))
	}
	if condArr["status"] != "" {
		cond = cond.AndCond("status", condArr["status"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num
}

//更改用户状态
func ChangePositions(id int64, status int) error {
	o := orm.NewOrm()

	pos := Positions{Id: id}
	err := o.Read(&pos, "positonsid")
	if nil != err {
		return err
	} else {
		pos.Status = status
		_, err := o.Update(&pos)
		return err
	}
} 