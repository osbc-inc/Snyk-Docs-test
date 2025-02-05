---
설명: 이 페이지는 Snyk Learn 액세스 제어를 어떻게 관리하는지에 대한 단계별 안내를 제공합니다.
---

# Snyk Learn 액세스 제어

{% hint style="info" %}
**릴리스 상태**

[Snyk Learn 과제](snyk-learn-assignments.md)는 학습 관리 add-on 제공의 일환으로 얼리 엑세스 중입니다. 자세한 정보는 Snyk 계정 팀에 문의하십시오.
{% endhint %}

**Snyk Learn 액세스 제어** 기능을 사용하면 Snyk 플랫폼과 동일한 [액세스 제어 모델](../../snyk-admin/user-roles/user-role-management.md)을 사용하여 Snyk Learn의 사용자 권한을 관리할 수 있습니다. 그룹 관리자는 Snyk Learn 과제와 관련된 권한 세트를 사용하여 사용자 권한을 관리할 수 있습니다. 이러한 역할은 조직의 사용자와 기능을 반영할 수 있습니다. 예를 들어, 귀하의 조직에서 교육 및 훈련 관리 기능에 대한 Snyk Learn의 권한을 교육 및 훈련 관리자로 제한할 수 있습니다.

{% hint style="info" %}
그룹 및 조직 관리자는 추가 권한 없이 기본적으로 Snyk Learn 보고서 및 과제에 액세스할 수 있습니다.
{% endhint %}

기본 조직 및 그룹 관리자 역할을 넘어서 과제에 대한 액세스 권한을 부여하려면 [사용자 지정 역할](../../snyk-admin/user-roles/custom-role-templates/snyk-learn-learning-admin.md)을 만들어야 합니다. 사용자 지정 역할은 과제를 관리하기 위한 권한 세트를 가져야 합니다. 고유 역할은 과제를 관리하기 위해 View Users 권한을 가져야 합니다. 다음은 과제를 관리하는 데 필요한 권한을 나열합니다:

* 조직 과제보기
* 조직 과제 편집
* 조직 계약자 생성
* 조직 과제 삭제

역할이 생성된 후, 그룹 또는 조직 관리자는 이 역할을 Snyk 사용자에게 부여할 수 있습니다.

자세한 내용은 [Snyk Learn 과제](snyk-learn-assignments.md)을 참조하십시오.
