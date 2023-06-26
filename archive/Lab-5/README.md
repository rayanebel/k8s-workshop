# LAB 5 - Monitorer la santé de vos conteneurs

Afin de savoir si nos applications conteneurisées sont saines et prêtes à servir le trafic, Kubernetes nous fournit des mecanismes de controle que l'on appel les **Health Checks**.

Ces mécanismes de controle sont exécuter par Kubelet pour déterminer quand redémarrer un conteneur _(LivenessProbe)_ et par les services pour déterminer si un pod doit recevoir du trafic ou non _(readinessProbe)_.

Dans ce lab, nous allons donc voir comment mettre en place ces mécanismes de controle au sein de notre application.

# Déploiement

Commençons par recréer notre application précedente en y ajoutant les mécanismes de contôle suivant : `LivenessProbe` et `ReadinessProbe`.

Pour ce faire, il vous suffit d'exécuter les commandes suivantes : 

```bash
$ cd Lab-5/
$ kubectl create ns k8s-lab-5
$ kubectl apply -f cookie-app-deployment.yml -n k8s-lab-5
```

Vérifions que notre application à bien été déployé

```bash
$ kubectl get pods -n k8s-lab-5

NAME                                      READY   STATUS    RESTARTS   AGE
cookie-app-healthcheck-7cdfd84cf9-zfgql   1/1     Running   0          21s
```

Regardons ensuite les détails de notre deployment afin de comprendre les mécanismes de controle que nous venons de mettre en place.

```bash
$ kubectl describe deployment cookie-app-healthcheck -n k8s-lab-5

Name:                   cookie-app-healthcheck
Namespace:              default
CreationTimestamp:      Sat, 01 Dec 2018 09:55:15 +0000
Labels:                 app=cookie

[...]

Pod Template:
  Labels:  app=cookie
  Containers:
   webapp:
    Image:        quay.io/coreos/example-app:v1.0
    Port:         80/TCP
    Host Port:    0/TCP
    Liveness:     http-get http://:80/ delay=2s timeout=1s period=5s #success=1 #failure=3
    Readiness:    http-get http://:80/ delay=10s timeout=1s period=5s #success=1 #failure=3
    Environment:  <none>
    Mounts:       <none>
  Volumes:        <none>

[...]
```

Sur le résultat ci-dessus, nous avons configurer un `LivenessProbe` qui fait un check HTTP sur le port 80 toutes les 5 secondes en laissant un délais de 2 secondes avant de commencer le premier check (démarrage du conteneur). Nous avons ensuite, un `ReadinessProbe` qui réalise le même test que le LivenessProbe sauf qu'il laisse un délais de 10 secondes avant de commencer le premier test.

On voit également que les tests ont reussis et que notre application est saine.

```bash
$ kubectl get pods -n k8s-lab-5

NAME                                      READY   STATUS    RESTARTS   AGE
cookie-app-healthcheck-7cdfd84cf9-zfgql   1/1     Running   0          21s
```

Faisons en sorte maintenant de mettre en echec le LivenessProbe que nous avons mis en place. Pour ce faire, editez votre deployment et changer le port dans la section LivenessProbe.

```bash
$ kubectl edit deployment cookie-app-healthcheck -n k8s-lab-5

apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  annotations:
    deployment.kubernetes.io/revision: "1"

[...]
spec:
[...]
  template:
    [...]
    spec:
       [...]
       livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 2
          periodSeconds: 5
          successThreshold: 1
          timeoutSeconds: 1
```

Enregistrez vos modifications et vérifiez ce qu'il se passe en affichant la liste de vos pods

```bash
kubectl get pod -n k8s-lab-5

NAME                                      READY   STATUS    RESTARTS   AGE
cookie-app-healthcheck-56c5f4bfb7-vwr5k   0/1     Running   3          51s
```

On voit ci-dessus que Kubelet tente de redémmarer indéfiniment notre conteneur (restarts).

On voit également que nous n'avons pas de pods disponible en affichant les détails de notre déploiement.

```bash
$ kubectl describe deployment cookie-app-healthcheck -n k8s-lab-5

Name:                   cookie-app-healthcheck
Namespace:              default
CreationTimestamp:      Sat, 01 Dec 2018 09:55:15 +0000
Labels:                 app=cookie
Selector:               app=cookie
Replicas:               1 desired | 1 updated | 1 total | 0 available | 1 unavailable

[...]
```

Faisons maintenant le même test pour le ReadinessProbe. Mais tout d'abord reinitialisez le lab en exécutant les commandes suivantes : 

```bash
$ kubectl delete -f cookie-app-deployment.yml -n k8s-lab-5

$ kubectl apply -f cookie-app-deployment.yml -n k8s-lab-5
```

Exposons ensuite notre application en utilisant un service de type ClusterIP

```bash
$ kubectl apply -f cookie-app-service.yml -n k8s-lab-5
```

Affichons maintenant la liste des endpoints (pods) vers lesquels notre service redirigera le trafic

```bash
$ kubectl describe svc cookie-app-clusterip -n k8s-lab-5 | grep -I "Endpoints"

Endpoints:         10.244.3.2:80
```

Modifions maintenant la configuration de notre readinessProbe afin de mettre en echec les tests.

```bash
$ kubectl edit deployment cookie-app-healthcheck -n k8s-lab-5

apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  annotations:
    deployment.kubernetes.io/revision: "1"

[...]
spec:
[...]
  template:
    [...]
    spec:
       [...]
       readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 10
          periodSeconds: 5
          successThreshold: 1
          timeoutSeconds: 1
```

Vérifions ensuite ce qu'il se passe

```bash
kubectl get pods -n k8s-lab-5

NAME                                      READY   STATUS    RESTARTS   AGE
cookie-app-healthcheck-59f945db47-bfjp2   0/1     Running   0          28s
```

Ici, nous voyons que notre application n'est pas prête. Si on affiche les détails de notre pod on peut également voir que le readinessProbe a echoué.

```bash
kubectl describe pod cookie-app-healthcheck-59f945db47-bfjp2 -n k8s-lab-5

[...]

Events:
  Type     Reason     Age               From               Message
  ----     ------     ----              ----               -------
  Normal   Scheduled  24s               default-scheduler  Successfully assigned default/cookie-app-healthcheck-59f945db47-bfjp2 to k8snode2
  Normal   Pulled     21s               kubelet, k8snode2  Container image "quay.io/coreos/example-app:v1.0" already present on machine
  Normal   Created    21s               kubelet, k8snode2  Created container
  Normal   Started    21s               kubelet, k8snode2  Started container
  Warning  Unhealthy  5s (x2 over 10s)  kubelet, k8snode2  Readiness probe failed: Get http://10.244.3.19:8080/: dial tcp 10.244.3.19:8080: connect: connection refused
```

Maintenant du côté de notre service si on affiche à nouveau la liste des endpoints nous pouvons voir que nous n'avons plus de pods en état de réceptionner le trafic

```bash
$ kubectl describe svc cookie-app-clusterip -n k8s-lab-5 | grep -I "Endpoints"

Endpoints:
```