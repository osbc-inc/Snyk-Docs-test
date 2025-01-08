# Elixir

## 적용 가능성

Elixir에 대한 Snyk은 **Snyk 오픈 소스에 대해서만** 지원됩니다.

Snyk 제품을 사용하여 응용 프로그램을 가져오고 테스트하거나 모니터링할 수 있는 언어 가용성을 확인하세요.

가능한 기능:

* SCM을 통한 응용 프로그램 가져오기: 불가
* CLI 및 IDE를 통한 응용 프로그램 테스트 또는 모니터링은 Snyk Open Source에서만 가능합니다.
* `pkg:hex`를 사용하여 응용 프로그램의 SBOM 테스트
* `pkg:hex`를 사용하여 응용 프로그램의 패키지 테스트

Snyk CLI를 사용하여 코드 분석을 위한 자세한 정보는 [Snyk 코드용 Snyk CLI](../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/)를 참조하십시오.

## 패키지 매니저 및 지원되는 파일 확장자

Snyk for Elixir는 패키지 매니저로 [Mix](https://hexdocs.pm/mix/Mix.html)/[Hex](https://hex.pm/)를 지원하며 패키지 레지스트리로 [hex.pm](https://hex.pm/)을 지원하며 어떠한 파일 형식도 지원하지 않습니다.

## 프레임워크 및 라이브러리

Snyk for Elixir에는 사용 가능한 프레임워크 및 라이브러리가 없습니다.

## 기능

보고서 기능은 Snyk for Elixir에서 지원됩니다.

## Elixir를 위한 Snyk CLI

{% hint style="info" %}
의존성을 스캔하려면 먼저 Elixir와 Mix를 설치해야 합니다. 자세한 내용은 [Elixir 설치 지침](https://elixir-lang.org/install.html)을 참조하세요.
{% endhint %}

Snyk은 [CLI](../snyk-cli/)를 사용하여 취약점을 테스트하여 Elixir 프로젝트의 보안을 스캔합니다.

Mix는 Elixir 프로젝트를 컴파일하고 테스트하고 만드는 빌드 도구입니다. Mix는 Hex 패키지 매니저와 통합하여 종속성을 관리합니다.

Snyk은 `mix.exs` 및 `mix.lock` 파일을 분석하여 프로젝트의 종속성 트리를 빌드합니다. `mix.lock` 파일이 있어야 하며 `mix.exs` 파일과 동기화되어야 합니다. Snyk이이 트리를 빌드한 후에는 패키지의 취약점을 종속성 트리 어디에서든 찾기 위해 [취약점 데이터베이스](https://snyk.io/vuln)를 사용합니다.

### **프로젝트 명명**

Snyk UI에서 프로젝트는 주로 `mix.exs` 파일의 `Mix.Project`에서 내보낸 `project/0` 함수의 `app` 키워드에 따라 명명됩니다.

이름을 재정의하려면 `--project-name` CLI 옵션을 사용하십시오.

### **Mix 어머렐라 프로젝트**

Mix 어머렐라 프로젝트를 테스트하는 경우, Snyk는 어머렐라 프로젝트임을 감지하고 자동으로 모든 하위 앱을 포함합니다.

주 `mix.exs`와 함께 각 앱 `mix.exs`는 Snyk UI에 별도의 프로젝트로 나타납니다. 앱의 경로에 따라 명명됩니다.

Snyk은 Mix 프로젝트에 나열된 모든 `:hex` 패키지 및 해당 종속성 및 취약점을 완전히 지원합니다.

Hex 지원에는 Elixir 및 Erlang 패키지가 모두 포함됩니다.

Snyk은 또한 `:path`, `:git`, `:github` 종속성에 대해 제한적인 지원을 제공하지만 해당 종속성 또는 취약점을 지원하지는 않습니다.

* `:path` 종속성은 이름별로 종속성 트리에 나타납니다.
* `:git` 및 `:github` 종속성은 저장소 URL 및 버전(`mix.exs` 파일에서 정의된 `:branch`, `:tag`, 또는 `:ref`)에 따라 종속성 트리에 나타납니다.

## IDE 및 CI/CD

통합 개발 환경에 대해서는 [Snyk IDEs](../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/)를 참조하십시오.

지속적인 통합/지속적인 전달 워크플로우를 사용하는 경우 자동화 소프트웨어와 통합을 기반으로 Snyk을사용하여 스캔할 수 있습니다.

## Elixir를 위한 Snyk 문제 해결

`asdf`를 사용할 때 버전을 설정하려면 `asdf global elixir <선택한 버전>`을 실행해야 합니다.

도움이 필요한 경우, [Snyk 지원팀에 문의](https://support.snyk.io)하십시오.
