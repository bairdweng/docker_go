### Docker 实战

* 编译可执行文件

  ```js
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
  ```

* 制作镜像 将main放入Docker/app中

  ```js
  // 构建镜像 miaoyou_server 镜像名称
  docker build -t miaoyou_server .
  ```

  > Dockerfile用于配置镜像

  ```js
  FROM golang
  MAINTAINER  bairdweng
  WORKDIR /go/src/
  COPY . .
  EXPOSE 8200
  ENTRYPOINT ["./app/main"] #入口
  
  ```

* 查看镜像并制作容器 

  ```
  docker images
  ```

  ```
  REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
  miaoyou_server      latest              b09e7f70e0ec        52 minutes ago      828MB
  <none>              <none>              09aff806183b        58 minutes ago      828MB
  ```

* 制作容器 宿主机的端口8200映射到容器的8200

  ```
  docker run --rm -idt -p 8200:8200 miaoyou_server
  ```

* nginx配置 -v 本地路径:虚拟机niginx路径 -d

  ```
  docker run --name nginx -p 80:80 -v /Users/bairdweng/Desktop/servers/docker_go/Docker/nginx/api.bairdweng.com.conf:/etc/nginx/conf.d/api.bairdweng.com.conf -d nginx
  
  ```

  > api.bairdweng.com.conf 的配置

  ```js
  server {
     listen      80;
     server_name api.bairdweng.com;
     location / {
          proxy_pass http://docker.for.mac.host.internal:8200;
     }
  }
  ```

  