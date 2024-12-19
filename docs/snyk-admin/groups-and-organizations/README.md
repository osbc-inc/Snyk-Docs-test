# 그룹 및 조직

{% hint style="info" %}
**기능 가용성**\
Snyk 그룹은 엔터프라이즈 플랜에서만 사용 가능합니다. 자세한 정보는 [플랜 및 가격](https://snyk.io/plans/)을 참조하십시오.
{% endhint %}

## 그룹, 조직 및 프로젝트

Snyk에는 Snyk 스캔과 기능에 대한 액세스를 제어할 수 있는 계층 구조가 있습니다. 계층 구조의 주요 부분은 다음 다이어그램에서 설명되어 있습니다:

<figure><img src="../../.gitbook/assets/image (1) (1) (1) (2).png" alt="그룹, 조직 및 프로젝트"><figcaption><p>그룹, 조직 및 프로젝트</p></figcaption></figure>

* [**그룹**](groups/): 대부분 그룹은 전체 Snyk 사용자를 포괄하지만 대기업은 여러 그룹을 가질 수 있습니다. 그룹에는 여러 조직을 포함할 수 있습니다.
* [**조직**](organizations/): 조직은 팀과 같은 특정 비즈니스 영역을 나타냅니다. 조직에는 여러 프로젝트를 포함할 수 있습니다.
* [**프로젝트**](../snyk-projects/)**:** 프로젝트는 Snyk가 문제를 스캔하는 항목을 기반으로 설정됩니다. 각 프로젝트는 스캔 결과를 보여줍니다. 해당 프로젝트에서 문제를 어떻게 스캔할지 정의하도록 프로젝트를 구성할 수 있습니다.

Snyk에는 또한 [조직 내 사용자 관리](organizations/manage-users-in-organizations.md) 및 [그룹 내 사용자 관리](groups/manage-users-in-a-group.md) 기능이 있습니다. Snyk API v1을 사용하여 [조직에 사용자를 할당](../user-management-with-the-api/provision-users-to-organizations-using-the-api.md)하거나 [그룹과 조직에서 회원을 제거](../user-management-with-the-api/remove-members-from-groups-and-orgs-using-the-api.md)할 수 있습니다.

사용자를 추가하거나 [Snyk 그룹의 세션 길이를 구성](groups/configure-session-length-for-a-snyk-group.md)하기 위해 조직 액세스 요청을 사용할 수 있습니다.

새로운 사용자가 추가되거나 예기치 않은 활동을 분석하려면 Snyk REST API를 통해 조직 또는 그룹에 대한 사용자 주도 활동의 감사 로그를 [검색](../user-management-with-the-api/retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md)할 수 있습니다.

이 섹션의 설명서는 [그룹](groups/) 및 [조직](organizations/)을 다룹니다.