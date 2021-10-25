tar czvf solver.tar.gz match_solver
scp solver.tar.gz root@47.104.220.230:~/match_control/.
ssh root@47.104.220.230 "cd match_control;tar xvf solver.tar.gz;"
rm solver.tar.gz
