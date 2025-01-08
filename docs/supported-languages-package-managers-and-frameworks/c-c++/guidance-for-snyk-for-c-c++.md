# C/C++용 Snyk에 대한 지침

이 페이지에서는 언어 및 패키지 관리자에 대한 고려 사항을 검토하여 기술 스택에서 Snyk을 효과적으로 적용할 수 있도록 도와줍니다.

## 코드 분석

* Snyk은 분석을 수행하기 위해 컴파일하거나 빌드를 필요로하지 않습니다.
* Snyk Code는 소스 코드를 직접 분석합니다.
* 컴포넌트를 사전에 컴파일하는 경우 스캔 중에 소스를 사용할 수 있도록 해야 합니다.

## 오픈 소스 및 라이선스

npm 또는 Maven과 같은 패키지 관리자의 경우, Snyk은일반적으로 `snyk test` 및 `snyk monitor`의 관리되는 오픈 소스 기능을 사용합니다. C/C++의 경우, Snyk은 `--unmanaged`를 추가하여 관리되지 않는 종속성을 지원합니다.

{% hint style="info" %}
Snyk은 빌드에 편입되지 않으며 스캔을 수행하기 위해 빌드에 의존하지 않습니다. Snyk은 소스 코드에서 분석을 수행합니다.
{% endhint %}

* 오픈 소스 소스 코드가 있어야 합니다.
* Snyk은 파일을 지문 처리하고 이를 Snyk 데이터베이스와 비교하여 패키지, 버전, 라이선스 및 취약점을 식별합니다.

## Snyk 통합 및 일반적인 사용 패턴

### IDE

#### Snyk Code와 함께

추가 옵션이 필요하지 않습니다. IDE 내의 Snyk 플러그인은 결과를 표시하기 위한 뷰를 제공합니다.

**Snyk Open Source와 함께**

IDE 설정의 **추가 매개변수**에서 `--unmanaged` 옵션을 입력하여 C/C++ 오픈 소스 종속성을 스캔할 수 있습니다.

<div align="left"><figure><img src="https://lh6.googleusercontent.com/1j-2sJjuVejBJ6nARpaAx2uhdhqT7G3XyNCGZqFxBXJV9ujqRHBYiwInr_mFT7SH-fnhG6iUysKxzYKluPG1f3xUKyb2q-JycA_0QevtaS3hdm4I7-QT7M5benqzWkIe5N-7L3czV-F84_xUR5yl7k0" alt="의존성 스캔"><figcaption><p>의존성 스캔</p></figcaption></figure></div>

### CLI 팁과 트릭

#### 코드베이스

Snyk는 분석을 수행하기 위해 빌드에 의존하지 않습니다. 소스 코드만 필요합니다.

터미널에서 소스 코드의 디렉터리를 열고 다음 명령을 실행하세요:

```
snyk code test
```

{% hint style="info" %}
컴포넌트를 사전에 컴파일하는 경우에도 최상의 결과와 커버리지를 얻으려면 소스 코드가 여전히 있어야 합니다.
{% endhint %}

보고용으로, [snyk-to-html](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md) 플러그인을 사용하여 보고서 아티팩트를 생성할 수 있습니다. 또한 프로그래밍 방식으로 결과에 액세스하기 위해 JSON 및 SARIF 출력 기능이 있으며 각각 `--json` 및 `--sarif` 옵션을 사용할 수 있습니다. 자세한 내용은 [테스트 결과를 JSON 또는 SARIF 파일로 내보내기](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/view-snyk-code-cli-results.md#export-test-results)를 참조하세요.

#### 오픈 소스 라이브러리

C/C++의 경우, 오픈 소스를 사용하여 라이선스 준수 문제와 관련된 알려진 보안 문제를 분석하려면 `--unmanaged` 옵션을 사용하세요.

```
snyk test --unmanaged
```

세부 내용은 [Snyk for C/C++](./)를 참조하세요.

* 테스트를 위해서는 오픈 소스 소스 코드가 있어야 합니다. 소스 코드는 vendor 폴더에 위치할 수 있습니다.
* 오픈 소스를 사전에 컴파일하는 경우에도 오픈 소스 코드가 여전히 있어야 합니다. Snyk이 기존 지식 베이스와의 정확한 비교를 수행하기 위해서는 오픈 소스 코드가 유지되어야 합니다.

모니터링 및 보고 공유에 대해 동일한 명령을 사용하세요:

```
snyk monitor --unmanaged --org=<org-id>
```

여기서 `org-id`는 Snyk 웹 인터페이스의 조직 설정에서 찾을 수 있습니다. 조직 ID는 필수가 아니지만 사용하는 것이 강력히 권장됩니다. 와 마찬가지로 [snyk-to-html](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md) 플러그인을 사용하여 보고서 아티팩트를 생성할 수 있습니다.

* 개별 또는 개인 스캔의 경우 CLI 또는 IDE를 사용하여 결과를 업로드하려면 `snyk monitor --unmanaged` 명령을 사용하세요.
  * 그러나 Snyk은은 개별 스캔이 잡음을 유발하지 않도록 개인 폴더로 결과를 보내고 프로젝트 설정에서 예약된 스캔을 비활성화하는 것이 좋다고 권장합니다.
  * 이는 라이선스 및 정책 정보를 확인할 수 있습니다.
* CI/CD와 같은 자동화된 스캔의 경우 `snyk monitor --unmanaged`를 사용하여 결과를 선택한 조직에 전송하세요. 이는 라이선스 및 정책 정보를 확인할 수 있습니다.

#### 종속성 목록

코드베이스에서 오픈 소스 스캔을 수행할 때 `--print-deps`를 사용하여 코드베이스에서 발견된 종속성의 자세한 목록을 얻을 수 있습니다.

이는 주어진 매치의 신뢰 수준을 식별하는 추가 이점이 있습니다. 신뢰 수준이 상당히 낮아지면 (< 90% 신뢰 수준), 해당 파일이 수정된 것이라고 생각할 수 있으며 원본 소스가 아닐 수 있으므로 해당 사항을 조사해야 합니다.

```
snyk test --unmanaged --print-deps
```

다음 이미지에서 이 목록은 문제 목록보다 먼저 인쇄됩니다:

<figure><img src="https://lh5.googleusercontent.com/x4y1uIQ2fCFX956f1eP4664i6VKEgK6eOOddlAZ4p4WnQWJu1t_ugSOpL394KEnuzSIPRs08gNAsmjvPa-GAV0C-975esRdy0EPDY7WImG1-SXSOFO0TIAVfh_Jp2DLYc6bm7iZu55UbE3Boh4TNk_I" alt="의존성 목록 보기"><figcaption><p>의존성 목록 보기</p></figcaption></figure>

#### 라이선스 정책

{% hint style="info" %}
**기능 제공 여부**\
[라이선스 완수](../../scan-with-snyk/snyk-open-source/scan-open-source-libraries-and-licenses/open-source-license-compliance.md) 기능은 Snyk 팀 또는 엔터프라이즈 [플랜](https://snyk.io/plans)에서 사용할 수 있습니다.
{% endhint %}

이 기능을 사용하면 회사가 오픈 소스 응용 프로그램에 대한 라이선스 정책을 작성하여 사용이 승인되지 않은 라이선스를 표시할 수 있습니다. Snyk이 승인되지 않은 라이선스에 일치하는 항목을 감지하면 알림을 보냅니다. 해당 알림에는 라이선스의 이름과 라이선스 정책 텍스트가 포함됩니다.

라이선스 정책 텍스트는 관리자에 의해 라이선스 문제와 연결되어 있습니다. 이 텍스트는 라이선스 문제에 대한 사용자 정의 지침이나 해당 라이선스 문제가 정책에 어떻게 위반되는지에 대한 정보를 제공합니다.

다음은 화면 하단에 있는 라이선스 정책 텍스트 예제를 보여줍니다. 라이선스가 발견된 경우 해결 방법에 대한 지침이 제공됩니다.

<div align="left"><figure><img src="https://lh4.googleusercontent.com/lIn5JFEyaZaTNMVenBoeGIgTpC6YHxpmAjK947z5ISPlHV1rlOvPNCLyzXxsGNj65AAlGn6ff9dF4lHVsVFYMaKXWC939tasD91k98xcDv_Ske6Dz7goMXl5lByyqg6ptvvqaK0UEqLSdzUU9GKrW4U" alt="라이선스 정책 텍스트 예제"><figcaption><p>라이선스 정책 텍스트 예제</p></figcaption></figure></div>

#### 대체 테스트 옵션

고급 종속성 관리 전략을 개발하는 경우 일반적이고 자주 사용되는 패키지 관리자를 사용하지 않을 수 있습니다. 이를 위해 C++의 경우, 애플리케이션에 포함된 오픈 소스 패키지와 버전을 알고 있지만 소스 코드가 없는 경우 [패키지의 이슈 목록 보기](../../snyk-api/reference/issues.md#orgs-org_id-packages-purl-issues) 엔드포인트를 사용하여 분석할 수 있습니다.

### 옵션 및 플러그인

* 로컬에서 또는 빌드할 때 보고서를 생성하는 데 도움이 되는 경우 [snyk-to-html 플러그인](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-to-html.md)을 참조하세요.
* 프로그래밍 방식으로 액세스할 수 있는 출력을 생성하는 데 사용할 수 있는 `--json` 및 `--sarif` 옵션을 확인하세요.
* 고급 필터링 옵션에 대해 알아보려면 [snyk-filter](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md)를 확인하세요.
