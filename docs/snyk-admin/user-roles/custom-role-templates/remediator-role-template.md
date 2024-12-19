# Remediator 역할 템플릿

이 것은 조직 수준의 역할이며, Remediator는 Snyk에서 특정 영역과 기능만 볼 수 있으며 PR, 프로젝트 등을 만들 수 없습니다.

Remediator는 문제 평가 및 우선 순위 설정에서 개발자를 돕는 역할을 합니다.

## 그룹 수준 권한

이 템플릿은 조직 수준 역할이며 그룹 수준 권한이 없습니다.

## 조직 수준 권한

이 역할을 생성하려면 관련 카테고리에서 다음 권한을 활성화하십시오:

### 조직 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화됨?</th></tr></thead><tbody><tr><td>조직 보기</td><td>true</td></tr><tr><td>조직 편집</td><td>false</td></tr><tr><td>조직 제거</td><td>false</td></tr></tbody></table>

### 프로젝트 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화됨?</th></tr></thead><tbody><tr><td>프로젝트 보기</td><td>true</td></tr><tr><td>프로젝트 추가</td><td>false</td></tr><tr><td>프로젝트 편집</td><td>true</td></tr><tr><td>프로젝트 상태 편집</td><td>false</td></tr><tr><td>프로젝트 테스트</td><td>true</td></tr><tr><td>프로젝트 이동</td><td>true</td></tr><tr><td>프로젝트 제거</td><td>false</td></tr><tr><td>프로젝트 히스토리 보기</td><td>true</td></tr><tr><td>프로젝트 통합 편집</td><td>false</td></tr><tr><td>프로젝트 속성 편집</td><td>false</td></tr><tr><td>Jira 이슈 보기</td><td>true</td></tr><tr><td>Jira 이슈 생성</td><td>true</td></tr><tr><td>프로젝트 태그 편집</td><td>true</td></tr></tbody></table>

### 프로젝트 무시 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화됨?</th></tr></thead><tbody><tr><td>프로젝트 무시 보기</td><td>true</td></tr><tr><td>프로젝트 무시 생성</td><td>false</td></tr><tr><td>프로젝트 무시 편집</td><td>false</td></tr><tr><td>프로젝트 무시 제거</td><td>false</td></tr></tbody></table>

아래 남은 권한 카테고리는 모두 비활성화된 상태여야 합니다:

* 감사 로그 관리
* 과금 관리
* 수집 관리
* 컨테이너 이미지 관리
* 자격 관리
* 통합 관리
* Kubernetes 통합 관리
* 패키지 관리
* 프로젝트 풀 리퀘스트 관리
* 보고서 관리
* 서비스 계정 관리
* Snyk 앱 관리
* Snyk 클라우드 관리
* Snyk 미리보기 관리
* 사용자 관리
* 웹훅 관리