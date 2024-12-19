# Terraform 변수 지원 (현재 IaC)

{% hint style="info" %}
이 페이지는 현재 IaC에만 해당됩니다.
{% endhint %}

Terraform (TF) 변수 지원은 현재 CLI에서만 사용 가능합니다. Snyk은 현재 다음을 지원합니다:

- [입력 변수](https://www.terraform.io/language/values/variables)
- [로컬 값](https://www.terraform.io/language/values/locals)

Snyk은 현재 [출력 값](https://www.terraform.io/language/values/outputs)을 지원하지 않습니다.

CLI는 모든 디렉토리를 스캔하고 지원되는 TF 파일을 포함하는 각 디렉토리를 자체 모듈로 처리합니다. 변수를 포함하는 각 모듈은 적절하게 참조됩니다.

지원되는 TF 파일 형식은 `.tf`, `.tfvars`, `.auto.tfvars`입니다. Snyk은 현재 환경 변수나 `--var` CLI 옵션을 사용하여 설정 및 정의된 변수를 지원하지 않습니다.

스캔은 [변수 정의 우선순위](https://www.terraform.io/language/values/variables#variable-definition-precedence)를 TF가 다루는 방식과 동일하게 다룹니다.

외부 변수 정의 파일을 `--var-file` 옵션을 사용하여 로드할 수 있습니다. 예를 들어:

`snyk iac test myproject/staging/networking --var-file=myproject/vars.tf`

이는 `myproject` 디렉토리에서 `vars.tf` 정의 파일을 로드하고, 해당 변수가 존재하는 경우 참조하고, 이를 예제에서 스캔된 경로인 `myproject/staging/networking`의 컨텍스트에 적용합니다.

더 많은 정보를 보려면 `IAC test`의 [도움말](../../../../snyk-cli/commands/iac-test.md)을 참고하세요.

## 지원되는 Terraform 표현식

현재 지원되는 표현식은 다음과 같습니다:

- [산술 및 논리 연산자](https://www.terraform.io/language/expressions/operators)
- [문자열 및 템플릿](https://www.terraform.io/language/expressions/strings#strings-and-templates)
- [조건 표현식](https://www.terraform.io/language/expressions/conditionals)
- [For 표현식](https://www.terraform.io/language/expressions/for)
- [Splat 표현식](https://www.terraform.io/language/expressions/splat)

## 지원되는 Terraform 함수

현재 지원되는 함수는 다음과 같습니다:

- 숫자 함수 - 모든 함수
- 문자열 함수 - `lower`, `regex`, `regexall`, `replace`, `substr`, `title`, `upper`를 제외한 모든 함수
- 컬렉션 함수 - `chunklist`, `concat`, `distinct`, `flatten`, `length`, `merge`, `reverse`, `sort`
- 인코딩 함수 - `csvdecode`, `jsondecode`, `jsonencode`
- 날짜 및 시간 함수 - `formatdate`, `timeadd`

## Terraform 변수 예제

### **올바른 우선순위에서 변수 처리**

다음 예제에서 새 리소스를 구성하고, `remote_user_addr`라는 변수를 사용하여 `cidr_blocks` 값을 설정하고 있습니다.

이 변수는 `variables.tf` 파일 내에서 기본값으로 정의되어 있지만, `terraform.tfvars` 파일 내에서 값이 재정의되고 있습니다.

최종적으로 값은 `0.0.0.0/0`로 설정되며, 이로 인해 CLI에서 이슈가 발생합니다.

```hcl
vpc.tf

resource "aws_security_group_rule" "ssh" {
  type              = "ingress"
  from_port         = 22
  to_port           = 22
  protocol          = "tcp"
  cidr_blocks       = [var.remote_user_addr]
  security_group_id = aws_security_group.allow.id
}
```

```hcl
variables.tf

variable "remote_user_addr" {
  type = string
  default = "11.0.0.0/24"
}
```

```hcl
terraform.tfvars

remote_user_addr = "0.0.0.0/0"
```

## **변수를 사용하는 조건 표현식**

다음 예제에서는 조건 표현식을 사용하여 로컬 및 입력 변수를 함께 사용하고 있습니다.

우리는 `local.test`가 0과 같은지 확인하고, `cidr_blocks`를 해당 값에 따라 설정하고 있습니다.

우리의 경우 `local.test`가 0과 같으며, 값은 `var.remote_user_addr`의 값으로 설정되어 CLI에서 이슈가 발생합니다.

```hcl
vpc.tf

resource "aws_security_group_rule" "ssh" {
  type              = "ingress"
  from_port         = 22
  to_port           = 22
  protocol          = "tcp"
  cidr_blocks       = local.test == 0 ? [var.remote_user_addr] : ["11.0.0.0/24"]
  security_group_id = aws_security_group.allow.id
}

locals {
  test = 0
}
```

```hcl
variables.tf

variable "remote_user_addr" {
  type = string
  default = "0.0.0.0/0"
}
```