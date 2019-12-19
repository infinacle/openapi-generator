# coding: utf-8

"""
    OpenAPI Petstore

    This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import
import re  # noqa: F401
import sys  # noqa: F401

import six  # noqa: F401

from petstore_api.model_utils import (  # noqa: F401
    ModelComposed,
    ModelNormal,
    ModelSimple,
    date,
    datetime,
    file_type,
    int,
    none_type,
    str,
    validate_get_composed_info,
)


class TypeHolderDefault(ModelNormal):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.

    Attributes:
      allowed_values (dict): The key is the tuple path to the attribute
          and the for var_name this is (var_name,). The value is a dict
          with a capitalized key describing the allowed value and an allowed
          value. These dicts store the allowed enum values.
      attribute_map (dict): The key is attribute name
          and the value is json key in definition.
      discriminator_value_class_map (dict): A dict to go from the discriminator
          variable value to the discriminator class name.
      openapi_types (dict): The key is attribute name
          and the value is attribute type.
      validations (dict): The key is the tuple path to the attribute
          and the for var_name this is (var_name,). The value is a dict
          that stores validations for max_length, min_length, max_items,
          min_items, exclusive_maximum, inclusive_maximum, exclusive_minimum,
          inclusive_minimum, and regex.
      additional_properties_type (tuple): A tuple of classes accepted
          as additional properties values.
    """

    allowed_values = {
    }

    openapi_types = {
        'string_item': (str,),  # noqa: E501
        'number_item': (float,),  # noqa: E501
        'integer_item': (int,),  # noqa: E501
        'bool_item': (bool,),  # noqa: E501
        'array_item': ([int],),  # noqa: E501
        'date_item': (date,),  # noqa: E501
        'datetime_item': (datetime,),  # noqa: E501
    }

    validations = {
    }

    additional_properties_type = None

    @staticmethod
    def discriminator():
        return None

    attribute_map = {
        'string_item': 'string_item',  # noqa: E501
        'number_item': 'number_item',  # noqa: E501
        'integer_item': 'integer_item',  # noqa: E501
        'bool_item': 'bool_item',  # noqa: E501
        'array_item': 'array_item',  # noqa: E501
        'date_item': 'date_item',  # noqa: E501
        'datetime_item': 'datetime_item',  # noqa: E501
    }

    @staticmethod
    def _composed_schemas():
        return None

    required_properties = set([
        '_data_store',
        '_check_type',
        '_from_server',
        '_path_to_item',
        '_configuration',
    ])

    def __init__(self, array_item, string_item='what', number_item=1.234, integer_item=-2, bool_item=True, _check_type=True, _from_server=False, _path_to_item=(), _configuration=None, **kwargs):  # noqa: E501
        """type_holder_default.TypeHolderDefault - a model defined in OpenAPI

        Args:
            array_item ([int]):

        Keyword Args:
            string_item (str): defaults to 'what', must be one of ['what']  # noqa: E501
            number_item (float): defaults to 1.234, must be one of [1.234]  # noqa: E501
            integer_item (int): defaults to -2, must be one of [-2]  # noqa: E501
            bool_item (bool): defaults to True, must be one of [True]  # noqa: E501
            _check_type (bool): if True, values for parameters in openapi_types
                                will be type checked and a TypeError will be
                                raised if the wrong type is input.
                                Defaults to True
            _path_to_item (tuple/list): This is a list of keys or values to
                                drill down to the model in received_data
                                when deserializing a response
            _from_server (bool): True if the data is from the server
                                False if the data is from the client (default)
            _configuration (Configuration): the instance to use when
                                deserializing a file_type parameter.
                                If passed, type conversion is attempted
                                If omitted no type conversion is done.
            date_item (date): [optional]  # noqa: E501
            datetime_item (datetime): [optional]  # noqa: E501
        """

        self._data_store = {}
        self._check_type = _check_type
        self._from_server = _from_server
        self._path_to_item = _path_to_item
        self._configuration = _configuration

        self.string_item = string_item
        self.number_item = number_item
        self.integer_item = integer_item
        self.bool_item = bool_item
        self.array_item = array_item
        for var_name, var_value in six.iteritems(kwargs):
            setattr(self, var_name, var_value)
