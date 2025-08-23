# Nginx on GKE with Static IP and Multi-Domain SSL

This guide describes how to deploy Nginx on a Google Kubernetes Engine (GKE) cluster, expose it with a regional static external IP, and manage SSL certificates for multiple domains using Kubernetes secrets.

---

## Background

In the original POC, Nginx was deployed on a GCP VM to route per-site and static requests to Cloudflare R2. This guide adapts that approach for a production-ready deployment on GKE, leveraging Kubernetes best practices for scalability, reliability, and maintainability.

---

## Prerequisites
- A GKE cluster (regional, recommended)
- `kubectl` and `gcloud` CLI tools configured
- Sufficient IAM permissions to create addresses, secrets, and deploy resources
- SSL certificates and private keys for all your domains

---

### Allocate a Regional Static External IP

Reserve a static IP in the same region as your GKE cluster:

```bash
gcloud compute addresses create nginx-static-ip-dev --region=europe-west1
```

---

### Prepare SSL Certificates as a Kubernetes Secret

Store all your PEM and key files in a single Kubernetes secret. This allows Nginx to serve multiple domains with SNI:

```bash
kubectl create secret generic nginx-certs-meepo-secret \
  --from-file=meepoadmin.site.pem \
  --from-file=meepoadmin.site.key \
  --from-file=meepo.site.pem \
  --from-file=meepo.site.key \
  --from-file=a9game.site.pem \
  --from-file=a9game.site.key \
  --from-file=ngambling.tech.pem \
  --from-file=ngambling.tech.key \
  -n meepo-dev
```
- **Best practice:** Name the files in the secret exactly as per the domain name.

---

### Deploy Nginx with Static IP and SSL

- Use a Kubernetes Deployment for Nginx, mounting your config and the `nginx-certs` secret as a volume at `/etc/nginx/cert`.
- Expose both port 80 (for health checks and HTTP-to-HTTPS redirect) and 443 (for HTTPS) in your Deployment and Service.
- Set loadbalancerIP to the static ip allocated. Please note: annotation with static ip address name might not work for some reason.

```yaml
spec:
  type: LoadBalancer
  loadBalancerIP: "34.76.0.26"
```

---

### DNS Configuration
- Point your domain(s) to the reserved static IP address.
- Example:
  - `A record: yourdomain.com -> <STATIC_IP>`

---

## References
- [Kubernetes: Exposing Services](https://kubernetes.io/docs/concepts/services-networking/service/)
- [GKE: Reserving Static IP Addresses](https://cloud.google.com/kubernetes-engine/docs/how-to/alias-ips#static)