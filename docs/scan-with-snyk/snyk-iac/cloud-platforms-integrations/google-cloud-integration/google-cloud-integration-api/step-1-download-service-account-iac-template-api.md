# 단계 1: 서비스 계정 IaC 템플릿(API) 다운로드

구글 프로젝트의 리소스 구성을 스캔하기 위한 권한을 제공하는 타이트한 범위의 구글 서비스 계정을 선언하는 인프라 구성(IaC) 템플릿을 다운로드해야 Cloud Environment를 생성할 수 있습니다.

이 템플릿은 구글 클라우드 프로젝트를 위해 일련의 [구글 서비스 API](https://cloud.google.com/service-usage/docs/enabled-service)를 활성화합니다. 이는 Snyk가 프로젝트 리소스를 스캔하기 위해 필요한 API를 사용할 수 있도록 보장합니다.

이 IaC 템플릿을 사용하여 [단계 2: 구글 서비스 계정 생성(API)](step-2-create-the-google-service-account-api.md)에서 역할을 프로비저닝할 것입니다.

## IaC 템플릿 검색

[Snyk API](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#post-/orgs/-org\_id-/cloud/permissions)에서 IaC 템플릿을 검색하려면 로그인한 상태에서 Snyk 조직 수준 [서비스 계정](https://docs.snyk.io/features/user-and-group-management/structure-account-for-high-application-performance/service-accounts#set-up-a-service-account)에 대한 API 토큰이 필요합니다. 이 서비스 계정은 Org Admin 역할을 갖습니다.

1. [Snyk 웹 UI](https://app.snyk.io/)에서 **Settings (톱니바퀴 아이콘) > General > Organization ID**로 이동하여 조직 ID를 복사합니다.
2. 다음 형식으로 Snyk API에 요청을 보냅니다:

```bash
curl -X POST \
'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/permissions?version=2022-12-21~beta' \
-H 'Authorization: token YOUR-API-TOKEN' \
-H 'Content-Type:application/vnd.api+json' -d '{
    "data": {
        "attributes": {
            "type": "tf",
            "platform": "google"
        },
        "type": "permissions"
    }
}'
```

{% hint style="info" %}
앞의 예는 [curl](https://curl.se/)을 사용한 것이지만 [Postman](https://www.postman.com/) 또는 [HTTPie](https://httpie.io/)와 같은 다른 API 클라이언트를 사용할 수 있습니다.
{% endhint %}

응답은 다음과 같이 JSON 문서가 됩니다(길이가 줄어듦):

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "00000000-0000-0000-0000-000000000000",
    "type": "permissions",
    "attributes": {
      "data": "variable \"project_id\"<...>",
      "type": "tf"
    }
  }
}
```

## JSON 역 이스케이프

앞서 나온 출력의 `data.attributes.data` 필드는 Google 서비스 계정이 포함된 Terraform 템플릿이 포함된 이스케이프된 JSON 문자열입니다.

리소스를 프로비저닝하기 위해 템플릿을 사용하려면 JSON을 역 이스케이프해야 합니다. 이 작업은 다음 중 하나로 수행할 수 있습니다:

* [`jq` 사용](step-1-download-service-account-iac-template-api.md#use-jq)
* [내용 수동 변환](step-1-download-service-account-iac-template-api.md#transform-the-content-manually)

### `jq` 사용

1. [jq를 다운로드하고 설치](https://stedolan.github.io/jq/download/)합니다.
2. 템플릿을 검색하기 위해 API 요청을 제출할 때 명령어 끝에 다음을 추가합니다:

```bash
| jq -r .data.attributes.data > snyk_google_iac_template.tf
```

이렇게 하면 템플릿이 `snyk_google_iac_template.tf` 파일에 올바르게 포맷팅되어 현재 작업 디렉토리에 저장됩니다.

### 내용 수동 변환

1. API 응답에서 `data.attributes.data` 내용을 복사합니다(값의 가장 처음과 마지막의 이중 따옴표 제외). 이중 따옴표는 킴.
2. `FreeFormatter.com`과 같은 도구에 이 문자열을 붙여넣어 JSON을 역 이스케이프합니다.
3. 역 이스케이프된 Terraform 출력을 새 `.tf` 파일로 저장합니다.

## 다음 단계

다음 단계는 Snyk를 위해 구글 서비스 계정을 생성하는 것입니다.