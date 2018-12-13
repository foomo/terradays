data "external" "data" {
  program = ["bash", "${path.module}/scripts/data.sh"]

  query = {
    filename = "${var.filename}"
  }
}