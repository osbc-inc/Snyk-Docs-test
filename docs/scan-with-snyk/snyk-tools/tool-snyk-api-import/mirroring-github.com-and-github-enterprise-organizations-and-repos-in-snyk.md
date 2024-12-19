# Snyk에서 GitHub.com 및 GitHub Enterprise 조직 및 저장소의 미러링

Snyk에서 GitHub 및 GitHub Enterprise 저장소 전체를 가져오기 위해 사용 가능한 유틸리티에서 네 가지 명령을 사용할 수 있습니다. 계속하기 위해 GitHub 토큰과 Snyk 토큰을 환경 변수로 구성해야 합니다.

## GitHub 및 GitHub Enterprise 저장소 가져오는 일반적인 단계

자세한 정보는 개별 설명서 페이지를 참조하십시오. 일반적인 단계는 다음과 같습니다:

1. `export GITHUB_TOKEN=***` 및 `export SNYK_TOKEN=***`
2. 조직 데이터 생성, 예를 들어, `snyk-api-import orgs:data --source=github --groupId=<snyk_group_id>` 전체 지침: [Snyk에서 조직 생성](creating-organizations-in-snyk.md)
3. Snyk에서 조직 생성: `snyk-api-import orgs:create --file=orgs.json` [Snyk에서 조직 생성](creating-organizations-in-snyk.md)의 전체 지침을 따르면, 가져오기에 필요한 Snyk 조직 ID 및 통합 ID가 포함된 `snyk-created-orgs.json` 파일을 생성할 수 있습니다.
4. 가져오기 데이터 생성: `snyk-api-import import:data --orgsData=snyk-created-orgs.json --source=github --integrationType=github` 전체 지침: [가져오기 대상 데이터 생성](creating-import-targets-data-for-import-command.md)
5. 가져오기 실행: `DEBUG=*snyk* snyk-api-import import` 전체 지침: [가져오기 시작](kicking-off-an-import.md)

## 미러링하는 동안 새로운 저장소 및 조직만 다시 가져오기

초 기 가져오기가 완료된 후, 주기적으로 새로운 저장소를 확인하고 이를 Snyk에 추가할 수 있도록 다음 단계를 따릅니다. 이 절차는 저장소를 가져오는 이전 단계와 유사합니다.

1. `export GITHUB_TOKEN=***` 및 `export SNYK_TOKEN=***`
2. Snyk에서 조직 데이터 생성하고 `--skipEmptyOrg`를 사용하여 저장소가 없는 것을 건너뛰기 : `snyk-api-import orgs:data --source=github --groupId=<snyk_group_id> --skipEmptyOrg` 전체 지침: [Snyk에서 조직 생성](creating-organizations-in-snyk.md)
3. Snyk에서 조직 생성하고 `--noDuplicateNames` 매개변수로 이미 만들어진 조직을 건너뛰기: `snyk-api-import orgs:create --file=orgs.json --noDuplicateNames` [Snyk에서 조직 생성](creating-organizations-in-snyk.md)의 전체 지침을 따르면, 가져오기에 필요한 Snyk 조직 ID 및 통합 ID가 포함된 `snyk-created-orgs.json` 파일을 생성할 수 있습니다.
4. 가져오기 데이터 생성: `snyk-api-import import:data --orgsData=snyk-created-orgs.json --source=github --integrationType=github` 전체 지침: [가져오기 대상 데이터 생성](creating-import-targets-data-for-import-command.md)
5. 선택 사항 - 이전에 가져온 로그를 생성하여 그룹 내 모든 이전에 가져온 저장소를 건너뛰기: `snyk-api-import-macos list:imported --integrationType=<integration-type> --groupId=<snyk_group_id>` 전체 지침: [가져오기 시작](kicking-off-an-import.md)
6. 가져오기 실행: `DEBUG=*snyk* snyk-api-import import` 전체 지침: [가져오기 시작](kicking-off-an-import.md)