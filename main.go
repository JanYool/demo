package main // main包 = 可执行程序入口

// ============================================
// 一、分组声明 import —— 一个好习惯
// ============================================
import (
	"errors" // error 类型专用包
	"fmt"    // 格式化 I/O
)

// ============================================
// 二、常量 const（编译期确定，程序运行期不可改）
// ============================================
const Pi = 3.1415926           // 常量可指定超高精度（比如 200 位小数），分配给 float32/float64 时自动截断
const AppName = "GoDemo"       // 字符串常量
const Pi32 float32 = 3.1415926 // 也可以显式指明类型

// ============================================
// 三、iota 枚举 —— Go 的"自动编号器"
// ============================================
const (
	StatusNew    = iota // 0，新 const 块从 0 开始
	StatusActive        // 1，隐式 = iota，每行 +1
	StatusDone          // 2
)

const v = iota // 遇到新的 const 关键字，iota 重置！这里 v == 0

const (
	a       = iota             // 0
	b       = "B"              // "B"（打断了 iota 自增，但计数器继续）
	c       = iota             // 2（iota 继续计数）
	d, e, f = iota, iota, iota // d=3, e=3, f=3（同一行 iota 值相同）
	g       = iota             // 4
)

// ============================================
// 四、包级别变量（只能用 var，不能用 :=）
// ============================================
var (
	GlobalCounter int = 0      // 大写开头 = 可导出（public），其他包能访问
	internalNote      = "私有变量" // 小写开头 = 不可导出（private），只能本包用
)

// ============================================
// 五、定义一个可供外部使用的结构体（后续演示零值）
// ============================================
type Demo struct {
	Name  string
	Count int
}

func main() {
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║   Go 语言基础 —— 02.2 知识Demo     ║")
	fmt.Println("╚══════════════════════════════════════╝")

	// ========================================
	// 知识点1：变量定义 —— var，类型放变量名后面
	// ========================================
	fmt.Println("\n━━━ 1. 变量定义 ━━━")

	var age int = 25             // 完整形式：var 名字 类型 = 值
	var name = "张三"              // 省略类型，编译器自动推导为 string
	var isStudent bool           // 只声明不赋值 → 零值 false
	var x, y, z int = 10, 20, 30 // 同时声明多个同类型变量
	var m, n = 3.14, "hello"     // 多变量，类型各自推导

	fmt.Println("age:", age, "| name:", name, "| isStudent:", isStudent)
	fmt.Println("x:", x, "y:", y, "z:", z)
	fmt.Println("m:", m, "| n:", n)

	// ========================================
	// 知识点2：简短声明 := —— 只能在函数内部用！
	// ========================================
	fmt.Println("\n━━━ 2. 简短声明 := ━━━")

	score := 98.5               // 自动推导 float64
	language := "Go"            // 自动推导 string
	a1, a2, a3 := 1, "two", 3.0 // 多变量同时简短声明

	fmt.Println("score:", score, "| language:", language)
	fmt.Println("a1:", a1, "a2:", a2, "a3:", a3)

	// ========================================
	// 知识点3：空白标识符 _ —— 丢弃不需要的值
	//     Go 要求声明的变量必须被使用，否则编译报错
	//     用 _ 可以"吃掉"不想要的值
	// ========================================
	fmt.Println("\n━━━ 3. 空白标识符 _ ━━━")

	_, phone := 404, "13800138000" // 404 被丢弃，只保留 phone
	fmt.Println("phone:", phone)

	// ========================================
	// 知识点4：内置基础类型全家桶
	// ========================================
	fmt.Println("\n━━━ 4. 内置基础类型 ━━━")

	// --- 布尔 ---
	var flag bool = true // bool: 只有 true / false，默认 false
	var enabled, disabled = true, false
	fmt.Println("bool:", flag, enabled, disabled)

	// --- 整数 ---
	var i1 int = 100  // int: 有符号整数（长度因平台而异）
	var i2 uint = 200 // uint: 无符号整数
	var i3 int8 = 127 // int8: -128~127
	var i4 byte = 65  // byte = uint8（0~255）
	var i5 rune = '中' // rune = int32，存 Unicode 码点

	fmt.Println("int:", i1, "| uint:", i2, "| int8:", i3)
	fmt.Println("byte:", i4, "→ 字符:", string(i4)) // 65 → 'A'
	fmt.Println("rune:", i5, "→ 字符:", string(i5)) // 20013 → '中'

	// --- 不同整数类型不能互操作！必须显式转换 ---
	var small int8 = 10
	var big int32 = int32(small) // ✅ 显式转换
	fmt.Println("int8→int32 显式转换:", big, "(不转换就编译报错)")

	// --- 浮点 ---
	var f32 float32 = 3.14                   // float32
	var f64 float64 = 3.14159265358979323846 // float64（浮点默认类型）
	f64auto := 2.71828                       // := 推导浮点 → 默认 float64
	fmt.Println("float32:", f32, "| float64:", f64, "| 自动推导:", f64auto)

	// --- 复数 ---
	var c64 complex64 = 5 + 5i    // complex64: 32位实部 + 32位虚部
	var c128 complex128 = 10 + 2i // complex128: 默认复数类型
	fmt.Println("complex64:", c64, "| complex128:", c128)

	// --- 字符串（UTF-8，不可变）---
	var s1 string = "Hello, 世界" // 双引号
	var s2 string = ""          // 空字符串（零值）
	s3 := "短声明字符串"
	fmt.Println("string:", s1, "| 空串:", s2, "|", s3)

	// ========================================
	// 知识点5：字符串不可变 → 想改怎么办？
	//     s[0] = 'x' ❌ 编译报错
	//     正确做法：string → []byte → 改 → string
	// ========================================
	fmt.Println("\n━━━ 5. 字符串不可变 & 修改方法 ━━━")

	original := "hello"
	raw := []byte(original) // string → []byte 切片
	raw[0] = 'H'            // 修改第一个字符
	modified := string(raw) // []byte → string
	fmt.Println("原始:", original, "→ 修改:", modified)

	// --- 字符串拼接用 + ---
	say := "Hello,"
	world := " Go世界"
	combined := say + world
	fmt.Println("拼接:", combined)

	// --- 切片方式"修改"字符串 ---
	s := "hello"
	s = "c" + s[1:] // 取 s[1:]（即 "ello"），前面拼上 "c"
	fmt.Println("切片式修改:", s)

	// --- 反引号：多行原始字符串，不转义，原样输出 ---
	multiLine := `这是第一行
这是第二行
\t\n 原样输出，不会被转义`
	fmt.Println("多行字符串:\n" + multiLine)

	// ========================================
	// 知识点6：error 类型 —— Go 处理错误的方式
	// ========================================
	fmt.Println("\n━━━ 6. error 类型 ━━━")

	err := errors.New("这是一个错误信息")
	if err != nil {
		fmt.Println("捕获到错误:", err)
	}

	// ========================================
	// 知识点7：常量与 iota 验证
	// ========================================
	fmt.Println("\n━━━ 7. 常量 / iota 枚举 ━━━")

	fmt.Println("Pi:", Pi, "| Pi32:", Pi32, "| AppName:", AppName)
	fmt.Println("StatusNew:", StatusNew, "StatusActive:", StatusActive, "StatusDone:", StatusDone)
	fmt.Println("iota 重置 v:", v)
	fmt.Println("iota 打断示例: a=", a, "b=", b, "c=", c, "d=", d, "e=", e, "f=", f, "g=", g)

	// ========================================
	// 知识点8：大写/小写 = 导出/私有
	// ========================================
	fmt.Println("\n━━━ 8. 命名规则 ━━━")
	fmt.Println("GlobalCounter（大写开头）= 可导出，其他包可访问")
	fmt.Println("internalNote（小写开头）= 私有，仅本包可访问:", internalNote)

	// ========================================
	// 知识点9：array —— 固定长度数组
	// ========================================
	fmt.Println("\n━━━ 9. array（固定数组）━━━")

	var arr1 [5]int // 声明长度为5的int数组，零值填充
	arr1[0] = 42    // 下标从0开始
	arr1[4] = 99
	fmt.Println("arr1:", arr1, "| 长度:", len(arr1))

	arr2 := [3]int{1, 2, 3}      // := 声明并初始化
	arr3 := [5]int{1, 2, 3}      // 前3个指定，其余为0
	arr4 := [...]int{10, 20, 30} // ... 让编译器自动算长度
	fmt.Println("arr2:", arr2, "| arr3:", arr3, "| arr4:", arr4, "len:", len(arr4))

	// --- 多维数组 ---
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("二维数组 matrix:", matrix)

	// --- 数组是值类型，传参会复制整个副本 ---
	arrCopy := arr2 // 完整拷贝
	arrCopy[0] = 999
	fmt.Println("arr2[0]:", arr2[0], "| arrCopy[0]:", arrCopy[0], "(值拷贝，互不影响)")

	// ========================================
	// 知识点10：slice —— 动态数组（引用类型）
	// ========================================
	fmt.Println("\n━━━ 10. slice（动态切片）━━━")

	// --- 创建 slice ---
	var sEmpty []int                // 声明空 slice（nil）
	slice1 := []byte{'a', 'b', 'c'} // 直接初始化
	slice2 := []int{1, 2, 3, 4, 5}  // := 方式
	fmt.Println("nil slice:", sEmpty, "len:", len(sEmpty), "(此时为 nil)")
	fmt.Println("slice1:", slice1, "| slice2:", slice2)

	// --- 从数组切出 slice ---
	arr := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	aSlice := arr[2:5] // arr[2] 到 arr[4]（不含 arr[5]），即 c,d,e
	bSlice := arr[3:5] // arr[3] 和 arr[4]，即 d,e
	fmt.Println("数组:", string(arr[:]))
	fmt.Println("aSlice = arr[2:5]:", string(aSlice))
	fmt.Println("bSlice = arr[3:5]:", string(bSlice))

	// --- slice 简便操作 ---
	fmt.Println("arr[:3]:", string(arr[:3])) // 等价 arr[0:3]，从头到第3个
	fmt.Println("arr[7:]:", string(arr[7:])) // 等价 arr[7:10]，从第7到最后
	fmt.Println("arr[:]:", string(arr[:]))   // 等价 arr[0:10]，全部

	// --- 从 slice 再切 slice ---
	middle := arr[3:7]      // d,e,f,g → len=4, cap=7
	sub := middle[1:3]      // e,f
	extended := middle[0:5] // slice 可在 cap 范围内扩展：d,e,f,g,h
	fmt.Println("middle:", string(middle), "| sub:", string(sub))
	fmt.Println("extended(利用cap扩展):", string(extended))

	// --- slice 是引用类型：改一处，处处变 ---
	ref1 := []int{10, 20, 30, 40, 50}
	ref2 := ref1 // 指向同一个底层数组
	ref2[0] = 999
	fmt.Println("ref1[0]:", ref1[0], "| ref2[0]:", ref2[0], "(引用类型，共享底层)")

	// --- len / cap / append / copy ---
	nums := []int{1, 2, 3}
	fmt.Println("初始 nums:", nums, "len:", len(nums), "cap:", cap(nums))

	nums = append(nums, 4, 5) // 追加元素
	fmt.Println("append 后:", nums, "len:", len(nums), "cap:", cap(nums))

	dest := make([]int, len(nums)) // 创建目标 slice
	copied := copy(dest, nums)     // 复制元素
	fmt.Println("copy 到 dest:", dest, "| 复制了", copied, "个元素")

	// --- 三参数 slice [low:high:max] ---
	base := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	three := base[2:4:7] // len=4-2=2, cap=7-2=5
	fmt.Println("三参数 slice:", three, "len:", len(three), "cap:", cap(three))

	// ========================================
	// 知识点11：map —— 字典/映射（引用类型）
	// ========================================
	fmt.Println("\n━━━ 11. map（字典）━━━")

	// --- 需要用 make 初始化 ---
	numbers := make(map[string]int) // key=string, value=int
	numbers["one"] = 1
	numbers["ten"] = 10
	numbers["three"] = 3
	fmt.Println("numbers:", numbers)
	fmt.Println("第三个数字:", numbers["three"])

	// --- 直接初始化赋值 ---
	rating := map[string]float32{
		"C":      5.0,
		"Go":     4.5,
		"Python": 4.5,
		"C++":    2.0,
	}
	fmt.Println("rating:", rating)

	// --- 判断 key 是否存在 ---
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# 评分:", csharpRating)
	} else {
		fmt.Println("C# 不在 rating 里 (ok=false)")
	}

	// --- delete 删除元素 ---
	delete(rating, "C")
	fmt.Println("删除 C 后:", rating)

	// --- map 是引用类型 ---
	m1 := make(map[string]string)
	m1["Hello"] = "Bonjour"
	m2 := m1 // 指向同一底层
	m2["Hello"] = "你好"
	fmt.Println("m1[\"Hello\"]:", m1["Hello"], "(引用类型，共享底层)")

	// --- len 返回 key 数量 ---
	fmt.Println("rating 中 key 数量:", len(rating))

	// ========================================
	// 知识点12：make vs new
	//     new(T) → 返回 *T 指针（零值填充）
	//     make → 只用于 slice/map/channel，返回初始化后的 T
	// ========================================
	fmt.Println("\n━━━ 12. make vs new ━━━")

	// new：返回指针，分配零值
	ptr := new(int)
	fmt.Println("new(int) 的值:", *ptr, "（零值）")

	ptr2 := new(Demo)
	fmt.Println("new(Demo):", ptr2, "Name 零值:", ptr2.Name, "Count 零值:", ptr2.Count)

	// make：只用于 slice/map/channel，返回初始化后的值（不是指针）
	ms := make([]int, 3, 5) // len=3, cap=5
	mm := make(map[string]int)
	mc := make(chan int, 1) // 缓冲为1的 channel
	fmt.Println("make([]int):", ms, "len:", len(ms), "cap:", cap(ms))
	fmt.Println("make(map):", mm)
	fmt.Println("make(chan):", mc)

	// ========================================
	// 知识点13：零值 —— 未初始化的默认值
	// ========================================
	fmt.Println("\n━━━ 13. 零值（Zero Values）━━━")

	var zeroInt int
	var zeroInt8 int8
	var zeroFloat float64
	var zeroBool bool
	var zeroStr string
	var zeroByte byte
	var zeroRune rune
	var zeroSlice []int
	var zeroMap map[string]int

	fmt.Println("int 零值:", zeroInt)
	fmt.Println("int8 零值:", zeroInt8)
	fmt.Println("float64 零值:", zeroFloat)
	fmt.Println("bool 零值:", zeroBool)
	fmt.Println("string 零值:", `"`+zeroStr+`"`, "(空字符串)")
	fmt.Println("byte 零值:", zeroByte)
	fmt.Println("rune 零值:", zeroRune)
	fmt.Println("slice 零值:", zeroSlice, "(nil? ", zeroSlice == nil, ")")
	fmt.Println("map 零值:", zeroMap, "(nil? ", zeroMap == nil, ")")

	fmt.Println("\n══════════════════════════════════════")
	fmt.Println("  演示完毕！对照 02.2.md 复习即可。")
	fmt.Println("══════════════════════════════════════")
}
