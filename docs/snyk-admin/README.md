# 관리자

관리는 다음 기능을 포함합니다:

- [그룹 및 조직 관리](groups-and-organizations/)
- [Snyk 프로젝트 관리 및 사용](snyk-projects/)
- [조직에서 사용자 관리](groups-and-organizations/organizations/manage-users-in-organizations.md) 및 [그룹에서 사용자 관리](groups-and-organizations/groups/manage-users-in-a-group.md)
- [사용자 역할 관리](user-roles/)
- [알림 관리](manage-notifications.md)
- [설정 관리](groups-and-organizations/group-and-organization-settings.md)

이 페이지에서는 다음 주제를 다룹니다:

- [계정, 그룹, 조직, 대상 및 프로젝트](./#accounts-groups-organizations-targets-and-projects)
- [사용자 유형](./#user-types)
- [관리자 도구](./#snyk-admin-tools)

{% hint style="info" %}
**기능 가용성**

일부 기능인 그룹과 같은 기능은 특정 요금제에서만 사용할 수 있습니다. 자세한 내용은 [요금제 및 가격 정책](https://snyk.io/plans/)을 참조하세요.\

Enterprise 설정에 대한 자세한 내용은 [Enterprise 설정](../enterprise-setup/)을 참조하세요.
{% endhint %}

## 계정, 그룹, 조직, 대상 및 프로젝트

Snyk에는 스캔 및 기타 Snyk 기능에 대한 액세스를 제어하는 계층 구조가 있습니다.

- **계정:** 사용자는 Snyk 계정에 로그인하여 스캔 및 설정 및 스캔 결과 보기 또는 수정을 할 수 있습니다.
- [**그룹**](groups-and-organizations/groups/): 고객은 하나 이상의 Snyk 그룹을 갖습니다. 그룹은 회사 전체에 해당할 수 있습니다. 대기업은 부서, 부문 또는 회사의 다른 부분과 일치하는 여러 그룹을 가질 수 있습니다. 그룹에는 여러 조직이 포함될 수 있습니다.
- [**조직**](groups-and-organizations/organizations/): Snyk 그룹은 하나 이상의 Snyk 조직을 포함합니다. 조직은 팀과 같은 특정 비즈니스 영역을 나타냅니다. 조직에는 여러 프로젝트가 포함될 수 있습니다.
- [**대상**](snyk-projects/#target): 각 대상은 Snyk로 가져온 저장소를 대상으로 하여 스캔하고 다시 테스트합니다.
- [**프로젝트**](snyk-projects/): 프로젝트는 Snyk가 문제를 스캔하는 항목을 기반으로 설정되며, 매니페스트 파일 등의 결과를 보여줍니다. 프로젝트에서 어떻게 문제를 스캔할지 정의하기 위해 프로젝트를 구성할 수 있습니다. [시작하기](getting-started/)를 참조하세요.

Snyk 관리자는 그룹 및 조직을 설정합니다. 자세한 내용은 [그룹 및 조직 관리](groups-and-organizations/)을 참조하세요. 대상 및 [프로젝트](snyk-projects/)는 Snyk 사용자가 개발 프로젝트를 Snyk로 가져와 스캔할 때 생성됩니다.

## 사용자 유형

Snyk에는 다음과 같은 사전 정의된 사용자 유형이 있습니다:

- 조직 관리자
- 조직 공동 작업자
- 그룹 관리자
- 그룹 뷰어
- 그룹 구성원

자세한 내용은 [사전 정의된 역할](user-roles/pre-defined-roles.md)을 참조하십시오. 각 역할과 관련된 권한을 확인할 수 있습니다.

{% hint style="info" %}
**기능 가용성**

그룹 수준 역할은 Enterprise 요금제에서만 사용할 수 있습니다. 자세한 내용은 [요금제 및 가격 정책](https://snyk.io/plans/)을 참조하세요.&nbsp;
{% endhint %}

## 관리자 도구

Snyk는 그룹, 조직 및 사용자 역할 및 권한 관리, 알림 및 설정을 관리할 수 있는 도구를 제공합니다.

### 사용자 및 권한 관리

그룹에서 사용자 및 권한을 관리할 수 있습니다. 자세한 내용은 [사용자 및 권한 관리](user-roles/user-role-management.md)을 참조하십시오.

<figure><img src="../.gitbook/assets/image (245) (1) (1) (1).png" alt="멤버 관리 인터페이스"><figcaption><p>멤버 관리 인터페이스</p></figcaption></figure>

### 그룹 및 조직 관리

Snyk 그룹 및 조직은 팀 간의 협력을 유지하는 데 도움이 됩니다. 자세한 내용은 [그룹 및 조직 관리](groups-and-organizations/)를 참조하십시오.

### 알림 정의

본인 및 당신의 조직의 이메일 알림을 관리할 수 있습니다. 자세한 내용은 [알림 관리](manage-notifications.md)를 참조하십시오.

<figure><img src="../.gitbook/assets/image (6) (2).png" alt="이메일 알림 관리 인터페이스"><figcaption><p>이메일 알림 관리 인터페이스</p></figcaption></figure>

### 설정 관리

귀하의 Snyk 계정을 귀하의 작업 프로세스에 맞게 사용자 정의할 수 있습니다. 자세한 내용은 [설정 관리](groups-and-organizations/group-and-organization-settings.md)를 참조하십시오.