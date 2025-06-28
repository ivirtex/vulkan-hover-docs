# vkGetPhysicalDeviceFeatures(3) Manual Page

## Name

vkGetPhysicalDeviceFeatures - Reports capabilities of a physical device



## [](#_c_specification)C Specification

To query supported features, call:

Warning

This functionality is deprecated by [Vulkan Version 1.1](#versions-1.1). See [Deprecated Functionality](#deprecation-gpdp2) for more information.

```c++
// Provided by VK_VERSION_1_0
void vkGetPhysicalDeviceFeatures(
    VkPhysicalDevice                            physicalDevice,
    VkPhysicalDeviceFeatures*                   pFeatures);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device from which to query the supported features.
- `pFeatures` is a pointer to a [VkPhysicalDeviceFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures.html) structure in which the physical device features are returned. For each feature, a value of `VK_TRUE` specifies that the feature is supported on this physical device, and `VK_FALSE` specifies that the feature is not supported.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceFeatures-physicalDevice-parameter)VUID-vkGetPhysicalDeviceFeatures-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceFeatures-pFeatures-parameter)VUID-vkGetPhysicalDeviceFeatures-pFeatures-parameter  
  `pFeatures` **must** be a valid pointer to a [VkPhysicalDeviceFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures.html) structure

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceFeatures)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0