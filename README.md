# WebScr 是一款访问URL进行截图的工具

# 项目简介

WebScr是一款基于Go语言开发的工具，它能批量访问URL并截取首页图的工具

项目地址：https://github.com/Msup5/WebScr

# 选项

```
参数:
  -h, --help  	查看帮助
  -f          	URL 文件路径
  -o          	保存路径
  -t          	线程, 默认5
  -r          	重试次数, 默认3
  -wT         	等待时间, 默认60s
  -eT         	额外等待时间, 默认10s
  -re         	获取 ip, title, status,server, 默认false
  --html-output 输出html文件
```

# 使用说明
使用前请先安装Chrome浏览器

下载文件 zip 进行解压, cmd运行WebScr.exe 即可

简单用法

```
WebScr.exe -f url.txt                              访问url.txt文件内的URL并截图
WebScr.exe -f url.txt -t 20                        访问url.txt文件内的URL并截图, 且设置线程为20
WebScr.exe -f url.txt --html-output example.html   将所有结果输出为html文件（推荐）
```

其他用法

```
WebScr.exe -f url.txt -o C:\image                     访问target.txt文件内的URL并截图, 且图片保存路径为C:\image (自动创建image目录)
WebScr.exe -f url.txt -o C:\image -t 20 -r 5 -wT 80   访问url.txt文件内的URL并截图, 且图片保存路径为C:\image，并设置线程为20, 重试次数5, 等待时间80s
WebScr.exe -f url.txt -o C:\image -t 20 -re true      访问url.txt文件内的URL并截图, 且图片保存路径为C:\image，并且获取目标 ip, title, status, server, 保存路径为results/results.csv
```

# 运行截图

```
WebScr.exe -f url.txt  访问 url.txt 文件并截图
```

![image](https://github.com/Msup5/WebScr/blob/main/docs/2025-03-10_20-19-09.png)

```
WebScr.exe -f url.txt --html-output example.html  将所有结果输出成 html 文件
```

![image](https://github.com/Msup5/WebScr/blob/main/docs/2025-03-10_19-18-30.png)

![image](https://github.com/Msup5/WebScr/blob/main/docs/2025-03-10_20-07-00.png)

```
WebScr.exe -f url.txt -re true  获取 ip, title, status,server 并输出csv文件
```

![image](https://github.com/Msup5/WebScr/blob/main/docs/2025-03-10_20-10-25.png)

# 免责声明

本工具仅用于学习，严禁用于任何非法活动。使用本文所述技术前，请确保已获得目标系统所有者的明确授权。任何滥用信息造成的法律

责任及后果均由使用者自行承担，作者不承担任何责任。

# 更新
## 2025 更新
- 2025/3/11: 增加了跳过证书验证, 修复了一些 Bug
- 2025/3/10: 增加了输出 html 文件功能

- 2025/3/9: 增加了一个请求功能和小改动

- 2025/3/8: 增加了一些功能

- 2025/3/5: 针对 1.0.3 版本进行一次大更新, 增加了更多的功能
# 旧版本
## v1.0.3
无需下载整个文件, 更简便。存在已知BUG, 未修复且不再维护
https://github.com/Msup5/WebScr/releases/tag/1.0.3
