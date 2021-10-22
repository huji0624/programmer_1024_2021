cd data_generator;
GOOS=linux GOARCH=amd64 go build
scp data_generator root@47.104.220.230:~/match_control/.
#scp data_generator debug@47.104.227.116:~/.
rm data_generator


#  debug@47.104.227.116
#  hiPG.624


