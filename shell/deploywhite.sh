#!/bin/bash

#./deploywhite.sh 10.231.105.97 jinyu_3F_entry_time false

ip=$1

devicename=$2

en=$3

response=$(curl -X POST -H "Content-Type: application/json" -d "{\"pageNum\":1,\"name\":\"${devicename}\",\"pageSize\":\"1000\"}" "http://${ip}/cube/webapi/v1/device/channel/search")

  echo "response: $response"
# 使用jq提取uuid字段的值
uuids=$(echo "$response" | jq -r '.data.list[].uuid')



# 循环遍历uuid值并打印
for uuid in $uuids; do
  echo "$uuid"
  # 使用提取到的uuid发送新的POST请求
  data="{\"list\":[{\"channelUuid\":\"${uuid}\",\"algoType\":\"glider-faced-video-share\",\"config\":{\"faceDeploy\":{\"customSettings\":{\"mode\":1},\"enable\":${en},\"facePedBind\":false,\"groupUuids\":[\"PG-0d27ab7bd2bd476e926c\"],\"ruleCustomName\":\"人脸\",\"ruleId\":1}}}]}"
  
  
  curl -s -X POST -H "Content-Type: application/json" -d "$data" "http://${ip}/cube/webapi/v1/deploy/upsertByChannel"
#  response=$(curl -s -X POST -H "Content-Type: application/json" -d "$data" "http://${ip}/cube/webapi/v1/deploy/upsertByChannel")
#  echo "response: $response"
done
