# (추천) 사용자 정의 규칙 사양(단위 테스트) 작성

사용자 정의 규칙 사양은 사용자 정의 규칙에 대한 단위 테스트입니다. 사용자 정의 규칙을 반복하면서 단위 테스트를 유지하는 것이 좋습니다. 이러한 단위 테스트에는 일반적으로 유효하고 유효하지 않은 테스트 케이스를 나타내는 Terraform 샘플을 포함합니다. 다음은 태그 요구 사항을 강제하는 사용자 정의 규칙에 대한 예제 사용자 정의 규칙 사양입니다:

{% code title="/spec/rules/REQUIRED_S3_BUCKET_TAGS" %}
```
provider "aws" {
  region = "us-east-1"
}

resource "aws_s3_bucket" "valid" {
  bucket = "valid-bucket"

  tags = {
    owner           = "devteam1"
    classification  = "public"
  }
}

resource "aws_s3_bucket" "invalid" {
  bucket = "invalid-bucket"

  tags = {
    owner           = "devteam5"
    classification  = "available"
  }
}
```
{% endcode %}

## 사용자 정의 규칙 사양 템플릿 만들기

다음 CLI 명령을 사용하세요:

```
snyk iac rules init
```

이 명령은 인터랙티브 프롬프트 세트를 소개하고 프로젝트, 규칙, 규칙 사양(테스트용), 그리고 리소스 관계를 설정하도록 허용합니다. `규칙 사양`을 선택하면, 규칙 사양 템플릿을 생성하기 위한 인터랙티브 프롬프트 세트가 제공됩니다. 이를 통해 테라폼, 클라우드포메이션, 애저 리소스 매니저 또는 쿠버네티스를 기반으로 한 템플릿이 생성됩니다.

**인터랙티브 프롬프트의 예시 출력:**

```
What do you want to initialize? rule spec
Choose a rule ID: RULE_1
Spec name: RULE_SPEC_1
Input type: tf
```

## 사용자 정의 규칙 사양 작성

사용자 정의 규칙 사양 템플릿이 생성되면 선택한 입력 유형(예: Terraform)과 유효한 구성과 무효한 구성을 채워 넣어야 합니다.​ ​​​