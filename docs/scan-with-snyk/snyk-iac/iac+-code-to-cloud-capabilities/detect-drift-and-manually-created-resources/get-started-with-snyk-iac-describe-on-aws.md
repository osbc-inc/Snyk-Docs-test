# Snyk IaC를 AWS에서 설명하고 시작하기

## **단계 1: 환경에 대한 AWS 인증 구성**

[`snyk iac describe`](../../../../snyk-cli/commands/iac-describe.md) 명령은 올바르게 실행되려면 클라우드 제공업체로의 인증이 필요합니다. 이를 위해 가능한 한 최소한의 읽기 전용 액세스 권한만 필요합니다. IAM 사용자를 위한 기본적으로 내장된 AWS `ReadOnlyAccess` IAM 정책을 시작 지점으로 사용할 수 있습니다.

`snyk iac describe`는 AWS의 표준 인증 방법인 `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, 및 `AWS_REGION` [환경 변수](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html#envvars-list)와 같은 것들을 재사용할 수 있습니다. 이러한 값이 설정되어 있으면 Snyk CLI가 자동으로 이를 활용하여 AWS에서 인증을 수행합니다.

또한 `~/.aws/credentials`에 [AWS 프로필](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html)을 구성하고 `AWS_PROFILE` 환경 변수를 사용할 수 있습니다.

## **단계 2: `describe` 명령을 사용하여 드리프트 보고**

### **미관리 리소스**

Snyk IaC는 미관리 리소스의 드리프트를 보고할 수 있습니다. 미관리 리소스란 클라우드 제공업체에는 있지만 Terraform 상태에는 없는 리소스를 말합니다. 이러한 리소스들을 Terraform으로 가져오거나 IaaS 계정에서 삭제할 수 있습니다.

관리 중인 리소스의 드리프트를 감지하는 자세한 정보는 [Command: plan](https://developer.hashicorp.com/terraform/cli/commands/plan)에서 Terraform CLI 문서를 참조하세요.&#x20;

### Terraform 상태 파일 선택

클라우드 환경에서 발생하는 드리프트를 이해하기 위해 환경의 상태를 하나 이상의 Terraform 상태 파일(`.tfstate`)과 비교합니다.

상태 파일은 로컬에 있을 수도 있고 S3 버킷에 있을 수도 있습니다. Terraform Cloud도 있지만 이는 이 문서의 범위를 벗어납니다.

`--from` 옵션은 `.tfstate` 파일의 경로를 결정하는 데 도움을 줍니다.

단일 로컬 Terraform 상태 파일의 경우 다음 명령을 사용합니다:

`$ snyk iac describe --from="tfstate://path/to/terraform.tfstate"`

특정 디렉토리에서 자동으로 발견된 모든 Terraform 상태를 로드하려면 다음과 같이 glob 패턴을 사용할 수 있습니다:

`$ snyk iac describe --from="tfstate://path/to/**/*.tfstate"`

S3 백엔드에 저장된 단일 Terraform 상태의 경우:

`$ snyk iac describe --from="tfstate+s3://my-bucket/path/to/state.tfstate"`

`--from` 옵션에 나열하여 여러 Terraform 상태 파일을 집계할 수도 있습니다. 로컬 디렉토리를 스캔하여 다양한 파일을 선택하거나, 다양한 소스로부터 다양한 경로를 사용할 수 있습니다. 두 가지 특정 Terraform 상태를 선택하려면 다음을 실행합니다:

`$ snyk iac describe --from="tfstate://path/to/terraform_S3.tfstate,tfstate://path/to/terraform_VPC.tfstate"`

### 드리프트 결과 및 다음 단계

#### 베이스라인 작성

한 번 `snyk iac describe`를 실행하고 현재 IaC 인프라의 커버리지 보고서를 받았습니다. 확인한 모든 문제를 해결하고 변경하지 않을 알려진 차이점만 남았을 때, 차이점을 다음 스캔에서 표시되지 않도록 하는 베이스라인을 만들 수 있습니다.

두 가지 옵션이 있습니다: 특정 리소스를 무시하거나 여러 리소스를 무시합니다.

**여러 리소스 무시**

`describe` 명령의 출력을 사용하여 [`.snyk 제외 정책을 업데이트](../../../../snyk-cli/commands/iac-update-exclude-policy.md)합니다:

`$ snyk iac describe --json --all | snyk iac update-exclude-policy`

**특정 리소스 무시**

특정 리소스를 무시하려면 `.snyk` 파일을 수정하여 리소스 세부 정보를 `exclude` 목록에 추가해야 합니다. 자세한 정보는 [드리프트를 위한 리소스 무시](ignore-resources-for-drift.md)를 참조하십시오.

이제 `snyk iac describe`를 주기적으로 실행하여 새로운 리소스가 IaC 배포 외부에서 생성될 때 알림을 받을 수 있습니다.