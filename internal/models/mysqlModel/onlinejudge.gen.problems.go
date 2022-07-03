package mysqlModel

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ProblemsMgr struct {
	*_BaseMgr
}

// ProblemsMgr open func
func ProblemsMgr(db *gorm.DB) *_ProblemsMgr {
	if db == nil {
		panic(fmt.Errorf("ProblemsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProblemsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("problems"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProblemsMgr) GetTableName() string {
	return "problems"
}

// Reset 重置gorm会话
func (obj *_ProblemsMgr) Reset() *_ProblemsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ProblemsMgr) Get() (result Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProblemsMgr) Gets() (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ProblemsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Problems{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAutoID AutoID获取 自增主键
func (obj *_ProblemsMgr) WithAutoID(autoID int64) Option {
	return optionFunc(func(o *options) { o.query["AutoID"] = autoID })
}

// WithProblemID ProblemID获取 问题id
func (obj *_ProblemsMgr) WithProblemID(problemID string) Option {
	return optionFunc(func(o *options) { o.query["ProblemID"] = problemID })
}

// WithTitle Title获取 问题标题
func (obj *_ProblemsMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["Title"] = title })
}

// WithContent Content获取 问题正文描述
func (obj *_ProblemsMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["Content"] = content })
}

// WithInputLayout InputLayout获取 输入格式描述
func (obj *_ProblemsMgr) WithInputLayout(inputLayout string) Option {
	return optionFunc(func(o *options) { o.query["InputLayout"] = inputLayout })
}

// WithOutputLayout OutputLayout获取 输出格式描述
func (obj *_ProblemsMgr) WithOutputLayout(outputLayout string) Option {
	return optionFunc(func(o *options) { o.query["OutputLayout"] = outputLayout })
}

// WithTimeLimit TimeLimit获取 最大运行时间
func (obj *_ProblemsMgr) WithTimeLimit(timeLimit int) Option {
	return optionFunc(func(o *options) { o.query["TimeLimit"] = timeLimit })
}

// WithMemoryLimit MemoryLimit获取 最大运行内存
func (obj *_ProblemsMgr) WithMemoryLimit(memoryLimit int) Option {
	return optionFunc(func(o *options) { o.query["MemoryLimit"] = memoryLimit })
}

// WithTestNum TestNum获取 测试个数
func (obj *_ProblemsMgr) WithTestNum(testNum int) Option {
	return optionFunc(func(o *options) { o.query["TestNum"] = testNum })
}

// WithStatus Status获取 问题测试是否上传
func (obj *_ProblemsMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["Status"] = status })
}

// WithCreateTime CreateTime获取 创建时间
func (obj *_ProblemsMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["CreateTime"] = createTime })
}

// WithIsDeleted IsDeleted获取 是否删除
func (obj *_ProblemsMgr) WithIsDeleted(isDeleted bool) Option {
	return optionFunc(func(o *options) { o.query["IsDeleted"] = isDeleted })
}

// GetByOption 功能选项模式获取
func (obj *_ProblemsMgr) GetByOption(opts ...Option) (result Problems, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProblemsMgr) GetByOptions(opts ...Option) (results []*Problems, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAutoID 通过AutoID获取内容 自增主键
func (obj *_ProblemsMgr) GetFromAutoID(autoID int64) (result Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}

// GetBatchFromAutoID 批量查找 自增主键
func (obj *_ProblemsMgr) GetBatchFromAutoID(autoIDs []int64) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`AutoID` IN (?)", autoIDs).Find(&results).Error

	return
}

// GetFromProblemID 通过ProblemID获取内容 问题id
func (obj *_ProblemsMgr) GetFromProblemID(problemID string) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`ProblemID` = ?", problemID).Find(&results).Error

	return
}

// GetBatchFromProblemID 批量查找 问题id
func (obj *_ProblemsMgr) GetBatchFromProblemID(problemIDs []string) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`ProblemID` IN (?)", problemIDs).Find(&results).Error

	return
}

// GetFromTitle 通过Title获取内容 问题标题
func (obj *_ProblemsMgr) GetFromTitle(title string) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`Title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 问题标题
func (obj *_ProblemsMgr) GetBatchFromTitle(titles []string) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`Title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromContent 通过Content获取内容 问题正文描述
func (obj *_ProblemsMgr) GetFromContent(content string) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`Content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 问题正文描述
func (obj *_ProblemsMgr) GetBatchFromContent(contents []string) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`Content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromInputLayout 通过InputLayout获取内容 输入格式描述
func (obj *_ProblemsMgr) GetFromInputLayout(inputLayout string) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`InputLayout` = ?", inputLayout).Find(&results).Error

	return
}

// GetBatchFromInputLayout 批量查找 输入格式描述
func (obj *_ProblemsMgr) GetBatchFromInputLayout(inputLayouts []string) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`InputLayout` IN (?)", inputLayouts).Find(&results).Error

	return
}

// GetFromOutputLayout 通过OutputLayout获取内容 输出格式描述
func (obj *_ProblemsMgr) GetFromOutputLayout(outputLayout string) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`OutputLayout` = ?", outputLayout).Find(&results).Error

	return
}

// GetBatchFromOutputLayout 批量查找 输出格式描述
func (obj *_ProblemsMgr) GetBatchFromOutputLayout(outputLayouts []string) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`OutputLayout` IN (?)", outputLayouts).Find(&results).Error

	return
}

// GetFromTimeLimit 通过TimeLimit获取内容 最大运行时间
func (obj *_ProblemsMgr) GetFromTimeLimit(timeLimit int) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`TimeLimit` = ?", timeLimit).Find(&results).Error

	return
}

// GetBatchFromTimeLimit 批量查找 最大运行时间
func (obj *_ProblemsMgr) GetBatchFromTimeLimit(timeLimits []int) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`TimeLimit` IN (?)", timeLimits).Find(&results).Error

	return
}

// GetFromMemoryLimit 通过MemoryLimit获取内容 最大运行内存
func (obj *_ProblemsMgr) GetFromMemoryLimit(memoryLimit int) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`MemoryLimit` = ?", memoryLimit).Find(&results).Error

	return
}

// GetBatchFromMemoryLimit 批量查找 最大运行内存
func (obj *_ProblemsMgr) GetBatchFromMemoryLimit(memoryLimits []int) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`MemoryLimit` IN (?)", memoryLimits).Find(&results).Error

	return
}

// GetFromTestNum 通过TestNum获取内容 测试个数
func (obj *_ProblemsMgr) GetFromTestNum(testNum int) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`TestNum` = ?", testNum).Find(&results).Error

	return
}

// GetBatchFromTestNum 批量查找 测试个数
func (obj *_ProblemsMgr) GetBatchFromTestNum(testNums []int) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`TestNum` IN (?)", testNums).Find(&results).Error

	return
}

// GetFromStatus 通过Status获取内容 问题测试是否上传
func (obj *_ProblemsMgr) GetFromStatus(status bool) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`Status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 问题测试是否上传
func (obj *_ProblemsMgr) GetBatchFromStatus(statuss []bool) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`Status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过CreateTime获取内容 创建时间
func (obj *_ProblemsMgr) GetFromCreateTime(createTime time.Time) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`CreateTime` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_ProblemsMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`CreateTime` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromIsDeleted 通过IsDeleted获取内容 是否删除
func (obj *_ProblemsMgr) GetFromIsDeleted(isDeleted bool) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`IsDeleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找 是否删除
func (obj *_ProblemsMgr) GetBatchFromIsDeleted(isDeleteds []bool) (results []*Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`IsDeleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ProblemsMgr) FetchByPrimaryKey(autoID int64) (result Problems, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Problems{}).Where("`AutoID` = ?", autoID).Find(&result).Error

	return
}
