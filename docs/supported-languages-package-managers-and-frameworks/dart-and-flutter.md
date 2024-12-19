# Dart and Flutter

## 적용 가능성

Snyk는 Dart 및 Flutter에 대해 **오픈 소스에만 지원**합니다.

Snyk 제품을 사용하여 애플리케이션을 가져오거나 테스트하고 모니터링할 수 있는 언어 가용성을 확인하십시오.

사용 가능한 기능:

- SCM을 통해 앱 가져오기: N/A
- CLI 및 IDE를 통해 앱 테스트하거나 모니터링하기: N/A
- `pkg:pub`를 사용하여 앱의 SBOM 테스트
- `pkg:pub`를 사용하여 앱의 패키지 테스트

## 패키지 관리자 및 지원되는 파일 확장자

Snyk는 Dart 및 Flutter에 대해 Pub를 패키지 관리자로 지원하며 [pub.dev](https://pub.dev/)를 패키지 레지스트리로 지원하며 어떠한 파일 형식도 지원하지 않습니다.

## 프레임워크 및 라이브러리

Snyk는 Dart 및 Flutter에 대해 프레임워크 및 라이브러리를 지원하지 않습니다.

## 기능

Snyk는 Dart 및 Flutter에 대해 어떠한 Snyk 기능도 지원하지 않습니다.

## API를 사용하여 pub 패키지 테스트

{% hint 스타일="info" %}
Snyk API는 엔터프라이즈 플랜에서만 사용할 수 있습니다. 자세한 정보는 [플랜 및 가격](https://snyk.io/plans)을 참조하십시오.
{% endhint %}

Snyk는 API 엔드포인트 [패키지에 대한 문제 목록](../snyk-api/reference/issues.md#orgs-org_id-packages-purl-issues)를 사용하여 Pub 패키지 관리자에서 오픈 소스 패키지를 테스트하는 것을 지원합니다:

```
GET /orgs/{org_id}/packages/{purl}/issues 엔드포인트
```

이 엔드포인트는 패키지의 알려진 취약점을 반환합니다. 자세한 정보는 [패키지에 대한 문제 목록](../snyk-api/how-to-use-snyk-sbom-and-list-issues-apis/list-issues-for-a-package.md) 페이지를 참조하십시오.

## SBOM CLI를 사용하여 pub 패키지 테스트

[SBOM CLI](../snyk-cli/commands/sbom.md)를 사용하여 SBOM을 테스트할 수도 있습니다. 먼저 SBOM 파일을 생성해야 합니다. 예를 들어, 다음과 같이 `cdxgen`을 사용하여 SBOM을 추출하고 이를 다음과 같이 Snyk CLI로 보낼 수 있습니다:

```
cdxgen -t pub -o pub-sbom.json \
  && snyk sbom test --experimental --file pub-sbom.json
```

## Flutter 앱에서 플랫폼 의존성 (iOS, macOS, Android) 테스트

Flutter 애플리케이션은 종종 애널리틱스, 하드웨어 액세스 또는 기존 기능 통합과 같은 하위 수준 작업을 처리하기 위해 네이티브 플랫폼 의존성에 의존합니다. 이러한 의존성은 기능을 확장하기 위해 pub 패키지를 통해 추가하거나 Gradle 또는 Cocoapods와 같은 빌드 시스템에 직접 통합될 수 있습니다.

Snyk의 일반 오픈 소스 지원으로 이러한 패키지를 스캔할 수 있지만, CLI 도구에서 이러한 패키지를 이용할 수 있도록 모든 앱 빌드가 필요합니다.

모든 관련 플랫폼에 대해 애플리케이션을 빌드해 시작하십시오. 이렇게 하면 `pub`이 필요한 모든 패키지를 가져오고 Flutter 빌드 시스템이 네이티브 빌드 시스템에 필요한 링크를 설정합니다.

```
flutter build apk --debug
flutter build ios --debug --no-codesign
flutter build macos --debug
```

다음으로, 네이티브 의존성을 스캔하기 위해 `snyk monitor` 명령을 실행하십시오.

```
snyk monitor --all-projects --exclude=example,.symlinks
```

`--exclude` 매개변수를 사용하여 중복을 제거하고 예제 애플리케이션을 무시합니다. 예제 애플리케이션은 플러그인 소스 코드 일부이지만 정규 애플리케이션 빌드에 포함되지 않습니다.

이제 Snyk 웹 UI에서 서드파티 플러그인에 의해 도입된 포함하여 모든 네이티브 의존성을 볼 수 있을 것입니다.

<figure><img src="../.gitbook/assets/image (571).png" alt=""><figcaption><p>Flutter 앱에서 의존성을 보여주는 Snyk 프로젝트 페이지</p></figcaption></figure>

도움이 필요한 경우, [Snyk 지원팀과 연락](https://support.snyk.io)하십시오.