# vkGetPrivateData(3) Manual Page

## Name

vkGetPrivateData - Retrieve data associated with a Vulkan object



## [](#_c_specification)C Specification

To retrieve application-defined data from a slot associated with a Vulkan object, call:

```c++
// Provided by VK_VERSION_1_3
void vkGetPrivateData(
    VkDevice                                    device,
    VkObjectType                                objectType,
    uint64_t                                    objectHandle,
    VkPrivateDataSlot                           privateDataSlot,
    uint64_t*                                   pData);
```

or the equivalent command

```c++
// Provided by VK_EXT_private_data
void vkGetPrivateDataEXT(
    VkDevice                                    device,
    VkObjectType                                objectType,
    uint64_t                                    objectHandle,
    VkPrivateDataSlot                           privateDataSlot,
    uint64_t*                                   pData);
```

## [](#_parameters)Parameters

- `device` is the device that created the object
- `objectType` is a [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html) specifying the type of object data is associated with.
- `objectHandle` is a handle to the object data is associated with.
- `privateDataSlot` is a handle to a [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html) specifying location of private data pointer storage.
- `pData` is a pointer to specify where application-defined data is returned. `0` will be written in the absence of a previous call to `vkSetPrivateData` using the object specified by `objectHandle`.

## [](#_description)Description

Note

Due to platform details on Android, implementations might not be able to reliably return `0` from calls to `vkGetPrivateData` for [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) objects on which `vkSetPrivateData` has not previously been called. This erratum is exclusive to the Android platform and objects of type [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html).

Valid Usage

- [](#VUID-vkGetPrivateData-objectType-04018)VUID-vkGetPrivateData-objectType-04018  
  `objectHandle` **must** be `device` or a child of `device`
- [](#VUID-vkGetPrivateData-objectHandle-09498)VUID-vkGetPrivateData-objectHandle-09498  
  `objectHandle` **must** be a valid handle to an object of type `objectType`

Valid Usage (Implicit)

- [](#VUID-vkGetPrivateData-device-parameter)VUID-vkGetPrivateData-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetPrivateData-objectType-parameter)VUID-vkGetPrivateData-objectType-parameter  
  `objectType` **must** be a valid [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html) value
- [](#VUID-vkGetPrivateData-privateDataSlot-parameter)VUID-vkGetPrivateData-privateDataSlot-parameter  
  `privateDataSlot` **must** be a valid [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html) handle
- [](#VUID-vkGetPrivateData-pData-parameter)VUID-vkGetPrivateData-pData-parameter  
  `pData` **must** be a valid pointer to a `uint64_t` value
- [](#VUID-vkGetPrivateData-privateDataSlot-parent)VUID-vkGetPrivateData-privateDataSlot-parent  
  `privateDataSlot` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_EXT\_private\_data](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_private_data.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html), [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPrivateData).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0