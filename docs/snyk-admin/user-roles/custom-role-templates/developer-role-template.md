# 개발자 역할 템플릿

이 조직 수준 역할은 스캔 결과를 검토하고 문제를 해결하며 프로젝트 테스트를 시작하는 데 도움을 줍니다. 이 역할을 가진 사용자는 조직 및 프로젝트를 볼 수 있습니다.

보통 Snyk를 배포할 때, 개발자는 Snyk PR 체크를 재정의할 수 있는 권한을 가질 수 있지만, 이 권한은 개발자들이 Snyk IDE 확장 기능을 사용하고 SDLC 초기에 문제를 해결하기 시작하는 데 편안해지면 취소할 수 있습니다. 마찬가지로, 개발자들이 프로젝트를 추가할 수 있도록 시작한 다음에 이 권한을 팀 리드로 제한할 수도 있습니다.

## 그룹 수준 권한

이 템플릿은 조직 수준 역할을 위한 것이며 그룹 수준 권한은 없습니다.

## 조직 수준 권한

이 역할을 만들려면 해당 카테고리에서 다음 권한을 활성화하세요:

### 조직 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화?</th></tr></thead><tbody><tr><td>조직 보기</td><td>true</td></tr><tr><td>조직 편집</td><td>false</td></tr><tr><td>조직 제거</td><td>false</td></tr></tbody></table>

### 프로젝트 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화?</th></tr></thead><tbody><tr><td>프로젝트 보기</td><td>true</td></tr><tr><td>프로젝트 추가</td><td>true</td></tr><tr><td>프로젝트 편집</td><td>true</td></tr><tr><td>프로젝트 상태 편집</td><td>true</td></tr><tr><td>프로젝트 테스트</td><td>true</td></tr><tr><td>프로젝트 이동</td><td>true</td></tr><tr><td>프로젝트 제거</td><td>true</td></tr><tr><td>프로젝트 이력 보기</td><td>true</td></tr><tr><td>프로젝트 통합 편집</td><td>false</td></tr><tr><td>프로젝트 속성 편집</td><td>true</td></tr><tr><td>Jira 이슈 보기</td><td>true</td></tr><tr><td>Jira 이슈 생성</td><td>true</td></tr><tr><td>프로젝트 태그 편집</td><td>true</td></tr></tbody></table>

### 프로젝트 무시 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화?</th></tr></thead><tbody><tr><td>프로젝트 무시 보기</td><td>true</td></tr><tr><td>프로젝트 무시 생성</td><td>true</td></tr><tr><td>프로젝트 무시 편집</td><td>true</td></tr><tr><td>프로젝트 무시 제거</td><td>true</td></tr></tbody></table>

### 프로젝트 pull request 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화?</th></tr></thead><tbody><tr><td>pull request 생성</td><td>true</td></tr><tr><td>pull request 체크 성공으로 표시</td><td>true</td></tr></tbody></table>

### Snyk 클라우드 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화?</th></tr></thead><tbody><tr><td>환경 보기</td><td>false</td></tr><tr><td>환경 생성</td><td>false</td></tr><tr><td>환경 삭제</td><td>false</td></tr><tr><td>환경 업데이트</td><td>false</td></tr><tr><td>스캔 보기</td><td>true</td></tr><tr><td>스캔 생성</td><td>true</td></tr><tr><td>리소스 보기</td><td>false</td></tr><tr><td>아티팩트 보기</td><td>false</td></tr><tr><td>아티팩트 생성</td><td>false</td></tr><tr><td>사용자 지정 룰 보기</td><td>false</td></tr><tr><td>사용자 지정 룰 생성</td><td>false</td></tr><tr><td>사용자 지정 룰 편집</td><td>false</td></tr><tr><td>사용자 지정 룰 삭제</td><td>false</td></tr></tbody></table>

아래에 나열된 권한 카테고리의 남은 권한은 모두 비활성화로 설정되어야 합니다:

* 감사 로그 관리
* 요금 관리
* 수집 관리
* 컨테이너 이미지 관리
* 자격 부여 관리
* 통합 관리
* Kubernetes 통합 관리
* 패키지 관리
* 보고서 관리
* 서비스 계정 관리
* Snyk 앱 관리
* Snyk 미리 보기 관리
* 사용자 관리
* 웹훅 관리