#!/bin/bash

echo "正在启动依赖项..."
echo "------------------------------"
if systemctl is-active --quiet mysqld; then
    echo "Mysql 已运行。"
else
    echo "（警告！）（警告！）Mysql 服务未运行！（警告！）（警告！）"
fi

echo "------------------------------"
screen -dmS minio_session minio server ~/minio-data --console-address :9001
echo "MinIO 已运行。"

echo "------------------------------"   
screen -dmS etcd_session etcd
echo "etcd 已运行。"

echo "------------------------------"
echo "测试 Redis 中，请查看："
valkey-cli ping

echo "------------------------------"
cd ~/kafka_2.13-3.9.0
screen -dmS zookeeper bash -c "bin/zookeeper-server-start.sh config/zookeeper.properties"
screen -dmS kafka bash -c "bin/kafka-server-start.sh config/server.properties"
sleep 1.5

running_processes=$(ps aux | grep kafka_2.13-3.9.0 | wc -l)
if [ "$running_processes" -eq 3 ]; then
    echo "Kafka 已运行。"
else
    echo "（警告！）（警告！）启动 Kafka 时出现问题！（警告！）（警告！）"
fi

echo "------------------------------"
echo "所有服务已在后台启动。"
