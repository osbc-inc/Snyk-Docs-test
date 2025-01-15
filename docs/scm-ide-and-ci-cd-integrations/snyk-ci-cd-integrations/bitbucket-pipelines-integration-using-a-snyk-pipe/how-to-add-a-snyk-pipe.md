# Snyk 파이프를 추가하는 방법

다음 단계를 따라 Snyk 파이프를 추가하세요:

1. 기존 파이프라인을 편집하거나 새 파이프라인을 생성하는 동안 Snyk 파이프를 추가하세요. 자세한 내용은 [파이프라인](https://confluence.atlassian.com/bitbucket/configure-bitbucket-pipelines-yml-792298910.html) 및 [파이프](https://support.atlassian.com/bitbucket-cloud/docs/pipes/)에 대한 Bitbucket 문서를 참조하세요. Snyk 파이프를 추가할 때는 나머지 단계에서 제시된 지침을 따르세요.
2. Bitbucket 파이프라인 편집기를 사용하여 `.yml` 파일 구성을 업데이트하고 올바른 언어를 선택하며 Snyk 파이프를 추가할 때 Bitbucket 파이프 빌드 디렉토리를 사용하세요.
3. 빌드 단계가 모두 끝난 후 Snyk 파이프를 Bitbucket 편집기 인터페이스에 붙여넣으세요. 빌드 단계는 다음과 같은 명령어입니다: `npm install / composer install / bundle install / dotnet restore / docker build`
4. **npm publish** 또는 **docker push**와 같은 배포 단계 **이전에** 파이프를 붙여넣었는지 확인하세요.
5. **SNYK\_TOKEN** 및 **LANGUAGE**와 같은 필수 변수를 구성하세요.
6. (선택 사항) **DONT\_BREAK\_BUILD** 및 **SEVERITY\_THRESHOLD**의 취약점 발견 시 파이프 실행 실패 여부를 선택하고 필요한 경우 **MONITOR**를 활성화하세요. 자세한 내용은 [Snyk 파이프 매개변수 및 값](snyk-pipe-parameters-and-values-bitbucket-cloud.md)를 참조하세요.
7. Snyk이 파이프라인 명령어에 포함된 후에는 해당 저장소에서 매니페스트 파일(예: `package.json`, `package-lock.json`)을 찾아 스캔을 수행합니다.

결과는 다음과 같이 Bitbucket 파이프라인 출력 인터페이스에 표시됩니다:

![Bitbucket 파이프라인 출력 인터페이스 - 여기서 Snyk은 취약점을 발견하여 파이프라인이 실패한 것입니다](<../../../.gitbook/assets/Screenshot 2023-10-03 at 13.08.45.png>)

{% hint style="info" %}
빌드가 실패하면 **MONITOR**가 **True**로 설정되어 있더라도 Snyk는 프로젝트가 성공적으로 빌드될 때까지 모니터 단계로 진행하지 않습니다. Snyk의 모니터링을 활성화하려면 **DONT\_BREAK\_BUILD**를 **True**로 설정하세요. 파이프를 스캔 단계에서 실패시킬 심각도 임계값을 파이프에 알리기 위해 **SEVERITY\_THRESHOLD**를 사용할 수 있습니다. 자세한 내용은 [Snyk 파이프 매개변수 및 값](snyk-pipe-parameters-and-values-bitbucket-cloud.md)를 참조하세요.
{% endhint %}
