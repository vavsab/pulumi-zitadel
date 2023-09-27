// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing a GitLab IdP on the organization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := zitadel.NewOrgIdpGitlab(ctx, "default", &zitadel.OrgIdpGitlabArgs{
// 			OrgId:        pulumi.Any(data.Zitadel_org.Default.Id),
// 			ClientId:     pulumi.String("15765e..."),
// 			ClientSecret: pulumi.String("*****abcxyz"),
// 			Scopes: pulumi.StringArray{
// 				pulumi.String("openid"),
// 				pulumi.String("profile"),
// 				pulumi.String("email"),
// 			},
// 			IsLinkingAllowed:  pulumi.Bool(false),
// 			IsCreationAllowed: pulumi.Bool(true),
// 			IsAutoCreation:    pulumi.Bool(false),
// 			IsAutoUpdate:      pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// terraform # The resource can be imported using the ID format `<id[:org_id][:client_secret]>`, e.g.
//
// ```sh
//  $ pulumi import zitadel:index/orgIdpGitlab:OrgIdpGitlab imported '123456789012345678:123456789012345678:1234567890abcdef'
// ```
type OrgIdpGitlab struct {
	pulumi.CustomResourceState

	// client id generated by the identity provider
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolOutput `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolOutput `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolOutput `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolOutput `pulumi:"isLinkingAllowed"`
	// Name of the IDP
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the organization
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
}

// NewOrgIdpGitlab registers a new resource with the given unique name, arguments, and options.
func NewOrgIdpGitlab(ctx *pulumi.Context,
	name string, args *OrgIdpGitlabArgs, opts ...pulumi.ResourceOption) (*OrgIdpGitlab, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
	}
	if args.IsAutoCreation == nil {
		return nil, errors.New("invalid value for required argument 'IsAutoCreation'")
	}
	if args.IsAutoUpdate == nil {
		return nil, errors.New("invalid value for required argument 'IsAutoUpdate'")
	}
	if args.IsCreationAllowed == nil {
		return nil, errors.New("invalid value for required argument 'IsCreationAllowed'")
	}
	if args.IsLinkingAllowed == nil {
		return nil, errors.New("invalid value for required argument 'IsLinkingAllowed'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource OrgIdpGitlab
	err := ctx.RegisterResource("zitadel:index/orgIdpGitlab:OrgIdpGitlab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrgIdpGitlab gets an existing OrgIdpGitlab resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrgIdpGitlab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrgIdpGitlabState, opts ...pulumi.ResourceOption) (*OrgIdpGitlab, error) {
	var resource OrgIdpGitlab
	err := ctx.ReadResource("zitadel:index/orgIdpGitlab:OrgIdpGitlab", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrgIdpGitlab resources.
type orgIdpGitlabState struct {
	// client id generated by the identity provider
	ClientId *string `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret *string `pulumi:"clientSecret"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation *bool `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate *bool `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed *bool `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed *bool `pulumi:"isLinkingAllowed"`
	// Name of the IDP
	Name *string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
}

type OrgIdpGitlabState struct {
	// client id generated by the identity provider
	ClientId pulumi.StringPtrInput
	// client secret generated by the identity provider
	ClientSecret pulumi.StringPtrInput
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolPtrInput
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolPtrInput
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolPtrInput
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolPtrInput
	// Name of the IDP
	Name pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayInput
}

func (OrgIdpGitlabState) ElementType() reflect.Type {
	return reflect.TypeOf((*orgIdpGitlabState)(nil)).Elem()
}

type orgIdpGitlabArgs struct {
	// client id generated by the identity provider
	ClientId string `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret string `pulumi:"clientSecret"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation bool `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate bool `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed bool `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed bool `pulumi:"isLinkingAllowed"`
	// Name of the IDP
	Name *string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
}

// The set of arguments for constructing a OrgIdpGitlab resource.
type OrgIdpGitlabArgs struct {
	// client id generated by the identity provider
	ClientId pulumi.StringInput
	// client secret generated by the identity provider
	ClientSecret pulumi.StringInput
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolInput
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolInput
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolInput
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolInput
	// Name of the IDP
	Name pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayInput
}

func (OrgIdpGitlabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orgIdpGitlabArgs)(nil)).Elem()
}

type OrgIdpGitlabInput interface {
	pulumi.Input

	ToOrgIdpGitlabOutput() OrgIdpGitlabOutput
	ToOrgIdpGitlabOutputWithContext(ctx context.Context) OrgIdpGitlabOutput
}

func (*OrgIdpGitlab) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgIdpGitlab)(nil)).Elem()
}

func (i *OrgIdpGitlab) ToOrgIdpGitlabOutput() OrgIdpGitlabOutput {
	return i.ToOrgIdpGitlabOutputWithContext(context.Background())
}

func (i *OrgIdpGitlab) ToOrgIdpGitlabOutputWithContext(ctx context.Context) OrgIdpGitlabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpGitlabOutput)
}

// OrgIdpGitlabArrayInput is an input type that accepts OrgIdpGitlabArray and OrgIdpGitlabArrayOutput values.
// You can construct a concrete instance of `OrgIdpGitlabArrayInput` via:
//
//          OrgIdpGitlabArray{ OrgIdpGitlabArgs{...} }
type OrgIdpGitlabArrayInput interface {
	pulumi.Input

	ToOrgIdpGitlabArrayOutput() OrgIdpGitlabArrayOutput
	ToOrgIdpGitlabArrayOutputWithContext(context.Context) OrgIdpGitlabArrayOutput
}

type OrgIdpGitlabArray []OrgIdpGitlabInput

func (OrgIdpGitlabArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgIdpGitlab)(nil)).Elem()
}

func (i OrgIdpGitlabArray) ToOrgIdpGitlabArrayOutput() OrgIdpGitlabArrayOutput {
	return i.ToOrgIdpGitlabArrayOutputWithContext(context.Background())
}

func (i OrgIdpGitlabArray) ToOrgIdpGitlabArrayOutputWithContext(ctx context.Context) OrgIdpGitlabArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpGitlabArrayOutput)
}

// OrgIdpGitlabMapInput is an input type that accepts OrgIdpGitlabMap and OrgIdpGitlabMapOutput values.
// You can construct a concrete instance of `OrgIdpGitlabMapInput` via:
//
//          OrgIdpGitlabMap{ "key": OrgIdpGitlabArgs{...} }
type OrgIdpGitlabMapInput interface {
	pulumi.Input

	ToOrgIdpGitlabMapOutput() OrgIdpGitlabMapOutput
	ToOrgIdpGitlabMapOutputWithContext(context.Context) OrgIdpGitlabMapOutput
}

type OrgIdpGitlabMap map[string]OrgIdpGitlabInput

func (OrgIdpGitlabMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgIdpGitlab)(nil)).Elem()
}

func (i OrgIdpGitlabMap) ToOrgIdpGitlabMapOutput() OrgIdpGitlabMapOutput {
	return i.ToOrgIdpGitlabMapOutputWithContext(context.Background())
}

func (i OrgIdpGitlabMap) ToOrgIdpGitlabMapOutputWithContext(ctx context.Context) OrgIdpGitlabMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpGitlabMapOutput)
}

type OrgIdpGitlabOutput struct{ *pulumi.OutputState }

func (OrgIdpGitlabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgIdpGitlab)(nil)).Elem()
}

func (o OrgIdpGitlabOutput) ToOrgIdpGitlabOutput() OrgIdpGitlabOutput {
	return o
}

func (o OrgIdpGitlabOutput) ToOrgIdpGitlabOutputWithContext(ctx context.Context) OrgIdpGitlabOutput {
	return o
}

// client id generated by the identity provider
func (o OrgIdpGitlabOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpGitlab) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o OrgIdpGitlabOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpGitlab) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// enable if a new account in ZITADEL should be created automatically on login with an external account
func (o OrgIdpGitlabOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpGitlab) pulumi.BoolOutput { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enable if a the ZITADEL account fields should be updated automatically on each login
func (o OrgIdpGitlabOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpGitlab) pulumi.BoolOutput { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enable if users should be able to create a new account in ZITADEL when using an external account
func (o OrgIdpGitlabOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpGitlab) pulumi.BoolOutput { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enable if users should be able to link an existing ZITADEL user with an external account
func (o OrgIdpGitlabOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpGitlab) pulumi.BoolOutput { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// Name of the IDP
func (o OrgIdpGitlabOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpGitlab) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o OrgIdpGitlabOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpGitlab) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o OrgIdpGitlabOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrgIdpGitlab) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

type OrgIdpGitlabArrayOutput struct{ *pulumi.OutputState }

func (OrgIdpGitlabArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgIdpGitlab)(nil)).Elem()
}

func (o OrgIdpGitlabArrayOutput) ToOrgIdpGitlabArrayOutput() OrgIdpGitlabArrayOutput {
	return o
}

func (o OrgIdpGitlabArrayOutput) ToOrgIdpGitlabArrayOutputWithContext(ctx context.Context) OrgIdpGitlabArrayOutput {
	return o
}

func (o OrgIdpGitlabArrayOutput) Index(i pulumi.IntInput) OrgIdpGitlabOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrgIdpGitlab {
		return vs[0].([]*OrgIdpGitlab)[vs[1].(int)]
	}).(OrgIdpGitlabOutput)
}

type OrgIdpGitlabMapOutput struct{ *pulumi.OutputState }

func (OrgIdpGitlabMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgIdpGitlab)(nil)).Elem()
}

func (o OrgIdpGitlabMapOutput) ToOrgIdpGitlabMapOutput() OrgIdpGitlabMapOutput {
	return o
}

func (o OrgIdpGitlabMapOutput) ToOrgIdpGitlabMapOutputWithContext(ctx context.Context) OrgIdpGitlabMapOutput {
	return o
}

func (o OrgIdpGitlabMapOutput) MapIndex(k pulumi.StringInput) OrgIdpGitlabOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrgIdpGitlab {
		return vs[0].(map[string]*OrgIdpGitlab)[vs[1].(string)]
	}).(OrgIdpGitlabOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpGitlabInput)(nil)).Elem(), &OrgIdpGitlab{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpGitlabArrayInput)(nil)).Elem(), OrgIdpGitlabArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpGitlabMapInput)(nil)).Elem(), OrgIdpGitlabMap{})
	pulumi.RegisterOutputType(OrgIdpGitlabOutput{})
	pulumi.RegisterOutputType(OrgIdpGitlabArrayOutput{})
	pulumi.RegisterOutputType(OrgIdpGitlabMapOutput{})
}
