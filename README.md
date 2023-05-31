This is a backend application for managing leaderboards for games.

## Getting Started

First, run the development server (make sure you have [make](https://www.gnu.org/software/make/) installed):

```bash
make dev-run-api
```

You can make requests to [http://localhost:8080/api/v1](http://localhost:8080/api/v1) in order to create/update leaderboards.

## Routes
* POST /leaderboards

Creates new leaderboard

* POST /participants

Creates new leaderboard participant

* GET /leaderboards/{leaderboard_id}

Gets leaderboard information

* GET /leaderboards/{leaderboard_id}/participants/{external_id}

Gets leaderboard participant information

* GET /leaderboards

Lists leaderboards

* GET /leaderboards/{leaderboard_id}/participants

Lists participants

* DELETE /leaderboard/{leaderboard_id}

Deletes leaderboard

* DELETE /leaderboards/{leaderboard_id}/participants/{external_id}

Deletes leaderboard participant

* PATCH /leaderboard/{leaderboard_id}/reset

Resets leaderboard

* PATCH /leaderboards/{leaderboard_id}/participants/{external_id}/score

Updates leaderboard participant score
