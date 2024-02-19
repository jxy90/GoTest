#!/bin/bash

 response=$(curl -X POST -H "Content-Type: application/json" -d '{"pageNum":1,"name":"fire","pageSize":"100"}' http://10.117.74.93/cube/webapi/v1/device/channel/search)

  echo "response: $response"
# 使用jq提取uuid字段的值
uuids=$(echo "$response" | jq -r '.data.list[].uuid')

# 循环遍历uuid值并打印
for uuid in $uuids; do
  echo "$uuid"
  # 使用提取到的uuid发送新的POST请求
  data="{\"list\":[{\"channelUuid\":\"$uuid\",\"algoType\":\"supersafety-megflow-object-nv-cu101\",\"config\":{\"algorithmDeploy\":{\"taskInfo\":{\"custom\":null,\"groupUuids\":[],\"name\":\"\",\"rulesParams\":[{\"areas\":[{\"areaId\":1,\"areaType\":\"POLYGON\",\"custom\":{},\"points\":[{\"x\":0.011450382,\"y\":0.024879841},{\"x\":0.011450382,\"y\":0.97709924},{\"x\":0.9872773,\"y\":0.97936106},{\"x\":0.9936387,\"y\":0.020356234}]}],\"enable\":true,\"eventType\":\"FIRE\",\"extendParams\":{\"cooldownDuration\":5,\"duration\":1,\"level\":\"ALARM_LEVEL\",\"targetMax\":1,\"targetMin\":0,\"targetTypes\":[\"FIRE\"],\"threshold\":{\"FIRE\":0.5}},\"labels\":{},\"masks\":[],\"ruleCustomName\":\"火焰报警\",\"ruleId\":1}]}}}}]}"
  response=$(curl -s -X POST -H "Content-Type: application/json" -d "$data" http://10.117.74.93/cube/webapi/v1/deploy/upsertByChannel)
  
  
  echo "response: $response"
done
