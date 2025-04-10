# Deployment instruction in dev environment

`aggregator-gateway` already exists in dev environment to route the traffic to `dev` namespace. The microservices for meepo are in `meepo-dev` namespace and we need to create a dedicate gateway `meepo-gateway` in this namespace to route the traffic for meepo microservices. Please follow steps below.


## Create the tls certificate in kubernetes cluster

``` bash
# kubectl create secret tls meepoapi-xyz-tls --cert='meepoapi.xyz.crt' --key='meepoapi.xyz.key'
```

## Reserve static ip address for ingress traffic

```bash
# gcloud compute addresses create gke-meepo-dev-ingress-ip --region=europe-west1 --network-tier=STANDARD
# gcloud compute addresses describe gke-meepo-dev-ingress-ip --region=europe-west1 // to get the ip address allocated
```

## Apply the kube yaml files in deploy folder

```bash
# kubectl apply -f deploy/gateway-dev.yaml    // create the gateway
# kubectl apply -f deploy/httproute-dev.yaml  // create the httproute
```