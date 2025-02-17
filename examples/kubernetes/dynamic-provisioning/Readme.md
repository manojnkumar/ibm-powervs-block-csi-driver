# Dynamic Volume Provisioning
This example shows how to create a PowerVS volume and consume it from container dynamically.

## Prerequisites

1. Kubernetes 1.13+ (CSI 1.0).

2. The [powervs-ebs-csi-driver](https://github.com/ppc64le-cloud/powervs-csi-driver) is installed.

## Usage

1. Create a sample app along with the StorageClass and the PersistentVolumeClaim:
```
kubectl apply -f specs/
```

2. Validate the volume was created and `volumeHandle` contains an PowerVS volumeID:
```
kubectl describe pv
```

3. Validate the pod successfully wrote data to the volume:
```
kubectl exec -it app cat /data/out.txt
```

4. Cleanup resources:
```
kubectl delete -f specs/
```