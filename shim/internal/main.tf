output "filename" {
  value = "${path.module}/artifact/{{ .Artifact }}"
}

output "source_code_hash" {
  value = "${base64sha256(file("${path.module}/artifact/{{ .Artifact }}"))}"
}
