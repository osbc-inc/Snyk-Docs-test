# Incident Responder role template

이 조직 수준 역할은 특정 유형의 문제를 찾고 변경을 가하는 데 Snyk의 핵심 기능에 신속하게 액세스해야 합니다.

## 그룹 수준 권한

이 템플릿은 조직 수준 역할이며 그룹 수준 권한은 없습니다.

## 조직 수준 권한

이 역할을 만들려면 관련 카테고리에서 다음 권한을 활성화하세요:

### 조직 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화?</th></tr></thead><tbody><tr><td>조직 보기</td><td>true</td></tr><tr><td>조직 편집</td><td>false</td></tr><tr><td>조직 제거</td><td>false</td></tr></tbody></table>

### 프로젝트 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화?</th></tr></thead><tbody><tr><td>프로젝트 보기</td><td>true</td></tr><tr><td>프로젝트 추가</td><td>true</td></tr><tr><td>프로젝트 편집</td><td>true</td></tr><tr><td>프로젝트 상태 편집</td><td>true</td></tr><tr><td>프로젝트 테스트</td><td>true</td></tr><tr><td>프로젝트 이동</td><td>true</td></tr><tr><td>프로젝트 제거</td><td>true</td></tr><tr><td>프로젝트 기록 보기</td><td>true</td></tr><tr><td>프로젝트 통합 편집</td><td>false</td></tr><tr><td>프로젝트 속성 편집</td><td>true</td></tr><tr><td>Jira 문제 보기</td><td>true</td></tr><tr><td>Jira 문제 생성</td><td>true</td></tr><tr><td>프로젝트 태그 편집</td><td>true</td></tr></tbody></table>

### 프로젝트 무시 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화?</th></tr></thead><tbody><tr><td>프로젝트 무시 보기</td><td>true</td></tr><tr><td>프로젝트 무시 생성</td><td>true</td></tr><tr><td>프로젝트 무시 편집</td><td>true</td></tr><tr><td>프로젝트 무시 제거</td><td>true</td></tr></tbody></table>

### 프로젝트 풀 리퀘스트 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화?</th></tr></thead><tbody><tr><td>풀 리퀘스트 생성</td><td>true</td></tr><tr><td>풀 리퀘스트 검사 성공으로 표시</td><td>true</td></tr></tbody></table>

### 보고서 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화?</th></tr></thead><tbody><tr><td>조직 보고서 보기</td><td>true</td></tr></tbody></table>

아래 나열된 권한 카테고리의 나머지 권한은 모두 비활성화로 설정되어야 합니다:

* 감사 로그 관리
* 요금 관리
* 수집 관리
* 컨테이너 이미지 관리
* 권한 관리
* 통합 관리
* Kubernetes 통합 관리
* 패키지 관리
* 서비스 계정 관리
* Snyk 앱 관리
* Snyk 클라우드 관리
* Snyk 미리보기 관리
* 사용자 관리
* 웹훅 관리