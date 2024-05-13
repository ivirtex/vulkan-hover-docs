# VkConditionalRenderingFlagBitsEXT(3) Manual Page

## Name

VkConditionalRenderingFlagBitsEXT - Specify the behavior of conditional
rendering



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[vkCmdBeginConditionalRenderingEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginConditionalRenderingEXT.html)::`flags`,
specifying the behavior of conditional rendering, are:

``` c
// Provided by VK_EXT_conditional_rendering
typedef enum VkConditionalRenderingFlagBitsEXT {
    VK_CONDITIONAL_RENDERING_INVERTED_BIT_EXT = 0x00000001,
} VkConditionalRenderingFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_CONDITIONAL_RENDERING_INVERTED_BIT_EXT` specifies the condition
  used to determine whether to discard rendering commands or not. That
  is, if the 32-bit predicate read from `buffer` memory at `offset` is
  zero, the rendering commands are not discarded, and if non zero, then
  they are discarded.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_conditional_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_conditional_rendering.html),
[VkConditionalRenderingFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConditionalRenderingFlagsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkConditionalRenderingFlagBitsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
