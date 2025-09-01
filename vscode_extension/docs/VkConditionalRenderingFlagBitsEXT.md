# VkConditionalRenderingFlagBitsEXT(3) Manual Page

## Name

VkConditionalRenderingFlagBitsEXT - Specify the behavior of conditional rendering



## [](#_c_specification)C Specification

Bits which **can** be set in [vkCmdBeginConditionalRenderingEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginConditionalRenderingEXT.html)::`flags`, specifying the behavior of conditional rendering, are:

```c++
// Provided by VK_EXT_conditional_rendering
typedef enum VkConditionalRenderingFlagBitsEXT {
    VK_CONDITIONAL_RENDERING_INVERTED_BIT_EXT = 0x00000001,
} VkConditionalRenderingFlagBitsEXT;
```

## [](#_description)Description

- `VK_CONDITIONAL_RENDERING_INVERTED_BIT_EXT` specifies the condition used to determine whether to discard rendering commands or not. That is, if the 32-bit predicate read from `buffer` memory at `offset` is zero, the rendering commands are not discarded, and if non zero, then they are discarded.

## [](#_see_also)See Also

[VK\_EXT\_conditional\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_conditional_rendering.html), [VkConditionalRenderingFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConditionalRenderingFlagsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkConditionalRenderingFlagBitsEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0