# debezium

A simple demonstration of how to implement Debezium and Apache Kafka connection using Golang

## How to run
```
$ docker ps 
$ docker exec -it e50170826eab bash -l
root@e50170826eab:/# psql -U admin-user

admin-user=# create database debit_card;
CREATE DATABASE

admin-user=# \c debit_card;
You are now connected to database "debit_card" as user "admin-user".

debit_card=# CREATE TABLE “debit_card” (
debit_card(#     "Id" SERIAL PRIMARY KEY NOT NULL,
debit_card(#     "Name" VARCHAR(100) NOT NULL,
debit_card(#     "status" VARCHAR(20) NOT NULL
debit_card(# );
CREATE TABLE

debit_card=# INSERT INTO "debit_card" ("Name", "Status")
VALUES ('TRANG', "active");
```
```
% go run cmd/main.go
2024/07/18 14:37:51 context.Background env development root dir
2024/07/18 14:37:51 Listening ...
2024/07/18 14:38:39 Message: postgres.public.debit_card[0]@6 {"before":null,"after":{"Id":7,"Name":"Laptop","StockQuantity":60},"source":{"version":"2.5.0.Final","connector":"postgresql","name":"postgres","ts_ms":1721288278238,"snapshot":"false","db":"debit_card","sequence":"[\"23939392\",\"23940328\"]","schema":"public","table":"debit_card","txId":572,"lsn":23940328,"xmin":null},"op":"c","ts_ms":1721288278668,"transaction":null}
```

## References
https://github.com/mehmetcantas/go-kafka-debezium

## Services logs
```
postgres                     | LOG:  starting logical decoding for slot "product_slot"
kafka_connect_with_debezium  | 2024-07-18 07:18:55,550 WARN   ||  [Producer clientId=connector-producer-debezium_connector-0] Error while fetching metadata with correlation id 6 : {postgres.public.debit_card=LEADER_NOT_AVAILABLE}   [org.apache.kafka.clients.NetworkClient]
kafka                        | [2024-07-18 07:18:55,553] INFO [Controller id=1] New topics: [HashSet(postgres.public.debit_card)], deleted topics: [HashSet()], new partition replica assignment [Set(TopicIdReplicaAssignment(postgres.public.debit_card,Some(O7AfmUxOSCuTlheknsdHpQ),Map(postgres.public.debit_card-0 -> ReplicaAssignment(replicas=1, addingReplicas=, removingReplicas=))))] (kafka.controller.KafkaController)
kafka                        | [2024-07-18 07:18:55,559] TRACE [Broker id=1] Received LeaderAndIsr request LeaderAndIsrPartitionState(topicName='postgres.public.debit_card', partitionIndex=0, controllerEpoch=1, leader=1, leaderEpoch=0, isr=[1], partitionEpoch=0, replicas=[1], addingReplicas=[], removingReplicas=[], isNew=true, leaderRecoveryState=0) correlation id 11 from controller 1 epoch 1 (state.change.logger)
kafka                        | [2024-07-18 07:18:55,572] TRACE [Broker id=1] Cached leader info UpdateMetadataPartitionState(topicName='postgres.public.debit_card', partitionIndex=0, controllerEpoch=1, leader=1, leaderEpoch=0, isr=[1], zkVersion=0, replicas=[1], offlineReplicas=[]) for partition postgres.public.debit_card-0 in response to UpdateMetadata request sent by controller 1 epoch 1 with correlation id 12 (state.change.logger)
```
