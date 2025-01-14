# AWS API Gateway: Snyk을 Slack에 연결하는 POST 메서드 추가

Slack이 받을 페이로드에는 메시지가 포함되어 있으므로 메시지를 수신하고 유효한 메시지인지 확인한 후 AWS Lambda 함수로 전송할 POST 메서드를 생성합니다.

다음 단계를 따라 POST 메서드를 추가합니다:

1. 만든 AWS API Gateway로 이동합니다.
2. **Resources** 를 클릭합니다.
3. 메서드를 생성하려면 **Actions** -> **Create Method** -> **Post** 로 이동합니다.
4.  Lambda 함수와 함께 작동하도록 AWS API Gateway를 구성하려면 인접한 Lambda 함수 상자에 Gateway를 추가하여 생성합니다:\
    **Lambda Function Integration** 유형을 선택합니다.\
    **Default Timeout** 를 선택합니다.

    <figure><img src="https://lh3.googleusercontent.com/3WjrkRdG1_TnfQ5w-9Ivg6J0xjic4znbfN3_76HX6quIGo5sydsEub8aMXrv9_MQsfAorYc4gUOwgIGK9JOpu0ysmI_dXFFtwlRk6LarMYu5xEgOHsJ2_9qHgKdw4Kf3MTFKX2v2EkBD5e80zC9tEZXUnFJnCfPLbaGCGv2h4omcpK10ntHdYvaVBA" alt=""><figcaption><p>AWS Lambda 함수 상자</p></figcaption></figure>
5. **Resources** 에서 새로운 **POST** 메서드를 클릭합니다.
6.  AWS Gateway POST 메서드 실행 화면의 오른쪽 상단에 있는 **Integration Request** 를 클릭합니다.

    <figure><img src="https://lh5.googleusercontent.com/_Prq2fJ7F-NE4jEiw1tqYIn0Bq-HTG0_wahTwkrod8zisAkjtKmL3O1Y0c8XEh2iYeibdkh1jWYR3V_jGvdWCbUEfE5LXd7I7cTovohFD81-NFGTvesu1jIFGKjRIWm88dAG_qcgKBQVMO7YrHvVcnERYFvr91I18K36137u2z4suVA_3P_xj8aCpQ" alt=""><figcaption><p>AWS Gateway POST 메서드 실행</p></figcaption></figure>
7. 맨 아래로 스크롤하여 **application/json Content-Type** 으로 **Mapping Template** 를 추가합니다. 다음 코드를 템플릿에 추가합니다:\
   `{`\
   `"method": "$context.httpMethod",`\
   `"body" : $input.json('$'),`\
   `"headers": {`\
   `#foreach($param in $input.params().header.keySet())`\
   `"$param": "$util.escapeJavaScript($input.params().header.get($param))"`\
   `#if($foreach.hasNext),#end #end`\
   `}`\
   `}`
8.  결과 화면이 이러한 항목을 반영하는지 확인합니다.

    <figure><img src="https://lh6.googleusercontent.com/d0AynUJWVROWc0ff2EnY_NAT7kqkrvBThXMw8d9D0StX1KKoig7ol2uDqLoMOCt7UBP7C3RYiSUrcZlg9yglP08fVXf8WBxOvGHtV25hw5PsfQC_lWfoDJkl38kIaqBNxIdg_k7W4Qg5jvQ-faPp4ySOF5j15vWRxCEjxzvAIhsHpl3s3dE2lolJdg" alt=""><figcaption><p>코드가 포함된 AWS API Gateway POST 요청 매핑 템플릿</p></figcaption></figure>
