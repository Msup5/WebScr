# WebScr 是一款访问URL进行截图的工具

# 项目简介

WebScr是一款使用Go语言开发的工具，它能批量访问URL并截取首页图的工具

项目地址：https://github.com/Msup5/WebScr

# 选项

```
参数:
  -h, --help  查看帮助
  -f          URL 文件
  -o          保存路径
  -t          线程, 默认5
```



# 使用说明

简单用法

```
WebScr.exe -u target.txt           访问target.txt文件内的URL并截图
WebScr.exe -u target.txt -t 20     访问target.txt文件内的URL并截图, 且设置线程为20
```

其他用法

```
WebScr.exe -u target.txt -o C:\image     访问target.txt文件内的URL并截图, 且图片保存路径为C:\image (自动创建image目录)
```

# 免责声明

本工具仅用于学习，严禁用于任何非法活动。使用本文所述技术前，请确保已获得目标系统所有者的明确授权。任何滥用信息造成的法律责任及后果均由使用者自行承担，作者不承担任何责任。
