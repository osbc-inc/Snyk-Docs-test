# SCM 통합을 위한 배포 권장사항

## 권장 배포 순서

Snyk를 회사 구조에 원활하게 배포하려면 계획된 롤아웃이 필요합니다. 아래 단계를 따라 한 번에 통합 단계를 줄여 소프트웨어 개발 라이프 사이클(SDLC)에서 마찰을 발생시키지 않도록 조치하세요.

| 배포 단계                                                                                                                                                                                                                                                                           | 구성 및 테스트                                                                                                                                                                                                                                                                                                                                                                                                     | 결과                                                                                                                                                                                                                                                                                                                                                            |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <p><a href="deployment-recommendations-for-scm-integrations.md#stage-1-set-up-your-scm-integration"><strong>단계 1:</strong></a><br>첫 번째 Snyk 조직에서 조직 수준의 SCM 통합 설정.<br><br>Snyk AppRisk에서 그룹 수준의 SCM 통합 설정.</p>                                                   | <p>해당 SCM 통합에 대한 <a href="../#organization-level-snyk-scm-integrations">설치 문서</a>를 참조하고, <a href="../#user-permissions-and-access-scope-requirements">사용자 권한 및 액세스 범위 요구 사항</a> 섹션을 읽어서 회사 구조 내 올바른 사용자에게 접근 가능한지 확인하세요.<br><br>적용 가능한 AppRisk SCM 통합에 대한 <a href="../#group-level-snyk-apprisk-scm-integrations">설치 문서</a>를 참조하세요.</p> | <p>구성된 통합으로 프로젝트를 가져오기 준비.<br><br>AppRisk는 Git 저장소를 감지하여 Snyk 스캔을 롤아웃 진행상황을 추적할 수 있습니다.</p>                                                                                                                                                                                         |
| <p><a href="deployment-recommendations-for-scm-integrations.md#stage-2-import-projects"><strong>단계 2:</strong></a><br>통합된 SCM에서 모든 관련 프로젝트 가져오기.</p>                                                                                          | <p>• 공개 및 비공개 repo 테스트</p><p>• 모든 사용자 알림 비활성화</p><p>• Snyk PR 상태 확인 비활성화</p><p>• 자동 수정 PR 비활성화</p><p>• 자동 종속성 업그레이드 PR 비활성화</p><p>• Dockerfile 기본 이미지 업데이트 비활성화</p>                                                                                                                                                                                                         | <p>Snyk는 가져온 저장소를 스캔하여 모든 지원되는 파일을 감지하고 해당 프로젝트로 생성합니다. 이러한 기능을 비활성화하면 초기 설정과 사용에 잠재적인 방해 요소를 제거하여 예정된 스캔과 보고서의 높은 가시성을 통해 배포 차단 없이 진행률을 볼 수 있습니다.</p> |
| <p><a href="deployment-recommendations-for-scm-integrations.md#stage-3-enable-snyk-test-on-prs"><strong>단계 3:</strong></a><br>Snyk 우선 순위 보고 능력을 활용하여 고정 계획 구축.</p>                                                                             | [Snyk Priority Score](../../../manage-risk/prioritize-issues-for-fixing/priority-score.md) 및 [Snyk Risk Score](../../../manage-risk/prioritize-issues-for-fixing/risk-score.md)에 대해 자세히 알아보고 어떻게 이러한 기능을 활용하여 수정해야 할 문제의 우선 순위를 결정하는 데 도움을 받을 수 있는지 확인하세요.                                                                                                                                                                                          | 개발자와 보안 간의 정렬으로 어떤 문제를 해결해야 하는지 및 어떤 상황에서 언제 문제를 해결해야 하는지에 대한 결정을 내리고 수정 프로세스를 최적화합니다.                                                                                                                                                                                                                      |
| <p><a href="deployment-recommendations-for-scm-integrations.md#stage-4-enable-blocking-prs"><strong>단계 4:</strong></a><br>Snyk PR 확인 및 자동 수정 PR 설정을 구성하여 사용 가능한 해결 방법으로 실시간으로 문제를 개발자에게 알립니다.</p>                                        | [Snyk PR 상태 확인](../../../scan-with-snyk/pull-requests/)을 활성화하여 개발자가 PR에 높은 심각도 문제와 해결방법을 포함하는지 여부를 알 수 있도록 합니다.                                                                                                                                                                                                                                                                                                                   | 개발자가 저장소에 병합되기 전에 새로운 문제를 쉽게 볼 수 있도록하여 새로운 취약성의 추가 수를 줄이고 병합을 차단할지 여부에 대한 결정을 내릴 수 있도록 합니다.                                                                                                                                                                                                               |
| <p><a href="deployment-recommendations-for-scm-integrations.md#stage-5-automatic-fix-prs"><strong>단계 5:</strong></a><br>개발자가 새로운 취약성을 도입하지 못하도록 합니다.</p>                                                                                         | SCM 도구에서 브랜치 보호를 활성화하여 실패한 Snyk PR 상태 확인이 존재하는 PR이 병합되지 않도록 합니다.                                                                                                                                                                                                                                                                                                                                                                            | 이 변경으로 Snyk가 새로운 취약성을 감지하면 저장소로 병합되지 않도록 합니다. 시간이 지나면 픽스 가능 여부를 고려하지 않고 모든 심각도가 높거나 중요한 문제를 포함하는 PR에 민감도를 조정할 수 있습니다.                                                                                                                                     |
| <p><a href="deployment-recommendations-for-scm-integrations.md#stage-6-dependency-upgrade-prs"><strong>단계 6:</strong></a><br>Snyk 자동 기능을 활성화하여 문제 해결 및 의존성을 최신 상태로 유지하여 기술적 보안 빚을 감소시킵니다.</p>                                | <p><a href="https://docs.snyk.io/scan-using-snyk/pull-requests/snyk-fix-pull-or-merge-requests#automated-snyk-prs">자동 수정 PR</a>를 활성화합니다.<br><a href="https://docs.snyk.io/scan-using-snyk/pull-requests/snyk-fix-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs">자동 종속성 업그레이드</a>를 활성화합니다.</p>                                                                                                                                      | Snyk는 심각도가 가장 높은 문제를 수정하기 위해 자동으로 PR을 만듭니다. 의존성을 최신으로 유지하면 미래의 디자인 문제와 긴급 수정을 줄일 수 있으며, 이를 조사하고 해결하는 데 필요한 시간이 소요될 수 있습니다.                                                                                                                                      |

## 단계 1: 조직 및 그룹 수준 SCM 통합 설정

Snyk는 GitHub, GitHub Enterprise, Bitbucket Cloud 등을 포함한 조직 수준 SCM 통합 기능을 제공합니다.

세부 정보는 [Snyk 통합 설정](../../../getting-started/#set-up-a-snyk-integration)을 참조하십시오.

Snyk AppRisk를 위한 그룹 수준 SCM 통합 기능도 제공됩니다. GitHub, GitLab, Azure DevOps, Bitbucket 등이 이에 해당합니다.

자세한 내용은 [그룹 수준 Snyk AppRisk SCM 통합](../#group-level-snyk-apprisk-scm-integrations)을 참조하십시오.

### **저장소의 SCM 권한**

Snyk UI를 사용하여 시작된 작업(예: Fix PR 열기 또는 프로젝트 재테스트)은 작업 수행 사용자를 위해 수행됩니다. 따라서 이러한 작업을 수행하려면 자체 SCM 사용자나 서비스 계정과 연결해야 합니다. 이렇게 하면 Snyk가 이러한 작업을 수행할 저장소에 필요한 권한을 부여받게 됩니다.

이러한 권한에 대한 자세한 내용은 선택한 SCM 통합에 대한 [사용자 권한 및 액세스 범위 요구 사항](../#user-permissions-and-access-scope-requirements)을 확인하십시오.

### **알림 설정 변경**

기본적으로 Snyk는 Project의 종속성에 새로운 문제 또는 수정 사항을 발견할 때마다 모든 Org 사용자에게 이메일을 보내며, 조직 전체의 보안 상태에 대한 주간 업데이트 제공합니다. Org에 여러 프로젝트를 가져오기 위해 많은 프로젝트를 계획 중이라면 사용자에게 여러 전자 메일 알림이 보내지지 않도록하기 위해 해당 조직의 모든 알림을 비활성화하는 것을 고려하십시오.

사용자가 받는 이메일을 사용자 지정하려면 **Settings > Notifications > Notification Setting**으로 이동하십시오. 여기서 하는 변경 사항은 조직의 모든 구성원에 영향을 미치지만 구성원은 자신의 사용자 계정 설정에서 이러한 기본 설정을 재정의할 수 있습니다.

가져올 알림을 비활성화하려면 Org의 모든 사용자에 대한 알림을 해제하려면 해당 알림 상자를 선택 취소하십시오. 자세한 내용은 [알림 관리](../../../snyk-admin/manage-notifications.md)를 참조하십시오.

## 단계 2: 프로젝트 가져오기

Snyk UI의 **Projects** 페이지로 이동하고 **Add projects**를 선택한 다음 Snyk로 가져올 저장소를 선택하고 **Add selected repositories**를 클릭합니다.

* Snyk는 선택한 저장소에서 의존성 파일(예: **package.json**)을 스캔하기 시작하고 전체 디렉토리 트리에서 이 파일을 프로젝트로 가져옵니다.
* Snyk는 루트 폴더와 정의된 사용자 정의 파일 위치를 평가합니다. 매니페스트 또는 구성 파일이 없는 경우, Snyk는 가져올 파일이 없음을 알립니다.
* Snyk는 매니페스트 파일(프로젝트)을 감지하고 테스트한 다음 결과를 표시합니다.\
  가져온 프로젝트는 저장소 이름 아래에 나타납니다.\
  프로젝트가 가져온 후에는 계속해서 취약성을 확인합니다.

{% hint style="info" %}
프로젝트가 가져온 것을 확인하려면 통합에 대한 **Add project** 가져오기 페이지로 이동하십시오. 가져온 프로젝트는 저장소 이름 옆에 ✔로 표시됩니다.
{% endhint %}

세부 정보는 [프로젝트 가져오기](../../../getting-started/#import-a-project-to-scan-and-identify-issues)를 참조하십시오.

## 단계 3: Snyk PR에서 테스트 활성화

### **PR 테스트 설정 및 워크플로우**

기본적으로 Snyk는 감시 중인 저장소에 제출된 모든 풀 리퀘스트를 스캔하고 결과 및 추천 사항을 하나의 보안 확인과 라이선스 확인으로 묶어서 표시합니다.

### **상태 세부 정보**

**Details** 링크를 클릭하여 Snyk 검사의 상태를 표시합니다. 상태 옵션은 다음과 같습니다:

* **Success**: 문제가 식별되지 않았으며 모든 확인 항목이 통과함
* **Processing**: Snyk 테스트가 완료될 때까지 표시되는 상태
* **Failure**: 해결하지 않은 문제로 확인 항목이 통과되지 않은 상태
* **Error**: 다음과 같은 문제가 발생했음을 나타냅니다:
  * Manifest 파일이 동기화되지 않을 수 있음
  * Snyk가 매니페스트 파일을 읽지 못함
  * Snyk가 매니페스트 파일을 찾지 못함

### **PR 확인 설정 관리**

관리자는 각 SCM 통합에 대한 Snyk [PR Checks](../../../scan-with-snyk/pull-requests/pull-request-checks/) 설정을 조직 수준에서 관리한 다음 이러한 설정을 적용하여 해당 통합의 모든 프로젝트 또는 선택한 특정 프로젝트에 설정할 수 있습니다. 이 기능이 활성화되는지 여부를 구성하고 Snyk가 PR 확인을 실패시키는 조건을 설정할 수 있습니다.

통합 수준의 PR Check 설정([통합 수준에서 PR Checks 구성](../../../scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md#configure-pr-checks-at-the-integration-level)) 및 프로젝트 수준의 PR Check 설정([프로젝트 수준에서 PR Checks 구성](../../../scan-with-snyk/pull-requests/pull-request-checks/configure-pull-request-checks.md#configure-pr-checks-at-the-project-level))에 대한 자세한 내용은 이 프로세스의 세부 정보를 확인합니다. &#x20;

{% hint style="info" %}
Snyk 라이선스 정책을 사용하여 Snyk PR에 라이선스 문제가 없는지 확인할 수 있습니다.\
자세한 내용은 [Licenses](../../../scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/open-source-license-compliance.md)를 참조하십시오.
{% endhint %}

### **초기 단계: 가시성 확보 및 실패 조건 설정**

SCM 통합을 처음 롤아웃할 때 Snyk는 개발자가 Snyk 커밋 확인을 보게 만들기 위해 PR을 실패시키지 않고 PR을 테스트하도록 Snyk가 설정되기를 권장합니다.

1. 통합을 위한 PR 테스트를 조직 수준에서 또는 특정 프로젝트에 적용할 것인지 결정합니다.
2. PR Check 설정 관리란에서 설명된대로 실패 조건을 설정합니다([PR 확인 설정 관리](deployment-recommendations-for-scm-integrations.md#manage-pr-check-settings)에서 설명).

## 단계 4: PR 차단 활성화

Snyk을 소프트웨어 개발 생애 주기(SDLC)에 통합하고 개발자 인식이 잘 구축된 후, 보안 태세를 개선하기 위해 더 엄격한 정책을 적용할 수 있습니다. 예를 들어:

* **낮은 우선순위 프로젝트**: 수정 가능한 새로운 고위험 문제에 대해서만 PR을 실패시킬 수 있습니다.
* **중간 우선순위 프로젝트**: 고위험 문제에 대해서만 PR을 실패시킬 수 있습니다.
* **높은 우선순위 프로젝트(PCI/GDPR 준수)**: 모든 문제에 대해 PR을 실패시킬 수 있습니다.

{% hint style="info" %}
취약점 심각도를 내부 정책에 맞추려면 보안 정책을 사용하여 문제의 심각도를 변경하고 이를 관련 프로젝트 속성에 연결하세요. 자세한 내용은 [보안 정책](../../../manage-risk/policies/security-policies/)을 참조하세요.
{% endhint %}

## 단계 5: 자동 수정 PR

Snyk은 프로젝트를 매일 또는 매주 스캔합니다. 새로운 취약점이 발견되면 Snyk은 이메일을 보내고 수정 사항이 포함된 자동 PR을 리포지토리에 생성합니다.

조직의 모든 프로젝트에 대해 자동 수정 PR 설정을 구성하려면, 해당 조직을 선택하고 **조직 설정** > **통합 > 설정 편집**으로 이동하십시오.

<div align="center"><figure><img src="../../../.gitbook/assets/image (4) (1) (1) (1).png" alt="자동 수정 풀 요청 설정" width="563"><figcaption><p>자동 수정 풀 요청 설정</p></figcaption></figure></div>

{% hint style="info" %}
설정은 또한 조직 내 특정 프로젝트를 선택하고 **설정** 탭으로 이동하여 프로젝트별로 구성할 수 있습니다. 자세한 내용은 [프로젝트 설정 보기 및 편집](https://docs.snyk.io/snyk-admin/snyk-projects/view-and-edit-project-settings)을 참조하세요.
{% endhint %}

Snyk은 개발자가 자동 수정 PR을 사용하는 방법에 익숙하지 않다면 해당 패치를 자동 수정 PR에서 제외할 것을 권장합니다.

개발자에게 자동 수정 PR에 표시된 병합 권장 사항 레이블을 고려하도록 요청하세요.

{% hint style="info" %}
Snyk의 자동 수정 PR은 새로운 문제에 대해서만 생성됩니다.
{% endhint %}

만약 SCM이 GitHub이고 Snyk Broker를 사용하지 않는 경우, 기본적으로 Snyk은 모든 조직 사용자의 자격 증명을 회전시켜 자동 수정 PR을 엽니다. 필요에 따라 이를 변경하고 사용자 자격 증명을 설정하여 자동 수정 PR을 열 수 있습니다. 자세한 내용은 [고정된 GitHub 계정에서 수정 및 업그레이드 풀 요청 열기](../../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/opening-fix-and-upgrade-pull-requests-from-a-fixed-github-account.md)를 참조하세요.

## 단계 6 - 의존성 업그레이드 PR

그룹이 보안 기술 부채 해결을 시작할 준비가 되면, Snyk을 구성하여 의존성을 업그레이드하는 풀 요청(PR)을 자동으로 생성하도록 할 수 있습니다.

### **자동 업그레이드 PR 작동 방식**

1. 프로젝트 통합이 구성되고 사용자가 자동 업그레이드 PR을 활성화합니다.
2. Snyk은 프로젝트를 가져오면서 스캔을 수행하고 프로젝트를 정기적으로 모니터링합니다.
3. 각 스캔에서 Snyk이 의존성의 새 버전을 식별하면:
   * Snyk은 Snyk 프로젝트 설정에 따라 주기적으로 자동 업그레이드 PR을 생성합니다.
   * Snyk은 다른 열린 Snyk PR에서 이미 변경된(업그레이드되거나 패치된) 의존성에 대해 새로운 업그레이드 PR을 열지 않습니다.
   * Snyk은 각 의존성에 대해 별도의 PR을 엽니다.
   * Snyk은 열려 있는 Snyk PR이 5개 이상인 리포지토리에서는 업그레이드 PR을 생성하지 않습니다. 열린 PR의 제한이 도달하면 새로운 PR이 생성되지 않습니다. 이 숫자는 설정에서 1에서 10 사이로 설정할 수 있습니다. 이 제한은 업그레이드 PR을 생성할 때만 적용되며, 수정 PR에는 적용되지 않습니다.
   * 기본적으로 Snyk은 패치 및 마이너 업그레이드만 추천하지만, 주요 버전 업그레이드는 설정에서 이 기능을 활성화하면 가능합니다.
   * 최신 버전이 프로젝트에서 이미 발견된 취약점을 포함하고 있다면 Snyk은 업그레이드를 추천하지 않습니다.
   * Snyk은 21일 이하의 최신 버전에 대해 업그레이드를 추천하지 않습니다. 이는 기능적 버그를 유발한 후 배포되지 않거나, 악의적인 의도를 가진 사람이 제어권을 잃은 계정에서 배포된 버전을 방지하기 위함입니다.

### **지원되는 언어 및 리포지토리**

Snyk은 GitHub, GitHub Enterprise Server 및 BitBucket Cloud를 통해 npm, Yarn 및 Maven-Central 프로젝트에 대해 자동 업그레이드 PR을 지원합니다. Snyk Broker와 함께 사용이 지원됩니다. Broker와 함께 사용하려면 관리자가 먼저 v4.55.0 이상으로 업그레이드해야 합니다.

### **특정 프로젝트에 대해 자동 의존성 업그레이드 PR 활성화**

조직 수준에서 구성된 PR 설정을 재정의하고 프로젝트 수준에서 PR 설정을 설정하려면:

1. 자동 업그레이드 PR을 활성화하려는 조직을 열고 **프로젝트** 탭으로 이동합니다.
2. 관련 프로젝트를 선택하고 확장한 후, 관련 대상을 선택하고 톱니바퀴 아이콘을 사용하여 **설정**을 엽니다.
3. 설정 영역에서 왼쪽 패널 메뉴에서 통합 설정을 클릭하여 해당 프로젝트에 대해 고유한 설정을 적용합니다.
4. 로드된 설정에서 **자동 의존성 업그레이드 풀 요청**으로 스크롤하여 **사용 안 함**을 클릭합니다.
5. 나타나는 옵션에서:
   1. Snyk은 리포지토리당 최대 10개의 열린 PR을 생성합니다. 이 숫자를 더 제한하려면 드롭다운 목록에서 최대 PR 수를 선택하십시오. 자세한 내용은 [자동 PR을 통한 의존성 업그레이드](../../../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs-upgrade-prs/)를 참조하세요.
   2. **무시할 의존성** 필드에 자동 기능의 일부로 처리되지 않아야 할 의존성의 정확한 이름을 입력합니다. 이 필드는 소문자만 허용합니다.
   3. **의존성 업그레이드 설정**을 클릭한 후, Snyk이 이 프로젝트를 스캔할 때마다 Snyk은 스캔 결과에 따라 자동으로 업그레이드 PR을 제출합니다. 기존 Snyk 업그레이드 PR이나 수정 PR에 대해 새 버전이 릴리스되면, 새로운 PR을 생성하기 전에 기존 PR이 닫히거나 병합되어야 합니다.