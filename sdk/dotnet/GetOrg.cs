// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Zitadel
{
    public static class GetOrg
    {
        /// <summary>
        /// Datasource representing an organization in ZITADEL, which is the highest level after the instance and contains several other resource including policies if the configuration differs to the default policies on the instance.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zitadel = Pulumi.Zitadel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = Zitadel.GetOrg.Invoke(new()
        ///     {
        ///         Id = "123456789012345678",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["org"] = @default.Apply(getOrgResult =&gt; getOrgResult),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOrgResult> InvokeAsync(GetOrgArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrgResult>("zitadel:index/getOrg:getOrg", args ?? new GetOrgArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource representing an organization in ZITADEL, which is the highest level after the instance and contains several other resource including policies if the configuration differs to the default policies on the instance.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zitadel = Pulumi.Zitadel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = Zitadel.GetOrg.Invoke(new()
        ///     {
        ///         Id = "123456789012345678",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["org"] = @default.Apply(getOrgResult =&gt; getOrgResult),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetOrgResult> Invoke(GetOrgInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetOrgResult>("zitadel:index/getOrg:getOrg", args ?? new GetOrgInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrgArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetOrgArgs()
        {
        }
        public static new GetOrgArgs Empty => new GetOrgArgs();
    }

    public sealed class GetOrgInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetOrgInvokeArgs()
        {
        }
        public static new GetOrgInvokeArgs Empty => new GetOrgInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrgResult
    {
        /// <summary>
        /// ID of the organization
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the org.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Primary domain of the org
        /// </summary>
        public readonly string PrimaryDomain;
        /// <summary>
        /// State of the org, supported values: ORG*STATE*UNSPECIFIED, ORG*STATE*ACTIVE, ORG*STATE*INACTIVE, ORG*STATE*REMOVED
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetOrgResult(
            string id,

            string name,

            string primaryDomain,

            string state)
        {
            Id = id;
            Name = name;
            PrimaryDomain = primaryDomain;
            State = state;
        }
    }
}
