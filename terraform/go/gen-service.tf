resource "kubernetes_service" "rdkgen" {
  metadata {
    name = "rdkgen"

    labels = {
      app = "rdkgen"
    }
  }

  spec {
    selector = {
      app = "rdkgen"
    }
    port {
      name        = "50051"
      port        = 50051
      target_port = "50051"
    }

    type = "NodePort"
  }
}

