# golang标准库strconv包的学习

### func [Atoi](https://github.com/golang/go/blob/master/src/strconv/atoi.go?name=release#195)

```
func Atoi(s string) (i int, err error)
```

`Atoi` 函数将字符串转换为整数

```
 s := "42"
    i, err := strconv.Atoi(s)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("The integer is: %d\n", i) // The integer is: 42
```

用的特别多，string转int

### func [Itoa](https://github.com/golang/go/blob/master/src/strconv/itoa.go?name=release#24)

```
func Itoa(i int) string
```

```
i := 42
    s := strconv.Itoa(i)
    fmt.Println(s)
```

Int 转string

### func [ParseFloat](https://github.com/golang/go/blob/master/src/strconv/atof.go?name=release#533)

```
func ParseFloat(s string, bitSize int) (f float64, err error)
```

解析一个表示浮点数的字符串并返回其值。

```
 str := "3.1415926"
    f, err := strconv.ParseFloat(str, 64)
    if err != nil {
        fmt.Println("parse float error:", err)
        return
    }
    fmt.Printf("string %s to float64 %f\n", str, f)
```

stinrg转float64

### func [ParseInt](https://github.com/golang/go/blob/master/src/strconv/atoi.go?name=release#150)

```
func ParseInt(s string, base int, bitSize int) (i int64, err error)
```

`ParseInt` 函数将字符串解析为整数，支持不同进制的转换（2~36），返回解析后的 `int64` 类型整数。

```
// 将 10 进制的字符串转为整数
	i1, err1 := strconv.ParseInt("12345", 10, 64)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(i1) // 12345

	// 将 16 进制的字符串转为整数
	i2, err2 := strconv.ParseInt("1E240", 16, 64)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(i2) // 123456

	// 将 2 进制的字符串转为整数
	i3, err3 := strconv.ParseInt("11000000111001", 2, 64)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(i3) // 12345
```

string转int64一般用这个

我感觉常用的就这两个



