ansible pg -m copy -a "src=/root/match_control/match_solver dest=/root/."
ansible pg -m shell -a "cd match_solver;ls | xargs -I % chmod a+x %"
ansible pg -m shell -a "cd match_solver;ls | xargs -I % mv % ../%"
