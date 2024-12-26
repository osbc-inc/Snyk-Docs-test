# 드리프트 감지 및 수동으로 생성된 리소스

는 관리되지 않는 리소스의 드리프트를 보고할 수 있습니다. 관리되지 않는 리소스란 클라우드 제공업체에는 있지만 Terraform 상태에는 없는 리소스를 말합니다. 이러한 리소스를 Terraform으로 가져오거나 IaaS 계정에서 삭제할 수 있습니다.

관리되는 리소스의 드리프트를 감지하는 정보는 [Command: plan](https://developer.hashicorp.com/terraform/cli/commands/plan)을 Terraform CLI 문서에서 확인하실 수 있습니다.&#x20;

이 그룹의 페이지에서 제공되는 정보는 'snyk iac describe' 명령어의 사용을 지원합니다. 다음과 같은 정보를 제공합니다:

* [AWS에서  Describe 시작하기](get-started-with-snyk-iac-describe-on-aws.md)
* [클라우드 제공업체 구성](configure-cloud-providers/)
* [지원되는 리소스](supported-resources/)
* [IaC describe 명령어 예시](iac-describe-command-examples.md)
* [필터 규칙](filter-rules.md)
* [드리프트를 위한 리소스 무시하기](ignore-resources-for-drift.md)
* [IAC 소스 사용법](iac-sources-usage.md)