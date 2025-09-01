# VkMemoryBarrierAccessFlags3KHR(3) Manual Page

## Name

VkMemoryBarrierAccessFlags3KHR - Structure specifying additional access flags



## [](#_c_specification)C Specification

The `VkMemoryBarrierAccessFlags3KHR` structure is defined as:

```c++
// Provided by VK_KHR_maintenance8
typedef struct VkMemoryBarrierAccessFlags3KHR {
    VkStructureType      sType;
    const void*          pNext;
    VkAccessFlags3KHR    srcAccessMask3;
    VkAccessFlags3KHR    dstAccessMask3;
} VkMemoryBarrierAccessFlags3KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcAccessMask3` is a [VkAccessFlags3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags3KHR.html) mask of access flags to be included in the [first access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes).
- `dstAccessMask3` is a [VkAccessFlags3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags3KHR.html) mask of access flags to be included in the [second access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkMemoryBarrierAccessFlags3KHR-sType-sType)VUID-VkMemoryBarrierAccessFlags3KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_BARRIER_ACCESS_FLAGS_3_KHR`
- [](#VUID-VkMemoryBarrierAccessFlags3KHR-srcAccessMask3-parameter)VUID-VkMemoryBarrierAccessFlags3KHR-srcAccessMask3-parameter  
  `srcAccessMask3` **must** be a valid combination of [VkAccessFlagBits3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits3KHR.html) values
- [](#VUID-VkMemoryBarrierAccessFlags3KHR-dstAccessMask3-parameter)VUID-VkMemoryBarrierAccessFlags3KHR-dstAccessMask3-parameter  
  `dstAccessMask3` **must** be a valid combination of [VkAccessFlagBits3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits3KHR.html) values

## [](#_see_also)See Also

[VK\_KHR\_maintenance8](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance8.html), [VkAccessFlags3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags3KHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryBarrierAccessFlags3KHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0