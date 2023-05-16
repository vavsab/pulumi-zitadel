// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing the membership of a user on an granted project, defined with the given role.
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
// 		_, err := zitadel.NewProjectGrantMember(ctx, "projectGrantMember", &zitadel.ProjectGrantMemberArgs{
// 			OrgId:     pulumi.Any(zitadel_org.Org.Id),
// 			ProjectId: pulumi.Any(zitadel_project.Project.Id),
// 			GrantId:   pulumi.Any(zitadel_project_grant.Project_grant.Id),
// 			UserId:    pulumi.Any(zitadel_human_user.Granted_human_user.Id),
// 			Roles: pulumi.StringArray{
// 				pulumi.String("PROJECT_GRANT_OWNER"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ProjectGrantMember struct {
	pulumi.CustomResourceState

	// ID of the grant
	GrantId pulumi.StringOutput `pulumi:"grantId"`
	// ID of the organization which owns the resource
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// ID of the project
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// List of roles granted
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// ID of the user
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewProjectGrantMember registers a new resource with the given unique name, arguments, and options.
func NewProjectGrantMember(ctx *pulumi.Context,
	name string, args *ProjectGrantMemberArgs, opts ...pulumi.ResourceOption) (*ProjectGrantMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GrantId == nil {
		return nil, errors.New("invalid value for required argument 'GrantId'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	var resource ProjectGrantMember
	err := ctx.RegisterResource("zitadel:index/projectGrantMember:ProjectGrantMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectGrantMember gets an existing ProjectGrantMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectGrantMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectGrantMemberState, opts ...pulumi.ResourceOption) (*ProjectGrantMember, error) {
	var resource ProjectGrantMember
	err := ctx.ReadResource("zitadel:index/projectGrantMember:ProjectGrantMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectGrantMember resources.
type projectGrantMemberState struct {
	// ID of the grant
	GrantId *string `pulumi:"grantId"`
	// ID of the organization which owns the resource
	OrgId *string `pulumi:"orgId"`
	// ID of the project
	ProjectId *string `pulumi:"projectId"`
	// List of roles granted
	Roles []string `pulumi:"roles"`
	// ID of the user
	UserId *string `pulumi:"userId"`
}

type ProjectGrantMemberState struct {
	// ID of the grant
	GrantId pulumi.StringPtrInput
	// ID of the organization which owns the resource
	OrgId pulumi.StringPtrInput
	// ID of the project
	ProjectId pulumi.StringPtrInput
	// List of roles granted
	Roles pulumi.StringArrayInput
	// ID of the user
	UserId pulumi.StringPtrInput
}

func (ProjectGrantMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectGrantMemberState)(nil)).Elem()
}

type projectGrantMemberArgs struct {
	// ID of the grant
	GrantId string `pulumi:"grantId"`
	// ID of the organization which owns the resource
	OrgId string `pulumi:"orgId"`
	// ID of the project
	ProjectId string `pulumi:"projectId"`
	// List of roles granted
	Roles []string `pulumi:"roles"`
	// ID of the user
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a ProjectGrantMember resource.
type ProjectGrantMemberArgs struct {
	// ID of the grant
	GrantId pulumi.StringInput
	// ID of the organization which owns the resource
	OrgId pulumi.StringInput
	// ID of the project
	ProjectId pulumi.StringInput
	// List of roles granted
	Roles pulumi.StringArrayInput
	// ID of the user
	UserId pulumi.StringInput
}

func (ProjectGrantMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectGrantMemberArgs)(nil)).Elem()
}

type ProjectGrantMemberInput interface {
	pulumi.Input

	ToProjectGrantMemberOutput() ProjectGrantMemberOutput
	ToProjectGrantMemberOutputWithContext(ctx context.Context) ProjectGrantMemberOutput
}

func (*ProjectGrantMember) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectGrantMember)(nil)).Elem()
}

func (i *ProjectGrantMember) ToProjectGrantMemberOutput() ProjectGrantMemberOutput {
	return i.ToProjectGrantMemberOutputWithContext(context.Background())
}

func (i *ProjectGrantMember) ToProjectGrantMemberOutputWithContext(ctx context.Context) ProjectGrantMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectGrantMemberOutput)
}

// ProjectGrantMemberArrayInput is an input type that accepts ProjectGrantMemberArray and ProjectGrantMemberArrayOutput values.
// You can construct a concrete instance of `ProjectGrantMemberArrayInput` via:
//
//          ProjectGrantMemberArray{ ProjectGrantMemberArgs{...} }
type ProjectGrantMemberArrayInput interface {
	pulumi.Input

	ToProjectGrantMemberArrayOutput() ProjectGrantMemberArrayOutput
	ToProjectGrantMemberArrayOutputWithContext(context.Context) ProjectGrantMemberArrayOutput
}

type ProjectGrantMemberArray []ProjectGrantMemberInput

func (ProjectGrantMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectGrantMember)(nil)).Elem()
}

func (i ProjectGrantMemberArray) ToProjectGrantMemberArrayOutput() ProjectGrantMemberArrayOutput {
	return i.ToProjectGrantMemberArrayOutputWithContext(context.Background())
}

func (i ProjectGrantMemberArray) ToProjectGrantMemberArrayOutputWithContext(ctx context.Context) ProjectGrantMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectGrantMemberArrayOutput)
}

// ProjectGrantMemberMapInput is an input type that accepts ProjectGrantMemberMap and ProjectGrantMemberMapOutput values.
// You can construct a concrete instance of `ProjectGrantMemberMapInput` via:
//
//          ProjectGrantMemberMap{ "key": ProjectGrantMemberArgs{...} }
type ProjectGrantMemberMapInput interface {
	pulumi.Input

	ToProjectGrantMemberMapOutput() ProjectGrantMemberMapOutput
	ToProjectGrantMemberMapOutputWithContext(context.Context) ProjectGrantMemberMapOutput
}

type ProjectGrantMemberMap map[string]ProjectGrantMemberInput

func (ProjectGrantMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectGrantMember)(nil)).Elem()
}

func (i ProjectGrantMemberMap) ToProjectGrantMemberMapOutput() ProjectGrantMemberMapOutput {
	return i.ToProjectGrantMemberMapOutputWithContext(context.Background())
}

func (i ProjectGrantMemberMap) ToProjectGrantMemberMapOutputWithContext(ctx context.Context) ProjectGrantMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectGrantMemberMapOutput)
}

type ProjectGrantMemberOutput struct{ *pulumi.OutputState }

func (ProjectGrantMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectGrantMember)(nil)).Elem()
}

func (o ProjectGrantMemberOutput) ToProjectGrantMemberOutput() ProjectGrantMemberOutput {
	return o
}

func (o ProjectGrantMemberOutput) ToProjectGrantMemberOutputWithContext(ctx context.Context) ProjectGrantMemberOutput {
	return o
}

// ID of the grant
func (o ProjectGrantMemberOutput) GrantId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectGrantMember) pulumi.StringOutput { return v.GrantId }).(pulumi.StringOutput)
}

// ID of the organization which owns the resource
func (o ProjectGrantMemberOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectGrantMember) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// ID of the project
func (o ProjectGrantMemberOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectGrantMember) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// List of roles granted
func (o ProjectGrantMemberOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProjectGrantMember) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// ID of the user
func (o ProjectGrantMemberOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectGrantMember) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type ProjectGrantMemberArrayOutput struct{ *pulumi.OutputState }

func (ProjectGrantMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectGrantMember)(nil)).Elem()
}

func (o ProjectGrantMemberArrayOutput) ToProjectGrantMemberArrayOutput() ProjectGrantMemberArrayOutput {
	return o
}

func (o ProjectGrantMemberArrayOutput) ToProjectGrantMemberArrayOutputWithContext(ctx context.Context) ProjectGrantMemberArrayOutput {
	return o
}

func (o ProjectGrantMemberArrayOutput) Index(i pulumi.IntInput) ProjectGrantMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectGrantMember {
		return vs[0].([]*ProjectGrantMember)[vs[1].(int)]
	}).(ProjectGrantMemberOutput)
}

type ProjectGrantMemberMapOutput struct{ *pulumi.OutputState }

func (ProjectGrantMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectGrantMember)(nil)).Elem()
}

func (o ProjectGrantMemberMapOutput) ToProjectGrantMemberMapOutput() ProjectGrantMemberMapOutput {
	return o
}

func (o ProjectGrantMemberMapOutput) ToProjectGrantMemberMapOutputWithContext(ctx context.Context) ProjectGrantMemberMapOutput {
	return o
}

func (o ProjectGrantMemberMapOutput) MapIndex(k pulumi.StringInput) ProjectGrantMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectGrantMember {
		return vs[0].(map[string]*ProjectGrantMember)[vs[1].(string)]
	}).(ProjectGrantMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectGrantMemberInput)(nil)).Elem(), &ProjectGrantMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectGrantMemberArrayInput)(nil)).Elem(), ProjectGrantMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectGrantMemberMapInput)(nil)).Elem(), ProjectGrantMemberMap{})
	pulumi.RegisterOutputType(ProjectGrantMemberOutput{})
	pulumi.RegisterOutputType(ProjectGrantMemberArrayOutput{})
	pulumi.RegisterOutputType(ProjectGrantMemberMapOutput{})
}