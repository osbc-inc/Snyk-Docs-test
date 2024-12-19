# Legacy custom mapping

이 옵션을 구성하려면 SAML 속성 또는 OIDC 클레임 내에서 `roles` 배열을 전송하여 다음 중 **하나**의 패턴을 준수하도록하세요:

snyk-groupadmin

- 이 역할 매핑은 사용자에게 그룹 관리자 역할을 할당합니다.
- **groupadmin** 은 이 역할을 가지는 모든 사용자를 해당 사용자가 지정된 그룹에 할당된 모든 그룹에 그룹 관리자로 설정하여 해당 그룹에 속하는 모든 조직을 조직 관리자 권한으로 부여합니다.

snyk-groupviewer

- 이 역할 매핑은 그룹 뷰어 역할을 갖는 사용자를 할당하며 해당 그룹, 보고서 및 해당 그룹과 관련된 모든 조직에 대해 읽기 전용 액세스 권한을 부여합니다.

snyk-{그룹ID}

- 이 역할 매핑은 지정된 그룹 아래의 모든 조직 공동 작업자 역할을 사용자에게 할당합니다.
- **그룹ID**는 Snyk에서 그룹 수준에서 찾을 수 있는 그룹에 대한 ID 문자열입니다: `https://app.snyk.io/group/<그룹 ID>` 또는 **그룹 드롭다운 -> 설정 -> 일반 -> 그룹 ID**.

snyk-{orgslug}-{role}

- 이 역할 매핑은 `orgslug`로 지정된 Snyk 조직에 대해 지정된 역할인 공동 작업자 또는 관리자 또는 사용자 정의 역할을 갖는 사용자를 할당합니다.
- `orgslug`은 Snyk에서 조직 이름의 고유 식별자입니다.
  - `orgslug`를 찾는 방법: `https://app.snyk.io/org/{orgslug}` 또는 엔드포인트 [그룹 내 모든 조직 나열](../../../snyk-api/reference/groups-v1.md#group-groupid-orgs)을 사용하여 찾을 수 있습니다.
  - **참고**: 대부분의 경우 `orgslug`는 조직의 이름입니다. 그러나 예외가 있을 수 있습니다.
  - 참고: `orgslug`는 최대 60자까지의 값을 가질 수 있습니다.
- **role**:
  - 표준 역할을 사용하는 경우, `{role}`은 `collaborator` 또는 `admin` 중 하나여야 합니다**.**
  - 사용자 정의 역할을 사용할 수도 있으며, `{role}`에는 정규화된 이름을 사용해야 합니다. 자세한 내용은 [사용자 정의 SSO에서의 역할](../../../snyk-admin/user-roles/user-role-management.md#roles-in-custom-sso)을 참조하세요.

{% hint style="info" %}
사용자는 조직당 하나의 역할만 매핑해야 합니다. 조직에 대해 여러 역할을 매핑하는 것은 지원되지 않으며 예기치 않은 동작을 유발할 수 있습니다.
{% endhint %}

## Roles 배열 매핑 형식

사용자에게 그룹 관리자 역할을 할당하려면 다음 형식을 사용하세요:

```
{
    "roles": [
        "snyk-groupadmin"
    ]
}
```

사용자에게 그룹 뷰어 역할을 할당하려면 다음 형식을 사용하세요:

```
{
    "roles": [
        "snyk-groupviewer"
    ]
}
```

사용자에게 조직 공동 작업자 역할을 할당하려면 다음 형식을 사용하세요:

```
{
    "roles": [
        "snyk-{groupID}"
    ]
}
```

사용자에게 조직 관리자 또는 조직 공동 작업자 역할을 할당하려면 역할 배열에 대한 다음 형식을 사용하세요. **참고**: 조직별로 다른 역할을 할당할 수 있습니다. 다음 예제는 사용자를 `orgslug` 조직에서 조직 관리자로 할당하고 `orgslug2` 조직에서는 공동 작업자로 할당합니다.

```
{
    "roles": [
        "snyk-{orgslug}-admin",
        "snyk-{orgslug2}-collaborator"
    ]
}
```

사용자에게 사용자 정의 역할을 할당하려면 역할 배열에 대한 다음 형식을 사용하세요. 각 조직에 대해 다른 역할을 할당하고 다른 조직에 대해 표준 및 사용자 정의 역할을 혼합해서 사용할 수 있습니다.

```
{
    "roles": [
        "snyk-{orgslug}-admin",
        "snyk-{orgslug2}-collaborator",
        "snyk-{orgslug3}-developer_readonly"
    ]
}
```

{% hint style="info" %}
시스템은 배열 대신 역할의 쉼표로 구분된 목록도 지원합니다.

```
{
  "roles": "snyk-{orgslug}-admin,snyk-{orgslug2}-collaborator"
}
```
{% endhint %}

## 역할 배열 매핑 예시

다음 예는 매핑 규칙에 따라 Snyk 사용자에게 역할을 할당하는 방법을 보여줍니다.

- 고객은 ABC로 지정되어 있으며 ABC라는 그룹이 하나 있습니다.
- 고객은 Snyk 내에서 Application-SecurityScanner1, Partner-Plugins 및 Application-Payments라는 세 가지 조직을 갖고 있습니다.
- 고객은 비즈니스 개발, 엔지니어링, 보안 및 제품이라는 네 개의 팀을 가지고 있습니다. 각 팀은 다른 요구 사항을 가지고 있습니다:
  - 비즈니스 개발팀은 ABC 그룹 및 Partner-Plugins 조직에 대해 Org Admin으로만 액세스해야 합니다.
  - 엔지니어링팀은 ABC 그룹, Application-SecurityScanner1 조직을 Org Admin으로, Partner-Plugins 조직을 Org Admin으로, 그리고 Application-Payments를 Org Collaborator로 액세스해야 합니다.
  - 보안팀은 ABC 그룹을 그룹 관리자로, 그리고 세 조직을 Org Admin으로 액세스해야 합니다.
  - 제품 팀은 ABC 그룹 및 세 조직을 Org Collaborator로 액세스해야 합니다.

비즈니스 개발팀의 경우, Snyk는 snyk-{orgslug}-{role} 매핑을 사용합니다:

```
{
    "roles": [
        "snyk-partner-plugins-admin"
    ]
}
```

엔지니어링팀의 경우, Snyk는 snyk-{orgslug}-{role} 매핑을 사용합니다:

```
{
    "roles": [
        "snyk-application-securityscanner1-admin",
        "snyk-partner-plugins-admin",
        "snyk-application-payments-collaborator"
    ]
}
```

보안팀의 경우, Snyk는 snyk-groupadmin 매핑을 사용합니다:

```
{
    "roles": [
        "snyk-groupadmin"
    ]
}
```

제품팀의 경우, Snyk는 snyk-{groupID} 매핑을 사용합니다. 이때 groupID의 값을 삽입해야 합니다.

```
{
    "roles": [
        "snyk-{groupID}"
    ]
}
```

## 사용자 정의 매핑하에 있는 역할에 대한 요약 다이어그램

<figure><img src="../../../.gitbook/assets/image (373).png" alt=""><figcaption></figcaption></figure>