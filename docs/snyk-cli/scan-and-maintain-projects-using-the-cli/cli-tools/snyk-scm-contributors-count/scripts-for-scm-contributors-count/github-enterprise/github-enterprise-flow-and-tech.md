# GitHub Enterprise - Flow and Tech

## Flow <a href="#flow" id="flow"></a>

1. SCM에서 액세스 권한을 가진 조직을 가져와서 orgs 목록을 만듭니다. 조직을 모두 가져올지 여부는 `fetchAllOrgs` 플래그에 따라 결정됩니다. `one`/`some`/`all` 중 하나를 가져옵니다.
2. 가져온/제공된 orgs 아래 `one` 또는 `all` repos를 가져옵니다.
3. 가져온/제공된 repo/s에 대한 커밋을 가져와서 기여자 목록을 만듭니다.
4. 기여자별로 repo/s의 커밋 수를 계산합니다.
5. 제외 파일에 지정된 기여자를 제거합니다 (`the exclusionFilePath` 플래그가 설정되어 유효한 경로가 제공된 경우).
6. 결과를 출력합니다.

## GitHub Enterprise API endpoints used <a href="#azure-api-endpoints-used" id="azure-api-endpoints-used"></a>

* GitHub Enterprise에서 orgs 가져오기: `api/v3/organizations` (`fetchAllOrgs` 플래그가 **설정**된 경우) 또는 `api/v3/user/orgs` (`fetchAllOrgs` 플래그가 **설정되지 않은** 경우)
* 가져온/제공된 orgs 목록과 관련 있는 repo/s 목록 가져오기: `api/v3/orgs/{Org}/repos`
* 가져온/제공된 repo/s 목록에 대한 커밋 가져오기: `api/v3/repos/{Org}/{Repo}/commits`