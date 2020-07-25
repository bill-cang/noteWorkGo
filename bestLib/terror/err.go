package terror

type ERRORCODE uint32

const (
	/*系统级*/
	SysErr ERRORCODE = 0

	/*控制层：control*/
	RspErr          ERRORCODE = 1000
	ErrRspAuthFiled ERRORCODE = 1001
	ErrRspBadParam  ERRORCODE = 1002
	ErrRspInternal  ERRORCODE = 1003

	/*持久层：mysql*/
	MysqlErr             ERRORCODE = 2000
	ErrMysqlDataNotFound ERRORCODE = 2001

	/*缓存层：redis*/
	RedisErr ERRORCODE = 3000

	/*其他：OthErr*/
	OthErr                  ERRORCODE = 4000
	ErrCodeDecode           ERRORCODE = 4001
	ErrCodeEncode           ERRORCODE = 4002
	ErrCodeServiceNotExists ERRORCODE = 4003
	ErrCodeNetwork          ERRORCODE = 4004
)

var (
	//ErrDecodeFail 解包错误
	ErrDecodeFail = New(ErrCodeDecode, "fail to decode")
	//ErrEncodeFail 封包错误
	ErrEncodeFail = New(ErrCodeEncode, "fail to encode")
	//ErrServiceNotExists 服务不存在
	ErrServiceNotExists = New(ErrCodeServiceNotExists, "service not exists")
	//ErrNetwork 网络错误
	ErrNetwork = New(ErrCodeNetwork, "network error")
	//ErrAuthFail 认证失败
	ErrAuthFail = New(ErrRspAuthFiled, "fail to auth")
	//ErrBadParam 参数错误
	ErrBadParam = New(ErrRspBadParam, "bad parameter")
	//ErrDataNotFound 数据找不到
	ErrDataNotFound = New(ErrMysqlDataNotFound, "data not found")
	//ErrInternal 内部错误
	ErrInternal = New(ErrRspInternal, "internal error")
)
