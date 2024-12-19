# 필터 규칙

필터 규칙을 사용하여 리소스를 설명하고 무시할 수 있습니다. 포함 및 제외 논리를 모두 사용할 수 있습니다.

필터 규칙을 사용하면 워크플로에 리소스 집합을 포함하거나 제외하는 복잡한 포함 및 제외 표현식을 작성할 수 있습니다. 이 기능은 표현 언어 [JMESPath](https://jmespath.org)에 의해 구동됩니다.

필터는 다음 필드를 포함하는 정규화된 'struct'에 적용됩니다.

- Type: 리소스의 유형, 예를 들어 `aws_s3_bucket`
- Id: 리소스의 ID, 예를 들어 `my-bucket-name`

다음은 필터 규칙의 예입니다.

```
# 검색에는 S3 버킷만 포함됩니다.
$ snyk iac describe --filter="Type=='aws_s3_bucket'"
# 또는 (따옴표 사이의 쉘 특수 문자를 이스케이프해야 함에 주의하십시오)
$ snyk iac describe --filter=$'Type==\'aws_s3_bucket\''
# 'my-bucket-name'이라는 s3 버킷만 제외합니다.
$ snyk iac describe --filter=$'Type==\'aws_s3_bucket\' && Id!=\'my-bucket-name\''
# ID 접두사가 'terraform-'인 버킷을 무시합니다.
$ snyk iac describe --filter=$'!(Type==\'aws_s3_bucket\' && starts_with(Id, \'terraform-\'))'
# ID 접미사가 '-test'인 버킷을 무시합니다.
$ snyk iac describe --filter=$'!(Type==\'aws_s3_bucket\' && ends_with(Id, \'-test\'))'
```