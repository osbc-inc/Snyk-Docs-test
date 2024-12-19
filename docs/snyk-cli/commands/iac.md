# IaC

## 사용법

`snyk iac <COMMAND> [<OPTIONS>] [<PATH>]`

## 설명

`snyk iac` 명령은 Infrastructure as Code 파일에서 보안 문제를 찾아 보고하고; 인프라 구성의 이탈과 관리되지 않는 리소스를 감지, 추적 및 경보하며; .driftigore 파일을 생성합니다.

자세한 정보는 [Snyk CLI for IaC](https://docs.snyk.io/snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac)를 참조하십시오.

## `snyk iac` 명령어 및 도움말 문서

모든 `snyk iac` 명령어는 도움말 옵션과 함께 여기에 나열되어 있습니다:

* [iac test](iac-test.md); `iac test --help`: 알려진 보안 문제를 테스트합니다.
* [iac capture](iac-capture.md); `iac capture --help`: Terraform 상태 구성에 액세스하여 매핑 아티팩트를 생성합니다.
* [iac describe](iac-describe.md); `iac describe --help`: 인프라 구성의 이탈과 관리되지 않는 클라우드 리소스를 감지합니다.\
  예시: `snyk iac describe --only-unmanaged`
* [iac rules init](iac-rules-init.md); `iac rules init --help`: 새로운 사용자 정의 규칙 프로젝트 구조, 기존 사용자 정의 규칙 프로젝트의 새로운 규칙, 기존 사용자 정의 규칙 프로젝트의 새로운 사양 또는 기존 사용자 정의 규칙 프로젝트의 새로운 관계를 초기화합니다.
* [iac rules test](iac-rules-test.md); `iac rules test --help`: Rego로 작성된 모든 테스트를 실행합니다.
* [iac rules push](iac-rules-push.md); `iac rules push --help`: Rego로 작성된 규칙을 번들로 만들어 변경 사항을 Snyk 플랫폼에 업로드합니다.
* [iac update-exclude-policy](iac-update-exclude-policy.md); `iac update-exclude-policy --help`: 클라우드 리소스에 대한 `.snyk` 제외 사항을 자동으로 생성합니다.