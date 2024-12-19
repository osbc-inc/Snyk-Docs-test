# Snyk API를 사용하여 서비스 계정 관리

[Snyk REST API](../../snyk-api/reference/serviceaccounts.md)를 사용하여 서비스 계정을 관리할 수 있습니다.

{% hint style="info" %}
모든 이러한 작업을 수행하려면 특정 권한이 필요합니다. 자세한 내용은 [서비스 계정-역할 선택](./#select-a-role)을 참조하십시오.
{% endhint %}

## 서비스 계정 속성

`id` - 서비스 계정의 ID입니다.

`name` - 서비스 계정의 사용하기 편한 이름입니다.

`auth_type` - 서비스 계정의 인증 전략입니다. 다음 옵션이 있습니다:

- `api_key` - 서비스 계정이 일반 Snyk API 키를 사용합니다.
- `oauth_client_secret` - 서비스 계정이 클라이언트 시크릿으로 가져온 [OAuth 2.0 액세스 토큰](./#service-accounts-using-oauth-2.0)을 사용합니다.
- `oauth_private_key_jwt` - 서비스 계정이 개인 키로 서명된 JWT로 가져온 [OAuth 2.0 액세스 토큰](./#service-accounts-using-oauth-2.0)을 사용합니다.

`role_id` - 서비스 계정의 역할로, 해당 계정이 가지는 권한을 정의합니다. 사용 가능한 역할은 [group-id-roles](../../snyk-api/reference/groups-v1.md#group-groupid-roles) 엔드포인트를 사용하여 찾을 수 있습니다.

`jwks_url` - 서명된 JWT 요청을 확인하는 데 사용되는 공개 키를 호스팅하는 JWKs URL입니다. 이는 `https`이어야 합니다. `auth_type`이 `oauth_private_key_jwt`인 경우에만 필요합니다.

`access_token_ttl_seconds` - 생성된 액세스 토큰이 유효한 시간(초)입니다. 설정되지 않으면 기본값은 1시간입니다. `auth_type`이 `oauth_client_secret` 또는 `oauth_private_key_jwt`인 경우에만 필요합니다.

## 그룹 수준 서비스 계정 관리

### 그룹 내 서비스 계정 목록 가져오기

**요청**: `GET https://api.snyk.io/rest/groups/{groupId}/service_accounts`

**API 엔드포인트:** [그룹 서비스 계정 목록 가져오기](../../snyk-api/reference/serviceaccounts.md#groups-group\_id-service\_accounts-1)

이 [페이지별](../../snyk-api/rest-api/about-the-rest-api.md#pagination) 호출은 각각 서비스 계정을 설명하는 객체 배열을 반환합니다.

### 그룹에 서비스 계정 생성

**요청**: `POST https://api.snyk.io/rest/groups/{groupId}/service_accounts`

**API 엔드포인트:** [그룹에 서비스 계정 생성](../../snyk-api/reference/serviceaccounts.md#groups-group\_id-service\_accounts)

이 호출은 새로운 서비스 계정을 생성합니다. 요청의 JSON 형식 본문에 `role_id`를 전달하여 서비스 계정이 사용할 수 있는 권한을 정의합니다. 이 역할 ID는 [group-id-roles](../../snyk-api/reference/groups-v1.md#group-groupid-roles) 엔드포인트를 사용하여 찾을 수 있습니다. 역할은 여러 서비스 계정에 재사용할 수 있습니다.

### 그룹에서 서비스 계정 가져오기

**요청**: `GET https://api.snyk.io/rest/groups/{groupId}/service_accounts/{serviceAccountId}`

**API 엔드포인트:** [그룹 서비스 계정 가져오기](../../snyk-api/reference/serviceaccounts.md#groups-group\_id-service\_accounts-serviceaccount\_id-1)

이 호출은 특정 서비스 계정을 설명하는 세부 정보를 반환합니다.

### 그룹의 서비스 계정 업데이트

**요청**: `PATCH https://api.snyk.io/rest/groups/{groupId}/service_accounts/{serviceAccountId}`

**API 엔드포인트:** [그룹 서비스 계정 업데이트](../../snyk-api/reference/serviceaccounts.md#groups-group\_id-service\_accounts-serviceaccount\_id)

이 호출은 특정 서비스 계정의 세부 정보(현재는 서비스 계정의 이름)를 업데이트합니다.

### 그룹에서 서비스 계정 삭제

**요청**: `DELETE https://api.snyk.io/rest/groups/{groupId}/service_accounts/{serviceAccountId}`

**API 엔드포인트:** [그룹 서비스 계정 삭제](../../snyk-api/reference/serviceaccounts.md#groups-group\_id-service\_accounts-serviceaccount\_id-secrets)

이 호출은 지정된 서비스 계정을 영구적으로 삭제하고 해당 자격 증명을 취소합니다.

### 그룹의 서비스 계정 클라이언트 시크릿 관리

**요청**: `POST https://api.snyk.io/rest/groups/{groupId}/service_accounts/{serviceAccountId}/secrets`

**API 엔드포인트:** [그룹 서비스 계정의 클라이언트 시크릿 관리](../../snyk-api/reference/serviceaccounts.md#groups-group\_id-service\_accounts-serviceaccount\_id-secrets)

이 호출을 사용하여 `oauth_client_secret` 서비스 계정의 클라이언트 시크릿을 관리할 수 있습니다. 다음 작업을 수행할 수 있습니다:

- `create` - 새로운 클라이언트 시크릿 생성. 서비스 계정당 최대 두 개의 활성 시크릿을 가질 수 있습니다.
- `delete` - 기존 클라이언트 시크릿 삭제. 이는 요청 본문에 `client_secret`를 넣어야 합니다. 기존 클라이언트 시크릿을 삭제하면 무효화됩니다. 서비스 계정은 적어도 하나의 활성 시크릿을 가져야 하며, 마지막 시크릿으로 삭제를 호출하면 실패합니다.
- `replace` - 기존 클라이언트 시크릿을 동시에 삭제하고 새 시크릿을 생성합니다. `client_secret`이 침해된 경우 이 옵션이 권장됩니다.

## 조직 수준 서비스 계정 관리

### 조직 내 서비스 계정 목록 가져오기

**요청**: `GET https://api.snyk.io/rest/orgs/{orgId}/service_accounts`

**API 엔드포인트:** [조직 서비스 계정 목록 가져오기](../../snyk-api/reference/serviceaccounts.md#orgs-org\_id-service\_accounts-1)

이 [페이지별](../../snyk-api/rest-api/about-the-rest-api.md#pagination) 호출은 각각 서비스 계정을 설명하는 객체 배열을 반환합니다.

### 조직에 서비스 계정 생성

**요청**: `POST https://api.snyk.io/rest/orgs/{orgId}/service_accounts`

**API 엔드포인트:** [조직에 서비스 계정 생성](../../snyk-api/reference/serviceaccounts.md#orgs-org\_id-service\_accounts)

이 호출은 새로운 서비스 계정을 생성합니다. 요청의 JSON 형식 본문에 `role_id`를 전달하여 서비스 계정이 사용할 수 있는 권한을 정의합니다. 이 `role id`는 [group-id-roles](../../snyk-api/reference/groups-v1.md#group-groupid-roles) 엔드포인트를 사용하여 찾을 수 있습니다. 역할은 여러 서비스 계정에 재사용할 수 있습니다.

### 조직에서 서비스 계정 가져오기

**요청**: `GET https://api.snyk.io/rest/orgs/{orgId}/service_accounts/{serviceAccountId}`

**API 엔드포인트:** [조직 서비스 계정 가져오기](../../snyk-api/reference/serviceaccounts.md#orgs-org\_id-service\_accounts-serviceaccount\_id-1)

이 호출은 특정 서비스 계정을 설명하는 세부 정보를 반환합니다.

### 조직의 서비스 계정 업데이트

**요청**: `PATCH https://api.snyk.io/rest/orgs/{orgId}/service_accounts/{serviceAccountId}`

**API 엔드포인트:** [조직 서비스 계정 업데이트](https://apidocs.snyk.io/?version=2023-09-07#patch-/orgs/-org\_id-/service\_accounts/-serviceaccount\_id-)

이 호출은 특정 서비스 계정의 세부 정보를 업데이트할 수 있습니다. 서비스 계정의 이름이 업데이트됩니다.

### 조직에서 서비스 계정 삭제

**요청**: `DELETE https://api.snyk.io/rest/orgs/{orgId}/service_accounts/{serviceAccountId}`

**API 엔드포인트:** [조직 내의 서비스 계정 삭제](../../snyk-api/reference/serviceaccounts.md#orgs-org\_id-service\_accounts-serviceaccount\_id-2)

이 호출은 지정된 서비스 계정을 영구적으로 삭제합니다.

### 조직의 서비스 계정 클라이언트 시크릿 관리

**요청**: `POST https://api.snyk.io/rest/orgs/{orgId}/service_accounts/{serviceAccountId}/secrets`

**API 엔드포인트:** [조직의 서비스 계정 클라이언트 시크릿 관리](../../snyk-api/reference/serviceaccounts.md#orgs-org\_id-service\_accounts-serviceaccount\_id-secrets)

이 호출을 사용하여 `oauth_client_secret` 서비스 계정의 클라이언트 시크릿을 관리할 수 있습니다. 다음 작업을 수행할 수 있습니다:

- `create` - 새로운 클라이언트 시크릿 생성. 서비스 계정당 최대 두 개의 활성 시크릿을 가질 수 있습니다.
- `delete` - 기존 클라이언트 시크릿 삭제. 이는 요청 본문에 `client_secret`를 넣어야 합니다. 기존 클라이언트 시크릿을 삭제하면 무효화됩니다. 서비스 계정은 적어도 하나의 활성 시크릿을 가져야 하며, 마지막 시크릿으로 삭제를 호출하면 실패합니다.
- `replace` - 기존 클라이언트 시크릿을 동시에 삭제하고 새 시크릿을 생성합니다. `client_secret`이 침해된 경우 이 옵션이 권장됩니다.