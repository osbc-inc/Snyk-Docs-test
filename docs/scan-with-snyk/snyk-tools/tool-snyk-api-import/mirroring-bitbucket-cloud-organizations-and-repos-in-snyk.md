# Snyk에서 Bitbucket Cloud 조직과 저장소 반영하기

Snyk에서 Bitbucket Cloud 저장소 전체를 가져오기 위해 제공되는 유틸리티에서 네 가지 명령어를 사용할 수 있습니다. 계속 진행하려면 Bitbucket Cloud 사용자 이름과 암호, 그리고 Snyk 토큰을 환경 변수로 구성해야 합니다.

## Bitbucket Cloud 저장소 가져오는 일반적인 단계

자세한 정보는 개별 문서 페이지를 참조하십시오. 일반적인 단계는 다음과 같습니다:

1. `export BITBUCKET_CLOUD_USERNAME=***`, `export BITBUCKET_CLOUD_PASSWORD=***` 및 `export SNYK_TOKEN=***` 설정
2. 조직 데이터 생성, 예를 들어, `snyk-api-import orgs:data --source=bitbucket-cloud --groupId=<snyk_group_id>` 전체 지침: [Snyk에서 조직 만들기](creating-organizations-in-snyk.md)
3. Snyk에서 조직 생성: `snyk-api-import orgs:create --file=orgs.json` [Snyk에서 조직 만들기](creating-organizations-in-snyk.md)의 전체 지침을 따르면 Snyk 조직 ID 및 통합 ID가 포함된 `snyk-created-orgs.json` 파일이 생성됩니다.
4. 가져오기 데이터 생성: `snyk-api-import import:data --orgsData=snyk-created-orgs.json --source=bitbucket-cloud --integrationType=bitbucket-cloud` [가져오기 목표 데이터 생성](creating-import-targets-data-for-import-command.md)의 전체 지침
5. 가져오기 실행: `DEBUG=*snyk* snyk-api-import import` [가져오기 시작](kicking-off-an-import.md)\`\`

## 반올림 중에 새 저장소 및 조직만 재반영

초기 가져오기가 완료되면 주기적으로 새 저장소를 확인하고 이러한 단계를 따라 Snyk에 추가되도록 할 수 있습니다. 이 절차는 이전 저장소를 가져오는 단계와 유사합니다.

1. `export BITBUCKET_CLOUD_USERNAME=***`, `export BITBUCKET_CLOUD_PASSWORD=***` 및 `export SNYK_TOKEN=***` 설정
2. Snyk에서 조직 데이터 생성하고, `--skipEmptyOrg`를 사용하여 저장소가 없는 조직을 건너뛰기: `snyk-api-import orgs:data --source=bitbucket-cloud --groupId=<snyk_group_id> --skipEmptyOrg` [Snyk에서 조직 만들기](creating-organizations-in-snyk.md)의 전체 지침
3. Snyk에서 조직 생성하고, `--noDuplicateNames` 매개변수를 사용하여 이미 생성된 조직은 건너뛰기: `snyk-api-import orgs:create --file=orgs.json --noDuplicateNames` [Snyk에서 조직 만들기](creating-organizations-in-snyk.md)의 전체 지침을 따르면 Snyk 조직 ID 및 통합 ID가 포함된 `snyk-created-orgs.json` 파일이 생성됩니다.
4. 가져오기 데이터 생성: `snyk-api-import import:data --orgsData=snyk-created-orgs.json --source=bitbucket-cloud --integrationType=bitbucket-cloud` [가져오기 목표 데이터 생성](creating-import-targets-data-for-import-command.md)의 전체 지침
5. 선택 사항 - 이전에 가져온 로그를 생성하여 그룹에서 이전에 가져온 모든 저장소를 건너뛰기: `snyk-api-import-macos list:imported --integrationType=<integration-type> --groupId=<snyk_group_id>` [가져오기 시작](kicking-off-an-import.md)\`\`
6. 가져오기 실행: `DEBUG=*snyk* snyk-api-import import` [가져오기 시작](kicking-off-an-import.md)