---
description: snyk-api-import 도구를 사용하여 가져오기 파일을 생성하고 사용하는 방법
---

# 가져오기 파일 생성 및 사용

## 가져오기 파일 작동 방식

이 기능은 Bitbucket 및 Azure에서만 작동합니다.

`snyk-scm-contributors-count` 도구가 Snyk 계정 및 SCM 계정 모두에 연결되면 도구는 Snyk가 모니터링하는 프로젝트/저장소와 그렇지 않은 것을 알아냅니다.

`importConfDir` 및 `importFileRepoType` 플래그를 명령에 적용하면 도구가 가져오기 파일을 생성하여 미모니터링 저장소 데이터로 채웁니다. 이 파일은 [**snyk-api-import** ](creating-and-using-the-import-file.md#using-the-snyk-api-import-tool)도구와 함께 사용하여 빠진 저장소를 Snyk 계정에 가져올 수 있습니다.

* Snyk 토큰이 내보내진 경우 및 관련 Snyk 계정에 도구가 스캔하는 특정 SCM에 대한 통합이 설정된 경우 도구는 `snyk-api-import` 도구에서 필요한 OrgID 및 IntegrationID를 찾아 일치시키려고 시도하고 가져오기 파일에 자동으로 추가합니다.
* Snyk 토큰이 내보내지지 않았거나 사용자가 아직 Snyk 계정이 없는 경우, 이 기능을 사용하여 SCM의 모든 저장소를 매핑하고 `snyk-api-import` 도구에서 나중에 사용할 가져오기 파일을 생성할 수 있습니다. 이 경우나 도구가 OrgID 또는 IntegrationID를 찾을 수 없는 경우 도구는 사용자에게 이러한 ID를 제공하라는 프롬프트를 표시한 다음 가져오기 파일에 자동으로 추가합니다.

## 가져오기 플래그

`importConfDir` - 이 플래그는 미모니터링 저장소에 대한 쿼리를 수행하고 가져오기 JSON 파일을 생성할 폴더 경로(쓰기 권한 포함)를 기대합니다. 예를 들어:

```
snyk-scm-contributors-count <command> --token TOKEN -- importConfDir /snyk/import/
```

기본적으로 이 명령은 SCM을 스캔할 때 찾은 모든 유형의 미모니터링 저장소로 JSON 가져오기 파일을 채웁니다.

`importFileRepoType` - 이 플래그는 `all,` `private`, 또는 `public`의 값을(대소문자 구분하지 않음) 설정하여 가져오기 파일을 주어진 저장소 유형의 데이터로 채울 수 있습니다. 예를 들어:

```
snyk-scm-contributors-count <command> --token TOKEN -- importConfDir /snyk/import/ --importFileRepoType 'private'
```

{% hint style="info" %}
가져오기 파일은 사용자로부터 OrgID 및 IntegrationID를 제공받아야 올바른 조직 및 통합으로 가져오기할 수 있습니다.

도구는 이 두 값을 Snyk에서 찾으려고 시도할 것이며(전달된 SNYK\_TOKEN이 있고 Snyk의 org 매핑이 SCM과 일치하는 경우), 도구가 해당 값을 찾지 못하면 사용자에게 명령줄에서 이 ID를 제공하라는 프롬프트가 표시됩니다.

사용자가 OrgID 및 IntegrationID 값을 한 번 제공한 후에는 이러한 값을 가져오기 파일의 모든 항목에 대해 설정됩니다(즉, 모든 가져온 저장소가 Snyk에서 동일한 조직 안에 있게 됨).
{% endhint %}

## snyk-api-import 도구 사용

snyk-api-import 도구는 Snyk 계정으로 새 저장소를 안전하고 견고하게 가져오는 데 도움을 줍니다.

이 도구는 가져올 저장소 데이터가 포함된 json 파일을 필요로 합니다. 이 파일은 앞 섹션에서 설명한 대로 **snyk-contributors-count** 도구에 의해 자동으로 생성될 수 있습니다.

## 자세한 정보

**snyk-api-import** 도구에 대한 자세한 정보는 다음을 참조하십시오:

* [https://github.com/snyk/snyk-api-import](https://github.com/snyk/snyk-api-import)
* [https://github.com/snyk/snyk-api-import/blob/master/docs/import.md](https://github.com/snyk/snyk-api-import/blob/master/docs/import.md)