# vkGetPhysicalDeviceExternalFenceProperties(3) Manual Page

## Name

vkGetPhysicalDeviceExternalFenceProperties - Function for querying external fence handle capabilities.



## [](#_c_specification)C Specification

Fences **may** support import and export of their [payload](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-fences-payloads) to external handles. To query the external handle types supported by fences, call:

```c++
// Provided by VK_VERSION_1_1
void vkGetPhysicalDeviceExternalFenceProperties(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceExternalFenceInfo*    pExternalFenceInfo,
    VkExternalFenceProperties*                  pExternalFenceProperties);
```

or the equivalent command

```c++
// Provided by VK_KHR_external_fence_capabilities
void vkGetPhysicalDeviceExternalFencePropertiesKHR(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceExternalFenceInfo*    pExternalFenceInfo,
    VkExternalFenceProperties*                  pExternalFenceProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device from which to query the fence capabilities.
- `pExternalFenceInfo` is a pointer to a [VkPhysicalDeviceExternalFenceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalFenceInfo.html) structure describing the parameters that would be consumed by [vkCreateFence](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateFence.html).
- `pExternalFenceProperties` is a pointer to a [VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceProperties.html) structure in which capabilities are returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceExternalFenceProperties-physicalDevice-parameter)VUID-vkGetPhysicalDeviceExternalFenceProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceExternalFenceProperties-pExternalFenceInfo-parameter)VUID-vkGetPhysicalDeviceExternalFenceProperties-pExternalFenceInfo-parameter  
  `pExternalFenceInfo` **must** be a valid pointer to a valid [VkPhysicalDeviceExternalFenceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalFenceInfo.html) structure
- [](#VUID-vkGetPhysicalDeviceExternalFenceProperties-pExternalFenceProperties-parameter)VUID-vkGetPhysicalDeviceExternalFenceProperties-pExternalFenceProperties-parameter  
  `pExternalFenceProperties` **must** be a valid pointer to a [VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceProperties.html) structure

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceProperties.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceExternalFenceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalFenceInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceExternalFenceProperties).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0