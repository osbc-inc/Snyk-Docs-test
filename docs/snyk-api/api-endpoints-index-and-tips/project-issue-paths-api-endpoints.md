# 프로젝트 이슈 경로 API Endpoints

다음은 [모든 프로젝트 이슈 경로 나열](../reference/projects-v1.md#org-orgid-project-projectid-issue-issueid-paths) 및 [모든 프로젝트 스냅샷 이슈 경로 나열](../reference/snapshots-v1.md#org-orgid-project-projectid-history-snapshotid-issue-issueid-paths)의 API 참조에 제공된 정보에 대한 추가 정보를 제공합니다.

`paths` 엔드포인트는 이슈가 도입된 경로의 세부 정보를 제공합니다.

`paths` 엔드포인트로의 **요청**은 `GET` 요청입니다. 이 엔드포인트는 다음 URL에서 사용할 수 있습니다:

- `https://api.snyk.io/v1/org/<orgId>/project/<projectId>/issue/<issueId>/paths` ([모든 프로젝트 이슈 경로 나열](../reference/projects-v1.md#org-orgid-project-projectid-issue-issueid-paths)). 이는 프로젝트의 최신 테스트에서 이슈에 대한 경로를 반환합니다.
- `https://api.snyk.io/v1/org/<orgId>/project/<projectId>/history/<snapshotId>/issue/<issueId>/paths` ([모든 프로젝트 스냅샷 이슈 경로 나열](../reference/snapshots-v1.md#org-orgid-project-projectid-history-snapshotid-issue-issueid-paths)). 이는 프로젝트의 특정 테스트에서 이슈에 대한 경로를 반환합니다.

두 경로 엔드포인트 모두 페이지네이션을 허용하는 쿼리 문자열을 사용할 수 있습니다. 예를 들어, `?page=2&perPage=500`와 같이 설정할 수 있습니다.

기본적으로, 100개의 결과가 포함된 첫 번째 페이지가 반환됩니다. 페이지 당 최대 1,000개의 결과를 요청할 수 있습니다.

**응답**은 다음 구조를 가지고 있습니다:

```json
{
    "snapshotId": "6d5813be-7e6d-4ab8-80c2-1e3e2a454553",
    "paths": [...],
    "total": 193,
    "links": {
        "prev": "<https://snyk.io/>...",
        "next": "<https://snyk.io/>...",
        "last": "<https://snyk.io/>..."
    }
}
```

- `snapshotId`는 경로가 반환된 프로젝트 스냅샷의 ID입니다.
- `total`은 스냅샷에서 해당 이슈에 대한 전체 경로 수입니다.
- `links`는 응답 페이지 간 탐색을 위한 편리한 링크를 제공합니다. `links.next` 및 `links.prev`는 해당 페이지가 존재할 경우에만 제공됩니다. 예를 들어, 결과의 마지막 페이지를 검색하면 `next` 링크가 포함되지 않습니다.
- `paths`는 배열로, 각 요소는 의존성 트리를 통한 경로입니다. 각 경로는 예를 들어 다음과 같은 패키지 설명자의 배열입니다:

```json
{
    "paths": [
        [
            { "name": "lodash", "version": "4.17.4", "fixVersion": "4.17.20" }
        ],
        [ 
            { "name": "babel-template", "version": "6.26.0", "fixVersion": "6.26.0" },
            { "name": "lodash", "version": "4.17.10" }
        ]
    ]
}
```

이 예에서 이슈는 `lodash` 패키지의 두 가지 다른 버전을 통해 도입됩니다. 하나는 직접 종속성이고, 다른 하나는 `babel-template`을 통해 간접적입니다.

가장 짧은 경로가 먼저 제공됩니다. 이슈가 프로젝트 자체에 적용되는 경우 해당 문제가 경로의 유일한 요소로 나타납니다. 종속성에 적용되는 문제의 경우 각 경로는 직접 종속성으로 시작합니다.

`fixVersion` 속성은 해당 경로가 업그레이드 가능할 때 첫 번째 요소에 제공됩니다. `version` 속성과 `fixVersion` 속성이 동일한 경우, 업그레이드는 전이적 종속성을 다시 잠근다는 것을 의미합니다.