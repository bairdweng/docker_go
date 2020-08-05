
#!/ bin / bash
eval 'cd `dirname $0`'
# 编译
eval 'docker build -t bairdweng/miaoyou_server .'
# 提交
eval 'docker commit miaoyou_server bairdweng/miaoyou_server'
# 推送
eval 'docker push bairdweng/miaoyou_server'