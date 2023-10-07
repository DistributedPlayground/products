provider "aws" {
    region = "us-west-2"
}

resource "aws_vpc" "dev_product_vpc" {
    cidr_block = "10.0.0.0/16"
    enable_dns_support = true
    enable_dns_hostnames = true

    tags = {
        Name = "dev_product_vpc"
    }
}

resource "aws_subnet" "dev_product_subnet" {
    vpc_id = aws_vpc.dev_product_vpc.id
    cidr_block = "10.0.1.0/24"
    availability_zone = "us-west-2a"

    tags = {
        Name = "dev_product_subnet"
    }
}

resource "aws_subnet" "dev_product_subnet2" {
    vpc_id = aws_vpc.dev_product_vpc.id
    cidr_block = "10.0.2.0/24"
    availability_zone = "us-west-2b"

    tags = {
        Name = "dev_product_subnet2"
    }
}