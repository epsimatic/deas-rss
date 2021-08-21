package actions

import "net/http"

func (as *ActionSuite) Test_RoutesHandler() {
	res := as.HTML("/routes").Get()

	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Welcome to Buffalo")
}
