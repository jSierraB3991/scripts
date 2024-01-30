# main.tf

provider "google" {
  credentials = file(var.gcp_svc_key)
  project     = var.gcp_project
  region      = var.gcp_region
}


# Configuración de las instancias virtuales
resource "google_compute_instance" "vm_instance_1" {
  name         = "vm-1"
  machine_type = "n1-standard-1" # Cambia al tipo de máquina que desees
  zone         = var.gcp_region # Cambia a tu zona preferida

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10" # Cambia a tu imagen preferida
    }
  }

  network_interface {
    network = "default"
  }

  metadata_startup_script = <<-EOF
    #!/bin/bash
    apt-get update
    apt-get install -y docker.io
    curl -L "https://packages.gitlab.com/install/repositories/runner/gitlab-runner/script.deb.sh" | bash
    apt-get install -y gitlab-runner
    usermod -aG docker gitlab-runner
    systemctl restart docker
    systemctl enable docker
    systemctl restart gitlab-runner
    systemctl enable gitlab-runner
  EOF
}

resource "google_compute_instance" "vm_instance_2" {
  name         = "vm-2"
  machine_type = "n1-standard-1" # Cambia al tipo de máquina que desees
  zone         = var.gcp_region # Cambia a tu zona preferida

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10" # Cambia a tu imagen preferida
    }
  }

  network_interface {
    network = "default"
  }

  metadata_startup_script = <<-EOF
    #!/bin/bash
    apt-get update
    apt-get install -y docker.io
    curl -L "https://packages.gitlab.com/install/repositories/runner/gitlab-runner/script.deb.sh" | bash
    apt-get install -y gitlab-runner
    usermod -aG docker gitlab-runner
    systemctl restart docker
    systemctl enable docker
    systemctl restart gitlab-runner
    systemctl enable gitlab-runner
  EOF
}

# Configuración de la base de datos PostgreSQL
resource "google_sql_database_instance" "db_instance" {
  name             = "postgresql-instance"
  database_version = "POSTGRES_13"
  region           = var.gcp_region # Cambia a tu región preferida

  settings {
    tier = "db-custom-1-3840" # Cambia al tier que desees
  }
}

resource "google_sql_database" "db" {
  name     = "mi_basededatos"
  instance = google_sql_database_instance.db_instance.name
}

# Configuración de la red para permitir acceso a todas las instancias
resource "google_compute_firewall" "allow_all" {
  name    = "allow-all"
  network = "default"

  allow {
    protocol = "-1"
    ports    = []
  }

  source_ranges = ["0.0.0.0/0"]
}