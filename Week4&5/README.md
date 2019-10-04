# CLI 命令行实用程序开发基础

## 命令行准则

### 输入
应该允许输入来自以下两种方式：

 - 在命令行上指定的文件名。例如：
```bash
$ command input_file
```
 - 标准输入（stdin），缺省情况下为终端（也就是用户的键盘）。用户输入 Control-D（文件结束指示符）前输入的所有内容都成为 command 的输入。例如：

```bash
$ command
```
- 使用 shell 操作符“<”（重定向标准输入），也可将标准输入重定向为来自文件，这里，command 会读它的标准输入，不过 shell／内核已将其重定向，所以标准输入来自 input_file。如下所示：
```bash
$ command < input_file
```
- 使用 shell 操作符“|”（pipe）也可以使标准输入来自另一个程序的标准输出，如下所示：
```bash
$ other_command | command
```
### 输出
- 输出应该被写至标准输出，缺省情况下标准输出同样也是终端（也就是用户的屏幕）：

```bash
$ command
```
 - 同样，使用 shell 操作符“>”（重定向标准输出）可以将标准输出重定向至文件。这里，command 仍然写至它的标准输出，不过 shell／内核将其重定向，所以输出写至 output_file。
```bash
$ command > output_file
```
 - 或者，还是使用“|”操作符，command 的输出可以成为另一个程序的标准输入，shell／内核安排 command 的输出成为 other_command 的输入，如下所示
```bash
$ command | other_command
```
### 参数
 - 强制选项

  ​	`-sNumber`和`-eNumber`作为强制选项。
  ​	selpg 要求用户用两个命令行参数`-sNumber`（例如，`-s10`表示从第 10 页开始）和`-eNumber`（例如，`-e20`表示在第 20 页结束）指定要抽取的页面范围的起始页和结束页。selpg 对所给的页号进行合理性检查；换句话说，它会检查两个数字是否为有效的正整数以及结束页是否不小于起始页。这两个选项，`-sNumber`和`-eNumber`是强制性的，而且必须是命令行上在命令名 selpg 之后的头两个参数。比方说：

  ```bash
  $ selpg -s10 -e20 ...
  ```
- 可选选项一：
  `-lNumber`和`-f`作为可选选项
  `selpg`可以处理两种输入文本： 

  - 类型一：

    ​	该类文本的页行数固定。这是缺省类型，因此不必给出选项进行说明。也就是说，如果既没有给出`-lNumber`也没有给出`-f`选项，则`selpg`会理解为页有固定的长度，即是每页 72 行。 
    ​	选择 72 作为缺省值是因为在行打印机上这是很常见的页长度。这样做的意图是将最常见的命令用法作为缺省值，这样用户就不必输入多余的选项。该缺省值可以用`-lNumber`选项覆盖，如下所示：

    ```bash
    $ selpg -s10 -e20 -l66 ...
    ```

    这表明页有固定长度，每页为 66 行。

  - 类型二：

    ​	该类型文本的页由 ASCII 换页字符（十进制数值为 12，在 C 中用“\f”表示）定界。该格式与“每页行数固定”格式相比的好处在于，当每页的行数有很大不同而且文件有很多页时，该格式可以节省磁盘空间。在含有文本的行后面，类型二的页只需要一个字符——换页——就可以表示该页的结束。打印机会识别换页符并自动根据在新的页开始新行所需的行数移动打印头。

    ​	将这一点与类型一比较：在类型一中，文件必须包含 `pageLen - currentPageLen ` 个新的行以将文本移至下一页，在这里`pageLen`是固定的页大小而`currentPageLen`是当前页上实际文本行的数目。在此情况下，为了使打印头移至下一页的页首，打印机实际上必须打印许多新行。这在磁盘空间利用和打印机速度方面效率都很低（尽管实际的区别可能不太大）。

    ```bash
    $ selpg -s10 -e20 -l66 ...
    ```

    这表明页有固定长度，每页为 66 行。

- 可选选项二：

  `-dDestination`作为可选选项

  selpg 还允许用户使用“`dDestination`选项将选定的页直接发送至打印机。这里，`Destination`应该是 lp 命令“-d”选项（请参阅“man lp”）可接受的打印目的地名称。该目的地应该存在 ― selpg 不检查这一点。在运行了带“-d”选项的 selpg 命令后，若要验证该选项是否已生效，请运行命令“lpstat -t”。该命令应该显示添加到“Destination”打印队列的一项打印作业。如果当前有打印机连接至该目的地并且是启用的，则打印机应打印该输出。这一特性是用 
  `popen()` 系统调用实现的，该系统调用允许一个进程打开到另一个进程的管道，将管道用于输出或输入。在下面的示例中，我们打开到命令

  ```bash
  $ lp -dDestination
  ```

  的管道以便输出，并写至该管道而不是标准输出：

  ```bash
  selpg -s10 -e20 -dlp1
  ```

  该命令将选定的页作为打印作业发送至 lp1 打印目的地。

## 程序说明

### 定义结构体

在slepg.c文件中，结构体的定义如下所示：

```c
struct selpg_args
{
	int start_page;
	int end_page;
	char in_filename[BUFSIZ];
	int page_len; /* default value, can be overriden by "-l number" on command line */
	int page_type; /* 'l' for lines-delimited, 'f' for form-feed-delimited */
					/* default is 'l' */
	char print_dest[BUFSIZ];
};
```

我们参照其给出的一系列变量，定义新的结构体如下所示：

```go
type selgpArgs struct {
	startPage int                   // start page [-s10]
	endPage int                     // end page   [-e10]
	inputFile string                // input file
	pageLen int                     // page length
	pageType string                 // row number [-l10] or line break [-f]
	printDest string                // print destination
}
```

主要包括开始页、结束页、输入文件的名字、页的长度、分页方式、输出文件的名字

### 读取解析参数

参数处理上，我们使用pflag包来解析命令的参数。

首先，我们要在本地安装：

```bash
go get github.com/spf13/pflag
```

然后进行包的导入：

```go
import (
    "github.com/spf13/pflag" 
    // ...
)
```

再通过如下步骤进行参数值的绑定。

```go
pflag.IntVarP(&args.startPage, "start page", "s", 0, "Start page of file")
pflag.IntVarP(&args.endPage, "end page", "e", 0, "End page of file")
pflag.IntVarP(&args.pageLen, "page length", "l", 72, "lines in one page")
pflag.StringVarP(&args.pageType, "page type", "f", "l", "flag splits page")
pflag.StringVarP(&args.printDest, "print destination", "d", "", "name of printer")
```

再调用`Parse()`函数让pflag 对标识和参数进行解析

```go
pflag.Parse()
```

### 拟定提示信息

补充提示信息，在`Usage`中定义好帮助函数：

```go
pflag.Usage = func() {
    fmt.Println("----------------------------Help----------------------------")
    fmt.Println("--usage: selpg -s[startPage] -e[endPage] [optional] [filename]")
    fmt.Println("[optional parameters] -l: number of lines per page [default: 72].")
    fmt.Println("[optional parameters] -f: line break.")
    fmt.Println("[optional parameters] -l and -f are mutual exclusion.")
    fmt.Println("[optional parameters] -d: the destination of output.")
    fmt.Println("[filename] input file [default: input from the console].")
}
```

对于传进来的参数，我们也需要判断是否符合规范，比如数字是否合法、逻辑上是否正确等等。

```go
if (args.startPage <= 0 || args.endPage <= 0) {
    fmt.Println("[Error] page number should be positive")
    os.Exit(1)
} else if (args.startPage > args.endPage) {
    fmt.Println("[Error] end page should be greater than start page")
    os.Exit(2)
} else if (args.pageLen <= 0) {
    fmt.Println("[Error] page length should be positive")
    os.Exit(3)
} else if (args.pageType == "f" && args.pageLen != 72) {
    fmt.Println("[Error] -l and -f are mutual exclusion")
    os.Exit(4)
}
```

### 从标准输入或文件输出至标准输出或文件

要判断究竟是标准输入还是文件输入，我们就利用我们定义的`inputFile`变量即可：

```go
if args.inputFile != "" {
	// input from file
} else {
	// input from standard input
}
```

在判断是标准输出还是文件输出的时候，我们也利用同样的方式：

```go
if args.printDest == "" {
    // output to screen
} else {
    // output to file
}
```

然后，再根据`arg.pageType`所对应的方式，确定是按照行分页还是按照分页符分页。在两种不同的方式下，将输出分别对应输出到指定位置即可。

## 程序测试

测试输入文件每行为对应的行序号

### 输入测试

- 指定文件名输入

  ```bash
  selpg -s1 -e2 -l3 a.txt 
  ```
  
  ![1](img\1.png)
  
- 缺省情况下的终端输入：

  ```bash
  selpg -s1 -e1 -l3
  ```

  ![2](img\2.png)

- 重定向标准输入：

  ```bash
  selpg -s1 -e2 -l3 < a.txt 
  ```

  ![3](img\3.png)

- 使标准输入来自另一个程序的标准输出：

  ```bash
  cat a.txt | selpg -s1 -e2 -l4 
  ```

  ![4](img\4.png)

### 输出测试

- 输出在屏幕上在前面已经演示过，此处不再演示。


- 将标准输出重定向至文件：

  ```bash
  selpg -s1 -e2 -l3 a.txt > b.txt
  ```

  ![5](img\5.png)

- 使输出成为另一个程序的标准输入：

  我们首先写另外一个能够读取输入并输出的小程序getInput：

  ```go
  package main
  
  import (
  	"os"
  	"fmt"
  	"bufio"
  )
  
  func main() {
  	input := bufio.NewScanner(os.Stdin)
  	fmt.Println("Get text:")
  
  	for input.Scan() {
  		line := input.Text()
  		if line == "" {
  			break
  		}
  		fmt.Println(line)
  	}
  }
  ```

  将其放入工作目录并安装好，这时再进行测试：

  ```bash
  selpg -s1 -e2 -l3 a.txt | getInput
  ```

  ![6](img\6.png)

​		

### 错误输出测试

- 将错误信息打印在屏幕上：

  ```bash
  selpg --s
  ```

  ![7](img\7.png)

- 将错误重定向至文件

  ```bash
  selpg --s > error.txt
  ```

  ![8](img\8.png)

- 将标准输出和标准错误都重定向至不同的文件

  ```bash
  selpg -s5 -e1 -l4 > b.txt > error.txt
  ```

  ![9](img\9.png)

### 其他参数测试

- `-f`参数测试

  由于在txt文档中没有分页符\f，所以我们临时将\f换成空，如果遇见空行则视为一个新页，修改原文档如下所示：

  ```
  1
  2
  
  3
  4
  5
  6
  
  7
  8
  9
  
  10
  
  11
  12
  ...
  ```

  使用命令：

  ```bash
  selpg -s2 -e3 -f a.txt
  ```

  ![10](img\10.png)

- `-d`参数测试

  使用命令：

  ```bash
  selpg -s2 -e3 -l2 -d b.txt a.txt
  ```

  ![11](img\11.png)