resource "kubernetes_deployment" "rdkgen" {
  metadata {
    name = "rdkgen"

    labels = {
      app = "rdkgen"
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "rdkgen"
      }
    }

    template {
      metadata {
        labels = {
          app = "rdkgen"
        }
      }

      spec {
        container {
          name  = "rdkgen"
          image = "randokiak/randokiak-gen:0.0.1"
          #imagePullPolicy = "Never" # for use with minikube

          env {
            name  = "PORT"
            value = "50051"
          }

        }
      }
    }
  }
}

