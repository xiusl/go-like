
import sys
import os

argv = sys.argv
if len(argv) <= 1:
    exit("waring: need a module name")

project = "go-like"
module = argv[1]
app_path = "../app/{}/service/".format(module)
internal_path = app_path + "internal/"

if len(argv) > 2:
    server = argv[2]
    if server == 'http':
        grpc_server = False
    if server == 'grpc':
        http_server = False

def check_path(path):
    if not os.path.exists(path):
        os.makedirs(path)
        print(path + "创建成功")

def read_template(file):
    temp = ""
    with open(file, 'r') as f:
        temp = f.read()
        temp = temp.replace("{project}", project)
        temp = temp.replace("{module}", module)
        temp = temp.replace("{Module}", module.capitalize())
    return temp

def gen_main():
    path = app_path + 'cmd/server'
    check_path(path)

    with open(path+'/main.go', "w") as f:
        f.write(read_template('main.txt'))
    with open(path+'/wire.go', "w") as f:
        f.write(read_template('wire.txt'))

def gen_configs():
    path = app_path + 'configs'
    check_path(path)

    file = '/configs.yaml'
    with open(path+file, "w") as f:
        f.write(read_template('configs.txt'))

def gen_biz():
    path = internal_path + 'biz'
    check_path(path)

    file = '/biz.go'
    with open(path+file, "w") as f:
        f.write("package biz\n\nimport \"github.com/google/wire\"\n\nvar ProviderSet = wire.NewSet()")

def gen_data():
    path = internal_path + 'data'
    check_path(path)

    file = '/data.go'
    with open(path+file, "w") as f:
        f.write("package data\n\nimport \"github.com/google/wire\"\n\nvar ProviderSet = wire.NewSet()")

def gen_conf():
    path = internal_path + 'conf'
    check_path(path)

    file = '/conf.proto'
    with open(path+file, "w") as f:
        f.write(read_template('conf.txt'))

def gen_server():
    path = internal_path + 'server'
    check_path(path)

    with open(path+'/server.go', "w") as f:
        f.write(read_template('server.txt'))

    with open(path+'/http.go', "w") as f:
        f.write(read_template('http.txt'))

    with open(path+'/grpc.go', "w") as f:
        f.write(read_template('grpc.txt'))

def gen_service():
    path = internal_path + 'service'
    check_path(path)

    with open(path+'/service.go', "w") as f:
        f.write(read_template('service.txt'))

def gen_api():
    path = '../api/{}/service/v1'.format(module)
    check_path(path)
    file = "{0}/{1}.proto".format(path, module)
    with open(file, "w") as f:
        f.write(read_template('api.txt'))


def gen_makefile():
    check_path(app_path)
    with open(app_path+'Makefile', "w") as f:
        f.write("include ../../../app_makefile")


def gen_proj():
    gen_main()
    gen_configs()
    gen_biz()
    gen_data()
    gen_conf()
    gen_server()
    gen_service()
    gen_makefile()
    gen_api()

from subprocess import run

if __name__ == '__main__':
    gen_proj()
    run("cd {}; make grpc; make proto;".format(app_path), shell=True)