# 환경 삭제

환경을 삭제하면 Snyk는 모든 관련 스캔, 이슈 및 리소스 기록을 제거합니다. 클라우드 환경의 경우 클라우드 제공업체에 있는 실제 리소스에는 영향을 주지 않습니다.

## 웹 UI

Snyk 웹 UI에서는 조직 **설정 (톱니바퀴 아이콘) > 클라우드 환경**으로 이동하여 환경을 제거할 수 있습니다. [인프라 코드+ 또는 클라우드 환경 제거](view-add-and-remove-environments.md#remove-an-iac+-or-cloud-environment)를 참조하십시오.

## API

Snyk API를 사용하여 환경을 제거하려면 다음 형식으로 [환경 삭제](https://apidocs.snyk.io/#delete-/orgs/-org\_id-/cloud/environments/-environment\_id-) 엔드포인트에 요청을 보냅니다. 환경 ID는 [환경 ID 찾기](find-an-environment-id.md) 페이지에 표시된 방법을 사용하여 찾을 수 있습니다.

```
curl -X DELETE \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/environments/YOUR-ENVIRONMENT-ID?version=2022-12-21~beta' \
  -H 'Authorization: token YOUR-SERVICE-ACCOUNT-TOKEN'
```

명령이 성공하면 출력이 없습니다.

## Snyk AWS IAM 역할 제거

환경을 제거해도 Snyk Identity and Access Management (IAM) 역할은 제거되지 않습니다. Amazon Web Services (AWS) 계정에 대한 Snyk의 액세스를 완전히 제거하려면 역할을 만든 동일한 인프라 코드 도구를 사용하여 IAM 역할을 삭제합니다:

* **Terraform:** [terraform destroy](https://www.terraform.io/cli/commands/destroy) 명령을 사용하여 역할을 삭제합니다.
* **CloudFormation:** [AWS Management Console](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-console-delete-stack.html) 또는 AWS CLI의 [delete-stack 명령](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/cloudformation/delete-stack.html)을 사용하여 CloudFormation 스택을 삭제합니다.

## Google 서비스 계정 제거

환경을 제거해도 Google 서비스 계정은 제거되지 않습니다. Google 프로젝트에 대한 Snyk의 액세스를 완전히 제거하려면 [Google Cloud 콘솔](https://cloud.google.com/iam/docs/creating-managing-service-accounts#iam-service-accounts-delete-console)이나 [CLI](https://cloud.google.com/iam/docs/creating-managing-service-accounts#iam-service-accounts-delete-gcloud)를 사용하여 Google 서비스 계정을 삭제합니다.

## Azure 앱 등록, 연합 식별 자격 증명 및 서비스 주체 제거

환경을 제거해도 Azure 앱 등록, 연합 식별 자격 증명 또는 서비스 주체가 제거되지 않습니다. Snyk의 Azure 구독 액세스를 완전히 제거하려면 다음과 같은 방법으로 인프라를 삭제합니다:

* **Terraform:** [terraform destroy](https://www.terraform.io/cli/commands/destroy) 명령을 사용합니다.
* **Azure CLI Bash 스크립트:** 먼저 앱 등록에 대한 리더 역할 할당을 삭제한 후 (자세한 내용은 문서: [Azure CLI](https://learn.microsoft.com/en-us/azure/role-based-access-control/role-assignments-remove#azure-cli) 또는 [Azure Portal](https://learn.microsoft.com/en-us/azure/role-based-access-control/role-assignments-remove#azure-portal) 참조), Azure 앱 등록을 삭제합니다 (자세한 내용은 문서: [Azure CLI](https://learn.microsoft.com/en-us/cli/azure/ad/app?view=azure-cli-latest#az-ad-app-delete) 또는 [Azure Portal](https://learn.microsoft.com/en-us/azure/active-directory/develop/howto-remove-app#remove-an-application-authored-by-you-or-your-organization) 참조).