resource "null_resource" "default" {
  provisioner "local-exec" {
    command = "echo '{\"0\":\"foo\",\"1\":\"bar\"}' > ${var.filename}"

    interpreter = ["bash", "-c"]
  }
}
