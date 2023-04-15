package sysModel

import (
	"errors"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/mysql"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/tools"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	gorm.Model
	UUID        uuid.UUID    `json:"uuid"`
	Username    string       `json:"username"`
	Password    string       `json:"-"`
	NickName    string       `json:"nickname"   gorm:"default:'User'"`
	HeaderImg   string       `json:"headerImg"  gorm:"defauly'http://www.henrongyi.top/avatar/lufu.jpg'"`
	Authority   SysAuthority `json:"authority"  gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId"`
	AuthorityId string       `json:"authorityId"gorm:"default:888"`
}

// 注册接口
func (u *SysUser) Regist() (err error, userInter *SysUser) {
	var user SysUser
	notResigt := mysql.DEFAULTDB.Where("username = ?", u.Username).First(&user).RecordNotFound()
	if !notResigt {
		//notresigt为false表明读取到了，不能注册
		return errors.New("用户名已注册"), nil
	} else {
		//附加uuid 密码md5简单加密 注册
		u.Password = tools.MD5V([]byte(u.Password))
		u.UUID = uuid.NewV4()
		err = mysql.DEFAULTDB.Create(u).Error
	}
	return err, u
}

// 修改用户密码
func (u *SysUser) ChangePassword(newPassword string) (err error, userinter *SysUser) {
	var user SysUser
	u.Password = tools.MD5V([]byte(u.Password))
	err = mysql.DEFAULTDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", tools.MD5V([]byte(newPassword))).Error
	return err, u
}

// 用户更新接口
func (u *SysUser) SetUserAuthority(uuid uuid.UUID, AuthorityId string) (err error) {
	err = mysql.DEFAULTDB.Where("uuid = ?", uuid).First(&SysUser{}).Update("authority_id", AuthorityId).Error
	return err
}

// 用户登录
func (u *SysUser) Login() (err error, userInter *SysUser) {
	var user SysUser
	u.Password = tools.MD5V([]byte(u.Password))
	err = mysql.DEFAULTDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	if err != nil {
		return nil, &user
	}
	err = mysql.DEFAULTDB.Where("authority_id = ?", user.AuthorityId).First(&user.Authority).Error
	return err, &user
}

// 用户头像上传更新地址
func (u *SysUser) UploadHeaderImg(uuid uuid.UUID, filePath string) (err error, userInter *SysUser) {
	var user SysUser
	err = mysql.DEFAULTDB.Where("uuid = ?", uuid).First(&user).Update("header_img", filePath).First(&user).Error
	return err, &user
}

// 分页获取数据
func (u *SysUser) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	//封装分页办法
	err, db, total := servers.PagingServer(u, info)
	if err != nil {
		return
	} else {
		var userList []SysUser
		err = db.Preload("Authority").Find(&userList).Error
		return err, userList, total
	}

}
