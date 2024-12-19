# Snyk CLI를 위한 오픈 소스

{{Snyk Open Source}}는 매니페스트 파일을 스캔합니다. 스캔을 통해 Snyk는 매니페스트 파일에 표시된 구조의 계층 구조 트리를 생성합니다. 직접 및 간접 (위계적) 의존성 및 다른 패키지가 소개된 지점을 포함합니다.

이 트리를 구축한 후, Snyk는 취약성 데이터베이스를 사용하여 의존성 트리의 어느 부분에서든 패키지에 있는 취약성을 찾습니다. Snyk를 사용하면 프로젝트를 분석하는 것이 원본에서 프로젝트를 수정하는 것보다 더 쉬워집니다. 취약한 패키지가 도입된 지점을 빠르게 식별할 수 있습니다.

## 오픈 소스 프로젝트 테스트

알려진 취약점을 찾으려면 프로젝트를 테스트하세요:
* 프로젝트가 포함된 폴더로 이동 (`cd ~/projects/myproj/`)
* `$ snyk test`를 실행합니다.

`snyk test` 명령은 모든 로컬 종속성을 식별하고 알려진 취약성을 위해 Snyk 서비스에 쿼리를 보냅니다. `snyk test`는 찾은 문제와 추가 정보를 표시합니다. snyk 테스트 결과에 대한 정보는 [{{Snyk Open Source}} CLI 결과 검토](review-the-snyk-open-source-cli-results.md)를 참조하세요.

{% hint style="info" %}
Node.js, Ruby 및 Java 프로젝트의 경우 `snyk test`는 수정 단계를 제안하기도 합니다.
{% endhint %}

## 프로젝트 유형을 감지하는 데 사용하는 Snyk 파일

`snyk test`를 실행하면 매니페스트 파일을 자동으로 감지하여 프로젝트 유형을 결정하려고 합니다. Snyk가 프로젝트 유형을 자동으로 감지하는 데 사용하는 파일은 다음과 같습니다. (다음에 나열된 것에만 국한되지 않음)

* yarn.lock
* package-lock.json
* package.json
* Gemfile.lock
* pom.xml
* build.gradle
* build.sbt
* Pipfile
* requirements.txt
* Gopkg.lock
* vendor/vendor.json
* obj/project.assets.json
* packages.config
* composer.lock
* build.gradle.kts
* go.mod

{% hint style="info" %}
여러 매니페스트 파일을 분석하려면 [여러 매니페스트 파일 스캔](use-options-to-customize-the-snyk-test-command.md#scan-multiple-manifest-files)을 참조하세요.
{% endhint %}

Snyk가 파일을 분석하고 트리를 작성하는 방식은 다음과 같은 요소에 따라 달라집니다:
* 매니페스트 파일 유형에 따라 결정되는 사용하는 [언어 및 패키지 관리자](../../../supported-languages-package-managers-and-frameworks/)
* 스캔 방법, [Snyk CLI](../../) 또는 [Git 리포지토리 통합](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)을 사용하여 프로젝트 가져오기

` snyk test`를 일반적으로 사용되는 옵션과 함께 실행하는 팁은 [snyk test 명령 사용을 위한 옵션 사용](use-options-to-customize-the-snyk-test-command.md)을 참조하십시오.

더 많은 지원되는 언어에 대한 정보는 [지원되는 언어, 패키지 관리자 및 프레임워크](../../../supported-languages-package-managers-and-frameworks/)를 참조하십시오.