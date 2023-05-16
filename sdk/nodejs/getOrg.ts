// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Datasource representing an organization in ZITADEL, which is the highest level after the instance and contains several other resource including policies if the configuration differs to the default policies on the instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const orgOrg = zitadel.getOrg({
 *     orgId: "177073608051458051",
 * });
 * export const org = orgOrg;
 * ```
 */
export function getOrg(args: GetOrgArgs, opts?: pulumi.InvokeOptions): Promise<GetOrgResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("zitadel:index/getOrg:getOrg", {
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrg.
 */
export interface GetOrgArgs {
    /**
     * The ID of this resource.
     */
    orgId: string;
}

/**
 * A collection of values returned by getOrg.
 */
export interface GetOrgResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name of the org
     */
    readonly name: string;
    /**
     * The ID of this resource.
     */
    readonly orgId: string;
}

export function getOrgOutput(args: GetOrgOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrgResult> {
    return pulumi.output(args).apply(a => getOrg(a, opts))
}

/**
 * A collection of arguments for invoking getOrg.
 */
export interface GetOrgOutputArgs {
    /**
     * The ID of this resource.
     */
    orgId: pulumi.Input<string>;
}