// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing a serviceaccount situated under an organization, which then can be authorized through memberships or direct grants on other resources.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumiverse/zitadel";
 *
 * const _default = new zitadel.MachineUser("default", {
 *     orgId: data.zitadel_org["default"].id,
 *     userName: "machine@example.com",
 *     description: "a machine user",
 *     withSecret: false,
 * });
 * ```
 *
 * ## Import
 *
 * terraform The resource can be imported using the ID format `<id:has_secret[:org_id][:client_id][:client_secret]>`, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/machineUser:MachineUser imported '123456789012345678:123456789012345678:true:my-machine-user:j76mh34CHVrGGoXPQOg80lch67FIxwc2qIXjBkZoB6oMbf31eGMkB6bvRyaPjR2t'
 * ```
 */
export class MachineUser extends pulumi.CustomResource {
    /**
     * Get an existing MachineUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MachineUserState, opts?: pulumi.CustomResourceOptions): MachineUser {
        return new MachineUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/machineUser:MachineUser';

    /**
     * Returns true if the given object is an instance of MachineUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MachineUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MachineUser.__pulumiType;
    }

    /**
     * Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
     */
    public readonly accessTokenType!: pulumi.Output<string | undefined>;
    /**
     * Value of the client ID if withSecret is true
     */
    public /*out*/ readonly clientId!: pulumi.Output<string>;
    /**
     * Value of the client secret if withSecret is true
     */
    public /*out*/ readonly clientSecret!: pulumi.Output<string>;
    /**
     * Description of the user
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Loginnames
     */
    public /*out*/ readonly loginNames!: pulumi.Output<string[]>;
    /**
     * Name of the machine user
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the organization
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * Preferred login name
     */
    public /*out*/ readonly preferredLoginName!: pulumi.Output<string>;
    /**
     * State of the user
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Username
     */
    public readonly userName!: pulumi.Output<string>;
    /**
     * Generate machine secret, only applicable if creation or change from false
     */
    public readonly withSecret!: pulumi.Output<boolean | undefined>;

    /**
     * Create a MachineUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MachineUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MachineUserArgs | MachineUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MachineUserState | undefined;
            resourceInputs["accessTokenType"] = state ? state.accessTokenType : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSecret"] = state ? state.clientSecret : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["loginNames"] = state ? state.loginNames : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["preferredLoginName"] = state ? state.preferredLoginName : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
            resourceInputs["withSecret"] = state ? state.withSecret : undefined;
        } else {
            const args = argsOrState as MachineUserArgs | undefined;
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["accessTokenType"] = args ? args.accessTokenType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["withSecret"] = args ? args.withSecret : undefined;
            resourceInputs["clientId"] = undefined /*out*/;
            resourceInputs["clientSecret"] = undefined /*out*/;
            resourceInputs["loginNames"] = undefined /*out*/;
            resourceInputs["preferredLoginName"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["clientId", "clientSecret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(MachineUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MachineUser resources.
 */
export interface MachineUserState {
    /**
     * Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
     */
    accessTokenType?: pulumi.Input<string>;
    /**
     * Value of the client ID if withSecret is true
     */
    clientId?: pulumi.Input<string>;
    /**
     * Value of the client secret if withSecret is true
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * Description of the user
     */
    description?: pulumi.Input<string>;
    /**
     * Loginnames
     */
    loginNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the machine user
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * Preferred login name
     */
    preferredLoginName?: pulumi.Input<string>;
    /**
     * State of the user
     */
    state?: pulumi.Input<string>;
    /**
     * Username
     */
    userName?: pulumi.Input<string>;
    /**
     * Generate machine secret, only applicable if creation or change from false
     */
    withSecret?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a MachineUser resource.
 */
export interface MachineUserArgs {
    /**
     * Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
     */
    accessTokenType?: pulumi.Input<string>;
    /**
     * Description of the user
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the machine user
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * Username
     */
    userName: pulumi.Input<string>;
    /**
     * Generate machine secret, only applicable if creation or change from false
     */
    withSecret?: pulumi.Input<boolean>;
}
