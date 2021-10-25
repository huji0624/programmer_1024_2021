ansible pg -m copy -a "src=/root/match_control/data_generator dest=/root/data_generator"
ansible pg -m shell -a "chmod a+x data_generator"
