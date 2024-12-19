# 단계 3: 클라우드 환경 (API) 생성 및 스캔

{% hint style="info" %}
**요약**\
Snyk 클라우드 IAM 역할을 생성했습니다. 이제 클라우드 환경을 생성하고 스캔할 수 있습니다.
{% endhint %}

클라우드 환경을 생성하고 스캔하기 위해 [Snyk API endpoint Create New Environment](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#post-/orgs/-org\_id-/cloud/environments)로 요청을 보내려면 API 요청 본문에 역할의 Amazon 리소스 이름 (ARN)을 제공해야 합니다.

## 역할 ARN 찾기

[역할 ARN 찾기](../aws-integration-web-ui/step-3-create-and-scan-a-snyk-cloud-environment-web-ui.md#find-the-role-arn)의 단계를 따르고 나서 여기로 돌아와 Snyk API 요청을 보내는 방법을 살펴보세요.

## Snyk API 요청 보내기

역할 ARN이 준비되면 다음 형식으로 클라우드 환경을 생성하는 Snyk API에 요청을 보냅니다.

```
curl -X POST \
'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/environments?version=2022-12-21~beta' \
-H 'Authorization: token YOUR-API-TOKEN' \
-H 'Content-Type:application/vnd.api+json' -d '{
  "data": {
    "attributes": {
      "kind": "aws",
      "name": "Example AWS Environment",
      "options": {
        "role_arn": "YOUR-ROLE-ARN"
      }
    },
    "type": "environment"
  }
}'
```

{% hint style="info" %}
위 예제는 [curl](https://curl.se/)을 사용했지만 [Postman](https://www.postman.com/) 또는 [HTTPie](https://httpie.io/)와 같은 다른 API 클라이언트를 사용할 수 있습니다.
{% endhint %}

## API 응답 이해

응답은 새로 생성한 클라우드 환경에 대한 세부 정보를 포함하는 JSON 문서입니다. 예를 들어:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "3b7ccff9-8900-4e54-0000-1234abcd1234",
    "type": "environment",
    "attributes": {
      "name": "Example AWS Environment",
      "options": {
        "role_arn": "arn:aws:iam::123412341234:role/snyk-cloud-role"
      },
      "native_id": "123412341234",
      "properties": {
        "account_id": "123412341234",
        "account_alias": "example"
      },
      "kind": "aws",
      "revision": 1,
      "created_at": "2022-07-31T00:50:49Z",
      "status": "in_progress",
      "updated_at": "2022-07-31T00:50:49Z"
    },
    "relationships": {
      "organization": {
        "data": {
          "id": "d70c1768-5675-0000-1234-abcd1234abcd",
          "type": "organization"
        },
        "links": {
          "related": "/orgs/d70c1768-5675-0000-1234-abcd1234abcd?version=2022-12-21~beta"
        }
      }
    }
  }
}
```

Snyk은 환경을 생성하면 자동으로 스캔을 시작합니다.

{% hint style="info" %}
JSON 출력의 `data.attributes.status` 필드는 `in_progress`로 설정됩니다. 이는 Snyk가 환경을 생성하고 스캔을 시작했음을 의미합니다.
{% endhint %}

## 스캔이 완료되었는지 확인

스캔이 완료되었는지 확인하려면 환경 세부 정보를 얻기 위해 다음 형식의 API 요청을 보내어 스캔이 완료되었는지 확인할 수 있습니다. 환경을 생성할 때 JSON 출력의 `data.id` 필드에서 환경 ID를 찾을 수 있습니다.

```
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/environments?id=YOUR-ENVIRONMENT-ID&version=2022-12-21~beta' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

만약 JSON 출력에서 `data.attributes.status` 필드가 `success`로 설정되어 있다면, Snyk가 환경을 스캔을 완료했음을 나타냅니다.

환경을 다시 스캔하려면 [클라우드 환경 스캔](../../../../../scan-with-snyk/snyk-iac/getting-started-with-iac+-and-cloud-scans/snyk-environments/scan-a-cloud-environment.md)을 참조하세요.

## 다음은 무엇인가요?

이제 다음을 수행할 수 있습니다:

- Snyk가 찾은 클라우드 구성 문제를 확인하세요. [클라우드 및 IaC+ 문제](../../../../../scan-with-snyk/snyk-iac/getting-started-with-iac+-and-cloud-scans/manage-iac+-and-cloud-issues/)를 참조하세요.
- 클라우드 컨텍스트를 사용하여 취약점을 우선 순위로 지정하세요.