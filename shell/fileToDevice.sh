#!/bin/bash

#./fileToDevice.sh 10.231.105.97 /home/yulichun/mediaServer/cabin/alert_alarm/ rtsp://10.170.160.18:8554

# ip
ip=$1

# path
path=$2

# rtsp地址
rtsp_url=$3

# 进入目录  
cd "$path" || exit

split_string="${path#*/mediaServer/}"
  
echo "$split_string"
  
# 遍历目录下的所有文件  
for file in *; do  
  # 检查文件扩展名是否为 ".264"  
  if [[ $file == *.264 ]]; then  
    # 打印文件名
#    echo "$file"
    data="{\"name\":\"${file}\",\"deviceType\":4,\"deviceProtocol\":1,\"rtspUrl\":\"${rtsp_url}/${split_string}${file}\",\"connectMode\":1}"
    url="http://"$ip"/cube/webapi/v1/device/add"
    echo "$data"
    echo "$url"
    response=$(curl -s -X POST -H "Content-Type: application/json" -d "$data" $url)
    echo "$response"
  fi
done

