# vkSetPrivateData(3) Manual Page

## Name

vkSetPrivateData - Associate data with a Vulkan object



## [](#_c_specification)C Specification

To store application-defined data in a slot associated with a Vulkan object, call:

```c++
// Provided by VK_VERSION_1_3
VkResult vkSetPrivateData(
    VkDevice                                    device,
    VkObjectType                                objectType,
    uint64_t                                    objectHandle,
    VkPrivateDataSlot                           privateDataSlot,
    uint64_t                                    data);
```

or the equivalent command

```c++
// Provided by VK_EXT_private_data
VkResult vkSetPrivateDataEXT(
    VkDevice                                    device,
    VkObjectType                                objectType,
    uint64_t                                    objectHandle,
    VkPrivateDataSlot                           privateDataSlot,
    uint64_t                                    data);
```

## [](#_parameters)Parameters

- `device` is the device that created the object.
- `objectType` is a [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html) specifying the type of object to associate data with.
- `objectHandle` is a handle to the object to associate data with.
- `privateDataSlot` is a handle to a [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html) specifying location of private data storage.
- `data` is application-defined data to associate the object with. This data will be stored at `privateDataSlot`.

## [](#_description)Description

Valid Usage

- [](#VUID-vkSetPrivateData-objectHandle-04016)VUID-vkSetPrivateData-objectHandle-04016  
  `objectHandle` **must** be `device` or a child of `device`
- [](#VUID-vkSetPrivateData-objectHandle-04017)VUID-vkSetPrivateData-objectHandle-04017  
  `objectHandle` **must** be a valid handle to an object of type `objectType`

Valid Usage (Implicit)

- [](#VUID-vkSetPrivateData-device-parameter)VUID-vkSetPrivateData-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkSetPrivateData-objectType-parameter)VUID-vkSetPrivateData-objectType-parameter  
  `objectType` **must** be a valid [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html) value
- [](#VUID-vkSetPrivateData-privateDataSlot-parameter)VUID-vkSetPrivateData-privateDataSlot-parameter  
  `privateDataSlot` **must** be a valid [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html) handle
- [](#VUID-vkSetPrivateData-privateDataSlot-parent)VUID-vkSetPrivateData-privateDataSlot-parent  
  `privateDataSlot` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_private\_data](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_private_data.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html), [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkSetPrivateData)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0