// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing a GitHub Enterprise IdP on the organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumiverse/zitadel";
 *
 * const _default = new zitadel.OrgIdpGithubEs("default", {
 *     orgId: data.zitadel_org["default"].id,
 *     clientId: "86a165...",
 *     clientSecret: "*****afdbac18",
 *     scopes: [
 *         "openid",
 *         "profile",
 *         "email",
 *     ],
 *     authorizationEndpoint: "https://auth.endpoint",
 *     tokenEndpoint: "https://token.endpoint",
 *     userEndpoint: "https://user.endpoint",
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
 *  $ pulumi import zitadel:index/orgIdpGithubEs:OrgIdpGithubEs imported '123456789012345678:123456789012345678:123456789012345678:123456789012345678'
 * ```
 */
export class OrgIdpGithubEs extends pulumi.CustomResource {
    /**
     * Get an existing OrgIdpGithubEs resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgIdpGithubEsState, opts?: pulumi.CustomResourceOptions): OrgIdpGithubEs {
        return new OrgIdpGithubEs(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/orgIdpGithubEs:OrgIdpGithubEs';

    /**
     * Returns true if the given object is an instance of OrgIdpGithubEs.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrgIdpGithubEs {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrgIdpGithubEs.__pulumiType;
    }

    /**
     * the providers authorization endpoint
     */
    public readonly authorizationEndpoint!: pulumi.Output<string>;
    /**
     * client id generated by the identity provider
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * client secret generated by the identity provider
     */
    public readonly clientSecret!: pulumi.Output<string>;
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
     * the providers token endpoint
     */
    public readonly tokenEndpoint!: pulumi.Output<string>;
    /**
     * the providers user endpoint
     */
    public readonly userEndpoint!: pulumi.Output<string>;

    /**
     * Create a OrgIdpGithubEs resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrgIdpGithubEsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgIdpGithubEsArgs | OrgIdpGithubEsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgIdpGithubEsState | undefined;
            resourceInputs["authorizationEndpoint"] = state ? state.authorizationEndpoint : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["isAutoCreation"] = state ? state.isAutoCreation : undefined;
            resourceInputs["isAutoUpdate"] = state ? state.isAutoUpdate : undefined;
            resourceInputs["isCreationAllowed"] = state ? state.isCreationAllowed : undefined;
            resourceInputs["isLinkingAllowed"] = state ? state.isLinkingAllowed : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["scopes"] = state ? state.scopes : undefined;
            resourceInputs["tokenEndpoint"] = state ? state.tokenEndpoint : undefined;
            resourceInputs["userEndpoint"] = state ? state.userEndpoint : undefined;
        } else {
            const args = argsOrState as OrgIdpGithubEsArgs | undefined;
            if ((!args || args.authorizationEndpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorizationEndpoint'");
            }
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.clientSecret === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientSecret'");
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
            if ((!args || args.tokenEndpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tokenEndpoint'");
            }
            if ((!args || args.userEndpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userEndpoint'");
            }
            resourceInputs["authorizationEndpoint"] = args ? args.authorizationEndpoint : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args ? args.clientSecret : undefined;
            resourceInputs["isAutoCreation"] = args ? args.isAutoCreation : undefined;
            resourceInputs["isAutoUpdate"] = args ? args.isAutoUpdate : undefined;
            resourceInputs["isCreationAllowed"] = args ? args.isCreationAllowed : undefined;
            resourceInputs["isLinkingAllowed"] = args ? args.isLinkingAllowed : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["scopes"] = args ? args.scopes : undefined;
            resourceInputs["tokenEndpoint"] = args ? args.tokenEndpoint : undefined;
            resourceInputs["userEndpoint"] = args ? args.userEndpoint : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrgIdpGithubEs.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrgIdpGithubEs resources.
 */
export interface OrgIdpGithubEsState {
    /**
     * the providers authorization endpoint
     */
    authorizationEndpoint?: pulumi.Input<string>;
    /**
     * client id generated by the identity provider
     */
    clientId?: pulumi.Input<string>;
    /**
     * client secret generated by the identity provider
     */
    clientSecret?: pulumi.Input<string>;
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
     * the providers token endpoint
     */
    tokenEndpoint?: pulumi.Input<string>;
    /**
     * the providers user endpoint
     */
    userEndpoint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrgIdpGithubEs resource.
 */
export interface OrgIdpGithubEsArgs {
    /**
     * the providers authorization endpoint
     */
    authorizationEndpoint: pulumi.Input<string>;
    /**
     * client id generated by the identity provider
     */
    clientId: pulumi.Input<string>;
    /**
     * client secret generated by the identity provider
     */
    clientSecret: pulumi.Input<string>;
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
     * the providers token endpoint
     */
    tokenEndpoint: pulumi.Input<string>;
    /**
     * the providers user endpoint
     */
    userEndpoint: pulumi.Input<string>;
}
