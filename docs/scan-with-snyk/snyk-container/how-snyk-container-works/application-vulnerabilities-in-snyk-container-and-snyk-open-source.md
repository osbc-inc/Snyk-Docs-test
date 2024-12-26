# {Snyk} Container 및 {Snyk} Open Source의 응용 프로그램 취약점

Snyk} Container}는 컨테이너 내의 응용 프로그램 취약점을 감지하며 Snyk Open Source 기능과 중복됩니다.

일반적으로 Snyk가 동일한 Manifest 파일에서 종속성 그래프를 구축하고 있으면, } 스캔 결과는 일반적으로 동일합니다.

그러나 결과는 생태계 및 개발자가 응용 프로그램을 빌드하는 방식에 따라 크게 다를 수 있습니다. 컨테이너 내의 응용 프로그램은 컴파일된 응용 프로그램입니다. 그래서 일부 생태계에서, Snyk Open Source는 보다 자세한 Manifest를 스캔하여 보다 정확한 종속성 그래프를 만들 수 있습니다:

- **`golang` 프로젝트에 대한 Snyk 컨테이너**: Snyk는 Snyk Open Source처럼 종속성 목록을 액세스할 수 없습니다. 그렇기 때문에 Snyk Container는 바이너리를 역으로 분석하며 결과는 Snyk Open Source와 약간 다릅니다.
- **`npm` 패키지를 Snyk 컨테이너로 사용하는 경우**: Snyk는 종속성 목록에 액세스할 수 있습니다. 결과는 일반적으로 Snyk Open Source의 결과와 동일합니다. 자세한 내용은 [Open Source and licensing](../../../supported-languages-package-managers-and-frameworks/javascript/#open-source-and-licensing)를 참조하십시오.
- **`java` 애플리케이션에 대한 Snyk 컨테이너**: Open Source에서는 관리되지 않는 jar 파일을 포함하는 것이 가능합니다(참조: [Scan all unmanaged jar files](../../../snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files.md)). 따라서 결과가 Snyk Container와 다릅니다.
  - Snyk Container에서는 스캔이 이미지에서 찾은 모든 jar를 통과합니다(참조: [Detect application vulnerabilities in container images](../use-snyk-container/detect-application-vulnerabilities-in-container-images.md)). 또한, jar를 빌드하는 여러 가지 방법이 있으며, 사용된 방법이 Snyk Container가 종속성을 찾는 방식에 영향을 줍니다.
  - Snyk Open Source에서는 종속성의 여러 가지 잠재 버전이 있는 경우, 패키지 관리자 종속성 해결 논리가 한 버전만 선택되도록 보장합니다. 그러나 Snyk Container에서는 풀린 jar에 종속성의 다른 버전이 포함될 수 있으며, 모두 컨테이너에 존재하기 때문에 모든 버전이 보고됩니다.