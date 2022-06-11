curl http://localhost:9000/api/leaderboard -X POST -u user:pass
curl -u user:pass http://localhost:9000/api/player -X POST -H "Content-Type:application/json" --data '{"leaderboard_id":"6f5f6b7e-7067-4de9-b86a-97b6c0013c39"}'
curl -u user:pass -H "Content-Type:application/json" http://localhost:9000/api/player/score --data '{"player_id":"e9d93ada-ffca-429b-a4fd-2e68bd83b700", "amount":10}' -X PATCH
curl -u user:pass -H "Content-Type:application/json" http://localhost:9000/api/leaderboard -X DELETE --data '{"leaderboard_id":"6f5f6b7e-7067-4de9-b86a-97b6c0013c39"}'
