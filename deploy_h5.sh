cd h5
yarn build
tar czvf dist.tar.gz dist
scp dist.tar.gz root@47.104.220.230:~/.
ssh root@47.104.220.230 "tar xvf dist.tar.gz;rm -rf h5/dist;mv dist h5/.;"
