
#!/ bin / bash
# 打包到docker指定的目录
eval 'cd `dirname $0`'
eval 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./docker/miaoyou_server/app ./main/main.go'
echo "编译成功..."
eval 'cd ./docker/'
echo "开始执行docker-compose"
eval 'docker-compose stop && docker rm miaoyou_server_miaoyou_server_1'
eval 'docker-compose up -d --build'
