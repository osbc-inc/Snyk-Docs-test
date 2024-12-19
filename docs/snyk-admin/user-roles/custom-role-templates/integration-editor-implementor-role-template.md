# 통합 편집자/구현자 역할 템플릿

이 역할은 다수의 제 3자 도구를 통합하고 처리하는 관련 권한을 갖는 그룹 수준의 역할입니다.

Integration Editor/Implementor를 생성할 때, 이 사용자가 API를 구현하는지 여부를 정의합니다. 이 템플릿은 Organization-level **Service Accounts** 권한을 사용하여 API 구현을 포함합니다. 사용자 지정 역할에 필요하지 않은 경우 **Service Accounts** 권한을 포함하지 마십시오.

## 그룹 수준의 권한

이 역할을 만들려면 해당 카테고리에서 다음 권한을 활성화하십시오:

### Snyk 앱 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화 여부</th></tr></thead><tbody><tr><td>앱 보기</td><td>true</td></tr><tr><td>앱 설치</td><td>true</td></tr><tr><td>앱 편집</td><td>true</td></tr></tbody></table>

아래 나열된 권한의 나머지 카테고리는 해당 권한 모두를 비활성화로 설정해야 합니다:

* 그룹 관리
* 조직 관리
* 응용 위험 관리
* 감사 로그 관리
* IaC 설정 관리
* 통찰력 관리
* 문제 관리
* 보고서 관리
* 액세스 요청 관리
* 역할 관리
* 보안 및 라이선스 정책
* 서비스 계정 관리
* Snyk 미리 보기 관리
* SSO 설정 관리
* 태그 관리
* 사용자 관리

## 조직 수준의 권한

이 역할을 만들려면 해당 카테고리에서 다음 권한을 활성화하십시오:

### 조직 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화 여부</th></tr></thead><tbody><tr><td>조직 보기</td><td>true</td></tr><tr><td>조직 편집</td><td>false</td></tr><tr><td>조직 제거</td><td>false</td></tr></tbody></table>

### 통합 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화 여부</th></tr></thead><tbody><tr><td>통합 보기</td><td>true</td></tr><tr><td>통합 편집</td><td>true</td></tr></tbody></table>

### Kubernetes 통합 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화 여부</th></tr></thead><tbody><tr><td>Kubernetes 리소스 게시</td><td>true</td></tr></tbody></table>

### 프로젝트 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화 여부</th></tr></thead><tbody><tr><td>프로젝트 보기</td><td>true</td></tr><tr><td>프로젝트 추가</td><td>true</td></tr><tr><td>프로젝트 편집</td><td>true</td></tr><tr><td>프로젝트 상태 편집</td><td>true</td></tr><tr><td>프로젝트 테스트</td><td>true</td></tr><tr><td>프로젝트 이동</td><td>true</td></tr><tr><td>프로젝트 제거</td><td>true</td></tr><tr><td>프로젝트 이력 보기</td><td>false</td></tr><tr><td>프로젝트 통합 편집</td><td>true</td></tr><tr><td>프로젝트 속성 편집</td><td>false</td></tr><tr><td>Jira 이슈 보기</td><td>false</td></tr><tr><td>Jira 이슈 생성</td><td>false</td></tr><tr><td>프로젝트 태그 편집</td><td>false</td></tr></tbody></table>

### 프로젝트 풀 리퀘스트 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화 여부</th></tr></thead><tbody><tr><td>풀 리퀘스트 생성</td><td>true</td></tr><tr><td>풀 리퀘스트 체크 항목 성공으로 표시</td><td>true</td></tr></tbody></table>

### 보고서 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화 여부</th></tr></thead><tbody><tr><td>조직 보고서 보기</td><td>true</td></tr></tbody></table>

### 서비스 계정 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화 여부</th></tr></thead><tbody><tr><td>서비스 계정 보기</td><td>true</td></tr><tr><td>서비스 계정 생성</td><td>true</td></tr><tr><td>서비스 계정 편집</td><td>true</td></tr><tr><td>서비스 계정 제거</td><td>true</td></tr></tbody></table>

### Snyk 앱 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화 여부</th></tr></thead><tbody><tr><td>앱 보기</td><td>true</td></tr><tr><td>앱 설치</td><td>true</td></tr><tr><td>앱 생성</td><td>true</td></tr><tr><td>앱 편집</td><td>true</td></tr><tr><td>앱 삭제</td><td>true</td></tr></tbody></table>

### 웹훅 관리

<table><thead><tr><th>권한</th><th data-type="checkbox">활성화 여부</th></tr></thead><tbody><tr><td>발신 웹훅 보기</td><td>true</td></tr><tr><td>발신 웹훅 생성</td><td>true</td></tr><tr><td>발신 웹훅 제거</td><td>true</td></tr></tbody></table>

아래 나열된 권한의 나머지 카테고리는 해당 권한 모두를 비활성화로 설정해야 합니다:

* 감사 로그 관리
* 청구 관리
* 수집 관리
* 컨테이너 이미지 관리
* 자격 증명 관리
* 통합 관리
* 패키지 관리
* 프로젝트 무시 관리
* Snyk 클라우드 관리
* Snyk 미리 보기 관리
* 사용자 관리