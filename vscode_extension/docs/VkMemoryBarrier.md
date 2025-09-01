# VkMemoryBarrier(3) Manual Page

## Name

VkMemoryBarrier - Structure specifying a global memory barrier



## [](#_c_specification)C Specification

The `VkMemoryBarrier` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkMemoryBarrier {
    VkStructureType    sType;
    const void*        pNext;
    VkAccessFlags      srcAccessMask;
    VkAccessFlags      dstAccessMask;
} VkMemoryBarrier;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcAccessMask` is a bitmask of [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) specifying a [source access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks).
- `dstAccessMask` is a bitmask of [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) specifying a [destination access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks).

## [](#_description)Description

The first [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) is limited to access types in the [source access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks) specified by `srcAccessMask` and, if a [VkMemoryBarrierAccessFlags3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrierAccessFlags3KHR.html) is passed in `pNext`, `srcAccessMask3`.

The second [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) is limited to access types in the [destination access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks) specified by `dstAccessMask` and, if a [VkMemoryBarrierAccessFlags3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrierAccessFlags3KHR.html) is passed in `pNext`, `dstAccessMask3`.

Valid Usage (Implicit)

- [](#VUID-VkMemoryBarrier-sType-sType)VUID-VkMemoryBarrier-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_BARRIER`
- [](#VUID-VkMemoryBarrier-pNext-pNext)VUID-VkMemoryBarrier-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkMemoryBarrier-srcAccessMask-parameter)VUID-VkMemoryBarrier-srcAccessMask-parameter  
  `srcAccessMask` **must** be a valid combination of [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) values
- [](#VUID-VkMemoryBarrier-dstAccessMask-parameter)VUID-VkMemoryBarrier-dstAccessMask-parameter  
  `dstAccessMask` **must** be a valid combination of [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) values

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAccessFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPipelineBarrier.html), [vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryBarrier).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0