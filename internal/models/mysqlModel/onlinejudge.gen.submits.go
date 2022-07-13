package mysqlModel

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SubmitsMgr struct {
	*_BaseMgr
}

// SubmitsMgr open func
func SubmitsMgr(db *gorm.DB) *_SubmitsMgr {
	if db == nil {
		panic(fmt.Errorf("SubmitsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SubmitsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("submits"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SubmitsMgr) GetTableName() string {
	return "submits"
}

// Reset 重置gorm会话
func (obj *_SubmitsMgr) Reset() *_SubmitsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SubmitsMgr) Get() (result Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SubmitsMgr) Gets() (results []*Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SubmitsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Submits{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAutoID AutoID获取 自增主键
func (obj *_SubmitsMgr) WithAutoID(autoID int64) Option {
	return optionFunc(func(o *options) { o.query["AutoID"] = autoID })
}

// WithSubmitID SubmitID获取
func (obj *_SubmitsMgr) WithSubmitID(submitID string) Option {
	return optionFunc(func(o *options) { o.query["SubmitID"] = submitID })
}

// WithUserID UserID获取 用户id
func (obj *_SubmitsMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["UserID"] = userID })
}

// WithProblemID ProblemID获取 问题id
func (obj *_SubmitsMgr) WithProblemID(problemID string) Option {
	return optionFunc(func(o *options) { o.query["ProblemID"] = problemID })
}

// WithSubmitStatus SubmitStatus获取 提交状态
func (obj *_SubmitsMgr) WithSubmitStatus(submitStatus string) Option {
	return optionFunc(func(o *options) { o.query["SubmitStatus"] = submitStatus })
}

// WithIsDeleted IsDeleted获取
func (obj *_SubmitsMgr) WithIsDeleted(isDeleted bool) Option {
	return optionFunc(func(o *options) { o.query["IsDeleted"] = isDeleted })
}

// WithCreateTime CreateTime获取 创建时间
func (obj *_SubmitsMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["CreateTime"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *_SubmitsMgr) GetByOption(opts ...Option) (result Submits, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SubmitsMgr) GetByOptions(opts ...Option) (results []*Submits, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAutoID 通过AutoID获取内容 自增主键
func (obj *_SubmitsMgr) GetFromAutoID(autoID int64) (result Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}

// GetBatchFromAutoID 批量查找 自增主键
func (obj *_SubmitsMgr) GetBatchFromAutoID(autoIDs []int64) (results []*Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where("`AutoID` IN (?)", autoIDs).Find(&results).Error

	return
}

// GetFromSubmitID 通过SubmitID获取内容
func (obj *_SubmitsMgr) GetFromSubmitID(submitID string) (results []*Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where("`SubmitID` = ?", submitID).Find(&results).Error

	return
}

// GetBatchFromSubmitID 批量查找
func (obj *_SubmitsMgr) GetBatchFromSubmitID(submitIDs []string) (results []*Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where("`SubmitID` IN (?)", submitIDs).Find(&results).Error

	return
}

// GetFromUserID 通过UserID获取内容 用户id
func (obj *_SubmitsMgr) GetFromUserID(userID string) (results []*Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where("`UserID` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户id
func (obj *_SubmitsMgr) GetBatchFromUserID(userIDs []string) (results []*Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where("`UserID` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromProblemID 通过ProblemID获取内容 问题id
func (obj *_SubmitsMgr) GetFromProblemID(problemID string) (results []*Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where("`ProblemID` = ?", problemID).Find(&results).Error

	return
}

// GetBatchFromProblemID 批量查找 问题id
func (obj *_SubmitsMgr) GetBatchFromProblemID(problemIDs []string) (results []*Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where("`ProblemID` IN (?)", problemIDs).Find(&results).Error

	return
}

// GetFromSubmitStatus 通过SubmitStatus获取内容 提交状态
func (obj *_SubmitsMgr) GetFromSubmitStatus(submitStatus string) (results []*Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where("`SubmitStatus` = ?", submitStatus).Find(&results).Error

	return
}

// GetBatchFromSubmitStatus 批量查找 提交状态
func (obj *_SubmitsMgr) GetBatchFromSubmitStatus(submitStatuss []string) (results []*Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where("`SubmitStatus` IN (?)", submitStatuss).Find(&results).Error

	return
}

// GetFromIsDeleted 通过IsDeleted获取内容
func (obj *_SubmitsMgr) GetFromIsDeleted(isDeleted bool) (results []*Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where("`IsDeleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_SubmitsMgr) GetBatchFromIsDeleted(isDeleteds []bool) (results []*Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where("`IsDeleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCreateTime 通过CreateTime获取内容 创建时间
func (obj *_SubmitsMgr) GetFromCreateTime(createTime time.Time) (results []*Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where("`CreateTime` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_SubmitsMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where("`CreateTime` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SubmitsMgr) FetchByPrimaryKey(autoID int64) (result Submits, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Submits{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}
