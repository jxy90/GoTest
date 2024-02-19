#!/bin/bash
  
#./delAllDeploy.sh 10.231.105.97

ip=$1


response=$(curl -X POST -H "Content-Type: application/json" -d "{\"pageNum\":1,\"pageSize\":\"1000\"}" "http://${ip}/cube/webapi/v1/deploy/query")

  echo "response: $response"
# 使用jq提取uuid字段的值
uuids=$(echo "$response" | jq -r '.data.list[].channelUuid')



# 循环遍历uuid值并打印
for uuid in $uuids; do
  echo "$uuid"
  # 使用提取到的uuid发送新的POST请求
  data="{\"channelUuid\":\"${uuid}\"}"


  curl -s -X POST -H "Content-Type: application/json" -d "$data" "http://${ip}/cube/webapi/v1/deploy/deleteByChannel"
#  response=$(curl -s -X POST -H "Content-Type: application/json" -d "$data" "http://${ip}/cube/webapi/v1/deploy/upsertByChannel")
#  echo "response: $response"
done
