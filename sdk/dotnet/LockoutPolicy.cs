// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Zitadel
{
    /// <summary>
    /// Resource representing the custom lockout policy of an organization.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Zitadel = Pulumi.Zitadel;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var lockoutPolicy = new Zitadel.LockoutPolicy("lockoutPolicy", new()
    ///     {
    ///         OrgId = zitadel_org.Org.Id,
    ///         MaxPasswordAttempts = 5,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/lockoutPolicy:LockoutPolicy")]
    public partial class LockoutPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
        /// </summary>
        [Output("maxPasswordAttempts")]
        public Output<int> MaxPasswordAttempts { get; private set; } = null!;

        /// <summary>
        /// Id for the organization
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;


        /// <summary>
        /// Create a LockoutPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LockoutPolicy(string name, LockoutPolicyArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/lockoutPolicy:LockoutPolicy", name, args ?? new LockoutPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LockoutPolicy(string name, Input<string> id, LockoutPolicyState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/lockoutPolicy:LockoutPolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LockoutPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LockoutPolicy Get(string name, Input<string> id, LockoutPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new LockoutPolicy(name, id, state, options);
        }
    }

    public sealed class LockoutPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
        /// </summary>
        [Input("maxPasswordAttempts", required: true)]
        public Input<int> MaxPasswordAttempts { get; set; } = null!;

        /// <summary>
        /// Id for the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        public LockoutPolicyArgs()
        {
        }
        public static new LockoutPolicyArgs Empty => new LockoutPolicyArgs();
    }

    public sealed class LockoutPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correct or the password is reset.
        /// </summary>
        [Input("maxPasswordAttempts")]
        public Input<int>? MaxPasswordAttempts { get; set; }

        /// <summary>
        /// Id for the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        public LockoutPolicyState()
        {
        }
        public static new LockoutPolicyState Empty => new LockoutPolicyState();
    }
}