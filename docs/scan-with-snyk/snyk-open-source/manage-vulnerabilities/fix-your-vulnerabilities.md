# 취약점 수정

Snyk은 직접 종속성을 취약성이 없는 버전으로 업그레이드하거나 취약성을 패치하여 취약성을 수정하는 데 도움을 줍니다. Snyk이 프로젝트를 스캔한 후, 스캔 결과를 통해 명확한 제안과 설명을 통해 코드의 문제를 해결할 수 있습니다.

Snyk 오픈 소스를 사용하면 다음을 수행할 수 있습니다:

* [Snyk 웹 사용자 인터페이스에서 스캔 결과보기](fix-your-vulnerabilities.md#view-scan-results-on-the-snyk-web-ui)
* [Snyk CLI를 사용하여 스캔 결과 기반 취약점 수정](fix-your-vulnerabilities.md#fixing-vulnerabilities-based-on-scan-results-using-snyk-cli)
* [수정 적용](fix-your-vulnerabilities.md#apply-fixes)

## Snyk 웹 사용자 인터페이스에서 스캔 결과보기

Snyk 오픈 소스에서 각 탭에 대해 업그레이드 및 패치가 권장되는 프로젝트 세부 정보의 수정 제안 영역에서 다음과 같이 결과가 표시됩니다:

* 수정 가능한 전체 패키지 수가 탭 제목에 표시됩니다.
* 패키지별로 취약점 그룹이 업그레이드 또는 권장 사항으로 표시됩니다.
* 패키지가 나열되고 해당 패키지에 영향을주는 전체 취약성 목록을 표시하도록 확장할 수 있습니다.
* 종속성에서 발견된 모든 취약성을 페이지 아래쪽에서 함께 표시하며 문제를 우선순위로 지정하고 필요한 경우 수정을 시작할 수 있도록 도와주는 문맥 정보가 제공됩니다.

<figure><img src="../../../.gitbook/assets/Screenshot 2023-03-15 at 12.14.06.png" alt="웹 UI의 스캔 결과"><figcaption><p>웹 UI의 스캔 결과</p></figcaption></figure>

## 수정 권고 사항 보기

고치기 권고 사항 영역은 프로젝트 세부 정보 페이지에 나타납니다. Snyk은 다음 중 하나의 솔루션을 제공합니다:

* 원래 패키지로 업그레이드
* 패키지를 고정화하고 취약한 버전을 직접 종속성으로 가져 오지 않도록 지정된 버전의 패키지를 설치.
* 문제를 해결하기 위한 Snyk 정밀 패치. 패키지의 취약점을 해결하기 위한 업그레이드가 현재 사용할 수없는 경우, Snyk은 문제를 해결하는 패치를 제공합니다.

요약 영역은 패키지별로 조언을 그룹화하며 최상의 수정 가능한 조언을 기준으로 표시됩니다. 이러한 요약 목록의 조언에는 다음과 같은 각 패키지에 대한 세부 정보가 포함됩니다:

* 해당 패키지에 영향을 주는 모든 취약성 이름 및 심각성 세부 정보
* 권장 수정,이 패키지 및 해당 목록된 취약성을위한 권장 수정 링크 : 업그레이드 할 구체적인 버전 또는 패치의 이름

<figure><img src="../../../.gitbook/assets/Screenshot 2021-10-12 at 14.08.13.png" alt="업그레이드 문제 탭"><figcaption><p>업그레이드 문제 탭</p></figcaption></figure>

<figure><img src="../../../.gitbook/assets/Screenshot 2021-10-12 at 14.10.00 (1).png" alt="업그레이드 가능 및 수정 가능한 문제 탭"><figcaption><p>업그레이드 가능하고 패치 가능한 문제 탭</p></figcaption></figure>

프로젝트 세부 정보 페이지에서 아래쪽에 추가 조언 및 세부 정보를 찾을 수도 있습니다:

* **Issues** 탭에서 취약성 당 전체 설명
* **Dependencies** 탭에서 프로젝트 종속성 트리 전체를 찾아 영향을 받는 경로를 명확하게 시각화 할 수 있습니다.

## Snyk CLI를 사용하여 스캔 결과 기반 취약점 수정

취약성 수정에 대한 자세한 내용은 [Snyk CLI 사용하여 취약성 수정](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/fix-vulnerabilities-using-the-snyk-cli.md)를 참조하십시오.

## 수정 적용

수정을 적용하려면 다음을 수행할 수 있습니다:

* 관련 프로젝트 페이지의 특정 [문제 카드](../../../snyk-admin/snyk-projects/issue-card-information.md)에서 **이 취약점 수정**을 클릭하십시오.
* [소스 코드 통합](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)을 사용 중이라면:
  * 프로젝트 페이지에서 **수정 PR 열기**를 클릭합니다.
  * 새로운 수정사항을 사용할 수 있을 때 [자동 풀 리퀘스트](../../pull-requests/snyk-pull-or-merge-requests/create-automatic-prs-for-new-fixes-fix-prs.md)를 사용하여 취약성을 수정하는 데 도움이 되는 새로운 수정사항이 있는지 확인하십시오.

{% hint style="info" %}
**자동 수정 PRs**\
새로운 수정 가능한 취약점이 발견되면 Snyk가 자동으로 새 풀 리퀘스트를 여는 시도할 수 있습니다. 세부 정보는 [새 수정사항을 위한 자동 풀 리퀘스트 생성](../../pull-requests/snyk-pull-or-merge-requests/create-automatic-prs-for-new-fixes-fix-prs.md)에서 확인할 수 있습니다.
{% endhint %}
