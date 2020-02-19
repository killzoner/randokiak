resource "kubernetes_service" "elasticsearch" {
  metadata {
    name = "elasticsearch"

    labels = {
      app = "elasticsearch"
    }
  }

  spec {
    port {
      name        = "9200"
      port        = 9200
      target_port = "9200"
    }

    selector = {
      app = "elasticsearch"
    }

    type = "NodePort"
  }
}

