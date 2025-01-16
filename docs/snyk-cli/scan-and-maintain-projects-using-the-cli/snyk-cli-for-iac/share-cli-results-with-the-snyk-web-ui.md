# Snyk 웹 UI로 CLI 결과 공유

알려진 구성 문제를 해결하기 위해 [CLI](../../) `snyk iac test` 명령을 사용할 수 있습니다.

이러한 문제를 Snyk 웹 UI에서 볼 수 있도록하려면 다음 CLI 명령을 실행하십시오:

`snyk iac test myproject --report`

{% hint style="info" %}
[사용자 지정 규칙](../../../scan-with-snyk/snyk-iac/build-your-own-iac-custom-rules/current-iac-custom-rules/) 및 결과 공유 기능을 함께 사용하는 것은 현재 지원되지 않습니다.

Snyk은 파일 내용을 네트워크를 통해 공유하지 않으며, 방금 스캔된 구성 문제에 대한 필요한 메타데이터만 공유합니다.
{% endhint %}

## `snyk iac test --report` 예제 출력

```
> snyk iac test myproject --report

Testing arm-file.tf...


Infrastructure as code issues:
  ✗ VM Agent is not provisioned automatically for Windows [Low Severity] [SNYK-CC-AZURE-667] in Compute
    introduced by resource > azurerm_virtual_machine[my_terraformvm] > os_profile_windows_config > provision_vm_agent


Organization:      my.org
Type:              Terraform
Target file:       arm-file.tf
Project name:      myproject
Open source:       no
Project path:      myproject

Known issues were found when testing arm-file.tf, 1 issue found

Your test results are available at: https://app.snyk.io/org/my.org/projects 에서 myproject 이름으로 보고서를 가져올 수 있습니다.
```

이는 현재 구성 문제에 대한 스냅샷을 Snyk 대시보드로 보내어 Snyk 웹 UI에서 확인할 수 있습니다.

## Snyk 웹 UI에서 스냅샷보기

Snyk 웹 UI에 로그인하여 조직 프로젝트 페이지로 이동하여 스캔된 프로젝트의 최신 스냅샷을 확인할 수 있습니다.

<figure><img src="../../../.gitbook/assets/image (349) (1) (1) (1) (1) (1).png" alt="프로젝트 페이지에 새롭게 스캔된 프로젝트 목록"><figcaption><p>프로젝트 페이지에 새롭게 스캔된 프로젝트 목록</p></figcaption></figure>

또한 해당 프로젝트를 열어서 프로젝트 세부 정보를 확인할 수 있습니다:

<figure><img src="../../../.gitbook/assets/image (106) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (2).png" alt="프로젝트 세부 정보"><figcaption><p>프로젝트 세부 정보</p></figcaption></figure>

## **무시하기**

Snyk 웹 UI를 사용하거나 프로젝트 스캔 시 `.snyk` 정책 파일을 함께 생성하여 문제를 무시할 수 있습니다. 자세한 내용은 [.snyk 정책 파일을 사용한 Iac 무시](iac-ignores-using-the-.snyk-policy-file.md)을 참조하십시오.

{% hint style="info" %}
`.snyk` 정책 파일을 사용하여 무시된 문제는 Snyk 웹 UI에서 다시 보이지 않습니다.
{% endhint %}

## 프로젝트 태그

`--project-tags` 옵션을 사용하여 스캔된 프로젝트에 태그를 부착할 수 있습니다. 이 옵션은 쉼표로 구분된 각 태그가 키-값 쌍인 목록을 수용합니다. 키와 값은 `=` 기호로 구분됩니다. `--project-tags` 옵션은 `--report`와 함께 사용할 때에만 유효합니다.

다음 예제는 `department`와 `team` 태그를 각각 `platform` 및 `persistence` 값으로 스캔된 프로젝트에 부착하는 방법을 보여줍니다.

```
> snyk iac test myproject --report \
    --project-tags=department=platform,team=persistence
```

## 프로젝트 속성

`--project-business-criticality`, `--project-environment`, 및 `--project-lifecycle` 옵션을 사용하여 스캔된 프로젝트에 속성을 설정할 수 있습니다. 이러한 옵션은 `--report`와 함께 사용될 때에만 유효합니다.

* `--project-business-criticality`는 `critical`, `high`, `medium`, `low` 값 중에서 선택할 수 있는 쉼표로 구분된 목록을 수용합니다.
* `--project-environment`는 `frontend`, `backend`, `internal`, `external`, `mobile`, `saas`, `onprem`, `hosted`, `distributed` 값 중에서 선택할 수 있는 쉼표로 구분된 목록을 수용합니다.
* `--project-lifecycle`은 `production`, `development`, `sandbox` 값 중에서 선택할 수 있는 쉼표로 구분된 목록을 수용합니다.

다음 예제는 각 스캔된 프로젝트에 대해 비즈니스 중요도를 `high`, 환경을 `frontend` 및 `internal`, 그리고 라이프사이클을 `development`로 설정하는 방법을 보여줍니다.

```
> snyk iac test myproject --report \
    --project-business-criticality=high \
    --project-environment=frontend,internal \
    --project-lifecycle=development
```

## 대상 참조

`--target-reference` 옵션을 사용하여 스캔된 프로젝트에 대한 대상 참조를 설정할 수 있습니다. 이 옵션은 `--report`와 함께 사용될 때에만 유효합니다.

다음 예제는 스캔된 프로젝트에 대한 대상 참조를 현재 Git 브랜치 이름으로 설정하는 방법을 보여줍니다.

```
snyk iac test myproject --report \
    --target-reference="$(git branch --show-current)"
```
