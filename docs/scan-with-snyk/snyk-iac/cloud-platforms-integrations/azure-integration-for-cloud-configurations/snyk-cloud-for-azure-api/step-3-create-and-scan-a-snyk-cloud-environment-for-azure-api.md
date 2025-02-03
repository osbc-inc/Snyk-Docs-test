# 3단계: Azure용 클라우드 환경 만들기 및 스캔(API)

{% hint style="info" %}
**요약**\
Snyk을 위해 Azure 앱 등록, 연합 ID 자격증명, 및 서비스 주체를 생성했습니다. 이제 클라우드 환경을 생성하고 스캔할 수 있습니다.
{% endhint %}

Snyk API에 요청을 보내어 Azure 클라우드 환경을 생성하고 스캔하려면 API 요청 본문에 구독 ID, 테넌트 ID, 애플리케이션 ID를 제공해야 합니다.

## Snyk API 요청 보내기

클라우드 환경을 생성하기 위해 다음 형식의 Snyk API 요청을 보냅니다:

```
curl -X POST \
'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/environments?version=2022-12-21~beta' \
-H 'Authorization: token YOUR-API-TOKEN' \
-H 'Content-Type:application/vnd.api+json' -d '{
  "data": {
    "type": "environments",
    "attributes": {
      "options": {
        "subscription_id": "YOUR-SUBSCRIPTION-ID",
        "tenant_id": "YOUR-TENANT-ID",
        "application_id": "YOUR-APPLICATION-ID"
      },
      "kind": "azure",
      "name": "OPTIONAL-NAME"
    }
  }
}'
```

{% hint style="info" %}
위 예제는 [curl](https://curl.se/)을 사용하였지만 [Postman](https://www.postman.com/) 또는 [HTTPie](https://httpie.io/)와 같은 다른 API 클라이언트를 사용할 수 있습니다.
{% endhint %}

## API 응답 이해

응답은 새로 생성된 클라우드 환경에 대한 세부 정보를 포함하는 JSON 문서입니다. 예를 들어:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "e25a5ef1-1e96-1234-0000-1234abcd1234",
    "type": "environment",
    "attributes": {
      "name": "Example Azure Environment",
      "options": {
        "tenant_id": "00000000-0000-0000-1234-12345678abcd",
        "application_id": "12345678-1234-0000-0000-09876543dcba",
        "subscription_id": "abcd1234-abcd-1234-0000-000000000000"
      },
      "native_id": "abcd1234-abcd-1234-0000-000000000000",
      "properties": {
        "subscription_id": "abcd1234-abcd-1234-0000-000000000000",
        "subscription_name": "Example User"
      },
      "kind": "azure",
      "revision": 1,
      "created_at": "2023-02-06T06:34:05Z",
      "status": "in_progress",
      "updated_at": "2023-02-06T06:34:05Z"
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

환경을 생성하면 Snyk이 자동으로 스캔을 시작합니다.

{% hint style="info" %}
JSON 출력의 `data.attributes.status` 필드는 `in_progress`로 설정됩니다. 이는 Snyk이 환경을 생성하고 스캔을 시작했음을 의미합니다.
{% endhint %}

스캔이 완료되었는지 확인하려면 [스캔이 완료되었는지 확인](https://docs.snyk.io/integrations/cloud-platforms/getting-started-with-snyk-cloud-aws/snyk-cloud-for-aws-api/step-3-create-and-scan-a-snyk-cloud-environment#check-to-see-if-the-scan-is-finished)을 참조하십시오.

환경을 수동으로 다시 스캔하려면 [클라우드 환경 스캔](../../../getting-started-with-iac+-and-cloud-scans/snyk-environments/scan-a-cloud-environment.md)을 참조하십시오.

## 다음 단계는?

다음을 수행할 수 있습니다:

* Snyk이 발견한 클라우드 구성 문제를 확인하세요. [클라우드 및 IaC+ 문제](../../../getting-started-with-iac+-and-cloud-scans/manage-iac+-and-cloud-issues/)를 참조하세요.
* 클라우드 컨텍스트를 사용하여 취약점을 우선 순위로 지정하세요.
