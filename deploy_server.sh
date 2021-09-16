cd server
GOOS=linux GOARCH=amd64 go build
ssh root@47.104.220.230 "cat server/pid | xargs kill -9;"
scp server root@47.104.220.230:~/server/.
ssh root@47.104.220.230 "cd server;sh run.sh"
rm server
