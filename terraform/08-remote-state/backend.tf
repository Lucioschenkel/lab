terraform {
  backend "s3" {
    bucket         = "ztm-terraform-mastery"
    key            = "s3_backend.tfstate"
    dynamodb_table = "s3-state-lock"
    region         = "us-east-1"
  }
}
