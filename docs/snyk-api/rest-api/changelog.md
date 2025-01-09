## 2024-10-15 - Updated 2024-12-09

### GET - `/orgs/{org_id}` - Updated
- 응답 속성 `data`가 상태 `200`에 대해 선택 사항으로 변경되었습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 속성 `jsonapi`가 상태 `200`에 대해 선택 사항으로 변경되었습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 속성 `links`가 상태 `200`에 대해 선택 사항으로 변경되었습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 상태 `200`의 응답에서 선택 사항이었던 속성 `links/first`를 제거했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 상태 `200`의 응답에서 선택 사항이었던 속성 `links/last`를 제거했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 상태 `200`의 응답에서 선택 사항이었던 속성 `links/next`를 제거했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 상태 `200`의 응답에서 선택 사항이었던 속성 `links/prev`를 제거했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 상태 `200`의 응답에서 선택 사항이었던 속성 `links/related`를 제거했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 상태 `409`로 비 성공 응답을 추가했습니다.

- 상태 `200`의 응답에 선택 사항 속성 `data/attributes/created_at`를 추가했습니다.

- 상태 `200`의 응답에 선택 사항 속성 `data/attributes/updated_at`를 추가했습니다.

- 상태 `200`의 응답 속성 `data/attributes`가 필수로 변경되었습니다.

- 상태 `200`에 대한 `data/type` 응답의 속성 패턴 `^[a-z][a-z0-9]*(_[a-z][a-z0-9]*)*$`가 추가되었습니다.

## 2024-10-15 - Updated 2024-11-28

### GET - `/orgs/{org_id}/projects/{project_id}/sbom` - Updated
- 새로운 선택 사항인 `query` 요청 매개변수 `exclude`를 추가했습니다.

## 2024-10-15 - Updated 2024-11-06

### GET - `/orgs/{org_id}/packages/{purl}/issues` - Updated
- 응답 상태 `200`에 대한 `data/items/attributes/coordinates/items/representations/items/` 응답 속성 `anyOf` 목록에서 `#/components/schemas/ResourcePathRepresentation`, `#/components/schemas/PackageRepresentation`를 제거했습니다.

### POST - `/orgs/{org_id}/packages/issues` - Updated
- 응답 상태 `200`에 대한 `data/items/attributes/coordinates/items/representations/items/` 응답 속성 `anyOf` 목록에서 `#/components/schemas/ResourcePathRepresentation`, `#/components/schemas/PackageRepresentation`를 제거했습니다.

## 2024-10-15 - Updated 2024-10-31

### GET - `/orgs/{org_id}/packages/{purl}/issues` - Updated
- 상태 `200`에 대한 `data/items/attributes/coordinates/items/representations/items/` 응답 속성 `anyOf` 목록에 `#/components/schemas/ResourcePathRepresentation`, `#/components/schemas/PackageRepresentation`를 추가했습니다.

### POST - `/orgs/{org_id}/packages/issues` - Updated
- 상태 `200`에 대한 `data/items/attributes/coordinates/items/representations/items/` 응답 속성 `anyOf` 목록에 `#/components/schemas/ResourcePathRepresentation`, `#/components/schemas/PackageRepresentation`를 추가했습니다.

## 2024-10-15 - Updated 2024-10-30

### GET - `/orgs/{org_id}/issues` - Updated
- 응답 상태 `200`에 대한 `data/items/attributes/coordinates/items/reachability` 응답 속성에 새로운 `function` enum 값을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 `data/items/attributes/coordinates/items/reachability` 응답 속성에 새로운 `no-info` enum 값을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 `data/items/attributes/coordinates/items/reachability` 응답 속성에 새로운 `not-applicable` enum 값을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 `data/items/attributes/coordinates/items/reachability` 응답 속성에 새로운 `package` enum 값을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)

### GET - `/orgs/{org_id}/issues/{issue_id}` - Updated
- 응답 상태 `200`에 대한 `data/attributes/coordinates/items/reachability` 응답 속성에 새로운 `function` enum 값을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 `data/attributes/coordinates/items/reachability` 응답 속성에 새로운 `no-info` enum 값을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 `data/attributes/coordinates/items/reachability` 응답 속성에 새로운 `not-applicable` enum 값을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 `data/attributes/coordinates/items/reachability` 응답 속성에 새로운 `package` enum 값을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)

### GET - `/groups/{group_id}/issues` - Updated
- 응답 상태 `200`에 대한 `data/items/attributes/coordinates/items/reachability` 응답 속성에 새로운 `function` enum 값을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 `data/items/attributes/coordinates/items/reachability` 응답 속성에 새로운 `no-info` enum 값을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 `data/items/attributes/coordinates/items/reachability` 응답 속성에 새로운 `not-applicable` enum 값을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 `data/items/attributes/coordinates/items/reachability` 응답 속성에 새로운 `package` enum 값을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)

### GET - `/groups/{group_id}/issues/{issue_id}` - Updated
- 응답 상태 `200`에 대한 `data/attributes/coordinates/items/reachability` 응답 속성에 새로운 `function` enum 값을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 `data/attributes/coordinates/items/reachability` 응답 속성에 새로운 `no-info` enum 값을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 `data/attributes/coordinates/items/reachability` 응답 속성에 새로운 `not-applicable` enum 값을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 `data/attributes/coordinates/items/reachability` 응답 속성에 새로운 `package` enum 값을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)

## 2024-10-15

### API 간소화된 버전 관리

앞으로 Snyk는 각 안정성에 대해 하나씩이 아닌 각 버전-날짜마다 하나의 API 명세를 노출할 것입니다. 새로운 버전의 Snyk API는 중단 변경으로 인해 필요할 때만 게시될 것입니다. 새로운 버전에 대해서는 베타 버전의 경우 날짜만을 지정해야 하며, `2024-10-15~beta`가 아닌 `2024-10-15`를 지정해야 합니다. 이 변경으로 인해 기존 버전에는 영향을 미치지 않습니다. 이 새로운 접근 방식은 다가오는 새로운 버전에만 적용됩니다.

## 2024-08-25 - Updated 2024-10-10

### GET - `/self` - Updated
- 응답 상태 `200`에 대한 `data/attributes` 응답 속성 `anyOf` 목록에 `#/components/schemas/User20240422`, `#/components/schemas/ServiceAccount20240422`를 추가했습니다.

- 응답 상태 `200`에 대한 `data/attributes` 응답 속성 `anyOf` 목록에서 `#/components/schemas/ServiceAccount`를 제거했습니다.

### GET - `/orgs/{org_id}/projects` - Updated
- 응답 상태 `200`에 대한 `data/items/relationships/target` 응답 속성 `oneOf` 목록에 `#/components/schemas/ProjectRelationshipsTarget20230215`를 추가했습니다. 

- 응답 상태 `200`에 대한 `data/items/relationships/target` 응답 속성 `oneOf` 목록에서 `#/components/schemas/ProjectRelationshipsTarget`를 제거했습니다.

### PATCH - `/orgs/{org_id}/projects/{project_id}` - Updated
- 응답 상태 `200`에 대한 `data/relationships/target` 응답 속성 `oneOf` 목록에 `#/components/schemas/ProjectRelationshipsTarget20230215`를 추가했습니다. 

- 응답 상태 `200`에 대한 `data/relationships/target` 응답 속성 `oneOf` 목록에서 `#/components/schemas/ProjectRelationshipsTarget`를 제거했습니다.

### GET - `/orgs/{org_id}/projects/{project_id}` - Updated
- 응답 상태 `200`에 대한 `data/relationships/target` 응답 속성 `oneOf` 목록에 `#/components/schemas/ProjectRelationshipsTarget20230215`를 추가했습니다.

- 응답 상태 `200`에 대한 `data/relationships/target` 응답 속성 `oneOf` 목록에서 `#/components/schemas/ProjectRelationshipsTarget`를 제거했습니다.

### GET - `/orgs/{org_id}/packages/{purl}/issues` - Updated
- 응답 상태 `200`에 대한 선택 사항이었던 속성 `data/items/attributes/coordinates/items/representation`을 제거했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 선택 사항이었던 속성 `data/items/attributes/key`을 제거했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 선택 사항이었던 속성 `data/items/attributes/slots/exploit`을 제거했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 선택 사항 속성 `data/items/attributes/severities/items/type`를 추가했습니다.

- 응답 상태 `200`에 대한 선택 사항 속성 `data/items/attributes/severities/items/version`를 추가했습니다.

- 응답 상태 `200`에 대한 선택 사항 속성 `data/items/attributes/slots/exploit_details`를 추가했습니다.

- 응답 상태 `200`에 대한 필수 속성 `data/items/attributes/coordinates/items/representations`를 추가했습니다.

### POST - `/orgs/{org_id}/packages/issues` - Updated
- 응답 상태 `200`에 대한 선택 사항이었던 속성 `data/items/attributes/slots/exploit`을 제거했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 선택 사항 속성 `data/items/attributes/severities/items/type`를 추가했습니다.

- 응답 상태 `200`에 대한 선택 사항 속성 `data/items/attributes/severities/items/version`를 추가했습니다.

- 응답 상태 `200`에 대한 선택 사항 속성 `data/items/attributes/slots/exploit_details`를 추가했습니다.

### GET - `/orgs/{org_id}/invites` - Updated
- 응답 상태 `200`에 대한 `data/items/attributes/role` 응답 속성의 type/format이 `string`/`uuid`에서 `string`/``으로 변경되었습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 `data/items/type` 응답 속성에서 `org_invitation` enum 값을 제거했습니다.

## 2024-08-25 - Updated 2024-09-11

### POST - `/orgs/{org_id}/apps` - Updated
- 새로운 필수 요청 속성 `name`을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 새로운 필수 요청 속성 `redirect_uris`을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 새로운 필수 요청 속성 `scopes`을 추가했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 요청 속성 `data`를 제거했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 새로운 선택 사항 요청 속성 `access_token_ttl_seconds`을 추가했습니다.

- 새로운 선택 사항 요청 속성 `context`를 추가했습니다.

### GET - `/orgs/{org_id}/apps` - Updated
- 응답 상태 `200`에 대한 `data/items/attributes/redirect_uris` 응답 속성의 minItems가 `1`에서 `0`으로 감소했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`에 대한 `data/items/attributes/client_id` 응답 속성이 필수로 변경되었습니다.

- 응답 상태 `200`에 대한 `data/items/attributes/redirect_uris` 응답 속성이 필수로 변경되었습니다.

### PATCH - `/orgs/{org_id}/apps/{client_id}` - Updated
- 응답 상태 `200`에 대한 `data/attributes/redirect_uris` 응답 속성의 minItems가 `1`에서 `0`으로 감소했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 요청 속성 `data`를 제거했습니다.
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 새로운 선택 사항 요청 속성 `access_token_ttl_seconds`을 추가했습니다.

- 새로운 선택 사항 요청 속성 `name`을 추가했습니다attribute`가 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/user/data/attributes/email` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/user/data/attributes/name` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/user/data/attributes/username` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/user/data/id` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.



### GET - `/groups/{group_id}/memberships` - 업데이트됨
- `data/items/relationships/group` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/group/data/attributes` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/group/data/attributes/name` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/group/data/id` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/role` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/role/data/attributes` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/role/data/attributes/name` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/role/data/id` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/user` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/user/data/attributes` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/user/data/attributes/email` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/user/data/attributes/name` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/user/data/attributes/username` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/user/data/id` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.


## 2024-08-25 - 업데이트됨 2024-08-30

### POST - `/orgs/{org_id}/memberships` - 업데이트됨
- `data/relationships/org` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/org/data/attributes` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/org/data/attributes/name` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/org/data/id` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/role` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/role/data/attributes` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/role/data/attributes/name` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/role/data/id` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/user` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/user/data/attributes` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/user/data/attributes/email` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/user/data/attributes/name` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/user/data/attributes/username` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.

- `data/relationships/user/data/id` 응답 속성이 `201` 상태에서 필수로 변경되었습니다.



### GET - `/orgs/{org_id}/memberships` - Updated
- `data/items/relationships/org` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/org/data/attributes` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/org/data/attributes/name` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/org/data/id` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/role` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/role/data/attributes` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/role/data/attributes/name` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/role/data/id` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/user` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/user/data/attributes` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/user/data/attributes/email` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/user/data/attributes/name` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/user/data/attributes/username` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.

- `data/items/relationships/user/data/id` 응답 속성이 `200` 상태에서 필수로 변경되었습니다.


## 2024-08-25

### POST - `/orgs/{org_id}/memberships` - 추가됨
- 역할을 가지고 있는 사용자를 위한 조직 멤버십 생성



### GET - `/orgs/{org_id}/memberships` - 추가됨
- 조직의 모든 멤버십 반환



### PATCH - `/orgs/{org_id}/memberships/{membership_id}` - 추가됨
- 역할을 가지고 있는 사용자를 위한 조직 멤버십 업데이트



### DELETE - `/orgs/{org_id}/memberships/{membership_id}` - 추가됨
- 그룹에서 사용자의 멤버십 제거



### GET - `/groups/{group_id}/org_memberships` - 추가됨
- 그룹 사용자의 조직 멤버십 목록 가져오기



### POST - `/groups/{group_id}/memberships` - 추가됨
- 역할을 가지고 있는 사용자를 위한 그룹 멤버십 생성



### GET - `/groups/{group_id}/memberships` - 추가됨
- 그룹의 모든 멤버십 반환



### PATCH - `/groups/{group_id}/memberships/{membership_id}` - 추가됨
- 그룹 멤버십에서 역할 업데이트



### DELETE - `/groups/{group_id}/memberships/{membership_id}` - 추가됨
- 그룹에서 멤버십 삭제

## 2024-08-22

### GET - `/orgs/{org_id}/projects/{project_id}/sbom` - 업데이트됨
- `200` 상태의 응답에서 필수 속성인 `bomFormat` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `200` 상태의 응답에서 필수 속성인 `components` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `200` 상태의 응답에서 필수 속성인 `dependencies` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `200` 상태의 응답에서 필수 속성인 `metadata` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `200` 상태의 응답에서 선택적 속성인 `components` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `query` 요청 매개변수 `format`에 새로운 enum 값 `cyclonedx1.5+json` 추가

- `query` 요청 매개변수 `format`에 새로운 enum 값 `cyclonedx1.5+xml` 추가

- `query` 요청 매개변수 `format`에 새로운 enum 값 `cyclonedx1.6+json` 추가

- `query` 요청 매개변수 `format`에 새로운 enum 값 `cyclonedx1.6+xml` 추가


## 2024-08-15

### GET - `/orgs/{org_id}/audit_logs/search` - 업데이트됨
- `query` 요청 매개변수 `size`에 기본값 `100.00` 추가
![Badge](https://img.shields.io/badge/Breaking-yellow)


### GET - `/groups/{group_id}/audit_logs/search` - 업데이트됨
- `query` 요청 매개변수 `size`에 기본값 `100.00` 추가
![Badge](https://img.shields.io/badge/Breaking-yellow)

## 2024-06-21 - 업데이트됨 2024-06-27

### POST - `/orgs/{org_id}/collections` - 업데이트됨
- `201` 상태의 응답에서 `data/attributes/name` 속성의 maxLength를 `255`에서 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `201` 상태의 응답에서 `data/attributes/name` 속성의 minLength를 `1`에서 `0`으로 감소
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `201` 상태의 응답에서 `data/attributes/name` 속성 패턴 `^([a-zA-Z0-9 _\-\/:.])+$` 제거



### GET - `/orgs/{org_id}/collections` - 업데이트됨
- `200` 상태의 응답에서 `data/items/attributes/name` 속성의 maxLength를 `255`에서 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `200` 상태의 응답에서 `data/items/attributes/name` 속성의 minLength를 `1`에서 `0`으로 감소
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `200` 상태의 응답에서 `data/items/attributes/name` 속성 패턴 `^([a-zA-Z0-9 _\-\/:.])+$` 제거



### PATCH - `/orgs/{org_id}/collections/{collection_id}` - 업데이트됨
- `200` 상태의 응답에서 `data/attributes/name` 속성의 maxLength를 `255`에서 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `200` 상태의 응답에서 `data/attributes/name` 속성의 minLength를 `1`에서 `0`으로 감소
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `200` 상태의 응답에서 `data/attributes/name` 속성 패턴 `^([a-zA-Z0-9 _\-\/:.])+$` 제거



### GET - `/orgs/{org_id}/collections/{collection_id}` - 업데이트됨
- `200` 상태의 응답에서 `data/attributes/name` 속성의 maxLength를 `255`에서 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `200` 상태의 응답에서 `data/attributes/name` 속성의 minLength를 `1`에서 `0`으로 감소
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `200` 상태의 응답에서 `data/attributes/name` 속성 패턴 `^([a-zA-Z0-9 _\-\/:.])+$` 제거


## 2024-06-21 - 업데이트됨 2024-06-25

### PATCH - `/orgs/{org_id}` - 업데이트됨
- 요청 속성 `data/type`이 enum 값 목록으로 제한됨
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 요청 속성 `data/attributes`가 필수로 변경됨
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 요청 속성 `data/id`가 필수로 변경됨
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 요청 속성 `data/type`이 필수로 변경됨
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 속성 `data/type`에 새로운 `org` enum 값을 추가하고 응답 상태 `200`에 대해
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 요청 속성 `data/type`에 새로운 `org` enum 값을 추가함

- 요청 속성 `data/type`의 패턴 `^[a-z][a-z0-9]*(_[a-z][a-z0-9]*)*$` 삭제

- 응답 속성 `data/type`의 패턴 `^[a-z][a-z0-9]*(_[a-z][a-z0-9]*)*$`가 응답 상태 `200`에 대해 삭제됨


## 2024-06-21

### POST - `/orgs/{org_id}/invites` - 업데이트됨
- 요청 속성 `data/relationships` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)

## 2024-06-18

### POST - `/groups/{group_id}/settings/pull_request_template` - 업데이트됨
- 요청 속성 `data/attributes/branch_name` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `201` 상태의 응답에서 선택적인 속성 `data/attributes/branch_name` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)


### GET - `/groups/{group_id}/settings/pull_request_template` - 업데이트됨
- `200` 상태의 응답에서 선택적인 속성 `data/attributes/branch_name` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)

## 2024-06-06

### GET - `/orgs/{org_id}/projects` - 업데이트됨
- `200` 상태의 응답에서 선택적인 속성 `data/items/attributes/settings/auto_dependency_upgrade/is_inherited` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)


### PATCH - `/orgs/{org_id}/projects/{project_id}` - 업데이트됨
- `200` 상태의 응답에서 선택적인 속성 `data/attributes/settings/auto_dependency_upgrade/is_inherited` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)


### GET - `/orgs/{org_id}/projects/{project_id}` - 업데이트됨
- `200` 상태의 응답에서 선택적인 속성 `data/attributes/settings/auto_dependency_upgrade/is_inherited` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)

## 2024-05-23

### DELETE - `/self/apps/installs/{install_id}` - 업데이트됨
- API 작업 ID `deleteUserAppInstallByID` 제거 및 `deleteUserAppInstallById`로 대체
  

### DELETE - `/orgs/{org_id}/apps/installs/{install_id}` - 업데이트됨
- API 작업 ID `deleteAppOrgInstallByID` 제거 및 `- `/orgs/{org_id}/audit_logs/search` - 업데이트됨
  - `query` 요청 매개변수 `event` 삭제
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
  - `query` 요청 매개변수 `exclude_event` 삭제
  ![Badge](https://img.shields.io/badge/Breaking-yellow)
  - 새로운 선택적`query` 요청 매개변수 `events` 추가
  - 새로운 선택적 `query` 요청 매개변수 `exclude_events` 추가

### GET - `/groups/{group_id}/audit_logs/search` - 업데이트됨
- `query` 요청 매개변수 `event` 삭제
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `query` 요청 매개변수 `exclude_event` 삭제
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 새로운 선택적`query` 요청 매개변수 `events` 추가
- 새로운 선택적`query` 요청 매개변수 `exclude_events` 추가

## 2024-04-22

### GET - `/self` - 추가됨
- 요청을 실행하는 사용자에 관한 정보를 검색합니다.

### GET - `/orgs/{org_id}/projects` - 업데이트됨
- 응답 상태가 `200`인 경우, `data/items/relationships/target` 응답 프로퍼티 `oneOf` 목록에 `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` 추가

### PATCH - `/orgs/{org_id}/projects/{project_id}` - 업데이트됨
- 응답 상태가 `200`인 경우, `data/relationships/target` 응답 프로퍼티 `oneOf` 목록에 `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` 추가

### GET - `/orgs/{org_id}/projects/{project_id}` - 업데이트됨
- 응답 상태가 `200`인 경우, `data/relationships/target` 응답 프로퍼티 `oneOf` 목록에 `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` 추가

## 2024-02-28

### GET - `/orgs` - 업데이트됨
- `query` 요청 매개변수 `name`에 대해 maxLength를 `100`으로 설정
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `query` 요청 매개변수 `slug`에 대해 maxLength를 `100`으로 설정
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `query` 요청 매개변수 `slug`에 pattern `^[\w.-]+$` 추가
![Badge](https://img.shields.io/badge/Breaking-yellow)
- `200` 상태로 응답할 때 선택적 속성인 `data/items/attributes/access_requests_enabled` 추가

### PATCH - `/orgs/{org_id}` - 업데이트됨
- `200` 상태로 응답할 때 선택적 속성인 `data/attributes/access_requests_enabled` 추가

### GET - `/orgs/{org_id}` - 업데이트됨
- `200` 상태로 응답할 때 선택적 속성인 `data/attributes/access_requests_enabled` 추가

### GET - `/groups/{group_id}/orgs` - 추가됨
- 그룹에 속한 모든 조직의 목록을 페이지별로 가져옵니다.
기본적으로 이 엔드포인트는 조직 이름의 알파벳순으로 조직을 반환합니다.

## 2024-02-21

### GET - `/orgs/{org_id}/targets` - 추가됨
- 조직의 대상 목록을 가져옵니다.

### GET - `/orgs/{org_id}/targets/{target_id}` - 추가됨
- 조직의 지정된 대상을 가져옵니다.

### DELETE - `/orgs/{org_id}/targets/{target_id}` - 추가됨
- 지정된 대상을 삭제합니다.

### GET - `/orgs/{org_id}/projects` - 업데이트됨
- 응답 상태가 `200`인 경우, `data/items/relationships/target` 응답 프로퍼티 `oneOf` 목록에서 `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` 제거

### PATCH - `/orgs/{org_id}/projects/{project_id}` - 업데이트됨
- 응답 상태가 `200`인 경우, `data/relationships/target` 응답 프로퍼티 `oneOf` 목록에서 `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` 제거

### GET - `/orgs/{org_id}/projects/{project_id}` - 업데이트됨
- 응답 상태가 `200`인 경우, `data/relationships/target` 응답 프로퍼티 `oneOf` 목록에서 `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` 제거

## 2024-01-23

### GET - `/orgs/{org_id}/issues` - 추가됨
- 조직의 이슈 목록을 가져옵니다.

### GET - `/orgs/{org_id}/issues/{issue_id}` - 추가됨
- 이슈를 가져옵니다.

### GET - `/groups/{group_id}/issues` - 추가됨
- 그룹의 이슈 목록을 가져옵니다.

### GET - `/groups/{group_id}/issues/{issue_id}` - 추가됨
- 이슈를 가져옵니다.

## 2024-01-04

### POST - `/custom_base_images` - 업데이트됨
- `data/attributes/versioning_schema` 요청 프로퍼티 `oneOf` 목록에서 `#/components/schemas/VersioningSchemaDateType` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태가 `201`인 경우, `data/attributes/versioning_schema` 응답 프로퍼티 `oneOf` 목록에서 `#/components/schemas/VersioningSchemaDateType` 제거

### PATCH - `/custom_base_images/{custombaseimage_id}` - 업데이트됨
- `data/attributes/versioning_schema` 요청 프로퍼티 `oneOf` 목록에서 `#/components/schemas/VersioningSchemaDateType` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태가 `200`인 경우, `data/attributes/versioning_schema` 응답 프로퍼티 `oneOf` 목록에서 `#/components/schemas/VersioningSchemaDateType` 제거

### GET - `/custom_base_images/{custombaseimage_id}` - 업데이트됨
- 응답 상태가 `200`인 경우, `data/attributes/versioning_schema` 응답 프로퍼티 `oneOf` 목록에서 `#/components/schemas/VersioningSchemaDateType` 제거

## 2023-11-06

### DELETE - `/orgs/{org_id}/projects/{project_id}` - 추가됨
- 프로젝트 ID로 조직에서 하나의 프로젝트를 삭제합니다.

## 2023-11-03

### GET - `/self/apps/{app_id}/sessions` - 추가됨
- 앱에 대한 활성 OAuth 세션 목록을 가져옵니다.

### DELETE - `/self/apps/{app_id}/sessions/{session_id}` - 추가됨
- 활성 사용자 앱 세션을 폐기합니다.

### GET - `/self/apps/installs` - 추가됨
- 사용자가 설치한 앱 목록을 가져옵니다.

### DELETE - `/self/apps/installs/{install_id}` - 추가됨
- 설치 ID로 앱 액세스 권한을 폐기합니다.

### POST - `/orgs/{org_id}/apps` - 업데이트됨
- 새로운 필수 요청 프로퍼티 `data` 추가
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 요청 프로퍼티 `access_token_ttl_seconds` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 요청 프로퍼티 `context` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 요청 프로퍼티 `name` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 요청 프로퍼티 `redirect_uris` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 요청 프로퍼티 `scopes` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)

### GET - `/orgs/{org_id}/apps` - 업데이트됨
- 상태가 `200`인 경우, 응답 프로퍼티 `data/items/attributes/client_id`가 선택적 요소가 됨
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 상태가 `200`인 경우, 응답 프로퍼티 `data/items/attributes/redirect_uris`가 선택적 요소가 됨
![Badge](https://img.shields.io/badge/Breaking-yellow)

### PATCH - `/orgs/{org_id}/apps/{client_id}` - 업데이트됨
- 새로운 필수 요청 프로퍼티 `data` 추가
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 상태가 `200`인 경우, 응답 프로퍼티 `data/attributes/client_id`가 선택적 요소가 됨
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 상태가 `200`인 경우, 응답 프로퍼티 `data/attributes/redirect_uris`가 선택적 요소가 됨
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 요청 프로퍼티 `access_token_ttl_seconds` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 요청 프로퍼티 `name` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 요청 프로퍼티 `redirect_uris` 제거
![Badge](https://img.shields.io/badge/Breaking-yellow)

### GET - `/orgs/{org_id}/apps/{client_id}` - 업데이트됨
- 상태가 `200`인 경우, 응답 프로퍼티 `data/attributes/client_id`가 선택적 요소가 됨
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 상태가 `200`인 경우, 응답 프로퍼티 `data/attributes/redirect_uris`가 선택적 요소가 됨
![Badge](https://img.shields.io/badge/Breaking-yellow)

### POST - `/orgs/{org_id}/apps/installs` - 추가됨
- 조직에 Snyk 앱을 설치하고, 해당 앱은 비대화식 인증(예: 클라이언트 자격 증명)을 사용해야 합니다.

### GET - `/orgs/{org_id}/apps/installs` - 추가됨
- 조직에 설치된 앱 목록을 가져옵니다.

### DELETE - `/orgs/{org_id}/apps/installs/{install_id}` - 추가됨
- 설치 ID로 Snyk 조직에 대한 앱 권한을 폐기합니다.

### POST - `/orgs/{org_id}/apps/installs/{install_id}/secrets` - 추가됨
- 대화식이 아닌 Snyk 앱 인스톨에 대한 클라이언트 시크릿을 관리합니다.

### POST - `/orgs/{org_id}/apps/creations` - 추가됨
- 조직을 위해 새로운 Snyk 앱을 생성합니다.

### GET - `/orgs/{org_id}/apps/creations` - 추가됨
- 조직이 생성한 앱 목록을 가져옵니다.

### PATCH - `/orgs/{org_id}/apps/creations/{app_id}` - 추가됨
- 앱 ID를 사용하여 앱 생성 속성을 업데이트합니다.

### GET - `/orgs/{org_id}/apps/creations/{app_id}` - 추가됨
- 앱 ID에 해당하는 Snyk 앱을 가져옵니다.

### DELETE - `/orgs/{org_id}/apps/creations/{app_id}` - 추가됨
- 앱 ID에 해당하는 앱을 삭제합니다.

### POST - `/orgs/{org_id}/apps/creations/{app_id}/secrets` - 추가됨
- Snyk 앱의 클라이언트 시크릿을 관리합니다.

### POST - `/groups/{group_id}/apps/installs` - 추가됨
- 그룹에 Snyk 앱을 설치하고, 해당 앱은 비대화식 인증(예: 클라이언트 자격 증명)을 사용해야 합니다.

### GET - `/groups/{group_id}/apps/installs` - 추가됨
- 그룹에 설치된 앱 목록을 가져옵니다.

### DELETE - `/groups/{group_id}/apps/installs/{install_id}` - 추가됨
- 설치 ID로 Snyk 그룹에 대한 앱 권한을 폐기합니다.

### POST - `/groups/{group_id}/apps/installs/{install_id}/secrets` - 추가됨
- 대화식이 아닌 Snyk 앱 인스톨에 대한 클라이언트 시크릿을 관리합니다.

## 2023-11-02

### GET - `/orgs/{org_id}/container_images` - 추가됨
- 컨테이너 이미지의 인스턴스 목록을 나열합니다.

### GET - `/orgs/{org_id}/container_images/{image_id}` - 추가됨
- 컨테이너 이미지의 인스턴스를 가져옵니다.

### GET - `/orgs/{org_id}/container_images/{image_id}/relationships/image_target_refs` - 추가됨
- 컨테이너 이미지의 이미지 대상 참조 인스턴스 목록을 가져옵니다.

## 2023-09-13

### GET - `/orgs/{org_id}/projects` - 업데이트됨
- 응답 상태가 `200`인 경우, `data/items/relationships/target` 응답 프로퍼티 `oneOf` 목록에 `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` 추가

### PATCH - `/orgs/{org_id}/projects/{project_id}` - 업데이트됨
- 응답 상태가 `200`인 경우, `data/relationships/target` 응답 프로퍼티 `oneOf` 목록에 `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` 추가

### GET - `/orgs/{org_id}/projects/{project_id}` - 업데이트됨
- 응답 상태가 `200`인 경우, `data/relationships/target` 응답 프로퍼티 `oneOf` 목록에 `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` 추가

## 2023-09-12

### GET - `/orgs/{org_id}/projects` - 업데이트됨
- 응답 상태가 `200`인 경우, `data/items/relationships/target` 응답 프로퍼티 `oneOf` 목록에서 `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` 제거

### PATCH - `/orgs/{org_id}/projects/{project_id}` - 업데이트됨
- 응답 상태가 `200`인 경우, `data/relationships/target` 응답 프로퍼티 `oneOf` 목록에서 `#/components/schemas/Relationship, #/components/schemas/ProjectRelationshipsTarget` 제거

### GET - `/orgs/{- org.org_source.create
- org.org_source.delete
- org.org_source.edit
- org.policy.edit
- org.project_filter.create
- org.project_filter.delete
- org.project.add
- org.project.attributes.edit
- org.project.delete
- org.project.edit
- org.project.fix_pr.auto_open
- org.project.fix_pr.manual_open
- org.project.ignore.create
- org.project.ignore.delete
- org.project.ignore.edit
- org.project.monitor
- org.project.pr_check.edit
- org.project.remove
- org.project.settings.delete
- org.project.settings.edit
- org.project.stop_monitor
- org.project.tag.add
- org.project.tag.remove
- org.project.test
- org.request_access_settings.edit
- org.sast_settings.edit
- org.service_account.create
- org.service_account.delete
- org.service_account.edit
- org.settings.feature_flag.edit
- org.target.create
- org.target.delete
- org.user.add
- org.user.invite
- org.user.invite.accept
- org.user.invite.revoke
- org.user.invite_link.accept
- org.user.invite_link.create
- org.user.invite_link.revoke
- org.user.leave
- org.user.provision.accept
- org.user.provision.create
- org.user.provision.delete
- org.user.remove
- org.user.role.create
- org.user.role.delete
- org.user.role.details.edit
- org.user.role.edit
- org.user.role.permissions.edit
- org.webhook.add
- org.webhook.delete
- user.org.notification_settings.edit

### GET - `/groups/{group_id}/audit_logs/search` - 추가됨
- 그룹에 대한 감사 로그 검색. 일부 조직 수준 이벤트와 다음과 같은 그룹 수준 이벤트를 지원함:
  - api.access
  - group.cloud_config.settings.edit
  - group.create
  - group.delete
  - group.edit
  - group.notification_settings.edit
  - group.org.add
  - group.org.remove
  - group.policy.create
  - group.policy.delete
  - group.policy.edit
  - group.request_access_settings.edit
  - group.role.create
  - group.role.delete
  - group.role.edit
  - group.service_account.create
  - group.service_account.delete
  - group.service_account.edit
  - group.settings.edit
  - group.settings.feature_flag.edit
  - group.sso.add
  - group.sso.auth0_connection.create
  - group.sso.auth0_connection.edit
  - group.sso.create
  - group.sso.delete
  - group.sso.edit
  - group.sso.membership.sync
  - group.sso.remove
  - group.tag.create
  - group.tag.delete
  - group.user.add
  - group.user.remove
  - group.user.role.edit

## 2023-09-07

### POST - `/orgs/{org_id}/service_accounts` - 추가됨
- 조직을 위한 서비스 계정을 생성합니다. 서비스 계정은 Snyk API에 액세스하는 데 사용될 수 있습니다.

### GET - `/orgs/{org_id}/service_accounts` - 추가됨
- 조직을 위한 모든 서비스 계정을 가져옵니다.

### PATCH - `/orgs/{org_id}/service_accounts/{serviceaccount_id}` - 추가됨
- ID별로 조직 수준의 서비스 계정 이름을 업데이트합니다.

### GET - `/orgs/{org_id}/service_accounts/{serviceaccount_id}` - 추가됨
- ID별로 조직 수준의 서비스 계정을 가져옵니다.

### DELETE - `/orgs/{org_id}/service_accounts/{serviceaccount_id}` - 추가됨
- 조직에서 서비스 계정을 삭제합니다.

### POST - `/orgs/{org_id}/service_accounts/{serviceaccount_id}/secrets` - 추가됨
- 서비스 계정의 클라이언트 시크릿을 관리합니다.

### POST - `/groups/{group_id}/service_accounts` - 추가됨
- 그룹을 위한 서비스 계정을 생성합니다. 서비스 계정은 Snyk API에 액세스하는 데 사용될 수 있습니다.

### GET - `/groups/{group_id}/service_accounts` - 추가됨
- 그룹을 위한 모든 서비스 계정을 가져옵니다.

### PATCH - `/groups/{group_id}/service_accounts/{serviceaccount_id}` - 추가됨
- ID별로 그룹의 서비스 계정 이름을 업데이트합니다.

### GET - `/groups/{group_id}/service_accounts/{serviceaccount_id}` - 추가됨
- ID별로 그룹 수준의 서비스 계정을 가져옵니다.

### DELETE - `/groups/{group_id}/service_accounts/{serviceaccount_id}` - 추가됨
- ID에 따라 그룹 수준의 서비스 계정을 영구적으로 삭제합니다.

### POST - `/groups/{group_id}/service_accounts/{serviceaccount_id}/secrets` - 추가됨
- 서비스 계정 ID로 그룹 서비스 계정의 클라이언트 시크릿을 관리합니다.

## 2023-08-28

### GET - `/orgs/{org_id}/projects` - Updated
- `names_start_with`라는 새로운 선택적 `query` 요청 매개변수를 추가함

- `target_file`라는 새로운 선택적 `query` 요청 매개변수를 추가함

- `target_reference`라는 새로운 선택적 `query` 요청 매개변수를 추가함

- `target_runtime`라는 새로운 선택적 `query` 요청 매개변수를 추가함

### PATCH - `/orgs/{org_id}/projects/{project_id}` - Updated
- `user_id`라는 `query` 요청 매개변수를 삭제함
![Badge](https://img.shields.io/badge/Breaking-yellow)

## 2023-08-21

### POST - `/orgs/{org_id}/packages/issues` - Updated
- 응답에 선택적 속성 `meta`를 `200` 상태로 추가함

### POST - `/custom_base_images` - 추가됨
- 사용자 정의 기본 이미지를 만들려면 먼저 기본 이미지를 Snyk로 가져와야 합니다.
CLI, UI 또는 API를 통해 수행할 수 있습니다.

이 엔드포인트는 이미지를 사용자 정의 기본 이미지로 표시합니다. 즉, 이미지는 Snyk가 기본 이미지 업그레이드를 권장할 수 있는 이미지 풀에 추가됩니다.

첫 번째 이미지가 리포지토리에 추가되면 이 엔드포인트에 버전 스키마를 전달할 수 없습니다.
버전 스키마를 업데이트하려면 PATCH 엔드포인트를 사용해야 합니다.

### GET - `/custom_base_images` - 추가됨
- 정렬 및 필터링을 지원하는 사용자 정의 기본 이미지 목록을 가져옵니다.
org_id 또는 group_id 매개변수 중 하나는 성공적인 인가를 위해 설정되어야 합니다.

### PATCH - `/custom_base_images/{custombaseimage_id}` - 추가됨
- 사용자 정의 기본 이미지의 속성을 업데이트함

### GET - `/custom_base_images/{custombaseimage_id}` - 추가됨
- 사용자 정의 기본 이미지를 가져옵니다.

### DELETE - `/custom_base_images/{custombaseimage_id}` - 추가됨
- 사용자 정의 기본 이미지 리소스를 삭제함 (관련 컨테이너 프로젝트에는 영향을 주지 않음)

## 2023-06-22

### GET - `/orgs/{org_id}/settings/sast` - 추가됨
- 조직의 SAST 설정을 검색함

## 2023-05-29

### GET - `/orgs` - 추가됨
- 액세스할 수 있는 조직 목록을 페이징하여 가져옵니다.

### PATCH - `/orgs/{org_id}` - 추가됨
- 조직의 세부 정보를 업데이트함

### GET - `/orgs/{org_id}` - 추가됨
- 조직의 전체 세부 정보를 가져옵니다.

## 2023-04-28

### POST - `/orgs/{org_id}/invites` - Updated
- 새로운 필수 요청 속성 `data`를 추가함
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 상태 `201`의 응답 중 `data/attributes/role` 속성 유형/형식을 `string`/``에서 `string`/`uuid`로 변경함
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 요청 속성 `email`을 제거함
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 요청 속성 `role`을 제거함
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `201`의 `data/type` 속성에 새로운 `org_invitation` 열거 값을 추가함
![Badge](https://img.shields.io/badge/Breaking-yellow)

### GET - `/orgs/{org_id}/invites` - Updated
- 상태 `200`의 응답 중 `data/items/attributes/role` 속성 유형/형식을 `string`/``에서 `string`/`uuid`로 변경함
![Badge](https://img.shields.io/badge/Breaking-yellow)
- 응답 상태 `200`의 `data/items/type` 속성에 새로운 `org_invitation` 열거 값을 추가함
![Badge](https://img.shields.io/badge/Breaking-yellow)

## 2023-04-17

### POST - `/orgs/{org_id}/packages/issues` - 추가됨
- 모든 고객에게 이 엔드포인트가 제공되지 않음. 관심이 있다면 지원팀에 문의하십시오. Package URL (purl)로 식별된 일괄 패키지에 대한 문제를 쿼리함. 직접 취약점만 반환되며 종속성에서의 간섭 취약점은 반환되지 않습니다.

## 2023-03-20

### GET - `/orgs/{org_id}/projects/{project_id}/sbom` - 추가됨
- 이 엔드포인트를 사용하면 소프트웨어 프로젝트의 SBOM 문서를 검색할 수 있습니다.
다음 형식을 지원합니다:
* JSON 형식의 CycloneDX 버전 1.4 (`format`을 `cyclonedx1.4+json`로 설정)
* XML 형식의 CycloneDX 버전 1.4 (`format`을 `cyclonedx1.4+xml`로 설정)
* JSON 형식의 SPDX 버전 2.3 (`format`을 `spdx2.3+json`로 설정)

기본적으로 빈 JSON:API 응답으로 응답함.

## 2023-02-15

### GET - `/orgs/{org_id}/projects` - 추가됨
- 조직의 모든 프로젝트 목록을 표시함.

### PATCH - `/orgs/{org_id}/projects/{project_id}` - 추가됨
- 프로젝트 ID로 조직의 프로젝트 하나를 업데이트함.

### GET - `/orgs/{org_id}/projects/{project_id}` - 추가됨
- 프로젝트 ID로 조직의 프로젝트 하나를 가져옵니다.

## 2022-12-14

### POST - `/orgs/{org_id}/slack_app/{bot_id}` - 추가됨
- 특정 테넌트를 위해 새로운 Slack 알림 기본 설정을 생성함.

### GET - `/orgs/{org_id}/slack_app/{bot_id}` - 추가됨
- 제공된 테넌트 ID 및 봇 ID에 대한 Slack 통합 기본 알림 설정을 가져옵니다.

### DELETE - `/orgs/{org_id}/slack_app/{bot_id}` - 추가됨
- 주어진 Slack 앱 통합을 제거함.

### GET - `/orgs/{org_id}/slack_app/{bot_id}/projects` - 추가됨
- 프로젝트를 위한 Slack 알림 설정 오버라이드를 표시함. 이 설정은 테넌트의 기본 설정을 재정의함.

### POST - `/orgs/{org_id}/slack_app/{bot_id}/projects/{project_id}` - 추가됨
- 프로젝트를 위한 Slack 설정 오버라이드를 생성함.

### PATCH - `/orgs/{org_id}/slack_app/{bot_id}/projects/{project_id}` - 추가됨
- 프로젝트를 위한 Slack 알림 설정을 업데이트함.

### DELETE - `/orgs/{org_id}/slack_app/{bot_id}/projects/{project_id}` - 추가됨
- 프로젝트를 위한 Slack 설정 오버라이드를 제거함.

## 2022-11-14

### GET - `/orgs/{org_id}/invites` - 추가됨
- 조직에 대한 보류 중인 사용자 초대 목록을 나타냅니다.

### DELETE - `/orgs/{org_id}/invites/{invite_id}` - 추가됨
- 조직에 대한 보류 중인 사용자 초대를 취소함.

## 2022-11-07

### GET - `/orgs/{org_id}/slack_app/{tenant_id}/channels` - 추가됨
- 이 조직에 Snyk Slack 앱이 설정되어야 하며, Snyk Slack 앱에서 액세스할 수 있는 채널 목록을 가져옵니다. 현재 이 컬렉션을 순방향으로만 페이징할 수 있으며 prev 링크가 생성되지 않으며 ending_before 매개변수가 작동하지 않을 수 있음.

### GET - `/orgs/{org_id}/slack_app/{tenant_id}/channels/{channel_id}` - 추가됨
- 이 조직에 Snyk Slack 앱이 설정되어야 하며, 제공된 Slack 채널 ID에 대한 Slack 채널 이름을 반환함.

## 2022-09-15

### GET - `/orgs/{org_id}/packages/{purl}/issues` - 추가됨
- Package URL (purl)로 식별된 특정 패키지 버전에 대한 문제를 쿼리함. Snyk는 직접적인 취약점만 반환합니다. 종속성에서의 간섭 취약점은 반환되지 않으며, 이는 맥락에 따라 다를 수 있습니다.

## 2022-06-01

### POST - `/orgs/{org_id}/invites` - 추가됨
- 역할을 가진 사용자에게 조직으로 초대함.

## 2022-03-11

### GET - `/self/apps` - 추가됨
- 사용자를 대신하여 작동할 수 있는 앱 목록 가져오기.

### DELETE - `/self/apps/{app_id}` - 추가됨
- 앱 ID를 사용하여 앱의 액세스 취소.

### POST - `/orgs/{org_id}/apps` - 추가됨
- 조직에 새로운 앱 생성. 사용 중단됨, 대신 `/orgs/{org_id}/apps/creations` 사용 권장.

### GET - `/orgs/{org_id}/apps` - 추가됨
- 조직에서 생성한 앱 목록 가져오기. 사용 중단됨, 대신 `/orgs/{org_id}/apps/creations` 사용 권장.

### PATCH - `/orgs/{org_id}/apps/{client_id}` - 추가됨
- 앱 속성 업데이트. 사용 중단됨, 대신 `/orgs/{org_id}/apps/creations/{app_id}` 사용 권장.

### GET - `/orgs/{org_id}/apps/{client_id}` - 추가됨
- 클라이언트 ID를 사용하여 앱 가져오기. 사용 중단됨, 대신 `/orgs/{org_id}/apps/creations/{app_id}` 사용 권장.

### DELETE - `/orgs/{org_id}/apps/{client_id}` - 추가됨
- 앱 ID를 사용하여 앱 삭제. 사용 중단됨, 대신 `/orgs/{org_id}/apps/creations/{app_id}` 사용 권장.

### POST - `/orgs/{org_id}/apps/{client_id}/secrets` - 추가됨
- 앱의 클라이언트 비밀 관리. 사용 중단됨, 대신 `/orgs/{org_id}/apps/creations/{app_id}/secrets` 사용 권장.

### GET - `/orgs/{org_id}/app_bots` - 추가됨
- 조직에 권한을 부여받은 앱 봇 목록 가져오기. 사용 중단됨, 대신 `/orgs/{org_id}/apps/installs` 사용 권장.

### DELETE - `/orgs/{org_id}/app_bots/{bot_id}` - 추가됨
- 앱 봇의 권한 취소. 사용 중단됨, 대신 `/orgs/{org_id}/apps/installs/{install_id}` 사용 권장.

## 2021-12-09

### PATCH - `/orgs/{org_id}/settings/iac` - 추가됨
- 조직의 Infrastructure as Code 설정 업데이트.

### GET - `/orgs/{org_id}/settings/iac` - 추가됨
- 조직의 Infrastructure as Code 설정 가져오기.

### PATCH - `/groups/{group_id}/settings/iac` - 추가됨
- 그룹의 Infrastructure as Code 설정 업데이트.

### GET - `/groups/{group_id}/settings/iac` - 추가됨
- 그룹의 Infrastructure as Code 설정 가져오기.