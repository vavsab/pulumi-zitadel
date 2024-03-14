// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel/internal"
)

// Datasource representing multiple API applications belonging to a project.
func GetApplicationApis(ctx *pulumi.Context, args *GetApplicationApisArgs, opts ...pulumi.InvokeOption) (*GetApplicationApisResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetApplicationApisResult
	err := ctx.Invoke("zitadel:index/getApplicationApis:getApplicationApis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplicationApis.
type GetApplicationApisArgs struct {
	// Name of the application
	Name string `pulumi:"name"`
	// Method for querying applications by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
	NameMethod *string `pulumi:"nameMethod"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// ID of the project
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getApplicationApis.
type GetApplicationApisResult struct {
	// A set of all IDs.
	AppIds []string `pulumi:"appIds"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the application
	Name string `pulumi:"name"`
	// Method for querying applications by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
	NameMethod *string `pulumi:"nameMethod"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// ID of the project
	ProjectId string `pulumi:"projectId"`
}

func GetApplicationApisOutput(ctx *pulumi.Context, args GetApplicationApisOutputArgs, opts ...pulumi.InvokeOption) GetApplicationApisResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetApplicationApisResult, error) {
			args := v.(GetApplicationApisArgs)
			r, err := GetApplicationApis(ctx, &args, opts...)
			var s GetApplicationApisResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetApplicationApisResultOutput)
}

// A collection of arguments for invoking getApplicationApis.
type GetApplicationApisOutputArgs struct {
	// Name of the application
	Name pulumi.StringInput `pulumi:"name"`
	// Method for querying applications by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
	NameMethod pulumi.StringPtrInput `pulumi:"nameMethod"`
	// ID of the organization
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
	// ID of the project
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (GetApplicationApisOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationApisArgs)(nil)).Elem()
}

// A collection of values returned by getApplicationApis.
type GetApplicationApisResultOutput struct{ *pulumi.OutputState }

func (GetApplicationApisResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationApisResult)(nil)).Elem()
}

func (o GetApplicationApisResultOutput) ToGetApplicationApisResultOutput() GetApplicationApisResultOutput {
	return o
}

func (o GetApplicationApisResultOutput) ToGetApplicationApisResultOutputWithContext(ctx context.Context) GetApplicationApisResultOutput {
	return o
}

func (o GetApplicationApisResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetApplicationApisResult] {
	return pulumix.Output[GetApplicationApisResult]{
		OutputState: o.OutputState,
	}
}

// A set of all IDs.
func (o GetApplicationApisResultOutput) AppIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetApplicationApisResult) []string { return v.AppIds }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetApplicationApisResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationApisResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the application
func (o GetApplicationApisResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationApisResult) string { return v.Name }).(pulumi.StringOutput)
}

// Method for querying applications by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
func (o GetApplicationApisResultOutput) NameMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationApisResult) *string { return v.NameMethod }).(pulumi.StringPtrOutput)
}

// ID of the organization
func (o GetApplicationApisResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApplicationApisResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// ID of the project
func (o GetApplicationApisResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationApisResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApplicationApisResultOutput{})
}