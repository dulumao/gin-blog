# 使用 Gin 搭建的博客API

特性：
- `RESTful` API
- 支持配置文件
- `gorm` 集成
- 使用 `go.mod` 包管理
- `docker-compose` 一键部署
- `swagger` 自动生成文档

# 安装
**Golang 版本需1.11以上**

安装：
```bash
cd $GOPATH/src && git clone github.com/zengxianxue/gin-blog.git && cd gin-blog
```
依赖包：
```bash
go mod tidy
```
> 需要科学上网

编译：
```bash
./build.sh
```
# Docker 部署

启动容器：
```bash
docker-compose up -d
```
访问：`127.0.0.1:8080/docs/index.html` 输出`swagger`生成的文档

停止服务：
```bash
docker-compose down
```

# 参考资料
- [Docker参考](https://yeasy.gitbooks.io/docker_practice/introduction/)
- [docker-compose参考](https://yeasy.gitbooks.io/docker_practice/compose/)
- [博客参考：EDDYCJY](https://github.com/EDDYCJY/go-gin-example)
