# flume_exporter
Prometheus exporter for flume

To run it:

```bash
./flume_exporter [flags]
```

Help on flags:
```bash
./flume_exporter --help
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

### Using Docker

```
docker run -d -p 9360:9360   ...
```