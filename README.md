# flume_exporter
Prometheus exporter for flume.

To run it:

```bash
make build

./flume_exporter [flags]
```

Help on flags:
```bash
./flume_exporter --help
```

Configuration: config.yml
```
agents:
- name: "flume-agents"
  enabled: true
# multiple urls can be separated by ,  
  urls: ["http://localhost:36001/metrics"]   
```

### Using Docker
Default
```
docker run -d -p 9360:9360 zhijunwoo/flume_exporter:latest
```

Specified configuration
```
docker run -d \
    -p 9360:9360 \
    -v `pwd`/config.yml:/etc/flume_exporter/config.yml \
    -name flume_exporter \
    zhijunwoo/flume_exporter:latest
```

### monitoring metrics
#### Sources
- AppendAcceptedCount
- AppendBatchAcceptedCount
- AppendBatchReceivedCount
- AppendReceivedCount
- ChannelWriteFail
- EventAcceptedCount
- EventReadFail
- EventReceivedCount
- GenericProcessingFail
- KafkaCommitTimer
- KafkaEmptyCount
- KafkaEventGetTimer
- OpenConnectionCount

#### Channels
- ChannelCapacity
- ChannelSize
- CheckpointBackupWriteErrorCount
- CheckpointWriteErrorCount
- EventPutAttemptCount
- EventPutErrorCount
- EventPutSuccessCount
- EventTakeAttemptCount
- EventTakeErrorCount
- EventTakeSuccessCount
- KafkaCommitTimer
- KafkaEventGetTimer
- KafkaEventSendTimer
- Open
- RollbackCounter
- Unhealthy

- RollbackCount
- ChannelFillPercentage

#### Sinks
- BatchCompleteCount
- BatchEmptyCount
- BatchUnderflowCount
- ChannelReadFail
- ConnectionClosedCount
- ConnectionCreatedCount
- ConnectionFailedCount
- EventDrainAttemptCount
- EventDrainSuccessCount
- EventWriteFail
- KafkaEventSendTimer
- RollbackCount

#### Grafana Dashboard
Grafana Dashboard ID: 10736  
name: Flume Exporter Metrics Overview For Prometheus
For details of the dashboard please see [Flume Exporter Metrics](https://grafana.com/grafana/dashboards/10736)

### 新分支说明
- feature/read-flume-process
    主要作用：通过在Linux主机上执行`ps -ef | grep Dflume.monitoring.port`命令获取当前正在运行的有监控的flume进程，
    然后自动把他们监控起来，不需要每次启停一些程序后还要修改配置文件