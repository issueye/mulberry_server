package logic

import (
	"errors"
	"mulberry/app/admin/model"
	"mulberry/app/admin/requests"
	"mulberry/app/admin/service"
	commonModel "mulberry/common/model"
	"mulberry/global"
)

// 创建数据
func CreateRole(r *requests.CreateRole) error {

	srv := service.NewRole()

	data, err := srv.GetByField("code", r.Code)
	if err != nil {
		return err
	}

	if data.ID != 0 {
		return errors.New("角色编码已存在")
	}

	info := &model.Role{
		Code:   r.Code,
		Name:   r.Name,
		Remark: r.Remark,
	}

	return service.NewRole().Create(info)
}

// 更新数据
func UpdateRole(r *requests.UpdateRole) error {
	data := make(map[string]any)
	data["code"] = r.Code
	data["name"] = r.Name
	data["remark"] = r.Remark

	return service.NewRole().UpdateByMap(uint(r.Id), data)
}

// 根据ID查询数据
func GetRoleById(id uint) (*model.Role, error) {
	return service.NewRole().GetById(id)
}

// 根据条件查询数据
func ListRole(condition *commonModel.PageQuery[*requests.QueryRole]) (*commonModel.ResPage[model.Role], error) {
	return service.NewRole().ListRole(condition)
}

// 删除数据
func DeleteRole(id uint) error {
	return service.NewRole().Delete(id)
}

func InitRoles() {
	Roles := []*model.Role{
		model.NewRole("9001", "管理员"),
	}

	for _, Role := range Roles {
		RoleIsNotExistAdd(Role)
	}
}

func RoleIsNotExistAdd(Role *model.Role) {
	RoleSrv := service.NewUser()
	info, err := RoleSrv.GetRoleByName(Role.Name)
	if err != nil {
		global.Logger.Sugar().Errorf("查询角色失败，失败原因：%s", err.Error())
		return
	}

	if info.ID == 0 {
		err = RoleSrv.AddRole(Role)
		if err != nil {
			global.Logger.Sugar().Errorf("添加角色失败，失败原因：%s", err.Error())
			return
		}
	}
}
