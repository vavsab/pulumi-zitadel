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
    /// <summary>
    /// Resource representing a domain of the organization.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Zitadel = Pulumiverse.Zitadel;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Zitadel.Domain("default", new()
    ///     {
    ///         OrgId = data.Zitadel_org.Default.Id,
    ///         IsPrimary = false,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform # The resource can be imported using the ID format `name[:org_id]`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import zitadel:index/domain:Domain imported 'example.com:123456789012345678'
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/domain:Domain")]
    public partial class Domain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Is domain primary
        /// </summary>
        [Output("isPrimary")]
        public Output<bool?> IsPrimary { get; private set; } = null!;

        /// <summary>
        /// Is domain verified
        /// </summary>
        [Output("isVerified")]
        public Output<bool> IsVerified { get; private set; } = null!;

        /// <summary>
        /// Name of the domain
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        /// <summary>
        /// Validation type
        /// </summary>
        [Output("validationType")]
        public Output<int> ValidationType { get; private set; } = null!;


        /// <summary>
        /// Create a Domain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Domain(string name, DomainArgs? args = null, CustomResourceOptions? options = null)
            : base("zitadel:index/domain:Domain", name, args ?? new DomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Domain(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/domain:Domain", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Domain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Domain Get(string name, Input<string> id, DomainState? state = null, CustomResourceOptions? options = null)
        {
            return new Domain(name, id, state, options);
        }
    }

    public sealed class DomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is domain primary
        /// </summary>
        [Input("isPrimary")]
        public Input<bool>? IsPrimary { get; set; }

        /// <summary>
        /// Name of the domain
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        public DomainArgs()
        {
        }
        public static new DomainArgs Empty => new DomainArgs();
    }

    public sealed class DomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is domain primary
        /// </summary>
        [Input("isPrimary")]
        public Input<bool>? IsPrimary { get; set; }

        /// <summary>
        /// Is domain verified
        /// </summary>
        [Input("isVerified")]
        public Input<bool>? IsVerified { get; set; }

        /// <summary>
        /// Name of the domain
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Validation type
        /// </summary>
        [Input("validationType")]
        public Input<int>? ValidationType { get; set; }

        public DomainState()
        {
        }
        public static new DomainState Empty => new DomainState();
    }
}
