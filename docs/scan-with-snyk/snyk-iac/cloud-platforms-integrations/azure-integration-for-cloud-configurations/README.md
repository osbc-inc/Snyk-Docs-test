# 클라우드 구성을 위한 Azure 통합

Snyk은 [Microsoft Azure](https://azure.microsoft.com/en-us/) 구독과 통합하여 클라우드 구성에서 문제를 찾고 취약성을 우선순위에 따라 설정하는 데 도움이 되는 클라우드 컨텍스트를 생성합니다.

다음 방법을 사용하여 Azure 구독을 Snyk에 연동할 수 있습니다:

* [Snyk 웹 UI](azure-integration-web-ui/)
* [Snyk API](snyk-cloud-for-azure-api/)

## 클라우드 구성을 위한 Azure 통합 사전 요구 사항

클라우드 구성을 위한 Azure 통합을 추가하려면 다음이 필요합니다:

* Snyk 비즈니스 또는 엔터프라이즈 [플랜](https://snyk.io/plans/)
* 적절한 기능 플래그가 할당된 새로운 Snyk 조직(Snyk 연락처에 의해 할당됨)
* Snyk 그룹 관리자 또는 조직 관리자 [역할](../../../../snyk-admin/user-roles/pre-defined-roles.md)
* 다음 리소스를 만들 권한을 갖는 [Microsoft Azure](https://azure.microsoft.com/en-us/) 구독 및 관련 자격 증명에 대한 액세스:
  * [Active Directory (AD) 응용 프로그램 등록](https://learn.microsoft.com/en-us/azure/active-directory/develop/app-objects-and-service-principals#application-registration)
  * [페더레이션 신원 자격 증명](https://learn.microsoft.com/en-us/azure/active-directory/develop/workload-identity-federation)\
    이 [리소스](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/application_federated_identity_credential#api-permissions)를 만드는 데 Terraform을 사용하는 경우, 사용자는 Application Administrator 또는 Global Administrator 디렉터리 역할 중 하나를 가져야 합니다.
  * 읽기 전용 권한을 가진 [서비스 주체](https://learn.microsoft.com/en-us/azure/active-directory/develop/app-objects-and-service-principals#service-principal-object)\
    이 [리소스](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/service_principal)를 만드는 데 Terraform을 사용하는 경우, 사용자는 Application Administrator 또는 Global Administrator 디렉터리 역할 중 하나를 가져야 합니다.
* [Terraform CLI](https://www.terraform.io/downloads) 또는 Azure CLI([로컬](https://learn.microsoft.com/en-us/cli/azure/) 또는 [Cloud Shell](https://portal.azure.com/#home))에 액세스하여 Terraform 또는 Bash 스크립트를 통해 위의 리소스를 생성할 수 있어야 합니다.\
  Terraform 또는 Azure CLI를 로컬에서 사용하는 경우 Azure 자격 증명으로 구성해야 합니다. [Terraform](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs#authenticating-to-azure-active-directory) 또는 [Azure CLI](https://learn.microsoft.com/en-us/cli/azure/authenticate-azure-cli)의 지침을 참조하십시오.
* **API 전용:** Snyk API를 사용하기 위한 조직 수준의 [서비스 계정](../../../../enterprise-setup/service-accounts/)이 필요합니다.
* **API 전용:** [curl](https://curl.se/), [HTTPie](https://httpie.io/), 또는 [Postman](https://www.postman.com/)과 같은 API 클라이언트가 필요합니다.
* **API 전용, 선택 사항:** [jq](https://stedolan.github.io/jq/)는 Terraform 템플릿이나 Bash 스크립트가 포함된 JSON을 해제하는 데 사용됩니다.