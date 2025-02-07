# 타사 종속성 검사(SCA, Snyk 오픈 소스)

Eclipse 플러그인 버전 2.0.0 및 이후에서는 Snyk이 Eclipse의 원본 흐름과 보다 깊은 통합, 인라인 강조, 문제 통합, 그리고 호버로 문제에 대한 정보를 소개합니다. 다음에서는 서드파티 종속성에서 발견된 보안 취약점에 대한 이러한 모든 정보를 보여줍니다:

1. 취약한 패키지가 강조표시됩니다 (빨간색 squiggly line), 이 패키지에 심각한 보안 취약성이 있음을 나타내줍니다. 마우스오버하면 모든 정보가 표시되며, 스크롤하여 추가 정보를 확인하거나 링크를 클릭할 수 있습니다. 취약성이 있는 곳 바로 옆에 취해야 할 조치와 방법이 제시됩니다.
2. **문제** 뷰와의 통합을 볼 수 있으며, **문제** 뷰를 사용하여 문제를 필터링하고 그룹화하는 데 유용합니다. Snyk은 또한 문제가 있는 라인을 표시하며, 문제 뷰에서 문제를 클릭하면 해당 위치로 이동합니다.
3. 왼쪽에 있는 기둥 아이콘과 오른쪽에 있는 파일 맵 하이라이트 (우선순위에 맞는 색상과 일치)을 확인할 수 있습니다.

{% hint style="info" %}
호버 정보는 JavaEditor 및 GenericEditor에만 제한되며, 이는 Wild Web Developer와 같은 플러그인의 기본 편집기입니다.
{% endhint %}

<figure><img src="../../../.gitbook/assets/image (267) (1) (1) (1).png" alt=""><figcaption></figcaption></figure>

## **컨텍스트 메뉴**

우클릭 메뉴 옵션은 다음을 포함합니다:

**이슈 무시** - 다음 30일 동안 무시할 특정 이슈에 마우스를 올리고 컨텍스트 메뉴에 액세스합니다.

**Snyk 테스트** - 전체 워크스페이스에 대해 Snyk 테스트 실행.

**환경설정** - Snyk Vuln 스캐너 환경 설정에 직접 액세스하고 업데이트합니다.

## **축소된 Snyk 보기**

**제목:** 프로젝트의 이름.

**종속성:** 각 프로젝트에 대해 발견된 취약점 및 영향 받는 경로 수에 대한 요약.

## **확장된 Snyk 보기**

**제목:** 프로젝트에 영향을 미치는 취약점의 전체 이름으로, Snyk 데이터베이스에 있는 취약점의 설명 및 완전한 세부 정보와 링크됩니다.

**종속성:** 프로젝트의 직접 종속성 패키지의 이름 (명시적으로 설치한 패키지)이 취약점에 영향을 받는 것이 직접 또는 간접적인지를 나타냅니다.

모든 세부 정보가 단일 행에 나타나며, 종속성 (코드에서 명시적으로 사용된 패키지의 이름) 및 패키지 (실제로 취약점이 포함된 패키지의 이름) 열은 동일한 패키지의 이름을 나타냅니다:

{% embed url="https://files.gitbook.com/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-MdwVZ6HOZriajCf5nXH%2Fuploads%2Fgit-blob-4cdd086d6be47b598fc1a9a52c63023d59cff825%2Fuuid-e7accdc1-7495-e7a5-7a64-2403b066cb03-en.png?alt=media&token=e3bf024a-ba92-4b76-87be-b728d7edf092" %}
Eclipse 결과 세부사항
{% endembed %}

화살표가 나타나며 해당하는 모든 세부 정보를 그룹화하여 다음 예제와 유사하게 표시됩니다:

{% embed url="https://files.gitbook.com/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-MdwVZ6HOZriajCf5nXH%2Fuploads%2Fgit-blob-85e429be9a965c2dc534817a648773176a724531%2Fuuid-c71f67d1-80a3-7485-b33b-e602a1a5050e-en.png?alt=media&token=99e95293-bb37-4fed-8388-d9cb56a73092" %}
Eclipse 결과 행에 나타나는 화살표
{% endembed %}

## 축소 및 확장 모드에서의 예제

다음은 **축소된 모드**에서 종속성이 표시되는 것을 보여줍니다. 프로젝트가 간접적인 취약점에 영향을 받는 경우:

{% embed url="https://files.gitbook.com/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-MdwVZ6HOZriajCf5nXH%2Fuploads%2Fgit-blob-85e429be9a965c2dc534817a648773176a724531%2Fuuid-c71f67d1-80a3-7485-b33b-e602a1a5050e-en.png?alt=media&token=99e95293-bb37-4fed-8388-d9cb56a73092" %}
축소된 모드, 간접 취약점
{% endembed %}

이 예에서:

패키지 X는 패키지 Y를 사용하며, 패키지 Y는 패키지 Z를 사용합니다.

패키지 Z에는 교차 사이트 스크립트(XSS) 취약점이 포함돼 있으며, 이것이 간접적으로 프로젝트에 영향을 줍니다.

종속성 (코드에서 명시적으로 사용한 패키지의 이름)은 패키지 X이며, 패키지 필드에는 패키지 Z(실제로 취약점이 포함된 패키지의 이름)가 표시됩니다.

다음은 **확장된 모드**에서 종속성이 표시되는 방식을 보여줍니다. 프로젝트가 간접적인 취약점에 영향을 받는 경우:

행에서 화살표를 클릭하여 직접 종속성부터 취약한 패키지까지의 전체 경로를 확장하여 볼 수 있습니다.

{% embed url="https://files.gitbook.com/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-MdwVZ6HOZriajCf5nXH%2Fuploads%2Fgit-blob-992b169b89e7f3c45782fdeb47b205e3c0a95af8%2Fuuid-35658aaf-3359-80c2-c094-41a34c7863cc-en.png?alt=media&token=53c91ccc-f9bc-4ba7-a55f-8def3aa50d86" %}
확장된 모드, 간접 취약점
{% endembed %}

이전 화면에서 전체 경로가 나타납니다:

\[패키지 X의 이름]-->\[패키지 Y의 이름]-->\[패키지 Z의 이름]

**패키지:** 취약점에 직접 영향을 받는 프로젝트의 패키지 이름. 이전 화면에서:

* 종속성은 패키지 X로 표시됩니다 - 개발자가 코드에서 명시적으로 사용하는 패키지입니다
* 패키지 필드에는 취약점이 포함된 패키지인 패키지 Z가 표시됩니다.

**수정:** 문제 해결을 위해 업그레이드할 수 있는 패키지의 이름 및 버전.

## Snyk 결과 패널이 보이지 않는 경우

에이브러포트 close the Snyk Results panel by accident, or for some reason you do not see it, you can enable it as follows:

Windows -> Show View -> Other...을 차례로 탐색하십시오.

![Show View, Other](<../../../.gitbook/assets/Screenshot 2022-05-13 at 12.04.07.png>)

**Show View** 대화상자 창에서 Snyk를 검색하십시오.

![Show View dialog window](<../../../.gitbook/assets/Screenshot 2022-05-13 at 12.02.06 (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2) (4).png>)

이제 Snyk Results 패널을 볼 수 있어야 합니다:

![Snyk Results panel](<../../../.gitbook/assets/Screenshot 2022-05-13 at 12.02.18.png>)
