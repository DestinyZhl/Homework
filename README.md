# homework2.2 
1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 localhost/healthz 时，应返回200


# homework 1009模块三作业：
1.构建本地镜像。
2.编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。
3.将镜像推送至 Docker 官方镜像仓库。
4.通过 Docker 命令本地启动 httpserver。
5.通过 nsenter 进入容器查看IP配置。

作业需编写并提交 Dockerfile 及源代码。