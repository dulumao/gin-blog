# 使用 alpine 创建纯净版 Go 镜像
# 编译时需设置： CGO_ENABLED=0
FROM alpine

# 定义环境变量
ENV APP_DIR=/gin-blog \
    APP_PORT=8080

# 设置镜像工作目录
WORKDIR $APP_DIR

# 拷贝文件
COPY gin-blog $APP_DIR
COPY conf/ $APP_DIR/conf

# 导出端口
EXPOSE $APP_PORT

#ENTRYPOINT ["./entrypoint.sh"]
CMD ["/gin-blog/gin-blog"]