# API용으로 인증

Snyk API를 사용하려면 Snyk에서 API 토큰을 받아야 합니다. API 토큰은 Snyk에 등록하고 로그인한 후 개인 [일반 계정 설정](https://app.snyk.io/account)에서 찾을 수 있습니다. **키** 필드에서 **표시하려면 클릭**하세요. 그런 다음 API 키를 강조 표시하고 복사하세요.

새 API 토큰을 원하는 경우 **폐기 및 재생성**을 선택하세요. 이렇게 하면 이전 API 토큰이 무효화됩니다. 자세한 내용은 [Snyk API 토큰의 폐기 및 재생성](revoke-and-regenerate-a-snyk-api-token.md)을 참조하십시오.

API를 직접 사용할 때 다음 예시 요청에 API 토큰을 `Authorization` 헤더에 제공하며, `API_TOKEN`을 여러분의 API 토큰으로 대체하십시오.

```bash
curl --request GET \
--url "https://api.snyk.io/rest/self?version=2024-06-10" \
--header "Content-Type: application/vnd.api+json" \
--header "Authorization: token API_TOKEN"
```

API를 [Snyk Apps](../../how-to-use-snyk-apps-apis/)를 통해 사용하는 경우, 다음과 같이 `access_token`을 `bearer`로 선행하고 `Authorization` 헤더에 제공하세요.

```
Authorization: bearer ACCESS_TOKEN
```

그렇지 않으면 `401 Unauthorized` 응답이 반환됩니다.

```http
HTTP/1.1 401 Unauthorized

{
    "status": "401",
    "code": "Unauthorized"
}
```

API 토큰을 사용해야 하는 경우와 서비스 계정 토큰을 사용해야 하는 경우에 대한 정보는 [인증을 위한 API](./)를 참조하세요.