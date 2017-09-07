package users

import (
	"fmt"
	"opms/models"
	"opms/utils"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserPermissions struct {
	Id          int64 `orm:"pk;column(userid);"`
	Permissions string
	Model       string
	Modelc      string
}

func (this *UserPermissions) TableName() string {
	return models.RegisterModel(new(UserPermissions))
}

func GetPermission(id int64) string {
	var err error
	var name string
	err = utils.GetCache("GetPermissionsName.id."+fmt.Sprintf("%d", id), &name)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		var permission UserPermissions
		o := orm.NewOrm()
		o.QueryTable(models.TableName("users_permissions")).Filter("userid, id").One(&permission, &permission)
		name = permission.Permissions
		utils.SetCache("GetPermission.id."+fmt.Sprintf("%d", id), name, cache_expire)
	}
	return name
}
