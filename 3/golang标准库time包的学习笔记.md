# golang标准库time包的学习笔记

# time

### type [ParseError](https://github.com/golang/go/blob/master/src/time/format.go?name=release#598)

```go
type ParseError struct {
    Layout     string
    Value      string
    LayoutElem string
    ValueElem  string
    Message    string
}
```

```go
func (e *ParseError) Error() string
```

感觉没用过啊，解析时间字符串时出现的错误，有个err方法返回错误string

### type [Weekday](https://github.com/golang/go/blob/master/src/time/time.go?name=release#115)

```go
const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```

这么看星期天是0😢

有个string方法，返回星期几的名字

```go
func (d Weekday) String() string
```

### type [Month](https://github.com/golang/go/blob/master/src/time/time.go?name=release#79)

```go
const (
    January Month = 1 + iota
    February
    March
    April
    May
    June
    July
    August
    September
    October
    November
    December
)
```

跟weekday差不多吧，也有个string方法

```go
func (m Month) String() stringf
```

### type [Location](https://github.com/golang/go/blob/master/src/time/zoneinfo.go?name=release#15)

```go
var Local *Location = &localLoc
```

`time.Local` 本地

```go
var UTC *Location = &utcLoc
```

`time.UTC` 通用，零时区

#### func [LoadLocation](https://github.com/golang/go/blob/master/src/time/zoneinfo.go?name=release#273)

```go
func LoadLocation(name string) (*Location, error)
```

> 如果name是""或"UTC"，返回UTC；
>
> 如果name是"Local"，返回Local；
>
> 否则name应该是IANA时区数据库里有记录的地点名（该数据库记录了地点和对应的时区），如"America/New_York"

特地查了一下，中国的时区名字是`"Asia/Shanghai"`,即UTC+8

#### func [FixedZone](https://github.com/golang/go/blob/master/src/time/zoneinfo.go?name=release#89)

```go
func FixedZone(name string, offset int) *Location
```

> FixedZone使用给定的地点名name和时间偏移量offset（单位秒）创建并返回一个Location

偏移时区的吧，感觉几乎不会用

#### func (*Location) [String](https://github.com/golang/go/blob/master/src/time/zoneinfo.go?name=release#83)

```
func (l *Location) String() string
```

> String返回对时区信息的描述，返回值绑定为LoadLocation或FixedZone函数创建l时的name参数。

返回时区name

### type [Time](https://github.com/golang/go/blob/master/src/time/time.go?name=release#34)

```go
type Time struct {
	// wall and ext encode the wall time seconds, wall time nanoseconds,
	// and optional monotonic clock reading in nanoseconds.
	//
	// From high to low bit position, wall encodes a 1-bit flag (hasMonotonic),
	// a 33-bit seconds field, and a 30-bit wall time nanoseconds field.
	// The nanoseconds field is in the range [0, 999999999].
	// If the hasMonotonic bit is 0, then the 33-bit field must be zero
	// and the full signed 64-bit wall seconds since Jan 1 year 1 is stored in ext.
	// If the hasMonotonic bit is 1, then the 33-bit field holds a 33-bit
	// unsigned wall seconds since Jan 1 year 1885, and ext holds a
	// signed 64-bit monotonic clock reading, nanoseconds since process start.
	wall uint64
	ext  int64

	// loc specifies the Location that should be used to
	// determine the minute, hour, month, day, and year
	// that correspond to this Time.
	// The nil location means UTC.
	// All UTC times are represented with loc==nil, never loc==&utcLoc.
	loc *Location
}

```

`time.Time`结构体定义:

- `wall`：墙钟时间，包含三个部分：1位标志位（用于指示是否使用单调时钟）、33位秒数、30位纳秒数。
- `ext`：额外的时间，包含两个部分：1位标志位（用于指示是否使用单调时钟）、63位纳秒数。如果`wall`中的标志位为0，则`ext`存储的是自纪元以来的完整秒数；如果`wall`中的标志位为1，则`ext`存储的是单调时钟的纳秒数。
- `loc`：时区信息，用于确定时间点对应的具体日期和时间。如果`loc`为nil，则表示该时间点采用UTC时区。

通过`time.Format`等函数将时间点格式化为字符串,这个用的多

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    fmt.Println(t.Format("2006-01-02 15:04:05")) //小口决，612345
}

```



### func [Date](https://github.com/golang/go/blob/master/src/time/time.go?name=release#1022)

```go
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
```

参数说明：

- `year`：年份，如2023。
- `month`：月份，使用`time.Month`类型表示，取值范围为`time.January`到`time.December`。
- `day`：日期，1到31之间的整数。
- `hour`：小时，0到23之间的整数。
- `min`：分钟，0到59之间的整数。
- `sec`：秒数，0到59之间的整数。
- `nsec`：纳秒数，0到999999999之间的整数。
- `loc`：时区，==如果为`nil`则表示使用UTC时区==，否则表示使用指定的时区。

Date返回一个时区为loc、当地时间为：

```go
year-month-day hour:min:sec + nsec nanoseconds
```

`time.Date()`函数返回一个`time.Time`类型的值，表示指定的日期和时间。例如，下面创建了一个表示2023年5月4日12点0分0秒的时间点：

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Date(2023, time.May, 4, 12, 0, 0, 0, time.UTC)
    fmt.Println(t)
}

```

#### func [Now](https://github.com/golang/go/blob/master/src/time/time.go?name=release#784)

```go
func Now() Time
```

Now返回当前本地时间。

这个可以用的太多了，经典`time.Now()`

#### func [Parse](https://github.com/golang/go/blob/master/src/time/format.go?name=release#711)

```go
func Parse(layout, value string) (Time, error)
```

Parse解析一个格式化的时间字符串并返回它代表的时间。layout定义了参考时间：

```go
Mon Jan 2 15:04:05 -0700 MST 2006
```

参数说明：

- `layout`：解析模板，用于指定字符串的格式。
- `value`：待解析的字符串。

`layout`参数用于指定字符串的格式，常见的格式如下：

| 布局（layout）              | 含义                                              |
| --------------------------- | ------------------------------------------------- |
| "2006-01-02"                | 日期，如"2023-05-04"                              |
| "15:04:05"                  | 时间，如"12:30:00"                                |
| "2006-01-02 15:04:05"       | 日期和时间，如"2023-05-04 12:30:00"               |
| "2006-01-02T15:04:05Z07:00" | 带时区的日期和时间，如"2023-05-04T12:30:00+08:00" |

`time.Parse()`函数会将字符串按照指定的格式解析成`time.Time`类型的值。如果解析成功，它会返回一个表示时间的`time.Time`类型的值；否则，它会返回一个错误。

例子：

```go
 layout := "2006-01-02"
    value := "2023-05-04"
    t, err := time.Parse(layout, value)
    if err != nil {
        fmt.Println("parse error:", err)
        return
    }
    fmt.Println(t.Year()) //打印2023
```

#### func [ParseInLocation](https://github.com/golang/go/blob/master/src/time/format.go?name=release#720)

```
func ParseInLocation(layout, value string, loc *Location) (Time, error)
```

它与 `Parse` 函数的区别在于，它可以指定解析出的时间值对应的时区。

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    loc, _ := time.LoadLocation("America/New_York")
    t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-05-03 10:15:30", loc)
    fmt.Println(t) // 2023-05-03 10:15:30 -0400 EDT

}

```

 将Unix 时间（自 1970 年 1 月 1 日 UTC 起的秒数和纳秒数）转换为 `Time` 类型的时间值。

#### func (Time) [Location](https://github.com/golang/go/blob/master/src/time/time.go?name=release#813)

```go
func (t Time) Location() *Location
```

> Location返回t的地点和时区信息。

#### func (Time) [Zone](https://github.com/golang/go/blob/master/src/time/time.go?name=release#823)

```go
func (t Time) Zone() (name string, offset int)
```

> Zone计算t所在的时区，返回该时区的规范名（如"CET"）和该时区相对于UTC的时间偏移量（单位秒）。

#### func (Time) [IsZero](https://github.com/golang/go/blob/master/src/time/time.go?name=release#243)

```go
func (t Time) IsZero() bool
```

IsZero报告t是否代表Time零值的时间点，January 1, year 1, 00:00:00 UTC。

==这个判断零点==

#### func (Time) [Local](https://github.com/golang/go/blob/master/src/time/time.go?name=release#796)

```go
func (t Time) Local() Time
```

> Local返回采用本地和本地时区，但指向同一时间点的Time。

#### func (Time) [UTC](https://github.com/golang/go/blob/master/src/time/time.go?name=release#790)

```go
func (t Time) UTC() Time
```

> UTC返回采用UTC和零时区，但指向同一时间点的Time。

#### func (Time) [In](https://github.com/golang/go/blob/master/src/time/time.go?name=release#804)

```go
func (t Time) In(loc *Location) Time
```

> In返回采用loc指定的地点和时区，但指向同一时间点的Time。==如果loc为nil会panic==。

该time的local的

#### func (Time) [Unix](https://github.com/golang/go/blob/master/src/time/time.go?name=release#830)

```go
func (t Time) Unix() int64
```

> Unix将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位秒）。

一般用这个搞随机值

#### func (Time) [UnixNano](https://github.com/golang/go/blob/master/src/time/time.go?name=release#838)

```go
func (t Time) UnixNano() int64
```

> UnixNano将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位纳秒）。如果纳秒为单位的unix时间超出了int64能表示的范围，结果是未定义的。注意这就意味着Time零值调用UnixNano方法的话，结果是未定义的。

#### func (Time) [Equal](https://github.com/golang/go/blob/master/src/time/time.go?name=release#74)

```go
func (t Time) Equal(u Time) bool
```

> 判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息。

比较时间是否相等，别`==`

#### func (Time) [Before](https://github.com/golang/go/blob/master/src/time/time.go?name=release#65)

```go
func (t Time) Before(u Time) bool
```

> 如果t代表的时间点在u之前，返回真；否则返回假。

#### func (Time) [After](https://github.com/golang/go/blob/master/src/time/time.go?name=release#60)

```go
func (t Time) After(u Time) bool
```

> 如果t代表的时间点在u之后，返回真；否则返回假。

这两个时间比先后的

**下面省略time.hour\minute.....**

#### func (Time) [Add](https://github.com/golang/go/blob/master/src/time/time.go?name=release#613)

```go
func (t Time) Add(d Duration) Time
```

Add返回时间点t+d。

#### func (Time) [AddDate](https://github.com/golang/go/blob/master/src/time/time.go?name=release#658)

```go
func (t Time) AddDate(years int, months int, days int) Time
```

> AddDate返回增加了给出的年份、月份和天数的时间点Time。例如，时间点January 1, 2011调用AddDate(-1, 2, 3)会返回March 4, 2010。
>
> AddDate会将结果规范化，类似Date函数的做法。因此，举个例子，给时间点October 31添加一个月，会生成时间点December 1。（从时间点November 31规范化而来）

time的加法

#### func (Time) [Sub](https://github.com/golang/go/blob/master/src/time/time.go?name=release#631)

```go
func (t Time) Sub(u Time) Duration
```

计算时间差

#### func (Time) [Round](https://github.com/golang/go/blob/master/src/time/time.go?name=release#1107)

```go
func (t Time) Round(d Duration) Time
```

其中，方法接收者 `t` 表示要进行舍入操作的时间值，参数 `d` 表示要舍入到的时间单位。注意，只有对 `Duration` 值的数值部分进行舍入，其时间单位部分不变。

#### func (Time) [Truncate](https://github.com/golang/go/blob/master/src/time/time.go?name=release#1096)

```go
func (t Time) Truncate(d Duration) Time
```

跟Round差不多，不过是做的截断

#### func (Time) [Format](https://github.com/golang/go/blob/master/src/time/format.go?name=release#414)

```go
func (t Time) Format(layout string) string
```

#### func [ParseDuration](https://github.com/golang/go/blob/master/src/time/format.go?name=release#1158)

```go
func ParseDuration(s string) (Duration, error)
```

```go
duration, err := time.ParseDuration("1h30m")
    if err != nil {
        fmt.Println("Error parsing duration:", err)
        return
    }
    fmt.Println(duration)
```

#### func [Since](https://github.com/golang/go/blob/master/src/time/time.go?name=release#646)

```
func Since(t Time) Duration
```

> Since返回从t到现在经过的时间，等价于time.Now().Sub(t)。

### type [Timer](https://github.tcom/golang/go/blob/master/src/time/sleep.go?name=release#45)

```
type Timer struct {
    C <-chan Time
    // 内含隐藏或非导出字段
}
```

Timer类型代表单次时间事件。当Timer到期时，当时的时间会被发送给C，除非Timer是被AfterFunc函数创建的。

`Timer`类型表示在未来的某个时间点发送一个时间值的时间。这通常用于程序的超时功能，或者基于时间的事件的调度和同步。当一个`Timer`触发时，它会发送一个事件，表示时间已经过去了。如果需要在未来的某个时间点触发一个事件，可以使用`time.After`函数。

#### func [NewTimer](https://github.com/golang/go/blob/master/src/time/sleep.go?name=release#61)

```
func NewTimer(d Duration) *Timer
```

> NewTimer创建一个Timer，它会在最少过去时间段d后到期，向其自身的C字段发送当时的时间。

新建timer，在至少持续时间`d`之后向其自己的通道发送当前时间。`Timer`将在发送到通道之前保持阻塞状态，因此如果`d`为零或负数，则将在调用`NewTimer`时立即向通道发送时间。

### type [Ticker](https://github.com/golang/go/blob/master/src/time/tick.go?name=release#11)

```
type Ticker struct {
    C <-chan Time // 周期性传递时间信息的通道
    // 内含隐藏或非导出字段
}
```

Ticker保管一个通道，并每隔一段时间向其传递"tick"。

#### func [NewTicker](https://github.com/golang/go/blob/master/src/time/tick.go?name=release#21)

```
func NewTicker(d Duration) *Ticker
```

NewTicker返回一个新的Ticker，该Ticker包含一个通道字段，并会每隔时间段d就向该通道发送当时的时间。它会调整时间间隔或者丢弃tick信息以适应反应慢的接收者。如果d<=0会panic。关闭该Ticker可以释放相关资源。

#### func (*Ticker) [Stop](https://github.com/golang/go/blob/master/src/time/tick.go?name=release#45)

```
func (t *Ticker) Stop()
```

Stop关闭一个Ticker。在关闭后，将不会发送更多的tick信息。Stop不会关闭通道t.C，以避免从该通道的读取不正确的成功。

区别：`Timer`是一次性定时器，它在指定的时间间隔之后只会触发一次，而且在触发之后，计时器就会停止计时，需要通过调用`Stop`方法手动停止定时器。可以使用`NewTimer`函数创建一个`Timer`实例。

`Ticker`则是周期性定时器，它会每隔指定的时间间隔触发一次，而且在触发后会自动重新开始计时。可以使用`NewTicker`函数创建一个`Ticker`实例。

可以通过调用`Ticker.C`和`Timer.C`方法获取一个通道，当定时器到期时，该通道会被激活，并发送一个时间值。通过监听该通道，可以实现定时触发事件的功能。

### func [Sleep](https://github.com/golang/go/blob/master/src/time/sleep.go?name=release#9)

```
func Sleep(d Duration)
```

Sleep阻塞当前go程至少d代表的时间段。d<=0时，Sleep会立刻返回。

用的很多，休眠当前的go程

### func [After](https://github.com/golang/go/blob/master/src/time/sleep.go?name=release#101)

```
func After(d Duration) <-chan Time
```

After会在另一线程经过时间段d后向返回值发送当时的时间。等价于NewTimer(d).C。

Example

### func [Tick](https://github.com/golang/go/blob/master/src/time/tick.go?name=release#51)

```
func Tick(d Duration) <-chan Time
```

Tick是NewTicker的封装，只提供对Ticker的通道的访问。如果不需要关闭Ticker，本函数就很方便。