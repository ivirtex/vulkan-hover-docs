# vkGetPhysicalDeviceProperties(3) Manual Page

## Name

vkGetPhysicalDeviceProperties - Returns properties of a physical device



## [](#_c_specification)C Specification

To query general properties of physical devices once enumerated, call:

Warning

This functionality is deprecated by [Vulkan Version 1.1](#versions-1.1). See [Deprecated Functionality](#deprecation-gpdp2) for more information.

```c++
// Provided by VK_VERSION_1_0
void vkGetPhysicalDeviceProperties(
    VkPhysicalDevice                            physicalDevice,
    VkPhysicalDeviceProperties*                 pProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the handle to the physical device whose properties will be queried.
- `pProperties` is a pointer to a [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html) structure in which properties are returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceProperties-physicalDevice-parameter)VUID-vkGetPhysicalDeviceProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceProperties-pProperties-parameter)VUID-vkGetPhysicalDeviceProperties-pProperties-parameter  
  `pProperties` **must** be a valid pointer to a [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html) structure

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0