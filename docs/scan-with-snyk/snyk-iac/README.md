# Snyk IaC

Snyk 인프라스트럭처 코드(IaC)를 사용하면 클라우드 인프라 구성을 배포 전후에 보안할 수 있습니다.

{% hint style="info" %}
이 문서에서 **Current IaC**로 지칭되는 버전은 일반적으로 사용 가능한  버전입니다.

**IaC+**로 지칭되는 버전은 클로즈 베타로 제공되었고 해당 클로즈 베타에서 사용 중인 고객들을 대상으로합니다. 다른 고객은 Current IaC 문서를 참조해야 합니다. 추가 정보를 제공할 예정입니다.
{% endhint %}

{Snyk IaC의 두 버전 모두 다음을 수행할 수 있습니다:

- [HashiCorp Terraform](scan-your-iac-source-code/scan-terraform-files/)을 위한 안전한 구성 작성
- [AWS CloudFormation](scan-your-iac-source-code/scan-cloudformation-files/)을 위한 안전한 구성 작성
- [Kubernetes](scan-your-iac-source-code/scan-kubernetes-configuration-files/)와 [Azure Resource Manager (ARM)](scan-your-iac-source-code/scan-arm-configuration-files.md)을 위한 안전한 구성 작성 - IDE, SCM, CLI 및 Terraform Cloud/Enterprise 작업 흐름.
- 문제 보기 및 [수정 권고](getting-started-with-current-iac.md) 받아들여 코드에 직접 변경을 가해 애플리케이션이 제품 단계에 도달하기 전에 수정할 수 있습니다.
- 클라우드에서 [drift 감지](iac+-code-to-cloud-capabilities/detect-drift-and-manually-created-resources/) 및 수동으로 생성된 리소스 감지 가능
- AWS, Azure 및 Google Cloud 환경에서 미구성을 위해 클라우드 환경을 스캔하고 테스트하며 온보딩합니다.

IaC+는 {Snyk IaC의 클라우드 스캔 기능도 동일하게 구동하는 새 엔진 및 규칙 세트 위에 구축되어 다음을 가능하게 합니다:

- 모든 IaC 작업 흐름에서 Azure Resource Manager와 같은 언어에 대한 일관된 지원 포함
- Terraform에 대한 다중 파일 분석 추가 (모듈 및 변수 파일 지원)
- CIS 벤치마크, PCI, SOC 2 등 여러 준수 표준에 매핑된 확장된 보안 규칙 세트 사용
- 모든 IaC 작업 흐름에서 일관되게 작동하는 Rego를 사용하여 관리되는 사용자 지정 규칙 지원
- {Snyk Code와 일치하는 SCM에 대한 프로젝트 도입 (단일 IaC 파일이 아닌 저장소 전체의 문제 캡처)
- IaC+ SCM 프로젝트에 대한 주기적 (매일 또는 매주) 스캔 지원
- IaC+ 및 클라우드 문제에 대한 새로운 조직 전체 클라우드 문제 페이지 사용자는 해당 문제에 맞게 문제를 그룹화하고 필터링 및 검사하여 특정 문제에 대한 관련 리소스의 구성을 검사하고 문제에 대한 작업을 수행할 수 있습니다.

또한 IaC+는 "코드에서 클라우드" 사용 사례를 지원하며 다음 작업이 가능하도록합니다:

- 클라우드에서 생성된 클라우드 리소스에 대한 IaC 소스 코드에서 직접 [클라우드 문제 수정](iac+-code-to-cloud-capabilities/fix-cloud-issues-in-iac.md) 링크 클라우드 문제를 IaC 템플릿에 연결하여 잘못 구성된 클라우드 리소스를 배포한 IaC 소스 코드에서 직접 문제를 수정합니다.
- [배포된 인프라에서 컨텍스트 적용하여](iac+-code-to-cloud-capabilities/add-cloud-context-to-your-iac-tests.md) IaC 테스트의 거짓 양성 억제
- Terraform의 모든 워크플로우 (IaC에서 클라우드로)에 대해 전체 소프트웨어 개발 생명 주기 동안 동일한 사용자 정의 규칙 적용
- IaC 파일에서 생성된 IaC 및 클라우드 리소스 목록 보기 [List resources](https://apidocs.snyk.io/?version=2023-09-20%7Ebeta#get-/orgs/-org\_id-/cloud/resources)를 사용하여 API 엔드포인트에서 생성

지원되는 IaC 언어 및 클라우드 제공자 목록은 [지원되는 IaC 및 클라우드 제공자](supported-iac-languages-cloud-providers-and-cloud-resources/)를 참조하십시오.