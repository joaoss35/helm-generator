# Helm Generator

Helm can be a little complex sometimes, specially if you are new to the tool and need a quick start. As such, Helm Generator is a tool that simplifies the process of fetching a Helm chart for you.
Instead of having to add the helm repo yourself, or even fetching it for a possible fork, Helm-Generator abstracts that complexity for you.
The helm-generator binary allows you to save a specific helm-chart, of a specific version to a directory of your choice.

## Usage

```shell
helm-generator --chart <chart-name> --version <chart-version> --dest <destination-dir>
```

The following flags are available:

-  **`--chart`**: name of the Helm chart to fetch
-   **`--version`**: version of the Helm chart to fetch
-   **`--dest`**: destination directory to save the Helm chart
-   **`--help`**: display help message


## Requirements 
- [Helm](https://helm.sh) 3.5.0 or higher
  

## Usage Example

Fetch the latest version of the `mysql` Helm chart to the current directory:
```shell
helm-generator --chart mysql --dest .
```

Fetch a specific version of the `mysql` Helm chart to the charts directory:
```shell
helm-generator --chart mysql --version 3.9.0 --dest charts
```

## Supported charts

At the moment, only the [Bitnami](https://bitnami.com/) helm repo is supported, which means that you can fetch the following charts:

- `airflow`
- `apache`
- `appsmith`
- `argo-cd`
- `argo-workflows`
- `aspnet-core`
- `cassandra`
- `cert-manager`
- `clickhouse`
- `common`
- `concourse`
- `consul`
- `contour-operator`
- `discourse`
- `dokuwiki`
- `drupal`
- `ejbca`
- `elasticsearch`
- `etcd`
- `external-dns`
- `flink`
- `fluentd`
- `flux`
- `ghost`
- `gitea`
- `grafana`
- `haproxy`
- `harbor`
- `influxdb`
- `jaegar`
- `jasperreports`
- `jenkins`
- `joomla`
- `jupyterhub`
- `keycloak`
- `kiam`
- `kibana`
- `kong`
- `kube-prometheus`
- `kube-state-metrics`
- `kubeapps`
- `kubernetes-event-exporter`
- `logstash`
- `magento`
- `mariadb-galera`
- `mariadb`
- `mastodon`
- `matomo`
- `mediawiki`
- `memchached`
- `metallb`
- `metrics-server`
- `minio`
- `mongodb-sharded`
- `mongodb`
- `moodle`
- `multus-cni`
- `mxnet`
- `mysql`
- `nats`
- `nginx-ingress-controller`
- `nginx`
- `node-exporter`
- `oauth2-proxy`
- `odoo`
- `opencart`
- `osclass`
- `parse`
- `phpbb`
- `phpmyadmin`
- `pinniped`
- `postgresql-ha`
- `postgresql`
- `prestashop`
- `pytorch`
- `rabbitmq-cluster-operator`
- `rabbitmq`
- `redis-cluster`
- `redis`
- `redmine`
- `schema-registry`
- `sealed-secrets`
- `solr`
- `sonarqube`
- `spark`
- `spring-cloud-dataflow`
- `suitecrm`
- `supabase`
- `tensorflow-resnet`
- `thanos`
- `tomcat`
- `wavefront-hpa-adapter`
- `wavefront-prometheus-storage-adapter`
- `wavefront`
- `whereabouts`
- `wildfly`
- `wordpress`
- `zookeeper`
