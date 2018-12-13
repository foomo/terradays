module "write" {
  source = "./modules/write"

  # input variables
  filename = "${path.root}/data.json"
}

module "read" {
  source = "./modules/read"

  # input variables
  filename = "${module.write.filename}"
}

#module "echo" {
#  source = "./modules/echo"
#
#  # input variables
#  data = "${module.read.data}"
#}
