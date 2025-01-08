# 사용자 지정 매핑

커스텀 매핑을 통해 식별 공급자(Identity Provider, IdP)가 제공하는 데이터에 기반하여 사용자를 Snyk 그룹 및 조직에 동적으로 할당하여 확장 가능한 사용자 프로비저닝 및 액세스 모델을 구현할 수 있습니다.

{% hint style="info" %}
이 옵션을 구현하려면 Snyk 계정 팀과 협력하십시오.
{% endhint %}

Snyk 내 역할 및 권한에 대해 자세히 알아보려면 [사전 정의된 역할](../../../snyk-admin/user-roles/pre-defined-roles.md)을 참조하십시오. 또한 [사용자 역할 관리](../../../snyk-admin/user-roles/user-role-management.md)도 확인하십시오.

## 커스텀 매핑 요구 사항

* [SSO 설정을 위한 자원](../set-up-snyk-single-sign-on-sso.md#resources-for-sso-setup)에서 찾을 수 있는 적절한 IdP(식별 공급자)를 위한 SSO 정보 워크시트를 완료하십시오.
* IdP에서 사용자 역할을 할당하는 데 사용할 커스텀 속성을 적절히 구성하십시오. [예: 역할 배열 매핑](./#example-roles-array-mapping)을 참조하십시오.

## 커스텀 매핑 옵션

이 페이지에 설명된 Snyk의 최신 커스텀 매핑 옵션은 사용자들에게 그룹 수준의 사용자 정의 역할 및 사전 정의 역할을 부여할 수 있는 능력을 포함하여 유연성이 향상됩니다.

Snyk에서는 여전히 [레거시 커스텀 매핑](legacy-custom-mapping.md) 옵션을 지원합니다.

## Snyk와의 역할 배열 매핑

IdP에서 `roles`를 문자열 배열로 전달해야 합니다. 다양한 IdP를 위한 이를 설정하는 방법에 대한 [예시](examples-setting-up-custom-mapping-for-idps/)가 제공됩니다.

구성하는 방법에 대한 자세한 내용은 식별 공급자 설명서를 참조하십시오.

## 커스텀 매핑 어설션

이 섹션은 Snyk이 Snyk 그룹 및 조직 내의 Snyk 역할에 매핑될 때 기대하는 역할 어설션을 문서화합니다.

### 역할 어설션 형식

역할 어설션은 다음 형식으로 Snyk에 제공되어야 합니다:

`snyk:{scope}:{target}:{role}`

여기서:

* `snyk`은 역할 매핑을 위한 고정 접두사입니다. **필수 요소**입니다.
* `scope`는 `org` 또는 `group` 중 하나일 수 있습니다. **필수 요소**입니다. 역할 매핑에 유효한 scope가 포함되지 않은 경우 무시됩니다.
* `target`는 역할이 부여될 `org` 또는 `group`의 **slug**이어야 하거나 **와일드카드**여야 합니다.
* `role`은 필요한 역할의 표준화된 이름입니다.

자세한 내용은 [스러그](./#slugs) 섹션에서 확인하십시오.

{% hint style="warning" %}
사용자는 조직이나 그룹 당 하나의 역할만 매핑해야 합니다. 원격지 문자가 사용되는 경우를 제외하고 여러 역할을 매핑하는 것은 지원되지 않으며 예기치 않은 동작을 유발할 수 있습니다.
{% endhint %}

{% hint style="info" %}
SSO에 속한 조직에서 역할을 부여받은 사용자는 역할 어설션에 명시된 그룹 수준 역할이 없는 경우 해당 그룹의 **그룹 멤버** 그룹 수준 역할이 암시적으로 할당됩니다. 이는 해당 사용자가 그룹의 멤버가 되도록 하는 데 필요한 미리 정의된 그룹 수준 역할로 해당됩니다.
{% endhint %}

### 역할 어설션 예시

* `snyk:group:*:group_admin`: 사용자에게 SSO 연결에 연결된 모든 그룹에 대한 사전 정의 **그룹 관리자** 역할을 할당합니다.
* `snyk:group::custom:sys_admin`: 사용자에게 SSO 연결에 연결된 모든 그룹에 대한 사용자 정의 그룹 수준 역할 `Sys Admin`을 할당합니다.

### 역할 어설션 배열 예시

사용자에 대한 일련의 역할 어설션 예시는 다음과 같습니다:

```json
{
    "roles": [
        "snyk:group:*:group_viewer",
        "snyk:org:development:org_admin",
        "snyk:org:test-org-N58YhztauHcaMiNfvi5fbL:custom:developer_readonly"
    ]
}
```

{% hint style="info" %}
시스템은 배열 대신 역할을 쉼표로 구분한 목록도 지원합니다.

```json
{
  "roles": "snyk:group:*:group_viewer, snyk:org:development:org_admin, snyk:org:test-org-N58YhztauHcaMiNfvi5fbL:custom:developer_readonly" 
}
```
{% endhint %}

이러한 어선션을 통해 사용자에게 할당됩니다:

* 모든 그룹에 대한 사전 정의 **그룹 뷰어** 그룹 수준 역할. 부여되는 권한에 대한 자세한 내용은 [사전 정의 역할](../../../snyk-admin/user-roles/pre-defined-roles.md)을 참조하십시오.
* **조직 관리자** 조직 수준 역할(Deployment라는 이름의 조직에 대해).
* **개발자 읽기 전용**이름이 **테스트 조직**, 슬러그가 `test-org-N58YhztauHcaMiNfvi5fbL`인 조직에 대해 사용자 정의 조직 수준 역할.

## 와일드카드

커스텀 매핑은 모든 조직이나 그룹에 대해 동일한 역할을 할당하는 어설션을 허용하는 와일드카드를 도입했습니다.

와일드카드를 사용한 어설션은 특정 대상과 함께 어설션이 사용되는 경우보다 낮은 우선 순위를 가집니다.

{% hint style="info" %}
SSO 연결 내에서 조직에 역할이 부여되는 경우 역할 어설션에서 그룹 수준 역할이 명시적으로 지정되어 있지 않은 경우 해당 그룹의 **그룹 멤버** 그룹 수준 역할이 암시적으로 할당됩니다. 이는 해당 사용자가 그룹의 멤버가 되도록 하는 데 필요한 미리 정의된 그룹 수준 역할로 해당됩니다.
{% endhint %}

## 슬러그

유효한 역할 어설션을 얻기 위해 와일드카드가 사용되지 않는 경우 조직 또는 그룹 슬러그가 필요할 수 있습니다. 슬러그는 Snyk 내에서 조직 또는 그룹의 정식 이름입니다.

조직 슬러그를 찾으려면 **Settings** 페이지로 이동하여 **General** 설정 아래에서 조직 슬러그 값을 확인할 수 있습니다. 이 값을 복사하여 커스텀 매핑의 역할 어설션에 사용할 수 있습니다.

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252Ff3juPcOGIxRqh3P26JgM%252Fsettings_org_slug.png%3Falt%3Dmedia%26token%3D1dda021b-5144-4118-a6e4-96ffb197cc29&#x26;width=768&#x26;dpr=1&#x26;quality=100&#x26;sign=40c5747c&#x26;sv=2" alt=""><figcaption><p>조직 일반 설정 페이지, 조직 슬러그 표시</p></figcaption></figure>

그룹의 슬러그를 찾으려면 그룹 설정으로 이동하여 General Settings 아래에서 그룹 슬러그를 찾을 수 있습니다.

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252FlzQHxchjbOFAC7i7L8ZN%252Fimage.png%3Falt%3Dmedia%26token%3D05d9b3ad-7b61-439e-aed0-029add524e7c&#x26;width=768&#x26;dpr=1&#x26;quality=100&#x26;sign=51c6f646&#x26;sv=2" alt=""><figcaption><p>그룹 일반 설정 페이지, 그룹 슬러그 표시</p></figcaption></figure>

## 역할 표준화된 이름

커스텀 매핑에서 사용할 역할의 표준화된 이름을 찾으려면 먼저 역할이 Snyk 그룹에 존재하는지 확인하십시오. **Group Settings** > **Member Roles** > **Role**로 이동하여 이를 확인할 수 있습니다.

자세한 내용 및 특히 사용자 정의 역할에 대한 자세한 내용은 [사용자 역할 관리](../../../snyk-admin/user-roles/user-role-management.md)를 참조하십시오.

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252FvHpPz3hxVKUW5pxPcB26%252Fimage.png%3Falt%3Dmedia%26token%3D1340330c-585e-4b03-aeff-a2d986b0fc26&#x26;width=768&#x26;dpr=1&#x26;quality=100&#x26;sign=45bd9bbd&#x26;sv=2" alt=""><figcaption><p>조직 관리자 역할의 역할 세부 정보 페이지</p></figcaption></figure>

## 사전 정의 역할 슬러그

Snyk에는 [사전 정의 역할](../../../snyk-admin/user-roles/pre-defined-roles.md) 세트가 있습니다. 해당 정규화된 이름은 다음과 같습니다:

| 역할 유형 | 역할 이름            | 역할 슬러그             |
| ----- | ---------------- | ------------------ |
| 조직    | Org Admin        | `org_admin`        |
| 조직    | Org Collaborator | `org_collaborator` |
| 그룹    | Group Admin      | `group_admin`      |
| 그룹    | Group Viewer     | `group_viewer`     |
| 그룹    | Group Member     | `group_member`     |
