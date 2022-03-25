
* 启动该项目前，请先执行 mysql 脚本，位置在 `script\create_tables.sql`
* 数据库 Addr User Pass 切换成您自己的，位置在 `dal\dal.go` 文件中  
* 运行以下命令 
````
    $ go mod vendor  
    $ go run . 


    curl http://localhost:8090/album?id=1 

    curl http://localhost:8090/album?id=23
````