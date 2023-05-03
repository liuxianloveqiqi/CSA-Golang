# golang标准库strings学习

### func [Contains](https://github.com/golang/go/blob/master/src/strings/strings.go?name=release#112)

```go
func Contains(s, substr string) bool
```

判断字符串中是否包含子字符串

### func [ContainsAny](https://github.com/golang/go/blob/master/src/strings/strings.go?name=release#117)

```
func ContainsAny(s, chars string) bool
```

判断字符串s是否包含字符串chars中的任一字符

### func [Count](https://github.com/golang/go/blob/master/src/strings/strings.go?name=release#65)

```
func Count(s, sep string) int
```

返回字符串s中有几个不重复的sep子串。

做leetcode时用的比较多

### func [Index](https://github.com/golang/go/blob/master/src/strings/strings.go?name=release#127)

```
func Index(s, sep string) int
```

找index的

### func [Title](https://github.com/golang/go/blob/master/src/strings/strings.go?name=release#489)

```
func Title(s string) string
```

返回s中每个单词的首字母都改为标题格式的字符串拷贝。

改标题的

### func [ToLower](https://github.com/golang/go/blob/master/src/strings/strings.go?name=release#437)

```
func ToLower(s string) string
```

返回将所有字母都转为对应的小写版本的拷贝。

### func [ToUpper](https://github.com/golang/go/blob/master/src/strings/strings.go?name=release#434)

```
func ToUpper(s string) string
```

返回将所有字母都转为对应的大写版本的拷贝。

### func [Repeat](https://github.com/golang/go/blob/master/src/strings/strings.go?name=release#424)

```
func Repeat(s string, count int) string
```

返回count个s串联的字符串。

免得手写，重复字符的

### func [Replace](https://github.com/golang/go/blob/master/src/strings/strings.go?name=release#638)

```
func Replace(s, old, new string, n int) string
```

字符串 `s` 中的前 `n` 个旧字符串 `old` 替换为新字符串 `new`，返回新的字符串。如果 `n` 小于 0，则会替换所有旧字符串。

```
s := "hello, hello, world"
fmt.Println(strings.Replace(s, "hello", "hi", 1))    // 输出：hi, hello, world
fmt.Println(strings.Replace(s, "hello", "hi", -1))   // 输出：hi, hi, world
fmt.Println(strings.Replace(s, "l", "x", 2))         // 输出：hexxo, hello, world
fmt.Println(strings.Replace(s, "l", "x", -1))        // 输出：hexxo, hexxo, worxd

```

### func [Trim](https://github.com/golang/go/blob/master/src/strings/strings.go?name=release#586)

```
func Trim(s string, cutset string) string
```

返回将s前后端所有cutset包含的utf-8码值都去掉的字符串。

### func [TrimSpace](https://github.com/golang/go/blob/master/src/strings/strings.go?name=release#613)

```
func TrimSpace(s string) string
```

返回将s前后端所有空白（unicode.IsSpace指定）都去掉的字符串。

这个用的比较多，去前后的空格

### func [Split](https://github.com/golang/go/blob/master/src/strings/strings.go?name=release#294)

```
func Split(s, sep string) []string
```

`s`参数是需要分割的字符串，`sep`参数是分割符。函数返回的是一个字符串切片，每个元素是分割后得到的一个子串，切片的长度就是分割后得到的子串的个数。

```
str := "apple,banana,orange"
	arr := strings.Split(str, ",")
	fmt.Printf("%#v", arr) // []string{"apple", "banana", "orange"}
```

一般都是用`,`或者`“ ”`分割

### func [Join](https://github.com/golang/go/blob/master/src/strings/strings.go?name=release#349)

```
func Join(a []string, sep string) string
```

将一系列字符串连接为一个字符串，之间用sep来分隔。

拼接字符串的，用的很多

#### func [NewReader](https://github.com/golang/go/blob/master/src/strings/reader.go?name=release#144)

```
func NewReader(s string) *Reader
```

#### func (*Reader) [Len](https://github.com/golang/go/blob/master/src/strings/reader.go?name=release#24)

```
func (r *Reader) Len() int
```

`Len()` 方法返回 `Reader` 剩余未读取的字节数。该方法不会导致数据读取。如果已到达 `EOF`，则 `Len()` 返回 0。

```
reader := strings.NewReader("hello world")
	fmt.Println(reader.Len()) // 输出：11
```

#### func (*Reader) [Seek](https://github.com/golang/go/blob/master/src/strings/reader.go?name=release#103)

```
func (r *Reader) Seek(offset int64, whence int) (int64, error)
```

`Seek()` 方法用于更改当前读取位置

- `offset`：相对于 whence 的偏移量，可以是负数。
- `whence`：从哪个位置开始偏移。有以下 3 种值可供选择：
  - `io.SeekStart`：相对于文件开始位置。
  - `io.SeekCurrent`：相对于当前位置。
  - `io.SeekEnd`：相对于文件结尾位置。

`Seek()` 方法返回新的偏移量和可能的错误。

#### func [NewReplacer](https://github.com/golang/go/blob/master/src/strings/replace.go?name=release#31)

```
func NewReplacer(oldnew ...string) *Replacer
```

`NewReplacer`函数创建并返回一个Replacer对象，该对象可用于在字符串中替换指定的字符串。它的参数是一系列的字符串对，表示要替换的旧字符串和相应的新字符串。在替换过程中，将依次用新字符串替换旧字符串，直到所有旧字符串都被替换完为止。

#### func (*Replacer) [Replace](https://github.com/golang/go/blob/master/src/strings/replace.go?name=release#78)

```
func (r *Replacer) Replace(s string) string
```

下面给个例子

```
r := strings.NewReplacer("Hello", "Hi", "world", "gopher")
	fmt.Println(r.Replace("Hello, world!")) // 输出: Hi, gopher!
```

