# Terraform Cloud 통합을 통한 Snyk IaC 실행 작업

{% hint style="info" %}
Terraform Cloud 실행 작업은 Terraform Cloud Team 및 Governance 티어에서 사용할 수 있습니다. Terraform Cloud 무료 티어는 실행 작업을 지원하지 않습니다.
{% endhint %}

## Terraform Cloud 개요

[Terraform Cloud](https://www.terraform.io/cloud) (TFC)는 HashiCorp에서 제공하는 유상 SaaS 플랫폼으로, Terraform을 사용하는 팀에게 제품 준비 상태 관리 및 지속적 전달 기능을 제공합니다. 이를 통해 팀은 다음을 수행할 수 있습니다:

- 버전 관리 및 기본 제공을 통한 클라우드에서 Terraform 상태 관리
- 인프라에 대한 변경 사항을 검토하고 승인하여 인프라에 대한 변경 사항을 승인하는 팀의 중앙 집중화된 장소 보유
- Terraform을 사용하여 클라우드 인프라에 변경 사항 적용을 자동화하는 방식으로 클라우드 공급 업체에 대한 원격 작업을 Terraform Cloud가 관리하는 것과 같이 CI/CD 파이프라인과 유사한 방식으로 클라우드 인프라 관리

## **Terraform Cloud 내 Snyk 통합 개요**

Terraform Cloud에는 Run Tasks 기능이 제공됩니다. 이 기능은 Terraform 계획에 대한 특정 권한을 가진 고객에게 제공됩니다. 이 권한은 Team 플랜 이상에 선택적으로 추가할 수 있습니다. TFC에서 "run"은 결국 검토, 승인 및 적용할 Terraform 계획을 생성하는 TFC에서의 실행 단위를 나타냅니다.

Run Tasks 기능은 외부 통합이 "run" 이벤트에 연결되어 상호작용하고, 통과 또는 실패 여부를 결정하기 위한 상태를 제공할 수 있도록 허용합니다.

Snyk은 2021년 5월에 [Terraform 사용자가 주요 클라우드 공급업체의 Snyk 보안 정책에 대해 Terraform 계획 JSON 출력를 스캔할 수 있도록](https://snyk.io/blog/prevent-cloud-misconfigurations-hashicorp-terraform-snyk-iac/) 지원을 도입했습니다.

Snyk 통합은 Terraform Cloud의 "run" 워크플로우를 Snyk Terraform 계획 스캔과 연결하며, 생성된 각 "run"마다 Snyk는 미스구성에 대해 Terraform 계획 자산을 스캔합니다.

Terraform Cloud 사용자([Team 및 Governance 티어](https://www.hashicorp.com/products/terraform/pricing)에 해당)는 Snyk에 가입하여 Terraform Cloud의 작업 공간과 연결하고 통합 설정을 할 수 있습니다. 그 후 Run Task 권한 추가 기능이있는 경우 Terraform Cloud를 통해 소프트웨어 개발 라이프사이클의 일부로 보안 미스구성을 추적, 관리 및 해결할 수 있습니다.

설정 및 사용 세부 정보는 다음 페이지를 참조하십시오:

- [IaC를 위한 Terraform Cloud 통합 설정](set-up-the-terraform-cloud-integration-for-iac.md)
- [IaC를 위한 Terraform Cloud 통합 사용 방법](how-to-use-the-terraform-cloud-integration-for-iac.md)