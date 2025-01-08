# PHP

## 적용 가능성

Snyk은 [코드 분석을 위한 PHP](php-for-code-analysis.md)와 [오픈 소스를 위한 PHP](php-for-open-source.md)를 지지합니다.

어플리케이션으로 가져올 수 있는 언어 가용성을 확인하고, Snyk 제품을 사용하여 테스트하거나 모니터링할 수 있습니다.

{% hint style="info" %}
**지원되는 PHP 버전**

PHP 버전 5.2부터 8.0까지 사용할 수 있습니다.
{% endhint %}

사용 가능한 기능:

* SCM 가져오기, Snyk Open Source, Snyk Code에서 사용 가능합니다.
* CLI 및 IDE를 통해 앱을 테스트하거나 모니터링할 수 있습니다. Snyk Open Source, Snyk Code에서 사용 가능합니다.
* `pkg:composer`를 사용하여 앱의 SBOM 테스트하기
* `pkg:composer`를 사용하여 앱의 패키지 테스트하기

## 패키지 관리자 및 지원되는 파일 확장자

Snyk은은 PHP를 위해 Composer를 패키지 관리자로 지원하며, [packagist.org](https://packagist.org/)를 패키지 레지스트리로 지원하며, 다음 파일 형식을 지원합니다:

* Snyk Open Source: `composer.json`, `composer.lock`
* Snyk Code: `.php`, `.phtml`, `.module`, `.inc`, `.install`,

## 프레임워크 및 라이브러리

다음 프레임워크 및 라이브러리가 Snyk에서 PHP를 지원합니다:

* Laravel - Partial
* llphant - Comprehensive
* openai-php/client - Comprehensive
* orhanerday/open-ai - Comprehensive
* Pclzip - Comprehensive
* Symfony - Partial
* theodo-group/llphant - Comprehensive

## 기능

다음 기능이 Snyk에서 PHP를 지원합니다:

| Snyk Open Source                      | Snyk Code                                               |
| ------------------------------------- | ------------------------------------------------------- |
| <ul><li>라이선스 스캔</li><li>보고서</li></ul> | <ul><li>보고서</li><li>사용자 정의 규칙</li><li>인터파일 분석</li></ul> |

## PHP를 위한 Snyk CLI

PHP를 실행할 때 사용할 수 있는 고유한 옵션이 없습니다.

## SCM 통합 및 PHP

PHP 프로젝트는 사용 가능한 Snyk SCM 통합에서 가져올 수 있습니다. 프로젝트가 가져와진 후, Snyk은 지원하는 매니페스트 파일을 기반으로 프로젝트를 분석합니다.

가져오기할 프로젝트를 선택한 후, Snyk은 이러한 매니페스트 파일을 기반으로 의존성 트리를 빌드합니다. 다음 파일 둘 다 필요합니다:

* `composer.json`
* `composer.lock`

만약 리포지토리에 `composer.lock` 파일이 없다면, 가져오기는 composer.json 매니페스트를 처리하지 않습니다.

기본적으로 Snyk는 프로덕션 의존성을 스캔합니다. Snyk 웹 UI를 통해 개발 의존성(`require_dev`)을 취약점 스캔에 포함할 것인지 여부를 구성할 수 있습니다.

언어 기본 설정을 업데이트하려면:

1. 계정에 로그인하고 관리하려는 관련 그룹 및 조직으로 이동합니다.
2. **설정**을 선택한 후 **언어**를 선택합니다.
3. PHP에 대한 **설정 편집**을 선택하고 PHP 프로젝트에서 개발 및 프로덕션 의존성을 모두 포함하도록 설정하기 위해 **개발 의존성 스캔**을 선택합니다.
4. **설정 업데이트**를 선택합니다.

이러한 설정은 새로 가져온 프로젝트와 다시 테스트된 모든 기존 프로젝트에 적용됩니다.

## Snyk for PHP 문제 해결

도움이 필요한 경우, [Snyk 지원팀에 문의](https://support.snyk.io)하십시오.
