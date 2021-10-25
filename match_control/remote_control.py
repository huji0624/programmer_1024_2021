# -*- coding: UTF-8 -*-
import time

import paramiko

# 定义主机配置
configs = [
]

# 保存clients
clients = []


# 连接
def connect(config):
    client = paramiko.SSHClient()
    client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    client.connect(hostname=config[0], port=22, username=config[1], password=config[2])
    return client


def connect_all():
    for config in configs:
        clients.append(connect(config=config))


def open_sessions():
    ret = []
    for client in clients:
        ret.append(client.get_transport().open_session())
    return ret

cmds = [
    #"python main.py",
    "./casino",
    #"java -jar cp.jar",
    "./dig1024",
    #"java -jar DuiShenMeDui.jar",
    #"./huang_dig",
    "ls",
    #"sh cs.sh",
    #"./lzy",
    #"./main",
    #"java -jar power.jar",
    #"./treasureHunt",
    #"python3 main.py",
    #"sh huangxin.sh",
    #"ls",
]

def nohup(cmd):
    return "nohup "+cmd+" >s.log 2>&1 &"
    #return "nohup "+cmd+" &"

if __name__ == '__main__':
    connect_all()
    cmd = raw_input("请输入开始命令:")
    if cmd=="go":
    	# 连接所有主机
    	sessions = open_sessions()
    	print(time.time())
    	# 执行远程命令
    	sessions[0].exec_command(nohup(cmds[0]))
    	sessions[1].exec_command(nohup(cmds[1]))
    	sessions[2].exec_command(nohup(cmds[2]))
    	print(time.time())
