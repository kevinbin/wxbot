docker run -itd --name wechatbot -v `pwd`/config.json:/app/config.json wechatbot:latest
sleep 2
docker exec -it wechatbot tail -f -n 30 /app/run.log