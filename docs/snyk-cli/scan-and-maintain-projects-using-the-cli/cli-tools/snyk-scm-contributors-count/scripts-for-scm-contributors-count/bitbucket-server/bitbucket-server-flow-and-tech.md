# Bitbucket Server - Flow 및 기술

## Flow <a href="#flow" id="flow"></a>

1. Snyk로부터 모니터링 중인 프로젝트를 가져옵니다 (만약 `skipSnykMonitoredRepos` 플래그가 **설정되지 않았고** `SNYK_TOKEN`이 export되었다면).
2. 자격 증명이 액세스할 수 있는 프로젝트들을 SCM에서 `one`/`some`/`all` 가져와 프로젝트 목록을 생성합니다.
3. 가져오거나 제공된 프로젝트 아래의 `one`/`all` 저장소를 가져옵니다.
4. Snyk에서 모니터링되지 않는 저장소를 제거합니다 (만약 `skipSnykMonitoredRepos` 플래그가 **설정되지 않았고** `SNYK_TOKEN`이 export되었다면) 그리고 저장소 목록을 생성합니다.
5. 모니터링되지 않는 저장소를 쉽게 Snyk 계정으로 가져오기 위한 가져오기 파일을 생성합니다 (만약 `importConfDir` 플래그가 설정되어 있다면).
6. 가져오거나 제공된 저장소에 대해 커밋을 가져와 기여자 목록을 생성합니다.
7. 기여자들이 한 저장소에 대한 커밋 수를 계산합니다.
8. 제외 파일에 명시된 기여자들을 제거합니다 (만약 `exclusionFilePath` 플래그가 설정되어 있고 유효한 텍스트 파일 경로가 제공되었다면).
9. 결과를 출력합니다.

## Bitbucket Server API endpoints used <a href="#bitbucket-cloud-api-endpoints-used" id="bitbucket-cloud-api-endpoints-used"></a>

* BB Cloud에서 워크스페이스가 **제공되지 않았을 때** 리포지토리를 가져오려면: `/rest/api/1.0/repos`
* BB Cloud에서 워크스페이스/들이 **제공되었을 때** 리포지토리를 가져오려면: `/rest/api/1.0/projects/{Project}/repos`
* 가져오거나 제공된 저장소 목록에 대해 커밋을 가져오기 위한 API 엔드포인트: `/rest/api/1.0/projects/{Project}/repos/{Repo}/commits`
