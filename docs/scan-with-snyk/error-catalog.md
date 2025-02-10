# 오류 카탈로그

## Snyk 오류 코드

아래 표에 나와있는 오류 코드는 [Snyk API](../snyk-api/)나 [CLI](../snyk-cli/)를 사용하면서 발생할 수 있는 코드를 설명합니다. API를 사용하는 중에 오류가 발생하면 적절한 [HTTP 상태 코드](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes)도 함께 표시될 것입니다. 오류 코드 없이 오류가 발생하면 적절한 조치를 결정하기 위해 HTTP 상태 코드를 사용하십시오.

***

## 수정

#### [PR-FAILURES-0001](error-catalog.md#pr-failures-0001)

**수정 시나리오가 지원되지 않음**

시나리오가 지원되지 않아 Snyk이 수정 PR을 열지 못했습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [PR-FAILURES-0002](error-catalog.md#pr-failures-0002)

**SCM 레이트 제한**

너무 많은 요청으로 SCM 레이트 제한이 초과되었습니다.

**HTTP 상태:** [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429)

#### [PR-FAILURES-0003](error-catalog.md#pr-failures-0003)

**미인증 액세스**

사용자 추가 및 허용된 역할에 대한 문서를 읽어주십시오. 미인증 액세스로 인해 요청이 실패했습니다.

**HTTP 상태:** [403](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403)

**도움링크:**

* [https://docs.snyk.io/snyk-admin/groups-and-organizations/organizations/manage-users-in-organizations](https://docs.snyk.io/snyk-admin/groups-and-organizations/organizations/manage-users-in-organizations)

#### [SNYK-PACKAGES-0001](error-catalog.md#snyk-packages-0001)

**지원되지 않는 생태계**

언어 또는 패키지 관리자가 지원되지 않습니다. 문서에서 지원되는 패키지 관리자를 참조하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**도움링크:**

* [https://docs.snyk.io/scan-using-snyk/pull-requests/snyk-fix-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs/upgrade-open-source-dependencies-with-automatic-prs#supported-languages-and-scms](https://docs.snyk.io/scan-using-snyk/pull-requests/snyk-fix-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs/upgrade-open-source-dependencies-with-automatic-prs#supported-languages-and-scms)

#### [SNYK-PACKAGES-0003](error-catalog.md#snyk-packages-0003)

**메타데이터를 찾을 수 없음**

패키지 메타데이터가 없거나 찾을 수 없습니다.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**도움링크:**

* [https://docs.snyk.io/scan-using-snyk/pull-requests/snyk-fix-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs/upgrade-private-dependencies-with-automatic-prs#private-packages-api](https://docs.snyk.io/scan-using-snyk/pull-requests/snyk-fix-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs/upgrade-private-dependencies-with-automatic-prs#private-packages-api)

#### [SNYK-PACKAGES-0005](error-catalog.md#snyk-packages-0005)

**패키지에 성숙 버전이 없음**

성숙 버전을 찾을 수 없어 권장 버전을 제공할 수 없습니다.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

#### [SNYK-PACKAGES-0006](error-catalog.md#snyk-packages-0006)

**권장 버전을 찾을 수 없음**

이 정책을 사용하여 패키지에 대한 권장 버전을 제공할 수 없습니다.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

#### [SNYK-PACKAGES-0007](error-catalog.md#snyk-packages-0007)

**패키지가 이미 최신 버전입니다**

이 패키지의 더 새로운 버전을 찾을 수 없습니다. 이미 최신 버전이므로 더 이상 업그레이드할 수 없습니다.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

#### [SNYK-PACKAGES-0008](error-catalog.md#snyk-packages-0008)

**버전 다운그레이드가 지원되지 않음**

패키지 버전을 다운그레이드할 수 없습니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-PACKAGES-0009](error-catalog.md#snyk-packages-0009)

**잘못된 버전**

올바른 버전이 아닙니다. Semver 형식에 맞지 않습니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**도움링크:**

* [https://semver.org/](https://semver.org/)

#### [SNYK-PR-TEMPLATE-0001](error-catalog.md#snyk-pr-template-0001)

**풀 리퀘스트 속성을 가져오지 못했습니다**

주어진 변수와 가져온 PR 템플릿을 사용하여 사용자 지정 풀 리퀘스트 템플릿 속성을 가져올 수 없습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**도움링크:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

#### [SNYK-PR-TEMPLATE-0002](error-catalog.md#snyk-pr-template-0002)

**찾을 수 없음**

풀 리퀘스트 템플릿을 찾을 수 없습니다. 이미 만들었나요? 풀 리퀘스트 템플릿 설정 방법에 대한 지시사항은 첨부된 링크를 확인하십시오.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**도움링크:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

#### [SNYK-PR-TEMPLATE-0003](error-catalog.md#snyk-pr-template-0003)

**풀 리퀘스트 템플릿을 컴파일하지 못했습니다**

사용자 지정 풀 리퀘스트 템플릿을 컴파일할 수 없습니다. 템플릿 내의 Snyk 변수를 확인하여 구문 오류를 확인하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**도움링크:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

#### [SNYK-PR-TEMPLATE-0004](error-catalog.md#snyk-pr-template-0004)

**풀 리퀘스트 속성을 구문 분석하지 못했습니다**

주어진 변수를 사용하여 사용자 지정 풀 리퀘스트 템플릿을 구문 분석할 수 없습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**도움링크:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

#### [SNYK-PR-TEMPLATE-0005](error-catalog.md#snyk-pr-template-0005)

**Snyk 변수를 대체한 후 YAML 파일을 로드하지 못했습니다**

사용자 지정 PR 템플릿에 Snyk 변수를 대체한 후 YAML 파일을 로드할 수 없습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

#### [SNYK-PR-TEMPLATE-0006](error-catalog.md#snyk-pr-template-0006)

**사용자 지정 PR 템플릿의 해시 생성에 실패했습니다**

고객 PR 파일 및 프로젝트 vulnIds를 사용하여 해시를 생성하지 못했습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**도움링크:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

#### [SNYK-PR-TEMPLATE-0007](error-catalog.md#snyk-pr-template-0007)

**풀 리퀘스트 템플릿을 만들 수 없습니다**

Snyk이 풀 리퀘스트 템플릿을 만들지 못했습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**도움링크:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

#### [SNYK-PR-TEMPLATE-0008](error-catalog.md#snyk-pr-template-0008)

**풀 리퀘스트 템플릿을 가져오지 못했습니다**

Snyk이 풀 리퀘스트 템플릿을 가져오지 못했습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**도움링크:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

#### [SNYK-PR-TEMPLATE-0009](error-catalog.md#snyk-pr-template-0009)

**풀 리퀘스트 템플릿을 삭제하지 못했습니다**

Snyk이 풀 리퀘스트 템플릿을 삭제하지 못했습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**도움링크:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

#### [SNYK-PR-TEMPLATE-0010](error-catalog.md#snyk-pr-template-0010)

**잘못된 페이로드**

풀 리퀘스트 템플릿 페이로드가 잘못되었습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

#### [SNYK-PR-TEMPLATE-0011](error-catalog.md#snyk-pr-template-0011)

**Snyk 변수를 대체한 후 JSON 파일을 로드하지 못했습니다**

사용자 지정 PR 템플릿에 Snyk 변수를 대체한 후 JSON 파일을 로드할 수 없습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**도움링크:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

#### [SNYK-PR-TEMPLATE-0012](error-catalog.md#snyk-pr-template-0012)

**기본 PR 템플릿을 렌더링하지 못했습니다**

기본 PR 템플릿을 렌더링할 수 없습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**도움링크:**

* [https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta](https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta)

***

## Snyk

#### [SNYK-0001](error-catalog.md#snyk-0001)

**서비스 일시적으로 제한됨**

요청 속도 제한이 초과되었습니다. 몇 분 후에 다시 시도해주십시오.

**HTTP 상태:** [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429)

#### [SNYK-0002](error-catalog.md#snyk-0002)

**서버 오류 응답**

서버가 요청 방법을 인식하지 못하거나 이행할 수 없을 때 발생합니다. 요청을 검토하고 다시 시도하십시오.

**HTTP 상태:** [501](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/501)

**도움링크:**

* [https://docs.snyk.io/snyk-api-info](https://docs.snyk.io/snyk-api-info)

#### [SNYK-0003](error-catalog.md#snyk-0003)

**클라이언트 요청을 처리할 수 없음**

클라이언트 오류(예: 잘못된 요청 구문, 크기가 너무 큼, 잘못된 요청 메시지 프레이밍 또는 속임수적인 요청 라우팅)로 인해 서버가 요청을 처리할 수 없습니다. 요청을 검토하고 다시 시도하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-0004](error-catalog.md#snyk-0004)

**서버 통신 오류**

요청 중에 서버 타임아웃이 발생했습니다. Snyk 상태를 확인한 후 다시 시도하십시오.

**HTTP 상태:** [504](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/504)

**도움링크:**

* [https://status.snyk.io/](https://status.snyk.io/)

#### [SNYK-0005](error-catalog.md#snyk-0005)

**인증 오류**

인증 자격 증명이 인식되지 않거나 사용자 액세스 권한이 부여되지 않았습니다. 자격 증명을 수정하고 다시 시도하거나 Snyk 관리자에게 액세스를 요청하십시오.

**HTTP 상태:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

#### [SNYK-0006](error-catalog.md#snyk-0006)

**테스트 한도에 도달함**

Snyk 플랜에서 최대 테스트 횟수에 도달했습니다. 이로 인해`org=ORG_ID`를 사용하여 `snyk config set org=ORG_ID`를 실행하십시오.

개별적인 `snyk monitor` 실행에서 이러한 전역 구성을 덮어쓰려면, `run snyk test --org=ORG_ID`나 `snyk monitor --org=ORG_ID`를 실행하십시오.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://docs.snyk.io/snyk-admin/snyk-projects/project-tags](https://docs.snyk.io/snyk-admin/snyk-projects/project-tags)

#### [SNYK-9999](error-catalog.md#snyk-9999)

**서버 오류로 요청이 충족되지 않음**

예기치 않은 오류로 서버가 요청을 처리할 수 없습니다. Snyk 상태를 확인한 후 다시 시도하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**도움링크:**

* [https://status.snyk.io/](https://status.snyk.io/)

***

## 사용자 정의 베이스 이미지

#### [SNYK-CBI-0001](error-catalog.md#snyk-cbi-0001)

**버전 지정 스키마가 태그를 지원하지 않음**

사용 중인 버전 지정 스키마는 지정된 태그를 지원하지 않습니다. 태그를 포함한 버전 지정 스키마를 업데이트하십시오.

사용자 정의 베이스 이미지의 태그가 올바르면 버전 지정 스키마를 수정해야 합니다. 저장소가 현재 Semver를 사용하고 있고 "1.2.5.7"이라는 새 태그를 추가해야 한다면 사용자 정의 버전 지정 스키마를 사용할 수 있습니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**도움링크:**

* [https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/use-custom-base-image-recommendations/versioning-schema-for-custom-base-images](https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/use-custom-base-image-recommendations/versioning-schema-for-custom-base-images)

#### [SNYK-CBI-0002](error-catalog.md#snyk-cbi-0002)

**필수 매개변수 누락**

ORG ID 또는 GROUP ID를 제공하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-CBI-0003](error-catalog.md#snyk-cbi-0003)

**프로젝트가 존재하지 않음**

프로젝트를 찾을 수 없습니다. 프로젝트가 존재하며 프로젝트에 액세스할 수 있는지, 제공된 ID가 프로젝트 ID이며 CBI ID가 아닌지 확인하십시오.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

#### [SNYK-CBI-0004](error-catalog.md#snyk-cbi-0004)

**프로젝트가 컨테이너 이미지가 아님**

프로젝트가 컨테이너 이미지가 아닙니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**도움링크:**

* [https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/use-custom-base-image-recommendations](https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/use-custom-base-image-recommendations)

#### [SNYK-CBI-0005](error-catalog.md#snyk-cbi-0005)

**그룹을 검색할 수 없음**

프로젝트의 org가 그룹에 속하지 않습니다. 커스텀 베이스 이미지를 사용하려면 프로젝트를 다시 만들어 그룹에 추가하거나 org에 그룹을 추가하십시오. 무료 사용자에게는 그룹 기능이 제공되지 않음을 유의하세요.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-CBI-0006](error-catalog.md#snyk-cbi-0006)

**요청의 값이 일치하지 않음**

요청 본문 ID와 요청 경로 ID가 일치하지 않습니다. 값이 동일한지 확인한 후 다시 시도하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-CBI-0007](error-catalog.md#snyk-cbi-0007)

**요청 본문을 업데이트할 수 없음**

요청 본문에 업데이트할 수 있는 속성이 포함되어 있지 않습니다. 필요한 속성을 제공한 후 다시 시도하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-CBI-0008](error-catalog.md#snyk-cbi-0008)

**잘못된 페이징 커서**

제공된 페이징 커서가 잘못되었습니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-CBI-0009](error-catalog.md#snyk-cbi-0009)

**버전으로 정렬할 수 없음**

Snyk이 버전별로 필터링할 수 없습니다. 저장소 필터를 제공한 후 다시 시도하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-CBI-0010](error-catalog.md#snyk-cbi-0010)

**버전 지정 스키마를 업데이트할 수 없음**

버전 지정 스키마가 저장소의 모든 이미지에 적용되지 않았습니다. 따라서 리소스가 업데이트되지 않았습니다. 제공된 버전 지정 스키마를 업데이트하여 저장소의 모든 태그가 새 스키마에 맞는지 확인하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-CBI-0011](error-catalog.md#snyk-cbi-0011)

**프로젝트가 이미 커스텀 베이스 이미지에 연결됨**

제공된 프로젝트 ID가 이미 다른 커스텀 베이스 이미지에 연결되어 있습니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-CBI-0012](error-catalog.md#snyk-cbi-0012)

**저장소에 대한 버전 지정 스키마가 없음**

저장소에 버전 지정 스키마가 없습니다. 이 이미지는 해당 저장소의 첫 번째 이미지입니다. 이 저장소의 현재 및 미래 이미지 형식과 일치하는 버전 지정 스키마를 제공하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-CBI-0013](error-catalog.md#snyk-cbi-0013)

**버전 지정 스키마를 적용할 수 없음**

저장소에 이미 버전 지정 스키마가 있습니다. "versioning\_schema" 속성을 제거하거나 버전 지정 스키마를 업데이트하려면 PATCH 엔드포인트를 사용하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-CBI-0014](error-catalog.md#snyk-cbi-0014)

**커스텀 베이스 이미지를 찾을 수 없음**

요청한 커스텀 베이스 이미지를 찾을 수 없습니다. 다시 시도하고 오류가 지속되면 Snyk 지원팀에 문의하십시오.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

#### [SNYK-CBI-0015](error-catalog.md#snyk-cbi-0015)

**커스텀 베이스 이미지가 존재하지 않음**

요청한 커스텀 베이스 이미지가 존재하지 않습니다.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**도움링크:**

* [https://status.snyk.io/](https://status.snyk.io/)

#### [SNYK-CBI-0016](error-catalog.md#snyk-cbi-0016)

**커스텀 베이스 이미지 업데이트 실패**

커스텀 베이스 이미지를 업데이트하는 동안 내부 오류가 발생했습니다. 다시 시도하고 오류가 지속되면 Snyk 지원팀에 문의하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**도움링크:**

* [https://status.snyk.io/](https://status.snyk.io/)

#### [SNYK-CBI-0017](error-catalog.md#snyk-cbi-0017)

**프로젝트 속성을 검색할 수 없음**

프로젝트 속성을 검색하는 동안 내부 오류가 발생했습니다. 다시 시도하고 오류가 지속되면 Snyk 지원팀에 문의하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**도움링크:**

* [https://status.snyk.io/](https://status.snyk.io/)

#### [SNYK-CBI-0018](error-catalog.md#snyk-cbi-0018)

**이미지 컬렉션을 검색할 수 없음**

이미지 컬렉션을 검색하는 동안 내부 오류가 발생했습니다. 다시 시도하고 오류가 지속되면 Snyk 지원팀에 문의하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**도움링크:**

* [https://status.snyk.io/](https://status.snyk.io/)

#### [SNYK-CBI-0019](error-catalog.md#snyk-cbi-0019)

**버전 지정 스키마를 생성할 수 없음**

제공된 버전 지정 스키마가 잘못되어 이미지를 생성할 수 없습니다. 올바르게 형식화된 버전 지정 스키마를 제공한 후 다시 시도하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

***

## CLI

#### [SNYK-CLI-0001](error-catalog.md#snyk-cli-0001)

**환경 설정을 설정할 수 없음**

지정된 환경을 사용할 수 없습니다. 결과적으로 구성이 변경되지 않습니다. 환경에 대한 올바른 사양을 제공한 후 다시 시도하십시오.

**HTTP 상태:** [200](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200)

**도움링크:**

* [https://docs.snyk.io/snyk-cli/commands/config-environment](https://docs.snyk.io/snyk-cli/commands/config-environment)

#### [SNYK-CLI-0002](error-catalog.md#snyk-cli-0002)

**일관성 없는 구성일 가능성**

환경 변수 또는 구성 파일을 통해 CLI를 여러 방법으로 구성할 수 있습니다. 한 매개변수를 여러 번 구성할 경우 의도하지 않은 결과를 초래할 수 있습니다. 구성된 환경 변수를 검토하고 모든 것이 의도된 대로인지 확인하십시오. 그렇다면 `--no-check`를 사용하여이 체크를 건너뛸 수 있습니다.

**HTTP 상태:** [200](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200)

**도움링크:**

* [https://docs.snyk.io/snyk-cli/commands/config-environment](https://docs.snyk.io/snyk-cli/commands/config-environment)

#### [SNYK-OS-7001](error-catalog.md#snyk-os-7001)

**Snyk API로의 요청이 타임아웃**

Snyk API로의 요청이 예기치 않게 시간 초과되었습니다. Snyk 상태를 확인한 후 다시 시도하십시오.

**HTTP 상태:** [504](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/504)

**도움링크:**

* [https://status.snyk.io/](https://status.snyk.io/)

***

## 코드

#### [SNYK-CODE-0001](error-catalog.md#snyk-code-0001)

**분석 파일 개수 제한 초과**

분석 대상이 시스템 제한을 초과하는 지원되는 파일 수를 가지고 있을 때 이 오류가 발생합니다.

파일 수를 줄이려면 `.snyk` 파일을 사용하여 특정 디렉토리나 파일을 무시하십시오. 또는 Snyk CLI를 사용하여 개별 하위 디렉토리를 별도로 분석할 수 있습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#code-analysis-snyk-code](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#code-analysis-snyk-code)
* [https://docs.snyk.io/scan-applications/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process](https://docs.snyk.io/scan-applications/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process)
* [https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli](https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli)

#### [SNYK-CODE-0002](error-catalog.md#snyk-code-0002)

**분석 결과 크기 제한 초과**

분석 대상이 현재 시스템 제한을 초과하는 결과를 생성할 때 이 오류가 발생합니다.

전체 결과 크기를 줄이려면 `.snyk` 파일을 사용하여 특정 디렉토리나 파일을 무시하십시오. 또는 Snyk CLI를 사용하여 개별 하위 디렉토리를 별도로 분석할 수 있습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://docs.snyk.io/scan-applications/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process](https://docs.snyk.io/scan-applications/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process)
* \[https://docs.snyk.io/snyk-cli/[https://docs.snyk.io/scan-using-snyk/snyk-code/configure-snyk-code#enable-snyk-code-in-snyk-web-ui](https://docs.snyk.io/scan-using-snyk/snyk-code/configure-snyk-code#enable-snyk-code-in-snyk-web-ui)

#### [SNYK-CODE-0006](error-catalog.md#snyk-code-0006)

**프로젝트가 지원되지 않음**

Snyk은 지원되는 파일을 찾지 못했습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://docs.snyk.io/getting-started/supported-languages-frameworks-and-feature-availability-overview#code-analysis-snyk-code](https://docs.snyk.io/getting-started/supported-languages-frameworks-and-feature-availability-overview#code-analysis-snyk-code)

#### [SNYK-CODE-0007](error-catalog.md#snyk-code-0007)

**그룹에 대해 이미 존재하는 규칙 확장**

같은 유형과 속성을 갖는 규칙 확장이 이미 제공된 그룹에 존재합니다.

기존의 규칙 확장을 수정하거나 새로운 유형이나 속성을 갖는 규칙 확장을 생성하세요.

**HTTP 상태:** [409](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/409)

#### [SNYK-CODE-0008](error-catalog.md#snyk-code-0008)

**조직 관계는 고유해야 함**

각 Org 관계는 규칙 확장에 대해 고유해야 합니다.

관계에 있는 각 Org가 다른 ID를 가져야 합니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-CODE-0009](error-catalog.md#snyk-code-0009)

**그룹 관계는 요청된 URL의 그룹과 일치해야 함**

다른 그룹에 규칙 확장을 연결할 수 없습니다.

관계에서의 그룹 ID가 요청 경로에 있는 그룹 ID와 일치하는지 확인하세요.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-CODE-0010](error-catalog.md#snyk-code-0010)

**관리 그룹 외의 조직**

{Snyk Code 규칙 확장을 관리 그룹 외의 Org에 연결할 수 없습니다.

관계에 있는 각 Org가 요청된 URL의 그룹 내에 있는지 확인하세요.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-CODE-0011](error-catalog.md#snyk-code-0011)

**규칙 확장 제한 도달함**

그룹에 허용된 발행된 규칙 확장의 최대 개수에 도달했습니다.

새로운 규칙 확장을 만들려면 기존 규칙 확장을 제거해야 합니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-CODE-0012](error-catalog.md#snyk-code-0012)

**요청에서 지원되지 않는 규칙 이름**

하나 이상의 규칙 이름이 규칙 확장에서 지원되지 않습니다.

지원되지 않는 규칙 이름을 제거한 후 다시 시도하세요.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

***

## 통합

#### [SNYK-INTEGRATION-0001](error-catalog.md#snyk-integration-0001)

**SCM 통합을 찾을 수 없음**

SCM 통합이 존재하는지 확인하고 올바르게 설정되어 있는지 확인하세요.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**도움링크:**

* [https://docs.snyk.io/scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations](https://docs.snyk.io/scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations)

***

## OpenAPI

#### [SNYK-OPENAPI-0001](error-catalog.md#snyk-openapi-0001)

**잘못된 요청**

서버가 유효하지 않거나 손상된 데이터로 인해 요청을 처리할 수 없습니다. 요청을 검토한 후 다시 시도하세요.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**도움링크:**

* [https://docs.snyk.io/snyk-api-info/getting-started-using-snyk-rest-api](https://docs.snyk.io/snyk-api-info/getting-started-using-snyk-rest-api)

#### [SNYK-OPENAPI-0002](error-catalog.md#snyk-openapi-0002)

**금지됨**

요청한 리소스에 대한 액세스가 금지되었습니다. 요청을 검토한 후 다시 시도하세요.

**HTTP 상태:** [403](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403)

#### [SNYK-OPENAPI-0003](error-catalog.md#snyk-openapi-0003)

**허용되지 않음**

서버가 제공된 Accept 헤더와 일치하는 응답을 제공할 수 없습니다. 요청을 검토한 후 다시 시도하세요.

**HTTP 상태:** [406](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/406)

#### [SNYK-OPENAPI-0004](error-catalog.md#snyk-openapi-0004)

**찾을 수 없음**

서버가 요청한 리소스를 찾을 수 없습니다. 요청을 검토한 후 다시 시도하세요.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

#### [SNYK-OPENAPI-0005](error-catalog.md#snyk-openapi-0005)

**허용되지 않는 메소드**

대상 엔드포인트가 요청 메소드를 지원하지 않습니다. 요청을 검토한 후 다시 시도하세요.

**HTTP 상태:** [405](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405)

#### [SNYK-OPENAPI-0006](error-catalog.md#snyk-openapi-0006)

**요청 엔티티가 너무 큼**

요청 엔티티가 서버 제한을 초과했습니다. 요청 엔티티 크기를 줄이고 다시 시도하세요.

**HTTP 상태:** [413](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/413)

#### [SNYK-OPENAPI-0007](error-catalog.md#snyk-openapi-0007)

**권한 없음**

요청한 리소스에 대한 인증 자격 증명이 없습니다. 유효한 자격 증명을 보내고 다시 시도하세요.

**HTTP 상태:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

**도움링크:**

* [https://docs.snyk.io/snyk-api-info/authentication-for-api](https://docs.snyk.io/snyk-api-info/authentication-for-api)

#### [SNYK-OPENAPI-0008](error-catalog.md#snyk-openapi-0008)

**지원되지 않는 미디어 유형**

요청의 미디어 형식이 지원되지 않습니다. 미디어 형식을 변경한 후 다시 시도하세요.

**HTTP 상태:** \[415] (https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/415)

***

## 오픈 소스 언어 및 패키지 매니저

#### [SNYK-OS-0001](error-catalog.md#snyk-os-0001)

**manifest 파일을 구문 분석할 수 없음**

제공된 manifest 파일이 유효하지 않은 구문을 갖거나 예상된 스키마와 일치하지 않아 구문 분석할 수 없습니다. manifest 파일을 검토한 후 다시 시도하세요.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-OS-0002](error-catalog.md#snyk-os-0002)

**lock 파일을 파싱할 수 없음**

제공된 lock 파일이 유효하지 않은 구문을 갖거나 예상된 스키마와 일치하지 않아 파싱할 수 없습니다. lock 파일을 검토한 후 다시 시도하세요.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-OS-0003](error-catalog.md#snyk-os-0003)

**알 수 없는 종속성 버전**

종속성 버전을 해석할 수 없습니다.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**도움링크:**

* [https://support.snyk.io/s/article/Could-not-determine-version-for-dependencies](https://support.snyk.io/s/article/Could-not-determine-version-for-dependencies)

#### [SNYK-OS-0004](error-catalog.md#snyk-os-0004)

**필수 요청 헤더 누락**

서버에서 필수 요청 헤더가 누락된 요청을 수신했습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-0005](error-catalog.md#snyk-os-0005)

**필수 요소가 누락된 페이로드**

서버가 요청을 처리할 수 없습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-0006](error-catalog.md#snyk-os-0006)

**파일을 처리할 수 없음**

종속성 서비스가 파일을 처리할 수 없습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-0007](error-catalog.md#snyk-os-0007)

**원본에서 파일을 가져올 수 없음**

소스 URL에서 파일을 가져올 수 없습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

#### [SNYK-OS-0008](error-catalog.md#snyk-os-0008)

**환경 변수 누락**

서버가 특정 환경 변수가 필요한 중요 작업을 수행했지만, 해당 변수가 설정되지 않았거나 현재 환경 내에서 액세스할 수 없습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

#### [SNYK-OS-0009](error-catalog.md#snyk-os-0009)

**현재 지원되지 않는 브로커 연결**

서비스가 현재 그룹에서 해당 브로커 연결을 통해 가져 오려고 시도 할 때 권한이나 자격 증명 오류가 발생했습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

#### [SNYK-OS-0010](error-catalog.md#snyk-os-0010)

**Snyk이 저장소 복제를 실패함**

제공된 자격 증명을 사용하여 코드를 복제하려는 동안 Git에서 치명적인 오류가 발생했습니다. 다음을 확인하세요:

* 제공된 자격 증명이 정확하고 너무 제한적이지 않은지 확인하십시오.
* 복제할 브랜치가 있는지 확인하십시오.
* 제공한 저장소가 인터넷에서 접근 가능하며 브로커를 통해 연결되어 있지 않은지 확인하십시오.

그리고 작업을 다시 시도하세요.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-DOTNET-0001](error-catalog.md#snyk-os-dotnet-0001)

**치료를 위한 지원되지 않는 매니페스트 파일 유형**

제공된 매니페스트 파일이 .NET에 대해 Snyk에서 지원되지 않습니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**도움링크:**

* [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/.net](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/.net)

#### [SNYK-OS-DOTNET-0002](error-catalog.md#snyk-os-dotnet-0002)

**지원되지 않는 타겟 프레임워크**

제공된 매니페스트 파일이 현재 지원되지 않는 `<TargetFramework>` 또는 `<TargetFrameworks>`를 정의합니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-DOTNET-0003](error-catalog.md#snyk-os-dotnet-0003)

**C# 코드에 정적 Main 함수가 누락됨**

이 에러는 실행 가능 파일을 생성하는 코드에서 올바른 서명을 갖는 정적 Main 함수를 찾을 수 없을 때 발생합니다. 또는 올바른 대소문자로 정의되지 않은 경우, 'Main'이라는 진입점 함수가 경우와 같이 문제가 발생합니다.

이 문제를 해결하려면 다음 사항을 확인하세요: 코드에 main 함수가 포함되어 있는지 확인하세요. .cs 파일이 있어야 합니다. 아래와 같이 main 함수가 있는지 확인하십시오.

```c#
namespace Example
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("hello world");
        }
    }
}
```

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://learn.microsoft.com/en-us/dotnet/csharp/misc/cs5001](https://learn.microsoft.com/en-us/dotnet/csharp/misc/cs5001)

#### [SNYK-OS-DOTNET-0004](error-catalog.md#snyk-os-dotnet-0004)

**dotnet CLI가 자가 포함될 수 있는 바이너리 생성에 실패함**

`dotnet publish --sc --framework <your-target-framework>`을 실행하여 자체 포함 바이너리를 생성하는 데 실패할 때 발생하는 오류입니다. Snyk은프로젝트의 종속성 트리를 충분히 결정하기 위해 이 명령을 실행해야 합니다. 이 명령이 실패하면 Snyk은진행할 수 없습니다.

발생 이유를 확인하려면 다음 단계를 수행하세요:

* 임시 폴더에 프로젝트의 깨끗한 버전을 체크아웃합니다.
* 프로젝트에 대해 `dotnet publish --sc --framework <your-target-framework>` 를 실행하고 이 단계가 실패하는지 확인하세요.

로컬에서 이 단계가 성공하는 경우, Snyk은 다른 버전의 .NET SDK를 실행 중일 수 있습니다. Snyk에게 사용할 .NET SDK의 버전을 알려주기 위해 Microsoft가 제공하는 \[global설정.

**HTTP 상태:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

**도움링크:**

* [https://github.com/microsoft/artifacts-credprovider#environment-variables](https://github.com/microsoft/artifacts-credprovider#environment-variables)

#### [SNYK-OS-DOTNET-0006](error-catalog.md#snyk-os-dotnet-0006)

**프로젝트 파일에서 누락된 MSBuild Condition 구성 요소**

`dotnet` 도구는 프로젝트 파일에서 하나 이상의 MSBuild 조건에 대한 책임이 있는 `.targets`, `.csproj` 또는 `.props` 파일을 찾지 못했습니다.

도구는 다음과 같은 오류를 만났습니다.

```
/path/to/file/project.csproj(33,13): error MSB4100: Expected "$(SomeCondition)" to evaluate to a boolean instead of "", in condition "!$(SomeCondition)".
```

이는 현재 복원 중인 프로젝트 파일 및 해당 프로젝트와 연결된 모든 프로젝트에서 조건 정의가 누락되어 있음을 의미합니다.

Snyk은 현재 저장소에서 액세스할 수 있는 프로젝트 파일 또는 Snyk에서 사용 가능한 비공개 종속성만을 스캔할 수 있습니다.

예를 들어, 코드가 다음 구조를 가지는 경우:

```title=project.targets
<Project>
  <PropertyGroup>
    <SomeCondition Condition="'$(SomeCondition)' == ''">false</SomeCondition>
  </PropertyGroup>
</Project>
```

그리고

```title=project.csproj
<Project Sdk='Microsoft.NET.Sdk'>
  <Import Project='..\external-libraries\some-library\project.targets' />
  <PropertyGroup>
    <TargetFrameworks>net8.0</TargetFrameworks>
  </PropertyGroup>
  <ItemGroup Condition='!$(SomeCondition)'>
    <PackageReference Include='Newtonsoft.Json' Version='13.0.3' />
  </ItemGroup>
</Project>
```

그리고 `external-libraries`가 현재 스캔 중인 저장소의 일부가 아니라면, Snyk은 이를 찾을 수 없습니다.

이 오류는 Snyk에서 알 수 없는 외부 도구를 사용하여 코드에 추가되거나 생성된 외부 라이브러리에 의존하는 경우 빌드 스텝으로 발생합니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://learn.microsoft.com/en-us/visualstudio/msbuild/msbuild-conditional-constructs](https://learn.microsoft.com/en-us/visualstudio/msbuild/msbuild-conditional-constructs)

#### [SNYK-OS-DOTNET-0007](error-catalog.md#snyk-os-dotnet-0007)

**매니페스트 파일에서 대상 프레임워크를 찾을 수 없음**

Snyk은 제공된 매니페스트 파일에서 `<TargetFramework>`를 감지하지 못했습니다.

`Directory.Build.props` 파일을 사용하여 대상 프레임워크를 결정하고 있다면, 해당 파일의 이름이 해당하도록 확인하십시오. 고객 SCM 네트워크의 성능 고려로 인해 Snyk은`.props` 파일에 대해 대소문자 구분하지 않는 검색을 수행하지 않습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://learn.microsoft.com/en-us/visualstudio/msbuild/customize-by-directory?view=vs-2022#directorybuildprops-and-directorybuildtargets](https://learn.microsoft.com/en-us/visualstudio/msbuild/customize-by-directory?view=vs-2022#directorybuildprops-and-directorybuildtargets)

#### [SNYK-OS-DOTNET-0008](error-catalog.md#snyk-os-dotnet-0008)

**global.json이 구버전 SDK를 타겟팅함**

Snyk은 현재 [Microsoft에서 지원하는](https://dotnet.microsoft.com/en-us/download/dotnet) 최신 채널의 .NET을 지원하지만, 각 지원되는 채널 내의 모든 SDK 버전을 보장하지는 **않습니다.**

지원되는 채널 내에서 Snyk은 주로 최신의 채널 중에서 출시된 대부분 또는 모든 SDK 버전을 지원하려고 합니다.

Microsoft에서 현재 지원하는 채널이 `8.0`, `7.0`, `6.0`인 경우, Snyk은 이러한 채널을 위해 출시된 _최신_ SDK들을 **지원**할 것입니다.

`8.0.3` 하위의 SDK 버전이 `8.0.203`, `8.0.202`, `8.0.103`인 경우, Snyk은 _모든_ SDK 버전을 지원하는 것을 보장할 수는 없지만 최선을 다할 것입니다. Snyk은 Microsoft에서 현재 출시된 SDK 버전 중의 최신 것을 **지원**할 것입니다.

채널 `8.0`이 현재 지원되는 가장 최신 채널인 경우, Snyk은 여러 버전의 특정 SDK를 .NET 6과 같은 이전에도 지원되는 여러 채널에 대해 **지원**할 수 없습니다.

#### 지원 매트릭스 예시

다음의 예시에서:

* Microsoft에서 현재 지원하는 .NET 채널이 `.NET 8.0`, `.NET 7.0` 및 `.NET 6.0`이라고 가정합니다.
* `.NET 8.0`의 최신 SDK 버전은 `8.0.203`입니다.

그렇다면:

|  채널 |         SDK         | End-of-Life |  지원 |
| :-: | :-----------------: | :---------: | :-: |
| 8.0 | 8.0.203 (채널의 최신 버전) |      No     | Yes |
| 8.0 |       8.0.202       |      No     | Yes |
| 8.0 |       8.0.103       |      No     | Yes |
|     |        (...)        |             |     |
| 7.0 | 7.0.407 (채널의 최신 버전) |      No     | Yes |
| 7.0 |       7.0.314       |      No     |  No |
|     |        (...)        |             |     |
| 6.0 |       6.0.420       |      No     | Yes |
| 6.0 |       6.0.128       |      No     |  No |
|     |         (..)        |             |     |
| 5.0 | 5.0.408 (채널의 최신 버전) |     Yes     |  No |
| 5.0 |       5.0.214       |     Yes     |  No |
|     |         (..)        |             |     |

#### 해결책

이 제한으로 인해 `.NET` 채널을 핀다면, 새로운 지침([rollForward](https://learn.microsoft.com/en-us/dotnet/core/tools/global-json#rollforward) 지시문 포함) 없이 `global.json` 파일에 버전을 지정한 고객들에 대한 스캔 실패가 발생할 수 있습니다.

이 문제를 해결하기 위해, `global.json` 파일에서 `rollFoward` 지시문을 `latestMajor`로 설정하여 `6.0`이 최신 .NET 채널이 아니어도 최신 버전의 SDK를 사용하여 코드를 스캔할 수 있도록 권장합니다:

```json
{
  "sdk": {
    "version": "6.0.101",
    "rollForward": "latestMajor"
  }
}
```

이는 여러분의 버전 핀을 무시하고 더 높은 SDK 버전을 사용하여 코드를 스캔하도록 Snyk에 허용할 것입니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://versionsof.net/core/](https://versionsof.net/core/)
* [https://dotnet.microsoft.com/en-us/download/dotnet](https://dotnet.microsoft.com/en-us/download/dotnet)
* [https://learn.microsoft.com/en-us/dotnet/core/tools/global-json#rollforward](https://learn.microsoft.com/en-us/dotnet/core/tools/global-json#rollforward)

#### [SNYK-OS-DOTNET-0009](error-catalog.md#snyk-os-dotnet-0009)

**필수 타입 또는 네임스페이스 레퍼런스가 누락된 프로젝트 빌드 실패**

스캔을 위해 솔루션을 빌드하려고 시도할 때, `dotnet` SDK가 매니페스트 파일에서 참조된 하나 이상의 프로젝트를 복원하지 못했습니다.

Snyk은 **대소문자를 구분하는** 파일시스템에서 이러한 빌드를 실행하므로, `<ProjectReference>../src/NS.Project.csproj</ProjectReference>` 및 `<ProjectReference>../src/ns.project.csproj</ProjectReference>`는 동일한 것이 아님을 주의하십시오.

이는 파일 시스템이 대소문자를 구분하지 않는 Mac 또는 Windows 빌드 파이프라인을 사용하는 고객에게 문제가 발생할 수 있습니다. 이 경우, 올바른 매니페스트 파일을 참조하는지 확인하고 자세한 내용은 Snyk 가져오기 로그를 확인하십시오.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/compiler-messages/assembly-references#missing-references](https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/compiler-messages/assembly-references#missing-references)

#### [SNYK-OS-GO-0001](error-catalog.md#snyk-os-go-0001)

**개인 모듈에 액세스하지 못함**

Snyk은 당신의 go.mod 파일 내의 개인 모듈에 액세스할 수 없었습니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**도움링크:**

* [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/go](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/go)

#### [SNYK-OS-GO-0002](error-catalog.md#snyk-os-go-0002)

**go mod 파일을 찾을 수 없음**

현재 디렉토리나 부모 디렉터리에서 go.mod 파일을 찾을 수 없습니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**도움링크:**

* [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/go](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/go)

#### [SNYK-OS-GO-0003](error-catalog.md#snyk-os-go-0003)

**OAuth 재인증 필요**

Snyk에서 종속성을 분석하기 위해 코드가 Git을 사용하여 격리된 환경에 복제되어야 합니다.

조직이 Snyk의 코드 액세스를 허용한 후에 SAML SSO를 활성화하거나 강제로 실행했고, 따라서 다시 인증해야 합니다.

보통은 올바르지 않은 자격 증명으로 귀하의 저장소를 `git clone`하려고 시도하면 반복하여 나타납니다. Git 클라우드 공급업체와의 인증 구성을 확인하고 다시 시도하십시오.

{% hint style="warning" %}
**에러가 폐기되었습니다**

이유: 이 에러는 반복을 피하기 위해 더 일반적인 네임스페이스로 이동되었습니다.

앞으로, [SNYK-OS-8004](error-catalog.md#snyk-os-8004)가 대신 사용될 것입니다.
{% endhint %}

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://docs.github.com/en/enterprise-cloud@latest/authentication/authenticating-with-saml-single-sign-on/about-authentication-with-saml-single-sign-on#about-oauth-apps-github-apps-and-saml-sso](https://docs.github.com/en/enterprise-cloud@latest/authentication/authenticating-with-saml-single-sign-on/about-authentication-with-saml-single-sign-on#about-oauth-apps-github-apps-and-saml-sso)

#### [SNYK-OS-GO-0004](error-catalog.md#snyk-os-go-0004)

**프로젝트 저장소에 필요한 파일이 누락됨**

의존성 그래프 생성에는 프로젝트 내에서 `go list -deps -json`를 실행해야 합니다. 이 작업이 실패하면 전체 의존성 그래프를 계속 생성할 수 없습니다.

이 오류는 정리(`go mod tidy`)가 필요하거나 프로젝트 배포 프로세스에 `protobuf`와 같은 지금 Snyk에서 지원하지 않는 코드 생성 단계가 포함되어 있음을 의미합니다.

이것이 적합한 경우인지 확인하기 위해 프로젝트를 깨끗한 환경에서 복제하고, `go list -deps -json`를 실행하여 작업이 실패하는지 확인하십시오.

Snyk이 코드를 성공적으로 처리할 수 없는 경우, Snyk CLI를 배포 파이프라인의 일부로 추가하십시오.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://docs.snyk.io/snyk-cli](https://docs.snyk.io/snyk-cli)
* [https://github.com/snyk/snyk-go-plugin](https://github.com/snyk/snyk-go-plugin)
* [https://github.com/golang/go/blob/master/src/cmd/go/internal/list/list.go](https://github.com/golang/go/blob/master/src/cmd/go/internal/list/list.go)

#### [SNYK-OS-GO-0005](error-catalog.md#snyk-os-go-0005)

**프로젝트 저장소에 일관되지 않은 vendoring 정보가 있음**

의존성 그래프를 생성하려면 프로젝트 내에서 `go list -deps -json`을 실행해야 합니다. 이 작업이 실패하면 전체 의존성 그래프를 계속 생성할 수 없습니다.

이 오류는 `vendor/modules.txt` 파일과 `go.mod` 파일 사이에 일관성이 없는 경우입니다. 이 문제를 해결하려면:

* `go mod vendor`
* `go mod tidy`

다음으로 이러한 변경 사항을 귀하의 저장소에 커밋하십시오. Snyk은 우리 쪽에서 의도적으로 코드를 조작하지 않으므로 이 작업을 자동으로 수행하지 않습니다.

이것이 적합한 경우인지 확인하기 위해 프로젝트를 깨끗한 환경에서 복제하고, `go list -deps -json`를 실행하여 작업이 실패하는지 확인하고, 그런 다음 위에 언급된 명령을 실행하고 파일에서 변경 사항이 보고되[https://go.dev/ref/mod#vcs](https://go.dev/ref/mod#vcs)

#### [SNYK-OS-GO-0008](error-catalog.md#snyk-os-go-0008)

**개인 종속성을 가져올 수 없음**

Go 도구가 개인 종속성 중 하나를 가져 오는 동안 권한 오류가 발생했습니다. Snyk에 로그인 할 때 사용한 통합 토큰이 올바르게 구성되어 있는지 확인하여 Snyk이 개인 종속성에 액세스 할 수 있도록합니다.

Snyk Go 통합은 프로젝트 스캔 중인 프로젝트와 동일한 조직 내에서 사용되는 개인 종속성 만 지원합니다.

이 오류는 요청 된 개인 종속성의 권한 부여 자격 증명에 대한 Snyk의 액세스를 제대로 할 수 없을 때 나타납니다.

**HTTP 상태:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

#### [SNYK-OS-GO-0009](error-catalog.md#snyk-os-go-0009)

**도구 체인을 사용할 수 없음**

Go 도구 체인을 다운로드할 수 없습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

#### [SNYK-OS-MAVEN-0001](error-catalog.md#snyk-os-maven-0001)

**속성 누락**

pom 객체에서 필요한 속성이 누락되었습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-MAVEN-0002](error-catalog.md#snyk-os-maven-0002)

**속성의 값 해결 불가**

대상 속성에 유효한 값으로 해결할 수 없습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-MAVEN-0003](error-catalog.md#snyk-os-maven-0003)

**속성의 버전을 해결할 수 없음**

대상 속성에 유효한 버전으로 결정할 수 없습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

#### [SNYK-OS-MAVEN-0004](error-catalog.md#snyk-os-maven-0004)

**POM 파일에서 순환 속성 감지**

Maven 프로젝트의 구성 파일(POM)에서 속성 간에 순환 종속성이 있어 올바른 해결 방법을 방해하고 오류를 야기합니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-MAVEN-0005](error-catalog.md#snyk-os-maven-0005)

**XML 파일 구문 분석 오류**

XML 파일을 구문 분석하는 중 오류가 발생했습니다. 이는 pom.xml 또는 maven-metadata.xml 중 하나를 참조할 수 있습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-MAVEN-0006](error-catalog.md#snyk-os-maven-0006)

**제공된 좌표가 잘못됨**

프로젝트에 제공된 좌표가 잘못되었습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-MAVEN-0007](error-catalog.md#snyk-os-maven-0007)

**그룹 건너 뛰기**

다시 지정된 좌표로 인해 특정 groupId를 건너 뜁니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-MAVEN-0008](error-catalog.md#snyk-os-maven-0008)

**Pom 파일을 찾을 수 없음**

Maven 저장소에서 pom 파일을 찾을 수 없습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-MAVEN-0009](error-catalog.md#snyk-os-maven-0009)

**POM에서 프로젝트 누락**

POM에서 프로젝트 요소가 누락되었습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-MAVEN-0010](error-catalog.md#snyk-os-maven-0010)

**입력 XML에서 대상 POM을 해결할 수 없음**

입력 XML에서 대상 POM을 해결할 수 없습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-MAVEN-0011](error-catalog.md#snyk-os-maven-0011)

**리포지토리에서 대상 POM을 해결할 수 없음**

리포지토리에서 대상 POM을 해결할 수 없습니다.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

#### [SNYK-OS-MAVEN-0012](error-catalog.md#snyk-os-maven-0012)

**빌드 파일 리포지토리를 가져올 수 없음**

빌드 파일 리포지토리를 가져올 수 없습니다.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

#### [SNYK-OS-MAVEN-0013](error-catalog.md#snyk-os-maven-0013)

**호스팅 된 git 정보를 만들 수 없음**

소스 URL을 생성할 수 없습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

#### [SNYK-OS-MAVEN-0014](error-catalog.md#snyk-os-maven-0014)

**버전 범위에 해당하는 릴리스된 버전이 없음**

지정된 버전 범위에 릴리스된 버전이 없습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-MAVEN-0015](error-catalog.md#snyk-os-maven-0015)

**지원되지 않는 소스**

사용된 소스가 fetcher에서 지원되지 않습니다. 지원되는 소스는 다음과 같습니다: github, bitbucket, gitlab.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-MAVEN-0016](error-catalog.md#snyk-os-maven-0016)

**종속성 트리 처리 중에 타임아웃 발생**

종속성 트리를 처리하는 동안 타임아웃이 발생했습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-MAVEN-0017](error-catalog.md#snyk-os-maven-0017)

**Snyk 조직 언어 설정에서 구성된 하나 이상의 Maven 리포지토리에 도달할 수 없음**

조직 언어 설정에 구성된 하나 이상의 Maven 리포지토리에 도달할 수 없습니다.

이 오류는 다음과 같은 이유로 발생할 수 있습니다:

* 브로커를 사용하는 경우 브로커 클라이언트에서 구성 오류가 있을 수 있습니다. 사용자 이름과 암호를 다시 확인하세요.
* 브로커 클라이언트와 Snyk 또는 브로커 클라이언트와 구성된 리포지토리 사이의 네트워크 연결성이 문제가 있을 수 있습니다. 방화벽 규칙을 확인하세요.

이 문제를 해결하려면이 오류 메시지의 구체적인 세부 정보를 참조하여 어떤 리포지토리가 문제를 일으키는지 확인하십시오.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**도움링크:**

* [https://docs.snyk.io/integrate-with-snyk/package-repository-integrations](https://docs.snyk.io/integrate-with-snyk/package-repository-integrations)

#### [SNYK-OS-NODEJS-0001](error-catalog.md#snyk-os-nodejs-0001)

**NPM 패키지에 대한 저장소를 찾을 수 없음**

NPM 패키지를 찾을 수 없습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-NODEJS-0002](error-catalog.md#snyk-os-nodejs-0002)

**NPM 레지스트리 URL을 구문 분석할 수 없음**

NPM 레지스트리 URL을 구문 분석할 수 없습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-NODEJS-0003](error-catalog.md#snyk-os-nodejs-0003)

**브로커가 해결한 URL을 찾을 수 없음**

브로커가 해결한 URL을 찾을 수 없습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-NODEJS-0004](error-catalog.md#snyk-os-nodejs-0004)

**브로커 URL을 대체할 수 없음**

잠금 파일의 모든 브로커 URL을 대체할 수 없습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-NODEJS-0005](error-catalog.md#snyk-os-nodejs-0005)

**잘못된 NPM 버전**

지원되지 않는 NPM 버전입니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-NODEJS-0006](error-catalog.md#snyk-os-nodejs-0006)

**Github에서 알 수없는 blob 인코딩**

Github에서 알 수없는 blob 인코딩입니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-NODEJS-0007](error-catalog.md#snyk-os-nodejs-0007)

**포크된 프로세스에서 결과가 없음**

포크된 프로세스에서 결과가 없습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

#### [SNYK-OS-NODEJS-0008](error-catalog.md#snyk-os-nodejs-0008)

**자식 프로세스 실행 오류**

자식 프로세스가 실행 중에 오류가 발생했습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

#### [SNYK-OS-NODEJS-0009](error-catalog.md#snyk-os-nodejs-0009)

**유효한 패키지 업데이트 없음**

시스템은 잠금 파일에 지정된 패키지에 대한 유효한 업데이트를 찾으려고 시도했지만 사용 가능한 것이 없습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-NODEJS-0010](error-catalog.md#snyk-os-nodejs-0010)

**종속성 업데이트 없음**

의존성에 대한 사용 가능한 업데이트가 없습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-NODEJS-0011](error-catalog.md#snyk-os-nodejs-0011)

**JSON 파일을 구문 분석할 수 없음**

JSON 파일을 구문 분석하는 동안 오류가 발생했습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-NODEJS-0012](error-catalog.md#snyk-os-nodejs-0012)

**Base64 인코딩할 수 없음**

Base64 인코딩을 시도하는 동안 오류가 발생했습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-NODEJS-0013](error-catalog.md#snyk-os-nodejs-0013)

**Base64 디코딩할 수 없음**

Base64 디코딩을 시도하는 동안 오류가 발생했습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-NODEJS-0014](error-catalog.md#snyk-os-nodejs-0014)

**지원되는 파일 없음**

지원되는 파일을 찾을 수 없습니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-OS-NODEJS-0015](error-catalog.md#snyk-os-nodejs-0015)

**잘못된 구성**

구성 매개 변수가 예상되는 데이터 유형을 충족하지 않습니다. 제공된 값이 올바른 데이터 유형인지 확인하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

#### [SNYK-OS-NODEJS-0016](error-catalog.md#snyk-os-nodejs-0016)

**동기화 오류**

프로젝트가 종속성 파일과 매니페스트 파일 간에 동기화되지 않을 수 있습니다. 이는 package.json이 수정되거나 업데이트되었지만 pnpm-lock.yaml은 그렇지 않은 경우 발생할 수 있습니다.

패키지 파일과 매니페스트 파일이 올바르게 동기화되도록 하려면 pnpm install을 실행하십시오.

경우에 따라 node\_modules 폴더를 삭제하고 pnpm-lock.yaml을 삭제한 후 다시 pnpm install을 실행하여 완전한 재설치를 강제해야 할 수 있습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://support.snyk.io/s/article/Out-of-sync-manifest--lockfile-in-the-project](https://support.snyk.io/s/article/Out-of-sync-manifest--lockfile-in-the-project)

#### [SNYK-OS-NODEJS-0017](error-catalog.md#snyk-os-nodejs-0017)

**지원되지 않는 pnpm lockfile 버전**

lockfile 버전이 지원되지 않습니다. pnpm에 대한 지원되는 lockfile 버전은 v5 및 v6입니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-OS-NODEJS-0019](error-catalog.md#snyk-os-nodejs-0019)

**Yarn 패키지를 찾을 수 없음**

Snyk이 Yarn 레지스트리에서 패키지를 찾을 수 없습니다.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

#### [SNYK-OS-NODEJS-0020](error-catalog.md#snyk-os-nodejs-0020)

**패키지 레지스트리에 도달할 수 없음**

Snyk이 노드 패키지 레지스트리에 도달할 수 없습니다.

**HTTP 상태:** [503](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503)

#### \[SNYK-OS-N## [SNYK-OS-PYTHON-0005](error-catalog.md#snyk-os-python-0005)

#### Manifest 파일에서 구문 오류 발견

Manifest 파일에는 패키지 이름이 잘못되거나 지원되지 않는 문자와 같은 구문 오류가 있습니다. Manifest 파일이 구문 규칙을 따르고 로컬로 설치할 수 있는지 확인하십시오.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0006](error-catalog.md#snyk-os-python-0006)

#### 지원되지 않는 Python 버전

패키지 중 하나 이상이 프로젝트 스캔에서 사용된 Python 버전과 일치하지 않는 Python 버전을 요구합니다. 적절한 Python 버전을 조직 Python 언어 설정에서 선택하십시오. 또는 Python 버전 선택 오버라이드를 위해 `.snyk` 파일을 추가하십시오.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0007](error-catalog.md#snyk-os-python-0007)

#### 패키지 버전 간 충돌

두 개 이상의 패키지가 해결할 수 없는 버전 요구 사항으로 인해 충돌이 발생했습니다. 두 패키지와 그들의 요구 사항이 충돌을 일으키지 않도록하고, Manifest 파일이 로컬로 설치될 수 있는지 확인하십시오.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0008](error-catalog.md#snyk-os-python-0008)

#### 패키지 중 하나 이상의 일치하는 배포본을 찾을 수 없음

패키지 중 하나 이상이 프로젝트 스캔에서 사용된 Python 버전과 일치하지 않는 Python 버전을 요구합니다. 적절한 Python 버전을 조직 Python 언어 설정에서 선택하십시오. 또는 Python 버전 선택 오버라이드를 위해 `.snyk` 파일을 추가하십시오.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0009](error-catalog.md#snyk-os-python-0009)

#### 패키지 설치 실패

일부 패키지는 시스템 종속성이 누락되었거나 컴파일 오류 또는 기타 패키지 특정 문제로 인해 설치 중 실패했습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0010](error-catalog.md#snyk-os-python-0010)

#### 지원되지 않는 Python 버전

패키지 중 하나 이상이 프로젝트 스캔에서 사용된 Python 버전과 일치하지 않는 Python 버전을 요구합니다. Pipfile의 requires 섹션에서 올바른 python 버전을 사용하도록 하십시오.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0011](error-catalog.md#snyk-os-python-0011)

#### 패키지 중 하나 이상의 일치하는 배포본을 찾을 수 없음

패키지 중 하나 이상이 프로젝트 스캔에서 사용된 Python 버전과 일치하지 않는 Python 버전을 요구합니다. Pipfile의 requires 섹션에서 올바른 python 버전을 사용하도록 하십시오.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

### [SNYK-OS-PYTHON-0012](error-catalog.md#snyk-os-python-0012)

#### 다운로드된 Python 종속성의 10GB 용량 제한 초과

Manifes 파일의 다운로드된 Python 종속성의 총 크기가 10GB 제한을 초과했습니다. 이는 PyTorch와 같은 NVIDIA 드라이버를 필요로 하는 Python 패키지들의 크기가 크기 때문에 종종 발생합니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://docs.snyk.io/supported-languages-package-managers-and-frameworks/python/git-repositories-and-python#pip-and-git-repositories](https://docs.snyk.io/supported-languages-package-managers-and-frameworks/python/git-repositories-and-python#pip-and-git-repositories)

***

## 빌드

### [SNYK-OS-8001](error-catalog.md#snyk-os-8001)

#### 잘못된 요청

선택한 생태계에 대한 제공된 요청 payload가 유효하지 않습니다. API 문서를 검토하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**도움링크:**

* [https://apidocs.snyk.io/](https://apidocs.snyk.io/)

### [SNYK-OS-8002](error-catalog.md#snyk-os-8002)

#### 빌드 환경을 찾을 수 없음

제공된 컨텍스트에 대한 빌드 환경을 찾을 수 없습니다. 먼저 빌드 환경을 생성했는지 확인하십시오.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-OS-8003](error-catalog.md#snyk-os-8003)

#### 지원되지 않는 생태계

언어나 패키지 관리자가 지원되지 않습니다. 지원되는 패키지 관리자를 확인하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**도움링크:**

* [https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#open-source-and-licensing-snyk-open-source](https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#open-source-and-licensing-snyk-open-source)

### [SNYK-OS-8004](error-catalog.md#snyk-os-8004)

#### OAuth 재인증 필요

Snyk이 종속성을 분석하기 위해 Git을 사용하여 저장소에서 코드를 복제합니다.

조직이 코드 접근 권한을 Snyk에 부여한 후 SAML SSO를 활성화하거나 강제한 경우 재인증이 필요합니다.

보통 잘못된 구성 자격 증명으로 리포지토리를 `git clone`하여 발생하는 오류입니다. Git 클라우드 제공업체와의 인증 구성을 확인하고 다시 시도하십시오.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://docs.github.com/en/enterprise-cloud@latest/authentication/authenticating-with-saml-single-sign-on/about-authentication-with-saml-single-sign-on#about-oauth-apps-github-apps-and-saml-sso](https://docs.github.com/en/enterprise-cloud@latest/authentication/authenticating-with-saml-single-sign-on/about-authentication-with-saml-single-sign-on#about-oauth-apps-github-apps-and-saml-sso)

***

## SBOM Export

### [SNYK-OS-9000](error-catalog.md#snyk-os-9000)

#### SBOM 생성 내보내기 서버 오류

SBOM 생성 중 예기치 않은 오류가 발생했습니다. 요청을 검토한 후 다시 시도하십시오. 오류가 지속되면 Snyk 지원팀에 문의하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OS-9001](error-catalog.md#snyk-os-9001)

#### 의존성 그래프 오류

의도치 않은 의존성 그래프 오류가 발생했습니다. 요청을 검토한 후 다시 시도하십시오. 오류가 지속되면 Snyk 지원팀에 문의하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OS-9002](error-catalog.md#snyk-os-9002)

#### 의존성 그래프 구문 분석 오류

예기치 않은 오류로 인해 의존성 그래프를 구문 분석할 수 없습니다. 요청을 검토한 후 다시 시도하십시오. 오류가 지속되면 Snyk 지원팀에 문의하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OS-9003](error-catalog.md#snyk-os-9003)

#### 프로젝트 유형 때문에 SBOM이 지원되지 않음

Snyk Open Source 또는 Snyk Container 프로젝트에 대해서만 SBOM을 지원합니다.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-OS-9004](error-catalog.md#snyk-os-9004)

#### SBOM이 지원되지 않음

온리 오픈 소스 프로젝트(Snyk Open Source)에 대해서만 SBOM을 지원합니다.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-OS-9005](error-catalog.md#snyk-os-9005)

#### 의존성 그래프 요청을 처리할 수 없음

서버가 완전한 데이터 부족으로 인해 요청을 처리할 수 없습니다. 요청을 검토하고 다시 시도하십시오.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

### [SNYK-OS-9006](error-catalog.md#snyk-os-9006)

#### API 토큰 누락으로 인한 인가 실패

API 토큰이 잘못 구성되었거나 만료되었습니다. API 토큰을 구성 또는 생성한 후 다시 시도하십시오.

**HTTP 상태:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

**도움링크:**

* [https://docs.snyk.io/snyk-api-info/revoking-and-regenerating-snyk-api-tokens](https://docs.snyk.io/snyk-api-info/revoking-and-regenerating-snyk-api-tokens)

### [SNYK-OS-9007](error-catalog.md#snyk-os-9007)

#### 클라이언트 요청을 처리할 수 없음

요청의 본문이 비어 있습니다. 요청을 검토한 후 다시 시도하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OS-9008](error-catalog.md#snyk-os-9008)

#### 잘못된 의존성 그래프

제공된 의존성 그래프가 유효하지 않습니다. 요청을 검토한 후 다시 시도하십시오.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

***

## 오픈 소스 관리되지 않음

### [SNYK-OSJVM-001](error-catalog.md#snyk-osjvm-001)

#### Maven 검색 서비스를 사용할 수 없음

상위 Maven 검색 서비스가 사용 중지되었습니다.

**HTTP 상태:** [503](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503)

**도움링크:**

* [https://search.maven.org](https://search.maven.org)
* [https://status.maven.org](https://status.maven.org)

### [SNYK-OSJVM-002](error-catalog.md#snyk-osjvm-002)

#### SHA1을 찾을 수 없음

제공된 SHA1에 대한 좌표를 찾을 수 없습니다. 보내는 데이터를 확인하고 다시 시도하십시오.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**도움링크:**

* [https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files](https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files)

***

## PURL 취약점

### [SNYK-OSSI-1040](error-catalog.md#snyk-ossi-1040)

#### 귀하의 기관은 이 조치를 실행할 권한이 없습니다

베타 기능에 대한 액세스 권한이 없는 것으로 보입니다. 액세스하려면 계정 담당자 또는 팀을 통해 베타 기능 액세스를 요청할 수 있습니다.

**HTTP 상태:** [403](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403)

### [SNYK-OSSI-1050](error-catalog.md#snyk-ossi-1050)

#### 권한 요청 실패

인증 중 예기치 않은 오류가 발생했습니다. 다시 시도하고 오류가 계속되면 지원팀에 문의하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-2010](error-catalog.md#snyk-ossi-2010)

#### 유효하지 않은 purl

purl이 유효한지 확인하십시오. 자세한 정보는 패키지 URL 사양 링크를 참조하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**도움링크:**

* [https://github.com/package-url/purl-spec/blob/master/PURL-SPECIFICATION.rst](https://github.com/package-url/purl-spec/blob/master/PURL-SPECIFICATION.rst)

### [SNYK-OSSI-2011](error-catalog.md#snyk-ossi-2011)

#### 네임 스페이스가 지정되지 않았습니다

네임 스페이스를 제공해야 하는 패키지 유형을 요청했습니다(예: 메이븐 그룹 ID). 패키지를 검색하려면 네임 스페이스를 제공하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**도움링크:**

* [https://github.com/package-url/purl-spec/blob/master/PURL-SPECIFICATION.rst](https://github.com/package-url/purl-spec/blob/master/PURL-SPECIFICATION.rst)

### [SNYK-OSSI-2020](error-catalog.md#snyk-ossi-2020)

#### 지원되지 않는 생태계

패키지 유형이 지원되지 않습니다. Snyk API의 패키지에 대한 목록 문제를 확인하십시오.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-2021](error-catalog.md#snyk-ossi-2021)

#### Purl 구성 요소가 필요함

purl 사양의 구성 요소 목록이 필요합니다. purl이 필요한 모든 구성 요소를 지정하지 않았습니다. 누## K-OSSI-2041

#### 유효하지 않은 페이지네이션 매개변수

페이지네이션 제한은 > 1이고 ≤ 1000이어야 하며, 오프셋은 ≥0이어야 합니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-2042](error-catalog.md#snyk-ossi-2042)

#### purls 제한 초과

요청에 포함된 purl 수가 서비스에서 설정한 1000의 제한을 초과합니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-2043](error-catalog.md#snyk-ossi-2043)

#### 문제 개수가 제한을 초과함

제공된 purls에 대해 발견된 문제의 수가 API에서 정의한 제한을 초과합니다. 단일 요청에서 보내는 purl 수를 줄이세요.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-2044](error-catalog.md#snyk-ossi-2044)

#### 예상 배포판이 존재해야 함

제공된 패키지 URL에 필요한 distro 한정자가 없습니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

**도움링크:**

* [https://docs.snyk.io/scan-containers/how-snyk-container-works/supported-operating-system-distributions#debian](https://docs.snyk.io/scan-containers/how-snyk-container-works/supported-operating-system-distributions#debian)

### [SNYK-OSSI-2045](error-catalog.md#snyk-ossi-2045)

#### 지원되지 않는 Debian 배포판

현재이 Debian 배포판은 지원되지 않습니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-2046](error-catalog.md#snyk-ossi-2046)

#### 예상 네임스페이스가 있어야 함

제공된 패키지 URL에 필요한 네임스페이스가 없습니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-2047](error-catalog.md#snyk-ossi-2047)

#### 지원되지 않는 공급업체

제공된 패키지 URL에 지원되는 공급업체가 포함되어 있지 않습니다. 나열된 공급업체 중 하나를 사용하고 다시 시도하세요.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-2048](error-catalog.md#snyk-ossi-2048)

#### 지원되지 않는 Alpine 배포판

현재이 Alpine 배포판은 지원되지 않습니다.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

***

## 오픈 소스 프로젝트 문제

### [SNYK-OSSI-OSPI-1001](error-catalog.md#snyk-ossi-ospi-1001)

#### 유효하지 않은 요청

요청의 본문을 확인하고 다시 시도하세요.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-OSPI-1002](error-catalog.md#snyk-ossi-ospi-1002)

#### 유효한 API 응답을 반환할 수 없음

이 문제는 예기치 않은 것이며, 서비스는 곧 회복될 것입니다. 오류가 지속되면 지원팀에 문의하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-OSPI-2001](error-catalog.md#snyk-ossi-ospi-2001)

#### 데이터 처리 실패

이 문제는 예기치 않은 것이며, 서비스는 곧 회복될 것입니다. 오류가 지속되면 지원팀에 문의하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-OSPI-3001](error-catalog.md#snyk-ossi-ospi-3001)

#### 문제 데이터 저장 실패

입력을 확인한 다음 다시 시도하세요. 오류가 지속되면 지원팀에 문의하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-OSPI-4001](error-catalog.md#snyk-ossi-ospi-4001)

#### 내부 서버 오류

이 문제는 예기치 않은 것이며, 서비스는 곧 회복될 것입니다. 오류가 지속되면 지원팀에 문의하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

***

## 오픈 소스 프로젝트 스냅샷

### [SNYK-OSSI-OSPSS-1001](error-catalog.md#snyk-ossi-ospss-1001)

#### 유효하지 않은 요청

요청의 본문을 확인하고 다시 시도하세요.

**HTTP 상태:** [400](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400)

### [SNYK-OSSI-OSPSS-1002](error-catalog.md#snyk-ossi-ospss-1002)

#### 유효한 API 응답을 반환할 수 없음

이 문제는 예기치 않은 것이며, 서비스는 곧 회복될 것입니다. 오류가 지속되면 지원팀에 문의하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-OSPSS-2001](error-catalog.md#snyk-ossi-ospss-2001)

#### 데이터 처리 실패

이 문제는 예기치 않은 것이며, 서비스는 곧 회복될 것입니다. 오류가 지속되면 지원팀에 문의하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-OSPSS-3001](error-catalog.md#snyk-ossi-ospss-3001)

#### 스냅샷 데이터 저장 실패

입력을 확인한 다음 다시 시도하세요. 오류가 지속되면 지원팀에 문의하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

### [SNYK-OSSI-OSPSS-4001](error-catalog.md#snyk-ossi-ospss-4001)

#### 내부 서버 오류

이 문제는 예기치 않은 것이며, 서비스는 곧 회복될 것입니다. 오류가 지속되면 지원팀에 문의하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

***

## 정책

### [SNYK-POLICY-0001](error-catalog.md#snyk-policy-0001)

#### 잘못된 구성을 가진 정책 적용할 수 없음

Snyk은 정책의 구성이 잘못되었기 때문에 테스트 실행 중에 정책을 적용할 수 없었습니다. 정책을 수정하고 다시 시도할 수 있습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://docs.snyk.io/manage-risk/policies](https://docs.snyk.io/manage-risk/policies)

***

## PRChecks

### [SNYK-PR-CHECK-0001](error-catalog.md#snyk-pr-check-0001)

#### 매니페스트 읽기 오류

Snyk이 1개 이상의 매니페스트 파일을 읽는 데 실패했습니다. 때로는 문제가 발생할 수 있습니다: 신뢰할 수없는 연결, 제3자 서비스 다운 및 Snyk이 프로젝트를 테스트하기 위해 필요한 파일을 읽을 수 없습니다.

이런 경우 다음을 시도할 수 있습니다:

* Pull Request / Merge Request를 열고 다시 열어 새로운 테스트를 시작합니다.
* 저장소를 제거하고 Snyk에 다시 추가합니다.

최종적으로 문제가 지속되면 지원팀에 문의해야 합니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**도움링크:**

* [https://support.snyk.io/s/article/Failed-to-read-manifest-file---Commit-Status](https://support.snyk.io/s/article/Failed-to-read-manifest-file---Commit-Status)

(이하 생략)### SBOM 테스트 오류

예기치 않은 오류가 발생했습니다. 요청을 확인한 후 다시 시도하십시오. 오류가 지속되면 Snyk 지원팀에 문의하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

#### [SNYK-SBOM-0002](error-catalog.md#snyk-sbom-0002)

**조직 ID 불일치**

요청된 조직 ID가 SBOM 테스트 ID의 소유자와 일치하지 않습니다.

이 오류는 제공된 조직 ID가 SBOM 테스트 실행 시 사용된 것과 다른 경우 발생합니다. 요청에 사용된 조직 ID가 SBOM 테스트 생성 시 사용된 조직 ID와 동일한지 확인하십시오.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

#### [SNYK-SBOM-0003](error-catalog.md#snyk-sbom-0003)

**SBOM 테스트를 찾을 수 없음**

Snyk이 요청된 SBOM 테스트를 찾을 수 없습니다.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

#### [SNYK-SBOM-0004](error-catalog.md#snyk-sbom-0004)

**SBOM 테스트 실패**

SBOM 테스트가 실패하였으며 결과가 없습니다.

이 오류는 실패한 SBOM 테스트의 결과를 요청할 때 발생합니다.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

#### [SNYK-SBOM-0005](error-catalog.md#snyk-sbom-0005)

**SBOM 테스트 결과 아직 처리 중**

SBOM 테스트가 아직 처리 중이며 결과가 없습니다.

이 오류는 SBOM 테스트 결과를 요청했지만 SBOM 테스트가 아직 처리 중인 경우 발생합니다.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

#### [SNYK-SBOM-0006](error-catalog.md#snyk-sbom-0006)

**알 수 없는 SBOM 형식**

Snyk은 SBOM 파일 형식을 인식하지 못합니다.

지원되는 형식의 SBOM 문서를 제공하십시오.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://docs.snyk.io/snyk-cli/commands/sbom-test](https://docs.snyk.io/snyk-cli/commands/sbom-test)

#### [SNYK-SBOM-0007](error-catalog.md#snyk-sbom-0007)

**SBOM 입력 처리할 수 없음**

Snyk이 SBOM 파일을 해독할 수 없습니다. 유효한 SBOM 문서를 제공하고 다시 시도하십시오.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://docs.snyk.io/snyk-cli/commands/sbom-test](https://docs.snyk.io/snyk-cli/commands/sbom-test)

#### [SNYK-SBOM-0008](error-catalog.md#snyk-sbom-0008)

**지원되지 않는 SBOM 형식**

지원되는 형식의 SBOM 문서를 제공하고 다시 시도하십시오.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://docs.snyk.io/snyk-cli/commands/sbom-test](https://docs.snyk.io/snyk-cli/commands/sbom-test)

#### [SNYK-SBOM-0009](error-catalog.md#snyk-sbom-0009)

**SBOM 분석 실패**

Snyk이 제공된 SBOM 입력을 처리하지 못하고 문제를 검사할 수 없었습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-SBOM-0010](error-catalog.md#snyk-sbom-0010)

**검사 가능한 패키지를 찾을 수 없음**

제공된 SBOM 문서에 Snyk 취약성 분석에서 지원되는 패키지가 포함되지 않았거나 패키지가 없습니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

***

## SCM

#### [SNYK-SCM-0001](error-catalog.md#snyk-scm-0001)

**통합 유형이 지원되지 않음**

제공된 통합은 SCM 저장소 액세스를 지원하지 않습니다.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

**도움링크:**

* [https://docs.snyk.io/scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations](https://docs.snyk.io/scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations)

#### [SNYK-SCM-0002](error-catalog.md#snyk-scm-0002)

**리비전을 해결할 수 없음**

Snyk이 제공한 SCM 리비전을 해결할 수 없었습니다. 유효한 리비전을 제공하십시오. 전체 커밋 ID 또는 기존 커밋 참조 중 하나여야 합니다.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

#### [SNYK-SCM-0003](error-catalog.md#snyk-scm-0003)

**통합 인증 실패**

Snyk이 SCM 공급업체와 인증하지 못했습니다. Snyk 통합에 대한 유효한 자격 증명을 사용하고 있는지 확인하십시오.

**HTTP 상태:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

#### [SNYK-SCM-0004](error-catalog.md#snyk-scm-0004)

**통합 승인 실패**

Snyk이 SCM 공급업체와 승인하지 못했습니다. 조직에 SAML SSO가 활성화되어 있거나 강제로 사용 중인 경우 OAuth 애플리케이션 `Snyk` 다시 승인하십시오.

**HTTP 상태:** [401](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401)

#### [SNYK-SCM-0005](error-catalog.md#snyk-scm-0005)

**파일이 너무 많음**

총 파일 수가 40000을 초과하여 저장소를 검색할 수 없었습니다.

파일 수를 줄이려면 특정 디렉토리나 파일을 무시하도록 `.snyk` 파일을 사용하십시오. 또는 개별 작업 서브디렉터리를 별도로 분석하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

#### [SNYK-SCM-0006](error-catalog.md#snyk-scm-0006)

**저장소 크기가 너무 큼**

저장소의 크기가 15 GB를 초과하여 저장소를 검색할 수 없었습니다.

전체 저장소 크기를 줄이기 위해 특정 디렉토리나 파일을 무시하도록 `.snyk` 파일을 사용하십시오. 또는 개별 작업 서브디렉터리를 별도로 분석하십시오.

**HTTP 상태:** [500](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500)

***

## Target

#### [SNYK-TARGET-0001](error-catalog.md#snyk-target-0001)

**대상을 찾을 수 없음**

Snyk이 가져온 대상을 해결할 수 없었습니다. Snyk이 대상을 만들었는지 확인하고 다시 시도하십시오.

**HTTP 상태:** [404](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404)

**도움링크:**

* [https://docs.snyk.io/snyk-admin/snyk-projects#target](https://docs.snyk.io/snyk-admin/snyk-projects#target)

#### [SNYK-TARGET-0002](error-catalog.md#snyk-target-0002)

**고유 대상을 찾을 수 없음**

Snyk이 단일 대상을 해결할 수 없었습니다. Snyk이 동일한 통합 및 저장소 URL 쌍에 대해 구성된 여러 대상을 찾았습니다. 고유 대상이 존재하는지 확인하십시오.

**HTTP 상태:** [422](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422)

**도움링크:**

* [https://docs.snyk.io/snyk-admin/snyk-projects#target](https://docs.snyk.io/snyk-admin/snyk-projects#target)

\--- 생성 일시: 2024년 12월 16일 09시 51분 05.835초 (협정 세계시)
