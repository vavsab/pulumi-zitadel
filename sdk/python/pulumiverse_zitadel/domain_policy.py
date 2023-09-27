# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DomainPolicyArgs', 'DomainPolicy']

@pulumi.input_type
class DomainPolicyArgs:
    def __init__(__self__, *,
                 smtp_sender_address_matches_instance_domain: pulumi.Input[bool],
                 user_login_must_be_domain: pulumi.Input[bool],
                 validate_org_domains: pulumi.Input[bool],
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DomainPolicy resource.
        :param pulumi.Input[bool] user_login_must_be_domain: User login must be domain
        :param pulumi.Input[bool] validate_org_domains: Validate organization domains
        :param pulumi.Input[str] org_id: ID of the organization
        """
        pulumi.set(__self__, "smtp_sender_address_matches_instance_domain", smtp_sender_address_matches_instance_domain)
        pulumi.set(__self__, "user_login_must_be_domain", user_login_must_be_domain)
        pulumi.set(__self__, "validate_org_domains", validate_org_domains)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)

    @property
    @pulumi.getter(name="smtpSenderAddressMatchesInstanceDomain")
    def smtp_sender_address_matches_instance_domain(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "smtp_sender_address_matches_instance_domain")

    @smtp_sender_address_matches_instance_domain.setter
    def smtp_sender_address_matches_instance_domain(self, value: pulumi.Input[bool]):
        pulumi.set(self, "smtp_sender_address_matches_instance_domain", value)

    @property
    @pulumi.getter(name="userLoginMustBeDomain")
    def user_login_must_be_domain(self) -> pulumi.Input[bool]:
        """
        User login must be domain
        """
        return pulumi.get(self, "user_login_must_be_domain")

    @user_login_must_be_domain.setter
    def user_login_must_be_domain(self, value: pulumi.Input[bool]):
        pulumi.set(self, "user_login_must_be_domain", value)

    @property
    @pulumi.getter(name="validateOrgDomains")
    def validate_org_domains(self) -> pulumi.Input[bool]:
        """
        Validate organization domains
        """
        return pulumi.get(self, "validate_org_domains")

    @validate_org_domains.setter
    def validate_org_domains(self, value: pulumi.Input[bool]):
        pulumi.set(self, "validate_org_domains", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)


@pulumi.input_type
class _DomainPolicyState:
    def __init__(__self__, *,
                 org_id: Optional[pulumi.Input[str]] = None,
                 smtp_sender_address_matches_instance_domain: Optional[pulumi.Input[bool]] = None,
                 user_login_must_be_domain: Optional[pulumi.Input[bool]] = None,
                 validate_org_domains: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering DomainPolicy resources.
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[bool] user_login_must_be_domain: User login must be domain
        :param pulumi.Input[bool] validate_org_domains: Validate organization domains
        """
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if smtp_sender_address_matches_instance_domain is not None:
            pulumi.set(__self__, "smtp_sender_address_matches_instance_domain", smtp_sender_address_matches_instance_domain)
        if user_login_must_be_domain is not None:
            pulumi.set(__self__, "user_login_must_be_domain", user_login_must_be_domain)
        if validate_org_domains is not None:
            pulumi.set(__self__, "validate_org_domains", validate_org_domains)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="smtpSenderAddressMatchesInstanceDomain")
    def smtp_sender_address_matches_instance_domain(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "smtp_sender_address_matches_instance_domain")

    @smtp_sender_address_matches_instance_domain.setter
    def smtp_sender_address_matches_instance_domain(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "smtp_sender_address_matches_instance_domain", value)

    @property
    @pulumi.getter(name="userLoginMustBeDomain")
    def user_login_must_be_domain(self) -> Optional[pulumi.Input[bool]]:
        """
        User login must be domain
        """
        return pulumi.get(self, "user_login_must_be_domain")

    @user_login_must_be_domain.setter
    def user_login_must_be_domain(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "user_login_must_be_domain", value)

    @property
    @pulumi.getter(name="validateOrgDomains")
    def validate_org_domains(self) -> Optional[pulumi.Input[bool]]:
        """
        Validate organization domains
        """
        return pulumi.get(self, "validate_org_domains")

    @validate_org_domains.setter
    def validate_org_domains(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "validate_org_domains", value)


class DomainPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 smtp_sender_address_matches_instance_domain: Optional[pulumi.Input[bool]] = None,
                 user_login_must_be_domain: Optional[pulumi.Input[bool]] = None,
                 validate_org_domains: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Resource representing the custom domain policy of an organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.DomainPolicy("default",
            org_id=data["zitadel_org"]["default"]["id"],
            user_login_must_be_domain=False,
            validate_org_domains=True,
            smtp_sender_address_matches_instance_domain=True)
        ```

        ## Import

        terraform # The resource can be imported using the ID format `<[org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/domainPolicy:DomainPolicy imported '123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[bool] user_login_must_be_domain: User login must be domain
        :param pulumi.Input[bool] validate_org_domains: Validate organization domains
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DomainPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing the custom domain policy of an organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.DomainPolicy("default",
            org_id=data["zitadel_org"]["default"]["id"],
            user_login_must_be_domain=False,
            validate_org_domains=True,
            smtp_sender_address_matches_instance_domain=True)
        ```

        ## Import

        terraform # The resource can be imported using the ID format `<[org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/domainPolicy:DomainPolicy imported '123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param DomainPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 smtp_sender_address_matches_instance_domain: Optional[pulumi.Input[bool]] = None,
                 user_login_must_be_domain: Optional[pulumi.Input[bool]] = None,
                 validate_org_domains: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DomainPolicyArgs.__new__(DomainPolicyArgs)

            __props__.__dict__["org_id"] = org_id
            if smtp_sender_address_matches_instance_domain is None and not opts.urn:
                raise TypeError("Missing required property 'smtp_sender_address_matches_instance_domain'")
            __props__.__dict__["smtp_sender_address_matches_instance_domain"] = smtp_sender_address_matches_instance_domain
            if user_login_must_be_domain is None and not opts.urn:
                raise TypeError("Missing required property 'user_login_must_be_domain'")
            __props__.__dict__["user_login_must_be_domain"] = user_login_must_be_domain
            if validate_org_domains is None and not opts.urn:
                raise TypeError("Missing required property 'validate_org_domains'")
            __props__.__dict__["validate_org_domains"] = validate_org_domains
        super(DomainPolicy, __self__).__init__(
            'zitadel:index/domainPolicy:DomainPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            smtp_sender_address_matches_instance_domain: Optional[pulumi.Input[bool]] = None,
            user_login_must_be_domain: Optional[pulumi.Input[bool]] = None,
            validate_org_domains: Optional[pulumi.Input[bool]] = None) -> 'DomainPolicy':
        """
        Get an existing DomainPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[bool] user_login_must_be_domain: User login must be domain
        :param pulumi.Input[bool] validate_org_domains: Validate organization domains
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DomainPolicyState.__new__(_DomainPolicyState)

        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["smtp_sender_address_matches_instance_domain"] = smtp_sender_address_matches_instance_domain
        __props__.__dict__["user_login_must_be_domain"] = user_login_must_be_domain
        __props__.__dict__["validate_org_domains"] = validate_org_domains
        return DomainPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="smtpSenderAddressMatchesInstanceDomain")
    def smtp_sender_address_matches_instance_domain(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "smtp_sender_address_matches_instance_domain")

    @property
    @pulumi.getter(name="userLoginMustBeDomain")
    def user_login_must_be_domain(self) -> pulumi.Output[bool]:
        """
        User login must be domain
        """
        return pulumi.get(self, "user_login_must_be_domain")

    @property
    @pulumi.getter(name="validateOrgDomains")
    def validate_org_domains(self) -> pulumi.Output[bool]:
        """
        Validate organization domains
        """
        return pulumi.get(self, "validate_org_domains")

