# vkGetImageMemoryRequirements2(3) Manual Page

## Name

vkGetImageMemoryRequirements2 - Returns the memory requirements for specified Vulkan object



## [](#_c_specification)C Specification

To determine the memory requirements for an image resource, call:

```c++
// Provided by VK_VERSION_1_1
void vkGetImageMemoryRequirements2(
    VkDevice                                    device,
    const VkImageMemoryRequirementsInfo2*       pInfo,
    VkMemoryRequirements2*                      pMemoryRequirements);
```

or the equivalent command

```c++
// Provided by VK_KHR_get_memory_requirements2
void vkGetImageMemoryRequirements2KHR(
    VkDevice                                    device,
    const VkImageMemoryRequirementsInfo2*       pInfo,
    VkMemoryRequirements2*                      pMemoryRequirements);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the image.
- `pInfo` is a pointer to a [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryRequirementsInfo2.html) structure containing parameters required for the memory requirements query.
- `pMemoryRequirements` is a pointer to a [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html) structure in which the memory requirements of the image object are returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetImageMemoryRequirements2-device-parameter)VUID-vkGetImageMemoryRequirements2-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetImageMemoryRequirements2-pInfo-parameter)VUID-vkGetImageMemoryRequirements2-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryRequirementsInfo2.html) structure
- [](#VUID-vkGetImageMemoryRequirements2-pMemoryRequirements-parameter)VUID-vkGetImageMemoryRequirements2-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html) structure

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryRequirementsInfo2.html), [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetImageMemoryRequirements2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0