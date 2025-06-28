# VkQueryPoolCreateFlagBits(3) Manual Page

## Name

VkQueryPoolCreateFlagBits - Bitmask specifying query pool properties



## [](#_c_specification)C Specification

Bits which **can** be set in [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html)::`flags`, specifying options for query pools, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkQueryPoolCreateFlagBits {
  // Provided by VK_KHR_maintenance9
    VK_QUERY_POOL_CREATE_RESET_BIT_KHR = 0x00000001,
} VkQueryPoolCreateFlagBits;
```

## [](#_description)Description

- `VK_QUERY_POOL_CREATE_RESET_BIT_KHR` specifies that queries in the query pool are initialized on creation and do not need to be reset before first use.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkQueryPoolCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueryPoolCreateFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0