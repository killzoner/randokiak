resource "kubernetes_service" "pulsar" {
  metadata {
    name = "pulsar"
  }

  spec {
    selector = {
      app = "pulsar"
    }
    port {
      name        = "admin-tcp"
      port        = 8080
      target_port = 8080
      protocol    = "TCP"
    }
    # port {
    #   name        = "admin-udp"
    #   port        = 8080
    #   target_port = 8080
    #   protocol    = "UDP"
    # }

    port {
      name        = "api-tcp"
      port        = 6650
      target_port = 6650
      protocol    = "TCP"
    }
    # port { 
    #   name        = "api-udp"
    #   port        = 6650
    #   target_port = 6650
    #   protocol    = "UDP"
    # }
    
    type = "NodePort"
  }
}

