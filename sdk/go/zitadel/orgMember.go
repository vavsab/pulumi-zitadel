// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing the membership of a user on an organization, defined with the given role.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-zitadel/sdk/go/zitadel"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := zitadel.NewOrgMember(ctx, "orgMember", &zitadel.OrgMemberArgs{
// 			OrgId:  pulumi.Any(zitadel_org.Org.Id),
// 			UserId: pulumi.Any(zitadel_human_user.Human_user.Id),
// 			Roles: pulumi.StringArray{
// 				pulumi.String("ORG_OWNER"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type OrgMember struct {
	pulumi.CustomResourceState

	// ID of the organization
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// List of roles granted
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// ID of the user
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewOrgMember registers a new resource with the given unique name, arguments, and options.
func NewOrgMember(ctx *pulumi.Context,
	name string, args *OrgMemberArgs, opts ...pulumi.ResourceOption) (*OrgMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	var resource OrgMember
	err := ctx.RegisterResource("zitadel:index/orgMember:OrgMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrgMember gets an existing OrgMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrgMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrgMemberState, opts ...pulumi.ResourceOption) (*OrgMember, error) {
	var resource OrgMember
	err := ctx.ReadResource("zitadel:index/orgMember:OrgMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrgMember resources.
type orgMemberState struct {
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// List of roles granted
	Roles []string `pulumi:"roles"`
	// ID of the user
	UserId *string `pulumi:"userId"`
}

type OrgMemberState struct {
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// List of roles granted
	Roles pulumi.StringArrayInput
	// ID of the user
	UserId pulumi.StringPtrInput
}

func (OrgMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*orgMemberState)(nil)).Elem()
}

type orgMemberArgs struct {
	// ID of the organization
	OrgId string `pulumi:"orgId"`
	// List of roles granted
	Roles []string `pulumi:"roles"`
	// ID of the user
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a OrgMember resource.
type OrgMemberArgs struct {
	// ID of the organization
	OrgId pulumi.StringInput
	// List of roles granted
	Roles pulumi.StringArrayInput
	// ID of the user
	UserId pulumi.StringInput
}

func (OrgMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orgMemberArgs)(nil)).Elem()
}

type OrgMemberInput interface {
	pulumi.Input

	ToOrgMemberOutput() OrgMemberOutput
	ToOrgMemberOutputWithContext(ctx context.Context) OrgMemberOutput
}

func (*OrgMember) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgMember)(nil)).Elem()
}

func (i *OrgMember) ToOrgMemberOutput() OrgMemberOutput {
	return i.ToOrgMemberOutputWithContext(context.Background())
}

func (i *OrgMember) ToOrgMemberOutputWithContext(ctx context.Context) OrgMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgMemberOutput)
}

// OrgMemberArrayInput is an input type that accepts OrgMemberArray and OrgMemberArrayOutput values.
// You can construct a concrete instance of `OrgMemberArrayInput` via:
//
//          OrgMemberArray{ OrgMemberArgs{...} }
type OrgMemberArrayInput interface {
	pulumi.Input

	ToOrgMemberArrayOutput() OrgMemberArrayOutput
	ToOrgMemberArrayOutputWithContext(context.Context) OrgMemberArrayOutput
}

type OrgMemberArray []OrgMemberInput

func (OrgMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgMember)(nil)).Elem()
}

func (i OrgMemberArray) ToOrgMemberArrayOutput() OrgMemberArrayOutput {
	return i.ToOrgMemberArrayOutputWithContext(context.Background())
}

func (i OrgMemberArray) ToOrgMemberArrayOutputWithContext(ctx context.Context) OrgMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgMemberArrayOutput)
}

// OrgMemberMapInput is an input type that accepts OrgMemberMap and OrgMemberMapOutput values.
// You can construct a concrete instance of `OrgMemberMapInput` via:
//
//          OrgMemberMap{ "key": OrgMemberArgs{...} }
type OrgMemberMapInput interface {
	pulumi.Input

	ToOrgMemberMapOutput() OrgMemberMapOutput
	ToOrgMemberMapOutputWithContext(context.Context) OrgMemberMapOutput
}

type OrgMemberMap map[string]OrgMemberInput

func (OrgMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgMember)(nil)).Elem()
}

func (i OrgMemberMap) ToOrgMemberMapOutput() OrgMemberMapOutput {
	return i.ToOrgMemberMapOutputWithContext(context.Background())
}

func (i OrgMemberMap) ToOrgMemberMapOutputWithContext(ctx context.Context) OrgMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgMemberMapOutput)
}

type OrgMemberOutput struct{ *pulumi.OutputState }

func (OrgMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgMember)(nil)).Elem()
}

func (o OrgMemberOutput) ToOrgMemberOutput() OrgMemberOutput {
	return o
}

func (o OrgMemberOutput) ToOrgMemberOutputWithContext(ctx context.Context) OrgMemberOutput {
	return o
}

// ID of the organization
func (o OrgMemberOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgMember) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// List of roles granted
func (o OrgMemberOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrgMember) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// ID of the user
func (o OrgMemberOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgMember) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type OrgMemberArrayOutput struct{ *pulumi.OutputState }

func (OrgMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgMember)(nil)).Elem()
}

func (o OrgMemberArrayOutput) ToOrgMemberArrayOutput() OrgMemberArrayOutput {
	return o
}

func (o OrgMemberArrayOutput) ToOrgMemberArrayOutputWithContext(ctx context.Context) OrgMemberArrayOutput {
	return o
}

func (o OrgMemberArrayOutput) Index(i pulumi.IntInput) OrgMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrgMember {
		return vs[0].([]*OrgMember)[vs[1].(int)]
	}).(OrgMemberOutput)
}

type OrgMemberMapOutput struct{ *pulumi.OutputState }

func (OrgMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgMember)(nil)).Elem()
}

func (o OrgMemberMapOutput) ToOrgMemberMapOutput() OrgMemberMapOutput {
	return o
}

func (o OrgMemberMapOutput) ToOrgMemberMapOutputWithContext(ctx context.Context) OrgMemberMapOutput {
	return o
}

func (o OrgMemberMapOutput) MapIndex(k pulumi.StringInput) OrgMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrgMember {
		return vs[0].(map[string]*OrgMember)[vs[1].(string)]
	}).(OrgMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrgMemberInput)(nil)).Elem(), &OrgMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgMemberArrayInput)(nil)).Elem(), OrgMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgMemberMapInput)(nil)).Elem(), OrgMemberMap{})
	pulumi.RegisterOutputType(OrgMemberOutput{})
	pulumi.RegisterOutputType(OrgMemberArrayOutput{})
	pulumi.RegisterOutputType(OrgMemberMapOutput{})
}