# 원격 IaC 사용자 지정 규칙 번들 사용

사용자 정의 규칙 번들을 생성한 후에는 [번들 푸시하기](../writing-rules-using-the-sdk/pushing-a-bundle.md) 단계에 따라 지원되는 OCI 레지스트리 중 하나로 번들을 배포할 수 있습니다.

사용자 지정 규칙 번들을 성공적으로 푸시한 후에는 다음 중 하나를 사용하여 번들을 사용하는 것을 강제할 수 있습니다.:

- [Snyk 설정](use-a-remote-iac-custom-rules-bundle.md#snyk-settings-and-remote-custom-rules-bundle)
- [Snyk API](use-a-remote-iac-custom-rules-bundle.md#snyk-api-and-remote-custom-rules-bundle)
- [환경 변수](use-a-remote-iac-custom-rules-bundle.md#environment-variables-and-remote-custom-rules-bundle)

이러한 옵션 중 하나를 사용하여 사용자 지정 규칙을 강제로 적용한 후에는 다음과 같이 Snyk CLI를 사용자 이름 및 암호로 구성하여 Snyk에서 OCI 레지스트리에서 풀 수 있도록 승인할 수 있습니다:

```plaintext
snyk config set oci-registry-username=<org registry username>
snyk config set oci-registry-password=<org registry password>
```

이렇게 하면 다음 Snyk 환경 변수가 설정됩니다.:

- `SNYK_CFG_OCI_REGISTRY_USERNAME`
- `SNYK_CFG_OCI_REGISTRY_PASSWORD`

이 구성을 완료한 후에는 다음과 같이 Snyk IaC 스캔을 실행할 수 있습니다. CLI는 백그라운드에서 구성된 컨테이너 레지스트리에 푸시된 번들을 가져올 것입니다.

```plaintext
snyk iac test <file>
```

결과 구성 스캔 이슈에는 기본 Snyk 규칙과 사용자 지정 규칙의 이슈가 포함됩니다. 또한 [IaC CLI 테스트 결과 이해](../../../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/understand-the-iac-cli-test-results/)도 참조하십시오.

{% hint style="info" %}
주어진 시간에 번들 경로를 정의하는 데 사용해야 할 메서드는 하나만 정의되어야 합니다. Snyk 설정 페이지 또는 [Snyk API](../../../../../snyk-api/reference/iacsettings.md)를 사용하여 사용자 지정 규칙 설정을 비활성화하도록 하거나 이전에 저장된 설정을 지우도록 하세요.
{% endhint %}

## Snyk 설정 및 원격 사용자 정의 규칙 번들

Snyk는 사용자 지정 규칙 설정을 구성하는 데 Snyk 설정 페이지를 사용하는 것을 권장합니다. 이는 이러한 설정이 수정될 때 매우 간단하게 업데이트할 수 있는 간단한 방법입니다.

{% hint style="info" %}
태그는 사용자 지정 규칙 번들의 버전 관리에 도움이 됩니다. 새 버전 태그를 설정하여 업데이트된 번들을 간단히 구성할 수 있습니다.
{% endhint %}

이러한 원격 번들을 구성하는 방법은 조직 및 그룹 수준 양쪽에서 가능합니다. 그룹에 대한 원격 번들을 구성하면 해당 그룹에 속한 모든 조직에 원격 번들이 적용됩니다.

원격 번들을 구성하려면:

- Infrastructure as Code 설정에서 **Rules** 섹션을 찾으십시오.

{% hint style="info" %}
조직 수준에서 원격 사용자 정의 규칙 번들을 구성하려면 `Settings` > `Infrastructure as Code`로 이동합니다.

그룹 수준에서 원격 사용자 정의 규칙 번들을 구성하려면 `Settings` > `Infrastructure as Code`로 이동합니다.
{% endhint %}

<figure><img src="../../../../../.gitbook/assets/image (161) (1) (1) (1) (1) (1) (2) (3).png" alt="원격 사용자 정의 규칙 번들 활성화"><figcaption><p>원격 사용자 정의 규칙 번들 활성화</p></figcaption></figure>

- **Rules** 토글을 사용하여 원격 번들의 구성을 활성화합니다. 이렇게 하면 양식이로드되어 아래 예에서 보여지는 레지스트리 URL 및 태그를 지정할 수 있습니다.

<figure><img src="../../../../../.gitbook/assets/image (102) (2) (1).png" alt="레지스트리 URL 및 태그 양식 지정"><figcaption><p>레지스트리 URL 및 태그 양식 지정</p></figcaption></figure>

- OCI 레지스트리 URL 및 원격 규칙 번들의 태그를 구성하고 **Save changes**를 클릭하여 변경 사항을 저장합니다.

<figure><img src="../../../../../.gitbook/assets/image (184) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1).png" alt="레지스트리 URL 및 태그 구성"><figcaption><p>레지스트리 URL 및 태그 구성</p></figcaption></figure>

이제 원격 사용자 지정 규칙 번들이 구성되어 IaC 파일을 테스트할 때 사용됩니다.

**그룹을 위해 원격 번들 구성을 재정의할 수 있습니다**.

기본적으로 그룹에 대한 원격 번들을 구성하면 해당 그룹에 속한 모든 조직에 대해 원격 번들이 적용됩니다. 따라서 그룹 구성이 업데이트되면 이러한 변경 사항이 해당 조직에 적용됩니다.

그러나 조직은 그룹 구성을 재정의하고 자체 번들 및 태그를 정의할 수 있습니다. 조직의 변경 사항은 그룹이 구성을 업데이트할 때 변경되지 않습니다.

그룹 구성을 재정의하려면 Infrastructure as Code 설정의 조직의 `Rules` 섹션으로 이동하십시오.

- 처음에는 조직의 그룹에서 상속된 구성으로 섹션이 채워집니다.

<figure><img src="../../../../../.gitbook/assets/image (165) (1) (1).png" alt="IaC 설정의 조직 규칙 섹션"><figcaption><p>IaC 설정의 조직 규칙 섹션</p></figcaption></figure>

- 구성을 조직에 맞도록 업데이트하고 **Save changes**를 클릭합니다.

<figure><img src="../../../../../.gitbook/assets/image (196) (1) (2).png" alt="조직 규칙 구성 업데이트됨"><figcaption><p>조직 규칙 구성 업데이트됨</p></figcaption></figure>

- 이제 그룹 수준의 구성은 조직의 이러한 사용자 정의 설정을 재정의하지 않고 있습니다.

그룹 기본값으로 복원하려면 **그룹 기본값으로 재설정** 버튼을 사용하십시오.

## Snyk API 및 원격 사용자 정의 규칙 번들

Snyk 설정 페이지를 통해 수동으로 설정을 업데이트하는 것이 시간이 많이 걸린다면, API를 사용하여 API 호출을 통해 어떤 변형의 사용자 정의 규칙 설정을 전송할 수 있는 Snyk API를 사용할 수 있습니다.

- 예를 들어, 그룹 수준에서 사용자 정의 규칙 번들을 구성하기 위해 [그룹을 위한 Infrastructure as Code 설정 업데이트하기](../../../../../snyk-api/reference/iacsettings.md#groups-group_id-settings-iac) 엔드포인트를 사용하여 다음 바디를 제공합니다:

```plaintext
{
   "data": {
         "type": "iac_settings",
         "attributes": {
           "custom_rules": {
             "oci_registry_url": "registry-1.docker.io/group-account/group-bundle-image",
             "oci_registry_tag": "1.3.14",
             "is_enabled": true
           }
       }
   }
}
```

- 태그만 업데이트하려는 경우 더 간단한 바디를 보낼 수 있습니다:

```plaintext
{
   "data": {
         "type": "iac_settings",
         "attributes": {
           "custom_rules": {
             "oci_registry_tag": "1.3.14"
           }
       }
   }
}
```

- 사용자 정의 규칙을 비활성화하려면 `is_enabled` 플래그를 보낼 수 있습니다:

```plaintext
{
   "data": {
         "type": "iac_settings",
         "attributes": {
           "custom_rules": {
             "is_enabled": false
           }
       }
   }
}
```

API는 변경 사항을 확인하기 위해 그룹 설정으로 회신합니다:

```plaintext
{
  "type": "iac_settings",
  "id": "<group id>",
  "attributes": {
    "custom_rules": {
      "oci_registry_url": "registry-1.docker.io/group-account/group-bundle-image",
      "oci_registry_tag": "1.3.14",
      "is_enabled": true
    },
   "updated": "2021-11-27T11:34:33.132Z"
  }
}
```

Snyk API를 사용하여 원격 번들 구성을 **재정의할 수 있습니다**.

설정 페이지와 마찬가지로 [그룹을 위한 Infrastructure as Code 설정 업데이트하기](../../../../../snyk-api/reference/iacsettings.md#groups-group_id-settings-iac) 엔드포인트를 사용하여 그룹의 모든 조직에 원격 번들을 적용할 수 있습니다. 조직은 API 호출을 사용하여 그룹의 구성을 재정의하고 자체 번들 및 태그를 정의할 수 있습니다.

- 그룹 구성을 재정의하려면, [조직을 위한 Infrastructure as Code 설정 업데이트하기](../../../../../snyk-api/reference/iacsettings.md#orgs-org_id-settings-iac) 엔드포인트를 호출하여 요청 바디에 다른 사용자 정의 규칙 번들과 태그를 제공하십시오:

```plaintext
{
   "data": {
         "type": "iac_settings",
         "attributes": {
           "custom_rules": {
             "oci_registry_url": "registry-1.docker.io/org-account/org-bundle-imageage",
             "oci_registry_tag": "1.3.15",
             "is_enabled": true
           }
       }
   }
}
```

- API는 변경 사항을 확인하기 위해 조직 설정 및 부모 섹션 아래의 그룹 설정을 회신하여 두 가지를 비교할 수 있습니다:

```plaintext
{
  "type": "iac_settings",
  "id": "<org id>",
  "attributes": {
    "custom_rules": {
      "oci_registry_url": "registry-1.docker.io/org-account/org-bundle-image",
      "oci_registry_tag": "1.3.15",
      "is_enabled": true
    },
   "updated": "2021-11-27T11:34:33.132Z",
   "parents": {
      "group": {
        "id": "<group id>",
        "type": "iac_settings",
        "attributes": {
          "custom_rules": {
            "oci_registry_url": "registry-1.docker.io/group-account/group-bundle-image",
            "oci_registry_tag": "1.3.14",
            "is_enabled": true
          },
          "updated": "2021-11-19T10:59:45.259Z"
        }
      }
    }
  }
}
```

- 그룹 설정으로 되돌리려면 다음 요청 바디를 제공하여 API를 호출하시면 되겠습니다:

```plaintext
{
   "data": {
         "type": "iac_settings",
         "attributes": {
           "custom_rules": {
             "inherit_from_parent": "group"
           }
       }
   }
}
```

- API는 변경 사항을 확인하기 위해 조직 설정 및 부모 섹션 아래의 그룹 설정을 회신합니다.

```plaintext
{
  "type": "iac_settings",
  "id": "<org id>",
  "attributes": {
    "custom_rules": {
      "oci_registry_url": "registry-1.docker.io/group-account/group-bundle-image",
      "oci_registry_tag": "1.3.14",
      "is_enabled": true,
      "inherit_from_parent": "group"
    },
   "updated": "2021-11-19T10:59:45.259Z",
   "parents"  {
      "group": {
        "id": "<group id>",
        "type": "iac_settings",
        "attributes": {
          "custom_rules": {
            "oci_registry_url": "registry-1.docker.io/group-account/group-bundle-image",
            "oci_registry_tag": "1.3.14",
            "is_enabled": true
          },
          "updated": "2021-11-19T10:59:45.259Z"
        }
      }
    }
  }
}
```

## 환경 변수 및 원격 사용자 정의 규칙 번들

조직을 위해 Snyk 구성을 사용하여 사용자 지정 규칙 번들의 위치를 구성할 수도 있습니다. 프로젝트 폴더에서 다음 명령을 사용하여 컨테이너 레지스트리를  CLI와 구성하세요:

```plaintext
snyk config set oci-registry-url=registry-1.docker.io/org-account/org-bundle-image:1.3.14
```

이렇게하면 Snyk 환경 변수 `SNYK_CFG_OCI_REGISTRY_URL`이 설정됩니다.

{% hint style="info" %}
OCI Registry URL이 유효한 URL인지 확인하십시오. 예를 들어 Docker Hub의 경우:

`registry-1.docker.io/org-account/org-bundle-image:1.3.14`

Snyk 설정 페이지에서 이전에 정의된 URL을 지우거나 사용자 정의 규칙을 비활성화해야 합니다. 주어진 시간에 번들 경로를 정의하는 데 사용해야 할 메서드는 하나만 정의되어야 합니다.
{% endhint %}

## 원격 사용자 정의 규칙 번들에 대한 문제 해결

-Ops-가 `-d` 옵션과 함께 실행하여 디버그 로그를 활성화하십시오.:

```plaintext
snyk iac test <file> -d
```

일부 가능한 문제점은 다음과 같습니다.:

- 잘못된 컨테이너 레지스트리 URL 제공. Docker Hub를 사용하는 경우 위의 노트를 참조하십시오. 에러는 다음과 같습니다.

```plaintext
디스크에 사용자 지정 번들을 다운로드할 수 없습니다.
원격 레지스트리 액세스를 확인하고 올바른 매개변수를 제공했는지 확인하십시오.
```

- 잘못된 자격 증명을 제공했습니다. 에러는:

```plaintext
인증 오류가 있었습니다. 잘못된 자격 증명이 제공되었습니다.
디스크에 사용자 지정 번들을 다운로드할 수 없습니다.
원격