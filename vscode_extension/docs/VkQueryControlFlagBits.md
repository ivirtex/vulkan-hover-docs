# VkQueryControlFlagBits(3) Manual Page

## Name

VkQueryControlFlagBits - Bitmask specifying constraints on a query



## [](#_c_specification)C Specification

Bits which **can** be set in [vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginQuery.html)::`flags`, specifying constraints on the types of queries that **can** be performed, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkQueryControlFlagBits {
    VK_QUERY_CONTROL_PRECISE_BIT = 0x00000001,
} VkQueryControlFlagBits;
```

## [](#_description)Description

- `VK_QUERY_CONTROL_PRECISE_BIT` specifies the precision of [occlusion queries](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#queries-occlusion).

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkQueryControlFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryControlFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueryControlFlagBits).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0