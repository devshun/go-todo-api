## Go 言語 で作る最小限の TODO リスト API

### **setup**

<br />

1. .env を作成

```
cp .env.sample .env
```

2. .env に環境変数を記載

3. コンテナを起動

```
docker-compose up -d --build
```

### **CRUD TODO**

CREATE

```
curl --request POST 'http://localhost:8080/todo'
```

READ

```
curl --location --request GET 'http://localhost:8080/todo/:id'
```

UPDATE

```
curl --location --request PUT 'http://localhost:8080/todo/:id'
```

DELETE

```
curl --location --request DELETE 'http://localhost:8080/todo/:id'
```
