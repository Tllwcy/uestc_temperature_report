# uestc_temperature_report

**电子科技大学（UESTC）每日体温填报**

**本项目仅适用于微信小程序 ->  "uestc学生情况报送“**

**项目通过 Github 的 Actions 实现每日定时自动填报体温**





## **💭前言**

辅导员每日多次督促体温打卡，催生了本项目

平时忙起来会有忘记上报的情况，于是撸出了这个项目，实现自动打卡





## **🌀简介**

+ 项目需要获取小程序的Cookie来实现自动体温打卡，一般采用客户端抓包的方式获取，获取方式请自行google
+ 支持多账号体温上报





## 🔨**使用方式**

1. Fork仓库
   + 点击右上角的`Fork`，将仓库Fork到自己的账号下
2. 获取Cookie
   + 使用Wireshark或Charles等抓包工具，获取cookie
3. 添加 Cookie 至 Secrets
   + 回到自己的项目页面，依次点击`Settings`-->`Secrets`-->`New secret`
   + `Name`中填入`COOKIE`，将抓包到的`Cookie`粘贴到`Value`中，点击`Add secret`添加
     + Name的框中只能填`COOKIE`，不要填其他
     + 如果有多个 Cookie，不同账号的`Cookie`值之间用`#`分隔，如：`Cookie1#Cookie2#Cookie3`
4. 启用Action
   + 回到自己的项目页面，点击上方的`Actions`，再点击左侧的`uestc_temperature_report`，再点击`Run workflow`

以上，项目部署完毕

+ 项目会在每日凌晨 0:30 左右进行自动上报，也可重新创建Action手动触发上报
+ 在`Actions`页面点击`uestc_temperature_report`-->`build`-->`Run Main`查看运行日志





## **注意！！！**

​	本项目仅用于交流学习目的，不会对您的任何损失负责，包括但不限于帐号异常，辅导员请喝茶，健康码泛黄（红），考研失败，第三次世界大战等
