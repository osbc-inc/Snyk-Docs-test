---
description: Terraform AWS 프로바이더의 지원 및 제한 사항.
---

# Terraform AWS 공급자 지원

{% hint style="success" %}
**AWS Terraform 프로바이더**의 **4.0.0** 버전에서는 S3 서비스가 정의되는 방식에 변경 사항이 도입되었습니다. v4.0에서 S3 서비스의 정의는 이제 여러 Terraform 내의 리소스 블록에 분산되어 있습니다. 여러 파일에 걸쳐 S3 버킷의 인스턴스를 정의했다면, 이 업데이트는 파괴적인 변경이며 Snyk IaC에서 보안 결과에 부정적인 영향을 줄 수 있습니다.
{% endhint %}

{% hint style="info" %}
테라폼 문서에서: _독립적인 리소스를 통해 S3 버킷 설정의 관리를 분배하는 데 도움을 주기 위해_ `aws_s3_bucket` _리소스의 다양한 인수 및 속성은 읽기 전용이 되었습니다. 이러한 인수에 의존하는 구성은 해당_ `aws_s3_bucket_*` _리소스를 사용하도록 업데이트해야 합니다._
{% endhint %}

Terraform v4.0.0으로 이주하려면 S3 서비스 정의를 재구성하고 다시 가져와야 합니다. 이 작업 방식에 따라 보안 결과의 폭이 제한될 수 있습니다.

자세한 내용은 [HashiCorp의 Terraform V4 업그레이드 가이드](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/guides/version-4-upgrade)를 참조하십시오.

## Terraform AWS 프로바이더 지원 예제

다음은 Terraform **v3.x.x**에서 ACL을 사용하여 S3 버킷을 정의하는 예시입니다. `s3.tf`라는 파일에 있습니다.

{% code title="s3.tf" %}
```hcl
resource "aws_s3_bucket" "example" {
  # ... 기타 구성 ...
  acl = "public-read-write"
}
```
{% endcode %}

S3 버킷의 정의는 한 리소스 블록에 있습니다. 이 파일을 `snyk iac test s3.tf`를 사용하여 스캔하면 너그러운 ACL 설정으로 인한 보안 문제를 확인할 수 있습니다.

프로바이더의 **v4.0.0**으로, 특정 S3 버킷 속성이 이제 자체 리소스에 정의됩니다. 이전 예제를 이어가면, ACL 속성이 자체 리소스로 이동했기 때문에 재구성된 Terraform은 다음과 같습니다.

{% code title="s3.tf" %}
```hcl
resource "aws_s3_bucket" "example" {
  # ... 기타 구성 ...
}

resource "aws_s3_bucket_acl" "example" {
  bucket = aws_s3_bucket.example.id
  acl    = "public-read-write"
}
```
{% endcode %}

이 파일을 `snyk iac test s3.tf`를 사용하여 스캔하면 너그러운 ACL 설정에 대한 이전과 동일한 결과가 나옵니다.

## Terraform AWS 프로바이더의 지원 제한 사항

코드를 다르게 구조화할 수 있습니다. 예를 들어 S3 버킷 정의와 해당 속성을 별도의 Terraform 파일에 두는 방식입니다.

{% code title="s3_bucket.tf" %}
```hcl
resource "aws_s3_bucket" "example" {
  # ... 기타 구성 ...
}
```
{% endcode %}

{% code title="s3_acl.tf" %}
```hcl
resource "aws_s3_bucket_acl" "example" {
  bucket = aws_s3_bucket.example.id
  acl    = "public-read-write"
}
```
{% endcode %}

이 파일들을 스캔하면 보안 문제가 발생하지 않거나 너그러운 ACL에 대한 거짓 양성 결과를 받을 수 있습니다. 이는 현재 Snyk이 두 리소스를 연결할 수 없기 때문입니다.

Snyk은 제품에 이 추가 사용 사례를 지원하도록 작업 중입니다.

일시적 해결책은 단일 Terraform 파일에 리소스를 정의하는 것입니다.
