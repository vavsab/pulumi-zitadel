// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing the custom privacy policy of an organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const privacyPolicy = new zitadel.PrivacyPolicy("privacyPolicy", {
 *     orgId: zitadel_org.org.id,
 *     tosLink: "https://google.com",
 *     privacyLink: "https://google.com",
 *     helpLink: "https://google.com",
 * });
 * ```
 */
export class PrivacyPolicy extends pulumi.CustomResource {
    /**
     * Get an existing PrivacyPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivacyPolicyState, opts?: pulumi.CustomResourceOptions): PrivacyPolicy {
        return new PrivacyPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/privacyPolicy:PrivacyPolicy';

    /**
     * Returns true if the given object is an instance of PrivacyPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivacyPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivacyPolicy.__pulumiType;
    }

    public readonly helpLink!: pulumi.Output<string>;
    /**
     * Id for the organization
     */
    public readonly orgId!: pulumi.Output<string>;
    public readonly privacyLink!: pulumi.Output<string>;
    public readonly tosLink!: pulumi.Output<string>;

    /**
     * Create a PrivacyPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivacyPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivacyPolicyArgs | PrivacyPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrivacyPolicyState | undefined;
            resourceInputs["helpLink"] = state ? state.helpLink : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["privacyLink"] = state ? state.privacyLink : undefined;
            resourceInputs["tosLink"] = state ? state.tosLink : undefined;
        } else {
            const args = argsOrState as PrivacyPolicyArgs | undefined;
            if ((!args || args.helpLink === undefined) && !opts.urn) {
                throw new Error("Missing required property 'helpLink'");
            }
            if ((!args || args.orgId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgId'");
            }
            if ((!args || args.privacyLink === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privacyLink'");
            }
            if ((!args || args.tosLink === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tosLink'");
            }
            resourceInputs["helpLink"] = args ? args.helpLink : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["privacyLink"] = args ? args.privacyLink : undefined;
            resourceInputs["tosLink"] = args ? args.tosLink : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrivacyPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivacyPolicy resources.
 */
export interface PrivacyPolicyState {
    helpLink?: pulumi.Input<string>;
    /**
     * Id for the organization
     */
    orgId?: pulumi.Input<string>;
    privacyLink?: pulumi.Input<string>;
    tosLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PrivacyPolicy resource.
 */
export interface PrivacyPolicyArgs {
    helpLink: pulumi.Input<string>;
    /**
     * Id for the organization
     */
    orgId: pulumi.Input<string>;
    privacyLink: pulumi.Input<string>;
    tosLink: pulumi.Input<string>;
}