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
...



## 更新记录
- v1.0.4：datatime、http、config、crypto、uuid/ulid
- v1.0.3：加载配置
- v1.0.1：雪花算法
- v1.0.0：bean属性拷贝



## 使用案例

安装

```go
go get -u github.com/wegoteam/wepkg@latest
```

### base
- 响应结构体
- 分页

### net/http
封装http请求的客户端

### config
加载配置：默认加载环境变量、配置文件、命令行参数


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
属性复制、结构体转map、map转结构体
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

贡献来源：

https://github.com/spf13/viper

https://github.com/redis/go-redis

https://github.com/jeevatkm/go-model

https://github.com/go-resty/resty

https://github.com/golang-module/carbon