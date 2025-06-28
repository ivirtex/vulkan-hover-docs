# vkGetPhysicalDeviceSparseImageFormatProperties2(3) Manual Page

## Name

vkGetPhysicalDeviceSparseImageFormatProperties2 - Retrieve properties of an image format applied to sparse images



## [](#_c_specification)C Specification

`vkGetPhysicalDeviceSparseImageFormatProperties2` returns an array of [VkSparseImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageFormatProperties2.html). Each element describes properties for one set of image aspects that are bound simultaneously for a `VkImage` created with the provided image creation parameters. This is usually one element for each aspect in the image, but for interleaved depth/stencil images there is only one element describing the combined aspects.

```c++
// Provided by VK_VERSION_1_1
void vkGetPhysicalDeviceSparseImageFormatProperties2(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceSparseImageFormatInfo2* pFormatInfo,
    uint32_t*                                   pPropertyCount,
    VkSparseImageFormatProperties2*             pProperties);
```

or the equivalent command

```c++
// Provided by VK_KHR_get_physical_device_properties2
void vkGetPhysicalDeviceSparseImageFormatProperties2KHR(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceSparseImageFormatInfo2* pFormatInfo,
    uint32_t*                                   pPropertyCount,
    VkSparseImageFormatProperties2*             pProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device from which to query the sparse image format properties.
- `pFormatInfo` is a pointer to a [VkPhysicalDeviceSparseImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSparseImageFormatInfo2.html) structure containing input parameters to the command.
- `pPropertyCount` is a pointer to an integer related to the number of sparse format properties available or queried, as described below.
- `pProperties` is either `NULL` or a pointer to an array of [VkSparseImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageFormatProperties2.html) structures.

## [](#_description)Description

`vkGetPhysicalDeviceSparseImageFormatProperties2` behaves identically to [vkGetPhysicalDeviceSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSparseImageFormatProperties.html), with the ability to return extended information by adding extending structures to the `pNext` chain of its `pProperties` parameter.

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-physicalDevice-parameter)VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-pFormatInfo-parameter)VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-pFormatInfo-parameter  
  `pFormatInfo` **must** be a valid pointer to a valid [VkPhysicalDeviceSparseImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSparseImageFormatInfo2.html) structure
- [](#VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-pPropertyCount-parameter)VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-pProperties-parameter)VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer to an array of `pPropertyCount` [VkSparseImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageFormatProperties2.html) structures

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceSparseImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSparseImageFormatInfo2.html), [VkSparseImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageFormatProperties2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceSparseImageFormatProperties2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0