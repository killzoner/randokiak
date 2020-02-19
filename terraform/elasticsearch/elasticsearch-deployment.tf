resource "kubernetes_deployment" "elasticsearch" {
  metadata {
    name = "elasticsearch"

    labels = {
      app = "elasticsearch"
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "elasticsearch"
      }
    }

    template {
      metadata {
        labels = {
          app = "elasticsearch"
        }
      }

      spec {
        container {
          name  = "elasticsearch"
          image = "docker.elastic.co/elasticsearch/elasticsearch:7.6.0"

          port {
            container_port = 9200
          }

          env {
            name  = "ES_JAVA_OPTS"
            value = "-Xms512m -Xmx512m"
          }

          env {
            name  = "bootstrap.memory_lock"
            value = "true"
          }

          env {
            name  = "cluster.name"
            value = "docker-cluster"
          }

          env {
            name  = "discovery.type"
            value = "single-node"
          }
        }
      }
    }
  }
}

