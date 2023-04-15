# gva
本项目主要是管理系统CURD、权限管理、用户管理、列表、分页、上传下载、日志包封装、文档自动化等等功能
# 技术选型
1.后端采用golang框架gin，快速搭建基础restful风格API
2.前端项目采用VUE框架，构建基础页面
3.数据库采用Mysql
4.使用redis实现记录当前活跃用户的jwt令牌并实现多点登录限制
5.使用swagger构建自动化文档
6.使用fsnotify和viper实现json格式配置文件
7.使用logrus实现日志记录
8.使用gorm实现对数据库的基本操作

# 测试建议
clone项目以后，把db文件导入自己创建的库后，最好前往七牛云申请自己的空间地址，
替换掉项目中的七牛云公钥，私钥，仓名和默认url地址，登陆以后为最高权限，动api权限或者菜单权限均有可能导致数据错乱，系统无法使用。

golang项目存放于Server文件夹下，
Server内部static/config存放mysql相关配置。可以根据自己的mysql数据库名 用户名 密码修改对应配置
vue项目存放于Vue文件夹下
