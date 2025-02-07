# 사전 정의된 역할

{% hint style="info" %}
**기능 가용성**

그룹 수준 역할은 Enterprise Enterprise 플랜에서만 사용할 수 있습니다. 자세한 내용은 [요금제 및 가격](https://snyk.io/plans/)을 참조하세요.
{% endhint %}

Snyk은 Snyk Web UI 또는 [Snyk API](../user-management-with-the-api/)를 사용하여 할당 및 관리할 수 있는 표준 사용자 역할 세트를 제공합니다. 사전 정의된 역할의 권한 세트는 사용자 지정할 수 없습니다. 대신, Snyk Web UI의 "Manage role"에서 사용자 지정 역할을 생성하는 것을 권장합니다.

Snyk에서 제공하는 사전 정의된 역할은 다음과 같습니다:

* **Organization Admin:** 팀 리드(Team Leads)에 해당하는 표준 역할입니다. 이 역할을 가진 사용자는 프로젝트를 추가 및 삭제하고, Snyk 검사를 무시하며, 그룹 멤버에게 조직 수준 역할을 부여할 수 있습니다.
* **Organization Collaborator:** 개발자(Developers)에 해당하는 표준 역할입니다. 소규모 팀이나 개발자 중심 조직 접근 방식에 적합합니다.
* **Group Admin:** 조직 내에서 Snyk 사용을 총괄하는 역할로, 그룹 및 조직 수준에서 모든 권한을 가집니다. 또한, Group Admin은 그룹에 속한 모든 조직의 Organization Admin이 되지만, 조직 수준 사용자 목록에는 표시되지 않습니다.
* **Group Viewer:** 그룹 수준에서 액세스할 수 있지만, Snyk에서 작업을 수행하려면 조직 수준의 권한이 필요한 사용자입니다. 이는 Snyk 온보딩 과정에서 그룹 권한과 기능을 이해하고, 배포 후 사용자 지정 그룹 역할을 설계하는 데 사용됩니다.
* **Group Member:** 온보딩 이후 사용자 지정 역할을 생성하지 않으려는 경우 Group Viewer에서 전환하는 비기능적 사용자 역할입니다. 이 역할에 부여된 권한은 사용자의 요구 사항에 따라 달라질 수 있으며, Snyk 담당자와 논의하여 결정됩니다. Snyk Web UI에서 "Manage Members" 섹션의 목록에서 해당 역할을 선택하여 할당된 권한을 확인할 수 있습니다.
* **Tenant Admin:** 모든 테넌트(Tenant) 제품과 설정에 액세스할 수 있는 사용자입니다. 이 역할은 계정 소유자 및 관리자 전용입니다.
* **Tenant Viewer:** 테넌트의 모든 사용자 목록 및 해당 테넌트에 설정된 모든 그룹과 조직을 조회할 수 있는 사용자입니다.
* **Tenant Member:** 모든 테넌트 사용자의 기본 역할이지만, 테넌트 수준의 옵션에는 액세스할 수 없습니다.

## 역할 유형

역할은 조직(Organization), 그룹(Group), 테넌트(Tenant) 수준에서 관리할 수 있습니다.

테넌트 수준 역할은 그룹 및 조직 수준 역할을 부여하지 않습니다.

그룹 수준 역할에는 조직 및 그룹 수준에서의 권한이 포함됩니다. 그룹 역할에 추가된 조직 권한은 해당 그룹 내 모든 조직에 자동으로 적용됩니다. 예를 들어, 사전 정의된 그룹 역할인 **Group Viewer**는 그룹에서 일부 보기 권한을 부여하며, 동시에 해당 그룹 내 모든 조직에 대한 읽기 전용 권한을 제공합니다.

조직 역할은 조직 수준에서만 권한을 부여합니다. 조직 역할은 특정 조직에 대한 특정 권한을 부여하는 데 유용합니다.

조직 및 그룹 수준 역할을 조합하여 사용자에게 필요한 액세스를 부여할 수 있습니다. 예를 들어, 특정 조직에 **Organization Admin** 권한을 부여하면서, 그룹 내 다른 모든 조직에 대해 읽기 전용 액세스를 제공하려면 해당 사용자에게 **Group Viewer** 역할을 그룹 수준에서 부여하고, 특정 조직에서 **Organization Admin** 역할을 부여할 수 있습니다.

## 조직 수준 권한

다음 표는 각 사전 정의된 역할에 적용되는 조직 수준 권한을 나타냅니다.

|  | Org Admin | Org Collaborator | Group Admin | Group Viewer |
|---|---|---|---|---|
| 조직 보기(View Organization) | x | x | x | x |
| 조직 편집(Edit Organization) | x |  | x |  |
| 조직 삭제(Remove Organization) | x |  | x |  |
| 조직 보고서 보기(View Organization Reports) | x | x | x | x |
| 프로젝트 보기(View Project) | x | x | x | x |
| 프로젝트 추가(Add Project) | x | x | x |  |
| 프로젝트 편집(Edit Project) | x | x | x |  |
| 프로젝트 상태(Project Status) | x | x | x |  |
| 프로젝트 테스트(Test Project) | x | x | x |  |
| 프로젝트 이동(Move Project) | x |  | x |  |
| 프로젝트 삭제(Remove Project) | x | x | x |  |
| 프로젝트 기록 보기(View Project History) | x | x | x | x |
| 프로젝트 통합 편집(Edit Project Integrations) | x |  | x |  |
| 프로젝트 속성 편집(Edit Project Attributes) | x |  | x |  |
| Jira 이슈 보기(View Jira Issues) | x | x | x | x |
| Jira 이슈 생성(Create Jira Issues) | x | x | x |  |
| 프로젝트 태그 편집(Edit Project Tags) | x | x | x |  |
| 프로젝트 무시 보기(View Project Ignores) | x | x | x | x |
| 프로젝트 무시 생성(Create Project Ignores) | x | x | x |  |
| 프로젝트 무시 편집(Edit Project Ignores) | x | x | x |  |
| 프로젝트 무시 제거(Remove Project Ignores) | x | x | x |  |
| Pull Requests 생성(Create Pull Requests) | x | x | x |  |
| Pull Request 체크 성공으로 표시(Mark Pull Request checks as successful) | x |  | x |  |
| 서비스 계정 보기(View Service Accounts) | x |  | x | x |
| 서비스 계정 생성(Create Service Accounts) | x |  | x |  |
| 서비스 계정 편집(Edit Service Accounts) | x |  | x |  |
| 서비스 계정 삭제(Remove Service Accounts) | x |  | x |  |
| 사용자 보기(View Users) | x | x | x | x |
| 사용자 초대(Invite Users) | x |  | x |  |
| 사용자 관리(Manage Users) | x |  | x |  |
| 사용자 추가(Add Users) | x |  | x |  |
| 사용자 승인(Provision Users) | x |  | x |  |
| 사용자 퇴사(User Leave) | x | x | x |  |
| 사용자 삭제(User Remove) | x |  | x |  |
| 통합 보기(View Integrations) | x | x | x | x |

## 조직 수준 권한

다음 표는 각 사전 정의된 역할에 적용되는 조직 수준 권한을 자세히 설명합니다.

|                       | 조직 관리자 | 조직 협력자 | 그룹 관리자 | 그룹 뷰어 |
|-----------------------|--------------|--------------|--------------|--------------|
| 조직 보기            | x            | x            | x            | x            |
| 조직 편집            | x            |              | x            |              |
| 조직 제거            | x            |              | x            |              |
| 조직 보고서 보기    | x            | x            | x            | x            |
| 프로젝트 보기       | x            | x            | x            | x            |
| 프로젝트 추가       | x            | x            | x            |              |
| 프로젝트 편집       | x            | x            | x            |              |
| 프로젝트 상태       | x            | x            | x            |              |
| 프로젝트 테스트     | x            | x            | x            |              |
| 프로젝트 이동       | x            |              | x            |              |
| 프로젝트 제거       | x            | x            | x            |              |
| 프로젝트 기록 보기 | x            | x            | x            | x            |
| 프로젝트 통합 편집 | x            |              | x            |              |
| 프로젝트 속성 편집 | x            |              | x            |              |
| Jira 이슈 보기      | x            | x            | x            | x            |
| Jira 이슈 생성      | x            | x            | x            |              |
| 프로젝트 태그 편집 | x            | x            | x            |              |
| 프로젝트 무시 보기 | x            | x            | x            | x            |
| 프로젝트 무시 생성 | x            | x            | x            |              |
| 프로젝트 무시 편집 | x            | x            | x            |              |
| 프로젝트 무시 제거 | x            | x            | x            |              |
| 풀 리퀘스트 생성     | x            | x            | x            |              |
| 풀 리퀘스트 확인 표시 | x            |              | x            |              |
| 컬렉션 보기         | x            | x            | x            | x            |
| 컬렉션 생성         | x            |              | x            |              |
| 컬렉션 편집         | x            |              | x            |              |
| 컬렉션 삭제         | x            |              | x            |              |
| 서비스 계정 보기    | x            |              | x            | x            |
| 서비스 계정 생성    | x            |              | x            |              |
| 서비스 계정 편집    | x            |              | x            |              |
| 서비스 계정 제거    | x            |              | x            |              |
| 사용자 보기         | x            | x            | x            | x            |
| 사용자 초대         | x            |              | x            |              |
| 사용자 관리         | x            |              | x            |              |
| 사용자 추가         | x            |              | x            |              |
| 사용자 할당         | x            |              | x            |              |
| 사용자 나가기       | x            | x            | x            |              |
| 사용자 제거         | x            |              | x            |              |
| 통합 보기           | x            | x            | x            | x            |
| 통합 편집           | x            |              | x            |              |
| 패키지 테스트       | x            | x            | x            |              |
| 청구 보기           | x            |              | x            |              |
| 청구 편집           | x            |              | x            |              |
| 권리 보기           | x            | x            | x            | x            |
| 미리보기 기능 보기 | x            | x            | x            |              |
| 미리보기 기능 편집 | x            |              | x            |              |
| 감사 로그 보기      | x            | x            | x            |              |
| 아웃바운드 웹훅 보기 | x            |              | x            |              |
| 아웃바운드 웹훅 생성 | x            |              | x            |              |
| 아웃바운드 웹훅 제거 | x            |              | x            |              |
| 앱 보기             | x            |              | x            |              |
| 앱 설치             | x            |              | x            |              |
| 앱 생성             | x            |              | x            |              |
| 앱 편집             | x            |              | x            |              |
| 앱 삭제             | x            |              | x            |              |
| 환경 보기           | x            | x            | x            | x            |
| 환경 생성           | x            |              | x            |              |
| 환경 삭제           | x            |              | x            |              |
| 환경 업데이트       | x            |              | x            |              |
| 스캔 보기           | x            | x            | x            | x            |
| 스캔 생성           | x            | x            | x            |              |
| 리소스 보기         | x            | x            | x            | x            |
| 아티팩트 보기       | x            | x            | x            | x            |
| 아티팩트 생성       | x            | x            | x            |              |
| 사용자 정의 룰 보기 | x            | x            | x            | x            |
| 사용자 정의 룰 생성 | x            | x            | x            |              |
| 사용자 정의 룰 편집 | x            | x            | x            |              |
| 사용자 정의 룰 제거 | x            | x            | x            |              |
| 컨테이너 이미지 보기 | x           |              | x            |              |
| Kubernetes 리소스 게시 | x         |              | x            |              |
| Snyk 학습 관리 | x               |              | x            |              |

## 그룹 수준 권한

이 표는 각 사전 정의된 역할에 적용되는 그룹 수준 권한을 자세히 설명합니다.

<table><thead><tr><th></th><th width="92">Org Admin</th><th width="128">Org Collaborator</th><th width="87">Group Admin</th><th>Group Viewer</th></tr></thead><tbody><tr><td>그룹 보기</td><td></td><td></td><td>x</td><td>x</td></tr><tr><td>그룹 세부 정보 편집</td><td></td><td></td><td>x</td><td></td></tr><tr><td>그룹 설정 보기</td><td></td><td></td><td>x</td><td></td></tr><tr><td>설정 편집</td><td></td><td></td><td>x</td><td></td></tr><tr><td>그룹 알림 설정 보기</td><td></td><td></td><td>x</td><td></td></tr><tr><td>그룹 알림 설정 편집</td><td></td><td></td><td>x</td><td></td></tr><tr><td>조직 보기</td><td></td><td></td><td>x</td><td>x</td></tr><tr><td>조직 추가</td><td></td><td></td><td>x</td><td></td></tr><tr><td>조직 제거</td><td></td><td></td><td>x</td><td></td></tr><tr><td>역할 읽기</td><td></td><td></td><td>x</td><td></td></tr><tr><td>역할 생성</td><td></td><td></td><td>x</td><td></td></tr><tr><td>역할 편집</td><td></td><td></td><td>x</td><td></td></tr><tr><td>역할 제거</td><td></td><td></td><td>x</td><td></td></tr><tr><td>사용자 보기</td><td></td><td></td><td>x</td><td>x</td></tr><tr><td>그룹에 사용자 추가</td><td></td><td></td><td>x</td><td></td></tr><tr><td>그룹 내 사용자 편집</td><td></td><td></td><td>x</td><td></td></tr><tr><td>사용자 제거</td><td></td><td></td><td>x</td><td></td></tr><tr><td>사용자 삭제</td><td></td><td></td><td>x</td><td></td></tr><tr><td>사용자 프로비저닝</td><td></td><td></td><td>x</td><td></td></tr><tr><td>역할 할당 및 할당 해제</td><td></td><td></td><td>x</td><td></td></tr><tr><td>서비스 계정 보기</td><td></td><td></td><td>x</td><td></td></tr><tr><td>서비스 계정 생성</td><td></td><td></td><td>x</td><td></td></tr><tr><td>서비스 계정 편집</td><td></td><td></td><td>x</td><td></td></tr><tr><td>서비스 계정 제거</td><td></td><td></td><td>x</td><td></td></tr><tr><td>감사 로그 보기</td><td></td><td></td><td>x</td><td></td></tr><tr><td>정책 보기</td><td></td><td></td><td>x</td><td></td></tr><tr><td>정책 생성</td><td></td><td></td><td>x</td><td></td></tr><tr><td>정책 편집</td><td></td><td></td><td>x</td><td></td></tr><tr><td>정책 삭제</td><td></td><td></td><td>x</td><td></td></tr><tr><td>보고서 보기</td><td></td><td></td><td>x</td><td>x</td></tr><tr><td>태그 보기</td><td></td><td></td><td>x</td><td>x</td></tr><tr><td>IaC 설정 보기</td><td></td><td></td><td>x</td><td></td></tr><tr><td>IaC 설정 편집</td><td></td><td></td><td>x</td><td></td></tr><tr><td>기능 플래그 보기</td><td></td><td></td><td>x</td><td></td></tr><tr><td>기능 플래그 편집</td><td></td><td></td><td>x</td><td></td></tr><tr><td>요청 액세스 설정 보기</td><td></td><td></td><td>x</td><td></td></tr><tr><td>요청 액세스 설정 편집</td><td></td><td></td><td>x</td><td></td></tr><tr><td>SSO 설정 보기</td><td></td><td></td><td>x</td><td></td></tr><tr><td>SSO 설정 편집</td><td></td><td></td><td>x</td><td></td></tr><tr><td>앱 보기</td><td></td><td></td><td>x</td><td></td></tr><tr><td>앱 설치</td><td></td><td></td><td>x</td><td></td></tr><tr><td>앱 편집</td><td></td><td></td><td>x</td><td></td></tr><tr><td>AppRisk 보기</td><td></td><td></td><td>x</td><td></td></tr><tr><td>AppRisk 편집</td><td></td><td></td><td>x</td><td></td></tr><tr><td>인사이트 접근</td><td></td><td></td><td>x</td><td>x</td></tr></tbody></table>

## 테넌트 수준 권한 <a href="#tenant-level-permissions" id="tenant-level-permissions"></a>

테넌트 권한은 **테넌트 멤버** 페이지에서 설정 및 관리됩니다. 개별적으로 선택할 수 있는 권한은 없으며, **Tenant Admin**, **Tenant Viewer**, **Tenant Member**의 테넌트 역할 이름만 존재합니다. 자세한 내용은 [테넌트에서 사용자 관리](Manage users in a Tenant)를 참조하세요.