# Go Web Server 
----------

## Function

### TCP
* client调用tree命令显示server的pwd
* 根据client提供路径下载文件到client的pwd
### UDP
* 基于UDP的授时服务

## HELP
### main
NAME:

      fdl - A Web Sever App

USAGE:

      GoWebServer [global options] command [command options] [arguments...]

COMMANDS:

      client   Run Client
   
      server   Run Server
   
      help, h  Shows a list of commands or help for one command
   

GLOBAL OPTIONS:

      --help, -h  show help
   
### client

NAME:

      GoWebServer client - Run Client

USAGE:

      GoWebServer client [command options] [arguments...]

OPTIONS:

      -p value  protocol tcp or udp
   
      -d value  Download file
   
      -l        List

      -t        time
   
### server
NAME:

      GoWebServer server - Run Server

USAGE:

      GoWebServer server [command options] [arguments...]

OPTIONS:

      -p value  protocol tcp or udp

## TODO
* 文件断点续传
* 文件多线程传输
* 使用协程管理udp的read、write 


