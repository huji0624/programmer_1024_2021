cd data_generator;
GOOS=linux GOARCH=amd64 go build
scp data_generator root@47.104.220.230:~/data_generator/.
rm data_generator

cd ..

cd server
GOOS=linux GOARCH=amd64 go build
ssh root@47.104.220.230 "cat server/pid | xargs kill -9;"
scp server root@47.104.220.230:~/server/.
ssh root@47.104.220.230 "cd server;sh run.sh"
rm server

#cd ..

#cd h5
#yarn build
#tar czvf dist.tar.gz dist
#scp dist.tar.gz root@47.104.220.230:~/.
#ssh root@47.104.220.230 "tar xvf dist.tar.gz;rm -rf h5/dist;mv dist h5/.;"
