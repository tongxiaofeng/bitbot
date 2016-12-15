package bitbot

type StrategyVarabileType int

const (
	NUMBER StrategyVarabileType = 1 + iota
	BOOLEAN
	STRING //string
	SELECTED
)

type StrategyVarabile struct {
	Name         string
	Description  string
	Hint         string
	VarabileType StrategyVarabileType
	DefaultValue string
}

type StrategyLang int

const (
	GOLANG StrategyLang = 1 + iota
	NODEJS
	PYTHON
)

type Strategy struct {
	ID                int
	Name              string
	Description       string
	Lang              StrategyLang
	Code              string
	StrategyVarabiles []StrategyVarabile

	LocalPath          string //local execute file relative path
	LocalCommandParams string
}

/*
字符串	:String
数字型	:Number
布尔型	:true或者false
选择型	:用'|'分开, 如aa|bb|cc表示列表有三个选项, 对应值为0,1,2
加密型	:String (策略参数将在本地加密后保存到服务器与秘钥保存原理一样)

参数支持定义显示条件, 比如想让a变量在b变量为1或者true的时候显示, 变量a就定义成a@b或者a@b==1
如果变量a想在变量b在选择第三个选项的时候显示, 变量a变量名就定义为a@b==3
@后面是定义语法格式为 变量名+比较符+值(数字)
操作符支持==, !=, >=, <=, 变量前直接加!表示取反，比如a@!b,指b为1或者true时候,a不显示

*/
