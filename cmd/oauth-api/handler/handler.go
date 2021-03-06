package handler

import (
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/log"
	"github.com/NeuronFramework/restful"
	"github.com/NeuronOauth/oauth/api/gen/restapi/operations"
	"github.com/NeuronOauth/oauth/models"
	"github.com/NeuronOauth/oauth/services"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

type OauthHandler struct {
	logger  *zap.Logger
	service *services.OauthService
}

func NewOauthHandler() (h *OauthHandler, err error) {
	h = &OauthHandler{}
	h.logger = log.TypedLogger(h)
	h.service, err = services.NewOauthService(&services.OauthServiceOptions{})
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (h *OauthHandler) BasicAuth(clientId string, password string) (interface{}, error) {
	c, err := h.service.ClientLogin(&restful.Context{}, clientId, password)
	return c, err
}

func (h *OauthHandler) Token(p operations.TokenParams, oauthClient interface{}) middleware.Responder {
	if oauthClient == nil {
		return errors.Unauthorized("client认证失败")
	}

	if p.GrantType == "authorization_code" {
		if p.Code == nil {
			return errors.InvalidParam("Code不能为空")
		}

		if p.RedirectURI == nil {
			return errors.InvalidParam("RedirectURI不能为空")
		}

		if p.ClientID == nil {
			return errors.InvalidParam("ClientID不能为空")
		}

		result, err := h.service.AuthorizeCodeGrant(restful.NewContext(p.HTTPRequest),
			*p.Code, *p.RedirectURI, *p.ClientID, oauthClient.(*models.OauthClient))
		if err != nil {
			return errors.Wrap(err)
		}

		return operations.NewTokenOK().WithPayload(fromTokenResponse(result))
	} else if p.GrantType == "refresh_token" {
		if p.RefreshToken == nil {
			return errors.InvalidParam("RefreshToken不能为空")
		}

		if p.Scope == nil {
			return errors.InvalidParam("Scope不能为空")
		}

		result, err := h.service.RefreshTokenGrant(restful.NewContext(p.HTTPRequest),
			*p.RefreshToken, *p.Scope, oauthClient.(*models.OauthClient))
		if err != nil {
			return errors.Wrap(err)
		}

		return operations.NewTokenOK().WithPayload(fromTokenResponse(result))
	} else {
		return errors.InvalidParam("GrantType未知的类型")
	}
}

func (h *OauthHandler) Me(p operations.MeParams) middleware.Responder {
	openId, err := h.service.Me(restful.NewContext(p.HTTPRequest), p.AccessToken)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewMeOK().WithPayload(openId)
}
