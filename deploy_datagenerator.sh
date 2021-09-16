cd data_generator;
GOOS=linux GOARCH=amd64 go build
scp data_generator root@47.104.220.230:~/data_generator/.
rm data_generator
