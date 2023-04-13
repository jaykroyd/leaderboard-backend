curl -u user:pass http://localhost:9000/api/leaderboards -X POST
curl -u user:pass -H "Content-Type:application/json" http://localhost:9000/api/leaderboards/6f5f6b7e-7067-4de9-b86a-97b6c0013c39/participants -X POST
curl -u user:pass -H "Content-Type:application/json" http://localhost:9000/api/leaderboards/6f5f6b7e-7067-4de9-b86a-97b6c0013c39/participants/e9d93ada-ffca-429b-a4fd-2e68bd83b700/score --data '{"amount":10}' -X PATCH
curl -u user:pass -H "Content-Type:application/json" http://localhost:9000/api/leaderboards/6f5f6b7e-7067-4de9-b86a-97b6c0013c39 -X DELETE
