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
    /// Resource representing the default notification policy.
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
    ///     var notificationPolicy = new Zitadel.DefaultNotificationPolicy("notificationPolicy", new()
    ///     {
    ///         PasswordChange = false,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/defaultNotificationPolicy:DefaultNotificationPolicy")]
    public partial class DefaultNotificationPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Send notification if a user changes his password
        /// </summary>
        [Output("passwordChange")]
        public Output<bool> PasswordChange { get; private set; } = null!;


        /// <summary>
        /// Create a DefaultNotificationPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DefaultNotificationPolicy(string name, DefaultNotificationPolicyArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/defaultNotificationPolicy:DefaultNotificationPolicy", name, args ?? new DefaultNotificationPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DefaultNotificationPolicy(string name, Input<string> id, DefaultNotificationPolicyState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/defaultNotificationPolicy:DefaultNotificationPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DefaultNotificationPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DefaultNotificationPolicy Get(string name, Input<string> id, DefaultNotificationPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new DefaultNotificationPolicy(name, id, state, options);
        }
    }

    public sealed class DefaultNotificationPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Send notification if a user changes his password
        /// </summary>
        [Input("passwordChange", required: true)]
        public Input<bool> PasswordChange { get; set; } = null!;

        public DefaultNotificationPolicyArgs()
        {
        }
        public static new DefaultNotificationPolicyArgs Empty => new DefaultNotificationPolicyArgs();
    }

    public sealed class DefaultNotificationPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Send notification if a user changes his password
        /// </summary>
        [Input("passwordChange")]
        public Input<bool>? PasswordChange { get; set; }

        public DefaultNotificationPolicyState()
        {
        }
        public static new DefaultNotificationPolicyState Empty => new DefaultNotificationPolicyState();
    }
}