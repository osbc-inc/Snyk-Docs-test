# Snyk 웹훅 생성

[웹훅 API 만들기](https://snyk.docs.apiary.io/#reference/webhooks/webhook-collection/create-a-webhook)를 사용하여 Snyk 웹훅을 생성하세요.

API를 사용하려면 Snyk 조직 ID, Snyk 인증 토큰 및 대상 웹훅 URL을 제공해야 합니다.

다음은 예시 요청입니다. 원하는 도구를 사용하여 요청을 보낼 수 있습니다.

```plaintext
POST https://api.snyk.io/v1/org/{SNYK-ORG-ID}/webhooks HTTP/2
Host: snyk.io
Authorization: token {SNYK-TOKEN}
Content-Type: application/json

{
    "url": "https://{URL}",
    "secret": "my-secret-string"
}
```

응답은 다음과 같습니다:

```json
{
  "id": "{SNYK-WEBHOOK-ID}",
  "url": "https://{URL}"
}
```

그런 다음 [웹훅 트리거하기 API](https://snyk.docs.apiary.io/#reference/webhooks/ping/ping-a-webhook)를 사용하여 Snyk 웹훅을 프로액티브하게 트리거하여 통합을 테스트할 수 있습니다:

```plaintext
POST https://api.snyk.io/v1/org/{SNYK-ORG-ID}/webhooks/{SNYK-WEBHOOK-ID}/ping HTTP/2
Host: snyk.io
Authorization: token {SNYK-TOKEN}
Content-Type: application/json
```

Azure Function과 Snyk 웹훅이 생성된 후 [New Relic Curated UI 및 Snyk 사용자 정의 대시보드](new-relic-curated-ui-and-snyk-custom-dashboard.md)를 사용할 수 있습니다.