# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetApplicationOidcsResult',
    'AwaitableGetApplicationOidcsResult',
    'get_application_oidcs',
    'get_application_oidcs_output',
]

@pulumi.output_type
class GetApplicationOidcsResult:
    """
    A collection of values returned by getApplicationOidcs.
    """
    def __init__(__self__, app_ids=None, id=None, name=None, name_method=None, org_id=None, project_id=None):
        if app_ids and not isinstance(app_ids, list):
            raise TypeError("Expected argument 'app_ids' to be a list")
        pulumi.set(__self__, "app_ids", app_ids)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_method and not isinstance(name_method, str):
            raise TypeError("Expected argument 'name_method' to be a str")
        pulumi.set(__self__, "name_method", name_method)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="appIds")
    def app_ids(self) -> Sequence[str]:
        """
        A set of all IDs.
        """
        return pulumi.get(self, "app_ids")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the application
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameMethod")
    def name_method(self) -> Optional[str]:
        """
        Method for querying applications by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
        """
        return pulumi.get(self, "name_method")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[str]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        ID of the project
        """
        return pulumi.get(self, "project_id")


class AwaitableGetApplicationOidcsResult(GetApplicationOidcsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationOidcsResult(
            app_ids=self.app_ids,
            id=self.id,
            name=self.name,
            name_method=self.name_method,
            org_id=self.org_id,
            project_id=self.project_id)


def get_application_oidcs(name: Optional[str] = None,
                          name_method: Optional[str] = None,
                          org_id: Optional[str] = None,
                          project_id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationOidcsResult:
    """
    Datasource representing multiple OIDC applications belonging to a project.


    :param str name: Name of the application
    :param str name_method: Method for querying applications by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
    :param str org_id: ID of the organization
    :param str project_id: ID of the project
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['nameMethod'] = name_method
    __args__['orgId'] = org_id
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zitadel:index/getApplicationOidcs:getApplicationOidcs', __args__, opts=opts, typ=GetApplicationOidcsResult).value

    return AwaitableGetApplicationOidcsResult(
        app_ids=pulumi.get(__ret__, 'app_ids'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        name_method=pulumi.get(__ret__, 'name_method'),
        org_id=pulumi.get(__ret__, 'org_id'),
        project_id=pulumi.get(__ret__, 'project_id'))


@_utilities.lift_output_func(get_application_oidcs)
def get_application_oidcs_output(name: Optional[pulumi.Input[str]] = None,
                                 name_method: Optional[pulumi.Input[Optional[str]]] = None,
                                 org_id: Optional[pulumi.Input[Optional[str]]] = None,
                                 project_id: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationOidcsResult]:
    """
    Datasource representing multiple OIDC applications belonging to a project.


    :param str name: Name of the application
    :param str name_method: Method for querying applications by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
    :param str org_id: ID of the organization
    :param str project_id: ID of the project
    """
    ...