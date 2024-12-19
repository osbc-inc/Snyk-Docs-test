# Snyk 브로커 도커 설치를 위한 고급 구성

{% hint style="info" %}
**기본값 이외의 지역을 위한 다중 테넌트 설정**\
기본값 이외의 지역에서 Snyk 브로커를 설정할 때는 특정 URL을 필요로 하는 추가 환경 변수가 필요합니다. URL 및 예제에 대한 자세한 내용은 [지역 호스팅 및 데이터 항복성, 브로커 URL](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)을 참조하십시오.
{% endhint %}

도커를 사용하여 설치하는 경우, 필요에 따라 브로커를 구성하기 위해 다음 지침을 따르십시오:

* [Docker를 사용하여 인증 방법 변경](changing-the-auth-method-with-docker.md)
* [Docker 및 Helm을 사용한 자격 증명 풀링](credential-pooling-with-docker-and-helm.md)
* [도커를 사용한 브로커 클라이언트용 HTTPS](https-for-broker-client-with-docker.md)
* [도커를 위한 내부 인증서를 사용하는 백엔드 요청](backend-requests-with-an-internal-certificate-for-docker.md)
* [도커를 사용한 프록시 지원](proxy-support-with-docker.md)
* [도커를 사용하여 인증서 검증 비활성화](disable-certificate-verification-with-docker.md)
* [도커를 사용하여 시크릿 마운트](mounting-secrets-with-docker.md)
* [Snyk Code - 브로커를 통한 클론 기능](../../git-clone-through-broker.md)
* [Snyk Open Source - 대용량 매니페스트 파일에 대한 스캔 (SCA)](snyk-open-source-scans-sca-of-large-manifest-files-docker-setup.md)
* [보안되지 않은 하향 모드](insecure-downstream-mode.md)