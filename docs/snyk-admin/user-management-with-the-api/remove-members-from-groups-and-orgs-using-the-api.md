# API를 사용하여 그룹 및 조직에서 멤버 제거

사용자 계정에서 그룹 및 조직에서 프로그래밍 방식으로 멤버를 제거하려면 다음 단계에 설명된대로 API를 사용할 수 있습니다. 이 API 호출은 서비스 계정을 제거하는 데 사용할 수 없습니다.

## 조직 멤버십 제거

### 단계 1: 조직 멤버 목록 가져오기

**요청**: `GET https://api.snyk.io/v1/org/{orgId}/members`

**엔드포인트**: [멤버 목록](../../snyk-api/reference/organizations-v1.md#org-orgid-members)

이 호출은 조직의 모든 비관리자 멤버의 배열을 반환합니다. 조직에서 제거해야 하는 각 사용자의 `id`를 저장합니다.

### 단계 2: 조직에서 멤버 제거

**요청**: `DELETE https://api.snyk.io/v1/org/{orgId}/members/update/{userId}`

**엔드포인트**: [조직에서 멤버 제거](../../snyk-api/reference/organizations-v1.md#org-orgid-members-userid-1)

각 사용자에 대해 이전에 저장한 사용자 id를 사용하여 그 사용자를 조직에서 제거하기 위해 엔드포인트를 호출합니다.

성공적인 요청의 응답은 `200 OK`입니다.

사용자가 제거되었는지 확인하기 위해 조직 멤버 페이지를 확인하세요.

{% hint style="info" %}
조직에서 사용자를 제거하면 조직이 그룹의 일부인 경우 사용자는 그룹 멤버로써 그룹에 계속 존재합니다. 사용자를 그룹에서 완전히 제거하려면 다음 섹션의 단계를 따르세요.
{% endhint %}

## 그룹 멤버십 제거

### 단계 1: 그룹 멤버 목록 가져오기

**요청**: `GET https://api.snyk.io/v1/group/groupId/members`

**엔드포인트**: [그룹의 모든 멤버 목록](../../snyk-api/reference/groups-v1.md#group-groupid-members)

이 호출은 그룹의 모든 멤버의 배열을 반환합니다. 그룹에서 제거해야 하는 각 사용자의 `id`를 저장합니다.

### 단계 2: 그룹에서 멤버 제거

**요청**: PATCH https://api.snyk.io/rest/groups/{group\_id}/users/{id}?version=2024-07-10\~beta

**엔드포인트**: [그룹에서 사용자 역할 업데이트](https://apidocs.snyk.io/?version=2024-09-04%7Ebeta&_gl=1*191l4f9*_gcl_aw*R0NMLjE3MjE0MDU5NzcuQ2p3S0NBanduZWkwQmhCLUVpd0FBMnh1QmlwWlhrR2JvVy16SGJLb0hGZDk4SU80TlprcGMtcjM4bk8yOXpFMXZFRUJVbHY1LWdnVm1Cb0NHY2dRQXZEX0J3RQ..*_ga*MTM5MDkzOTgyMC4xNzA0NzI3Nzk5*_ga_X9SH3KP7B4*MTcyMjI3NzI0OS40ODAuMS4xNzIyMjc5MjIxLjQ2LjAuMA..#patch-/groups/-group_id-/users/-id-) (Beta, 현재 버전 사용)

**본문:**

```postman_json
{
    "data": {
        "attributes": {
            "membership": null
        },
        "id": "<user-id>",
        "type": "user"
    }
}
```

각 사용자에 대해 이전에 저장한 사용자 id를 사용하여 해당 멤버를 그룹에서 제거하려면 엔드포인트를 호출합니다.

성공적인 요청의 응답은 `200 OK`입니다.

사용자가 제거되었는지 확인하기 위해 그룹 멤버 페이지를 확인합니다.

{% hint style="info" %}
그룹에서 사용자를 제거하면 사용자는 Snyk에 계속 존재합니다. 사용자와 관련된 모든 데이터를 완전히 삭제하려면 다음 섹션의 단계를 따르세요.
{% endhint %}

## 그룹 멤버 삭제

SSO 연결이 하나의 그룹에만 연결된 경우, 다음 호출을 사용하여 시스템에서 그룹 멤버를 완전히 삭제할 수 있습니다. 이 삭제 작업은 GDPR(일반 데이터 보호 규정) 요구 사항에도 부합합니다.

**요청**: DELETE `https://api.snyk.io/rest/groups/{group_id}/sso_connections/{sso_id}/users/{user_id}?version=2023-01-30~beta`

**엔드포인트**: [그룹 SSO 연결에서 사용자 삭제](https://apidocs.snyk.io/?version=2024-09-04%7Ebeta#delete-/groups/-group_id-/sso_connections/-sso_id-/users/-user_id-) (Beta, 현재 버전 사용)

{sso_id}는 Snyk 웹 UI에서 찾을 수 있습니다. **그룹** >**설정** >**SSO** >**단계 3**으로 이동하세요. 도움이 필요하면 계정 팀에 문의하세요.

<figure><img src="../../.gitbook/assets/Screenshot 2023-02-22 at 10.27.19.png" alt="Self Serve SSO 스크린, 단계 3, sso_id 강조"><figcaption><p>Self Serve SSO 스크린, 단계 3, sso_id 강조</p></figcaption></figure>

성공적인 요청의 응답은 `200 OK`입니다.

멤버가 삭제되었는지 확인하려면 다음 요청을 사용하세요.
`GET https://api.snyk.io/v1/user/userId`