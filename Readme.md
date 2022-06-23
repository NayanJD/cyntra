

## GCP Setup

### Creating a dev cluster

```
gcloud container clusters create dev-cluster --num-nodes=2
```

### Delete the dev cluster

```
gcloud container clusters delete dev-cluster
```


## Interacting with connectors

If the cluster is deployed in K8s cluster, run below in separate terminal

```
kubectl port-forward service/connect-lb 8083:8083
```

Adding postgres source connector

```
curl -i -X POST -H "Accept:application/json" \
    -H  "Content-Type:application/json" http://localhost:8083/connectors/ \
    -d @source.json
```

Adding elasticsearch sink connector

```
curl -i -X POST -H "Accept:application/json" \
    -H  "Content-Type:application/json" http://localhost:8083/connectors/ \
    -d @es-sink.json
```

Deleting a connector is something went wrong

```
curl -i -X DELETE -H "Accept:application/json" localhost:8083/connectors/elastic-sink
```

Getting list of connectors

```
curl -H "Accept:application/json" localhost:8083/connectors/
```

Getting status of a connector 

```
curl -X GET http://localhost:8083/connectors/product-connector/status 
```