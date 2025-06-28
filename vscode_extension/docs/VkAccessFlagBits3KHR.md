# VkAccessFlagBits3KHR(3) Manual Page

## Name

VkAccessFlagBits3KHR - Access flags for VkAccessFlags3KHR



## [](#_c_specification)C Specification

Bits which **can** be set in the `srcAccessMask3` and `dstAccessMask3` members of [VkMemoryBarrierAccessFlags3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrierAccessFlags3KHR.html), specifying access behavior, are:

```c++
// Provided by VK_KHR_maintenance8
// Flag bits for VkAccessFlagBits3KHR
typedef VkFlags64 VkAccessFlagBits3KHR;
static const VkAccessFlagBits3KHR VK_ACCESS_3_NONE_KHR = 0ULL;
```

## [](#_description)Description

- `VK_ACCESS_3_NONE_KHR` specifies no additional accesses.

## [](#_see_also)See Also

[VK\_KHR\_maintenance8](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance8.html), [VkAccessFlags3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags3KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccessFlagBits3KHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0