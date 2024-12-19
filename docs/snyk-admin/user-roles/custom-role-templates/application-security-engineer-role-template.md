# 어플리케이션 보안 엔지니어 역할 템플릿

이 기관 수준 역할은 프로젝트 및 무시 사항을 추가, 이동, 제거할 수 있으며 PR 체크를 성공으로 표시할 수 있습니다.

## 그룹 수준 권한

이 템플릿은 기관 수준 역할이며 그룹 수준 권한은 없습니다.

## 기관 수준 권한

이 역할을 만들려면 해당 카테고리에서 다음 권한을 활성화하십시오:

### 기관 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화됨?</th></tr></thead><tbody><tr><td>기관 보기</td><td>true</td></tr><tr><td>기관 편집</td><td>false</td></tr><tr><td>기관 제거</td><td>false</td></tr></tbody></table>

### 프로젝트 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화됨?</th></tr></thead><tbody><tr><td>프로젝트 보기</td><td>true</td></tr><tr><td>프로젝트 추가</td><td>true</td></tr><tr><td>프로젝트 편집</td><td>true</td></tr><tr><td>프로젝트 상태 편집</td><td>true</td></tr><tr><td>프로젝트 테스트</td><td>true</td></tr><tr><td>프로젝트 이동</td><td>true</td></tr><tr><td>프로젝트 제거</td><td>false</td></tr><tr><td>프로젝트 히스토리 보기</td><td>true</td></tr><tr><td>프로젝트 통합 수정</td><td>true</td></tr><tr><td>프로젝트 속성 편집</td><td>true</td></tr><tr><td>Jira 이슈 보기</td><td>true</td></tr><tr><td>Jira 이슈 생성</td><td>true</td></tr><tr><td>프로젝트 태그 편집</td><td>true</td></tr></tbody></table>

### 프로젝트 무시 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화됨?</th></tr></thead><tbody><tr><td>프로젝트 무시 보기</td><td>true</td></tr><tr><td>프로젝트 무시 생성</td><td>true</td></tr><tr><td>프로젝트 무시 편집</td><td>true</td></tr><tr><td>프로젝트 무시 제거</td><td>true</td></tr></tbody></table>

### 프로젝트 풀 리퀘스트 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화됨?</th></tr></thead><tbody><tr><td>풀 리퀘스트 생성</td><td>false</td></tr><tr><td>풀 리퀘스트 체크를 성공으로 표시</td><td>true</td></tr></tbody></table>

아래 나열된 권한 카테고리의 나머지 권한은 해당 카테고리 내의 모든 권한이 비활성화되도록 설정해야 합니다:

* 감사 로그 관리
* 결제 관리
* 컬렉션 관리
* 컨테이너 이미지 관리
* 권한 관리
* 통합 관리
* Kubernetes 통합 관리
* 패키지 관리
* 보고서 관리
* 서비스 계정 관리
* Snyk 앱 관리
* Snyk 클라우드 관리
* Snyk 미리 보기 관리
* 사용자 관리
* 웹훅 관리