package mysqlModel

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UsersMgr struct {
	*_BaseMgr
}

// UsersMgr open func
func UsersMgr(db *gorm.DB) *_UsersMgr {
	if db == nil {
		panic(fmt.Errorf("UsersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("users"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UsersMgr) GetTableName() string {
	return "users"
}

// Reset 重置gorm会话
func (obj *_UsersMgr) Reset() *_UsersMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UsersMgr) Get() (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UsersMgr) Gets() (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UsersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Users{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAutoID AutoID获取 自增主键
func (obj *_UsersMgr) WithAutoID(autoID int64) Option {
	return optionFunc(func(o *options) { o.query["AutoID"] = autoID })
}

// WithUserID UserID获取 用户id
func (obj *_UsersMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["UserID"] = userID })
}

// WithUserName UserName获取 用户名
func (obj *_UsersMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["UserName"] = userName })
}

// WithPassword Password获取 用户密码
func (obj *_UsersMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["Password"] = password })
}

// WithEmail Email获取 用户邮箱
func (obj *_UsersMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["Email"] = email })
}

// WithIsAdmin IsAdmin获取 是否是管理员
func (obj *_UsersMgr) WithIsAdmin(isAdmin bool) Option {
	return optionFunc(func(o *options) { o.query["IsAdmin"] = isAdmin })
}

// WithPassNum PassNum获取 通过题目数量
func (obj *_UsersMgr) WithPassNum(passNum int64) Option {
	return optionFunc(func(o *options) { o.query["PassNum"] = passNum })
}

// WithSubmitNum SubmitNum获取 提交次数
func (obj *_UsersMgr) WithSubmitNum(submitNum int64) Option {
	return optionFunc(func(o *options) { o.query["SubmitNum"] = submitNum })
}

// WithCreateTime CreateTime获取 创建时间
func (obj *_UsersMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["CreateTime"] = createTime })
}

// WithIsDeleted IsDeleted获取 是否删除
func (obj *_UsersMgr) WithIsDeleted(isDeleted int8) Option {
	return optionFunc(func(o *options) { o.query["IsDeleted"] = isDeleted })
}

// GetByOption 功能选项模式获取
func (obj *_UsersMgr) GetByOption(opts ...Option) (result Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UsersMgr) GetByOptions(opts ...Option) (results []*Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAutoID 通过AutoID获取内容 自增主键
func (obj *_UsersMgr) GetFromAutoID(autoID int64) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}

// GetBatchFromAutoID 批量查找 自增主键
func (obj *_UsersMgr) GetBatchFromAutoID(autoIDs []int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`AutoID` IN (?)", autoIDs).Find(&results).Error

	return
}

// GetFromUserID 通过UserID获取内容 用户id
func (obj *_UsersMgr) GetFromUserID(userID string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`UserID` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户id
func (obj *_UsersMgr) GetBatchFromUserID(userIDs []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`UserID` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromUserName 通过UserName获取内容 用户名
func (obj *_UsersMgr) GetFromUserName(userName string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`UserName` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找 用户名
func (obj *_UsersMgr) GetBatchFromUserName(userNames []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`UserName` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromPassword 通过Password获取内容 用户密码
func (obj *_UsersMgr) GetFromPassword(password string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`Password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 用户密码
func (obj *_UsersMgr) GetBatchFromPassword(passwords []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`Password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromEmail 通过Email获取内容 用户邮箱
func (obj *_UsersMgr) GetFromEmail(email string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`Email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找 用户邮箱
func (obj *_UsersMgr) GetBatchFromEmail(emails []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`Email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromIsAdmin 通过IsAdmin获取内容 是否是管理员
func (obj *_UsersMgr) GetFromIsAdmin(isAdmin bool) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`IsAdmin` = ?", isAdmin).Find(&results).Error

	return
}

// GetBatchFromIsAdmin 批量查找 是否是管理员
func (obj *_UsersMgr) GetBatchFromIsAdmin(isAdmins []bool) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`IsAdmin` IN (?)", isAdmins).Find(&results).Error

	return
}

// GetFromPassNum 通过PassNum获取内容 通过题目数量
func (obj *_UsersMgr) GetFromPassNum(passNum int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`PassNum` = ?", passNum).Find(&results).Error

	return
}

// GetBatchFromPassNum 批量查找 通过题目数量
func (obj *_UsersMgr) GetBatchFromPassNum(passNums []int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`PassNum` IN (?)", passNums).Find(&results).Error

	return
}

// GetFromSubmitNum 通过SubmitNum获取内容 提交次数
func (obj *_UsersMgr) GetFromSubmitNum(submitNum int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`SubmitNum` = ?", submitNum).Find(&results).Error

	return
}

// GetBatchFromSubmitNum 批量查找 提交次数
func (obj *_UsersMgr) GetBatchFromSubmitNum(submitNums []int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`SubmitNum` IN (?)", submitNums).Find(&results).Error

	return
}

// GetFromCreateTime 通过CreateTime获取内容 创建时间
func (obj *_UsersMgr) GetFromCreateTime(createTime time.Time) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`CreateTime` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_UsersMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`CreateTime` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromIsDeleted 通过IsDeleted获取内容 是否删除
func (obj *_UsersMgr) GetFromIsDeleted(isDeleted int8) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`IsDeleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找 是否删除
func (obj *_UsersMgr) GetBatchFromIsDeleted(isDeleteds []int8) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`IsDeleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UsersMgr) FetchByPrimaryKey(autoID int64) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}
