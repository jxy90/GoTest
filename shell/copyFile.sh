#!/bin/bash  

#./copyFile.sh jinyu_3F_entry_time.264 ~/mediaServer/other/ 8

# 校验参数个数是否正确
if [ $# -ne 3 ]; then
  echo "Usage: $0 <source_file> <target_dir> <copy_count>"
  exit 1
fi

# 源文件名  
source_file=$1  
  
# 目标目录  
target_dir=$2  
  
# 循环复制文件并添加后缀  
for ((i=1; i<=${3:-10}; i++))  
do
  echo "${target_dir}${source_file%.*}_$i.${source_file##*.}" 
  cp "$source_file" "${target_dir}${source_file%.*}_$i.${source_file##*.}"  
done
