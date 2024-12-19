# Terraform 파일

{% hint style="info" %}
현재 IaC의 CLI와 달리 IaC+의 CLI는 Terraform 변수와 모듈을 스캔할 수 있습니다. 이 변수와 모듈이 공개되었느지 비공개인지에 관계없이 `snyk iac test` 명령을 실행하기 전에 `terraform init`를 실행하고, CLI는 생성된 `.terraform` 파일을 읽습니다.
{% endhint %}

Snyk Infrastructure as Code를 사용하면 CLI를 사용하여 정적 구성 파일과 Terraform 계획 출력을 모두 스캔할 수 있습니다.

## Terraform 구성 파일

현재 IaC 및 IaC+를 사용하여 예를 들어 `main.tf`와 같은 구성 파일을 CLI를 사용하여 스캔할 수 있습니다.

현재 IaC에서는 외부 모듈을 현재 지원하지 않습니다. 변수 처리에 대한 자세한 내용은 [Terraform 변수 지원](../../../../scan-with-snyk/snyk-iac/scan-your-iac-source-code/scan-terraform-files/terraform-variables-support-current-iac.md)을 참조하십시오.

## 구성 파일 스캔

파일 이름 또는 전체 디렉토리를 지정할 수 있습니다.

```
snyk iac test main.tf
snyk iac test .
```

## Terraform 계획

Terraform 계획은 구성 파일을 작성하고 해당 변경 사항을 배포하기 전에 실행되는 단계입니다.

`$ terraform plan`은 대상 환경을 원하는 상태와 일치시키기 위해 필요한 변경 사항을 식별합니다.

이 계획 단계에서 대상 Terraform 배포에서 사용되는 모든 변수 및 Terraform 모듈이 고려됩니다.

사용자 정의 Terraform 모듈을 작성하고 해당 모듈을 배포 중에 참조하는 경우 Terraform 계획 출력에 포함되어 이에 따라 스캔됩니다. 즉, Terraform 계획 출력은 보안 측면에서 스캔할 수 있는 완전한 아티팩트를 제공합니다.

Snyk CLI 버전 1.594.0부터는 이 아티팩트를 {{Snyk IaC}} CLI를 사용하여 스캔할 수 있습니다.

이 파일은 처리하기 위해 Snyk로 보내지 않습니다. CLI를 사용하여 로컬에서 스캔되며 사용자의 장치를 벗어나지 않습니다.

## Terraform 계획 출력 스캔

유효한 JSON 파일로 저장해야 하는 Terraform 계획 출력 경로를 제공하십시오.

```
snyk iac test tf-plan.json
```

기본적으로 Snyk는 전체 인프라가 아닌 인프라에 적용될 변경 사항을 스캔합니다.

`--scan=` 옵션을 사용하여 이 동작을 변경할 수 있습니다.

* `--scan=resource-changes`는 기본 동작입니다. 이는 `$ terraform apply`를 실행했을 때 인프라에 적용될 변경 사항만을 스캔합니다.
* `--scan=planned-values`는 기존 인프라와 적용될 변경 사항의 결과를 제공하는 전체 계획 상태를 스캔합니다.

이미 JSON 파일로 저장된 Terraform 계획 출력이 없는 경우 다음 단계를 따르야 할 수 있습니다.

```
terraform plan -out=tfplan.binary
terraform show -json tfplan.binary > tf-plan.json
```

`tf-plan.json` 파일의 이름을 필요에 따라 지정할 수 있습니다.

이러한 파일은 민감하게 다루어져야하며 소스 제어에 커밋하는 것이 권장되지 않습니다.

## Terraform 스캔 문제 해결

{% hint style="info" %}
`--experimental` 옵션이 더 이상 Terraform 프로젝트를 테스트하는 데 필요하지 않습니다.
{% endhint %}

**정적 파일과 계획 출력을 스캔하는 데 차이점**이 있을 수 있습니다. 이는 다음과 관련이 있을 수 있습니다:

* **변수** - Terraform 계획 출력은 변수에 저장된 값들을 고려합니다.
* **Terraform 모듈** - Terraform 계획 출력은 사용 중인 Terraform 모듈에서 발견된 모든 구성 문제를 포함합니다.
* **차이** - 기본적으로 Terraform 계획 출력이 인프라에 적용될 변경 사항에 대해서만 구성 문제를 스캔하며 전체 배포가 아닙니다. 반면 정적 스캔은 모든 파일을 확인합니다. `--scan=planned-values` 옵션을 사용하여 다시 스캔을 실행해 보십시오.

이 정보에 기반해 설명할 수 없는 불일치 사항을 발견하면 Support에 [요청](https://support.snyk.io)을 제출하십시오.