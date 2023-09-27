// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing an Azure AD IdP on the organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumiverse/zitadel";
 *
 * const _default = new zitadel.OrgIdpAzureAd("default", {
 *     orgId: data.zitadel_org["default"].id,
 *     clientId: "9065bfc8-a08a...",
 *     clientSecret: "H2n***",
 *     scopes: [
 *         "openid",
 *         "profile",
 *         "email",
 *         "User.Read",
 *     ],
 *     tenantType: "AZURE_AD_TENANT_TYPE_ORGANISATIONS",
 *     emailVerified: true,
 *     isLinkingAllowed: false,
 *     isCreationAllowed: true,
 *     isAutoCreation: false,
 *     isAutoUpdate: true,
 * });
 * ```
 *
 * ## Import
 *
 * terraform # The resource can be imported using the ID format `<id[:org_id][:client_secret]>`, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/orgIdpAzureAd:OrgIdpAzureAd imported '123456789012345678:123456789012345678:12345678-1234-1234-1234-123456789012'
 * ```
 */
export class OrgIdpAzureAd extends pulumi.CustomResource {
    /**
     * Get an existing OrgIdpAzureAd resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgIdpAzureAdState, opts?: pulumi.CustomResourceOptions): OrgIdpAzureAd {
        return new OrgIdpAzureAd(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/orgIdpAzureAd:OrgIdpAzureAd';

    /**
     * Returns true if the given object is an instance of OrgIdpAzureAd.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrgIdpAzureAd {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrgIdpAzureAd.__pulumiType;
    }

    /**
     * client id generated by the identity provider
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * client secret generated by the identity provider
     */
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * automatically mark emails as verified
     */
    public readonly emailVerified!: pulumi.Output<boolean>;
    /**
     * enable if a new account in ZITADEL should be created automatically on login with an external account
     */
    public readonly isAutoCreation!: pulumi.Output<boolean>;
    /**
     * enable if a the ZITADEL account fields should be updated automatically on each login
     */
    public readonly isAutoUpdate!: pulumi.Output<boolean>;
    /**
     * enable if users should be able to create a new account in ZITADEL when using an external account
     */
    public readonly isCreationAllowed!: pulumi.Output<boolean>;
    /**
     * enable if users should be able to link an existing ZITADEL user with an external account
     */
    public readonly isLinkingAllowed!: pulumi.Output<boolean>;
    /**
     * Name of the IDP
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the organization
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * the scopes requested by ZITADEL during the request on the identity provider
     */
    public readonly scopes!: pulumi.Output<string[] | undefined>;
    /**
     * if tenant*id is not set, the tenant*type is used
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;
    /**
     * the azure ad tenant type
     */
    public readonly tenantType!: pulumi.Output<string | undefined>;

    /**
     * Create a OrgIdpAzureAd resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrgIdpAzureAdArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgIdpAzureAdArgs | OrgIdpAzureAdState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgIdpAzureAdState | undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["emailVerified"] = state ? state.emailVerified : undefined;
            resourceInputs["isAutoCreation"] = state ? state.isAutoCreation : undefined;
            resourceInputs["isAutoUpdate"] = state ? state.isAutoUpdate : undefined;
            resourceInputs["isCreationAllowed"] = state ? state.isCreationAllowed : undefined;
            resourceInputs["isLinkingAllowed"] = state ? state.isLinkingAllowed : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["scopes"] = state ? state.scopes : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["tenantType"] = state ? state.tenantType : undefined;
        } else {
            const args = argsOrState as OrgIdpAzureAdArgs | undefined;
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.clientSecret === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientSecret'");
            }
            if ((!args || args.emailVerified === undefined) && !opts.urn) {
                throw new Error("Missing required property 'emailVerified'");
            }
            if ((!args || args.isAutoCreation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'isAutoCreation'");
            }
            if ((!args || args.isAutoUpdate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'isAutoUpdate'");
            }
            if ((!args || args.isCreationAllowed === undefined) && !opts.urn) {
                throw new Error("Missing required property 'isCreationAllowed'");
            }
            if ((!args || args.isLinkingAllowed === undefined) && !opts.urn) {
                throw new Error("Missing required property 'isLinkingAllowed'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args ? args.clientSecret : undefined;
            resourceInputs["emailVerified"] = args ? args.emailVerified : undefined;
            resourceInputs["isAutoCreation"] = args ? args.isAutoCreation : undefined;
            resourceInputs["isAutoUpdate"] = args ? args.isAutoUpdate : undefined;
            resourceInputs["isCreationAllowed"] = args ? args.isCreationAllowed : undefined;
            resourceInputs["isLinkingAllowed"] = args ? args.isLinkingAllowed : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["scopes"] = args ? args.scopes : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["tenantType"] = args ? args.tenantType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrgIdpAzureAd.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrgIdpAzureAd resources.
 */
export interface OrgIdpAzureAdState {
    /**
     * client id generated by the identity provider
     */
    clientId?: pulumi.Input<string>;
    /**
     * client secret generated by the identity provider
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * automatically mark emails as verified
     */
    emailVerified?: pulumi.Input<boolean>;
    /**
     * enable if a new account in ZITADEL should be created automatically on login with an external account
     */
    isAutoCreation?: pulumi.Input<boolean>;
    /**
     * enable if a the ZITADEL account fields should be updated automatically on each login
     */
    isAutoUpdate?: pulumi.Input<boolean>;
    /**
     * enable if users should be able to create a new account in ZITADEL when using an external account
     */
    isCreationAllowed?: pulumi.Input<boolean>;
    /**
     * enable if users should be able to link an existing ZITADEL user with an external account
     */
    isLinkingAllowed?: pulumi.Input<boolean>;
    /**
     * Name of the IDP
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * the scopes requested by ZITADEL during the request on the identity provider
     */
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * if tenant*id is not set, the tenant*type is used
     */
    tenantId?: pulumi.Input<string>;
    /**
     * the azure ad tenant type
     */
    tenantType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrgIdpAzureAd resource.
 */
export interface OrgIdpAzureAdArgs {
    /**
     * client id generated by the identity provider
     */
    clientId: pulumi.Input<string>;
    /**
     * client secret generated by the identity provider
     */
    clientSecret: pulumi.Input<string>;
    /**
     * automatically mark emails as verified
     */
    emailVerified: pulumi.Input<boolean>;
    /**
     * enable if a new account in ZITADEL should be created automatically on login with an external account
     */
    isAutoCreation: pulumi.Input<boolean>;
    /**
     * enable if a the ZITADEL account fields should be updated automatically on each login
     */
    isAutoUpdate: pulumi.Input<boolean>;
    /**
     * enable if users should be able to create a new account in ZITADEL when using an external account
     */
    isCreationAllowed: pulumi.Input<boolean>;
    /**
     * enable if users should be able to link an existing ZITADEL user with an external account
     */
    isLinkingAllowed: pulumi.Input<boolean>;
    /**
     * Name of the IDP
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * the scopes requested by ZITADEL during the request on the identity provider
     */
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * if tenant*id is not set, the tenant*type is used
     */
    tenantId?: pulumi.Input<string>;
    /**
     * the azure ad tenant type
     */
    tenantType?: pulumi.Input<string>;
}
