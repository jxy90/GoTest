#!/bin/bash
  echo "输入清空stream环境的ip:"
  read ip
  echo "你输入的值是：$ip"

  response=$(curl -X GET -H "Content-Type: application/json"  http://$ip:8897/api/v1/streams?pageSize=1000)

  echo "response: $response"

# 使用jq提取id字段的值
ids=$(echo "$response" | jq -r '.streams[].taskInfo.id')
echo "ids:$ids"

# 循环获取每个id
for id in $ids; do
    echo "ID: $id"

  # 使用提取到的uuid发送新的POST请求
  url="http://$ip:8897/api/v1/streams/$id"
  response=$(curl -X DELETE -H "Content-Type: application/json" url)


  echo "response: $response"
done