# GitLab - 플로우 및 기술

## 플로우 <a href="#flow" id="flow"></a>

1. GitLab 또는 GitLab Server 모드 설정( `url` 플래그를 통해 호스트가 제공되었는지 여부에 따라).
2. 인증 자격 증명이 액세스 할 수 있는 SCM에서 `하나`/`여러 개`의 그룹을 가져와 그룹 목록을 생성합니다.
3. 가져온/제공된 그룹 아래 `하나`/`모든` 프로젝트를 가져옵니다.
4. 모니터링되지 않는 리포지토리를 Snyk 계정에 쉽게 가져오기 위한 가져오기 파일 생성(`importConfDir` 플래그가 설정된 경우).
5. 가져온/제공된 프로젝트의 커밋을 가져와 기여자 목록을 생성합니다.
6. 기여자들에 의한 프로젝트의 커밋 수를 계산합니다.
7. 제외 파일에 지정된 기여자를 제외합니다(`exclusionFilePath` 플래그가 설정되고 올바른 경로를 텍스트 파일로 제공한 경우).
8. 결과를 출력합니다.

## 사용된 GitLab API 엔드포인트 <a href="#bitbucket-cloud-api-endpoints-used" id="bitbucket-cloud-api-endpoints-used"></a>

* GitLab에서 그룹 경로를 가져오기 위해 그룹/들 이름이 제공된 경우: `api/v4/groups?all_available=true&search={groupName}`
* 호스트 URL이 **제공되지 않은** 경우 GitLab에서 프로젝트를 가져오기 위해: `/api/v4/projects?membership=true`
* 호스트 URL이 **제공된** 경우 GitLab Server에서 프로젝트를 가져오기 위해: `/api/v4/projects`
* 가져온/제공된 프로젝트 목록의 커밋을 가져오기 위해: `/api/v4/projects/{ProjectPath}/repository/commits?since=${threeMonthsDate}`