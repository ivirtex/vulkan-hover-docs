# vkGetPhysicalDeviceMemoryProperties2(3) Manual Page

## Name

vkGetPhysicalDeviceMemoryProperties2 - Reports memory information for the specified physical device



## [](#_c_specification)C Specification

To query memory properties, call:

```c++
// Provided by VK_VERSION_1_1
void vkGetPhysicalDeviceMemoryProperties2(
    VkPhysicalDevice                            physicalDevice,
    VkPhysicalDeviceMemoryProperties2*          pMemoryProperties);
```

or the equivalent command

```c++
// Provided by VK_KHR_get_physical_device_properties2
void vkGetPhysicalDeviceMemoryProperties2KHR(
    VkPhysicalDevice                            physicalDevice,
    VkPhysicalDeviceMemoryProperties2*          pMemoryProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the handle to the device to query.
- `pMemoryProperties` is a pointer to a [VkPhysicalDeviceMemoryProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties2.html) structure in which the properties are returned.

## [](#_description)Description

`vkGetPhysicalDeviceMemoryProperties2` behaves similarly to [vkGetPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceMemoryProperties.html), with the ability to return extended information in a `pNext` chain of output structures.

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceMemoryProperties2-physicalDevice-parameter)VUID-vkGetPhysicalDeviceMemoryProperties2-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceMemoryProperties2-pMemoryProperties-parameter)VUID-vkGetPhysicalDeviceMemoryProperties2-pMemoryProperties-parameter  
  `pMemoryProperties` **must** be a valid pointer to a [VkPhysicalDeviceMemoryProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties2.html) structure

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceMemoryProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceMemoryProperties2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0