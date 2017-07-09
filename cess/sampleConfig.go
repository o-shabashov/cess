package cess

const SampleConfig = `{
  "db": [
    {
      "engine": "mysql",
      "host": "127.0.0.1",
      "name": "db_name",
      "username": "db_user",
      "password": "db_password",
      "port": "3306"
    },
    {
      "engine": "mysql",
      "host": "my.db.host.com",
      "name": "db_name2",
      "username": "db_user",
      "password": "db_password",
      "port": "3306"
    },

  ],
  "api": [
    {
      "name": "POST API",
      "url": "https://my.post.api.com",
      "headers": {},
      "data": {
        "key": "value",
        "version": "2.0"
      },
      "test_action": "/",
      "test_method": "POST"
    },
    {
      "name": "REST",
      "url": "https://rest.my-api.com",
      "headers": {
        "client_id": "testclient",
        "client_secret": "testsecret"
      },
      "data": {},
      "test_action": "/catalog/tags",
      "test_method": "GET"
    }
  ]
}`
