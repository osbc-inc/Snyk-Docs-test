# 지원되는 IaC 언어, 클라우드 공급자 및 클라우드 리소스

{% hint style="info" %}
기타 Snyk 제품에서 지원되는 환경에 대해서는 [지원되는 언어 및 프레임워크](../../../supported-languages-package-managers-and-frameworks/)를 참조하십시오.
{% endhint %}

Snyk IaC를 사용하여 Terraform, Kubernetes Manifests, AWS CloudFormation, Azure Resource Manager (ARM) 및 Helm Charts를 위한 인프라 구성 파일의 문제를 찾고 보고 수정할 수 있습니다.

## 지원되는 IaC 언어

현재 IaC 대 IaC+의 언어 지원은 아래에 문서화되어 있습니다:

<table data-header-hidden><thead><tr><th width="271"></th><th width="218.33333333333331"></th><th width="224"></th></tr></thead><tbody><tr><td></td><td><strong>현재 IaC 지원</strong></td><td><strong>IaC+ 지원</strong></td></tr><tr><td><strong>Terraform (단일 파일)</strong></td><td>예</td><td>예</td></tr><tr><td><strong>Terraform (모듈)</strong></td><td>아니요</td><td>예</td></tr><tr><td><strong>Terraform (변수 파일)</strong></td><td>예 (CLI만)</td><td>예</td></tr><tr><td><strong>AWS CloudFormation</strong></td><td>예</td><td>예</td></tr><tr><td><strong>Azure Resource Manager</strong></td><td>예 (CLI만)</td><td>예 (모든 워크플로)</td></tr><tr><td><strong>Kubernetes Manifests</strong></td><td>예</td><td>예</td></tr><tr><td><strong>Helm Charts</strong></td><td>예 (SCM만)</td><td>곧 출시될 예정</td></tr></tbody></table>

### Terraform 모듈 스캐닝

기본적으로 IaC+에서는 Terraform 모듈을 스캔할 수 있습니다. 이를 위해 모듈이 로컬 파일로 이용 가능해야 합니다.'terraform' 디렉토리에 있는 경우입니다. 모듈을 가져오고( `terraform init` 사용) 스캔하려면 다음 명령을 실행하세요:

```
terrraform init && snyk iac test
```

## 지원되는 클라우드 제공업체

는 다음 클라우드 제공업체의 스캔을 지원합니다:

* [Amazon Web Services](../cloud-platforms-integrations/aws-integration/)
* [Azure](../cloud-platforms-integrations/azure-integration-for-cloud-configurations/)
* [Google Cloud](../cloud-platforms-integrations/google-cloud-integration/)

## 클라우드 스캔을 위한 지원되는 리소스

지원되는 클라우드 리소스 유형 목록을 확인하려면 각 클라우드 제공자의 문서를 참조하십시오:

* [AWS](supported-aws-resources-for-cloud-context.md)
* [Azure](supported-azure-resources-for-cloud-context.md)
* [Google](supported-google-resources-for-cloud-context.md)
