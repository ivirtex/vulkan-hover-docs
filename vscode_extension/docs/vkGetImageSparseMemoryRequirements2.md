# vkGetImageSparseMemoryRequirements2(3) Manual Page

## Name

vkGetImageSparseMemoryRequirements2 - Query the memory requirements for a sparse image



## [](#_c_specification)C Specification

To query sparse memory requirements for an image, call:

```c++
// Provided by VK_VERSION_1_1
void vkGetImageSparseMemoryRequirements2(
    VkDevice                                    device,
    const VkImageSparseMemoryRequirementsInfo2* pInfo,
    uint32_t*                                   pSparseMemoryRequirementCount,
    VkSparseImageMemoryRequirements2*           pSparseMemoryRequirements);
```

or the equivalent command

```c++
// Provided by VK_KHR_get_memory_requirements2
void vkGetImageSparseMemoryRequirements2KHR(
    VkDevice                                    device,
    const VkImageSparseMemoryRequirementsInfo2* pInfo,
    uint32_t*                                   pSparseMemoryRequirementCount,
    VkSparseImageMemoryRequirements2*           pSparseMemoryRequirements);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the image.
- `pInfo` is a pointer to a [VkImageSparseMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSparseMemoryRequirementsInfo2.html) structure containing parameters required for the memory requirements query.
- `pSparseMemoryRequirementCount` is a pointer to an integer related to the number of sparse memory requirements available or queried, as described below.
- `pSparseMemoryRequirements` is either `NULL` or a pointer to an array of `VkSparseImageMemoryRequirements2` structures.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetImageSparseMemoryRequirements2-device-parameter)VUID-vkGetImageSparseMemoryRequirements2-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetImageSparseMemoryRequirements2-pInfo-parameter)VUID-vkGetImageSparseMemoryRequirements2-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkImageSparseMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSparseMemoryRequirementsInfo2.html) structure
- [](#VUID-vkGetImageSparseMemoryRequirements2-pSparseMemoryRequirementCount-parameter)VUID-vkGetImageSparseMemoryRequirements2-pSparseMemoryRequirementCount-parameter  
  `pSparseMemoryRequirementCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetImageSparseMemoryRequirements2-pSparseMemoryRequirements-parameter)VUID-vkGetImageSparseMemoryRequirements2-pSparseMemoryRequirements-parameter  
  If the value referenced by `pSparseMemoryRequirementCount` is not `0`, and `pSparseMemoryRequirements` is not `NULL`, `pSparseMemoryRequirements` **must** be a valid pointer to an array of `pSparseMemoryRequirementCount` [VkSparseImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryRequirements2.html) structures

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImageSparseMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSparseMemoryRequirementsInfo2.html), [VkSparseImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryRequirements2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetImageSparseMemoryRequirements2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0