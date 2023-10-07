provider "aws" {
  region = "us-west-2"
}

module "eks" {
    source = "terraform-aws-modules/eks/aws"
    cluster_name = "products-dev"
    cluster_version = "1.20"
    subnets = [] # add VPC subnet IDs here

    node_groups = {
        eks_nodes = {
            desired_capacity = 2
            max_capacity = 3
            min_capacity = 1

            instance_type = "t3.medium"
            key_name = var.key_name
        }
    }
}