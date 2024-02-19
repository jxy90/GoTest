#!/bin/bash  

#./addDevice.sh 10.231.105.97 rtsp://10.231.105.97:18554/other/jinyu_3F_entry_time.264 8

# ip
ip=$1

# rtsp地址  
rtsp_url=$2

# 分割字符串并打印结果
base_name=${rtsp_url%.*}
filename=$(basename "$rtsp_url")  
extension=${filename##*.}

echo "$base_name"
echo "$extension"

# 循环请求添加设备
for ((i=1; i<=${3:-10}; i++))
do

        data="{\"name\":\"${filename%.*}${i}\",\"deviceType\":4,\"deviceProtocol\":1,\"rtspUrl\":\"${base_name}_${i}.${extension}\",\"connectMode\":1}"
        url="http://"$ip"/cube/webapi/v1/device/add"
        echo "$data"
        echo "$url"
        response=$(curl -s -X POST -H "Content-Type: application/json" -d "$data" $url)
        echo "$response"
done
