# CLI Tester role template

이 조직 수준의 역할은 `snyk test`와 `snyk monitor`를 사용할 수 있습니다.

## 그룹 수준의 권한

이 템플릿은 조직 수준의 역할로 그룹 수준의 권한이 없습니다.

## 조직 수준의 권한

이 역할을 생성하려면 관련 카테고리에서 다음 권한을 활성화하십시오:

### 조직 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화?</th></tr></thead><tbody><tr><td>조직 보기</td><td>true</td></tr><tr><td>조직 편집</td><td>false</td></tr><tr><td>조직 제거</td><td>false</td></tr></tbody></table>

### 패키지 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화?</th></tr></thead><tbody><tr><td>패키지 테스트</td><td>true</td></tr></tbody></table>

### 프로젝트 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화?</th></tr></thead><tbody><tr><td>프로젝트 보기</td><td>true</td></tr><tr><td>프로젝트 추가</td><td>true</td></tr><tr><td>프로젝트 편집</td><td>false</td></tr><tr><td>프로젝트 상태 편집</td><td>false</td></tr><tr><td>프로젝트 테스트</td><td>true</td></tr><tr><td>프로젝트 이동</td><td>false</td></tr><tr><td>프로젝트 제거</td><td>false</td></tr><tr><td>프로젝트 이력 보기</td><td>false</td></tr><tr><td>프로젝트 통합 수정</td><td>false</td></tr><tr><td>프로젝트 속성 편집</td><td>false</td></tr><tr><td>Jira 이슈 보기</td><td>false</td></tr><tr><td>Jira 이슈 생성</td><td>false</td></tr><tr><td>프로젝트 태그 편집</td><td>false</td></tr></tbody></table>

### Snyk 미리보기 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화?</th></tr></thead><tbody><tr><td>Snyk 미리보기 기능 보기</td><td>true</td></tr><tr><td>Snyk 미리보기 기능 편집</td><td>false</td></tr></tbody></table>

아래 나열된 권한 카테고리의 나머지 권한은 모두 비활성화해야 합니다:

* 감사 로그 관리
* 결제 관리
* 수집 관리
* 컨테이너 이미지 관리
* 자격 부여 관리
* 통합 관리
* Kubernetes 통합 관리
* 프로젝트 무시 관리
* 프로젝트 풀 요청 관리
* 보고서 관리
* 서비스 계정 관리
* Snyk 앱 관리
* Snyk 클라우드 관리
* 사용자 관리
* 웹훅 관리