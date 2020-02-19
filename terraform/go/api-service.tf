resource "kubernetes_service" "rdkapi" {
  metadata {
    name = "rdkapi"

    labels = {
      app = "rdkapi"
    }
  }

  spec {
    selector = {
      app = "rdkapi"
    }
    port {
      name        = "8082"
      port        = 8082
      target_port = "8082"
    }

    type = "NodePort"
  }
}

