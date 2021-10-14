import time

import paramiko

# 定义主机配置
configs = [
    ("47.104.220.230", "root", "hiKLD.624"),
    ("47.104.227.116", "root", "hiKLD.624")
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


if __name__ == '__main__':
    connect_all()
    while True:
        cmd = input("请输入开始命令:\n")
        # 连接所有主机
        sessions = open_sessions()
        print(time.time())
        # 执行远程命令
        for session in sessions:
            session.exec_command(cmd)
        print(time.time())
