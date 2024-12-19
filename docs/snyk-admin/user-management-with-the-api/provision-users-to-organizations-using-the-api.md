# API를 사용하여 조직에 사용자 할당하기

Provision 사용자 엔드포인트를 사용하면 사용자가 Snyk 플랫폼에 로그인하기 전에 싱글 사인온 사용자의 권한을 조직화하고 부여할 수 있습니다. 엔드포인트는 [조직에 사용자 할당](../../snyk-api/reference/organizations-v1.md#org-orgid-provision), [보류 중인 사용자 할당 목록](../../snyk-api/reference/organizations-v1.md#org-orgid-provision-1), 및 [보류 중인 사용자 할당 삭제](../../snyk-api/reference/organizations-v1.md#org-orgid-provision-2)입니다.

할당된 사용자는 초대를 수락할 필요가 없습니다. 할당된 사용자가 처음 Snyk에 로그인하면 모든 권한이 부여됩니다. 첫 번째 로그인 전에 사용자를 조직에 대량으로 추가하려면 [조직에 사용자 할당](../../snyk-api/reference/organizations-v1.md#org-orgid-provision) 엔드포인트를 사용할 수 있습니다.

## API를 사용하여 사용자 할당하기 위한 전제 조건

{% hint style="info" %}
API에서 서비스 계정을 초대하는 사용자나 할당받는 사용자로 사용할 수 없습니다.
{% endhint %}

* 할당받는 사용자는 이미 Snyk 시스템에 존재하지 않아야 합니다.
* 초대하는 사용자는 개인 토큰을 사용하여 API를 호출해야 합니다.
* Organizations에 속한 Snyk 그룹은 [싱글 사인온(Single Sign On, SSO)이 구성되어 있어야 합니다](../../enterprise-setup/single-sign-on-sso-for-authentication-to-snyk/).
* 초대하는 사용자와 할당받는 사용자 모두 SSO를 사용하여 로그인해야 합니다.
* 초대하는 사용자는 이러한 호출을 실행하기 위해 `Provision Users` 권한을 갖고 있어야 합니다. 모든 그룹 및 조직 관리자는 기본적으로 이 권한을 갖습니다.

<figure><img src="../../.gitbook/assets/Screenshot 2022-09-09 at 09.57.17.png" alt="Provision Users 권한 활성화"><figcaption><p>Provision Users 권한 활성화</p></figcaption></figure>

## Provision 사용자 API 사용 방법

다음은 Provision 사용자 엔드포인트의 사용 방법을 설명합니다. 엔드포인트에 대한 API 문서에서 엔드포인트에 대해 더 자세한 정보를 보려면 [조직에 사용자 할당](../../snyk-api/reference/organizations-v1.md#org-orgid-provision), [보류 중인 사용자 할당 목록](../../snyk-api/reference/organizations-v1.md#org-orgid-provision-1), 및 [보류 중인 사용자 할당 삭제](../../snyk-api/reference/organizations-v1.md#org-orgid-provision-2)를 참조하십시오.

### 조직에 사용자 할당

엔드포인트 [조직에 사용자 할당](../../snyk-api/reference/organizations-v1.md#org-orgid-provision)을 사용하여 특정 조직에 사용자를 할당하고 역할을 지정할 수 있습니다. 사용자가 처음으로 Snyk에 로그인하면 역할에 정의된 권한이 자동으로 할당됩니다.

**`POST`** `https://api.snyk.io/v1/org/orgId/provision`

**요청 모델:**

`{`

`"email": "test@example.com",`

`"rolePublicId": "",`

`"role": "ADMIN"`

`}`

**응답 모델:**

`{`

`"email": "test@example.com",`

`"rolePublicId": "",`

`"role": "ADMIN",`

`"created": Date`

`}`

{% hint style="info" %}
Enterprise 프로그램 사용자는 자체 사용자 정의 [멤버 역할](../user-roles/user-role-management.md)을 정의하고 할당을 위해`rolePublicId`를 사용할 수 있습니다.\
\
동일한 호출에서`role` 또는`rolePublicId` 중 하나만 사용할 수 있습니다.
{% endhint %}

### 보류 중인 사용자 할당 목록

엔드포인트 [보류 중인 사용자 할당 목록](../../snyk-api/reference/organizations-v1.md#org-orgid-provision-1)은 응답에 보류 중인 할당된 사용자를 반환합니다.

**`GET`** `https://api.snyk.io/v1/org/orgId/provision`

**응답 모델:**

`[`

`....`

`{`

`"email": "test@example.com",`

`"rolePublicId": "",`

`"role": "ADMIN",`

`"created": Date`

`},`

`....`

`]`

### 보류 중인 사용자 할당 삭제

보류 중인 할당 요청을 제거하려면 엔드포인트 [보류 중인 사용자 할당 삭제](../../snyk-api/reference/organizations-v1.md#org-orgid-provision-2)를 사용합니다.

**`DELETE`** `https://api.snyk.io/v1/org/orgId/provision`

쿼리 매개변수

* 이메일 (문자열) - 사용자의 이메일

**응답 모델:**

`{`

`"ok": true`

`}`