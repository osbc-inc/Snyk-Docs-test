# 백로그 이슈 및 알려진 취약점에 대한 자동 PR 생성(백로그 PR)

{% hint style="info" %}
**기능 가용성**

* **알려진 취약점을 위한 자동 PRs (백로그 PRs)** 기능은 다음 SCM 통합에서 지원됩니다: GitHub, GitHub Enterprise, GitHub Cloud App, Bitbucket Server, Bitbucket Cloud, Bitbucket Connect, GitLab 및 Azure Repos.
* **자동 Fix PR** 설정은 통합에 따라 다를 수 있습니다.
{% endhint %}

Snyk이 취약점에 대한 자동 PR을 생성할 때 다음 규칙이 적용됩니다:

* 프로젝트를 위해 **지금 재테스트**를 선택하면 수동으로 스캔이 실행되고 24시간 창이 스캔이 실행되었다고 표시됩니다. 다음 자동 스캔이 실행될 때까지 자동 PR이 생성되지 않습니다.
* **700 이상**의 [우선 순위 점수](../../../manage-risk/prioritize-issues-for-fixing/priority-score.md)를 가진 각 프로젝트당 하나의 풀 리퀘스트가 생성됩니다.
* 풀 리퀘스트는 **테스트 및 자동화된 풀 리퀘스트 빈도** 설정에 따라 생성됩니다.
  * **테스트 및 자동화된 풀 리퀘스트 빈도**를 업데이트하려면 **프로젝트**로 이동하고 열린 소스 프로젝트를 선택합니다.
  * **설정**으로 이동하고 드롭다운 목록에서 옵션을 선택합니다.

<figure><img src="../../../.gitbook/assets/Project testing and PR Checks frequency (1).png" alt="프로젝트 테스트 및 자동 PR 확인 빈도 설정"><figcaption><p>프로젝트 테스트 및 자동 PR 확인 빈도 설정</p></figcaption></figure>

마지막 24시간 창이 언제 시작되었는지 확인하려면 **반복 테스트에 의해 스냅샷 캡처**가 있는 프로젝트 이슈 카드를 확인하십시오.

<figure><img src="../../../.gitbook/assets/Test information with a focus on the latest snapshot taken.png" alt="반복 테스트에 의해 13시간 전에 캡처된 스냅샷"><figcaption><p>반복 테스트에 의해 13시간 전에 캡처된 스냅샷</p></figcaption></figure>

특정 스캔 결과를 확인하려면 **\[snyk] 취약성 경보**라는 제목의 이메일을 받은 편지함도 확인할 수 있습니다.

## 통합 수준에서 자동 Fix PR 구성

이미 Snyk와 통합된 특정 Git 리포지토리에 대해 자동 픽스 PR을 구성하려면 다음 단계를 따릅니다.

{% hint style="warning" %}
**자동 Fix PRs** 사용 시 더 큰 버전 이동이 발생할 수 있습니다.
{% endhint %}

이 구성 설정은 해당 조직의 모든 프로젝트에 적용됩니다. 또한 사용자 지정 설정을 가진 프로젝트로 구성을 확장할 수도 있습니다.

1. Snyk Web UI를 열고 **Settings > Integrations**로 이동합니다.
2. Git 리포지토리 통합 (SCM)을 선택합니다. 이 예제에서는 GitHub이 구성됩니다.
3. **Automatic Fix PRs** 아래에서 \*\*Known vulnerabilities (backlog)\*\*를 활성화합니다.\
   이는 프로젝트 백로그에서 이전에 선언된 취약점을 검색합니다.

<figure><img src="../../../.gitbook/assets/Automatic fix PRs settings for Git integration.png" alt="Git 통합을 위한 자동 픽스 PRs 설정."><figcaption><p>Git 통합을 위한 자동 픽스 PRs 설정</p></figcaption></figure>

4. 백로그 PRs를 위한 **Fix Strategy**를 선택합니다.
   * 기본적으로, 고정 전략은 취약점 수준의 단일 PR입니다. Snyk은 백로그의 문제마다 하루에 한 번씩 PR을 열어 가장 높은 취약점을 수정합니다.
   * **동일한 종속성의 모든 취약점을 단일 PR로 해결**을 선택할 수 있습니다. 이는 동일한 종속성 내의 가장 높은 우선 순위 취약점 및 해당 취약점을 해결하는 수정뿐만 아니라 다른 취약점에 대한 수정을 선택합니다.
5. **Save**를 클릭합니다.
6. (선택 사항) **모든 대체된 프로젝트에 적용하려면 변경 사항 저장 및 적용**을 선택합니다.\
   모든 프로젝트에 동일한 구성을 적용하려면 이 옵션을 사용하십시오. 이 옵션을 선택하면 모든 개별 프로젝트 설정이 **자동 fix PRs**에 대한 전역 설정으로 재정의됩니다. 자동 fix full requests에 대한 이전 프로젝트 설정을 갖고 있던 프로젝트에 대해 이 옵션을 선택하면 프로젝트 설정이 전역 설정으로 재정의됩니다.

<figure><img src="../../../.gitbook/assets/Fix all vulnerabilities for the same dependency in a single PR.png" alt="동일한 종속성의 모든 취약점을 해결하는 단일 PR."><figcaption><p>동일한 종속성의 모든 취약점을 해결하는 단일 PR</p></figcaption></figure>

## 프로젝트 수준에서 자동 Fix PR 구성

프로젝트가 전역 통합 설정에서 상속되는 대신 특정 프로젝트에서만 작동하도록 자동 Fix PR을 구성할 수 있습니다.

1. **Projects**로 이동하고 오픈 소스 프로젝트를 포함하는 [대상](../../../snyk-admin/snyk-projects/#target)을 확장합니다.
2. **Settings**로 이동하고 GitHub와 같은 통합을 선택합니다.
3. **Automatic fix pull requests** 섹션에서:
   * **이 프로젝트에만 맞춤화**를 선택합니다
   * **Known vulnerabilities (backlog)**&#xC744; 활성화합니다.
4. [통합 구성 설정](create-automatic-prs-for-backlog-issues-and-known-vulnerabilities-backlog-prs.md#configure-automatic-fix-prs-at-the-integration-level)에서 소개된 대로 Backlog PRs를 위한 **Fix Strategy**를 선택합니다.
5. **Save changes**를 클릭합니다.

<figure><img src="../../../.gitbook/assets/Automatic fix PRs settings at the Project level.png" alt="프로젝트 수준의 자동 Fix PRs 설정."><figcaption><p>프로젝트 수준의 자동 Fix PRs 설정</p></figcaption></figure>
