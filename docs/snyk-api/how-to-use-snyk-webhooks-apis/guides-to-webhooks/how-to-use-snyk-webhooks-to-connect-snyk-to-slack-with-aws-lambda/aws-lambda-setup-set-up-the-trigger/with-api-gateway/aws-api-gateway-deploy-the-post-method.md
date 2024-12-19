# AWS API Gateway: POST 메서드 배포

구성된 POST 메서드를 배포하여 AWS Lambda 함수가 정보를 수신할 수 있도록 할 수 있습니다.

다음 단계에 따라 POST 메서드를 배포하세요:

1. **리소스** 탭으로 이동합니다.
2. **POST**를 클릭합니다.
3. **작업** 탭에서 **API 배포**를 클릭합니다.

    <figure><img src="https://lh3.googleusercontent.com/MVnbbBF4_quh1tD-sWln5t7RdNn6kui43IRi_KHshS-jKEDVkFmsf9IpAI7Ly1Eo6ZPLnVx3WTZJ13qJdTKPWm9vr2FU1ERamyAo7N1-687QeGswSozAvB9eo8oyafqCdoDxt4nlGSDZBoyh2zf6ONWZDnN656UyodXV07glWvxCfBlkfPf7Sz8HLg" alt=""><figcaption><p>AWS API Gateway POST 메서드 리소스, 작업 탭에 API 배포 선택</p></figcaption></figure>
4. 새 API를 배포하려는 **배포 단계**를 선택합니다. 여기서는 **기본(default)** 단계를 선택합니다.

    <figure><img src="https://lh6.googleusercontent.com/xiLxfQ4yO5vb39TKW84JQe8X05sZ01stYMXtY9H8w-V2vad54nEtBI94mYQBUnGGMrmp0aEiMrn5OA9xtDnqH3BjS1UyrE0Bxsx6-Oui3XW5vxi15x0AN-rMZCWHgi2NEhNxOc-PkYbpFCJLn6n88wfDetGwi19ka0ZojM2cNLyEjeGPugScFtAcww" alt=""><figcaption><p>AWS Gateway 기본(default) 단계에 API 배포</p></figcaption></figure>
5. Lambda 함수로 이동하여 Lambda 트리거 구성에서 새 API 엔드포인트가 표시되는지 확인합니다.
6. API Gateway 상자에서 API 엔드포인트를 복사하여 Snyk 웹훅 설정에 사용합니다.

    <figure><img src="https://lh4.googleusercontent.com/EOoL3PCnKMj0HI6jkRdVsE44DwAcnFN8M8jM3Obp_FA5AXTryIHTMtGn66LlSTquVfH__0wVfjKV5bUTCxwgJzClgcdPqFTrtaq57NCd-eKBoSgFFHN49Fdqw8OsBLQai5pFsGQwGhcNpqIeto4fmXozicUeJ2A25wkh81HVmxrQH53IS-oEZZTlmQ" alt=""><figcaption><p>새 엔드포인트가 표시된 AWS Lambda 함수 트리거 구성</p></figcaption></figure>
7. 이제 API 엔드포인트가 저장되었으므로 Snyk 웹훅을 설정하세요.