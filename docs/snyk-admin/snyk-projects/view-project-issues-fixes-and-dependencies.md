# 프로젝트 문제, 수정 사항 및 종속성 보기

다음과 같은 프로젝트 정보를 Snyk 웹 UI에서 확인할 수 있습니다:

* [Issues](view-project-issues-fixes-and-dependencies.md#view-issues): 취약점과 오픈 소스 라이센스 문제 수
* [Fixes](view-project-issues-fixes-and-dependencies.md#view-fixes): 수정 조언
* [Dependencies](view-project-issues-fixes-and-dependencies.md#view-dependencies): 오픈 소스의 직접 및 간접(중첩된) 종속성의 총 개수

## 문제 보기

프로젝트 세부정보 페이지는 **Issues** 탭에서 고취 값 카드를 표시합니다. 제공되는 정보에는 취약점 및 오픈 소스 라이센스 문제가 포함됩니다.

<figure><img src="../../.gitbook/assets/Screenshot 2021-10-19 at 11.49.30.png" alt="프로젝트 세부사항 Issues 탭 및 필터"><figcaption><p>프로젝트 세부사항 Issues 탭 및 필터</p></figcaption></figure>

좌측 패널의 필터를 사용하여 문제 검색을 좁힐 수 있습니다. **Issue type**, **Severity**, **Fixability**, **Exploit Maturity**, **Status**에 따라 문제를 필터링할 수 있는 확인란을 선택하세요. 또한 **Priority Score** 슬라이더를 편집하여 표시되는 범위를 변경할 수 있습니다. 기본값은 0부터 1000까지입니다.

중요도 점수로 정렬된 주요 영역의 문제 카드에 문제 세부사항이 표시됩니다. 자세한 내용은 [Issue card 정보](issue-card-information.md)를 참조하세요.

{% hint style="info" %}
Snyk는 스캔 중 발견된 문제를 수정하는 기능을 제공합니다. 자세한 내용은 [취약점 수정하기](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/fix-your-vulnerabilities.md)를 참조하세요.
{% endhint %}

## 문제 세부정보 보기

문제를 클릭하여 [중요도 점수](../../manage-risk/prioritize-issues-for-fixing/priority-score.md)를 포함한 세부정보를 확인할 수 있습니다.

<div align="left"><figure><img src="../../.gitbook/assets/Screenshot 2023-06-13 at 08.43.14.png" alt="문제 세부정보 보기"><figcaption><p>문제 세부정보 보기</p></figcaption></figure></div>

* [이 유형의 취약점에 대해 알아보기](../../getting-started/snyk-learn/)를 클릭하여 Snyk Learn 교육을 참조하세요.
* 보다 자세한 정보를 보려면 **더 많은 정보 표시**를 클릭하여 [Snyk 취약점 데이터베이스](https://snyk.io/product/vulnerability-database/)에서 취약성에 대한 자세한 정보를 확인하세요:

<div align="left"><figure><img src="../../.gitbook/assets/Screenshot 2023-06-13 at 08.47.54.png" alt="Snyk 취약점 데이터베이스의 추가 정보"><figcaption><p>Snyk 취약점 데이터베이스의 추가 정보</p></figcaption></figure></div>

## 수정 보기

프로젝트의 간접 종속성에 대한 Snyk의 지식으로 인해 **Fixes** 탭에서 수정 조언을 제공할 수 있습니다:

<div align="left"><figure><img src="../../.gitbook/assets/Screenshot 2021-10-19 at 11.57.07.png" alt="문제 세부사항 Fixes 탭"><figcaption><p>프로젝트 세부사항 Fixes 탭</p></figcaption></figure></div>

세부 정보는 [취약점 수정하기](../../scan-with-snyk/snyk-open-source/manage-vulnerabilities/fix-your-vulnerabilities.md)에서 확인할 수 있습니다.

## Snyk 오픈 소스에서 의존성 보기

Snyk은 응용프로그램의 패키지 매니저를 사용하여 의존성 트리를 구축하고 오픈 소스의 프로젝트 문제 상세 페이지의 **Dependencies** 탭에 표시합니다. 이 탭은 어떤 구성요소가 취약점을 도입했는지 보여 주며 해당 종속성이 어떻게 응용프로그램에 도입되었는지 나타냅니다.

다음은 예시입니다:

<div align="left"><figure><img src="../../.gitbook/assets/Screenshot 2023-06-13 at 08.57.23.png" alt="문제 상세 페이지 의존성 탭"><figcaption><p>문제 상세 페이지 의존성 탭</p></figcaption></figure></div>
