# Snyk IaC를 위한 Terraform Enterprise 통합

## Terraform Enterprise 개요

[HashiCorp의 Terraform Enterprise](https://www.terraform.io/enterprise) (TFE)는 [Terraform Cloud](https://cloud.hashicorp.com/products/terraform) 응용 프로그램의 개인 인스턴스를 제공하는 제품으로, 리소스 제한이 없으며 감사 로깅 및 SAML 단일 로그인과 같은 기업급 아키텍처 기능을 추가로 제공합니다.

## Terraform Enterprise와 Snyk 통합 개요

Snyk는 Terraform Enterprise와의 통합이 Terraform Cloud를 위한 Snyk 통합과 정확히 동일하게 작동합니다. 통합 설정 방법에 대한 자세한 내용은 [Snyk integration for Terraform Cloud](terraform-cloud-integration-for-snyk-iac-using-run-tasks/) 페이지를 참조하십시오.

## Terraform Enterprise 통합을 위한 네트워크 요구 사항

Terraform Enterprise의 Snyk 통합은 Terraform Enterprise 인스턴스와 Snyk 플랫폼 간의 네트워크 연결성을 필요로 합니다. [Snyk integration for Terraform Cloud](terraform-cloud-integration-for-snyk-iac-using-run-tasks/)에 설명된 대로 통합 설정을 시도했지만 통합이 작동하지 않는 경우, 다음 단계가 문제를 식별하는 데 도움이 될 수 있습니다:

* Terraform Enterprise 인스턴스에서 Snyk API로의 연결을 확인하려면:
  * Terraform Enterprise 서버에 로그인합니다.
  * Snyk API 엔드포인트로 HTTP 요청을 수행합니다; 예를 들어 **curl**을 사용하여 HTTP 요청을 시작할 수 있습니다.
  * Snyk API 엔드포인트에서 401/Unauthorized 응답을 받았더라도 그것은 수용할 수 있는 응답입니다. 기본 네트워크 연결성만을 확인하고 있습니다.
* 또한 Snyk 서버가 Terraform Enterprise 인스턴스에 도달할 수 있도록 해야 합니다. 네트워크 구성(방화벽 등)이 Snyk 플랫폼으로부터의 네트워크 트래픽 수신을 허용하도록 설정되어 있는지 확인하십시오.