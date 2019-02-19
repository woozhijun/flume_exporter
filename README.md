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