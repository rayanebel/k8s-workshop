# LAB 0 - Déploiement d'un cluster Kubernetes via kubeadm

## Connexion au serveur bastion
Connectez vous à vos serveurs Azure en passant par le bastion. 

## Déploiement
### Initialisation du master.

Pour pouvoir intéragir avec azure, Kubernetes à besoin d'avoir un ensemble d'informations notamment des credentials.

Ouvrez le fichier `cloud.conf` et complétez les informations manquantes :

```bash
{
    "cloud":"AzurePublicCloud",
    "tenantId": <ID_TENANT_AZURE>,
    "aadClientId": <SERVICE_PRINCIPAL_ID>,
    "aadClientSecret": <SERVICE_PRINCIPAL_SECRET>,
    "subscriptionId": <SUBSCRIPTION_ID>,
    "resourceGroup": <K8S_RESOURCE_GROUP>,
    "location": <K8S_REGION>,
    "subnetName": <K8S_AZURE_SUBNET>,
    "securityGroupName": <K8S_AZURE_SECURITYGROUP>,
    "vnetName": <K8S_AZURE_VNET>,
    "vnetResourceGroup": <K8S_AZURE_VNET_RG>,
    "routeTableName": <K8S_AZURE_ROUTETABLE>,
    "primaryAvailabilitySetName": <K8S_AZURE_AVAILIBILITYSET>,
    "cloudProviderBackoff": false,
    "useManagedIdentityExtension": false,
    "useInstanceMetadata": true
}
```

Pour plus de détails, consultez la page suivante : https://github.com/kubernetes/cloud-provider-azure/blob/master/docs/cloud-provider-config.md

Copiez ensuite ce fichier dans `/etc/kubernetes`

Editez ensuite, le fichier `kubeadm-master.yml` et remplacer tous les champs entre `<>` avec vos valeurs. 

```
apiEndpoint:
  advertiseAddress: <IP_MACHINE_VIRTUELLE>
  bindPort: 6443
apiVersion: kubeadm.k8s.io/v1alpha3
bootstrapTokens:
- groups:
  - system:bootstrappers:kubeadm:default-node-token
  token: abcdef.0123456789abcdef
  ttl: 24h0m0s
  usages:
  - signing
  - authentication
kind: InitConfiguration
nodeRegistration:
  criSocket: /var/run/dockershim.sock
  name: <NOM_DE_LA_MACHINE>
  taints:
  - effect: NoSchedule
    key: node-role.kubernetes.io/master
  kubeletExtraArgs:
    cloud-provider: "azure"
    cloud-config: "/etc/kubernetes/cloud.conf"
---
apiVersion: kubeadm.k8s.io/v1alpha3
auditPolicy:
  logDir: /var/log/kubernetes/audit
  logMaxAge: 2
  path: ""
certificatesDir: /etc/kubernetes/pki
clusterName: kubernetes
controlPlaneEndpoint: ""
etcd:
  local:
    dataDir: /var/lib/etcd
    image: ""
imageRepository: k8s.gcr.io
kind: ClusterConfiguration
kubernetesVersion: v1.12.3
apiServerExtraArgs:
  cloud-provider: "azure"
  cloud-config: "/etc/kubernetes/cloud.conf"
apiServerExtraVolumes:
- name: cloud
  hostPath: "/etc/kubernetes/cloud.conf"
  mountPath: "/etc/kubernetes/cloud.conf"
controllerManagerExtraArgs:
  cloud-provider: "azure"
  cloud-config: "/etc/kubernetes/cloud.conf"
  address: 0.0.0.0
controllerManagerExtraVolumes:
- name: cloud
  hostPath: "/etc/kubernetes/cloud.conf"
  mountPath: "/etc/kubernetes/cloud.conf"
schedulerExtraArgs:
  address: 0.0.0.0
networking:
  dnsDomain: cluster.local
  podSubnet: "10.244.0.0/16"
  serviceSubnet: 10.96.0.0/12
unifiedControlPlaneImage: ""
---
apiVersion: kubeproxy.config.k8s.io/v1alpha1
bindAddress: 0.0.0.0
clientConnection:
  acceptContentTypes: ""
  burst: 10
  contentType: application/vnd.kubernetes.protobuf
  kubeconfig: /var/lib/kube-proxy/kubeconfig.conf
  qps: 5
clusterCIDR: "10.244.0.0/16"
configSyncPeriod: 15m0s
conntrack:
  max: null
  maxPerCore: 32768
  min: 131072
  tcpCloseWaitTimeout: 1h0m0s
  tcpEstablishedTimeout: 24h0m0s
enableProfiling: false
healthzBindAddress: 0.0.0.0:10256
hostnameOverride: ""
iptables:
  masqueradeAll: false
  masqueradeBit: 14
  minSyncPeriod: 0s
  syncPeriod: 30s
ipvs:
  excludeCIDRs: null
  minSyncPeriod: 0s
  scheduler: ""
  syncPeriod: 30s
kind: KubeProxyConfiguration
metricsBindAddress: 127.0.0.1:10249
mode: ""
nodePortAddresses: null
oomScoreAdj: -999
portRange: ""
resourceContainer: /kube-proxy
udpIdleTimeout: 250ms
---
address: 0.0.0.0
apiVersion: kubelet.config.k8s.io/v1beta1
authentication:
  anonymous:
    enabled: false
  webhook:
    cacheTTL: 2m0s
    enabled: true
  x509:
    clientCAFile: /etc/kubernetes/pki/ca.crt
authorization:
  mode: Webhook
  webhook:
    cacheAuthorizedTTL: 5m0s
    cacheUnauthorizedTTL: 30s
cgroupDriver: cgroupfs
cgroupsPerQOS: true
clusterDNS:
- 10.96.0.10
clusterDomain: cluster.local
configMapAndSecretChangeDetectionStrategy: Watch
containerLogMaxFiles: 5
containerLogMaxSize: 10Mi
contentType: application/vnd.kubernetes.protobuf
cpuCFSQuota: true
cpuCFSQuotaPeriod: 100ms
cpuManagerPolicy: none
cpuManagerReconcilePeriod: 10s
enableControllerAttachDetach: true
enableDebuggingHandlers: true
enforceNodeAllocatable:
- pods
eventBurst: 10
eventRecordQPS: 5
evictionHard:
  imagefs.available: 15%
  memory.available: 100Mi
  nodefs.available: 10%
  nodefs.inodesFree: 5%
evictionPressureTransitionPeriod: 5m0s
failSwapOn: true
fileCheckFrequency: 20s
hairpinMode: promiscuous-bridge
healthzBindAddress: 127.0.0.1
healthzPort: 10248
httpCheckFrequency: 20s
imageGCHighThresholdPercent: 85
imageGCLowThresholdPercent: 80
imageMinimumGCAge: 2m0s
iptablesDropBit: 15
iptablesMasqueradeBit: 14
kind: KubeletConfiguration
kubeAPIBurst: 10
kubeAPIQPS: 5
makeIPTablesUtilChains: true
maxOpenFiles: 1000000
maxPods: 110
nodeLeaseDurationSeconds: 40
nodeStatusUpdateFrequency: 10s
oomScoreAdj: -999
podPidsLimit: -1
port: 10250
registryBurst: 10
registryPullQPS: 5
resolvConf: /etc/resolv.conf
rotateCertificates: true
runtimeRequestTimeout: 2m0s
serializeImagePulls: true
staticPodPath: /etc/kubernetes/manifests
streamingConnectionIdleTimeout: 4h0m0s
syncFrequency: 1m0s
volumeStatsAggPeriod: 1m0s
```

Enfin, initialisez votre master en exécutant la commande ci-dessous :

```bash
$ kubeadm init --config kubeadm-master.yml

Your Kubernetes master has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

[...]

You can now join any number of machines by running the following on each node
as root:

  kubeadm join X.X.X.X:6443 --token xxxxxxxxxxxxxx --discovery-token-ca-cert-hash sha256:xxxxxxxxxxxxxxxxx
```

Exécuter les commandes suivantes pour récupérer le fichier de configuration qui va vous permettre d'intéragir avec votre cluster kubernetes.

```bash
  $ mkdir -p $HOME/.kube
  $ sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  $ sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

### Initialisation des noeuds (workers)

Pour commencer récuperez le fichier `cloud.conf` et copier sur le sur tous vos noeuds dans `/etc/kubernetes/` _(cf étape précédente)_

Tout d'abord, editer le fichier `kubeadm-worker.yml` et ajouter les éléments manquants : 

```
---
apiEndpoint:
  advertiseAddress: <IP_API_K8S>
  bindPort: 6443
apiVersion: kubeadm.k8s.io/v1alpha3
caCertPath: /etc/kubernetes/pki/ca.crt
clusterName: kubernetes
discoveryFile: ""
discoveryTimeout: 5m0s
discoveryToken: abcdef.0123456789abcdef
discoveryTokenAPIServers:
- <IP_API_K8S>:6443
discoveryTokenUnsafeSkipCAVerification: true
kind: JoinConfiguration
nodeRegistration:
  criSocket: /var/run/dockershim.sock
  name: <NAME_OF_THE_SERVER>
  kubeletExtraArgs:
    cloud-provider: "azure"
    cloud-config: "/etc/kubernetes/cloud.conf"
tlsBootstrapToken: abcdef.0123456789abcdef
token: abcdef.0123456789abcdef
```

Pour joindre permettre à l'un de vos serveurs de joindre le cluster, il vous suffit d'executer la commande suivante : 

```bash
$ kubeadm join --config kubeadm-worker.yml

This node has joined the cluster:
* Certificate signing request was sent to apiserver and a response was received.
* The Kubelet was informed of the new secure connection details.

Run 'kubectl get nodes' on the master to see this node join the cluster.
```

Si vous rencontrez l'erreur suivante : 

```bash
[preflight] running pre-flight checks
	[WARNING RequiredIPVSKernelModulesAvailable]: the IPVS proxier will not be used, because the following required kernel modules are not loaded: [ip_vs_rr ip_vs_wrr ip_vs_sh ip_vs] or no builtin kernel ipvs support: map[ip_vs:{} ip_vs_rr:{} ip_vs_wrr:{} ip_vs_sh:{} nf_conntrack_ipv4:{}]

you can solve this problem with following methods:
 1. Run 'modprobe -- ' to load missing kernel modules;
1. Provide the missing builtin kernel ipvs support
```

Executer la commande suivante : 

```bash
$ for i in ip_vs ip_vs_rr ip_vs_wrr ip_vs_sh nf_conntrack_ipv4; do modprobe $i; done
```

### Vérifications

Pour vérifier l'état de votre cluster, retournez sur l'un des masters et exécuter la commande suivante : 

```bash
$ kubectl get nodes

NAME         STATUS     ROLES    AGE    VERSION
k8smaster1   NotReady   master   3h8m   v1.12.3
k8snode1     NotReady   <none>   55s    v1.12.3
k8snode2     NotReady   <none>   28s    v1.12.3
```

On voit ici que les noeuds ne sont pas ready car, nous n'avons pas encore configurer le plugin reseau (CNI).

### Mise en place de flannel

Pour la partie networking, nous allons flannel qui est le plugin le plus simple à mettre en place.

Pour ce faire, il vous suffit d'exécuter sur l'un des masters la commande suivante : 

```bash
$ kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml

clusterrole.rbac.authorization.k8s.io/flannel created
clusterrolebinding.rbac.authorization.k8s.io/flannel created
serviceaccount/flannel created
configmap/kube-flannel-cfg created
daemonset.extensions/kube-flannel-ds-amd64 created
daemonset.extensions/kube-flannel-ds-arm64 created
daemonset.extensions/kube-flannel-ds-arm created
daemonset.extensions/kube-flannel-ds-ppc64le created
daemonset.extensions/kube-flannel-ds-s390x created
```

Une fois flannel deployé, vos noeuds seront `Ready`

```bash
$ kubectl get nodes
NAME         STATUS   ROLES    AGE     VERSION
k8smaster1   Ready    master   3h13m   v1.12.3
k8snode1     Ready    <none>   6m4s    v1.12.3
k8snode2     Ready    <none>   5m37s   v1.12.3
```

Une fois vos noeuds prêts, les composants de kubernetes doivent se déployer automatiquement.

Pour vérifier vous pouvez exécuter la commande suivante : 

```
$ kubectl get pods -n kube-system

NAME                                 READY   STATUS    RESTARTS   AGE
coredns-576cbf47c7-5rsww             1/1     Running   0          9m36s
coredns-576cbf47c7-sq79t             1/1     Running   0          9m36s
etcd-k8smaster1                      1/1     Running   0          3h14m
kube-apiserver-k8smaster1            1/1     Running   0          3h15m
kube-controller-manager-k8smaster1   1/1     Running   0          3h14m
kube-flannel-ds-amd64-2g2zn          1/1     Running   0          4m5s
kube-flannel-ds-amd64-dfqzv          1/1     Running   0          4m5s
kube-flannel-ds-amd64-vkcjc          1/1     Running   0          4m5s
kube-proxy-dbnxt                     1/1     Running   0          8m8s
kube-proxy-nnmrq                     1/1     Running   0          3h15m
kube-proxy-tdxvv                     1/1     Running   0          8m34s
kube-scheduler-k8smaster1            1/1     Running   0          3h14m
```

### Troubleshoot

En cas de problème, voici les commandes utiles à connaitre pour vous aider dans votre démarche : 

```bash
$ kubectl get pods -n kube-system
$ kubectl logs pod <POD_NAME> -n <NAMESPACE>
$ systemctl status kubelet
$ journalctl --no-pager -xu kubelet
$ kubectl cluster-info
$ kubectl get componentstatuses
```

A partir de maintenant, vous avez un cluster Kubernetes opérationnel. Vous pouvez commencer à déployer vos applications.
