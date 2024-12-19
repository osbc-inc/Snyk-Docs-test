# 앱 세부 정보 관리

## 조직이 생성한 앱 목록

Snyk 조직이 소유한 Snyk 앱 목록을 보려면 `GET` 요청을 `apps/creations` 엔드포인트로 보냅니다:

`https://api.snyk.io/rest/orgs/{orgId}/apps/creations?version={version}`

자세한 내용은 엔드포인트 [조직이 생성한 앱 목록 가져오기](../reference/apps.md#orgs-org\_id-apps-1)을(를) 참조하세요.

## 앱 세부 정보 업데이트

앱의 이름이나 설정한 리디렉션 URI 목록을 업데이트할 수 있습니다.

앱을 업데이트하려면 `PATCH` 요청을 `apps/creations{app_id}` 엔드포인트로 보냅니다:

`https://api.snyk.io/rest/orgs/{orgId}/apps/creations{app_id}?version={version}`

`app_id` 경로 매개변수는 [`앱 생성 목록 조회의 'GET' 요청](manage-app-details.md#list-apps-created-by-an-organization)`에 대한 응답에서의 `id`입니다.

자세한 내용은 엔드포인트 [앱 ID를 사용하여 이름, 리디렉션 URL 및 엑세스 토큰 유효 기간과 같은 앱 생성 속성 수정](../reference/apps.md#orgs-org\_id-apps-creations-app\_id)을(를) 참조하세요.

## 앱 삭제

Snyk 조직에서 앱을 삭제하려면 `apps/creations{app_id}` 엔드포인트로 `DELETE` 요청을 보냅니다:

`https://api.snyk.io/rest/orgs/{orgId}/apps/creations/{app_id}?version={version}`

`app_id` 경로 매개변수는 [`앱 생성 목록 조회의 'GET' 요청](manage-app-details.md#list-apps-created-by-an-organization)`에 대한 응답에서의 `id`입니다.

자세한 내용은 엔드포인트 [앱 ID로 앱 삭제](../reference/apps.md#orgs-org\_id-apps-client\_id-2)을(를) 참조하세요.

앱을 삭제하면 앱 자격 증명이 취소되고 모든 앱 설치가 제거됩니다. 활성 사용자가 있는 경우 앱을 통해 Snyk에 연결할 수 없게 됩니다.

## 앱 clientSecret 회전

앱이 생성된 후 `clientSecret`를 볼 수는 없습니다. 분실한 경우 `clientSecret`를 회전하여 새로 받을 수 있습니다.

`POST` 요청을 보내어 생성한 앱에 대한 secret 관리 요청을 수행할 수 있습니다. 엔드포인트는 `apps/creations{app_id}/secrets`입니다:

`https://api.snyk.io/rest/orgs/{orgId}/apps/creations/{app_id}/secrets?version={version}`

`app_id` 경로 매개변수는 [`앱 생성 목록 조회의 'GET' 요청](manage-app-details.md#list-apps-created-by-an-organization)`에 대한 응답에서의 `id`입니다.

자세한 내용은 엔드포인트 [Snyk 앱의 client secret 관리](../reference/apps.md#orgs-org\_id-apps-creations-app\_id-secrets)을(를) 참조하세요.

{% hint style="info" %}
설치한 클라이언트 자격 증명 앱에 대해서는 [비대화식 Snyk 앱 설치에 대한 클라이언트 비밀 관리](../reference/apps.md#groups-group\_id-apps-installs-install\_id-secrets) 엔드포인트를 참조하세요.
{% endhint %}

POST 요청의 본문에 의해 표시되는 세 가지 작업을 수행할 수 있습니다:

* 생성 `{"mode": "create"}`
* 삭제 `{"mode": "delete", "secret": "{clientSecret}"}`
* 교체 `{"mode": "replace"}`

Snyk은 시크릿을 회전할 때 다음 절차를 채택하는 것을 권장합니다:

1. `{"mode": "create"}`를 사용하여 새로운 시크릿 생성
2. 새로 생성된 시크릿으로 서비스 업데이트
3. `{"mode": "delete", "secret": "{secret}"}`를 사용하여 이전 시크릿 삭제

### clientSecret 생성

보통 작동 중에는 정기적으로 클라이언트 시크릿을 회전하는 것이 좋습니다. 프로세스를 시작하려면 새로운 시크릿을 만들어줄 엔드포인트에 `{"mode": "create"}` 요청 본문을 보냅니다. 이 호출의 반환 값은 새로 생성된 시크릿이 포함된 앱입니다. 새 시크릿과 기존 시크릿은 수동으로 대체되거나 삭제될 때까지 유효합니다. 또한 즉시 클라이언트 시크릿을 교체할 수 있습니다.

한 앱에는 최대 두 개의 활성 시크릿이 있을 수 있습니다. 최대 수의 활성 시크릿이 이미 있는 경우 `create`를 호출하면 실패합니다.

### clientSecret 교체

앱의 `clientSecret`가 유출된 경우 `{"mode": "replace"}`를 사용하여 새로 생성할 수 있습니다.

`clientSecret`를 교체하면 현재 시크릿이 즉시 무효화됩니다. 앱이 새로운 시크릿으로 구성을 업데이트할 때까지 앱은 Snyk에 연결할 수 없습니다.

### clientSecret 삭제

사용하지 않는 시크릿을 정리하려면 `{"mode": "delete", "secret": "{clientSecret}"}`를 사용하여 엔드포인트를 호출합니다. 여기서 `{clientSecret}`은 삭제하려는 클라이언트 시크릿입니다. 이 작업은 즉시 시크릿을 무효화하여 더 이상 사용할 수 없게 합니다.

한 앱은 최소한 하나의 활성 시크릿이 있어야 합니다. 마지막 시크릿으로 delete를 호출하면 실패합니다.