[![GoDoc](https://godoc.org/github.com/joshdk/tfbundle/bundle?status.svg)](https://godoc.org/github.com/joshdk/tfbundle/bundle)
[![Go Report Card](https://goreportcard.com/badge/github.com/joshdk/tfbundle)](https://goreportcard.com/report/github.com/joshdk/tfbundle)
[![CircleCI](https://circleci.com/gh/joshdk/tfbundle.svg?&style=shield)](https://circleci.com/gh/joshdk/tfbundle/tree/master)

# TFBundle

📦 Bundle a single artifact as a Terraform module

## Installing

You can fetch this library by running the following

    go get -u github.com/joshdk/tfbundle

## Usage

### As a command line tool

The `tfbundle` tool can be used to consume a given file (`lambda.zip` in this example), and generate a fully contained Terraform module (`module.tgz`).

```
$ ls
lambda.zip

$ tfbundle lambda.zip module.tgz

$ ls
lambda.zip module.tgz

$ tar -tf module.tgz
main.tf
artifact
artifact/lambda.zip
```

### As the resulting Terraform module

#### Configuration

The resulting module takes no inputs, just a `source`.

```hcl
module "artifact" {
  source = "module.tgz"
}
```

#### Outputs

The `filename` output contains the absolute path to the original bundled file, after being fetched with `terraform get`. Intended to be used as a passthrough when configuring [`aws.lambda_function.filename`](https://www.terraform.io/docs/providers/aws/r/lambda_function.html#filename).

```hcl
output "filename" {
  value = "${module.artifact.filename}"
}
```

The `source_code_hash` output contains a content hash for the original bundled file. Intended to be used as a passthrough when configuring [`aws.lambda_function.source_code_hash`](https://www.terraform.io/docs/providers/aws/r/lambda_function.html#source_code_hash).

```hcl
output "source_code_hash" {
  value = "${module.artifact.source_code_hash}"
}
```

## License

This library is distributed under the [MIT License](https://opensource.org/licenses/MIT), see LICENSE.txt for more information.
