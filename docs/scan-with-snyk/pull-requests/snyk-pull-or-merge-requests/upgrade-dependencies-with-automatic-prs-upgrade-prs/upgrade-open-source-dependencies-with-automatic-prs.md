# 오픈 소스 종속성 자동 PR로 업그레이드

Git 리포지토리를 가져오면 Snyk는 취약젬과 라이선스, 그리고 종속성 건강 문제를 스캔하여 지속적으로 감시합니다. Snyk는 구성 설정에 따라 자동으로 pull request(PR)를 생성합니다.

<figure><img src="../../../../.gitbook/assets/image (435).png" alt="Snyk Bot conversation card in GitHub reporting PR raised"><figcaption><p>Snyk Bot conversation card in GitHub reporting PR raised</p></figcaption></figure>

## 지원되는 언어 및 SCMs

Snyk는 npm, Yarn 및 Maven Central 리포지토리에 대한 **자동 종속성 업그레이드 pull request** 기능을 지원합니다. 지원하는 소스 컨트롤 관리자(SCM)는 다음과 같습니다: GitHub, GitHub Enterprise, GitHub Cloud App, Bitbucket Server, Bitbucket Cloud, Bitbucket Connect, GitLab 및 Azure Repos.

이 기능은 Snyk Broker를 사용할 수도 있습니다. 이 기능을 사용하려면 Snyk Broker를 버전 1.4.55.0 이상으로 업그레이드해야 합니다. 자세한 정보는 [Snyk Broker 클라이언트 업그레이드](../../../../enterprise-setup/snyk-broker/upgrade-the-snyk-broker-client.md)를 참조하십시오.

## 자동 종속성 (업그레이드) PRs

자동 종속성 또는 업그레이드 PRs 이는 다음과 같이 작동합니다.

1. [조직 수준의 통합 설정](upgrade-open-source-dependencies-with-automatic-prs.md#enabling-the-automatic-dependency-upgrade-prs-option-for-an-entire-organization)에서 **자동 종속성 업그레이드 pull request** 옵션을 활성화해야 합니다.
2. 리포지토리 가져오기 시, Snyk가 리포지토리를 스캔하고 결과를 제공합니다. 그 후 Snyk는 주기적으로 스캔하여 오픈 소스 프로젝트를 지속적으로 모니터링합니다. 다시 스캔하는 빈도는 프로젝트 설정에 설정된 일정에 따라 결정됩니다.
3. 스캔마다 종속성의 새 버전이 식별되면, Snyk는 자동 업그레이드 PRs를 생성합니다.
   * 이미 업그레이드되었거나 다른 열린 Snyk PR에서 패치된 종속성에 대해 Snyk는 새로운 업그레이드 PR을 열지 않습니다. 마찬가지로, 해당 문제에 대해 열리고 업그레이드가 이루어지기 전에 닫힌 PR에 대해서도 동일합니다.
   * Snyk는 각 종속성마다 별도의 PR를 엽니다.
   * 기본적으로 Snyk는 프로젝트가 다섯 개 이상의 열린 Snyk PR을 가지고 있다면 업그레이드 PR을 생성하지 않습니다. 새 PR은 생성되지 않습니다. 이 한도는 통합 또는 프로젝트 설정에서 1-10 사이로 설정할 수 있습니다. 이 한도는 업그레이드 PR에만 적용되며 Fix PR은 제한되지 않습니다.
   * 기본적으로 Snyk는 패치 및 소규모 업그레이드 만 권장합니다. 그러나 주 버전 업그레이드에 대한 권장 사항은 해당 기능이 활성화된 **설정** 페이지에서 요청할 수 있습니다.
   * 최신 버전이 프로젝트에서 아직 발견되지 않은 취약점을 포함하는 경우, Snyk는 업그레이드를 권장하지 않습니다.
   * 21일 미만의 버전 업그레이드에 대한 권장 사항은 Snyk가 하지 않습니다. 이렇게 함으로써 기능 버그가 도입되거나 후속으로 삭제된 버전 또는 훼손된 계정에서 릴리즈된 버전을 피할 수 있습니다.

## 자동 종속성 업그레이드 PRs 옵션 활성화 방법

Snyk를 구성하여 종속성 건강을 정기적으로 확인하고 종속성 업그레이드를 권장하며 업그레이드에 대한 PR을 자동으로 제출할 수 있습니다. 구성 후, Snyk는 스캔된 프로젝트에 대한 종속성이 사용 가능해지면 모든 필요한 종속성에 대해 자동으로 PR을 생성합니다.

기본적으로 프로젝트 설정은 조직 설정을 상속합니다. 그러나 조직 및 프로젝트 수준의 설정이 다른 경우 프로젝트 설정이 조직 설정을 재정의합니다.

{% hint style="info" %}
자동 종속성 업그레이드 PR은 다음 SCM 통합에만 사용할 수 있습니다: GitHub, GitHub Enterprise 및 Bitbucket Cloud.
{% endhint %}

### 조직 전체에 대한 자동 종속성 업그레이드 PRs 옵션 활성화 방법

조직 전체에 대한 자동 업그레이드 PR을 구성하려면 다음 단계를 따르십시오:

1. Snyk 웹 UI에서 필요한 조직을 엽니다.
2. **Settings > Organization Settings > Integrations**으로 이동하여 구성된 SCM을 찾고 해당 통합의 행 맨 끝에 있는 **Edit Settings**를 클릭합니다.

<figure><img src="../../../../.gitbook/assets/image (436).png" alt="통합 설정 편집"><figcaption><p>통합 설정 편집</p></figcaption></figure>

3. 선택한 통합의 **Settings** 페이지에서 **Automatic dependency upgrade PRs** 섹션으로 이동합니다.
4. 이 섹션에서 다음 작업을 수행합니다:
   * 슬라이더 - **Enable**로 변경.
   * **최대 열린 업그레이드 PR 수** – 프로젝트가 받을 수 있는 열린 Snyk PR 수를 정의합니다; 최대 10개입니다. 열린 PR 수 한도에 도달하면 새로운 업그레이드 PR은 생성되지 않습니다.
   * **Major version을 업그레이드 추천에 포함** – 권장 사항에서 주 버전 업그레이드를 포함할지 여부를 선택합니다. 기본적으로 추천 사항에는 패치 및 소버전만 포함됩니다.
   * **무시할 종속성** – **자동 업그레이드** 작업에 포함되지 말아야 하는 종속성의 정확한 이름을 입력합니다. 소문자만 입력할 수 있습니다.

<figure><img src="../../../../.gitbook/assets/image (437).png" alt="자동 종속성 업그레이드 PFs 활성화"><figcaption><p>자동 종속성 업그레이드 PFs 활성화</p></figcaption></figure>

5. 변경 사항을 저장 및 적용하려면 다음 중 하나를 **Save** 드롭다운에서 선택하세요:
   * **Save** – 변경 사항이 저장되어 조직에서 이 설정을 상속받은 모든 프로젝트에 적용됩니다. 사용자 정의 설정이 있는 프로젝트는 해당 변경사항에 영향받지 않습니다.
   * **Save changes and apply to all overridden Projects** – 변경 사항이 저장되어 조직의 모든 프로젝트에 적용됩니다. 사용자 정의 설정을 가진 프로젝트는 조직 설정을 상속하고 사용자 정의 설정이 무시됩니다.

이제부터 Snyk가 조직의 모든 프로젝트를 스캔할 때 새 업그레이드가 가능하면 자동으로 Upgrade PR을 제출합니다.

기존 Snyk Upgrade PR이나 Fix PR에 대해 더 높은 버전이 릴리스되면 기존 PR을 닫거나 병합해야만 Snyk가 새 PR을 만들 수 있습니다.

### 특정 프로젝트에 대한 자동 종속성 업그레이드 PRs 옵션 활성화 방법

프로젝트 수준의 설정이 조직 수준의 설정을 재정의합니다. 그러나 프로젝트 수준의 설정을 변경한 후 **모든 재정의된 프로젝트에 적용** 옵션으로 저장된 경우 조직 수준의 설정이 프로젝트의 사용자 정의 설정을 무시할 수 있습니다.

특정 프로젝트에 대한 자동 업그레이드 PR을 구성하려면 다음 단계를 따르세요:

1. Snyk 웹 UI에서 프로젝트가 포함된 조직을 엽니다.
2. 프로젝트 목록에서 자동 업그레이드 PR를 활성화하려는 **프로젝트**를 찾아 확장합니다.
3. 프로젝트 행 끝에 있는 **Project settings**을 클릭합니다.

<figure><img src="../../../../.gitbook/assets/image (438).png" alt="프로젝트 행의 프로젝트 설정"><figcaption><p>프로젝트 행의 프로젝트 설정</p></figcaption></figure>

4. **프로젝트 Settings** 페이지에서 사용 중인 통합을 선택합니다.

<figure><img src="../../../../.gitbook/assets/image (439).png" alt="프로젝트 설정 - 사용 중인 통합을 선택하세요"><figcaption><p>프로젝트 설정 - 사용 중인 통합을 선택하세요</p></figcaption></figure>

5. **Integration** 페이지로 스크롤하여 **자동 종속성 업그레이드 pull requests** 섹션으로 이동하고 다음 중 하나를 선택하세요:
   * **통합 설정에서 상속** – 조직의 통합 설정을 선택한 프로젝트에 적용합니다.\
     **조직에 대한 자동 종속성 업그레이드 PR 설정이 비활성화**된 경우 이 옵션은 프로젝트에도 비활성화됩니다.
   * **이 프로젝트에만 사용자 정의** – 프로젝트에 **자동 종속성 업그레이드 PR** 옵션의 특정 설정을 적용합니다. 이 경우:
     * 슬라이더를 **Enabled**로 변경합니다.
     * **Include major version in upgrade recommendation**에서 사용 가능한 옵션 중 하나를 선택하여 주 버전 업그레이드를 권장할지 여부를 정의합니다.\
       기본적으로 추천 사항에는 패치 및 소버전만 포함됩니다.
     * **Limit Snyk to this many dependency upgrade PRs open simultaneously**에서 프로젝트가 의존성 업그레이드 PR도 받기 위해 가질 수 있는 열린 Snyk PR 수를 정의합니다. 1부터 10까지의 숫자를 설정할 수 있습니다.\
       열린 PR 수 한도에 도달하면 새로운 업그레이드 PR은 생성되지 않습니다.\
       기본적으로 의존성 업그레이드 PR도 받기 위해서는 프로젝트에서 열린 PR이 최대 4개여야 합니다.
     * **Dependencies to ignore**에서 **자동 업그레이드** 작업에서 제외해야 하는 종속성의 정확한 이름을 입력합니다.\
       소문자만 입력할 수 있습니다.
     * 변경 사항을 저장하려면 **Update dependency upgrade settings**를 클릭합니다.

<figure><img src="../../../../.gitbook/assets/image (440).png" alt="프로젝트 수준의 자동 종속성 업그레이드 pull requests 설정"><figcaption><p>프로젝트 수준의 자동 종속성 업그레이드 pull requests 설정</p></figcaption></figure>

이러한 단계를 완료하면 Snyk가 프로젝트를 스캔하고 업그레이드가 가능하면 자동으로 Upgrade PR을 제출합니다. 기존 Snyk Upgrade PR이나 Fix PR에 대해 더 높은 버전이 릴리스되면 기존 PR을 닫거나 병합해야만 Snyk가 새 PR을 만들 수 있습니다.