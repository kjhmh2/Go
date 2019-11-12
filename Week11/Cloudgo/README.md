# 开发Web服务程序

### 框架

本次使用的web开发框架是Martini框架，Martini 是Go 语言的 Web 框架，使用 Go 的 net/http接口开发，类似 Sinatra 或者 Flask 之类的框架，也可使用自己的 DB 层、会话管理和模板。具有如下特性：

- 使用极其简单.
- 无侵入式的设计.
- 很好的与其他的Go语言包协同使用.
- 超赞的路径匹配和路由.
- 模块化的设计 - 容易插入功能件，也容易将其拔出来.
- 已有很多的中间件可以直接使用.
- 框架内已拥有很好的开箱即用的功能支持.
- 完全兼容http.HandlerFunc接口.

### 测试

- curl测试

  ```bash
  [kjhmh2@centos ~]$ curl -v http://localhost:1234
  * About to connect() to localhost port 1234 (#0)
  *   Trying ::1...
  * Connected to localhost (::1) port 1234 (#0)
  > GET / HTTP/1.1
  > User-Agent: curl/7.29.0
  > Host: localhost:1234
  > Accept: */*
  > 
  < HTTP/1.1 200 OK
  < Content-Type: text/html; charset=UTF-8
  < Date: Tue, 12 Nov 2019 13:23:16 GMT
  < Content-Length: 1539
  < 
  <!DOCTYPE html>
  <html lang = "en">
  <head>
  	<title>Jigsaw Puzzle</title>
  	<meta charset = "UTF-8">
  	<link rel = "icon" href = "image/icon.png">
  	<link rel = "stylesheet" type="text/css" href="jigsaw.css">
  	<script src = "js/jquery-3.3.1.min.js"></script>
  	<script src = "jigsaw.js"></script>
  </head>
  <body>
  	<h1>Jigsaw Puzzle</h1>
  	<div id = "result"></div>
  	<div id = "position">
  		<div id = "game_square"></div>
  			<div class = "common position1" id = "part1"></div>
  			<div class = "common position2" id = "part2"></div>
  			<div class = "common position3" id = "part3"></div>
  			<div class = "common position4" id = "part4"></div>
  			<div class = "common position5" id = "part5"></div>
  			<div class = "common position6" id = "part6"></div>
  			<div class = "common position7" id = "part7"></div>
  			<div class = "common position8" id = "part8"></div>
  			<div class = "common position9" id = "part9"></div>
  			<div class = "common position10" id = "part10"></div>
  			<div class = "common position11" id = "part11"></div>
  			<div class = "common position12" id = "part12"></div>
  			<div class = "common position13" id = "part13"></div>
  			<div class = "common position14" id = "part14"></div>
  			<div class = "common position15" id = "part15"></div>
  			<div class = "common position16" id = "part16"></div>
  		<div id = "time">Time:</div>
  		<div id = "time_shown">
  			<div id = "color"></div>
  		</div>
  		<button id = "start">Restart</button>
  		<button id = "hint">Hint</button>
  	</div>
  </body>
  * Connection #0 to host localhost left intact
  ```

- ab测试

  ```bash
  [kjhmh2@centos ~]$ ab -n 10000 -c 1000 http://localhost:1234/
  This is ApacheBench, Version 2.3 <$Revision: 1430300 $>
  Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
  Licensed to The Apache Software Foundation, http://www.apache.org/
  
  Benchmarking localhost (be patient)
  Completed 1000 requests
  Completed 2000 requests
  Completed 3000 requests
  Completed 4000 requests
  Completed 5000 requests
  Completed 6000 requests
  Completed 7000 requests
  Completed 8000 requests
  Completed 9000 requests
  Completed 10000 requests
  Finished 10000 requests
  
  
  Server Software:        
  Server Hostname:        localhost
  Server Port:            1234
  
  Document Path:          /
  Document Length:        1539 bytes
  
  Concurrency Level:      1000
  Time taken for tests:   7.895 seconds
  Complete requests:      10000
  Failed requests:        0
  Write errors:           0
  Total transferred:      16570000 bytes
  HTML transferred:       15390000 bytes
  Requests per second:    1266.60 [#/sec] (mean)
  Time per request:       789.518 [ms] (mean)
  Time per request:       0.790 [ms] (mean, across all concurrent requests)
  Transfer rate:          2049.56 [Kbytes/sec] received
  
  Connection Times (ms)
                min  mean[+/-sd] median   max
  Connect:        0  277 525.9     10    3017
  Processing:    18  258 368.2    165    2599
  Waiting:        2  205 376.7    103    2593
  Total:         40  535 690.0    213    3859
  
  Percentage of the requests served within a certain time (ms)
    50%    213
    66%    331
    75%   1087
    80%   1132
    90%   1254
    95%   1525
    98%   3223
    99%   3394
   100%   3859 (longest request)
  ```

  ab测试的参数说明如下：
  - -n 执行的请求数量

  - -c 并发请求个数

  - -t 测试所进行的最大秒数

  - -p 包含了需要POST的数据的文件

  - -T POST数据所使用的Content-type头信息

  - -k 启用HTTP KeepAlive功能，即在一个HTTP会话中执行多个请求，默认时，不启用KeepAlive功能

  - ...

    返回值包括：

  - Concurrency Level：并发数

  - Time taken for tests：完成所有请求总共花费的时间

  - Complete requests：成功请求的次数

  - Failed requests：失败请求的次数

  - Total transferred：总共传输的字节数

  - HTML transferred：实际页面传输的字节数

  - Requests per second：每秒请求数

  - Time per request: [ms] (mean)： 平均每个用户等待的时间

  - Time per request: [ms] (mean, across all concurrent requests) ：服务器处理的平均时间

  - Transfer rate：传输速率

