# IaC+ 및 클라우드의 주요 개념

{% hint style="info" %}
Snyk IaC+는 현재 클로즈 베타 상태이며 새로운 고객을 받아들이지 않습니다.

\
&#xNAN;_&#xD604;재 IaC를 사용하여 시작하는 방법_ 에 대한 자세한 내용은 [여기](https://docs.snyk.io/scan-using-snyk/snyk-iac/getting-started-with-current-iac)를 참조하십시오.
{% endhint %}

IaC+ 및 클라우드 스캔에는 [환경](key-concepts-for-iac+-and-cloud.md#environments) 및 [자원](key-concepts-for-iac+-and-cloud.md#resources)과 같이 Snyk 코어 개념과는 다른 여러 독특한 개념이 있습니다.

## 환경

Snyk의 **환경**은 다음과 같이 일치하는 조직 개념입니다:

* IaC+ 환경의 경우: SCM 저장소, CLI 테스트 보고서 또는 Terraform Cloud/Enterprise 실행 작업 보고서
* 클라우드 환경의 경우: Amazon Web Services (AWS) 계정, Azure 구독 또는 Google Cloud 프로젝트.

Snyk [프로젝트](../../../snyk-admin/snyk-projects/#project)와 달리, 환경에는 [리소스](key-concepts-for-iac+-and-cloud.md#resources)라고 하는 스캔 가능한 엔티티가 포함됩니다. 자원은 서로 관련이 있을 수 있으며, 한 자원은 다른 자원의 하위 자원 또는 동일 레벨 자원일 수 있습니다. 또한 자원은 테스트할 수 있는 속성을 갖고 있으며, 이러한 속성이 잘못 구성되면 이슈가 생성됩니다. 이로써 환경과 해당 자원은 프로젝트와는 다른 개념입니다.

Snyk 환경은 클라우드 제공업체와의 통합 설정도 포함합니다. 예를 들어, 각 환경은 다른 AWS 계정과의 통합을 나타낼 수 있습니다.

`/cloud/environments` Snyk REST API 엔드포인트를 사용하여 모든 환경 목록을 검색하고 이름 및 스캔 상태와 같은 속성으로 필터링할 수 있습니다.

다음 클라우드 제공자가 지원됩니다:

* [Amazon Web Services](https://aws.amazon.com/)
* [Microsoft Azure](https://azure.microsoft.com/en-us/)
* [Google Cloud](https://cloud.google.com/)

## 리소스

**리소스**는 AWS S3 버킷, Identity and Access Management (IAM) 역할 또는 Virtual Private Cloud (VPC) 플로우 로그와 같은 클라우드 인프라 엔티티입니다.

각 스캔에서 Snyk는 환경의 각 리소스의 구성 속성을 기록합니다.

`/cloud/resources` Snyk REST API 엔드포인트를 사용하여 조직의 모든 리소스 목록을 검색하고 환경 ID, 리소스 ID 또는 리소스 유형과 같은 속성으로 필터링할 수 있습니다.

클라우드 환경에서 지원되는 리소스 유형 목록은 다음을 참조하십시오:

* [AWS 리소스 지원되는 리소스](../supported-iac-languages-cloud-providers-and-cloud-resources/supported-aws-resources-for-cloud-context.md)
* [Azure 리소스 지원되는 리소스](../supported-iac-languages-cloud-providers-and-cloud-resources/supported-azure-resources-for-cloud-context.md)
* [Google 리소스 지원되는 리소스](../supported-iac-languages-cloud-providers-and-cloud-resources/supported-google-resources-for-cloud-context.md)

## 리소스 매핑

리소스 매핑은 클라우드 리소스와 IaC 리소스 간의 연결을 나타냅니다. Snyk는 지역 또는 CI 파이프라인에서 `snyk iac capture` 명령을 실행할 때 생성된 Terraform 상태 파일에서 매핑 아티팩트로 이러한 연결을 결정합니다. 매핑 아티팩트에는 Snyk가 리소스 매핑을 유도하는 데 사용하는 자원 ID와 같은 세부사항이 포함됩니다. Snyk는 매핑 아티팩트가 만들어지고 업데이트될 때 또는 클라우드 환경이 만들어지고 업데이트될 때 매핑 실행을 트리거하며, 이로써 Snyk 조직의 리소스 매핑을 만들거나 업데이트하거나 삭제합니다. 자세한 내용은 [IaC에서 클라우드 이슈 해결](../iac+-code-to-cloud-capabilities/fix-cloud-issues-in-iac.md)에서 확인하십시오.

## 규칙

보안 **규칙**은 보안 문제로 이어질 수 있는 클라우드 인프라 및 인프라스트럭처 코드 (IaC)의 잘못된 구성을 확인합니다. Snyk에는 IaC+ 환경에 적용할 수 있는 미리 정의된 [보안 규칙](https://security.snyk.io/rules/cloud/)이 있습니다.

예를 들어, _S3 버킷에는 모든 블록 공개 액세스 옵션이 활성화되어 있지 않음_ 이라는 규칙이 있습니다. Snyk는 AWS S3 버킷의 구성을 스캔하여 해당 규칙을 위반하고 데이터 침해에 취약한지 확인할 수 있습니다.

## 이슈

**이슈**는 보안 문제로 이어질 수 있는 잘못된 구성을 나타냅니다. 이는 자원 및 규칙과 관련이 있습니다. 예를 들어, AWS S3 버킷은 _S3 버킷에 모든 블록 공개 액세스 옵션이 활성화되지 않음_ 규칙을 통해 테스트될 수 있습니다. 버킷이 해당 규칙을 위반하면 Snyk는 클라우드 이슈를 열게 됩니다.

Snyk가 이슈를 생성한 후에는 이슈가 해결되기까지 열어두고, 해당 문제가 수정된 시점에 이슈가 닫힙니다.

Snyk 웹 UI에서 조직의 이슈를 볼 수 있습니다. [Snyk 웹 UI에서 IaC+ 이슈 보기](manage-iac+-and-cloud-issues/view-iac+-and-cloud-issues-in-the-snyk-web-ui.md)를 참조하십시오.

## 규정 준수 표준 <a href="#docs-internal-guid-e2e38027-7fff-9271-f2c0-e23677542f6e" id="docs-internal-guid-e2e38027-7fff-9271-f2c0-e23677542f6e"></a>

**규정 준수 표준**은 조직이 IT 시스템 및 인프라를 보안하는 지침과 제어를 설정하는 프레임워크입니다. 규정 준수 표준은 버전이 매겨지며, 버전은 다양한 속도로 출시됩니다. 예: NIST 800-53 (vRev5), CIS AWS Foundations Benchmark (v1.4.0). Snyk는 [Cloud Compliance Issues 보고서](../../../manage-issues/reporting/available-snyk-reports.md#cloud-compliance-issues-report)를 제공합니다.

자세한 내용은 [지원되는 규정 준수 표준](../view-iac+-and-cloud-compliance-reporting.md#supported-compliance-standards)을 참조하십시오.

## 규정 준수 제어 <a href="#docs-internal-guid-11e1473c-7fff-ea66-c8f4-16a826a82e6b" id="docs-internal-guid-11e1473c-7fff-ea66-c8f4-16a826a82e6b"></a>

**규정 준수 제어**는 규정 준수 표준의 특정 권고사항 또는 지침으로 조직이 시스템 또는 인프라를 안전하게 보호해야 하는 방법을 규정합니다. 예: CIS AWS Foundations Benchmark (v1.4.0)의 control 2.1.5는 _S3 버킷이 '공개 액세스 차단 (버킷 설정)'으로 구성되었을 것을 보장합니다_ . 이 제어를 준수하려면 조직은 모든 S3 버킷에 대해 "공개 액세스 차단" 설정을 활성화해야 합니다.

## 규정 준수 매핑

Snyk는 보안 [규칙](key-concepts-for-iac+-and-cloud.md#rules)을 규정 준수 제어에 매핑합니다. 이는 각 규칙이 하나 이상의 제어와 관련되고, 각 제어가 하나 이상의 규칙과 관련되어 있음을 의미합니다.

예를 들어, CIS AWS Foundations Benchmark (v1.4.0)의 control 2.1.5는 _'S3 버킷이 '모두 공개 액세스 차단 옵션'이 활성화되지 않음'_ 이라는 보안 규칙 [SNYK-CC-00195](https://security.snyk.io/rules/cloud/SNYK-CC-00195)에 매핑됩니다.
