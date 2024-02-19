#!/bin/bash


devicename=$1


response="{\"pageNum\":1,\"name\":\"${devicename}\",\"pageSize\":\"1000\"}" 

echo $response


