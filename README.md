[![CircleCI][circleci-badge]][circleci-link]
[![Go Report Card][go-report-card-badge]][go-report-card-link]
[![License][license-badge]][license-link]
[![Godoc][godoc-badge]][godoc-link]
[![CodeCov][codecov-badge]][codecov-link]
[![Releases][github-release-badge]][github-release-link]

# TFBundle

📦 Bundle a single artifact as a Terraform module

## Installing

### From source

You can install a development version of this tool by running:

```bash
$ go get -u github.com/joshdk/tfbundle
```

### Precompiled binary

Alternatively, you can download a precompiled [release][github-release-link] binary by running:

```bash
$ wget -q https://github.com/joshdk/tfbundle/releases/download/0.1.0/tfbundle_linux_amd64
$ sudo install tfbundle_linux_amd64 /usr/bin/tfbundle
```

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

The `etag` output contains an entity tag for the bundled file. Used as a pass-through when configuring [`aws.s3_bucket_object.etag`](https://www.terraform.io/docs/providers/aws/r/s3_bucket_object.html#etag).

```hcl
output "etag" {
  value = "${module.artifact.etag}"
}
```

The `filename` output contains the absolute path to the bundled file.

```hcl
output "filename" {
  value = "${module.artifact.filename}"
}
```

The `size` output contains the size in bytes of the bundled file.

```hcl
output "size" {
  value = "${module.artifact.size}"
}
```

The `source_code_hash` output contains a content hash for the bundled file. Used as a pass-through when configuring [`aws.lambda_function.source_code_hash`](https://www.terraform.io/docs/providers/aws/r/lambda_function.html#source_code_hash).

```hcl
output "source_code_hash" {
  value = "${module.artifact.source_code_hash}"
}
```

## License

This library is distributed under the [MIT License][license-link], see [LICENSE.txt][license-file] for more information.

[circleci-badge]:        https://circleci.com/gh/joshdk/tfbundle.svg?&style=shield
[circleci-link]:         https://circleci.com/gh/joshdk/workflows/tfbundle/tree/master
[go-report-card-badge]:  https://goreportcard.com/badge/github.com/joshdk/tfbundle
[go-report-card-link]:   https://goreportcard.com/report/github.com/joshdk/tfbundle
[license-badge]:         https://img.shields.io/badge/license-MIT-green.svg
[license-file]:          https://github.com/joshdk/tfbundle/blob/master/LICENSE.txt
[license-link]:          https://opensource.org/licenses/MIT
[godoc-badge]:           https://godoc.org/github.com/joshdk/tfbundle/bundle?status.svg
[godoc-link]:            https://godoc.org/github.com/joshdk/tfbundle/bundle
[codecov-badge]:         https://codecov.io/gh/joshdk/tfbundle/branch/master/graph/badge.svg
[codecov-link]:          https://codecov.io/gh/joshdk/tfbundle
[github-release-badge]:  https://img.shields.io/github/release/joshdk/tfbundle.svg
[github-release-link]:   https://github.com/joshdk/tfbundle/releases
