#参考博客https://gitee.com/lzsnmb/QiuBlog
# ginblog
gin+vue编写一个个人博客
1.安装gin框架 go get -u github.com/gin-gonic/gin
2.搭建基本框架目录
api(controller)，config(配置文件用goini)，middleware(中间间),routers(路由)，utils(工具)，
upload(上传下载)，web(前端)
获取goini依赖 go get gopkg.in/ini.v1
3.下载数据库工具
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

