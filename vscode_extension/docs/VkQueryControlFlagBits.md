# VkQueryControlFlagBits(3) Manual Page

## Name

VkQueryControlFlagBits - Bitmask specifying constraints on a query



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[vkCmdBeginQuery](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginQuery.html)::`flags`, specifying constraints
on the types of queries that **can** be performed, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkQueryControlFlagBits {
    VK_QUERY_CONTROL_PRECISE_BIT = 0x00000001,
} VkQueryControlFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_QUERY_CONTROL_PRECISE_BIT` specifies the precision of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-occlusion"
  target="_blank" rel="noopener">occlusion queries</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkQueryControlFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryControlFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueryControlFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
