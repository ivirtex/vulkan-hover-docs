# vkGetPhysicalDeviceFormatProperties(3) Manual Page

## Name

vkGetPhysicalDeviceFormatProperties - Lists physical device's format capabilities



## [](#_c_specification)C Specification

To query supported format features which are properties of the physical device, call:

Warning

This functionality is deprecated by [Vulkan Version 1.1](#versions-1.1). See [Deprecated Functionality](#deprecation-gpdp2) for more information.

```c++
// Provided by VK_VERSION_1_0
void vkGetPhysicalDeviceFormatProperties(
    VkPhysicalDevice                            physicalDevice,
    VkFormat                                    format,
    VkFormatProperties*                         pFormatProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device from which to query the format properties.
- `format` is the format whose properties are queried.
- `pFormatProperties` is a pointer to a [VkFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties.html) structure in which physical device properties for `format` are returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceFormatProperties-physicalDevice-parameter)VUID-vkGetPhysicalDeviceFormatProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceFormatProperties-format-parameter)VUID-vkGetPhysicalDeviceFormatProperties-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-vkGetPhysicalDeviceFormatProperties-pFormatProperties-parameter)VUID-vkGetPhysicalDeviceFormatProperties-pFormatProperties-parameter  
  `pFormatProperties` **must** be a valid pointer to a [VkFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties.html) structure

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceFormatProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0