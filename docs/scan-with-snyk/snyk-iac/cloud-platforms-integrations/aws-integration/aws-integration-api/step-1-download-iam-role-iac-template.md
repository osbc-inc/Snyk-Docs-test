# 단계 1: IAM 역할 IaC 템플릿 다운로드 (API)

클라우드 환경을 생성하기 전에,  (IaC) 템플릿을 다운로드해야 합니다. 이 템플릿은 읽기 전용 **Identity and Access Management** (IAM) 역할을 선언하며, 이 역할은 Snyk가 AWS 계정의 리소스 구성을 스캔할 수 있도록 할 수 있습니다.

이 IaC 템플릿은 [단계 2: Snyk IAM 역할 생성](../aws-integration-web-ui/step-2-create-the-snyk-iam-role.md)에서 역할을 프로비전하는 데 사용됩니다.

템플릿 형식을 선택할 수 있으며, [Terraform HCL](https://www.terraform.io/language/syntax/configuration) 또는 [AWS CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html)을 선택할 수 있습니다. IAM 권한은 두 템플릿 모두에서 동일하므로 작업하기 가장 편한 형식을 선택하십시오.

## IaC 템플릿 검색

[Snyk API endpoint Generate Cloud Provider Permissions](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#post-/orgs/-org\_id-/cloud/permissions)을 사용하여 IaC 템플릿을 가져오려면, Org Admin 역할을 가진 Organization 수준의 [서비스 계정](../../../../../enterprise-setup/service-accounts/)에 대한 API 토큰이 필요합니다.

1. [Snyk Web UI](https://app.snyk.io)에서 **Settings (톱니바퀴 아이콘) > General > Organization ID**로 이동하여 조직 ID를 복사합니다.
2. 다음 형식으로 Snyk API에 요청을 보내어 템플릿을 검색합니다. `INPUT-TYPE`을 Terraform의 경우 `tf`, CloudFormation의 경우 `cf`로 바꿉니다.

```bash
curl -X POST \
'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/permissions?version=2022-12-21~beta' \
-H 'Authorization: token YOUR-API-TOKEN' \
-H 'Content-Type:application/vnd.api+json' -d '{
    "data": {
        "attributes": {
            "type": "INPUT-TYPE",
            "platform": "aws"
        },
        "type": "permissions"
    }
}'
```

{% hint style="info" %}
앞의 예제는 [curl](https://curl.se/)을 사용했지만, [Postman](https://www.postman.com/) 또는 [HTTPie](https://httpie.io/)와 같은 API 클라이언트를 사용할 수 있습니다.
{% endhint %}

## API 응답 이해

응답은 아래와 같은 JSON 문서입니다 (길이가 줄어들었습니다).

Terraform 구성으로 예제 응답:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "00000000-0000-0000-0000-000000000000",
    "type": "permissions",
    "attributes": {
      "data": "data \"aws_iam_policy_document\"<...>",
      "type": "tf"
    }
  }
}
```

CloudFormation 템플릿으로 예제 응답:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "00000000-0000-0000-0000-000000000000",
    "type": "permissions",
    "attributes": {
      "data": "AWSTemplateFormatVersion:<...>",
      "type": "cf"
    }
  }
}
```

## JSON 이스케이프 해제

이전 출력의 `data.attributes.data` 필드는 IAM 역할과 정책을 포함하는 Terraform 또는 CloudFormation 템플릿을 포함하는 이스케이프된 JSON 문자열입니다.

리소스를 프로비저닝하는 데 템플릿을 사용하기 전에 JSON을 **이스케이프 해제**해야 합니다. 다음 방법으로 수행할 수 있습니다:

* [jq 사용](step-1-download-iam-role-iac-template.md#use-jq)
* [내용 수동 변환](step-1-download-iam-role-iac-template.md#transform-the-content-manually)

### `jq` 사용

1. [jq 다운로드 및 설치](https://stedolan.github.io/jq/download/)합니다.
2. 템플릿을 검색하기 위해 API 요청을 제출할 때, 다음을 명령어 끝에 추가합니다:

    ```
    | jq -r .data.attributes.data > snyk_iac_template
    ```

    이렇게 하면 형식에 맞게된 템플릿이 현재 작업 디렉토리에 `snyk_iac_template` 파일로 저장됩니다.
3. 파일을 `.tf` (Terraform) 또는 `.yaml` (CloudFormation) 확장명으로 변경합니다.

### 내용 수동 변환

1. API 응답에서 `data.attributes.data`의 내용을 복사합니다. 값의 시작과 끝에 있는 큰따옴표를 제외한 긴 문자열이어야 합니다. `data \"aws_iam_policy_document\"` (Terraform) 또는 `AWSTemplateFormatVersion` (CloudFormation)로 시작하는 문자열이어야 합니다.
2. 이 문자열을 [FreeFormatter.com](https://www.freeformatter.com/json-escape.html)과 같은 도구에 붙여넣어 JSON을 이스케이프 해제합니다.
3. 해제된 출력을 새로운 `.tf` 파일 (Terraform) 또는 `.yaml` 파일 (CloudFormation)로 저장합니다.

## 선택 사항: IAM 역할 이름 변경

기본적으로 Snyk IAM 역할의 이름은 `snyk-cloud-role`입니다. 조직이 특정 역할 이름 요구 사항을 가지고 있다면 Terraform 또는 CloudFormation 템플릿에서이 이름을 변경할 수 있습니다.

**Terraform**에서 역할 이름은 19번째 줄에 있습니다:

```
  name                = "snyk-cloud-role"
```

**CloudFormation**에서 역할 이름은 7번째 줄에 있습니다:

```
      RoleName: snyk-cloud-role
```

## 다음 단계

다음 단계는 다운로드한 템플릿을 사용하여 Snyk를 위해 IAM 역할 및 정책을 [생성하는 것](step-2-create-the-snyk-iam-role-api.md)입니다.