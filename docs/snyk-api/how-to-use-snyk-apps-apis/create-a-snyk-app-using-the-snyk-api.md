# Snyk API를 사용하여 Snyk 앱 생성하기

API 토큰과 `orgId`가 있는 경우, [조직을 위한 새로운 Snyk 앱 생성](../reference/apps.md#orgs-org\_id-apps-creations) 엔드포인트로 `POST` 요청을 보내어 Snyk API를 사용하여 Snyk 앱을 생성할 수 있습니다:

```
https://api.snyk.io/rest/orgs/{orgId}/apps/creations?version={version}
```

Snyk 앱을 생성하는 예시 CURL 요청:

```
curl -XPOST -H"Content-Type: application/vnd.api+json" \
 -H"Authorization: token <REPLACE_WITH_API_TOKEN>" \
 -d '{"data": { "attributes": {"name": "My Awesome Snyk App", "redirect_uris": ["https://example.com/callback"], "scopes": ["org.read"], "context": "tenant"}, "type": "app"}}' \
 "https://api.snyk.io/rest/orgs/<REPLACE_WITH_YOUR_ORGID>/apps/creations?version=2024-10-11"
```

요청 본문에는 `name`, `redirectURIs`, 및 [`scopes`](scopes-to-request.md)와 같은 새로운 앱의 세부 정보가 포함되어야 합니다.

응답에는 통합을 완료하는 데 필요한 세부 정보인 `clientId` 및 `clientSecret`가 포함됩니다. 이러한 값을 Snyk API 엔드포인트에서 앱 내에서 사용하며, 앱 구성의 일부로 저장하는 것을 고려하십시오.

{% hint style="info" %}
`clientSecret`는 앱을 인증하는 데 사용되므로 절대 공개해서는 안됩니다. 이 `clientSecret`는 단 한 번만 표시되므로 안전하고 비공개로 유지하십시오. 만약 누출되거나 분실된 경우, [앱의 clientSecret 회전](manage-app-details.md#rotate-app-clientsecret)이 가능합니다.
{% endhint %}