# 단계 3: Google을 위한 클라우드 환경을 생성하고 스캔

{% hint style="info" %}
**요약**\
Snyk를 위한 Google 서비스 계정을 생성했습니다. 이제 클라우드 환경을 생성하고 스캔할 수 있습니다.
{% endhint %}

[Snyk API](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#post-/orgs/-org\_id-/cloud/environments)에 요청을 보내어 클라우드 환경을 생성하고 스캔하려면 API 요청 본문에서 **Google 서비스 계정의 이메일 주소**와 **프로젝트 ID**를 제공해야 합니다.

## Snyk API 요청 보내기

다음 형식으로 Snyk API에 요청을 보내어 클라우드 환경을 만듭니다:

```bash
curl -X POST \
'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/environments?version=2022-12-21~beta' \
-H 'Authorization: token YOUR-API-TOKEN' \
-H 'Content-Type:application/vnd.api+json' -d '{
  "data": {
    "type": "environments",
    "attributes": {
      "options": {
        "identity_provider": "YOUR-IDENTITY-PROVIDER-URL",
        "service_account_email": "YOUR-SERVICE-ACCOUNT-EMAIL",
        "project_id": "YOUR-PROJECT-ID"
      },
      "kind": "google"
    }
  }
}'
```

{% hint style="info" %}
앞선 예시는 [curl](https://curl.se/)을 사용했지만 [Postman](https://www.postman.com/) 또는 [HTTPie](https://httpie.io/)와 같은 다른 API 클라이언트를 사용할 수 있습니다.
{% endhint %}

## API 응답 이해

응답은 새로 생성된 클라우드 환경에 대한 세부 정보가 포함된 JSON 문서입니다. 예를 들어:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "c61079b9-7260-4b2f-1234-abcd1234abcd",
    "type": "environment",
    "attributes": {
      "name": "my-project-demo",
      "options": {
        "project_id": "my-project-demo",
        "service_account_email": "snyk-cloud-mt-us-abcd1234@my-project-demo.iam.gserviceaccount.com"
      },
      "native_id": "my-project-demo",
      "properties": {
        "project_id": "my-project-demo",
        "project_name": "my-project-demo",
        "project_number": "123456789012"
      },
      "kind": "google",
      "revision": 1,
      "created_at": "2022-10-13T20:45:19Z",
      "status": "in_progress",
      "updated_at": "2022-10-13T20:45:19Z"
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

Snyk는 환경이 생성되면 자동으로 스캔을 시작합니다.

{% hint style="info" %}
JSON 출력의 `data.attributes.status` 필드가 `in_progress`로 설정됩니다. 이는 Snyk가 환경을 만들고 스캔을 시작했음을 의미합니다.
{% endhint %}

스캔이 완료되었는지 확인하려면 [스캔이 완료되었는지 확인](https://docs.snyk.io/integrations/cloud-platforms/aws-integration/snyk-cloud-for-aws-api/step-3-create-and-scan-a-snyk-cloud-environment#check-to-see-if-the-scan-is-finished)을 참조하십시오.

환경을 다시 스캔하려면 [클라우드 환경 스캔](../../../../../scan-with-snyk/snyk-iac/getting-started-with-iac+-and-cloud-scans/snyk-environments/scan-a-cloud-environment.md)을 참조하십시오.

{% hint style="info" %}
Google은 서비스 계정을 생성하는 데 60초 이상이 걸릴 수 있습니다. 서비스 계정을 생성한 후 즉시 환경을 만들려고 시도하면 **자격 증명을 확인할 수 없음 오류**가 발생할 수 있으니, 최소 60초를 기다린 후 다시 시도하십시오.
{% endhint %}

## 다음 단계는?

이제 다음을 수행할 수 있습니다:

* Snyk가 발견한 클라우드 구성 문제를 확인합니다. [클라우드 및 IaC+ 문제](../../../../../scan-with-snyk/snyk-iac/getting-started-with-iac+-and-cloud-scans/manage-iac+-and-cloud-issues/)를 참조하십시오.
* 클라우드 컨텍스트로 취약점을 우선순위대로 지정합니다.