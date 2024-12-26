# API 엔드포인트 인덱스 및 팁

이 문서의 인덱스 및 노트 섹션은 이 인덱스 외에도 특정 유스 케이스에 대한 해결책 및 [Snyk APIs 사용 시나리오](scenarios-for-using-the-snyk-api.md), 그리고 다음과 같은 Snyk API 엔드포인트 사용에 대한 상세 정보 페이지를 제공합니다:

* [API를 사용하여 프로젝트를 위한 조직 및 그룹 식별](organization-and-group-identification-for-projects-using-the-api.md)
* [V1 API 엔드포인트의 프로젝트 이슈 경로](project-issue-paths-api-endpoints.md)
* [API에서 수신하는 프로젝트 유형 응답](project-type-responses-from-the-api.md)

특정 API에 대한 자세한 내용은 다음 섹션을 참조하십시오:

* [Snyk 앱 API 사용 방법](../how-to-use-snyk-apps-apis/)
* [Snyk SBOM 및 이슈 목록 API 사용 방법](../how-to-use-snyk-sbom-and-list-issues-apis/)
* [Snyk 웹훅 API 사용 방법](../how-to-use-snyk-webhooks-apis/)

추가 정보는 [API 지원 기사](https://support.snyk.io/s/topic/0TOPU00000BgWMv4AN/snyk-api)를 참조하십시오.

이 인덱스에는 REST GA 및 beta 및 V1 API 엔드포인트의 범주 및 이름이 포함되어 있으며, 각 엔드포인트에 대한 참조 문서의 URL 및 사용 가능한 경우 관련 정보로의 링크가 함께 제공됩니다. REST는 기본값이며 GA는 beta가 명시되지 않는 한 상태입니다. V1 API는 해당되는 경우에 지정됩니다. 이 인덱스는 진행 중인 작업이며, 지속적으로 추가 정보가 추가됩니다.

## AccessRequests (beta)

### [액세스 요청 가져오기](https://apidocs.snyk.io/?beta=\&version=2024-10-15#get-/self/access_requests)

## Apps

**추가 정보:** [Snyk 앱](../how-to-use-snyk-apps-apis/)

### [귀하를 대신하여 작동하는 앱 목록 가져오기](../reference/apps.md#self-apps)

### [앱 취소하기](../reference/apps.md#self-apps-app_id)

### [앱에 대한 활성 OAuth 세션 목록 가져오기](../reference/apps.md#self-apps-app_id-sessions)

### [활성 사용자 앱 세션 취소하기](../reference/apps.md#self-apps-app_id-sessions-session_id)

### [사용자에게 설치된 앱 목록 가져오기](../reference/apps.md#self-apps-installs)

### [설치 ID별로 앱의 액세스 취소하기](../reference/apps.md#self-apps-installs-install_id)

**대체:** 오류 발생 앱 권한 취소

### 오류 발생 [조직을 위한 새로운 앱 생성](../reference/apps.md#orgs-org_id-apps)

**대체:** 새로운 Snyk 앱을 위한 조직 생성

**추가 정보:** [Snyk API를 사용하여 Snyk 앱 생성](../how-to-use-snyk-apps-apis/create-a-snyk-app-using-the-snyk-api.md)

### [조직에 의해 생성된 앱 목록 가져오기](../reference/apps.md#orgs-org_id-apps-1)

**대체:** 조직에 의해 생성된 앱 목록 가져오기

**추가 정보:** [앱 세부 정보 관리](../how-to-use-snyk-apps-apis/manage-app-details.md)

### 오류 발생 [앱 속성 업데이트(이름, 리다이렉션 URI 및 액세스 토큰 유지 시간)](../reference/apps.md#orgs-org_id-apps-client_id)

**대체:** App ID를 사용하여 이름, 리다이렉션 URI 및 액세스 토큰 유지 시간과 같은 앱 생성 속성 업데이트하기

### 오류 발생 [클라이언트 ID별 앱 가져오기](../reference/apps.md#orgs-org_id-apps-client_id-1)

**대체:** App ID로 Snyk 앱 가져오기

### 오류 발생 [앱 삭제](../reference/apps.md#orgs-org_id-apps-client_id-2)

**대체:** App ID로 Snyk 앱 삭제

### 오류 발생 [앱을 위한 클라이언트 비밀 관리](../reference/apps.md#orgs-org_id-apps-client_id-secrets)

**대체:** 상호 작용하지 않는 Snyk 앱 설치의 클라이언트 비밀 관리

### [이 조직에 Snyk 앱 설치하기](../reference/apps.md#orgs-org_id-apps-installs)

### [조직에 설치된 앱 목록 가져오기](../reference/apps.md#orgs-org_id-apps-installs-1)

**대체:** 조직에 승인된 앱 봇 목록 가져오기

**추가 정보:** [슬랙 앱 (Jira 통합)](../../integrate-with-snyk/jira-and-slack-integrations/slack-app.md) (슬랙 앱 봇 ID 찾기)

### [Snyk 조직에 대한 앱 권한 취소하기](../reference/apps.md#orgs-org_id-apps-installs-install_id)

**참고:** 설치 ID로 Snyk 그룹에 대한 앱 권한 취소하기

### [상호작용하지 않는 Snyk 앱 설치용 클라이언트 비밀 관리](../reference/apps.md#orgs-org_id-apps-installs-install_id-secrets)

**추가 정보:** [앱 세부 정보 관리](../how-to-use-snyk-apps-apis/manage-app-details.md)

### [조직을 위해 새로운 Snyk 앱 생성](../reference/apps.md#orgs-org_id-apps-creations)

**대체:** 조직을 위해 새로운 앱 생성하기

**추가 정보:** [Snyk API를 사용하여 Snyk 앱 생성](../how-to-use-snyk-apps-apis/create-a-snyk-app-using-the-snyk-api.md)

### 오류 발생 [조직에 의해 생성된 앱 목록 가져오기](../reference/apps.md#orgs-org_id-apps-creations-1)

**대체:** 조직에 의해 생성된 앱 목록 가져오기

### [앱 ID를 사용하여 이름, 리다이렉션 URI 및 액세스 토큰 유지 시간 업데이트하기](../reference/apps.md#orgs-org_id-apps-creations-app_id)

**대체:** 이름, 리다이렉션 URI 및 액세스 토큰 유지 시간과 같은 앱 속성 업데이트하기

**추가 정보:** [앱 세부 정보 관리](../how-to-use-snyk-apps-apis/manage-app-details.md)

### [App ID로 Snyk 앱 가져오기](../reference/apps.md#orgs-org_id-apps-creations-app_id)

**대체:** 클라이언트 ID로 앱 가져오기

### [App ID로 앱 삭제하기](../reference/apps.md#orgs-org_id-apps-creations-app_id-2)

**대체:** App ID로 앱 삭제하기

**추가 정보:** [앱 세부 정보 관리](../how-to-use-snyk-apps-apis/manage-app-details.md)

### [Snyk 앱의 클라이언트 비밀 관리](../reference/apps.md#orgs-org_id-apps-creations-app_id-secrets)

**추가 정보:** [앱 세부 정보 관리](../how-to-use-snyk-apps-apis/manage-app-details.md)

### 오류 발생 [조직에 승인된 앱 봇 목록 가져오기](../reference/apps.md#orgs-org_id-app_bots)

**대체:** [조직에 설치된 앱 목록 가져오기](https://apidocs.snyk.io/?#get-/orgs/-org_id-/apps/installs)

**추가 정보:** [슬랙 앱](../../integrate-with-snyk/jira-and-slack-integrations/slack-app.md) (Jira 통합용)

### 오류 발생 [앱 봇 권한 취소하기](../reference/apps.md#orgs-org_id-app_bots-bot_id)

**대체:** 설치 ID로 Snyk 그룹에 대한 앱 권한 취소하기

**참고:** [앱을 위한 액세스 취소](https://apidocs.snyk.io/?#delete-/self/apps/installs/-install_id-)

### [이 그룹에 Snyk 앱 설치하기](../reference/apps.md#groups-group_id-apps-installs)

### [그룹에 설치된 앱 목록 가져오기](../reference/apps.md#groups-group_id-apps-installs-1)

### [그룹에 설치된 앱에 대한 권한 취소](../reference/apps.md#groups-group_id-apps-installs-install_id)

### [상호작용하지 않는 Snyk 앱 설치용 클라이언트 비밀 관리](../reference/apps.md#groups-group_id-apps-installs-install_id-secrets)

**대체:** 앱용 클라이언트 비밀 관리

## 감사 로그

**추가 정보:** [조직이나 그룹을 위한 API로 사용자가 시작한 활동의 감사 로그 검색하기](../../snyk-admin/user-management-with-the-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md);\
[AWS CloudTrail Lake](../../integrate-with-snyk/event-forwarding/aws-cloudtrail-lake.md)

### [조직의 감사 로그 검색](../reference/audit-logs.md#orgs-org_id-audit_logs-search)

**추가 정보:** [조직이나 그룹을 위한 API로 사용자가 시작한 활동의 감사 로그 검색하기](../../snyk-admin/user-management-with-the-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md), [AWS CloudTrail Lake](../../integrate-with-snyk/event-forwarding/aws-cloudtrail-lake.md)

### [그룹의 감사 로그 검색](../reference/audit-logs.md#groups-group_id-audit_logs-search)

**추가 정보:** [감사 로그 API의 새로운 GA REST 버전을 통해 감사 로그를 효율적으로 필터링하기](https://updates.snyk.io/filter-through-your-audit-logs-more-efficiently-with-the-new-ga-rest-version-of-the-audit-logs-api-and-api-access-is-now-opt-in-291850) (제품 업데이트); [조직이나 그룹을 위한 API로 사용자가 시작한 활동의 감사 로그 검색하기](../../snyk-admin/user-management-with-the-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md)

## 감사 로그 (v1)

### 그룹 수준 감사 로그

[그룹의 감사 로그 검색](../reference/audit-logs.md#groups-group_id-audit_logs-search) 사용

### 조직 수준 감사 로그

[조직의 감사 로그 검색](../reference/audit-logs.md#orgs-org_id-audit_logs-search) 사용

## Cloud (beta)

### [환경 목록](https://apidocs.snyk.io/?beta=\&version=2024-10-15#get-/orgs/-org_id-/cloud/environments)

### [새로운 환경 생성](https://apidocs.snyk.io/?beta=\&version=2024-10-15#post-/orgs/-org_id-/cloud/environments)

### [환경 삭제](https://apidocs.snyk.io/?beta=\&version=2024-10-15#delete-/orgs/-org_id-/cloud/environments/-environment_id-)

### [환경 업데이트](https://apidocs.snyk.io/?beta=\&version=2024-10-15#patch-/orgs/-org_id-/cloud/environments/-environment_id-)

### [클라우드 제공자 권한 생성](https://apidocs.snyk.io/?beta=\&version=2024-10-15#post-/orgs/-org_id-/cloud/permissions)

### [리소스 목록](https://apidocs.snyk.io/?beta=\&version=2024-10-15#get-/orgs/-org_id-/cloud/resources)

[Snyk IaC](../../scan-with-snyk/snyk-iac/) (사용: IaC 및 클라우드 리소스의 인벤토리 보기)

### [스캔 목록](https://apidocs.snyk.io/?beta=\&version=2024-10-15#get-/orgs/-org_id-/cloud/scans)

### [스캔 생성](https://apidocs.snyk.io/?beta=\&version=2024-10-15#post-/orgs/-org_id-/cloud/scans)

### [스캔 가져오기](https://apidocs.snyk.io/?beta=\&version=2024-10-15#get-/orgs/-org_id-/cloud/scans/-scan_id-)

## 컬렉션

이 API를 사용하려면 프로젝트 히스토리 보기 권한이 필요합니다.

**추가 정보:** [프로젝트 컬렉션 그룹화](../../snyk-admin/snyk-projects/project-collections-groupings/)

### [컬렉션 생성](../reference/collection.md#orgs-org_id-collections)

### [컬렉션 가져오기](../reference/collection.md#orgs-org_id-collections-1)

### [컬렉션 편집](../reference/collection.md#orgs-org_id-collections-collection_id)

### [컬렉션 가져오기](../reference/collection.md#orgs-org_id-collections-collection_id-1)

### [컬렉션 삭제](../reference/collection.md#orgs-org_id-collections-collection_id-2)

### [컬렉션에 프로젝트 추가](../reference/collection.md#orgs-org_id-collections-collection_id-relationships-projects)

### [지정된 컬렉션에서 프로젝트 가져오기](../reference/collection.md#orgs-org_id-collections-collection_id-relationships-projects-1)

### [컬렉션에서 프로젝트 제거](../reference/collection.md#orgs-org_id-collections-collection_id-relationships-projects-2)

## ContainerImage

### [컨테이너 이미지의 인스턴스 목록](../reference/containerimage.md#orgs-org_id-container_images)

### [컨테이너 이미지의 인스턴스 가져오기](../reference/containerimage.md#orgs-org_id-container_images-image_id)

### [컨테이너 이미지의 대상 참조 이미지 인스턴스 목록](../reference/containerimage.md#orgs-org_id-container_images-image_id-relationships-image_target_refs)

## 사용자 지정 기본 이미지

**추가 정보:** [사용자 지정 기본 이미지 권장사항 사용](../../scan-with-snyk/snyk-container/use-snyk-container/use-custom-base-image-recommendations/)

### [기존 컨테이너 프로젝트에서 사용자 지정 기본 이미지 생성](../reference/custom-base-images.md#custom_base_images)

**추가 정보:** [사용자 지정 기본 이미지 권장사항 사용](../../scan-with-snyk/snyk-container/use-snyk-container/use-custom-base-image-recommendations/), 섹션 [생성된 프로젝트를 사용자 지정 기본 이미지로 표시](../../scan프로젝트에서 API를 사용하여 조직 및 그룹 식별

### [그룹의 모든 SSO 연결 가져오기](https://apidocs.snyk.io/?version=2024-10-15#get-/groups/-group_id-/sso_connections)

### [지정된 SSO 연결을 사용하는 모든 사용자 가져오기](https://apidocs.snyk.io/?version=2024-10-15#get-/groups/-group_id-/sso_connections/-sso_id-/users)

### [그룹 SSO 연결에서 사용자 삭제하기](https://apidocs.snyk.io/?version=2024-10-15#delete-/groups/-group_id-/sso_connections/-sso_id-/users/-user_id-)

**추가 정보**: [API를 사용하여 그룹 및 조직에서 회원 삭제](../../snyk-admin/user-management-with-the-api/remove-members-from-groups-and-orgs-using-the-api.md); [API를 사용하여 조직 또는 그룹의 사용자 주도 활동에 대한 감사 로그 검색](../../snyk-admin/user-management-with-the-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md)

## 그룹 (v1)

### [그룹 내 모든 태그 나열](../reference/groups-v1.md#group-groupid-tags)

**추가 정보**: [프로젝트 태그](../../snyk-admin/introduction-to-snyk-projects/project-tags.md)

### [그룹에서 태그 삭제](../reference/groups-v1.md#group-groupid-tags-delete)

**추가 정보**: [프로젝트 태그](../../snyk-admin/introduction-to-snyk-projects/project-tags.md)

### [그룹 설정 업데이트](../reference/groups-v1.md#group-groupid-settings)

### [그룹 설정 보기](../reference/groups-v1.md#group-groupid-settings-1)

### [그룹 내 모든 역할 나열](../reference/groups-v1.md#group-groupid-roles)

**추가 정보**: [API를 사용하여 회원 역할 업데이트](../../snyk-admin/user-management-with-the-api/update-member-roles-using-the-api.md);\
[Snyk API를 사용하여 서비스 계정 관리](../../enterprise-setup/service-accounts/manage-service-accounts-using-the-snyk-api.md)

### [그룹 내 모든 조직 나열](../reference/groups-v1.md#group-groupid-orgs)

**추가 정보**: [프로젝트를 위한 조직 및 그룹 식별](organization-and-group-identification-for-projects-using-the-api.md);\
[레거시 사용자 정의 맵핑](../../enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/custom-mapping/legacy-custom-mapping.md);\
[api-import 가져오기 대상 데이터 생성](../../scan-with-snyk/snyk-tools/tool-snyk-api-import/creating-import-targets-data-for-import-command.md);\
[시나리오: 지정된 그룹의 모든 프로젝트에 대한 프로젝트 스냅샷 검색](scenarios-for-using-the-snyk-api.md#retrieve-a-project-snapshot-for-every-project-in-a-given-group);\
[시나리오: 취약점에 영향을 받는 모든 프로젝트 찾기](scenarios-for-using-the-snyk-api.md#find-all-projects-affected-by-a-vulnerability)

### [그룹 내 조직에 구성원 추가](../reference/groups-v1.md#group-groupid-org-orgid-members)

### [그룹 내 모든 구성원 나열](../reference/groups-v1.md#group-groupid-members)

**추가 정보**: [API를 사용하여 그룹 및 조직에서 회원 삭제](../../snyk-admin/user-management-with-the-api/remove-members-from-groups-and-orgs-using-the-api.md);\
[시나리오: 주어진 목록의 모든 사용자를 회사가 소유한 조직(그룹의 모든 조직)에 할당합니다](scenarios-for-using-the-snyk-api.md#assign-all-users-in-a-given-list-to-all-the-organizations-a-company-has-all-organizations-in-a-group)

## 그룹

### [그룹 사용자의 조직 멤버십 목록 가져오기](../reference/groups.md#groups-group_id-org_memberships)

### [사용자에게 역할이 있는 그룹 멤버십 생성](../reference/groups.md#groups-group_id-memberships)

### [그룹의 모든 멤버십 가져오기](../reference/groups.md#groups-group_id-memberships-1)

### [그룹 멤버십에서 역할 업데이트](../reference/groups.md#groups-group_id-memberships-membership_id)

### [그룹에서 멤버십 삭제](../reference/groups.md#groups-group_id-memberships-membership_id-1)

## IacSettings

### [조직을 위한  설정 업데이트](../reference/groups-v1.md#group-groupid-org-orgid-members)

**추가 정보**: [원격 IaC 사용자 정의 규칙 번들 사용](../../scan-with-snyk/snyk-iac/build-your-own-iac-custom-rules/current-iac-custom-rules/use-iac-custom-rules-with-cli/use-a-remote-iac-custom-rules-bundle.md)

### [조직을 위한  설정 가져오기](../reference/iacsettings.md#orgs-org_id-settings-iac-1)

### [그룹을 위한  설정 업데이트](../reference/iacsettings.md#groups-group_id-settings-iac)

**추가 정보**: [원격 IaC 사용자 정의 규칙 번들 사용](../../scan-with-snyk/snyk-iac/build-your-own-iac-custom-rules/current-iac-custom-rules/use-iac-custom-rules-with-cli/use-a-remote-iac-custom-rules-bundle.md), [파이프라인 내의 IaC 사용자 정의 규칙](../../scan-with-snyk/snyk-iac/build-your-own-iac-custom-rules/current-iac-custom-rules/iac-custom-rules-within-a-pipeline.md); [원격 IaC 사용자 정의 규칙 번들 사용](../../scan-with-snyk/snyk-iac/build-your-own-iac-custom-rules/current-iac-custom-rules/use-iac-custom-rules-with-cli/use-a-remote-iac-custom-rules-bundle.md)

### [그룹을 위한  설정 가져오기](../reference/iacsettings.md#groups-group_id-settings-iac-1)

## Ignor...

[그하 세부연락](https://apidocs.snyk.io/?version=2024-10-15)## 토큰 프로비저닝

**더 많은 정보:** [설정에 필요한 토큰 얻기](../../enterprise-setup/snyk-broker/snyk-broker-code-agent/install-snyk-broker-code-agent-using-docker/obtain-the-required-tokens-for-setup.md) (Snyk Broker Code Agent)

## 초대

[사용자를 조직으로 초대](../reference/invites.md#orgs-org_id-invites)

[조직으로의 보류 중인 사용자 초대 목록](../reference/invites.md#orgs-org_id-invites-1)

[조직으로의 보류 중인 사용자 초대 취소](../reference/invites.md#orgs-org_id-invites-invite_id)

## 이슈

[패키지에 대한 이슈 목록](../reference/issues.md#orgs-org_id-packages-purl-issues)

**더 많은 정보:**
[Dart 및 Flutter](../../supported-languages-package-managers-and-frameworks/dart-and-flutter.md);\
[Rust](../../supported-languages-package-managers-and-frameworks/rust.md):\
[페이지를 위한 `Snyk for C++` 지침, 대체 테스트 옵션 섹션](../../supported-languages-package-managers-and-frameworks/c-c++/guidance-for-snyk-for-c-c++.md#alternate-testing-options);\
[Java 및 Kotlin을 위한 지침](../../supported-languages-package-managers-and-frameworks/java-and-kotlin/guidance-for-java-and-kotlin.md);\
[JavaScript 및 Node.js를 위한 지침, 비관리 JavaScript 섹션](../../supported-languages-package-managers-and-frameworks/javascript/best-practices-for-javascript-and-node.js.md#unmanaged-javascript);\
[패키지에 대한 이슈 목록 페이지](../how-to-use-snyk-sbom-and-list-issues-apis/list-issues-for-a-package.md)

[특정 집합의 패키지에 대한 이슈 목록](../reference/issues.md#orgs-org_id-packages-issues) (모든 고객에게 제공되지 않음)

[Org ID별 이슈 조회](../reference/issues.md#orgs-org_id-issues)

**더 많은 정보:**
[시나리오: 이슈 대량 무시](scenarios-for-using-the-snyk-api.md#bulk-ignore-issues);\
[조직 내 모든 프로젝트에서 의 모든 이슈 포함 목록](scenarios-for-using-the-snyk-api.md#list-all-issues-including-snyk-code-issues-in-all-the-projects-in-an-organization)

[이슈 확인](../reference/issues.md#orgs-org_id-issues-issue_id) (조직용)

[그룹 ID별 이슈 조회](../reference/issues.md#orgs-org_id-issues-issue_id)

**참고:** 해결책은 응답에 포함되어 있지 않습니다.

추가 정보: [Reachability](../../manage-risk/prioritize-issues-for-fixing/reachability-analysis.md)

[이슈 확인](../reference/issues.md#groups-group_id-issues-issue_id) (그룹용)

## Jira (v1)

[모든 Jira 이슈 목록](../reference/jira-v1.md#org-orgid-project-projectid-jira-issues)

**더 많은 정보:** [Jira 통합](../../integrate-with-snyk/jira-and-slack-integrations/jira-integration.md); [CI/CD 통합에서 Snyk} 테스트 및 snyk monitor](../../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-test-and-snyk-monitor-in-ci-cd-integration.md)

[Jira 이슈 생성](../reference/jira-v1.md#org-orgid-project-projectid-issue-issueid-jira-issuev)

**더 많은 정보:** [Jira 통합](../../integrate-with-snyk/jira-and-slack-integrations/jira-integration.md);\
[CI/CD 통합에서 Snyk} 테스트 및 snyk monitor](../../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/snyk-ci-cd-integration-deployment-and-strategies/snyk-test-and-snyk-monitor-in-ci-cd-integration.md)

## 라이선스 (v1)

[모든 라이선스 목록](../reference/licenses-v1.md)

## 모니터 (v1)

[모니터 Dep Graph](../reference/monitor-v1.md)

**더 많은 정보:** [Dep Graph API (Bazel)](../../scan-with-snyk/snyk-open-source/bazel/#dep-graph-api)

## 조직 (v1)

**더 많은 정보:** [웹훅 이벤트 및 페이로드](../how-to-use-snyk-webhooks-apis/webhooks.md)

[사용자가 속한 모든 조직 목록](../reference/organizations-v1.md#orgs)

**더 많은 정보:** [API를 사용하여 프로젝트의 조직 및 그룹 식별](organization-and-group-identification-for-projects-using-the-api.md);\
[시나리오: 어떤 이유로든 Broker 토큰 회전 또는 변경](scenarios-for-using-the-snyk-api.md#rotate-or-change-your-broker-token-for-any-reason)

[새 조직 생성](../reference/organizations-v1.md#org)

**더 많은 정보:** [가시성 설정 및 조직 템플릿 구성](../../implement-snyk/enterprise-implementation-guide/phase-2-configure-account/set-visibility-and-configure-an-organization-template/) (엔터프라이즈 구현 가이드 Phase 2, 계정 구성);\
[api-import: Snyk에서 조직 생성](../../scan-with-snyk/snyk-tools/tool-snyk-api-import/creating-organizations-in-snyk.md);\
[시나리오: 특정 그룹에 동일한 설정을 갖는 여러 개의 새 조직 생성](scenarios-for-using-the-snyk-api.md#create-multiple-new-organizations-that-all-have-the-same-settings-in-a-given-group)

[조직 삭제](../reference/organizations-v1.md#org-orgid)

[조직 설정 업데이트](../reference/organizations-v1.md#org-orgid-settings)

조직 설정의 유일한 편집 가능한 속성은 `requestAccess`입니다.

**더 많은 정보:** [특정 그룹에 동일한 설정을 갖는 여러 개의 새 조직 생성](scenarios-for-using-the-snyk-api.md#create-multiple-new-organizations-that-all-have-the-same-settings-in-a-given-group)

[조직 설정 보기](../reference/organizations-v1.md#org-orgid-settings-1)

**더 많은 정보:** [특정 그룹에 동일한 설정을 갖는 여러 개의 새 조직 생성](scenarios-for-using-the-snyk-api.md#create-multiple-new-organizations-that-all-have-the-same-settings-in-a-given-group)

[조직에 사용자 프로비저닝](../reference/organizations-v1.md#org-orgid-provision)

**더 많은 정보:** [API를 사용하여 조직에 사용자 프로비저닝](../../snyk-admin/user-management-with-the-api/provision-users-to-organizations-using-the-api.md):\
[시나리오: 첫 로그인 이전에 규모에 맞게 사용자를 조직에 추가](scenarios-for-using-the-snyk-api.md#add-users-to-organizations-at-scale-ahead-of-the-first-login)

[보류 중인 사용자 프로비저닝 목록](../reference/organizations-v1.md#org-orgid-provision-1)

**더 많은 정보:** [API를 사용하여 조직에 사용자 프로비저닝](../../snyk-admin/user-management-with-the-api/provision-users-to-organizations-using-the-api.md)

[보류 중인 사용자 프로비저닝 삭제](../reference/organizations-v1.md#org-orgid-provision-2)

**더 많은 정보:** [API를 사용하여 조직에 사용자 프로비저닝](../../snyk-admin/user-management-with-the-api/provision-users-to-organizations-using-the-api.md)

[알림 설정](../reference/organizations-v1.md#org-orgid-notification-settings)

**더 많은 정보:** [가져오기 대상 데이터 생성을 위한 API](../../scan-with-snyk/snyk-tools/tool-snyk-api-import/creating-import-targets-data-for-import-command.md);\
[도구: snyk-api-import](../../scan-with-snyk/snyk-tools/tool-snyk-api-import/)

[조직 알림 설정 가져오기](../reference/organizations-v1.md#org-orgid-notification-settings-1)

[모든 구성원 목록](../reference/organizations-v1.md#org-orgid-members)

**더 많은 정보:** [API를 사용하여 구성원 역할 업데이트](../../snyk-admin/user-management-with-the-api/update-member-roles-using-the-api.md); [API를 사용하여 그룹 및 조직에서 구성원 삭제](../../snyk-admin/user-management-with-the-api/remove-members-from-groups-and-orgs-using-the-api.md)

[조직 내 구성원 업데이트](../reference/organizations-v1.md#org-orgid-members-userid)

**더 많은 정보:** [사용자 역할 관리](../../snyk-admin/user-roles/user-role-management.md)

[조직에서 구성원 제거](../reference/organizations-v1.md#org-orgid-members-userid-1)

**더 많은 정보:** [API를 사용하여 그룹 및 조직에서 구성원 삭제](../../snyk-admin/user-management-with-the-api/remove-members-from-groups-and-orgs-using-the-api.md);\
[사용자 역할 관리](../../snyk-admin/user-roles/user-role-management.md)

[조직에서 구성원 역할 업데이트](../reference/organizations-v1.md#org-orgid-members-update-userid)

**더 많은 정보:** [사용자 역할 관리](../../snyk-admin/user-roles/user-role-management.md); [API를 사용하여 구성원 역할 업데이트](../../snyk-admin/user-management-with-the-api/update-member-roles-using-the-api.md)

[사용자 초대](../reference/organizations-v1.md#org-orgid-invite)

**더 많은 정보:** [API를 사용하여 구성원 역할 업데이트](../../snyk-admin/user-management-with-the-api/update-member-roles-using-the-api.md);\
[시나리오: 회사의 모든 조직(그룹 내 모든 조직)에 특정 목록의 모든 사용자 할당](scenarios-for-using-the-snyk-api.md#assign-all-users-in-a-given-list-to-all-the-organizations-a-company-has-all-organizations-in-a-group)

## Orgs (GA 및 beta)

[접근 가능한 모든 조직 목록](../reference/orgs.md#orgs)

**더 많은 정보:** [Snyk 앱 사전 요구 사항](../how-to-use-snyk-apps-apis/prerequisites-for-snyk-apps.md);\
[API를 사용하여 프로젝트의 조직 및 그룹 식별](organization-and-group-identification-for-projects-using-the-api.md)

[조직 업데이트](../reference/orgs.md#orgs-org_id)

[역할을 갖는 사용자를 위한 사용자의 org 멤버십 생성](../reference/orgs.md#orgs-org_id-memberships)

[조직의 모든 멤버십 가져오기](../reference/orgs.md#orgs-org_id-memberships-1)

[역할을 갖는 사용자의 조직 멤버십 업데이트](../reference/orgs.md#orgs-org_id-memberships-membership_id)

[그룹 내 모든 조직 목록](../reference/orgs.md#groups-group_id-orgs)

[ORG 가져오기](https://apidocs.snyk.io/?version=2024-10-15#get-/orgs/-org_id-) (beta)

**더 많은 정보:** [프로젝트를 위한 Org 및 그룹 식별](organization-and-group-identification-for-projects-using-the-api.md)

## 프로젝트 (v1)

**더 많은 정보:** [API에서 프로젝트 유형 응답](project-type-responses-from-the-api.md);\
[웹훅 이벤트 및 페이로드](../how-to-use-snyk-webhooks-apis/webhooks.md)

[프로젝트 업데이트](../reference/projects-v1.md#org-orgid-project-projectid)

[단일 프로젝트 검색](../reference/projects-v1.md#org-orgid-project-projectid-1)

**더 많은 정보:** [API에서 프로젝트 유형 응답](project-type-responses-from-the-api.md)

[프로젝트 삭제](../reference/projects-v1.md#org-orgid-project-projectid-2)

**더 많은 정보:** [어노테이티드 가져오기](../../scan-with-snyk/snyk-container/kubernetes-integration/annotated-import.md) (Kubernetes 통합 섹션); [API에서 프로젝트 유형 응답](project-type-responses-from-the-api.md); [시나리오: 새 컨테이너 이미지 가져오기](scenarios-for-using-the-snyk-api.md#import-fresh-container-images)

[프로젝트에 태그 추가](../reference/projects-v1.md#org-orgid-project-projectid-tags)

**더 많은 정보:** [프로젝트 태그](../../snyk-admin/introduction-to-snyk-projects/project-tags.md); [인사이트 설정: , Code, 그리고 Container Projects 연관 설정](../../manage-risk/prioritize-issues-for-fixing/set-up-insights-for-snyk-apprisk/set-up-insights-associating-snyk-open-source-code-and-container-projects.md);\
[시나리오: 어떤 이유로든 Broker 토큰 회전 또는 변경](scenarios-for-using-the-snyk-api.md#rotate-or-change-your-broker-token-for-any-reason)

[프로젝트에서 태그 제거](../reference/projects-v1.md#org-orgid-project-projectid-tags-remove)

**더 많은 정보:** [프로젝트 태그](../../snyk-admin/introduction-to-snyk-projects/project-tags.md)

[프로젝트 설정 업데이트](../reference/projects-v1.md#org-orgid-project-projectid-settings)

[프로젝트 설정 목록](../reference/projects-v1.md#org-orgid-project-projectid-settings-1)

[프로젝트 설정 삭제](../reference/projects-v1.md#org-orgid-project-projectid-settings-2)

[다른 조직으로 프로젝트 이동](../reference/projects-v1.md#org-orgid-project-projectid-move)

**더 많은 정보:** [시나리오: 프로젝트를 다른 조직으로 이동](scenarios-for-using-the-snyk-api.md#move-projects-from-one-organization-to-another)

[프로젝트 이슈 경로 목록](../reference/projects-v1.md#org-orgid-project-projectid-issue-issueid-paths)

**더 많은 정보:** [프로젝트 이슈 경로 API 엔드포인트](project-issue-paths-api-endpoints.md)

[프로젝트 종속성 그래프 가져오기](../reference/projects-v1.md#org-orgid-project-projectid-dep-graph)

[비활성화](../reference/projects-v1.md#org-orgid-project-projectid-deactivate) (프로젝트)

[적용 중(프로젝트) 속성](../reference/projects-v1.md#org-orgid-project-projectid-attributes)

API 엔드포인트인 속성 적용을 사용하여 프로젝트가 생성된 후 비즈니스 중요도, 라이프사이클 단계 및 환경을 포함한 Snyk 프로젝트에 속성을 설정할 수 있습니다. 다음과 같이 진행합니다:

* [가져오기 대상](../reference/import-projects-v1.md#org-orgid-integrations-integrationid-import) API 엔드포인트를 사용하여 프로젝트 가져오기
* [가져오기 대상](../reference/import-projects-v1.md#org-orgid이유를 그룹화합니다. Snyk 는 문제를 식별자별로 그룹화하여 List all aggregated issues endpoint에 대한 한 응답은 여러 경로를 통해 동일한 문제에 해당할 수 있습니다. 따라서 `ignoredReason`은 집계된 모든 문제에 걸쳐 적용되며 해당 단일 그룹화된 문제에 적용됩니다.

**추가 정보:** [시나리오: 조직 내 모든 프로젝트에 포함된 Snyk Code 이슈를 포함한 모든 이슈 나열](scenarios-for-using-the-snyk-api.md#list-all-issues-including-snyk-code-issues-in-all-the-projects-in-an-organization)

### [활성화](../reference/projects-v1.md#org-orgid-project-projectid-activate) (프로젝트)

## 프로젝트

### [주어진 조직 ID로 모든 프로젝트 나열](../reference/projects.md#orgs-org_id-projects)

`types`에 대한 쿼리 문자열 매개변수는 선택 사항입니다. 엔드포인트는 특정 프로젝트 유형을 강제하지 않으며 프로젝트 유형과 일치하지 않는 문자열을 입력하면 `일치하는 프로젝트 없음`을 반환합니다.

**추가 정보:** [Slack 앱 (Jira 통합)](../../integrate-with-snyk/jira-and-slack-integrations/slack-app.md) (사용: 프로젝트 ID 찾기);\
[Snyk 프로젝트](../../snyk-admin/snyk-projects/);\
[프로젝트 정보](../../snyk-admin/snyk-projects/project-information.md);\
[프로젝트 속성](../../snyk-admin/snyk-projects/project-attributes.md);\
[시나리오: 취약성에 영향을 받는 모든 프로젝트 찾기](scenarios-for-using-the-snyk-api.md#find-all-projects-affected-by-a-vulnerability);\
[시나리오: 조직 내 모든 프로젝트에 Snyk Code 이슈를 포함한 모든 이슈 나열](scenarios-for-using-the-snyk-api.md#list-all-issues-including-snyk-code-issues-in-all-the-projects-in-an-organization);\
[시나리오: 이슈 일괄 무시](scenarios-for-using-the-snyk-api.md#bulk-ignore-issues);\
[시나리오: Snyk 내 모든 프로젝트에 태그 지정](scenarios-for-using-the-snyk-api.md#tag-all-projects-in-snyk);\
[시나리오: 새로운 컨테이너 이미지 가져오기](scenarios-for-using-the-snyk-api.md#import-fresh-container-images);\
[시나리오: 리포지토리에서 새 프로젝트 감지 및 가져오기](scenarios-for-using-the-snyk-api.md#detect-and-import-new-projects-in-a-repository-into-a-target)

### [프로젝트 ID로 프로젝트 업데이트](../reference/projects.md#orgs-org_id-projects-project_id)

**추가 정보:** [프로젝트 설정 보기 및 편집](../../snyk-admin/snyk-projects/view-and-edit-project-settings.md);\
[스캐닝 시작](../../scan-with-snyk/start-scanning.md) (사용: 테스트 빈도 설정)

### [프로젝트 ID로 프로젝트 가져오기](../reference/projects.md#orgs-org_id-projects-project_id-1)

### [프로젝트 ID로 프로젝트 삭제](../reference/projects.md#orgs-org_id-projects-project_id-2)

## 풀 리퀘스트 템플릿

### [그룹용 풀 리퀘스트 템플릿 생성 또는 업데이트](../reference/pull-request-templates.md#groups-group_id-settings-pull_request_template)

**추가 정보:** [API를 사용하여 사용자 정의 PR 템플릿 만들고 관리하기](../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/customize-pr-templates/apply-a-custom-pr-template.md#create-and-manage-a-custom-pr-template-using-the-api)### 새로운 프로젝트 파일을 감지하고 정기적으로 Snyk에 가져오기 

[여기에서 자세한 내용을 확인하세요](scenarios-for-using-the-snyk-api.md#detect-new-projects-files-in-repositories-and-import-them-into-a-target-in-snyk-on-a-regular-basis);

### 시나리오: 정기적으로 Snyk에 새로운 프로젝트(파일)를 감지하고 대상으로 가져오기 

[여기에서 자세한 내용을 확인하세요](scenarios-for-using-the-snyk-api.md#identify-and-import-new-repositories-only)

## 타겟 ID로 타겟 가져오기

[여기에서 자세한 내용을 확인하세요](../reference/targets.md#orgs-org_id-targets-target_id)

## 타겟 ID로 타겟 삭제하기

[여기에서 자세한 내용을 확인하세요](../reference/targets.md#orgs-org_id-targets-target_id-1)

## 테스트 (v1)

**더 많은 정보:** [Java 및 Kotlin에 대한 가이드](../../supported-languages-package-managers-and-frameworks/java-and-kotlin/guidance-for-java-and-kotlin.md);\
[스캔 시작하기](../../scan-with-snyk/start-scanning.md);\
[오픈 소스 라이브러리 및 라이센스 스캔](../../scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/)

### package.json 및 yarn-lock 파일 테스트

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-yarn)

### sbt 파일 테스트

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-sbt)

### group id, artifact id, 버전별로 공개 패키지의 문제를 위한 sbt 테스트

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-sbt-groupid-artifactid-version)

### gemfile.lock 파일 테스트

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-rubygems)

### 이름 및 버전별로 공개 gem의 문제를 위한 테스트

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-rubygems-gemname-version)

### requirements.txt 파일 테스트 (pip)

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-pip)

### 이름과 버전별로 공개 (pip) 패키지의 문제를 위한 pip 테스트

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-pip-packagename-version)

### package.json 및 package-lock.json 파일 테스트

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-npm)

### 이름과 버전별로 공개 패키지의 문제를 위한 npm 테스트

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-npm-packagename-version)

**더 많은 정보:** [JavaScript 및 Node.js에 대한 지침, Unmanaged JavaScript 섹션](../../supported-languages-package-managers-and-frameworks/javascript/best-practices-for-javascript-and-node.js.md#unmanaged-javascript)

### maven 파일 테스트

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-maven)

### group id, artifact id, 버전별로 공개 패키지의 문제를 위한 maven 테스트

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-maven-groupid-artifactid-version)

**더 많은 정보:** [Java 및 Kotlin에 대한 가이드](../../supported-languages-package-managers-and-frameworks/java-and-kotlin/guidance-for-java-and-kotlin.md)

### gradle 파일 테스트 

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-gradle)

### group, 이름, 버전별로 공개 패키지의 문제를 위한 gradle 테스트

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-gradle-group-name-version)

### vendor.json 파일 테스트

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-govendor)

### Gopkg.toml 및 Gopkg.lock 파일 테스트

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-golangdep)

### Dep Graph 테스트

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-dep-graph)

**더 많은 정보:** [Dep Graph API](../../scan-with-snyk/snyk-open-source/bazel/dep-graph-api.md) (Bazel);\
[Unmanaged JavaScript](../../supported-languages-package-managers-and-frameworks/javascript/best-practices-for-javascript-and-node.js.md#unmanaged-javascript) (JavaScript 및 Node.js에 대한 지침);\
[스캔 시작하기](../../scan-with-snyk/start-scanning.md)

### composer.json 및 composer.lock 파일 테스트

[여기에서 자세한 내용을 확인하세요](../reference/test-v1.md#test-composer)

## 사용자 (v1)

### 사용자 세부 정보 가져오기

[여기에서 자세한 내용을 확인하세요](../reference/users-v1.md#user-userid)

### 내 세부 정보 가져오기

[여기에서 자세한 내용을 확인하세요](../reference/users-v1.md#user-me)

### 조직 알림 설정 수정

[여기에서 자세한 내용을 확인하세요](../reference/users-v1.md#user-me-notification-settings-org-orgid)

### 조직 알림 설정 가져오기 

[여기에서 자세한 내용을 확인하세요](../reference/users-v1.md#user-me-notification-settings-org-orgid-1)

### 프로젝트 알림 설정 수정

[여기에서 자세한 내용을 확인하세요](../reference/users-v1.md#user-me-notification-settings-org-orgid-project-projectid)

### 프로젝트 알림 설정 가져오기

[여기에서 자세한 내용을 확인하세요](../reference/users-v1.md#user-me-notification-settings-org-orgid-project-projectid-1)

## 사용자

### 나의 사용자 세부 정보

[여기에서 자세한 내용을 확인하세요](../reference/users.md)

### 그룹에서 사용자 역할 업데이트

[여기에서 자세한 내용을 확인하세요](https://apidocs.snyk.io/?version=2024-10-15#patch-/groups/-group_id-/users/-id-) (베타)

**참고:** 이 엔드포인트를 사용하여 그룹에서 사용자를 제거합니다.

**더 많은 정보:** [API를 사용하여 그룹 및 조직에서 구성원 제거하기](../../snyk-admin/user-management-with-the-api/remove-members-from-groups-and-orgs-using-the-api.md)

### ID로 사용자 가져오기

[여기에서 자세한 내용을 확인하세요](https://apidocs.snyk.io/?version=2024-10-15#get-/orgs/-org_id-/users/-id-) (베타)

## Webhooks (v1)

### 웹훅 생성

[여기에서 자세한 내용을 확인하세요](../reference/webhooks-v1.md#org-orgid-webhooks)

**더 많은 정보:** [시나리오: 특정 이벤트나 시간에, Snyk에서 코드베이스(소스 제어 관리)로의 모든 상호작용(풀 리퀘스트, 테스트) 비활성화하기](scenarios-for-using-the-snyk-api.md#for-a-specific-event-or-time-disable-all-interactions-pull-requests-tests-from-snyk-to-the-code-base)

### 웹훅 목록

[여기에서 자세한 내용을 확인하세요](../reference/webhooks-v1.md#org-orgid-webhooks-1)

**더 많은 정보:**

### 웹훅 검색

[여기에서 자세한 내용을 확인하세요](../reference/webhooks-v1.md#org-orgid-webhooks-webhookid)

### 웹훅 삭제

[여기에서 자세한 내용을 확인하세요](../reference/webhooks-v1.md#org-orgid-webhooks-webhookid-1)

**더 많은 정보:** [시나리오: 특정 이벤트나 시간에, Snyk에서 코드베이스(소스 제어 관리)로의 모든 상호작용(풀 리퀘스트, 테스트) 비활성화하기](scenarios-for-using-the-snyk-api.md#for-a-specific-event-or-time-disable-all-interactions-pull-requests-tests-from-snyk-to-the-code-base)

### 웹훅 ping

\\