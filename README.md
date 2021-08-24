# go-yoda
An example API, implementing PayPI authentication.
  
# Yoda API Live Example
<a href="https://app.paypi.dev/subscribe/c2VydmljZTozMDhlNTI2NS0xY2E4LTQ5ZTgtYWYzMi01MzdjZGYyYjNhN2Y=">
  <img src="https://staging.paypi.dev/assets/getKey.svg" alt="Logo" height="30">
</a>
</br>

HTTP
```HTTP
GET /?text=Failure is the greatest teacher HTTP/1.1
Host: yoda.apis.paypi.dev
Authorization: Bearer <Your PayPI Subscriber Secret>
```

cURL
```bash
curl --location --request GET 'yoda.apis.paypi.dev?text=Failure%20is%20the%20greatest%20teacher' \
--header 'Authorization: Bearer <Your PayPI Subscriber Secret>'
```



# What does it do?
go-yoda takes in english text and returns what Yoda would have said, for example "How is your day" turns to "Your day how is".
It uses [PayPI](https://paypi.dev) to authenticate users, simplifying key management and payments.

