# 드리프트를 무시하기 위한 리소스 제외

`.snyk` 정책 파일은 `snyk iac describe`에 의해 IaC(drift)로 간주되지 않도록 리소스를 제외하는 데 사용할 수 있습니다. 자세한 내용은 [`.snyk` 정책 파일 문서](../../../../manage-risk/policies/the-.snyk-file.md)를 참조하십시오.

리소스를 제외할 때는 단순히 일부 리소스를 제외하려면 `.snyk`를 사용하십시오. 더 복잡한 요구 사항이 있는 경우 필터 규칙을 사용하는 것을 고려하십시오. 자세한 내용은 [필터 규칙](filter-rules.md)을 참조하십시오.

`snyk iac describe` 명령을 시작하는 디렉토리에서 (일반적으로 IaC 리포의 루트인) `.snyk` 파일을 작성하십시오.

각 줄은 다음과 같은 구조여야 합니다:

* `리소스_유형.리소스_아이디`, 여기서 `리소스_아이디`는 해당 유형의 모든 리소스를 제외하기 위한 와일드카드입니다.
* `리소스_유형.리소스_아이디.필드_이름.경로`, 여기서 `리소스_아이디`는 특정 유형의 특정 필드에 대한 드리프트를 무시하기 위한 와일드카드이고 `경로`에도 와일드카드가 들어갈 수 있습니다.

## IaC 무시의 예시

IAM 사용자인 "_tfc-demo_" 하나를 무시합니다.

{% code title=".snyk" %}
```yaml
# Snyk (https://snyk.io) 정책 파일, 알려진 취약점을 수정하거나 무시합니다.
version: v1.22.1
exclude:
  iac-drift:
    - aws_iam_user.tfc-demo
```
{% endcode %}

모든 S3 버킷 드리프트를 무시합니다.

{% code title=".snyk" %}
```yaml
# Snyk (https://snyk.io) 정책 파일, 알려진 취약점을 수정하거나 무시합니다.
version: v1.22.1
exclude:
  iac-drift:
    - aws_s3_bucket.*
```
{% endcode %}

`.snyk` 정책 파일은 규칙의 부정을 지원합니다. 이를 통해 특정 유형을 제외한 모든 것을 무시할 수 있습니다. 다음 예시에서는 오직 S3 버킷만 무시되지 않습니다:

{% code title=".snyk" %}
```yaml
# Snyk (https://snyk.io) 정책 파일, 알려진 취약점을 수정하거나 무시합니다.
version: v1.22.1
exclude:
  iac-drift:
    - '*'
    - '!aws_s3_bucket'
```
{% endcode %}

특정 IAM 정책 첨부 (AWSServiceRoleForRDS)를 해당 ARN(arn:aws:iam::aws:policy/aws-service-role/AmazonRDSServiceRolePolicy)을 사용하여 무시합니다.

{% code title=".snyk" %}
```yaml
# Snyk (https://snyk.io) 정책 파일, 알려진 취약점을 수정하거나 무시합니다.
version: v1.22.1
exclude:
  iac-drift:
    - aws_iam_policy_attachment.AWSServiceRoleForRDS-arn:aws:iam::aws:policy/aws-service-role/AmazonRDSServiceRolePolicy
```
{% endcode %}

아래와 같이 my-bucket이라는 S3 버킷을 무시합니다.

```
# Snyk (https://snyk.io) 정책 파일, 알려진 취약점을 수정하거나 무시합니다.
version: v1.22.1
exclude:
  iac-drift:
      # my-bucket이라는 S3 버킷을 무시합니다
    - aws_s3_bucket.my-bucket
      # 모든 aws_instance 리소스를 무시합니다
    - aws_instance.*
      # 모든 람다 함수의 환경을 무시합니다
    - aws_lambda_function.*.environment
      # aws_iam_role.AmazonSSMRoleForInstances 및 aws_iam_role.AWSServiceRoleForAmazonSSM과 같이 Amazon이 포함된 리소스를 무시합니다
    - *role.*Amazon*
      # my-lambda-name 람다 함수의 lastModified를 무시합니다
    - aws_lambda_function.my-lambda-name.last_modified
```

## 필터 규칙을 우선하는 방법[​](https://docs.driftctl.com/0.22.0/usage/filtering/driftignore#precedence-over-filter-rules) <a href="#precedence-over-filter-rules" id="precedence-over-filter-rules"></a>

이 페이지에서 설명하는 리소스 무시 수단을 필터 규칙과 함께 사용할 수 있습니다.

**참고:** 필터 규칙에 의해 동일한 리소스가 포함되고 `.snyk` 파일 안에서 제외된 경우, `snyk iac describe`는 해당 리소스를 무시합니다.

## 드리프트 제외 규칙 자동 생성[​](https://docs.driftctl.com/0.22.0/usage/filtering/driftignore#automatically-generate-a-driftignore-file) <a href="#automatically-generate-a-driftignore-file" id="automatically-generate-a-driftignore-file"></a>

자세한 내용은 `snyk iac update-exclude-policy --help`를 실행하십시오.

이 명령은 모든 감지된 드리프트를 추가하여 `.snyk` 정책 파일을 생성해 모두 무시하는 데 도움을 줍니다.

예를 들어 모든 관리되지 않는 리소스를 한꺼번에 무시하려면 다음 명령을 실행하십시오:

```
$ snyk iac describe --json | snyk iac update-exclude-policy
```