# 클라우드 환경 스캔

Snyk는 클라우드 환경이 생성될 때 자동으로 스캔을 실행합니다. 그 이후로 Snyk는 24시간마다 한 번씩 환경을 스캔합니다. 또한 [Snyk API](https://apidocs.snyk.io/?version=2022-12-21%7Ebeta#post-/orgs/-org\_id-/cloud/scans)를 사용하여 언제든지 수동으로 새로운 스캔을 시작할 수도 있습니다.

## `jq` 사용

[jq 설치가 완료](https://stedolan.github.io/jq/download/)되어 있다면, [Snyk API 스캔 생성](https://apidocs.snyk.io/#post-/orgs/-org\_id-/cloud/scans) 엔드포인트에서 첫 번째 환경 ID를 검색하고 수동으로 스캔을 트리거하고 환경을 수동으로 스캔하려면 아래의 단일 명령을 실행할 수 있습니다:

```bash
SNYK_ORG_ID="당신의-조직-ID" && \
SNYK_API_TOKEN="당신의-API-토큰" && \
SNYK_ENV_ID=$(curl -s -X GET \
  "https://api.snyk.io/rest/orgs/${SNYK_ORG_ID}/cloud/environments?version=2022-12-21~beta<a data-footnote-ref href="#user-content-fn-1">&#x26;kind=aws,azure,google</a>" \
  -H "Authorization: token ${SNYK_API_TOKEN}" | jq -r '.data[0].id') && \
curl -X POST \
"https://api.snyk.io/rest/orgs/${SNYK_ORG_ID}/cloud/scans?version=2022-12-21~beta" \
-H "Authorization: token ${SNYK_API_TOKEN}" \
-H "Content-Type:application/vnd.api+json"  -d "{
  \"data\": {
    \"relationships\": {
      \"environment\": {
        \"data\": {
          \"id\": \"${SNYK_ENV_ID}\",
          \"type\": \"environment\"
        }
      }
    },
    \"type\": \"resource\"
  }
}"
```

`jq -r '.data[0].id`는 Snyk API [환경 목록](https://apidocs.snyk.io/#get-/orgs/-org\_id-/cloud/environments) 출력에서 **첫 번째** 환경의 ID를 반환하기 때문에, 이 기술은 특히 당신의 조직이 단일 환경을 가지고 있는 경우 유용합니다. 다른 환경을 스캔하려면 `0`을 다른 숫자로 변경할 수도 있습니다. 예를 들어, `jq -r '.data[1].id`는 출력에서 **두 번째** 환경의 ID를 반환합니다.

## `jq` 없이

만약 [jq가 설치되어](https://stedolan.github.io/jq/download/) 있지 않다면, Snyk API에 요청을 보내 모든 환경 ID를 반환하고 스캔하려는 환경의 ID를 찾은 다음 스캔을 수동으로 트리거하는 요청을 보낼 수 있습니다.

### 환경 ID 찾기

먼저, 스캔하려는 클라우드 환경의 ID를 찾으세요. 다음 형식으로 요청을 [해당](https://apidocs.snyk.io/#get-/orgs/-org\_id-/cloud/environments) 엔드포인트에 보냅니다:

```bash
curl -X GET \
  'https://api.snyk.io/rest/orgs/당신의-조직-ID/cloud/environments?version=2022-12-21~beta' \
  -H 'Authorization: token 당신의-API-토큰'
```

결과에서 `data.id` 속성을 찾으세요. 다음 간소화된 예제에서 ID는 `3b7ccff9-8900-4e54-0000-1234abcd1234`입니다:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "3b7ccff9-8900-4e54-0000-1234abcd1234",
    <length를 위해 일부 부분을 생략함>
  }
}
```

### 스캔 트리거

스캔을 수동으로 트리거하려면 다음 형식으로 [클라우드 스캔](https://apidocs.snyk.io/#post-/orgs/-org\_id-/cloud/scans) 엔드포인트로 요청을 보냅니다:

```bash
curl -X POST \
'https://api.snyk.io/rest/orgs/당신의-조직-ID/cloud/scans?version=2022-12-21~beta' \
-H 'Authorization: token 당신의-API-토큰' \
-H "Content-Type:application/vnd.api+json"  -d '{
  "data": {
    "relationships": {
      "environment": {
        "data": {
          "id": "당신의-환경-ID",
          "type": "environment"
        }
      }
    },
    "type": "resource"
  }
}'
```

참고: 이 예제에서는 [curl](https://curl.se/)을 사용했지만 [Postman](https://www.postman.com/)이나 [HTTPie](https://httpie.io/)와 같은 다른 API 클라이언트를 사용할 수도 있습니다.

## API 응답 이해

Snyk는 새로운 스캔에 대한 정보를 포함한 JSON 문서를 반환합니다. 예를 들어:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": {
    "id": "a7fa2167-58a8-4ac5-9999-0987dcba6543",
    "type": "scan",
    "attributes": {
      "created": "2022-08-07T04:59:58.639423469Z",
      "updated": null,
      "finished": null,
      "revision": 2,
      "kind": "user_initiated",
      "status": "queued"
    },
    "relationships": {
      "environment": {
        "data": {
          "id": "3b7ccff9-8900-4e54-0000-1234abcd1234",
          "type": "environment"
        },
        "links": {
          "related": "/orgs/d70c1768-5675-0000-1234-abcd1234abcd/cloud/environments?id=3b7ccff9-8900-4e54-0000-1234abcd1234&version=2022-12-21~beta"
        }
      },
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

다음은 API 응답에서 일부 주요 속성들입니다:
- `data.id`: 스캔 ID
- `data.attributes.status`: 스캔 상태

## 스캔 상태 확인

스캔 상태를 확인하려면 스캔 중인 환경의 세부 정보를 검색하세요. 다음 형식으로 [`/cloud/environments`](https://apidocs.snyk.io/#get-/orgs/-org\_id-/cloud/environments) 엔드포인트에 요청을 보냅니다:

```bash
curl -X GET \
  'https://api.snyk.io/rest/orgs/당신의-조직-ID/cloud/environments?id=당신의-환경-ID&version=2022-12-21~beta' \
  -H 'Authorization: token 당신의-API-토큰'
```

Snyk는 환경 세부정보를 담은 JSON 문서를 반환합니다. `data.attributes.status` 값을 찾아 스캔 상태를 확인하세요. 다음 간소화된 예제에서 상태는 `success`입니다:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": [
    {
      "id": "3b7ccff9-8900-4e54-0000-1234abcd1234",
      "type": "environment",
      "attributes": {
        "status": "success",
        <length를 위해 일부 부분을 생략함>
      }
    }
  ]
}
```

스캔 상태 값은 다음과 같습니다:
- `queued`: 스캔이 시작 예정
- `in_progress`: 스캔 진행 중
- `success`: 스캔 완료
- `error`: 스캔 에러; 잠시 기다렸다가 다시 스캔을 시도하세요

## 조직의 모든 스캔 보기

조직의 모든 스캔을 보려면 [스캔 목록](https://apidocs.snyk.io/#get-/orgs/-org\_id-/cloud/scans) 엔드포인트로 API 요청을 보냅니다:

```bash
curl -X GET \
'https://api.snyk.io/rest/orgs/당신의-조직-ID/cloud/scans?version=2022-12-21~beta' \
-H 'Authorization: token 당신의-API-토큰'
```

Snyk는 모든 스캔에 대한 정보를 담은 JSON 문서를 반환합니다. 예를 들어:

```json
{
  "jsonapi": {
    "version": "1.0"
  },
  "data": [
    {
      "id": "a7fa2167-58a8-4ac5-9999-0987dcba6543",
      "type": "scan",
      "attributes": {
        "created_at": "2022-08-04T22:14:47Z",
        "error": "",
        "finished_at": "2022-08-04T22:16:31Z",
        "kind": "user_initiated",
        "options": {
          "role_arn": "arn:aws:iam::123456789012:role/snyk-cloud-role"
        },
        "revision": 2,
        "status": "success",
        "updated_at": "2022-08-04T22:16:31Z"
      },
      "relationships": {
        "environment": {
          "data": {
            "id": "3b7ccff9-8900-4e54-0000-1234abcd1234",
            "type": "environment"
          },
          "links": {
            "related": "/orgs/d70c1768-5675-0000-1234-abcd1234abcd/cloud/environments?id=3b7ccff9-8900-4e54-0000-1234abcd1234&version=2022-12-21~beta"
          }
        },
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
    <length를 위해 일부 부분을 생략함>
  ]
}
```