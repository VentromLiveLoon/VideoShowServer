



# 视频展示服务器(VideoShow)

## 项目介绍

本项目使用beego框架实现了一个简单的流视频浏览和收藏浏览

本项目的视频流的实现使用zlmediakit 服务器搭建流服务器，使用rtmp协议上传视频到zlmediakit服务器，再通过本项目开发的VideoShow  Web端来进行观看。



## 项目依赖

zlmediakit 	server 流媒体服务器 （如何安装访问：https://github.com/ZLMediaKit/ZLMediaKit）

go			go语言开发环境

mysql		 数据库

ffmpeg	       视频的推流

## 项目配置

### zlmediakit 服务器配置
了解zlmediakit访问：https://github.com/ZLMediaKit/ZLMediaKit

zlmediakit服务器的config.init的配置最好如下所示,(不了解最好不要随意配置)

```ini
; auto-generated by mINI class {

[api]
apiDebug=1
defaultSnap=./www/logo.png
downloadRoot=./www
secret=1
snapRoot=./www/snap/

[cluster]
origin_url=
retry_count=3
timeout_sec=15

[ffmpeg]
bin=/usr/bin/ffmpeg
cmd=%s -re -i %s -c:a aac -strict -2 -ar 44100 -ab 48k -c:v libx264 -f flv %s
log=./ffmpeg/ffmpeg.log
restart_sec=0
snap=%s -i %s -y -f mjpeg -frames:v 1 -an %s

[general]
check_nvidia_dev=1
enableVhost=0
enable_ffmpeg_log=0
flowThreshold=1024
maxStreamWaitMS=15000
mediaServerId=your_server_id
mergeWriteMS=0
resetWhenRePlay=1
streamNoneReaderDelayMS=86400000
unready_frame_cache=100
wait_add_track_ms=3000
wait_track_ready_ms=10000

[hls]
broadcastRecordTs=0
deleteDelaySec=10
fastRegister=0
fileBufSize=65536
segDelay=1
segDur=1
segKeep=0
segNum=0
segRetain=5

[hook]
alive_interval=10.0
enable=0
on_flow_report=
on_http_access=
on_play=
on_publish=
on_record_mp4=
on_record_ts=
on_rtp_server_timeout=
on_rtsp_auth=
on_rtsp_realm=
on_send_rtp_stopped=
on_server_exited=
on_server_keepalive=
on_server_started=
on_shell_login=
on_stream_changed=
on_stream_none_reader=
on_stream_not_found=
retry=1
retry_delay=3.0
stream_changed_schemas=rtsp/rtmp/fmp4/ts/hls/hls.fmp4
timeoutSec=10

[http]
allow_cross_domains=1
allow_ip_range=127.0.0.1
charSet=utf-8
dirMenu=1
forbidCacheSuffix=
forwarded_ip_header=
keepAliveSecond=3000
maxReqSize=40960
notFound=<html><head><title>404 Not Found</title></head><body bgcolor="white"><center><h1>您访问的资源不存在！</h1></center><hr><center>ZLMediaKit(git hash:87cb488/2024-02-19T11:54:13+08:00,branch:master,build time:2024-02-19T12:04:37)</center></body></html>
port=8090
rootPath=./www
sendBufSize=65536
sslport=4430
virtualPath=

[multicast]
addrMax=239.255.255.255
addrMin=239.0.0.0
udpTTL=64

[protocol]
add_mute_audio=1
auto_close=0
continue_push_ms=15000
enable_audio=1
enable_fmp4=1
enable_hls=1
enable_hls_fmp4=0
enable_mp4=0
enable_rtmp=1
enable_rtsp=1
enable_ts=1
fmp4_demand=0
hls_demand=0
hls_save_path=./www
modify_stamp=2
mp4_as_player=0
mp4_max_second=3600
mp4_save_path=./www
paced_sender_ms=0
rtmp_demand=0
rtsp_demand=0
ts_demand=0

[record]
appName=record
fastStart=0
fileBufSize=65536
fileRepeat=0
sampleMS=500

[rtc]
externIP=
max_bitrate=0
min_bitrate=0
port=8000
preferredCodecA=PCMU,PCMA,opus,mpeg4-generic
preferredCodecV=H264,H265,AV1,VP9,VP8
rembBitRate=0
start_bitrate=0
tcpPort=8000
timeoutSec=15

[rtmp]
directProxy=1
enhanced=0
handshakeSecond=1500
keepAliveSecond=1500
port=1935
sslport=0

[rtp]
audioMtuSize=600
h264_stap_a=1
lowLatency=0
rtpMaxSize=10
videoMtuSize=1400

[rtp_proxy]
dumpDir=
gop_cache=1
h264_pt=98
h265_pt=99
opus_pt=100
port=10000
port_range=30000-35000
ps_pt=96
timeoutSec=15

[rtsp]
authBasic=0
directProxy=1
handshakeSecond=15
keepAliveSecond=15
lowLatency=0
port=5540
rtpTransportType=-1
sslport=0

[shell]
maxReqSize=1024
port=0

[srt]
latencyMul=4
pktBufSize=8192
port=9000
timeoutSec=5

; } ---

```



### VideoShow服务端配置

 在这里完成数据库的配置(conf/app.conf)

```ini
#项目配置
#[videoshow]
appname = VideoShow
httpport = 8080
runmode = dev
sessionon = true
#并发上传视频数量
routine_num = 100

#mysql的一些配置
#[mysql]
db_driver = mysql
db_host = 127.0.0.1
db_port = 3306

# 配置你的mysql的用户名和密码就可以了
db_user = $username
db_password = $password
db_name = video_show
db_charset = utf8

#zlmediakit服务器的相关设置,(不了解最好不要配置)
#[zlmediakit服务器的相关设置]
#上传视频流的ip地址
upload_zl_media_kit_media_server_ip = rtmp://127.0.0.1
#下载视频流的ip地址
download_zl_media_kit_media_server_ip = http://127.0.0.1:8090
#视频流的访问路径
zl_media_kit_data_path = /record/
#视频流的格式
video_type = hls_delay.m3u8


```

## 数据库加载

```mysql
drop database video_show;
create database video_show;
-- 本项目推荐的数据库名称
use video_show;

-- 视频流的名称
drop table video_path;
create table video_path(
    id bigint unsigned primary key auto_increment,
    name varchar(256) not null
);
-- 用户表
drop table user;
create table user(
    id int unsigned primary key auto_increment,
    name varchar(256) unique not null,
    password varchar(256) not null,
    -- 用户浏览视频流的id位置
    viewposition bigint unsigned
);
-- 用户的收藏表
drop table collect;
create table collect(
    id bigint unsigned primary key auto_increment,
    -- 用户id
    userid int unsigned  not null,
    -- 视频id
    videoid bigint unsigned  not null,
    index userid_idx (userid),
    index videoid_idx (videoid)
);
```



## 项目编译

```bash
go mod tidy && go build -o VideoShow main.go
```

## 启动ZLMediakit Server

```
cd zlmediaserver/ && ./MediaServer
```

## 上传视频流

```bash
./VideoShow scan /path/to/videos/directory
```

注意：扫描的目录的视频流不要太多，根据系统的性能扫描，还有推流的过程不显示，请不要关闭计算机。。。



## 项目运行

```
VideoShow
```



## 项目访问

```
http://127.0.0.1:8080
```



## 关于作者

```
#   #                   #                            
#   #                   #                            
#   #                   #                            
 # #    ####   # ###   ####   # ##   ####   # # ##   
 # #   #    #  ##   #   #      #    #    #  ## #  #  
 # #   ######  #    #   #      #    #    #  #  #  #  
  #    #       #    #   #      #    #    #  #  #  #  
  #    #    #  #    #   #  #   #    #    #  #  #  #  
  #     ####   #    #    ##    #     ####   #  #  #  
                                                     
            									-----------------------开发者::Ventrom_Live_Loon (QQ::2320653274)
                                                 				-------喜欢不要忘记给一个Start ！！:)
```

