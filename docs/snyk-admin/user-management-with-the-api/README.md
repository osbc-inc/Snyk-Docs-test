# API를 사용한 사용자 관리

{% hint style="info" %}
**기능 가용성**

Snyk API는 엔터프라이즈 요금제에서만 사용할 수 있습니다. 자세한 정보는 [요금제 및 가격](https://snyk.io/plans/)을 확인하세요.
{% endhint %}

Snyk 웹 UI와 Snyk API를 통해 사용자를 관리할 수 있습니다. 또한 API를 사용하여 [서비스 계정을 관리](../../enterprise-setup/service-accounts/manage-service-accounts-using-the-snyk-api.md)할 수도 있습니다.

* 초기 로그인 전에 SSO 사용자를 위해 특정 역할에 대한 권한을 구성하고 부여하기 위해 [provisioning 엔드포인트](provision-users-to-organizations-using-the-api.md)를 사용할 수 있습니다.
* [회원 역할을 업데이트](update-member-roles-using-the-api.md)하려면 API를 사용해야 합니다.
* 멤버 엔드포인트를 사용하여 프로그래밍 방식으로 [그룹 및 조직에서 사용자를 제거](remove-members-from-groups-and-orgs-using-the-api.md)할 수 있습니다.
* 사용자 활동을 모니터링하기 위해 그룹 레벨의 감사 로그 엔드포인트를 사용하여 지난 90일간의 [감사 로그](retrieve-audit-logs-of-user-initiated-activity-by-api-for-an-org-or-group.md)를 검색할 수 있습니다.