# 코드 분석 분해

리포지토리를 가져오면 가 가져온 코드 내의 취약점을 자동으로 테스트합니다. 단일 리포지토리의 모든 파일에서 감지된 취약점은 Snyk 프로젝트로 편성되어 Code Analysis로 라벨이 붙습니다. Code Analysis는 특정 리포지토리에 대한 테스트 결과를 제공하며, 해당 리포지토리 소스 코드에서 발견된 모든 취약점이 나열됩니다.

<figure><img src="../../../.gitbook/assets/Code analysis .png" alt="코드 분석 페이지 개요."><figcaption><p>코드 분석 페이지</p></figcaption></figure>

## 코드 분석 구성 요소

이 표는 코드 분석 프로젝트의 요소를 요약한 것입니다.

<table><thead><tr><th width="160">구성 요소</th><th>설명</th></tr></thead><tbody><tr><td>헤더</td><td>가져온 리포지토리의 세부 정보를 포함하며 Git 리포지토리 내 리포지토리로의 링크, 프로젝트 이름 및 프로젝트 탭: <strong>개요</strong>, <strong>이력</strong>, 및 <strong>설정</strong>을 보여줍니다.</td></tr><tr><td>프로젝트 요약 정보 영역</td><td>리포지토리 가져오기 및 마지막 테스트 날짜, 온디맨드 테스트 옵션을 위한 다시 테스트하기, 리포지토리를 가져온 사용자 이름, 프로젝트 소유자의 이름 및 분석된 코드 파일 수 및 분석되지 않은 파일 수를 표시합니다. <a href="./#retesting-code-repository">코드 리포지토리 재테스트하기</a> 보기</td></tr><tr><td><a href="../../../snyk-admin/snyk-projects/project-information.md">프로젝트 필터</a></td><td>표시된 문제를 필터링하기 위한 사전 정의된 기준 세트를 포함합니다.</td></tr><tr><td>취약점 문제</td><td>가져온 리포지토리에서 가 발견한 취약점을 포함합니다.</td></tr><tr><td><a href="breakdown-of-code-analysis.md#data-flow">데이터 흐름</a></td><td>코드에서 문제의 오염된 흐름을 표시합니다.</td></tr><tr><td><a href="breakdown-of-code-analysis.md#fix-analysis">고치기 분석</a></td><td>발견된 취약점 유형에 대한 추가 정보를 제공하며, 이 문제를 방지하는 데 가장 좋은 방법과 수정의 코드 예제를 제공합니다.</td></tr><tr><td>CWE</td><td>특정 취약점 유형의 CWE (Common Weakness Enumeration) ID와 이 취약점 유형이 설명된 CWE 웹사이트로의 링크를 포함합니다. <a href="breakdown-of-code-analysis.md#example-cwe-22-path-traversal">예: CWE-22: 경로 이탈</a> 보기</td></tr><tr><td><a href="./#open-repository-external-link">외부 링크로 리포지토리 열기</a></td><td>즉시 해결을 위한 통합된 Git 리포지토리로의 빠른 액세스 기능을 제공합니다.</td></tr><tr><td>가져온 사용자</td><td>분석된 리포지토리를 가져온 Git 리포지토리 사용자 이름입니다.</td></tr><tr><td><a href="../../../snyk-admin/snyk-projects/project-information.md">프로젝트 소유자</a></td><td>조직 내 프로젝트의 주인을 식별하는데 사용되며, 이는 Snyk 프로젝트에는 영향을 주지 않습니다. 필수 사항은 아니며 기본적으로 비어 있습니다. </td></tr><tr><td><a href="../../../snyk-admin/snyk-projects/project-attributes.md">환경</a></td><td>환경 속성은 클라이언트 측(Frontend)부터 서버 측(Backend)까지의 소프트웨어 컨텍스트를 설명하며, 모바일, 클라우드(SaaS), 온프레미스, 호스팅 서비스 및 분산 시스템과 같은 플랫폼을 포함합니다.</td></tr><tr><td><a href="../../../snyk-admin/snyk-projects/project-attributes.md">비즈니스 중요도</a></td><td>위험 점수 기능이 활성화된 경우 비즈니스 중요도 속성은 자동으로 점수에 영향을 미치며, 최상위 속성 수준으로 점수를 조정합니다. <a href="../../../manage-risk/prioritize-issues-for-fixing/risk-score.md">위험 점수</a> 보기</td></tr><tr><td><a href="../../../snyk-admin/snyk-projects/project-attributes.md">수명 주기</a></td><td>수명 주기 속성은 Pro, 개발 또는 샌드박스를 구분합니다.</td></tr><tr><td><a href="../../../snyk-admin/introduction-to-snyk-projects/project-tags.md">태그</a></td><td>Snyk 프로젝트 태그 기능을 사용하여 사용자 정의 메타데이터 태그를 추가, 제거 및 사용하여 프로젝트를 구성하고 필터링할 수 있습니다. 이는 프로젝트 관리 및 탐색을 단순하게 해줍니다.</td></tr><tr><td>분석 요약</td><td>리포지토리에서 분석된 코드 파일 수 및 분석된 파일의 총 리포지토리 코드 파일 중 분석된 파일의 백분율을 표시합니다.</td></tr><tr><td>리포 분해</td><td><ul><li><strong>분해된 파일</strong>: 에서 검토된 파일을 포함하여 인식된 확장자 및 프로그래밍 언어를 포함합니다.</li><li><strong>분해 안된 파일:</strong> 지원되지 않는 언어 또는 확장자로 인해 아직 분석되지 않은 텍스트 파일을 포함합니다.</li><li><strong>알 수 없음</strong>: 인식된 확장자가 없는 파일로, 멀티미디어 콘텐츠(사진 및 비디오), 이진, 자사 포맷 파일 또는 의 관심 영역 밖의 다른 형식을 포함할 수 있습니다.</li></ul></td></tr></tbody></table>

## 데이터 흐름

데이터 흐름은 코드에서 발견된 문제의 위치 및 응용 프로그램 전반에 걸쳐 흘러가는 방식을 보여줍니다. 코드에서 문제의 오염된 흐름을 보여주며, 소스에서 싱크로, 코드 흐름의 모든 단계의 코드 라인을 제시합니다.

소스는 잠재적인 문제의 입력 지점입니다. 이것은 사용자 또는 외부 장치가 데이터를 입력할 수 있는 애플리케이션의 지점으로, 이는 애플리케이션의 보안을 위반할 가능성이 있는 데이터를 입력해야 합니다. 예를 들어 SQL Injection 문제의 경우, 소스는 사용자가 입력하는 양식 또는 다른 데이터 입력 영역일 것입니다.

싱크는 코드에서 문제가 실행되는 연산입니다. 이 지점은 깨끗한 입력을 받아야 하며, 그렇지 않으면 악용될 수 있습니다. 예를 들어 SQL Injection 문제의 경우, 싱크는 DB에게 받은 입력에 따라 특정 조치를 수행하는 내부 작업일 것입니다.

{Snyk Code가 발견한 모든 문제에는 데이터 흐름이 있습니다. 예를 들어 하드코딩된 비밀 정보의 경우 문제의 소스가 데이터 흐름 페이지에 표시됩니다.

### 데이터 흐름 보기

1. Snyk 웹 UI에 로그인하고 [그룹 및 조직](../../../snyk-admin/groups-and-organizations/)을 선택합니다.
2. **프로젝트**로 이동하고 리포지토리 프로젝트를 포함하는 대상 폴더를 선택합니다.
3. **코드 분석** 프로젝트를 엽니다.
4. (옵션) [특정 취약점 문제를 검색하거나 필터링](./#vulnerability-issues)합니다.
5. 취약점 문제를 선택하고 **전체 세부 정보** > **데이터 흐름**으로 이동합니다.
6. 데이터 흐름 분석의 일환으로 다음 작업을 수행할 수 있습니다:

- 코드에서 소스에서 싱크로의 문제의 오염된 흐름을 보기. [데이터 흐름 분석 예](breakdown-of-code-analysis.md#data-flow-analysis-example) 보기
- [통합된 Git 리포지토리에서 데이터 흐름 외부 링크 열기](breakdown-of-code-analysis.md#open-data-flow-external-link)
- **무시** 버튼을 사용하여 열린 취약점 문제를 무시합니다. [문제 무시](../../../manage-risk/prioritize-issues-for-fixing/ignore-issues/) 보기

### **데이터 흐름 분석 예시**

다음 경로 이탈 문제에서 개발자는 입력을 살균하지 않았습니다. 이는 공격자가 파일 시스템의 모든 파일, 비밀번호 파일과 같은 민감한 데이터에 액세스하도록 경로 이탈 공격을 수행할 수 있게 합니다.

<figure><img src="../../../.gitbook/assets/Path Traversal vulnerability issues.png" alt="경로 이탈 취약점 문제 개요."><figcaption><p>경로 이탈 취약점 문제의 데이터 흐름</p></figcaption></figure>

### **데이터 흐름 외부 링크 열기**

Git 리포지토리에서 표시된 소스 코드를 열려면 오른쪽 패널 위의 파일 이름을 선택합니다. 이 예에선 파일 이름이 **routes/profileImageUrlUpload.ts**입니다.

<figure><img src="../../../.gitbook/assets/Open source code in data flow (1).png" alt="데이터 흐름 내의 소스 코드 외부 링크입니다."><figcaption><p>데이터 흐름 내의 소스 코드 외부 링크</p></figcaption></figure>

소스 코드는 통합된 Git 리포지토리에 나타나며, 취약점을 수정해야 하는 정확한 위치를 보여줍니다. 코드를 수정하여 취약점을 해결할 수 있습니다.

<figure><img src="../../../.gitbook/assets/Fix vulnerability in Git repository from data flow.png" alt="외부 소스 코드에 표시된 취약점입니다."><figcaption><p>외부 소스 코드에 표시된 취약점</p></figcaption></figure>

## 고치기 분석

고치기 분석을 통해 코드에서 발견된 취약점 문제를 해결할 수 있습니다. 발견된 취약점 유형에 대한 세부 정보, 이 문제를 방지하는 데 사용할 수 있는 모범 사례 및 전역 오픈 소스 커뮤니티의 코드 예제가 제공됩니다.

특정 취약점을 확인하는 더 자세한 정보를 찾으려면 CWE 링크를 열어 해당 취약점 유형에 대해 자세히 알아볼 수 있습니다. [CWE-22 ](breakdown-of-code-analysis.md#example-cwe-22-path-traversal) 및 [CWE-601 예시](breakdown-of-code-analysis.md#example-cwe-601-open-redirect)를 참조하세요.

일부 취약점은 이해, 해결 및 방지하는 대화형 레슨에 대한 링크도 포함합니다. [Snyk Learn](https://learn.snyk.io/)를 참조하세요.

<figure><img src="../../../.gitbook/assets/ - Results - Issues - Fix analysis page - 2.png" alt="경로 이탈 취약점의 고치기 분석 페이지"><figcaption><p>경로 이탈 취약점의 고치기 분석 페이지</p></figcaption></figure>

### 고치기 분석 보기

1. Snyk 웹 UI에 로그인하고 [그룹 및 조직](../../../snyk-admin/groups-and-organizations/)을 선택합니다.
2. **프로젝트**로 이동하고 리포지토리 프로젝트를 포함하는 대상 폴더를 선택합니다.
3. **코드 분석** 프로젝트를 엽니다.
4. (옵션) [특정 취약점 문제를 검색하거나 필터링](./#vulnerability-issues)합니다.
5. 취약점 문제를 선택하고 **전체 세부 정보** > **고치기 분석**으로 이동합니다.
6. 고치기 분석의 일환으로 다음 작업을 수행할 수 있습니다:

- 발견된 문제를 보고 방지 방법을 살펴봅니다.
- 전역 오픈 소스 커뮤니티의 코드 예제를 검토하고 검토하고 코드 샘플을 살펴봅니다.
- 통합된 Git 리포지토리에 나타나는 고치기 예제의 코드 차이를 보려면 [데이터 흐름 분석 외부 링크 열기](breakdown-of-code-analysis.md#data-flow-analysis-example)를 참조하세요.
- **무시** 버튼을 사용하여 열린 취약점 문제를 무시합니다. [문제 무시](../../../manage-risk/prioritize-issues-for-fixing/ignore-issues/)를 참조하세요.

**고치기 분석** 페이지를 통해 다음을 수행할 수 있습니다:

### **고치기 분석 외부 링크 열기**

Git 리포지토리에서 취약점에 대한 코드Snyk 취약점 데이터베이스에서 유사한 취약성을 식별한 점수

### 질적인 요소&#x20;

* 소스의 심각성, 직접적인지 간접적인지
* 싱크의 보편성과 영향
* 보안팀 경험 및 연구
* 고객 피드백

## 예: CWE-22: 경로 이탈

CWE-22 경로 이탈의 경우, 취약성이 테스트에서 발생하면 낮은 심각성입니다. 테스트에서 발생하지 않고 직접 소스에서 유래한 경우, 심각성은 높습니다. 그렇지 않으면 낮은 심각성입니다.

<figure><img src="../../../.gitbook/assets/image (2) (2).png" alt="우선순위 점수 CWE-22 경로 이탈을위한 의사결정 플로 차트"><figcaption><p>우선순위 점수 CWE-22 경로 이탈을위한 의사결정 플로 차트</p></figcaption></figure>

## 예: CWE-601: 오픈 리디렉션

CWE-2601 오픈 리디렉션의 경우, 취약성이 테스트에서 발생하면 낮은 심각성입니다. 테스트에서 발생하지 않고 직접 소스에서 유래한 경우, 중간 심각성입니다.

<figure><img src="../../../.gitbook/assets/image (5) (8).png" alt="우선순위 점수 CWE-601 오픈 리디렉션을위한 의사결정 플로 차트"><figcaption><p>우선순위 점수 CWE-601 오픈 리디렉션을위한 의사결정 플로 차트</p></figcaption></figure>