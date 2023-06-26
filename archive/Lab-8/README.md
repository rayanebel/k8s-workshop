# Lab 03 - Monitoring

L'objectif de ce lab est de mettre en place un mécanisme vous permettant de monitorer votre cluster Kubernetes. Pour celà, nous allons nous interresser à une solution en paritculier : Prometheus, une solution de monitoring open source capable de collecter, d'aggreger et de visualiser des métriques et d'alerter en cas de dysfonctionnement. Membre à part entière de la CNCF, Prometheus offre une interroperabilité parfaite avec Kubernetes et est donc le candidat parfait pour assurer cette fonction.

Dans ce lab, nous allons voir pas à pas comment mettre en place cette solution au sein de notre cluster et, pour nous faciliter les choses, nous allons utiliser un outil très particulier qui s'incrit parfaitement dans la démarche DevOps : **Prometheus-operator**, un outil open source développé par la communauté CoreOS qui va nous permettre d'intégrer Prometheus nativement à notre environnement Kubernetes et de gérer automatiquement son déploiement et sa configuration.

**Objets Kubernetes** : Deployment, Daemonsets, Secrets, Configmaps, Namespaces, Roles, ClusterRole, RoleBinding, ClusterRoleBinding, ServiceAccount...

## Prometheus Operator

```
cd prometheus-operator
kubectl apply -f namespace.yml
kubectl get ns
```
```
kubectl apply -f prometheus-operator-rbac.yml
kubectl apply -f prometheus-operator.yml

kubectl get pod -n monitoring

NAME                                  READY   STATUS    RESTARTS   AGE
prometheus-operator-bdf79ff67-sknp7   1/1     Running   0          1d

kubectl get deployment prometheus-operator -n monitoring

NAME                  DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
prometheus-operator   1         1         1            1           2d

kubectl get customresourcedefinitions | grep -I "monitoring.coreos.com"

alertmanagers.monitoring.coreos.com     2018-11-23T13:01:53Z
prometheuses.monitoring.coreos.com      2018-11-23T13:01:53Z
prometheusrules.monitoring.coreos.com   2018-11-23T13:01:53Z
servicemonitors.monitoring.coreos.com   2018-11-23T13:01:53Z
```

## Prometheus

```
kubectl apply -f prometheus-rbac.yml
kubectl apply -f prometheus.yml

kubectl get prometheus -n monitoring

NAME   CREATED AT
k8s    1d
```

## Exporter(s)
## Grafana