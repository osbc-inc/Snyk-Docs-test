# 분석 결과: 

{Snyk 오픈 소스 분석은 매 스캔마다 코드에서 발견된 취약점을 보여줍니다. 이 스캔은 백그라운드에서 실행되며 기본적으로 활성화되어 있습니다.

비주얼 스튜디오 코드 결과 화면의 **문제** 탭에서 프로젝트에서 발견된 모든 취약점을 확인할 수 있습니다.

##  편집기 창

편집기 창에서는 JavaScript, TypeScript 및 HTML로 코딩하는 동안 오픈 소스 모듈의 보안 취약점을 보여줍니다. 코드를 작성하는 동안, 가져오고 있는 모듈이 포함하는 취약점의 수 등과 같은 피드백을 코드 내에서 받을 수 있습니다. 편집기는 최상위 종속성 취약점만 노출하며 전체 취약점 목록은 사이드 패널을 참조하세요.

가져온 npm 패키지에서 보안 취약점을 찾고, 필요할 때 가져온 npm 패키지의 알려진 취약점을 즉시 확인할 수 있습니다:

<figure><img src="../../../../.gitbook/assets/image (345).png" alt="npm 패키지의 취약점"><figcaption><p>npm 패키지의 취약점</p></figcaption></figure>

`package.json` 파일에 있는 코드 내 취약점 수도 표시됩니다:

<figure><img src="../../../../.gitbook/assets/image (340).png" alt="취약점 수가 표시된 결과 화면"><figcaption><p>취약점 수가 표시된 결과 화면</p></figcaption></figure>

자주 사용하는 CDN(Content Delivery Networks)에서 JavaScript 패키지의 보안 취약점을 찾을 수 있습니다. 확장 프로그램은 프로젝트의 모든 HTML 파일을 스캔하고 즐겨 사용하는 CDN에서 가져온 모듈에 대한 취약점 정보를 표시합니다.

현재 지원되는 CDNs는 다음과 같습니다:

- unpkg.com
- ajax.googleapis.com
- cdn.jsdelivr.net
- cdnjs.cloudflare.com
- code.jquery.com
- maxcdn.bootstrapcdn.com
- yastatic.net
- ajax.aspnetcdn.com

<figure><img src="../../../../.gitbook/assets/oss-editor-html (1) (1).png" alt="CDN에서의 취약점"><figcaption><p>CDN에서의 취약점</p></figcaption></figure>

제공된 코드 작업을 트리거하여 가장 심각한 취약점으로 이동할 수 있습니다. 이 작업을 실행하면 자세한 정보를 보여주는 취약점 창이 열립니다:

<figure><img src="../../../../.gitbook/assets/Screenshot 2023-03-17 at 14.04.13.png" alt="코드 작업"><figcaption><p>코드 작업</p></figcaption></figure>

##  취약점 창

<figure><img src="../../../../.gitbook/assets/image (346).png" alt=" 취약점 창"><figcaption><p> 취약점 창</p></figcaption></figure>

오픈 소스 보안(OSS) 취약점 창은 취약 모듈에 대한 정보를 보여줍니다.

- 취약점을 자세히 설명하는 외부 리소스 링크(CVE, CWE, Snyk 취약점 DB)
- CVSS 점수 및 악용 성숙도
- 취약점이 시스템에 도입된 경로의 상세 정보
- 취약점 요약 및 수정 조언 요약