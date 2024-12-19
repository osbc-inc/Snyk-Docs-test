# AWS 통합

Snyk은 클라우드 설정에서 문제를 찾고 취약점을 우선순위를 정할 수 있도록 클라우드 컨텍스트를 생성하기 위해 귀하의 [Amazon Web Services (AWS)](https://aws.amazon.com/) 계정과 통합됩니다.

다음 방법을 사용하여 AWS 계정을 Snyk에 등록할 수 있습니다:

* [Snyk 웹 UI](aws-integration-web-ui/)
* [Snyk API](aws-integration-api/)

다음은 **AWS 통합에 필요한 사전 조건**입니다:

AWS 통합을 설정하려면 다음이 필요합니다:

* Snyk 엔터프라이즈 [플랜](https://snyk.io/plans/)
* Snyk 문의를 통해 적합한 기능 플래그가 할당된 새로운 Snyk 조직
* Snyk 그룹 관리자 또는 조직 관리자 [역할](../../../../snyk-admin/user-roles/pre-defined-roles.md)
* 읽기 전용 IAM 역할을 생성할 권한이 있는 [AWS](https://aws.amazon.com/) 계정 및 관련 자격 증명에 대한 액세스
* Snyk를 위해 IAM 역할을 생성하기 위해 Terraform 또는 AWS CloudFormation을 통해 Terraform CLI, AWS CLI 또는 AWS Management Console에 액세스
* Terraform 또는 AWS CLI를 사용하는 경우 AWS 자격 증명으로 구성했는지 확인하십시오. [Terraform](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#authentication-and-configuration)나 [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html)에 대한 지침을 참조하십시오.
* **API 전용:** Snyk API를 사용하기 위해 조직 수준의 [서비스 계정](../../../../enterprise-setup/service-accounts/)이 있어야 하며 Org Admin 역할을 사용해야 합니다.
* **API 전용:** [curl](https://curl.se/), [HTTPie](https://httpie.io/), [Postman](https://www.postman.com/)과 같은 API 클라이언트가 필요합니다.