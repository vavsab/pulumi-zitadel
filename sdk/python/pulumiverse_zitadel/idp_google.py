# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IdpGoogleArgs', 'IdpGoogle']

@pulumi.input_type
class IdpGoogleArgs:
    def __init__(__self__, *,
                 client_id: pulumi.Input[str],
                 client_secret: pulumi.Input[str],
                 is_auto_creation: pulumi.Input[bool],
                 is_auto_update: pulumi.Input[bool],
                 is_creation_allowed: pulumi.Input[bool],
                 is_linking_allowed: pulumi.Input[bool],
                 name: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a IdpGoogle resource.
        :param pulumi.Input[str] client_id: client id generated by the identity provider
        :param pulumi.Input[str] client_secret: client secret generated by the identity provider
        :param pulumi.Input[bool] is_auto_creation: enable if a new account in ZITADEL should be created automatically on login with an external account
        :param pulumi.Input[bool] is_auto_update: enable if a the ZITADEL account fields should be updated automatically on each login
        :param pulumi.Input[bool] is_creation_allowed: enable if users should be able to create a new account in ZITADEL when using an external account
        :param pulumi.Input[bool] is_linking_allowed: enable if users should be able to link an existing ZITADEL user with an external account
        :param pulumi.Input[str] name: Name of the IDP
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: the scopes requested by ZITADEL during the request on the identity provider
        """
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "client_secret", client_secret)
        pulumi.set(__self__, "is_auto_creation", is_auto_creation)
        pulumi.set(__self__, "is_auto_update", is_auto_update)
        pulumi.set(__self__, "is_creation_allowed", is_creation_allowed)
        pulumi.set(__self__, "is_linking_allowed", is_linking_allowed)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if scopes is not None:
            pulumi.set(__self__, "scopes", scopes)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Input[str]:
        """
        client id generated by the identity provider
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Input[str]:
        """
        client secret generated by the identity provider
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter(name="isAutoCreation")
    def is_auto_creation(self) -> pulumi.Input[bool]:
        """
        enable if a new account in ZITADEL should be created automatically on login with an external account
        """
        return pulumi.get(self, "is_auto_creation")

    @is_auto_creation.setter
    def is_auto_creation(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_auto_creation", value)

    @property
    @pulumi.getter(name="isAutoUpdate")
    def is_auto_update(self) -> pulumi.Input[bool]:
        """
        enable if a the ZITADEL account fields should be updated automatically on each login
        """
        return pulumi.get(self, "is_auto_update")

    @is_auto_update.setter
    def is_auto_update(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_auto_update", value)

    @property
    @pulumi.getter(name="isCreationAllowed")
    def is_creation_allowed(self) -> pulumi.Input[bool]:
        """
        enable if users should be able to create a new account in ZITADEL when using an external account
        """
        return pulumi.get(self, "is_creation_allowed")

    @is_creation_allowed.setter
    def is_creation_allowed(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_creation_allowed", value)

    @property
    @pulumi.getter(name="isLinkingAllowed")
    def is_linking_allowed(self) -> pulumi.Input[bool]:
        """
        enable if users should be able to link an existing ZITADEL user with an external account
        """
        return pulumi.get(self, "is_linking_allowed")

    @is_linking_allowed.setter
    def is_linking_allowed(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_linking_allowed", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the IDP
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        the scopes requested by ZITADEL during the request on the identity provider
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "scopes", value)


@pulumi.input_type
class _IdpGoogleState:
    def __init__(__self__, *,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 is_auto_creation: Optional[pulumi.Input[bool]] = None,
                 is_auto_update: Optional[pulumi.Input[bool]] = None,
                 is_creation_allowed: Optional[pulumi.Input[bool]] = None,
                 is_linking_allowed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering IdpGoogle resources.
        :param pulumi.Input[str] client_id: client id generated by the identity provider
        :param pulumi.Input[str] client_secret: client secret generated by the identity provider
        :param pulumi.Input[bool] is_auto_creation: enable if a new account in ZITADEL should be created automatically on login with an external account
        :param pulumi.Input[bool] is_auto_update: enable if a the ZITADEL account fields should be updated automatically on each login
        :param pulumi.Input[bool] is_creation_allowed: enable if users should be able to create a new account in ZITADEL when using an external account
        :param pulumi.Input[bool] is_linking_allowed: enable if users should be able to link an existing ZITADEL user with an external account
        :param pulumi.Input[str] name: Name of the IDP
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: the scopes requested by ZITADEL during the request on the identity provider
        """
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if is_auto_creation is not None:
            pulumi.set(__self__, "is_auto_creation", is_auto_creation)
        if is_auto_update is not None:
            pulumi.set(__self__, "is_auto_update", is_auto_update)
        if is_creation_allowed is not None:
            pulumi.set(__self__, "is_creation_allowed", is_creation_allowed)
        if is_linking_allowed is not None:
            pulumi.set(__self__, "is_linking_allowed", is_linking_allowed)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if scopes is not None:
            pulumi.set(__self__, "scopes", scopes)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        client id generated by the identity provider
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[pulumi.Input[str]]:
        """
        client secret generated by the identity provider
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter(name="isAutoCreation")
    def is_auto_creation(self) -> Optional[pulumi.Input[bool]]:
        """
        enable if a new account in ZITADEL should be created automatically on login with an external account
        """
        return pulumi.get(self, "is_auto_creation")

    @is_auto_creation.setter
    def is_auto_creation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_auto_creation", value)

    @property
    @pulumi.getter(name="isAutoUpdate")
    def is_auto_update(self) -> Optional[pulumi.Input[bool]]:
        """
        enable if a the ZITADEL account fields should be updated automatically on each login
        """
        return pulumi.get(self, "is_auto_update")

    @is_auto_update.setter
    def is_auto_update(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_auto_update", value)

    @property
    @pulumi.getter(name="isCreationAllowed")
    def is_creation_allowed(self) -> Optional[pulumi.Input[bool]]:
        """
        enable if users should be able to create a new account in ZITADEL when using an external account
        """
        return pulumi.get(self, "is_creation_allowed")

    @is_creation_allowed.setter
    def is_creation_allowed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_creation_allowed", value)

    @property
    @pulumi.getter(name="isLinkingAllowed")
    def is_linking_allowed(self) -> Optional[pulumi.Input[bool]]:
        """
        enable if users should be able to link an existing ZITADEL user with an external account
        """
        return pulumi.get(self, "is_linking_allowed")

    @is_linking_allowed.setter
    def is_linking_allowed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_linking_allowed", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the IDP
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        the scopes requested by ZITADEL during the request on the identity provider
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "scopes", value)


class IdpGoogle(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 is_auto_creation: Optional[pulumi.Input[bool]] = None,
                 is_auto_update: Optional[pulumi.Input[bool]] = None,
                 is_creation_allowed: Optional[pulumi.Input[bool]] = None,
                 is_linking_allowed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Resource representing a Google IDP on the instance.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.IdpGoogle("default",
            client_id="182902...",
            client_secret="GOCSPX-*****",
            is_auto_creation=False,
            is_auto_update=True,
            is_creation_allowed=True,
            is_linking_allowed=False,
            scopes=[
                "openid",
                "profile",
                "email",
            ])
        ```

        ## Import

        terraform # The resource can be imported using the ID format `<id[:client_secret]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/idpGoogle:IdpGoogle imported '123456789012345678:G1234567890123'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: client id generated by the identity provider
        :param pulumi.Input[str] client_secret: client secret generated by the identity provider
        :param pulumi.Input[bool] is_auto_creation: enable if a new account in ZITADEL should be created automatically on login with an external account
        :param pulumi.Input[bool] is_auto_update: enable if a the ZITADEL account fields should be updated automatically on each login
        :param pulumi.Input[bool] is_creation_allowed: enable if users should be able to create a new account in ZITADEL when using an external account
        :param pulumi.Input[bool] is_linking_allowed: enable if users should be able to link an existing ZITADEL user with an external account
        :param pulumi.Input[str] name: Name of the IDP
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: the scopes requested by ZITADEL during the request on the identity provider
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IdpGoogleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing a Google IDP on the instance.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.IdpGoogle("default",
            client_id="182902...",
            client_secret="GOCSPX-*****",
            is_auto_creation=False,
            is_auto_update=True,
            is_creation_allowed=True,
            is_linking_allowed=False,
            scopes=[
                "openid",
                "profile",
                "email",
            ])
        ```

        ## Import

        terraform # The resource can be imported using the ID format `<id[:client_secret]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/idpGoogle:IdpGoogle imported '123456789012345678:G1234567890123'
        ```

        :param str resource_name: The name of the resource.
        :param IdpGoogleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IdpGoogleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 is_auto_creation: Optional[pulumi.Input[bool]] = None,
                 is_auto_update: Optional[pulumi.Input[bool]] = None,
                 is_creation_allowed: Optional[pulumi.Input[bool]] = None,
                 is_linking_allowed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IdpGoogleArgs.__new__(IdpGoogleArgs)

            if client_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_id'")
            __props__.__dict__["client_id"] = client_id
            if client_secret is None and not opts.urn:
                raise TypeError("Missing required property 'client_secret'")
            __props__.__dict__["client_secret"] = client_secret
            if is_auto_creation is None and not opts.urn:
                raise TypeError("Missing required property 'is_auto_creation'")
            __props__.__dict__["is_auto_creation"] = is_auto_creation
            if is_auto_update is None and not opts.urn:
                raise TypeError("Missing required property 'is_auto_update'")
            __props__.__dict__["is_auto_update"] = is_auto_update
            if is_creation_allowed is None and not opts.urn:
                raise TypeError("Missing required property 'is_creation_allowed'")
            __props__.__dict__["is_creation_allowed"] = is_creation_allowed
            if is_linking_allowed is None and not opts.urn:
                raise TypeError("Missing required property 'is_linking_allowed'")
            __props__.__dict__["is_linking_allowed"] = is_linking_allowed
            __props__.__dict__["name"] = name
            __props__.__dict__["scopes"] = scopes
        super(IdpGoogle, __self__).__init__(
            'zitadel:index/idpGoogle:IdpGoogle',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_secret: Optional[pulumi.Input[str]] = None,
            is_auto_creation: Optional[pulumi.Input[bool]] = None,
            is_auto_update: Optional[pulumi.Input[bool]] = None,
            is_creation_allowed: Optional[pulumi.Input[bool]] = None,
            is_linking_allowed: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'IdpGoogle':
        """
        Get an existing IdpGoogle resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: client id generated by the identity provider
        :param pulumi.Input[str] client_secret: client secret generated by the identity provider
        :param pulumi.Input[bool] is_auto_creation: enable if a new account in ZITADEL should be created automatically on login with an external account
        :param pulumi.Input[bool] is_auto_update: enable if a the ZITADEL account fields should be updated automatically on each login
        :param pulumi.Input[bool] is_creation_allowed: enable if users should be able to create a new account in ZITADEL when using an external account
        :param pulumi.Input[bool] is_linking_allowed: enable if users should be able to link an existing ZITADEL user with an external account
        :param pulumi.Input[str] name: Name of the IDP
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: the scopes requested by ZITADEL during the request on the identity provider
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IdpGoogleState.__new__(_IdpGoogleState)

        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["client_secret"] = client_secret
        __props__.__dict__["is_auto_creation"] = is_auto_creation
        __props__.__dict__["is_auto_update"] = is_auto_update
        __props__.__dict__["is_creation_allowed"] = is_creation_allowed
        __props__.__dict__["is_linking_allowed"] = is_linking_allowed
        __props__.__dict__["name"] = name
        __props__.__dict__["scopes"] = scopes
        return IdpGoogle(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        """
        client id generated by the identity provider
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[str]:
        """
        client secret generated by the identity provider
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="isAutoCreation")
    def is_auto_creation(self) -> pulumi.Output[bool]:
        """
        enable if a new account in ZITADEL should be created automatically on login with an external account
        """
        return pulumi.get(self, "is_auto_creation")

    @property
    @pulumi.getter(name="isAutoUpdate")
    def is_auto_update(self) -> pulumi.Output[bool]:
        """
        enable if a the ZITADEL account fields should be updated automatically on each login
        """
        return pulumi.get(self, "is_auto_update")

    @property
    @pulumi.getter(name="isCreationAllowed")
    def is_creation_allowed(self) -> pulumi.Output[bool]:
        """
        enable if users should be able to create a new account in ZITADEL when using an external account
        """
        return pulumi.get(self, "is_creation_allowed")

    @property
    @pulumi.getter(name="isLinkingAllowed")
    def is_linking_allowed(self) -> pulumi.Output[bool]:
        """
        enable if users should be able to link an existing ZITADEL user with an external account
        """
        return pulumi.get(self, "is_linking_allowed")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the IDP
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        the scopes requested by ZITADEL during the request on the identity provider
        """
        return pulumi.get(self, "scopes")

