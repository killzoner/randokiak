resource "kubernetes_service" "webapp" {
  metadata {
    name = "webapp"

    labels = {
      app = "webapp"
    }
  }

  spec {
    selector = {
      app = "webapp"
    }
    port {
      name        = "4200"
      port        = 4200
      target_port = "4200"
      node_port   = 30001 # fixed nodePort example
    }

    type = "NodePort"
  }
}

