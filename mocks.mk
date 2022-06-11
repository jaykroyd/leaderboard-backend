mocks:
	@make bll-mocks
	@make dal-mocks

# Bll
bll-mocks:
	@make leaderboard-controller-bll-mock
	@make player-controller-bll-mock

leaderboard-controller-bll-mock:
	@mockgen -source=bll/leaderboard/controller.go -package mocks -mock_names="Controller=MockLeaderboardController" -destination mocks/leaderboard_controller_bll.go

player-controller-bll-mock:
	@mockgen -source=bll/player/controller.go -package mocks -mock_names="Controller=MockPlayerController" -destination mocks/player_controller_bll.go

# Dal
dal-mocks:
	@make leaderboard-dal-mock
	@make player-dal-mock

leaderboard-dal-mock:
	@mockgen -source=dal/leaderboard/dal.go -package mocks -mock_names="LeaderboardDAL=MockLeaderboardDAL" -destination mocks/leaderboard_dal.go

player-dal-mock:
	@mockgen -source=dal/player/dal.go -package mocks -mock_names="PlayerDAL=MockPlayerDAL" -destination mocks/player_dal.go
