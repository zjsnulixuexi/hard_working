# hard_working
vue+go+sqlite
src是前端部分，只上传了vue的src部分


Go函数实现部分
Main包作为程序的入口，执行
App包接受前端传输的json数据
Gorm包执行与数据库交互的功能，如对数据库中的数据进行增删改查
Router路由部分，前端传过来的json数据中存储着一个type为operate_type的数据，代表着               前端需要执行的功能，也对应着后端的相关函数
Py包里面的python文件由前端传给后端后，后端在服务器保存，供后续数据分析使用
Go.mod和go.sum是go语言的包管理工具和相关配置文件
Models是部分struct结构体，与数据库表对应

后端具体的执行流程：首先执行main.go文件后跳转到app.go文件，解析前端传过来的json数据后，找到operate_type字段，执行router包下的对应函数，而router对应函数会调用gorm包的相关函数，对数据库进行合理的操作，其中可能会执行部分python脚本
