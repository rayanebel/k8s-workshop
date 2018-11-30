# LAB 1 - Déploiement de votre première application

Dans ce lab, nous allons déployer notre première application sur kubernetes.

**Objets Kubernetes** : Pod

Pour ce faire, connectez vous sur le serveur ou se trouve votre configuration kubernetes vous permettant d'intéragir avec l'API. Par défaut, ce fichier se trouve sur les masters.

## Déploiement de l'application

Pour déployer votre première application, il vous suffit d'exécuter les commandes suivantes : 

```bash
cd Lab-1/
kubectl create namespace k8s-lab-1
kubectl apply -f cookie-app-pod.yml -n k8s-lab-1
```

Vérifier ensuite que votre application fonctionne en exécutant la commande suivante : 

```bash
kubectl get pod -n k8s-lab-1
NAME         READY   STATUS    RESTARTS   AGE
cookie-app   1/1     Running   0          54s
```

Afficher plus de détails sur votre déploiement vous pouvez exécuter la commande suivante : 

```bash
kubectl describe pod cookie-app -n k8s-lab-1

Name:               cookie-app
Namespace:          default
Priority:           0
PriorityClassName:  <none>
Node:               k8snode2/10.0.0.5
Start Time:         Fri, 30 Nov 2018 16:38:17 +0000
Labels:             app=cookie-app
Annotations:        kubectl.kubernetes.io/last-applied-configuration:
                      {"apiVersion":"v1","kind":"Pod","metadata":{"annotations":{},"labels":{"app":"cookie-app"},"name":"cookie-app","namespace":"default"},"spe...
Status:             Pending
IP:
Containers:
  cookie-app:
    Container ID:
    Image:          quay.io/coreos/example-app:v1.0
    Image ID:
    Port:           80/TCP
    Host Port:      0/TCP
    State:          Waiting
      Reason:       ContainerCreating
    Ready:          False
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-rplpx (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             False
  ContainersReady   False
  PodScheduled      True
Volumes:
  default-token-rplpx:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-rplpx
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute for 300s
                 node.kubernetes.io/unreachable:NoExecute for 300s
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  41s   default-scheduler  Successfully assigned default/cookie-app to k8snode2
  Normal  Pulling    39s   kubelet, k8snode2  pulling image "quay.io/coreos/example-app:v1.0"
  Normal  Pulled     3s    kubelet, k8snode2  Successfully pulled image "quay.io/coreos/example-app:v1.0"
```

Connectez-vous ensuite à votre conteneur en utilisant la commande suivante : 

```
$ kubectl exec -it cookie-app -n k8s-lab-1 -- bash
root@cookie-app:/# hostname
cookie-app
root@cookie-app:/#
```

Etant donné que notre application est une application web, nous pouvons vérifier son bon fonctionnement en executant une commande `curl` depuis un autre conteneur

```bash
$ kubectl get pod -o wide -n k8s-lab-1

NAME         READY   STATUS    RESTARTS   AGE   IP           NODE       NOMINATED NODE
cookie-app   1/1     Running   0          45m   10.244.3.2   k8snode2   <none>
```

**NB**: Récupérer l'adresse IP de votre pod et remplacez les '.' par des '-'

```bash
$ kubectl run curl --restart=Never --image=radial/busyboxplus:curl -n k8s-lab-1 -i --tty -- curl http://10-244-3-2.default.pod.cluster.local -I

HTTP/1.1 200 OK
Server: nginx/1.11.10
Date: Fri, 30 Nov 2018 17:21:54 GMT
Content-Type: text/html
Content-Length: 576
Last-Modified: Tue, 01 Jan 1980 00:00:00 GMT
Connection: keep-alive
ETag: "12cea600-240"
Accept-Ranges: bytes
```

Si vous voulez consulter les logs de votre application utiliser la commande suivante : 

```
$ kubectl logs <POD_NAME> -n k8s-lab-1
```

A ce stade, vous avez déployé votre première application. Bravo !