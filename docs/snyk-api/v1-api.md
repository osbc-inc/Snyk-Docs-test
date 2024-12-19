# V1 API

## V1 API 소개

{% hint 스타일 = "info" %}
Snyk API는 엔터프라이즈 플랜에서만 사용할 수 있습니다.

더 많은 정보는 [플랜 및 가격](https://snyk.io/plans)을 참조하세요.

V1 API는 이제 REST API에 중점을 두고 있기 때문에, 언젠가는 종료될 예정입니다.
{% endhint %}

V1 API를 사용하면 Snyk에서 정의한 문제에 대해 패키지를 테스트하고, 특정 워크플로우를 완수하기 위해 Snyk 프로세스를 자동화할 수 있습니다. 고객 및 파트너는 다음과 같은 작업을 수행할 수 있습니다:

* 취약점 데이터에 액세스
* 프로젝트 및 애플리케이션 스캔
* 보완 조언 수신
* 사용자 데이터 확인하여 사용자 정의 보안 솔루션 구축

V1 API 엔드포인트는 Snyk 사용자 문서의 [참조](reference/)에서 확인할 수 있습니다. 사용자 문서에서 업데이트가 이루어집니다. 사용자 문서로 이주된 엔드포인트는 [온라인](https://snyk.docs.apiary.io)에도 남아 있습니다.

## API URL

Snyk는 다음 지역에서 호스팅됩니다. 각 지역에는 고유한 기본 URL이 있습니다.

<table><thead><tr><th width="189">지역</th><th>기본 URL</th></tr></thead><tbody><tr><td>SNYK-US-01</td><td><code>https://api.snyk.io/v1/</code></td></tr><tr><td>SNYK-US-02</td><td><code>https://api.us.snyk.io/v1/</code></td></tr><tr><td>SNYK-EU-01 </td><td><code>https://api.eu.snyk.io/v1/</code> </td></tr><tr><td>SNYK-AU-01</td><td><code>https://api.au.snyk.io/v1/</code></td></tr></tbody></table>

{% hint 스타일 = "info" %}
이 API는 오직 HTTPS로만 사용 가능합니다. HTTP를 통해 호출하면 모든 요청에 대해 404를 반환합니다.
{% endhint %}

## 인가

이 API를 사용하려면 Snyk에서 토큰을 받아야 합니다. Snyk에 등록하고 로그인한 후에는 [개인 계정 설정](https://snyk.io/account/)에서 토큰을 찾을 수 있습니다. 자세한 내용은 [API를 위한 인증(Authentication for API)](rest-api/authentication-for-api/) 참조합니다.&#x20;

토큰을 `token` 다음에 토큰이 오는 `Authorization` 헤더에 제공하세요:

```
Authorization: token API_KEY
```

그렇지 않으면 401 "Unauthorized" 응답이 반환됩니다.

```
HTTP/1.1 401 Unauthorized

{
    "code": 401,
    "error": "Not authorised",
    "message": "Not authorised"
}
```

## 요청 속도 제한

Snyk는 고객에게 안정적인 경험을 제공하기 위해 V1 API 요청을 제한합니다.

V1 API는 **분당 2,000개의 요청**을 기본 속도 제한으로 가지고 있지만, 특정 엔드포인트에는 더 낮은 제한이 설정될 수 있습니다. 각 엔드포인트의 속도 제한을 보려면 참조 문서를 참조하세요.

제한을 초과하면 `429` 오류 응답을 받게 됩니다.

## 오류

V1 API는 오류 응답에 대해 표준 HTTP 오류 코드를 사용합니다.

```json
{
    "code": 404,
    "message": "Org 39db46b1-ad57-47e6-a87d-e34f6968030b was not found or you may not have the correct permissions to access the org.",
    "error": "Org 39db46b1-ad57-47e6-a87d-e34f6968030b was not found or you may not have the correct permissions to access the org."
}
```

에러 참조는 또한 서버 응답의 `x-error-reference` 헤더에서 제공됩니다.

예시 `500` 응답:

```sh
HTTP/1.1 500 Internal Server Error
x-error-reference: a45ec9c1-065b-4f7b-baf8-dbd1552ffc9f
Content-Type: application/json; charset=utf-8
Content-Length: 1848
Vary: Accept-Encoding
Date: Sun, 10 Sep 2017 06:48:40 GMT
```
