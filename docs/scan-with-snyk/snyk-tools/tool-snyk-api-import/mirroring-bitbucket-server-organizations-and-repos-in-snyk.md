# Snyk에서 Bitbucket Server 조직 및 저장소를 미러링하는 방법

Snyk에 Bitbucket Server 저장소 전체를 가져오기 위해 네 가지 명령어를 사용할 수 있습니다. 진행하려면 Bitbucket Server 토큰과 Snyk 토큰을 환경 변수로 구성해야 합니다.

## Bitbucket Server 저장소 가져오는 일반적인 단계

자세한 정보는 개별 문서 페이지를 참조하십시오. 일반적인 단계는 다음과 같습니다:

1. `export BITBUCKET_SERVER_TOKEN=***` 및 `export SNYK_TOKEN=***`
2. 조직 데이터 생성 예시: `snyk-api-import orgs:data --source=bitbucket-server --groupId=<snyk_group_id>` 자세한 지침: [Snyk에서 조직 생성하기](creating-organizations-in-snyk.md)
3. Snyk에서 조직 생성: `snyk-api-import orgs:create --file=orgs.json` [Snyk에서 조직 생성하기](creating-organizations-in-snyk.md)의 전체 지침을 따르면 Snyk 조직 ID 및 가져오기에 필요한 통합 ID를 포함한 `snyk-created-orgs.json` 파일이 생성됩니다.
4. 가져오기 데이터 생성: `snyk-api-import import:data --orgsData=snyk-created-orgs.json --source=bitbucket-server --integrationType=bitbucket-server` 자세한 지침: [가져오기 대상 데이터 생성](creating-import-targets-data-for-import-command.md)
5. 가져오기 실행: `DEBUG=*snyk* snyk-api-import import` 전체 지침: [가져오기 시작하기](kicking-off-an-import.md)\`\`

## 미러링하는 동안 새로운 저장소 및 조직만 다시 가져오는 방법

초기 가져오기가 완료되면 주기적으로 새로운 저장소를 확인하고 다음 단계를 따라 Snyk에 추가해야 합니다.

1. `export BITBUCKET_SERVER_TOKEN=***` 및 `export SNYK_TOKEN=***`
2. Snyk에서 조직 데이터 생성하고 저장소가 없는 경우 건너뛰려면 `--skipEmptyOrg`를 사용하여 `snyk-api-import orgs:data --source=bitbucket-server --groupId=<snyk_group_id> --skipEmptyOrg` 자세한 지침: [Snyk에서 조직 생성하기](creating-organizations-in-snyk.md)
3. Snyk에서 조직 생성하고 기존에 생성된 조직을 건너뛰려면 `--noDuplicateNames` 매개변수를 사용하여 `snyk-api-import orgs:create --file=orgs.json --noDuplicateNames` [Snyk에서 조직 생성하기](creating-organizations-in-snyk.md)의 전체 지침을 따르면 Snyk 조직 ID 및 가져오기에 필요한 통합 ID를 포함한 `snyk-created-orgs.json` 파일이 생성됩니다.
4. 가져오기 데이터 생성: `snyk-api-import import:data --orgsData=snyk-created-orgs.json --source=bitbucket-server --integrationType=bitbucket-server` 자세한 지침: [가져오기 대상 데이터 생성](creating-import-targets-data-for-import-command.md)
5. 선택 사항 - 이전에 가져온 로그를 생성하여 그룹 내 모든 이전에 가져온 저장소 건너뛰기: `snyk-api-import-macos list:imported --integrationType=<integration-type> --groupId=<snyk_group_id>` 전체 지침: [가져오기 시작하기](kicking-off-an-import.md)\`\`
6. 가져오기 실행: `DEBUG=*snyk* snyk-api-import import` 전체 지침: [가져오기 시작하기](kicking-off-an-import.md)\`\`