# LAB 2 - Accéder à vos applications via les Services

Dans ce lab, nous allons voir comment accéder à notre précédente application en utilisant les services Kubernetes.

**Rappel :**

Un service, est une structure qui permet de se connecter de façon fiable aux conteneurs qui se trouvent dans les pods, en utilisant une adresse IP virtuelle stable et un port spécifique _(ex: 8080)_. La correspondance réseau entre les services et les pods est assuré par le composant : **kube-proxy**

Il existe trois types de service : 

* ClusterIP : Accéder à votre application uniquement depuis votre cluster (pas d'accès externe).
* NodePort : Accéder à un service depuis l'extérieur via un port ouvert sur vos noeuds.
* Loadbalancer : Accéder à un service via un LoadBalancer externe. Ce type de LoadBalancer est seulement disponible si vos machines se trouvent chez un cloud provider tels que AWS, Azure ou GCP.

## Déploiement

Nous allons commencer par déployer un service de type ClusterIP. Pour ce faire, executez la commande suivante : 

```bash
cd Lab-2/
$ kubectl apply -f cookie-app-service-clusterip.yml -n k8s-lab-1
```

Vérifiez la création de votre service en exécutant la commande suivante : 

```bash
kubectl get svc -n k8s-lab-1

NAME                   TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE     SELECTOR
cookie-app-clusterip   ClusterIP   10.96.210.27   <none>        80/TCP    46s     app=cookie-app
```

Ici, nous avons notre service qui à pour adresse IP : 10.96.210.27 et qui redirige le trafic vers les pods qui ont le label **app=cookie**

Si vous voulez connaitre les labels associés à vos pods, vous pouvez exécuter les commandes suivantes : 

```bash
$ kubectl get pod -n k8s-lab-1 --show-labels
NAME                       READY   STATUS             RESTARTS   AGE   LABELS
cookie-app                 1/1     Running            0          70m   app=cookie-app

$ kubectl describe pod cookie-app -n k8s-lab-1
Name:               cookie-app
Namespace:          default
Priority:           0
PriorityClassName:  <none>
Node:               k8snode2/10.0.0.5
Start Time:         Fri, 30 Nov 2018 16:38:17 +0000
Labels:             app=cookie-app
Annotations:        kubectl.kubernetes.io/last-applied-configuration:
                      {"apiVersion":"v1","kind":"Pod","metadata":{"annotations":{},"labels":{"app":"cookie-app"},"name":"cookie-app","namespace":"default"},"spe...
Status:             Running
IP:                 10.244.3.2

[...]
```

Vous avez la possibilité de retrouver la liste des pods associés à un service en utilisant la commande suivante : 

```bash
$ kubectl describe svc cookie-app-clusterip -n k8s-lab-1 | grep -I "Endpoints"

Endpoints:         10.244.3.2:80
```

Ou alors, étant donné que lorsqu'un service kubernetes est déclare un objet de type `Endpoint` est aussi créé, vous avez également la possibilité de récupérer la liste des pods via cette commande : 

```bash
kubectl get endpoints cookie-app-clusterip -n k8s-lab-1 -o=jsonpath='{.subsets[*].addresses[*].ip}{"\n"}'

10.244.3.2
```

Créeons  maintenant un service de type NodePort. Pour ce faire exécuter la commande suivante :

```bash
$ kubectl apply -f cookie-app-service-nodeport.yml -n k8s-lab-1
```

Vérifiez ensuite la bonne création de votre service :

```bash
$ kubectl get svc cookie-app-nodeport -n k8s-lab-1

NAME                   TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
cookie-app-nodeport    NodePort    10.110.84.19   <none>        80:30001/TCP   4s
```

Ici, nous avons un service de type `NodePort` qui nous permet d'accéder à notre application via le port `30001` qui est ouvert sur tous les noeuds de notre cluster.

Pour vérifier, connectez vous en SSH sur l'un de vos noeuds et exécuter la commande suivante : 

```
$ netstat -nlp | grep 30001

tcp6  0  0 :::30001  :::*  LISTEN  18274/kube-proxy
```

Si vous souhaitez avoir des détails sur les endpoints _(pods)_ ciblés, vous pouvez reutiliser les commandes que l'on a vu plus haut.

Maintenant, il nous reste plus qu'a créer un service de type LoadBalancer. Pour ce faire, exécuter la commande suivante : 

```bash
$ kubectl apply -f cookie-app-service-lb.yml -n k8s-lab-1
```

Vérifiez ensuite la bonne création de votre service :

**NB**: Il faut attendre quelques minutes avant que le LoadBalancer côté Cloud Provider soit provisionné.

```bash
$ kubectl get svc cookie-app-nodeport -n k8s-lab-1

NAME                   TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
cookie-app-lb          LoadBalancer   10.111.175.155   137.xxx.xxx.xx     80:30008/TCP   84s
```

Ici, nous avons bien un service de type `LoadBalancer` qui à pour adresse IP publique 137.xxx.xxx.xx et qui réparti le traffic entre les différents noeuds du cluster sur le port 30008.

Tester maintenant l'accès à votre application en ouvrant votre navigateur préféré

![Application](./images/cookie-app.png)