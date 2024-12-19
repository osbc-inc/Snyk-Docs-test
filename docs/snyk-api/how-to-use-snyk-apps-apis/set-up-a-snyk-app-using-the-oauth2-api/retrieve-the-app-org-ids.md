# App Org IDs 가져오기

사용자는 단일 조직 또는 단일 그룹과 연결할 수 있습니다. 대부분의 Snyk API 엔드포인트는 수행 중인 작업을 인가하는 데 사용되는 경로의 `orgid`를 요구합니다. 자세한 내용은 [Snyk와 앱 통합](../#integrating-apps-with-snyk)을 참조하십시오.

당신의 앱에서 사용하는 `orgid`를 검색하려면, 다음 URL의 `orgs` 엔드포인트 [접근 가능한 조직 나열](https://apidocs.snyk.io/#get-/orgs)로 GET 요청을 보냅니다:

`https://api.snyk.io/rest/orgs?version={version}`

Snyk은 이 값을 저장하고 사용자 세부정보와 관련시키는 것을 권장합니다.