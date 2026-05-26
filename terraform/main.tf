terraform {
  required_version = ">= 1.5"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

variable "cluster_name" {
  default = "identity-security-fabric"
}

resource "aws_kms_key" "fabric" {
  description             = "Identity Security Fabric encryption key"
  deletion_window_in_days = 30
  enable_key_rotation     = true
  tags = {
    Name = "identity-fabric-key"
  }
}

resource "aws_s3_bucket" "worm_audit_log" {
  bucket = "identity-fabric-worm-audit-${data.aws_caller_identity.current.account_id}"
  tags = {
    Name = "identity-fabric-worm-audit"
  }
}

data "aws_caller_identity" "current" {}

output "kms_key_id" {
  value = aws_kms_key.fabric.id
}

output "s3_bucket" {
  value = aws_s3_bucket.worm_audit_log.id
}
