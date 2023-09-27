// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Datasource representing an LDAP IdP on the organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const default = zitadel.getOrgIdpLdap({
 *     orgId: data.zitadel_org["default"].id,
 *     id: "123456789012345678",
 * });
 * ```
 */
export function getOrgIdpLdap(args: GetOrgIdpLdapArgs, opts?: pulumi.InvokeOptions): Promise<GetOrgIdpLdapResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("zitadel:index/getOrgIdpLdap:getOrgIdpLdap", {
        "id": args.id,
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrgIdpLdap.
 */
export interface GetOrgIdpLdapArgs {
    /**
     * The ID of this resource.
     */
    id: string;
    /**
     * ID of the organization
     */
    orgId?: string;
}

/**
 * A collection of values returned by getOrgIdpLdap.
 */
export interface GetOrgIdpLdapResult {
    /**
     * User attribute for the avatar url
     */
    readonly avatarUrlAttribute: string;
    /**
     * Base DN for LDAP connections
     */
    readonly baseDn: string;
    /**
     * Bind DN for LDAP connections
     */
    readonly bindDn: string;
    /**
     * Bind password for LDAP connections
     */
    readonly bindPassword: string;
    /**
     * User attribute for the display name
     */
    readonly displayNameAttribute: string;
    /**
     * User attribute for the email
     */
    readonly emailAttribute: string;
    /**
     * User attribute for the email verified state
     */
    readonly emailVerifiedAttribute: string;
    /**
     * User attribute for the first name
     */
    readonly firstNameAttribute: string;
    /**
     * The ID of this resource.
     */
    readonly id: string;
    /**
     * User attribute for the id
     */
    readonly idAttribute: string;
    /**
     * enabled if a new account in ZITADEL are created automatically on login with an external account
     */
    readonly isAutoCreation: boolean;
    /**
     * enabled if a the ZITADEL account fields are updated automatically on each login
     */
    readonly isAutoUpdate: boolean;
    /**
     * enabled if users are able to create a new account in ZITADEL when using an external account
     */
    readonly isCreationAllowed: boolean;
    /**
     * enabled if users are able to link an existing ZITADEL user with an external account
     */
    readonly isLinkingAllowed: boolean;
    /**
     * User attribute for the last name
     */
    readonly lastNameAttribute: string;
    /**
     * Name of the IDP
     */
    readonly name: string;
    /**
     * User attribute for the nick name
     */
    readonly nickNameAttribute: string;
    /**
     * ID of the organization
     */
    readonly orgId?: string;
    /**
     * User attribute for the phone
     */
    readonly phoneAttribute: string;
    /**
     * User attribute for the phone verified state
     */
    readonly phoneVerifiedAttribute: string;
    /**
     * User attribute for the preferred language
     */
    readonly preferredLanguageAttribute: string;
    /**
     * User attribute for the preferred username
     */
    readonly preferredUsernameAttribute: string;
    /**
     * User attribute for the profile
     */
    readonly profileAttribute: string;
    /**
     * Servers to try in order for establishing LDAP connections
     */
    readonly servers: string[];
    /**
     * Wether to use StartTLS for LDAP connections
     */
    readonly startTls: boolean;
    /**
     * Timeout for LDAP connections
     */
    readonly timeout: string;
    /**
     * User base for LDAP connections
     */
    readonly userBase: string;
    /**
     * User filters for LDAP connections
     */
    readonly userFilters: string[];
    /**
     * User object classes for LDAP connections
     */
    readonly userObjectClasses: string[];
}

export function getOrgIdpLdapOutput(args: GetOrgIdpLdapOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrgIdpLdapResult> {
    return pulumi.output(args).apply(a => getOrgIdpLdap(a, opts))
}

/**
 * A collection of arguments for invoking getOrgIdpLdap.
 */
export interface GetOrgIdpLdapOutputArgs {
    /**
     * The ID of this resource.
     */
    id: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
}
