# WebScr 是一款访问URL进行截图的工具

# 项目简介

WebScr是一款使用Go语言开发的工具，它能批量访问URL并截取首页图的工具

项目地址：https://github.com/Msup5/WebScr

# 选项

```
参数:
  -h, --help  查看帮助
  -f          URL 文件路径
  -o          保存路径
  -t          线程, 默认5
  -r          重试次数, 默认3
  -wT         等待时间, 默认60s
  -eT         额外等待时间, 默认10s
  -re         获取 ip, title, status,server
```

# 使用说明
使用前请先安装Chrome浏览器
下载文件, cmd运行WebScr.exe 即可

简单用法

```
WebScr.exe -f target.txt           访问target.txt文件内的URL并截图
WebScr.exe -f target.txt -t 20     访问target.txt文件内的URL并截图, 且设置线程为20
```

其他用法

```
WebScr.exe -f target.txt -o C:\image     访问target.txt文件内的URL并截图, 且图片保存路径为C:\image (自动创建image目录)
WebScr.exe -f target.txt -o C:\image -t 20 -r 5 -wT 80 访问target.txt文件内的URL并截图, 且图片保存路径为C:\image，并设置线程为20, 重试次数5, 等待时间80s
WebScr.exe -f target.txt -o C:\image -t 20 -re true 访问target.txt文件内的URL并截图, 且图片保存路径为C:\image，并且获取目标 ip, title, status, server, 保存路径为results/results.csv
```

# 运行截图
![image](https://github.com/Msup5/WebScr/blob/main/docs/2025-03-08_11-22-01.png))

# 免责声明

本工具仅用于学习，严禁用于任何非法活动。使用本文所述技术前，请确保已获得目标系统所有者的明确授权。任何滥用信息造成的法律责任及后果均由使用者自行承担，作者不承担任何责任。

# 更新
## 2025 更新

更新了一个请求功能以及其他小改动

增加了一些功能
# 旧版本
无需下载整个文件, 存在已知BUG, 未修复且不再维护
https://github.com/Msup5/WebScr/releases/tag/1.0.3
