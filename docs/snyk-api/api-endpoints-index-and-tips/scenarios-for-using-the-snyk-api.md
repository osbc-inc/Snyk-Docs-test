# Snyk API 사용 시나리오

Snyk API 시나리오는 Snyk 애플리케이션을 API를 사용하여 작업을 수행하는 데 사용할 수 있는 프로시저를 식별합니다. 이 페이지에 나열된 시나리오는 Snyk 프로세스에 그룹화되어 있으며 [저장소](https://github.com/snyk-playground/cx-tools/tree/main/scripts)나 사용자 문서 사이트에서 제공됩니다. 링크가 포함되어 있습니다.

이러한 프로시저를 사용할 때 문제가 발생하면 기술 지원 담당자 또는 솔루션 엔지니어에게 문의하거나 [특정 티켓을 제출](https://support.snyk.io)하여 Snyk 지원팀에 문의하십시오.

## Snyk 조직 구조 관리

### 특정 그룹에서 모두 동일한 설정을 가진 여러 개의 새 조직 생성

시나리오: [create-multiple-orgs-and-copy-settings](https://github.com/snyk-playground/cx-tools/blob/main/scripts/create-multiple-orgs-and-copy-settings.md) (전체 프로시저)

**사용된 Endpoints:**\
[새 조직 만들기](../reference/organizations-v1.md#org)\
[조직 설정 보기](../reference/organizations-v1.md#org-orgid-settings-1)\
[조직 설정 업데이트](../reference/organizations-v1.md#org-orgid-settings)\
[설정과 자격 증명이 포함된 통합 복제](../reference/integrations-v1.md#org-orgid-integrations-integrationid-clone)

### 회사가 보유한 모든 조직에 대한 특정 목록의 사용자를 모든 조직에 할당

시나리오: [assign-users-to-all-orgs](https://github.com/snyk-playground/cx-tools/blob/main/scripts/assign-users-to-all-orgs.md) (전체 프로시저)

**사용된 Endpoints:**\
[그룹의 모든 회원 목록](../reference/groups-v1.md#group-groupid-members)\
[사용자 초대](../reference/organizations-v1.md#org-orgid-invite)

### 첫 로그인 전에 대규모로 사용자를 조직에 추가

시나리오: [API를 통한 조직에 사용자 할당](../../snyk-admin/user-management-with-the-api/provision-users-to-organizations-using-the-api.md)

**사용된 Endpoint:**\
[사용자를 조직에 할당](../reference/organizations-v1.md#org-orgid-provision)

## Snyk 프로젝트 가져오고 설정하기

### 식별 및 새 저장소만 가져오기

시나리오: [Identify-and-import-new-repos](https://github.com/snyk-playground/cx-tools/blob/main/scripts/Identify-and-import-new-repos.md) (전체 프로시저)

**사용된 Endpoints:**\
[조직 ID에 따른 대상 가져오기](../reference/targets.md#orgs-org_id-targets)\
[대상 가져오기](../reference/import-projects-v1.md#org-orgid-integrations-integrationid-import)

### 신선한 컨테이너 이미지 가져오기

시나리오: [import-new-container-images](https://github.com/snyk-playground/cx-tools/blob/main/scripts/import-new-container-images.md) (전체 프로시저)

**사용된 Endpoints:**\
[특정 조직의 모든 프로젝트 목록](../reference/projects.md#orgs-org_id-projects)\
[대상 가져오기](../reference/import-projects-v1.md#org-orgid-integrations-integrationid-import)\
[가져오기 작업 세부 정보 가져오기](../reference/import-projects-v1.md#org-orgid-integrations-integrationid-import-jobid)\
[프로젝트 삭제](../reference/projects-v1.md#org-orgid-project-projectid-2)

### 저장소의 새 프로젝트(파일) 식별 및 정기적으로 Snyk로 가져오기

시나리오: [Identify-and-import-new-repos](https://github.com/snyk-playground/cx-tools/blob/main/scripts/Identify-and-import-new-repos.md) (전체 프로시저)

**사용된 Endpoint:**\
[조직 ID에 따른 대상 가져오기](../reference/targets.md#orgs-org_id-targets)\
[대상 가져오기](../reference/import-projects-v1.md#org-orgid-integrations-integrationid-import)

## Snyk 프로젝트 관리

### Snyk의 모든 프로젝트 태그 지정

시나리오: [Snyk의 프로젝트 태그 지정](https://github.com/snyk-playground/cx-tools/blob/main/scripts/tag-snyk-projects.md) (전체 프로시저)

**사용된 Endpoints:**\
[특정 조직의 모든 프로젝트 목록](../reference/projects.md#orgs-org_id-projects)

### 한 조직에서 다른 조직으로 프로젝트 이동

시나리오: [다른 조직 간 프로젝트 이동](https://github.com/snyk-playground/cx-tools/blob/main/scripts/move-projects.md) (전체 프로시저)

{% hint style="info" %}
사용된 API 토큰은 그룹 관리자 액세스가 필요합니다.\
다른 그룹 간 조직 이동을 하는 경우, 두 그룹 모두 그룹 관리자 권한이 있는 개인 API 토큰을 사용해야 합니다. 서비스 계정은 다른 그룹 간 조직 이동을 할 수 없습니다.&#x20;

보고를 위한 기존 데이터는 손실됩니다.
{% endhint %}

**사용된 Endpoints:**\
[다른 조직으로 프로젝트 이동](../reference/projects-v1.md#org-orgid-project-projectid-move)

## SCM과 통합

### 특별한 이벤트 또는 시간에 Snyk에서 코드베이스(소스 제어 관리 저장소)로의 모든 상호 작용(풀 리퀘스트, 테스트) 비활성화

시나리오: [disable-all-interaction-from-snyk](https://github.com/snyk-playground/cx-tools/blob/main/scripts/disable-all-interaction-from-snyk.md) (전체 프로시저)

**사용된 Endpoints 대안 1: 다른 조직의 통합 가져오기 및 각 통합 설정 업데이트**\
[리스트](../reference/integrations-v1.md#org-orgid-integrations-1) (통합)\
[업데이트](../reference/integrations-v1.md#org-orgid-integrations-integrationid-settings) (통합 설정)\
[기존 통합 업데이트](../reference/integrations-v1.md#org-orgid-integrations-integrationid)

**사용된 Endpoints 대안 2:** **웹훅 접근 방법: Snyk 웹훅 가져오기를 통해 웹훅 ID를 가져와 웹훅 삭제**\
[웹훅 목록](../reference/webhooks-v1.md#org-orgid-webhooks-1)\
[웹훅 삭제](../reference/webhooks-v1.md#org-orgid-webhooks-webhookid-1)\
[웹훅 생성](../reference/webhooks-v1.md#org-orgid-webhooks)

## 문제 가져오기와 관리

### 특정 그룹의 모든 프로젝트에 대한 프로젝트 스냅샷 검색

시나리오: [Retrieve-project-snapshots](https://github.com/snyk-playground/cx-tools/blob/main/scripts/retrieve-projects-snapshots.md) (전체 프로시저)

**사용된 Endpoints:**\
[그룹 내 모든 조직 목록](../reference/groups-v1.md#group-groupid-orgs)\
[최신 문제 목록 가져오기](../reference/reporting-api-v1.md#reporting-issues-latest)

### 특정 취약점에 영향을 받는 모든 프로젝트 찾기

시나리오: [find-all-projects-affected-by-a-vuln.md](https://github.com/snyk-playground/cx-tools/blob/main/scripts/find-all-projects-affected-by-a-vuln.md) (전체 프로시저)

**사용된 Endpoints:**\
[문제 목록 가져오기](../reference/reporting-api-v1.md#reporting-issues)\
[그룹 내 모든 조직 목록](../reference/groups-v1.md#group-groupid-orgs)\
[특정 조직의 모든 프로젝트 목록](../reference/projects.md#orgs-org_id-projects)

### 대량 문제 무시

시나리오: [bulk-ignore-issues](https://github.com/snyk-playground/cx-tools/blob/main/scripts/bulk-ignore-issues.md) (전체 프로시저)

**사용된 Endpoints:**\
[특정 조직의 모든 프로젝트 목록](../reference/projects.md#orgs-org_id-projects)\
[최신 문제 목록 가져오기](../reference/reporting-api-v1.md#reporting-issues-latest) (모든 문제 가져오기, 코드 제외)\
[조직 ID별 문제 가져오기](../reference/issues.md#orgs-org_id-issues) (코드 문제 가져오기)

### 조직의 모든 프로젝트에서 모든 문제(Snyk Code 문제 포함) 목록

시나리오: [list-all-issues-for-a-snyk-org](https://github.com/snyk-playground/cx-tools/blob/main/scripts/list-all-issues-for-a-snyk-org.md) (전체 프로시저)

**사용된 Endpoints:**\
[특정 조직의 모든 프로젝트 목록](../reference/projects.md#orgs-org_id-projects)\
[모든 집계된 문제 목록](../reference/projects-v1.md#org-orgid-project-projectid-aggregated-issues) (코드 제외)\
[조직 ID별 문제 가져오기](../reference/issues.md#orgs-org_id-issues)\
REST 실험적 [ID 별  문제 가져오기](https://apidocs.snyk.io/?version=2022-04-06%7Eexperimental#get-/orgs/-org_id-/issues/detail/code/-issue_id-)\
[무시된 사항 검색](../reference/ignores-v1.md#org-orgid-project-projectid-ignore-issueid-2)