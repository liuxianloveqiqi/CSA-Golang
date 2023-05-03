# golang标准库json包的学习

就系列化和反序列化

### func [Marshal](https://github.com/golang/go/blob/master/src/encoding/json/encode.go?name=release#131)

```
func Marshal(v interface{}) ([]byte, error)
```

```
type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"Alice", 25}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
```

输出

```
{"Name":"Alice","Age":25}
```

如果结构体里没有json标签，并且字段首字母小写，那么将不会序列化

## json标签

- `omitempty`: 如果该字段的值是零值或者是空值，那么在编码的过程中就忽略这个字段，不进行编码。
- `string`: 将该字段编码为JSON字符串，而不是原本的类型。
- `number`: 将该字段编码为JSON数字，而不是原本的类型。
- `null`: 将该字段的值编码为null，而不是零值或者空值。
- `inline`: 将该字段内嵌到外层结构体中，而不是在JSON中单独编码。

## func [Unmarshal](https://github.com/golang/go/blob/master/src/encoding/json/decode.go?name=release#67)

```
func Unmarshal(data []byte, v interface{}) error {}
```

例如，如果我们有以下 JSON 字符串：

```
jsonCopy code
{"name": "Alice", "age": 30, "isStudent": true}
```

我们可以使用以下 Go 代码将其解码：

```
goCopy code
type Person struct {
    Name      string `json:"name"`
    Age       int    `json:"age"`
    IsStudent bool   `json:"isStudent"`
}

data := []byte(`{"name": "Alice", "age": 30, "isStudent": true}`)
var p Person
if err := json.Unmarshal(data, &p); err != nil {
    fmt.Println(err)
}
fmt.Println(p) //{Alice 30 true}

```