# IaC+ 및 클라우드 리소스 보기

조직 내의 IaC+ 및 클라우드 리소스의 모든 속성을 볼 수 있습니다. 이를 통해 클라우드 공급업체 계정 전체의 모든 리소스를 인벤토리로 만들거나 최근 스캔 중에 어떤 리소스의 기록된 상태를 볼 수 있습니다.

## 모든 리소스 보기

조직 내의 모든 리소스를 나열하려면 다음 형식으로 `/cloud/resources` 엔드포인트에 요청을 보냅니다.

```bash
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/resources?version=2022-12-21~beta' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

## API 응답 이해

Snyk는 조직 내의 모든 리소스에 대한 정보를 포함하는 JSON 문서를 반환합니다. 예를 들어:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "links": {
    "next": "/rest/orgs/d70c1768-5675-4f44-0000-1234abcd1234/cloud/resources?starting_after=eyJpZCI6IjY5ODA5MjNhLWU0ZTAtNDg3Mi04ZDAwLWRjZDEXAMPLEEXAMPLE&version=2022-04-13~experimental"
  },
  "data": [
    {
      "id": "23b3a46d-cdf7-526c-8888-1abc2abc3abc",
      "type": "resource",
      "attributes": {
        "environment_id": "ef5d85de-fb4f-4c42-1234-000000000000",
        "scan_id": "44f386a6-6ce8-4303-0000-098765432109",
        "created_at": "2022-08-07T05:34:24.272279Z",
        "updated_at": "2022-08-07T05:34:24.272279Z",
        "revision": 1,
        "kind": "cloud",
        "hash": "717cdff4218bf3d6abcdefEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLEEXAMPLE",
        "platform": "aws",
        "namespace": "us-east-1",
        "resource_type": "aws_s3_bucket",
        "resource_id": "example-bucket",
        "native_id": "arn:aws:s3:::example-bucket",
        "name": "example-bucket",
        "location": "us-east-1",
        "state": {
          "id": "example-bucket",
          "acl": "private",
          "arn": "arn:aws:s3:::example-bucket"
          <trimmed for length>
        },
        "tags": {}
      },
      "relationships": {
        "environment": {
          "data": {
            "id": "ef5d85de-fb4f-4c42-1234-000000000000",
            "type": "environment"
          },
          "links": {
            "related": "/orgs/d70c1768-5675-4f44-0000-1234abcd1234/cloud/environments?id=ef5d85de-fb4f-4c42-1234-000000000000&version=2022-12-21~beta"
          }
        },
        "organization": {
          "data": {
            "id": "d70c1768-5675-4f44-0000-1234abcd1234",
            "type": "organization"
          },
          "links": {
            "related": "/orgs/d70c1768-5675-4f44-0000-1234abcd1234?version=2022-12-21~beta"
          }
        },
        "scan": {
          "data": {
            "id": "a7fa2167-58a8-4ac5-9999-0987dcba6543",
            "type": "scan"
          },
          "links": {
            "related": "/orgs/d70c1768-5675-4f44-0000-1234abcd1234/cloud/scans?id=a7fa2167-58a8-4ac5-9999-0987dcba6543&version=2022-12-21~beta"
          }
        }
      }
    }
    <trimmed for length>
  ]
}
```

다음 표는 API 응답에서 중요한 속성 몇 가지를 나열합니다:

| 속성                             | 정의                                                                                          |
| -------------------------------- | -------------------------------------------------------------------------------------------- |
| `data.id`                        | 리소스 ID                                                                                    |
| `data.attributes.environment_id` | 리소스를 포함하는 환경의 ID                                                                   |
| `data.attributes.scan_id`        | 이 리소스가 처음 감지된 스캔의 ID                                                             |
| `data.attributes.resource_type`  | 리소스 유형                                                                                  |
| `data.attributes.native_id`      | 리소스의 클라우드 공급자에 대한 고유 식별자; AWS의 경우 Amazon 리소스 이름 (ARN)            |
| `data.attributes.state`          | 가장 최근 스캔 시점에서의 리소스 속성                                                        |

## 리소스 목록 필터링

쿼리 매개변수를 사용하여 리소스 목록을 필터링할 수 있습니다.

예를 들어, 단일 환경에서의 리소스를 반환하려면 URL에 `environment_id=YOUR-ENVIRONMENT-ID`를 추가합니다:

```bash
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/resources?environment_id=YOUR-ENVIRONMENT-ID&version=2022-12-21~beta' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

일부 매개변수를 사용하여 여러 값을 지정할 수 있습니다. AWS 리전 `us-east-1` 또는 `us-east-2`의 리소스를 반환하려면 URL에 `location=us-east-1,us-east-2`를 추가합니다:

```bash
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/resources?location=us-east-1,us-east-2&version=2022-12-21~beta' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

`&` 기호를 사용하여 쿼리 매개변수를 결합할 수 있습니다. 예를 들어, AWS S3 버킷 중 5개만 반환하려면 URL에 `resource_type=aws_s3_bucket&limit=5`를 추가합니다:

```bash
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/resources?resource_type=aws_s3_bucket&limit=5&version=2022-12-21~beta' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

지원되는 매개변수 목록은 [리소스 목록 API 문서](https://apidocs.snyk.io/#get-/orgs/-org\_id-/cloud/resources)를 참조하십시오.

## 특정 리소스에 대한 세부 정보 보기

Snyk API를 통해 단일 리소스에 대한 세부 정보를 보려면 다음 형식으로 요청을 보냅니다.

```bash
curl -X GET \
  'https://api.snyk.io/rest/orgs/YOUR-ORGANIZATION-ID/cloud/resources?id=YOUR-RESOURCE-ID&version=2022-12-21~beta' \
  -H 'Authorization: token YOUR-API-TOKEN'
```

Snyk는 선택한 리소스에 대한 정보를 포함하는 JSON 문서를 반환합니다. 이 정보는 [API 응답 이해](view-iac+-and-cloud-resources.md#understand-the-api-response)에 표시된 것과 동일합니다.