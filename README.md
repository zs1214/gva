# gva
本项目主要是管理系统CURD、权限管理、用户管理、列表、分页、上传下载、日志包封装、文档自动化等等功能
应用技术：Gin + JWT  鉴权 + 动态路由 +Redis + MySQL + Zap
1.后端采用golang框架gin，快速搭建基础restful风格API
2.前端项目采用VUE框架，构建基础页面
3.数据库采用Mysql
4.使用redis实现记录当前活跃用户的jwt令牌并实现多点登录限制
5.使用swagger构建自动化文档
6.使用fsnotify和viper实现json格式配置文件
7.使用logrus实现日志记录
8.使用gorm实现对数据库的基本操作

