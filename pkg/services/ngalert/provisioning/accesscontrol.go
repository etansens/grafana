package provisioning

import (
	"context"

	ac "github.com/grafana/grafana/pkg/services/accesscontrol"
	"github.com/grafana/grafana/pkg/services/auth/identity"
	"github.com/grafana/grafana/pkg/services/ngalert/accesscontrol"
	"github.com/grafana/grafana/pkg/services/ngalert/models"
	"github.com/grafana/grafana/pkg/services/ngalert/store"
)

type ruleService interface {
	HasAccess(ctx context.Context, user identity.Requester, evaluator ac.Evaluator) (bool, error)
	AuthorizeAccessToRuleGroup(ctx context.Context, user identity.Requester, rules models.RulesGroup) error
	AuthorizeRuleChanges(ctx context.Context, user identity.Requester, change *store.GroupDelta) error
}

func newRuleAccessControlService(ac *accesscontrol.RuleService) *provisioningRuleAccessControl {
	return &provisioningRuleAccessControl{
		ruleService: ac,
	}
}

type provisioningRuleAccessControl struct {
	ruleService
}

var _ ruleAccessControlService = &provisioningRuleAccessControl{}

func (p *provisioningRuleAccessControl) AuthorizeAccessToRuleGroup(ctx context.Context, user identity.Requester, rules models.RulesGroup) error {
	if can, err := p.CanReadAllRules(ctx, user); can || err != nil {
		return err
	}
	return p.ruleService.AuthorizeAccessToRuleGroup(ctx, user, rules)
}

func (p *provisioningRuleAccessControl) AuthorizeRuleChanges(ctx context.Context, user identity.Requester, change *store.GroupDelta) error {
	if can, err := p.CanWriteAllRules(ctx, user); can || err != nil {
		return err
	}
	return p.ruleService.AuthorizeRuleChanges(ctx, user, change)
}

func (p *provisioningRuleAccessControl) CanReadAllRules(ctx context.Context, user identity.Requester) (bool, error) {
	return p.HasAccess(ctx, user, ac.EvalAny(
		ac.EvalPermission(ac.ActionAlertingProvisioningRead),
		ac.EvalPermission(ac.ActionAlertingProvisioningReadSecrets),
	))
}

func (p *provisioningRuleAccessControl) CanWriteAllRules(ctx context.Context, user identity.Requester) (bool, error) {
	return p.HasAccess(ctx, user, ac.EvalPermission(ac.ActionAlertingProvisioningWrite))
}

// access control that simulate full access to read\write rules
type allAccessControlService struct {
}

var _ ruleAccessControlService = &allAccessControlService{}

func (a allAccessControlService) AuthorizeAccessToRuleGroup(ctx context.Context, user identity.Requester, rules models.RulesGroup) error {
	return nil
}

func (a allAccessControlService) AuthorizeRuleChanges(ctx context.Context, user identity.Requester, change *store.GroupDelta) error {
	return nil
}

func (a allAccessControlService) CanReadAllRules(ctx context.Context, user identity.Requester) (bool, error) {
	return true, nil
}

func (a allAccessControlService) CanWriteAllRules(ctx context.Context, user identity.Requester) (bool, error) {
	return true, nil
}
