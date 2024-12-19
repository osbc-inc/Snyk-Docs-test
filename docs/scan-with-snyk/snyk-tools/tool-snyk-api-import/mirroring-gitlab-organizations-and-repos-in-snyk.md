# Snyk에서 GitLab 조직 및 저장소 미러링

사용 가능한 유틸리티에서 네 개의 명령어를 사용하여 GitLab 저장소 전체를 Snyk에 가져올 수 있습니다. 계속하기 위해 GitLab 토큰과 Snyk 토큰을 환경 변수로 구성해야 합니다.

## GitLab 저장소 가져오기 일반 단계

자세한 정보는 개별 문서 페이지를 참조하십시오. 일반적인 단계는 다음과 같습니다:

1. `export GITLAB_TOKEN=***` 및 `export SNYK_TOKEN=***`
2. 조직 데이터 생성 예시: `snyk-api-import orgs:data --source=gitlab --groupId=<snyk_group_id>` 자세한 지침: [Snyk에서 조직 생성](creating-organizations-in-snyk.md)
3. Snyk에서 조직 생성: `snyk-api-import orgs:create --file=orgs.json` [Snyk에서 조직 생성](creating-organizations-in-snyk.md)의 전체 지침을 따르면, 가져오에 필요한 Snyk 조직 ID 및 통합 ID가 있는 `snyk-created-orgs.json` 파일이 생성됩니다.
4. 가져오 데이터 생성: `snyk-api-import import:data --orgsData=snyk-created-orgs.json --source=gitlab --integrationType=gitlab` 자세한 지침: [가져오 목표 데이터 생성](creating-import-targets-data-for-import-command.md)
5. 가져오 실행: `DEBUG=*snyk* snyk-api-import import` 전체 지침: [가져오 시작](kicking-off-an-import.md)

## 미러링하는 동안 새 저장소 및 조직만 다시 가져오기

초기 가져오가 완료되면 정기적으로 새 저장소를 확인하고 다음 단계를 따라 Snyk에 추가되었는지 확인할 수 있습니다. 이는 저장소를 가져오는 이전 단계와 유사합니다.

1. `export GITLAB_TOKEN=***` 및 `export SNYK_TOKEN=***`
2. Snyk에서 조직 데이터 생성하고 저장소가 없는 경우 건너 뜁니다. `--skipEmptyOrg`를 사용하여 `snyk-api-import orgs:data --source=gitlab --groupId=<snyk_group_id> --skipEmptyOrg` 자세한 지침: [Snyk에서 조직 생성](creating-organizations-in-snyk.md)
3. Snyk에서 조직 생성하되, 이미 생성된 경우 `--noDuplicateNames` 매개변수를 사용하여 건너 뜁니다. `snyk-api-import orgs:create --file=orgs.json --noDuplicateNames` [Snyk에서 조직 생성](creating-organizations-in-snyk.md)의 전체 지침을 따르면, 가져오에 필요한 Snyk 조직 ID 및 통합 ID가 있는 `snyk-created-orgs.json` 파일이 생성됩니다.
4. 가져오 데이터 생성: `snyk-api-import import:data --orgsData=snyk-created-orgs.json --source=gitlab --integrationType=gitlab` 자세한 지침: [가져오 목표 데이터 생성](creating-import-targets-data-for-import-command.md)
5. 선택 사항 - 이전에 가져온 로그 생성하여 그룹의 이전에 가져온 모든 저장소를 건너 뜁니다: `snyk-api-import-macos list:imported --integrationType=gitlab --groupId=<snyk_group_id>` 전체 지침: [가져오 시작](kicking-off-an-import.md)
6. 가져오 실행: `DEBUG=*snyk* snyk-api-import import` 전체 지침: [가져오 시작](kicking-off-an-import.md)