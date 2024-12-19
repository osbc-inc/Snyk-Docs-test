# 프로비저닝 옵션 선택

{% hint style="info" %}
Snyk 계정 팀과 협력하여 이 SSO 옵션 구현을 준비하세요.
{% endhint %}

회사의 새로운 사용자가 Snyk에 액세스하는 방법을 결정합니다:

* [모두에게 공개](choose-a-provisioning-option.md#open-to-all)
* [초대 필요](choose-a-provisioning-option.md#invitation-required)
* [사용자 정의 매핑](choose-a-provisioning-option.md#custom-mapping)

## 모두에게 공개

오픈 옵션을 선택하면, SSO를 사용하여 Snyk에 자동 액세스할 수 있습니다. 모든 사용자는 선택한 그룹 내의 모든 조직에 액세스할 수 있습니다. 사용자 계정은 동일한 역할로 프로비저닝되며 두 가지 옵션이 있습니다:

* **조직 관리자** 역할은 모든 새로운 사용자가 다른 조직 관리자 및 협업자를 관리하고, 그룹 보고서를 볼 수 있으며, 그룹 내의 조직과 조직 내에서 비관리적인 작업을 수행할 수 있습니다.
* **조직 협업자** 역할은 조직 내에서 비관리적인 작업을 수행할 수 있습니다.

Snyk 지원팀에게 새로운 사용자가 조직에서 **관리자** 또는 **협업자** 역할을 할 것인지 알려주세요. 선택한 역할은 모든 사용자에게 적용됩니다. 자세한 내용은 [미리 정의된 역할(Pre-defined roles)]( ../../snyk-admin/user-roles/pre-defined-roles.md)을 참조하세요.

## 초대 필요

초대 필요 또는 **그룹 구성원** 옵션을 선택할 경우, 관리자는 사용자를 초대하거나 사용자가 조직에 액세스를 요청할 수 있습니다.

조직에 사용자를 초대하는 두 가지 방법이 있습니다. 구성원을 초대(자세한 내용은 [조직에서 사용자 관리](../../snyk-admin/groups-and-organizations/organizations/manage-users-in-organizations.md) 참조)하거나 Snyk [API 사용자 초대 엔드포인트](https://snyk.docs.apiary.io/#reference/organizations/user-invitation-to-organization/invite-users)를 사용하여 과정을 자동화할 수 있습니다.

초대되지 않은 사용자가 SSO를 사용하여 로그인하면 Snyk에 액세스할 수 있지만, 관리자가 그들을 초대하거나 해당 조직에 수동으로 추가할 때까지 프로젝트를 볼 수 없습니다. 새로운 사용자가 [액세스를 요청할 수 있는](../../snyk-admin/groups-and-organizations/organizations/requests-for-access-to-an-organization.md) 적절한 연락처가 있는 조직 목록을 제공할 수 있습니다.

## 사용자 정의 매핑

{% hint style="info" %}
**기능 이용 가능성**\
Custom Mapping은 엔터프라이즈 플랜에서만 이용 가능합니다. 자세한 정보는 [요금제 및 가격](https://snyk.io/plans/)을 참조하세요.&#x20;
{% endhint %}

[사용자 정의 매핑 옵션](custom-mapping/)을 사용하여 사용자 계정을 사용자 정의 규칙으로 프로비저닝할 수 있습니다.

다른 그룹마다 SSO를 다르게 구성할 수 있습니다. 또한 ID 제공자로부터의 정보를 기반으로 사용자를 특정 조직 및 역할 할당에 매핑할 수도 있습니다.