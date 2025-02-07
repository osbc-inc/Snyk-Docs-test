# 통합 구성

## SDLC 통합 포인트 구성하기 (IDE, Git 저장소, CI/CD, 컨테이너 레지스트리)

{% hint style="info" %}
프로젝트 가져오기에 대한 자세한 정보는 [프로젝트 가져오기](../../phase-3-gain-visibility/import-projects.md)와 [롤아웃](../../phase-5-initial-rollout-to-team/)을 참조하십시오.
{% endhint %}

조직을 만들 때 Snyk를 소프트웨어 개발 생명주기(SDLC)에서 어디에 추가할지 결정해야 합니다.

Snyk는 취약점의 시각성을 우선 얻기를 권장하며, 관련된 기능은 처음에 비활성화된 채로 유지합니다.

첫 번째 단계는 사용 가능한 방법 중 하나를 사용하여 프로젝트를 가져오는 것이며, 기존 문제와 취약점에 대한 완전한 시각성을 확보한 후에 새로운 문제가 코드베이스에 추가되는 것을 방지하기 위해 점진적으로 게이팅 방법을 도입할 수 있습니다.

{% hint style="info" %}
**알림 설정 (이메일 알림)**

* Snyk는 사용자가 프로젝트를 가져오는 동안 곧바잇은 알림을 받지 않도록 모든 이메일 알림을 초기에 비활성화하는 것을 제안합니다.
* 새로운 조직에는 그룹 수준에서 이를 비활성화하고, 기존 조직에는 조직 수준에서 비활성화해야 합니다.
{% endhint %}

## 통합 도구

Snyk를 SDLC에 통합하여 시각성을 확보하는 가장 일반적인 도구에 대해 설명합니다.

### Git 저장소

Snyk는 여러 Git 저장소와 통합하여 코드의 문제와 취약점을 추적, 모니터링 및 해결하는 데 도움을 줄 수 있습니다.

온프레미스 Git 저장소를 사용하는 경우 통합을 가능하게 하려면 [Snyk 브로커](../../../../enterprise-setup/snyk-broker/)를 구성하고 실행해야 합니다.

저장소를 가져올 때 Snyk는 취약점을 스캔하고 기본 브랜치를 모니터링합니다.

모니터링은 다음을 포함합니다.

* 지정된 브랜치의 매일 스캔
* PR(풀 리퀘스트) 및 머지 체크

Git 저장소의 소스 제어 스캔은 지원되는 대부분의 언어에 적합하지만, 개인 패키지 관리자(예: Artifactory)를 사용하는 경우 해당 패키지도 스캔하려면 Snyk와 통합해야 합니다.

Git 저장소 통합을 통해 프로젝트를 수동으로 브라우저를 통해 가져올 수도 있고, [snyk-api-import](../../../../scan-with-snyk/snyk-tools/tool-snyk-api-import/) 도구를 사용하여 대량으로 저장소를 가져올 수도 있습니다. 또는 파이프라인에 삽입할 수 있는 특정 저장소를 가져오기 위한 엔드포인트 [Import Targets](../../../../snyk-api/reference/import-projects-v1.md#org-orgid-integrations-integrationid-import)를 사용할 수 있습니다.

{% hint style="info" %}
Snyk Enterprise 고객을 위해서는 **GitHub Enterprise 통합 카드**를 사용하는 것이 강력히 권장됩니다. 이 옵션을 사용하려면 GitHub Enterprise 고객이어야 할 필요는 없지만, 이 옵션을 사용하면 인터페이스의 접근에 대한 일관성 부분에서 OAuth가 아닌 개인 엑세스 토큰(PAT)을 사용할 수 있습니다.
{% endhint %}

Git 저장소에서 프로젝트를 가져온 경우, Snyk [PR 체크](../../../../scan-with-snyk/pull-requests/pull-request-checks/) 및 자동 PR 수정도 구성할 수 있습니다. 이를 통해 Git 저장소에서 PR을 제출할 때 자동으로 코드 변경 사항을 실시간으로 스캔하여 새로운 보안 문제가 코드베이스로 들어오는 것을 방지할 수 있습니다.

이를 통해 소프트웨어 개발 생명주기 초기에 모든 제출된 PR을 보안 문제에 대해 확인함으로써 스캔 및 시각성을 가능하게 할 수 있습니다.

{% hint style="info" %}
게이팅을 초기에 비활성화하려면 Snyk에서 프로젝트가 등록될 때 자동으로 구성된 매일 모니터링을 사용하고, 구성에서 PR/MR 체크를 비활성화하세요.

* Snyk이 제공하는 자동 수정
  * 자동 수정 PR
  * 자동 종속성 업데이트 PR
  * Snyk 취약점 패치
* PR 체크에서 확인할 수 있는 것
  * 오픈 소스 보안 및 라이선스
  * 코드 분석 (베타)
{% endhint %}

마찬가지로, 수정 및 업그레이드 PR 기능을 비활성화해야 할 수도 있습니다.

[Snyk AppRisk 통합](../configure-snyk-apprisk-integrations.md#setup-integrations)을 위해서는 Snyk 웹 UI의 Integration Hub 옵션을 사용하여 추가 설정과 구성이 필요합니다.

### CI/CD (빌드 파이프라인)

Snyk을 빌드 파이프라인 단계로 추가하여 취약한 애플리케이션 또는 컴포넌트(레지스트리)의 배포를 방지하여 애플리케이션을 보안하십시오.

CLI를 통해

* Scala, Gradle, GO와 같은 특정 패키지 관리자에 대한 최적의 정확성
* 빌드 환경이 개인 패키지에 액세스할 수 있는 경우 추가적인 통합 구성이 필요하지 않음
* 빌드가 실패하고 Snyk로 보고하거나 단순히 Snyk로 보고하는 프로덕션으로 푸시된 구성요소에 대한 시각성

사용할 수 있는 여러 [CI/CD 통합](../../../../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/)을 사용할 수도 있고, 테스트하는 방법에 대해 더 유연성을 가져오기 위해 파이프라인의 일부로 [Snyk CLI](../../../../snyk-cli/)를 사용할 수도 있습니다.

처음에는 Snyk이 제공하는 `monitor` 기능을 사용하여 정보를 가져와서 발견된 문제를 확인하는 것이 좋습니다. 이미 소스 제어 통합을 사용하여 이를 실현하고 있다면 진행하지 않아도 됩니다. 이후 새로운 취약점이 추가되는 것을 막기 위해 `test` 기능을 도입할 수 있으며, 처음에는 중요한 문제에서 빌드를 실패하도록하고 시간이 지남에 따라 실패 기준을 조정할 수 있습니다.

{% hint style="info" %}
`snyk iac test --report` 및 `snyk code test --report` (베타)를 사용하는 경우 문제 발견 시 빌드가 비정상 응답 코드로 중지될 수 있습니다.\
\
이를 수동으로 테스트하려면 `--report` 옵션을 포함하기 위해 빌드 단계를 항상 계속되도록 설정하거나 `snyk code test --report || true.`와 같이 항상 `or true`와 같은 대체 옵션을 추가해야 합니다. 구문의 구체적인 부분은 CLI가 실행되는 환경에 따라 다를 수 있습니다.
{% endhint %}

[Snyk 필터](https://docs.snyk.io/snyk-api/other-tools/tool-snyk-filter)나 새로운 문제를 강조하기 위한 [Snyk 델타](https://docs.snyk.io/snyk-api/other-tools/tool-snyk-delta)와 같은 플러그인을 사용하여 파이프라인을 구성하는 것이 인기가 있습니다.

다양한 파이프라인 통합의 데모는 [Snyk-Labs](https://github.com/snyk-labs/snyk-cicd-integration-examples)에서 찾을 수 있습니다.
