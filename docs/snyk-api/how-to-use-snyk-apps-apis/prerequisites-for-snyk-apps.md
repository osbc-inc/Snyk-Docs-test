# Snyk 앱을 위한 선행 조건

Snyk 앱을 생성하려면 Snyk API에 액세스해야 합니다. 시작하려면 [API에 대한 인증](../rest-api/authentication-for-api/authenticate-for-the-api.md) 단계를 따르세요.

또한 앱이 소유될 Snyk 조직의 ID(귀하의 `orgId`)를 검색해야 합니다. Snyk 웹 UI의 조직 설정에서 조직 ID를 가져오거나 엔드포인트 [액세스 가능한 조직 목록](../reference/orgs.md#orgs)을 사용하여 다음을 확인할 수 있습니다. \
`https://api.snyk.io/rest/orgs`에 Snyk API 토큰을 포함한 Authorization 헤더로 요청하세요.

{% hint style="info" %}
Snyk 앱은 유료 액세스를 구매한 사용자인지 여부와 관계없이 API에 대한 일급 액세스 권한을 갖습니다. 이 기능을 활용하려면 앱이 앱 내에서 API를 호출하기 위해 더 이상 사용되지 않는 `https://snyk.io/api/`보다 https://api.snyk.io/ 도메인의 API 엔드포인트를 사용해야 합니다.
{% endhint %}