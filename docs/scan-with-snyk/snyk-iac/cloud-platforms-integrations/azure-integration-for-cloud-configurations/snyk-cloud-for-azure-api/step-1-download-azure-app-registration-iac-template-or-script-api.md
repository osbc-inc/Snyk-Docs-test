# 단계 1: Azure 앱 등록 IaC 템플릿 또는 스크립트 (API) 다운로드

Azure 구독에 대한 클라우드 환경을 생성하기 전에 다음 리소스를 선언하는 Terraform 인프라스트럭처 코드(IaC) 템플릿 또는 Azure CLI Bash 스크립트를 **다운로드**해야 합니다:

* [액티브 디렉터리(AD) 애플리케이션 등록](https://learn.microsoft.com/en-us/azure/active-directory/develop/app-objects-and-service-principals#application-registration)
* [패더레이션 ID 자격 증명](https://learn.microsoft.com/en-us/azure/active-directory/develop/workload-identity-federation)
* [서비스 프린시팔](https://learn.microsoft.com/en-us/azure/active-directory/develop/app-objects-and-service-principals#service-principal-object)

이 인프라는 구독의 리소스 구성을 스캔하려는 Snyk에 대한 읽기 전용 권한을 부여합니다.

Step 2에서 Entra ID 앱 등록 (API)을 생성할 때 다운로드한 IaC 템플릿이나 Bash 스크립트를 사용하여 인프라를 **프로비전**합니다.

두 방법은 동일한 인프라를 생성하므로 작업하기 가장 편한 방법을 선택하십시오.

## IaC 템플릿 또는 스크립트 검색

Snyk API에서 IaC 템플릿을 검색하려면 Snyk 조직 수준의 [서비스 계정](../../../../../enterprise-setup/service-accounts/)에 대한 API 토큰과 Org Admin 역할이 필요합니다.

또한 등록하려는 Azure 구독의 구독 ID와 테넌트 ID가 필요합니다. 이 정보는 [Azure 설명서에 설명된 방법](https://learn.microsoft.com/en-us/azure/azure-portal/get-subscription-tenant-id)으로 찾을 수 있습니다.

1. [Snyk 웹 UI](https://app.snyk.io/)에서 **설정 (톱니바퀴 아이콘) > 일반 > 조직 ID**로 이동하여 조직 ID를 복사합니다.
2. 다음 양식에 따라 Snyk API에 요청을 보냅니다. `INPUT-TYPE`을 Terraform의 경우 `tf`, Bash의 경우 `bash`로 대체하세요:

```
curl -X POST \
'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/permissions?version=2022-12-21~beta' \
-H 'Authorization: token YOUR-API-TOKEN' \
-H 'Content-Type:application/vnd.api+json' -d '{
    "data": {
        "attributes": {
            "options": {
              "subscription_id": "YOUR-SUBSCRIPTION-ID",
              "tenant_id": "YOUR-TENANT-ID"
            },
            "type": "INPUT-TYPE",
            "platform": "azure"
        },
        "type": "permissions"
    }
}'
```

{% hint style="info" %}
앞서 나온 예제는 [curl](https://curl.se/)을 사용했지만 [Postman](https://www.postman.com/)이나 [HTTPie](https://httpie.io/)와 같은 다른 API 클라이언트를 사용할 수 있습니다.
{% endhint %}

Azure CLI를 로컬에서 실행하는 대신 [Azure Cloud Shell](https://portal.azure.com/#cloudshell/)을 사용하여 Bash 스크립트를 실행할 계획이라면 위 curl 명령을 Cloud Shell에서 실행하세요.

## API 응답 이해

응답은 다음과 같이 요약된 JSON 문서입니다.

Terraform 구성을 포함하는 예시 응답:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "00000000-0000-0000-0000-000000000000",
    "type": "permissions",
    "attributes": {
      "data": "provider \"azuread\"<...>",
      "type": "tf"
    }
  }
}
```

Bash 스크립트를 포함하는 예시 응답:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "00000000-0000-0000-0000-000000000000",
    "type": "permissions",
    "attributes": {
      "data": "objectId=$(az ad app create <...>",
      "type": "bash"
    }
  }
}
```

## JSON 문자열 이스케이프 해제

이전 출력의 `data.attributes.data` 필드는 Entra ID 앱 등록, 패더레이션 ID 자격 증명, 서비스 프린시팔을 포함하는 이스케이프된 JSON 문자열입니다.

리소스를 프로비전하기 전에 JSON을 **이스케이프 해제**해야 합니다. 이를 수행하는 방법은 다음과 같습니다:

* [jq 사용](step-1-download-azure-app-registration-iac-template-or-script-api.md#use-jq)
* [내용 수동 변환](step-1-download-azure-app-registration-iac-template-or-script-api.md#transform-the-content-manually)

### `jq` 사용

1. [jq 다운로드 및 설치](https://stedolan.github.io/jq/download/)
2. 템플릿 검색 중 API 요청을 보낼 때 다음을 명령어 끝에 추가하세요:\
   `| jq -r .data.attributes.data > snyk_azure_permissions`\
   이 명령을 실행하면 올바르게 형식이 지정된 템플릿이 현재 작업 디렉토리에 `snyk_azure_permissions` 파일에 저장됩니다.
3. 파일을 `.tf` (Terraform) 또는 `.sh` (Bash) 확장자로 이름을 변경하세요.

### 내용 수동 변환

1. API 응답에서 `data.attributes.data`의 내용을 복사합니다. 값의 매우 처음과 매우 끝에 있는 이중 따옴표를 제외하고 복사합니다. 이렇게 하면 `provider \"azuread\"` (Terraform) 또는 `objectId=$(az ad app create` (Bash)로 시작하는 긴 문자열을 얻게 됩니다.
2. 이 문자열을 [FreeFormatter.com](https://www.freeformatter.com/json-escape.html)과 같은 도구에 붙여넣어 JSON을 이스케이프 해제하세요.
3. 이스케이프 해제된 출력을 새 `.tf` 파일 (Terraform) 또는 `.sh` 파일 (Bash)로 저장하세요.

## 다음은?

다음 단계는 다운로드한 템플릿 또는 스크립트를 사용하여 Snyk의 Entra ID 앱 등록, 패더레이션 ID 자격 증명, 서비스 프린시팔을 생성하는 것입니다.