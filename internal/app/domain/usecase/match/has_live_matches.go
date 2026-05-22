package match_usecase

import "context"

func (u *MatchUseCase) HasLiveMatches(ctx context.Context) (bool, error) {
	hasLiveMatches, err := u.repo.HasLiveMatches(ctx)
	if err != nil {
		return false, err
	}

	return hasLiveMatches, nil
}
