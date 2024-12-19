# REST Search Audit Logs API와 GA REST Audit Logs API 마이그레이션 가이드

## REST Search Audit Logs API의 새로운 기능

OpenAPI 사양을 기반으로, Snyk REST API는 일관된, 친숙하며 사용하기 쉬운 API 프레임워크를 제공하도록 설계되었습니다. 새 API의 이점은 다음과 같습니다:

* 일관된 버전 관리
* 개선된 성능

Search Audit Logs REST API는 성능 및 일관적인 커서 기반 페이지네이션을 선호하며 오프셋 기반 페이지네이션을 사용하지 않도록 구체화했습니다. 우리의 v1 API는 페이지 오프셋 값이 증가할수록 느려지고 매우 높은 페이지 수에서 최종적으로 시간 초과됩니다. 우리는 이를 대폭 개선하여 페이지 수와 상관없이 더 빠른 속도로 결과를 지속적으로 반환하는 커서 기반 접근 방식을 사용하여 REST API에서 이를 크게 개선했습니다.

성능 개선뿐만 아니라, 더 세밀한 필터링을 위해 여러 개의 포함 또는 제외 이벤트를 제공할 수 있는 필터링 기능도 개선되었습니다. 이에 추가로, API는 일정 기간의 로그 검색만 가능했지만, 우리는 이를 하루의 분 단위 창으로 검색할 수 있도록 더 높은 세분성의 시간 스탬프 [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339)로 개선했습니다.

모든 검색 쿼리에 대해 이제 `` `api.access` `` 로그가 기본적으로 제외되며, 명시적으로 \`events\` 매개변수의 일부로 제공되지 않는 한 모든 검색 쿼리에서 제외됩니다. 이는 일반적으로 정보가 적고 고용량인 로그이며 명시적 조치 로그(e.g. group.create)로 대체되는데, 명시적 조치 로그에는 조치에 대한 보다 풍부한 문맥 정보가 포함되어 있으므로 모든 경우에 이를 선호해야 합니다.

다음은 v1 엔드포인트 및 해당 GA REST 같이 사용할 수 있는 내용입니다:

<table><thead><tr><th width="324">v1 그룹 및 조직 검색 감사 로그</th><th>GA REST API</th><th>참고</th></tr></thead><tbody><tr><td>(query) from (YYYY-MM-DD)</td><td>(query) from (<a href="https://datatracker.ietf.org/doc/html/rfc3339">RFC3339</a>)</td><td>시간 형식이 변경되었습니다</td></tr><tr><td>(query) to (YYYY-MM-DD)</td><td>(query) to (<a href="https://datatracker.ietf.org/doc/html/rfc3339">RFC3339</a>)</td><td>시간 형식이 변경되었습니다.</td></tr><tr><td>(query) page (Number)</td><td>(query) cursor (string)</td><td></td></tr><tr><td>(query) sortOrder</td><td>그대로 유지</td><td></td></tr><tr><td>(body) filters.userId (string)</td><td>(query) user_id (string)</td><td>요청 본문에서 쿼리 매개변수로 이동되었습니다.</td></tr><tr><td>(body) filters.email (string)</td><td>사용 중단됨</td><td>사용자 의견에 따라 제거되었습니다</td></tr><tr><td>(body) filters.event (string)</td><td>(query) events (쉼표로 구분된 목록)</td><td>매개변수 이름이 변경되고 쿼리로 이동되었습니다. 여러 값을 지원합니다.</td></tr><tr><td>(body) filters.excludeEvent (string)</td><td>(query) exclude_events (쉼표로 구분된 목록)</td><td>매개변수 이름이 변경되고 쿼리로 이동되었습니다. 여러 값을 지원합니다.</td></tr><tr><td>(body) filters.projectId (string)</td><td>(query) project_id (string)</td><td>요청 본문에서 쿼리 매개변수로 이동되었습니다.</td></tr></tbody></table>

## 응답

REST 마이그레이션의 일환으로 감사 로그 이벤트 페이로드는 변경되지 않았지만, 응답 구조는 표준화된 JSON API 응답에 맞게 다릅니다.

v1 응답은 배열 형태로 반환됩니다. 예시:

\`\`\`

`[`

  `{`

    `"groupId": "4a18d42f-0706-4ad0-b127-24078731fbea",`

    `"orgId": "4a18d42f-0706-4ad0-b127-24078731fbea",`

    `"userId": "4a18d42f-0706-4ad0-b127-24078731fbea",`

    `"projectId": null,`

    `"event": "group.edit",`

    `"content": {`

      `"before": {`

        `"name": "Group Previous Name"`

      `},`

      `"after": {`

        `"name": "Group Current Name"`

      `}`

    `},`

    `"created": "2017-04-11T21:00:00.000Z"`

  `}`

`]`

\`\`\`

REST 응답은 동일한 이벤트 페이로드 정보를 포함하지만 다음 형식으로 제공됩니다:

\`\`\`

`{`

  `"data": {`

    `"items": [`

      `{`

  `"groupId": "4a18d42f-0706-4ad0-b127-24078731fbea",`

  `"orgId": "4a18d42f-0706-4ad0-b127-24078731fbea",`

  `"userId": "4a18d42f-0706-4ad0-b127-24078731fbea",`

  `"projectId": null,`

  `"event": "group.edit",`

  `"content": {`

    `"before": {`

      `"name": "Group Previous Name"`

    `},`

    `"after": {`

      `"name": "Group Current Name"`

    `}`

  `},`

  `"created": "2017-04-11T21:00:00.000Z"`

`}`

\
\

  `],`

  `},`

  `"jsonapi": {`

    `"version": "1.0"`

  `},`

  `"links": {`

    `"next": "https://example.com/api/resource",`

  `}`

`}`

\`\`\`