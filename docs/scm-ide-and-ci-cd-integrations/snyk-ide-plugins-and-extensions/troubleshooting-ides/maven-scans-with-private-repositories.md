# 개인 저장소를 사용하는 Maven 스캔

## 문제 <a href="#problem" id="problem"></a>

Maven이 DE 플러그인 또는 CI/CD 파이프라인에서 개인 저장소를 참조하는 종속성으로 인해 실패합니다.

## 조사 <a href="#investigation" id="investigation"></a>

* 저장소를 사용하여 CLI를 독립 실행해보세요.
* Maven 프로필이나 Maven 설정이 빌드를 실행하는 데 사용되었는지 확인하세요 (예: `.mvn`에 위치).
* 파이프라인이 개인 소스에 액세스할 수 있는지 확인하세요.

종속성은 개인 저장소에서 다운로드되어야 하며 Snyk CLI는 다운로드가 실패했다는 오류 메시지를 보여줍니다.

## 해결책 <a href="#solution" id="solution"></a>

예를 들어, `-- -s your-home-directory/.m2/settings.xml` 또는 `-- -Dprofile=my-profile`와 같은 연결 정보를 포함하는 Maven 설정 또는 프로필을 설정할 수 있습니다.

이러한 매개변수는 IDE의 `추가 매개변수` 플러그인 설정을 통해 CLI로 전달할 수 있습니다. CI/CD의 경우, 만약 Snyk CLI가 직접 사용된다면 이를 파이프라인 인수로 전달해야 하며 또는 Docker 이미지에서 실행되는 경우 올바른 경로 매핑을 제공해야 합니다.