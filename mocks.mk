.PHONY: mocks
mocks:
	@make bll-mocks
	@make dal-mocks

# Bll
bll-mocks:
	@make leaderboard-controller-bll-mock
	@make participant-controller-bll-mock

leaderboard-controller-bll-mock:
	@mockgen -source=bll/leaderboard/interfaces.go -package mocks -mock_names="Controller=MockLeaderboardController" -destination mocks/leaderboard_controller_bll.go

participant-controller-bll-mock:
	@mockgen -source=bll/participant/controller.go -package mocks -mock_names="Controller=MockParticipantController" -destination mocks/participant_controller_bll.go

# Dal
dal-mocks:
	@make leaderboard-dal-mock
	@make participant-dal-mock

leaderboard-dal-mock:
	@mockgen -source=dal/leaderboard/dal.go -package mocks -mock_names="LeaderboardDAL=MockLeaderboardDAL" -destination mocks/leaderboard_dal.go

participant-dal-mock:
	@mockgen -source=dal/participant/dal.go -package mocks -mock_names="ParticipantDAL=MockParticipantDAL" -destination mocks/participant_dal.go
