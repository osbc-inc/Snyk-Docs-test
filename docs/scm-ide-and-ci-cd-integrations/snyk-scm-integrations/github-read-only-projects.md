# GitHub 읽기 전용 프로젝트

Snyk 계정에 새 통합을 추가하려면 먼저 통합을 설치할 수준 유형을 결정해야 합니다.

* [그룹 수준](github-read-only-projects.md#group-level-snyk-apprisk-integrations) - Snyk 애플리케이션에 통합을 추가하여 Snyk AppRisk Essentials 또는 Snyk AppRisk Pro에서 사용할 수 있습니다. Snyk AppRisk에 대한 통합을 설정하려면 그룹 수준의 통합 메뉴를 사용하세요.
* [조직 수준](github-read-only-projects.md#organization-level-snyk-integrations) - Snyk 애플리케이션에 모든 Snyk 제품을 위해 사용할 수 있는 통합을 추가합니다. Snyk AppRisk 제외.

Snyk은 GitHub 읽기 전용 프로젝트를 제공하여 조직에서 소유하지 않는 공개 GitHub 저장소를 모니터링할 수 있습니다.

## 조직 수준 - Snyk 통합

### GitHub 읽기 전용 프로젝트 작동 방식

읽기 전용 프로젝트를 추가하면 Snyk을 사용하여 문제를 미연에 방지하거나 해결할 필요가 없는 공개 저장소에서 의존성으로 고려 중인 프로젝트, 비즈니스 내에서 독립적 도구로 이미 사용 중인 프로젝트 또는 기타 저장소의 취약점을 추적할 수 있습니다.

저장소는 조직의 GitHub 자격 증명을 사용하여 매일 테스트됩니다. 이러한 자동화된 테스트는 Snyk 계획의 테스트 제한 부분으로 계산되지 않습니다.

Snyk GitHub 통합을 통해 가져온 프로젝트와는 달리, 읽기 전용 상태로 가져오거나 모니터링된 프로젝트는 다음을 할 수 없습니다.

* 병합 요청이 병합될 때 자동 재테스트 사용.
* 새로운 취약점이 도입되지 않도록 감지하고 선택적으로 차단하기 위해 발생한 모든 PR에 대해 커밋 테스트.
* 취약점을 해결하는 데 필요한 최소한의 변경을 추천하기 위해 [자동수정 PR](../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/create-automatic-prs-for-new-fixes-fix-prs.md) 사용.
* 종속성을 최신 상태로 유지하고 새로운 취약점을 방지하며 발견된 취약점을 간소화하기 위해 [자동 종속성 업그레이드 PR](../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs-upgrade-prs/) 사용.
* 사용자가 선택한 구체적인 문제를 해결하기 위해 Snyk를 통해 생성된 수동 Fix PR 사용.

### 공개 GitHub 저장소 모니터링 방법

다음 단계에 따라 **대시보드** 및 **프로젝트** 탭의 **Add project > Monitor public GitHub repos** 메뉴를 사용하거나 [공개 GitHub 저장소 모니터링](https://app.snyk.io/add/github-readonly)으로 이동하여 읽기 전용 프로젝트를 가져옵니다.

1. `owner/repository` 형식을 따라 모니터링할 공개 저장소를 입력합니다.
2. 유효한 저장소 이름을 입력한 경우 **+ Add repo**를 클릭합니다.\
   저장소는 빠르게 지원하는 매니페스트 파일을 위해 테스트됩니다.
3. 모니터링하려는 공개 저장소를 입력하고 **Import N repository/ies**를 선택합니다.

<figure><img src="../../.gitbook/assets/github_readonly_steps 2 &#x26; 3_18july2022.png" alt="Add repo 및 저장소 가져오기"><figcaption><p>Add repo 및 저장소 가져오기</p></figcaption></figure>

## 그룹 수준 - Snyk AppRisk 통합

Snyk AppRisk에 대한 GitHub 설정 가이드로 이동하여 [Snyk AppRisk를 위한 GitHub 설정 가이드](github-enterprise.md#github-setup-guide-for-snyk-apprisk)에서 Snyk AppRisk를 위한 GitHub 통합을 설정하는 방법에 대한 모든 세부 정보를 확인하세요.
