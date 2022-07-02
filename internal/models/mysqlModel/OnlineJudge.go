package mysqlModel

import (
	"time"
)

// Problems [...]
type Problems struct {
	AutoID       int64     `gorm:"primaryKey;column:AutoID;type:bigint;not null" json:"-"`                      // 自增主键
	ProblemID    string    `gorm:"column:ProblemID;type:varchar(20)" json:"problemId"`                          // 问题id
	Title        string    `gorm:"column:Title;type:varchar(255)" json:"title"`                                 // 问题标题
	Content      string    `gorm:"column:Content;type:varchar(4000)" json:"content"`                            // 问题正文描述
	InputLayout  string    `gorm:"column:InputLayout;type:varchar(1000)" json:"inputLayout"`                    // 输入格式描述
	OutputLayout string    `gorm:"column:OutputLayout;type:varchar(1000)" json:"outputLayout"`                  // 输出格式描述
	TimeLimit    int       `gorm:"column:TimeLimit;type:int" json:"timeLimit"`                                  // 最大运行时间
	MemoryLimit  int       `gorm:"column:MemoryLimit;type:int" json:"memoryLimit"`                              // 最大运行内存
	TestNum      int       `gorm:"column:TestNum;type:int" json:"testNum"`                                      // 测试个数
	CreateTime   time.Time `gorm:"column:CreateTime;type:datetime;default:CURRENT_TIMESTAMP" json:"createTime"` // 创建时间
	IsDeleted    bool      `gorm:"column:IsDeleted;type:tinyint(1);default:0" json:"isDeleted"`                 // 是否删除
}

// TableName get sql table name.获取数据库表名
func (m *Problems) TableName() string {
	return "problems"
}

// ProblemsColumns get sql column name.获取数据库列名
var ProblemsColumns = struct {
	AutoID       string
	ProblemID    string
	Title        string
	Content      string
	InputLayout  string
	OutputLayout string
	TimeLimit    string
	MemoryLimit  string
	TestNum      string
	CreateTime   string
	IsDeleted    string
}{
	AutoID:       "AutoID",
	ProblemID:    "ProblemID",
	Title:        "Title",
	Content:      "Content",
	InputLayout:  "InputLayout",
	OutputLayout: "OutputLayout",
	TimeLimit:    "TimeLimit",
	MemoryLimit:  "MemoryLimit",
	TestNum:      "TestNum",
	CreateTime:   "CreateTime",
	IsDeleted:    "IsDeleted",
}

// Users [...]
type Users struct {
	AutoID     int64     `gorm:"primaryKey;column:AutoID;type:bigint;not null" json:"-"`                      // 自增主键
	UserID     string    `gorm:"column:UserID;type:varchar(20)" json:"userId"`                                // 用户id
	UserName   string    `gorm:"column:UserName;type:varchar(100)" json:"userName"`                           // 用户名
	Password   string    `gorm:"column:Password;type:varchar(255)" json:"password"`                           // 用户密码
	Email      string    `gorm:"column:Email;type:varchar(100)" json:"email"`                                 // 用户邮箱
	IsAdmin    bool      `gorm:"column:IsAdmin;type:tinyint(1)" json:"isAdmin"`                               // 是否是管理员
	PassNum    int64     `gorm:"column:PassNum;type:bigint;default:0" json:"passNum"`                         // 通过题目数量
	SubmitNum  int64     `gorm:"column:SubmitNum;type:bigint;default:0" json:"submitNum"`                     // 提交次数
	CreateTime time.Time `gorm:"column:CreateTime;type:datetime;default:CURRENT_TIMESTAMP" json:"createTime"` // 创建时间
	IsDeleted  int8      `gorm:"column:IsDeleted;type:tinyint;default:0" json:"isDeleted"`                    // 是否删除
}

// TableName get sql table name.获取数据库表名
func (m *Users) TableName() string {
	return "users"
}

// UsersColumns get sql column name.获取数据库列名
var UsersColumns = struct {
	AutoID     string
	UserID     string
	UserName   string
	Password   string
	Email      string
	IsAdmin    string
	PassNum    string
	SubmitNum  string
	CreateTime string
	IsDeleted  string
}{
	AutoID:     "AutoID",
	UserID:     "UserID",
	UserName:   "UserName",
	Password:   "Password",
	Email:      "Email",
	IsAdmin:    "IsAdmin",
	PassNum:    "PassNum",
	SubmitNum:  "SubmitNum",
	CreateTime: "CreateTime",
	IsDeleted:  "IsDeleted",
}
