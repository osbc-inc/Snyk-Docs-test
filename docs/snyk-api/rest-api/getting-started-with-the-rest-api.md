# REST API를 시작하는 방법

`curl`을 사용하여 명령 줄에서 REST API에 간단한 호출을 하려면 다음 단계를 따르세요.

1. [Snyk](https://snyk.io/)에 로그인합니다.
2. 계정에서 왼쪽 탐색을 사용하여 프로젝트를 나열할 수 있는 **조직**을 찾습니다.
3. **조직 설정**으로 이동한 다음 **일반** 설정 페이지에서 **조직 ID**를 찾아 값 복사합니다.
4. 개인 [일반 계정 설정](https://app.snyk.io/account/)로 이동하고 **API 토큰**을 복사합니다. 자세한 지침은 [API용 인증](authentication-for-api/authenticate-for-the-api.md)을 참조하세요.
5. 요청을 수행하기 위해 `curl` 명령을 사용합니다. **조직 ID**와 **API 토큰**을 각각 대체하여 `{orgId}` 및 API\_TOKEN을 대체합니다. `version` 매개변수의 경우, Snyk은 현재 날짜를 사용하는 것을 권장합니다. 다음 예제가 있습니다.

```sh
curl --request GET \
--url "https://api.snyk.io/rest/orgs/{orgId}/projects?version=2024-06-10" \
--header "Content-Type: application/vnd.api+json" \
--header "Authorization: token API_TOKEN"
```

{% hint style="info" %}
API를 호출할 때 사용할 API URL은 각기 다른 지역별로 다릅니다. 전체 목록은 [API URLs](about-the-rest-api.md#api-urls)를 참조하세요.

예를 들어, `SNYK-US-02` 지역 API URL은 다음과 같습니다.

* **API V1:** https://api.us.snyk.io/v1/&#x20;
* **REST API:** https://api.us.snyk.io/rest/&#x20;
{% endhint %}

`target-reference` 매개변수를 사용하는 경우 URL 인코딩해야 합니다.

문제 또는 질문이 있으면 [Snyk 지원팀](https://support.snyk.io)에 문의하십시오.