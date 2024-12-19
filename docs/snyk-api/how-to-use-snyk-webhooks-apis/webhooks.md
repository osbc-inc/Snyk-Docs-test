# 웹훅 이벤트 및 페이로드

웹훅은 요청 본문 내에서 JSON으로 이벤트 페이로드를 JSON으로 포함한 `Content-Type`이 `application/json`으로 전달됩니다. 또한 다음 헤더를 보냅니다:

- `X-Snyk-Event` - 이벤트의 이름
- `X-Snyk-Transport-ID` - 이 전달을 식별하는 GUID
- `X-Snyk-Timestamp` - 이벤트가 발생한 ISO 8601 타임스탬프, 예: `2020-09-25T15:27:53Z`
- `X-Hub-Signature` - 요청 본문의 HMAC 16진수 다이제스트, 웹훅을 안전하게 유지하고 요청이 실제로 Snyk에서 온 것임을 보장합니다
- `User-Agent` - 요청의 원천 식별, 예: `Snyk-Webhooks/XXX`

서버가 페이로드를 수신토록 구성되면 요청이 구성한 엔드포인트로 보내진 모든 페이로드를 수신합니다. 보안상의 이유로 요청은 Snyk에서 오는 것으로 제한해야 합니다.

웹훅 이벤트를 수신토록 구성된 서버가 `X-Snyk-Event` 헤더를 확인해야 하며, 여러 이벤트 유형을 받을 수 있습니다.

## **핑 (ping)**

핑 이벤트는 새 웹훅이 생성된 후 발생하며, 핑 웹훅 API를 사용하여 수동으로 트리거될 수도 있습니다. 이는 Snyk로부터 웹훅이 데이터를 올바르게 받는지 테스트하는 데 유용합니다.

`ping` 이벤트는 다음 요청을 수행합니다:

```shell
POST /webhook-handler/snyk123 HTTP/1.1
Host: my.app.com
X-Snyk-Event: ping/v0
X-Snyk-Transport-ID: 998fe884-18a0-45db-8ae0-e379eea3bc0a
X-Snyk-Timestamp: 2020-09-25T15:27:53Z
X-Hub-Signature: sha256=7d38cdd689735b008b3c702edd92eea23791c5f6
User-Agent: Snyk-Webhooks/044aadd
Content-Type: application/json
{
  "webhookId": "d3cf26b3-2d77-497b-bce2-23b33cc15362"
}
```

## **프로젝트 스냅샷 (project_snapshot)**

이 이벤트는 기존 프로젝트가 테스트되고 새로운 스냅샷이 생성될 때마다 트리거됩니다. 프로젝트의 모든 테스트에 트리거되며, 새로운 문제가 있는지 여부와 관계없이 트리거됩니다. 이 이벤트는 새 프로젝트가 생성되거나 가져오는 경우에는 트리거되지 않습니다. 현재 지원되는 대상/스캔 유형은 오픈 소스 및 컨테이너입니다.

```sh
POST /webhook-handler/snyk123 HTTP/1.1
Host: my.app.com
X-Snyk-Event: project_snapshot/v0
X-Snyk-Transport-ID: 998fe884-18a0-45db-8ae0-e379eea3bc0a
X-Snyk-Timestamp: 2020-09-25T15:27:53Z
X-Hub-Signature: sha256=7d38cdd689735b008b3c702edd92eea23791c5f6
User-Agent: Snyk-Webhooks/044aadd
Content-Type: application/json
{
  "project": { ... }, // API 응답과 일치하는 프로젝트 객체
  "org": { ... }, // API 응답과 일치하는 조직 객체
  "group": { ... }, // API 응답과 일치하는 그룹 객체
  "newIssues": [], // API 응답과 일치하는 이슈 객체의 배열
  "removedIssues": [], // API 응답과 일치하는 이슈 객체의 배열
}
```

## **페이로드의 상세 예제**

### **프로젝트 (project)**

[프로젝트 (v1)](../reference/projects-v1.md) API 참조.

```json
"project": {
  "name": "snyk/goof",
  "id": "af137b96-6966-46c1-826b-2e79ac49bbd9",
  "created": "2018-10-29T09:50:54.014Z",
  ...
}
```

### **조직 (org)**

[조직 (v1)](../reference/organizations-v1.md) API 참조.

```json
"org": {
  "name": "My Org",
  ...
}
```

### **그룹 (group)**

[그룹 (v1)](../reference/groups-v1.md) API 참조.

```json
"group": {
  "name": "ACME Inc.",
   ...
}
```

### **이슈 (issue)**

```json
{
  "id": "npm:ms:20170412",
  ...
}
```