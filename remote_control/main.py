import paramiko

# 定义主机配置
configs = [
    ("47.104.220.230", "root", "hiKLD.624"),
    ("47.104.227.116", "root", "hiKLD.624")
]
# 保存连接 session
sessions = []


# 连接
def connect(config):
    client = paramiko.SSHClient()
    client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    client.connect(hostname=config[0], port=22, username=config[1], password=config[2])
    return client.get_transport().open_session()


def connect_all():
    sessions.clear()
    for config in configs:
        sessions.append(connect(config=config))


if __name__ == '__main__':
    while True:
        cmd = input("请输入开始命令:\n")
        #连接所有主机
        connect_all()
        #执行远程命令
        for session in sessions:
            session.exec_command(cmd)
