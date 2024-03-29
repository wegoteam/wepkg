# wepkg
## 基础组件

- 属性复制
- 分布式雪花算法
- 缓存（redis）
- 字符
- 文件
- json
- 时间日期
- 条件表达式
- 协程池
- 切片
- 日志
- 定时任务
- ...



## 更新记录
- v1.1.7：io/file、io/excel
- v1.0.6：id/random、job/cron
- v1.0.5：io/json、io/xml、log
- v1.0.4：datatime、http、config、crypto、id/uuid、id/ulid
- v1.0.3：config
- v1.0.2：id/snowflake
- v1.0.1：id/snowflake
- v1.0.0：bean



## 使用案例

安装

```go
go get -u github.com/wegoteam/wepkg@latest
```

### base
- 响应结构体
- 分页


### config
加载配置：默认加载环境变量、配置文件、命令行参数
- 默认配置文件加载顺序：命令行参数  > 默认配置文件目录(./config/config.yaml) 
    ```shell
    go run main.go --config=../config/config.yaml
    ```
    ```go
  c := config.GetConfig()
  var mysql = &MySQL{}
  err := c.Load("mysql", mysql)
  if err != nil {
      fmt.Errorf("Fatal error config file: %s \n", err)
  }
  fmt.Printf("mysql prop=%+v \n", mysql)
    ```
- 加载指定位置配置文件
    ```go
  c := config.NewConfig("config", "yaml", "", []string{"."})
  var mysql = &MySQL{}
  err := c.Load("mysql", mysql)
  if err != nil {
      fmt.Errorf("Fatal error config file: %s \n", err)
  }
  fmt.Printf("mysql prop=%+v \n", mysql)
    ```



### crypto
编码解码、加密解密和签名验签库
```go
func TestCrypto(t *testing.T) {
	fmt.Printf("hex 编码: %v\n", crypto.EncodeHex("hello world"))
	fmt.Printf("hex 编码: %s\n", crypto.EncodeHexToBytes([]byte("hello world")))
	fmt.Printf("hex 解码: %v\n", crypto.DecodeHex("68656c6c6f20776f726c64"))
	fmt.Printf("hex 解码: %s\n", crypto.DecodeHexToBytes([]byte("68656c6c6f20776f726c64")))
	fmt.Printf("base64 编码: %v\n", crypto.EncodeBase64("hello world"))
	fmt.Printf("base64 编码: %s\n", crypto.EncodeBase64ToBytes([]byte("hello world")))
	fmt.Printf("base64 解码: %v\n", crypto.DecodeBase64("aGVsbG8gd29ybGQ="))
	fmt.Printf("base64 解码: %s\n", crypto.DecodeBase64ToBytes([]byte("aGVsbG8gd29ybGQ=")))
	fmt.Printf("md5 编码: %v\n", crypto.EncryptMd5ToHex("hello world"))
	fmt.Printf("md5 编码: %v\n", crypto.EncryptMd5ToBase64("hello world"))
	fmt.Printf("md5 编码: %s\n", crypto.EncryptMd5ToHexBytes("hello world"))
	fmt.Printf("md5 编码: %s\n", crypto.EncryptMd5ToBase64Bytes("hello world"))
	fmt.Printf("sha1 编码: %v\n", crypto.EncryptSha1ToHex("hello world"))
	fmt.Printf("sha1 编码: %v\n", crypto.EncryptSha1ToBase64("hello world"))
	fmt.Printf("sha3-224 编码: %v\n", crypto.EncryptSha3ToHex("hello world", 224))
	fmt.Printf("sha3-224 编码: %v\n", crypto.EncryptSha3ToBase64("hello world", 224))
	fmt.Printf("sha3-256 编码: %v\n", crypto.EncryptSha3ToHex("hello world", 256))
	fmt.Printf("sha3-256 编码: %v\n", crypto.EncryptSha3ToBase64("hello world", 256))
	fmt.Printf("sha3-384 编码: %v\n", crypto.EncryptSha3ToHex("hello world", 384))
	fmt.Printf("sha3-384 编码: %v\n", crypto.EncryptSha3ToBase64("hello world", 384))
	fmt.Printf("sha3-512 编码: %v\n", crypto.EncryptSha3ToHex("hello world", 512))
	fmt.Printf("sha3-512 编码: %v\n", crypto.EncryptSha3ToBase64("hello world", 512))
	fmt.Printf("sha256 编码: %v\n", crypto.EncryptSha256ToHex("hello world"))
	fmt.Printf("sha256 编码: %v\n", crypto.EncryptSha256ToBase64("hello world"))

	publicKeyPkcs1, privateKeyPkcs1 := crypto.GenKeyPkcs1Pair()
	fmt.Printf("生成 PKCS1 格式的 RSA 密钥对: publicKeyPkcs1=%s privateKeyPkcs1=%s\n", publicKeyPkcs1, privateKeyPkcs1)
	publicKeyPkcs8, privateKeyPkcs8 := crypto.GenKeyPkcs8Pair()
	fmt.Printf("生成 PKCS8 格式的 RSA 密钥对: publicKeyPkcs8=%s privateKeyPkcs8=%s\n", publicKeyPkcs8, privateKeyPkcs8)

	fmt.Printf("验证 RSA 密钥对是否匹配 :%v\n", crypto.VerifyKeyPair(publicKeyPkcs1, privateKeyPkcs1))
	fmt.Printf("验证 RSA 密钥对是否匹配 :%v\n", crypto.VerifyKeyPair(publicKeyPkcs8, privateKeyPkcs8))
	fmt.Printf("验证是否是 RSA 公钥 :%v\n", crypto.IsPublicKey(publicKeyPkcs1))
	fmt.Printf("验证是否是 RSA 私钥 :%v\n", crypto.IsPrivateKey(privateKeyPkcs8))

	parsePublicKey, _ := crypto.ParsePublicKey(publicKeyPkcs1)
	fmt.Printf("解析公钥 :%v\n", parsePublicKey)

	parsePrivateKey, _ := crypto.ParsePrivateKey(privateKeyPkcs1)
	fmt.Printf("解析私钥 :%v\n", parsePrivateKey)

	exportPrivateKey, exportPrivateKeyErr := crypto.ExportPublicKey(publicKeyPkcs1)
	if exportPrivateKeyErr != nil {
		fmt.Errorf("exportPrivateKeyErr:%s\n", exportPrivateKeyErr.Error())
	}
	fmt.Printf("从 RSA 私钥里导出公钥 :%v\n", exportPrivateKey)
}
```
### datetime
时间日期
```go
func TestTime(t *testing.T) {
    fmt.Printf("当前时间：%v \n", datetime.Now())
    fmt.Printf("当前时间戳：%v \n", datetime.Timestamp())
    fmt.Printf("当前时间戳：%v \n", time.Now().Unix())
    fmt.Printf("当前毫秒级时间戳：%v \n", datetime.TimestampMilli())
    fmt.Printf("当前毫秒级时间戳：%v \n", time.Now().UnixNano())
    fmt.Printf("当前微秒级时间戳：%v \n", datetime.TimestampMicro())
    fmt.Printf("当前纳秒级时间戳：%v \n", datetime.TimestampNano())
    fmt.Printf("昨天时间：%v \n", datetime.Yesterday())
    fmt.Printf("明天时间：%v \n", datetime.Tomorrow())
    fmt.Printf("字符串转time：%v \n", datetime.Parse("2023-07-22 13:14:15"))
    fmt.Printf("当前时间转字符串：%v \n", datetime.ToString(time.Now()))
    fmt.Printf("当前时间转正则字符串：%v \n", datetime.Format(time.Now(), "Y-m-d H:i:s.U"))
    fmt.Printf("当前时间转正则字符串：%v \n", datetime.Layout(time.Now(), "2006-01-02 15:04:05.999"))
    fmt.Printf("当前时间改变年数：%v \n", datetime.ChangeYears(time.Now(), 1))
    fmt.Printf("当前时间改变年数：%v \n", datetime.ChangeYears(time.Now(), -1))
    fmt.Printf("当前时间改变月数：%v \n", datetime.ChangeMonths(time.Now(), 1))
    fmt.Printf("当前时间改变月数：%v \n", datetime.ChangeMonths(time.Now(), -1))
    fmt.Printf("当前时间改变天数：%v \n", datetime.ChangeDays(time.Now(), 1))
    fmt.Printf("当前时间改变天数：%v \n", datetime.ChangeDays(time.Now(), -1))
    fmt.Printf("当前时间改变小时：%v \n", datetime.ChangeHours(time.Now(), 1))
    fmt.Printf("当前时间改变小时：%v \n", datetime.ChangeHours(time.Now(), -1))
    fmt.Printf("当前时间改变分钟：%v \n", datetime.ChangeMinutes(time.Now(), 1))
    fmt.Printf("当前时间改变分钟：%v \n", datetime.ChangeMinutes(time.Now(), -1))
    fmt.Printf("当前时间改变秒数：%v \n", datetime.ChangeSeconds(time.Now(), 1))
    fmt.Printf("当前时间改变秒数：%v \n", datetime.ChangeSeconds(time.Now(), -1))
    fmt.Printf("当前时间改变毫秒数：%v \n", datetime.ChangeMilliseconds(time.Now(), 1))
    fmt.Printf("当前时间改变毫秒数：%v \n", datetime.ChangeMilliseconds(time.Now(), -1))
    fmt.Printf("两个时间相差的年数：%v \n", datetime.DiffYear(time.Now(), datetime.Parse("2023-07-22 13:14:15")))
    fmt.Printf("两个时间相差的年数：%v \n", datetime.DiffAbsYear(time.Now(), datetime.Parse("2023-07-22 13:14:15")))
    fmt.Printf("两个时间相差的月数：%v \n", datetime.DiffMonth(time.Now(), datetime.Parse("2023-07-22 13:14:15")))
    fmt.Printf("两个时间相差的月数：%v \n", datetime.DiffAbsMonth(time.Now(), datetime.Parse("2023-07-22 13:14:15")))
    fmt.Printf("两个时间相差的周数：%v \n", datetime.DiffWeek(time.Now(), datetime.Parse("2023-07-22 13:14:15")))
    fmt.Printf("两个时间相差的周数：%v \n", datetime.DiffAbsWeek(time.Now(), datetime.Parse("2023-07-22 13:14:15")))
    fmt.Printf("两个时间相差的天数：%v \n", datetime.DiffDay(time.Now(), datetime.Parse("2023-07-22 13:14:15")))
    fmt.Printf("两个时间相差的天数：%v \n", datetime.DiffAbsDay(time.Now(), datetime.Parse("2023-07-22 13:14:15")))
    fmt.Printf("两个时间相差的小时：%v \n", datetime.DiffHour(time.Now(), datetime.Parse("2023-07-22 13:14:15")))
    fmt.Printf("两个时间相差的小时：%v \n", datetime.DiffAbsHour(time.Now(), datetime.Parse("2023-07-22 13:14:15")))
    fmt.Printf("两个时间相差的分钟：%v \n", datetime.DiffMinute(time.Now(), datetime.Parse("2023-07-22 13:14:15")))
    fmt.Printf("两个时间相差的分钟：%v \n", datetime.DiffAbsMinute(time.Now(), datetime.Parse("2023-07-22 13:14:15")))
    fmt.Printf("两个时间相差的秒数：%v \n", datetime.DiffSecond(time.Now(), datetime.Parse("2023-07-22 13:14:15")))
    fmt.Printf("两个时间相差的秒数：%v \n", datetime.DiffAbsSecond(time.Now(), datetime.Parse("2023-07-22 13:14:15")))
    fmt.Printf("时间转换时间戳：%v \n", datetime.ToTimestamp(time.Now()))
    fmt.Printf("时间转换毫秒级时间戳：%v \n", datetime.ToTimestampMilli(time.Now()))
    fmt.Printf("时间转换微秒级时间戳：%v \n", datetime.ToTimestampMicro(time.Now()))
    fmt.Printf("时间转换纳秒级时间戳：%v \n", datetime.ToTimestampNano(time.Now()))
    fmt.Printf("判断时间是否有效：%v \n", datetime.IsEffective("0"))
}

```

### bean
- [x] 属性复制
- [x] 结构体转map
- [x] map转结构体
- [x] 结构体字段、tag、值获取
```go
type A struct {
    Age  int    `json:"age"`
    Name string `json:"name"`
}

type B struct {
    Name string
}

//结构体拷贝
func TestBeanCopy(t *testing.T) {
    a1 := A{Age: 100}
    a2 := A{Name: "C"}
    //拷贝属性：a2拷贝a1的属性
    errors := bean.BeanCopy(&a1, a2)
    a2.Name = "D"
    a1.Name = "E"
    fmt.Println("Errors:", errors)
    fmt.Println("a1的地址:", &a1)
    fmt.Println("a1:", a1)
    fmt.Println("a2的地址:", &a2)
    fmt.Println("a2:", a2)
}
//结构体转map
func TestBeanToMap(t *testing.T) {
    a := A{Age: 100, Name: "C"}
    toMap, err := bean.BeanToMap(a)
    fmt.Println("Errors:", err)
    fmt.Println("toMap:", toMap)
}

//结构体克隆
func TestBeanClone(t *testing.T) {
    a := &A{Age: 100, Name: "C"}
    clone, err := bean.BeanClone(a)
    clone.(*A).Name = "D"
    fmt.Println("Errors:", err)
    fmt.Println("a:", a)
    fmt.Println("clone:", clone)
    fmt.Println("a的地址:", &a)
    fmt.Println("clone的地址:", &clone)
}

//判断结构体是否为空
func TestIsZero(t *testing.T) {
    a := A{}
    zero := bean.IsZero(a)
    fmt.Println("zero:", zero)
}

//判断结构体是否有空值
func TestHasZero(t *testing.T) {
    existZero := bean.HasZero(A{Name: "1"})
    existZero2 := bean.HasZero(A{Name: "test", Age: 1})
    fmt.Println("existZero:", existZero)
    fmt.Println("existZero2:", existZero2)
}

//判断结构体指定字段是否有空值
func TestHasFieldsZero(t *testing.T) {
    field, existFieldsZero := bean.HasFieldsZero(A{}, "Name", "Age")
    fmt.Println("existFieldsZero:", existFieldsZero)
    fmt.Println("field:", field)
}

//获取结构体所有字段
func TestGetFields(t *testing.T) {
    field, err := bean.GetFields(A{})
    fmt.Println("err:", err)
    fmt.Println("field:", field)
}

//获取结构体指定字段的类型
func TestGetKind(t *testing.T) {
    kind, err := bean.GetKind(A{}, "Name")
    fmt.Println("err:", err)
    fmt.Println("kind:", kind)
}

//获取结构体指定字段的tag
func TestGetTag(t *testing.T) {
    tag, err := bean.GetTag(A{}, "Name")
    fmt.Println("err:", err)
    fmt.Println("tag:", tag.Get("json"))
}

//获取结构体所有字段的tag
func TestGetTags(t *testing.T) {
    tagsMap, err := bean.GetTags(A{})
    fmt.Println("err:", err)
    fmt.Println("tagsMap:", tagsMap)
}

//GetFieldVal
func TestGetFiled(t *testing.T) {
    field, err := bean.GetFieldVal(A{Name: "test"}, "Name")
    fmt.Println("err:", err)
    fmt.Println("field:", field)
}

//设置结构体指定字段的值
func TestSetFieldVal(t *testing.T) {
    var a = A{Name: "test"}
    fmt.Println("a:", a)
    err := bean.SetFieldVal(&a, "Name", "test2")
    fmt.Println("err:", err)
    fmt.Println("修改后a:", a)
}
```



### id
#### snowflake
- 雪花算法配置读取./config/config.yaml文件的配置
```yaml
#雪花算法配置
snowflake:
  #雪花计算方法,（1-漂移算法|2-传统算法），默认1
  method: 1
  #基础时间（ms单位），不能超过当前系统时间
  baseTime: 1582136402000
  #机器码，必须由外部设定，最大值 2^bitLength-1
  workerId: 1
  #机器码位长，默认值6，取值范围 [1, 15]（要求：序列数位长+机器码位长不超过22）
  bitLength: 6
  #序列数位长，默认值6，取值范围 [3, 21]（要求：序列数位长+机器码位长不超过22）
  seqBitLength: 6
```
- 设置NewSnowflakeOptions方法的参数配置雪花算法
```go
/**
Method：雪花计算方法,（1-漂移算法|2-传统算法），默认1
BaseTime：基础时间（ms单位），不能超过当前系统时间
WorkerId：机器码，必须由外部设定，最大值 2^WorkerIdBitLength-1
WorkerIdBitLength：机器码位长，默认值6，取值范围 [1, 15]（要求：序列数位长+机器码位长不超过22）
SeqBitLength：序列数位长，默认值6，取值范围 [3, 21]（要求：序列数位长+机器码位长不超过22）
MaxSeqNumber：最大序列数（含），设置范围 [MinSeqNumber, 2^SeqBitLength-1]，默认值0，表示最大序列数取最大值（2^SeqBitLength-1]）
MinSeqNumber：最小序列数（含），默认值5，取值范围 [5, MaxSeqNumber]，每毫秒的前5个序列数对应编号0-4是保留位，其中1-4是时间回拨相应预留位，0是手工新值预留位
TopOverCostCount：最大漂移次数（含），默认2000，推荐范围500-10000（与计算能力有关）
*/
func TestSnowflake(t *testing.T) {
	// 创建配置对象
	var options = snowflake.NewSnowflakeOptions(1)
	options.Method = 1
	options.WorkerIdBitLength = 6
	options.SeqBitLength = 6

	// 保存配置
	snowflake.SetSnowflakeOptions(options)

	for {
		//生成ID
		var newId = snowflake.GenSnowflakeId()
		fmt.Println(newId)
		time.Sleep(time.Second)
	}
}

```

使用方法
```go
/**
  使用默认配置生成
*/
func TestSnowflakeId(t *testing.T) {
    //返回字符串雪花算法ID
    var newStrId = GetSnowflakeId()
    
    fmt.Println(newStrId)
    
    //返回int64雪花算法ID
    newId := GenSnowflakeId()
    fmt.Println(newId)
}
```

#### uuid/ulid
```go
func TestUUID(t *testing.T) {
	fmt.Printf("uuid: %s\n", uuid.New())
	fmt.Printf("ulid: %s\n", ulid.New())
}
```

#### random
随机字符串，随机数字
```go
func TestRandom(t *testing.T) {
	randomStr := rand.RandomStr(10)
	randomNum := rand.RandomNum(10)
	fmt.Printf("randomStr: %s\n", randomStr)
	fmt.Printf("randomNum: %s\n", randomNum)
}
```

### log
日志记录：支持日志文件切割，日志级别，日志格式化，日志文件压缩，日志文件清理

配置文件设置方式：配置文件、手动设置
- 日志配置配置读取./config/config.yaml文件的配置
```yaml
#日志配置
logger:
  #日志输出格式，可选项：json、text
  format: text
  #日志级别，可选项：trace、debug、info、warn、error、panic、fatal
  level: info
  #日志输出位置，可选项：console、file；多个用逗号分隔
  output: console,file
  #日志文件名
  fileName: "./log/wegopkg.log"
  #日志文件最大大小，单位：MB
  maxSize: 10
  #日志文件最大保存时间，单位：天
  maxAge: 3
  #日志文件最大备份数量
  maxBackups: 50
```
- 设置SetLoggerConfig方法的参数配置日志
```go
func TestLogConfig(t *testing.T) {
	log.SetLoggerConfig(log.LoggerConfig{
		Level:  "trace",
		Format: "text",
		Output: "console",
	})
	config := log.GetLoggerConfig()
	fmt.Printf("config=%v\n", config)
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
}
```

使用方法
```go
func TestLog(t *testing.T) {
	log.Trace("Something very low level.")
	log.Tracef("Something very low level. %s", "test")
	log.Traceln("Something very low level.")

	log.Debug("Useful debugging information.")
	log.Debugf("Useful debugging information. %s", "test")
	log.Debugln("Useful debugging information.")

	log.Info("Something noteworthy happened!")
	log.Infof("Something noteworthy happened! %s", "test")
	log.Infoln("Something noteworthy happened!")

	log.Notice("Something unusual happened.")
	log.Noticef("Something unusual happened. %s", "test")
	log.Noticef("Something unusual happened.")

	log.Warn("You should probably take a look at this.")
	log.Warnf("You should probably take a look at this. %s", "test")
	log.Warnln("You should probably take a look at this.")

	log.Error("Something failed but I'm not quitting.")
	log.Errorf("Something failed but I'm not quitting. %s", "test")
	log.Errorln("Something failed but I'm not quitting.")

	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	log.Fatalf("Bye. %s", "test")
	log.Fatalln("Bye.")

	// Calls panic() after logging
	log.Panic("I'm bailing.")
	log.Panicf("I'm bailing. %s", "test")
	log.Panicln("I'm bailing.")

}
```


### net
-[x] http client的封装（get post put...）
-[ ] rpc
-[ ] websocket
-[ ] tcp
-[ ] udp
-[ ] grpc
-[ ] mqtt
-[ ] nats
#### net/http
封装http请求的客户端

```go
func TestDefaultClientPOST(t *testing.T) {
	client := http.BuildDefaultClient()
	var res string
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"roleName":""}`).
		SetResult(res).
		Post("http://localhost:18080/weflow/role/list")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Response Info:", resp)
	fmt.Println("Response Info:", res)
}

type Response[T any] struct {
	Code int    `json:"code"` // 0:成功，其他：失败
	Msg  string `json:"msg"`  // 错误信息
	Data T      `json:"data"` // 数据
}

type RoleInfoResult struct {
	ID         int64  `json:"id"`         // 唯一id
	RoleID     string `json:"roleID"`     // 角色id
	ParentID   string `json:"parentID"`   // 角色父id
	RoleName   string `json:"roleName"`   // 角色名称
	Status     int32  `json:"status"`     // 状态【1：未启用；2：已启用；3：锁定；】
	Remark     string `json:"remark"`     // 描述
	CreateUser string `json:"createUser"` // 创建人
	UpdateUser string `json:"updateUser"` // 更新人
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}

func TestGet(t *testing.T) {
	res1, err := http.Get[Response[[]RoleInfoResult]]("http://localhost:18080/weflow/role/list",
		map[string]string{
			"roleName": "",
		})
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Response Info:", res1)

	res2, err := http.GetString("http://localhost:18080/weflow/role/list",
		map[string]string{
			"roleName": "",
		})
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Response Info:", res2)
}

func TestPost(t *testing.T) {
	type Role struct {
		RoleName string `json:"roleName"`
	}
	var param = &Role{}
	res1, err := http.Post[Response[[]RoleInfoResult]]("http://localhost:18080/weflow/role/list", param)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Response Info:", res1)

	res2, err := http.PostString("http://localhost:18080/weflow/role/list", param)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Response Info:", res2)

	res3, err := http.PostForm[Response[[]RoleInfoResult]]("http://localhost:18080/weflow/role/list",
		map[string]string{
			"roleName": "",
		})
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Response Info:", res3)

	res4, err := http.PostFile[Response[any]]("http://localhost:18080/weflow/upload/file", "a.txt", "./testdata/a.txt")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Response Info:", res4)
	
	res5, err := http.PostFiles[Response[any]]("http://localhost:18080/weflow/upload/file", map[string]string{
		"a.txt": "./testdata/a.txt",
	})
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Response Info:", res5)
}
```

### io
- [x] 文件
- [x] json
- [x] xml
- [x] excel
- [ ] doc/pdf
- [x] 压缩字符串
- [ ] 压缩文件
#### io/json
json序列化和反序列化
```go
func TestFormat(t *testing.T) {
	type Role struct {
		RoleName string `json:"roleName"`
	}
	var param = &Role{
		RoleName: "admin",
	}
	marshal, err := json.Marshal(param)
	if err != nil {
		fmt.Errorf("json.Marshal err: %v", err)
	}
	fmt.Println(marshal)

	var role = &Role{}
	err = json.Unmarshal(marshal, role)
	if err != nil {
		fmt.Errorf("json.Unmarshal err: %v", err)
	}
	fmt.Println(role)
}
```

#### io/compress
字符串压缩
```go
func TestCompress(t *testing.T) {
	var dst []byte
	var source = []byte("test")
	encode := compress.Encode(dst, source)
	fmt.Printf("encode:%s\n", encode)
	fmt.Printf("dst encode:%s\n", dst)
	var src []byte
	decode, err := compress.Decode(encode, src)
	if err != nil {
		fmt.Errorf("err:%s\n", err.Error())
	}
	fmt.Printf("decode:%s\n", decode)
	fmt.Printf("src decode:%s\n", src)
}
```

#### io/file
文件操作：复制、移动等
```go
func TestFile(t *testing.T) {
	contentType, ext, parent := file.GetFileType("./testdata/a.txt")
	fmt.Printf("contentType=%v, ext=%v, parent=%v \n", contentType, ext, parent)
	fileType := file.GetFileExt("./testdata/a.txt")
	fmt.Printf("fileType=%v \n", fileType)
	isOn32bitArch := file.IsOn32bitArch()
	fmt.Printf("是否32位系统架构：%v \n", isOn32bitArch)
	isOn64bitArch := file.IsOn64bitArch()
	fmt.Printf("是否64位系统架构：%v \n", isOn64bitArch)
	isOnLinux := file.IsOnLinux()
	fmt.Printf("是否Linux系统：%v \n", isOnLinux)
	isOnMacOS := file.IsOnMacOS()
	fmt.Printf("是否MacOS系统：%v \n", isOnMacOS)
	isOnWindows := file.IsOnWindows()
	fmt.Printf("是否Windows系统：%v \n", isOnWindows)
	//file.ChangeExeDir()
	existDir := file.ExistDir("./testdata")
	fmt.Printf("是否存在目录：%v \n", existDir)
	existFile := file.ExistFile("./testdata/a.txt")
	fmt.Printf("是否存在文件：%v \n", existFile)
	existSymlink := file.ExistSymlink("./testdata/a.txt")
	fmt.Printf("是否存在软连接：%v \n", existSymlink)
	isDirEmpty, _ := file.IsDirEmpty("./testdata")
	fmt.Printf("目录是否为空：%v \n", isDirEmpty)
	isFileEmpty, _ := file.IsFileEmpty("./testdata/a.txt")
	fmt.Printf("文件是否为空：%v \n", isFileEmpty)
	size, _ := file.GetDirSize("./testdata")
	fmt.Printf("目录大小：%v \n", size)
	fileSize, _ := file.GetFileSize("./testdata/a.txt")
	fmt.Printf("文件大小：%v \n", fileSize)
	symlinkSize, _ := file.GetSymlinkSize("./testdata/a.txt")
	fmt.Printf("软连接大小：%v \n", symlinkSize)
	sameDirEntries, _ := file.SameDirEntries("./testdata", "./testdata")
	fmt.Printf("目录是否相同：%v \n", sameDirEntries)
	sameFileContent, _ := file.SameFileContent("./testdata/a.txt", "./testdata/a.txt")
	fmt.Printf("文件是否相同：%v \n", sameFileContent)
	sameSymlinkContent, _ := file.SameSymlinkContent("./testdata/a.txt", "./testdata/a.txt")
	fmt.Printf("软连接是否相同：%v \n", sameSymlinkContent)
	listDir, _ := file.ListDir("./testdata")
	fmt.Printf("目录列表：%v \n", listDir)
	listFile, _ := file.ListFile("./testdata")
	fmt.Printf("文件列表：%v \n", listFile)
	listSymlink, _ := file.ListSymlink("./testdata")
	fmt.Printf("软连接列表：%v \n", listSymlink)
	err := file.CopyDir("./testdata", "./testdata2")
	fmt.Printf("复制目录错误：%v \n", err)
	err = file.CopyFile("./testdata/a.txt", "./testdata2/a.txt")
	fmt.Printf("复制文件错误：%v \n", err)
	err = file.CopySymlink("./testdata/a.txt", "./testdata2/a.txt")
	fmt.Printf("复制软连接错误：%v \n", err)
	err = file.MoveDir("./testdata", "./testdata2")
	fmt.Printf("移动目录错误：%v \n", err)
	err = file.MoveFile("./testdata/a.txt", "./testdata2/a.txt")
	fmt.Printf("移动文件错误：%v \n", err)
	err = file.MoveSymlink("./testdata/a.txt", "./testdata2/a.txt")
	fmt.Printf("移动软连接错误：%v \n", err)
	listMatch, _ := file.ListMatch("./testdata", file.ListIncludeAll, "*.txt")
	fmt.Printf("匹配列表：%v \n", listMatch)
	joinPath := file.JoinPath("./testdata", "a.txt")
	fmt.Printf("拼接路径：%v \n", joinPath)
	exist := file.Exist("./testdata/a.txt")
	fmt.Printf("是否存在：%v \n", exist)
	notExist := file.NotExist("./testdata/a.txt")
	fmt.Printf("是否不存在：%v \n", notExist)
	err = file.MakeDir("./testdata2")
	fmt.Printf("创建目录错误：%v \n", err)
	//根据名称排序
	sort.Stable(file.SortListByName(listFile))
	//根据大小排序
	sort.Stable(file.SortListBySize(listFile))
	//根据修改时间排序
	sort.Stable(file.SortListByModTime(listFile))
}
```

#### io/excel
excel操作
```go

```

### job
任务调度
#### jon/cron
定时任务

```text
cron表达式语法：
 ┌──分钟（0 - 59）
 │ ┌──小时（0 - 23）
 │ │ ┌──日（1 - 31）
 │ │ │ ┌─月（1 - 12）
 │ │ │ │ ┌─星期（0 - 6，表示从周日到周六）
 │ │ │ │ │
 *  *  *  *  * 被执行的命令
 秒 分 时 日 月 周
```
使用方法
```go
func TestCron(t *testing.T) {
	jonID, _ := cron.AddJob("0/2 * * * * ?", JobTest1)
	cron.AddJob("0/2 * * * * ?", JobTest2)
	time.Sleep(time.Second * 10)
	cron.DelJob(jonID)
	time.Sleep(time.Second * 10)
}

func JobTest1() {
	fmt.Println("JobTest1")
}

func JobTest2() {
	fmt.Println("JobTest2")
}
```

## 引用列表
- https://github.com/spf13/viper
- https://github.com/redis/go-redis
- https://github.com/jeevatkm/go-model
- https://github.com/go-resty/resty
- https://github.com/golang-module/carbon
- https://github.com/qax-os/excelize