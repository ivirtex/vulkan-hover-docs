# VkConditionalRenderingBeginInfoEXT(3) Manual Page

## Name

VkConditionalRenderingBeginInfoEXT - Structure specifying conditional rendering begin information



## [](#_c_specification)C Specification

The `VkConditionalRenderingBeginInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_conditional_rendering
typedef struct VkConditionalRenderingBeginInfoEXT {
    VkStructureType                   sType;
    const void*                       pNext;
    VkBuffer                          buffer;
    VkDeviceSize                      offset;
    VkConditionalRenderingFlagsEXT    flags;
} VkConditionalRenderingBeginInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `buffer` is a buffer containing the predicate for conditional rendering.
- `offset` is the byte offset into `buffer` where the predicate is located.
- `flags` is a bitmask of [VkConditionalRenderingFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConditionalRenderingFlagsEXT.html) specifying the behavior of conditional rendering.

## [](#_description)Description

If the 32-bit value at `offset` in `buffer` memory is zero, then the rendering commands are discarded, otherwise they are executed as normal. If the value of the predicate in buffer memory changes while conditional rendering is active, the rendering commands **may** be discarded in an implementation-dependent way. Some implementations may latch the value of the predicate upon beginning conditional rendering while others may read it before every rendering command.

Valid Usage

- [](#VUID-VkConditionalRenderingBeginInfoEXT-buffer-01981)VUID-VkConditionalRenderingBeginInfoEXT-buffer-01981  
  If `buffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkConditionalRenderingBeginInfoEXT-buffer-01982)VUID-VkConditionalRenderingBeginInfoEXT-buffer-01982  
  `buffer` **must** have been created with the `VK_BUFFER_USAGE_CONDITIONAL_RENDERING_BIT_EXT` bit set
- [](#VUID-VkConditionalRenderingBeginInfoEXT-offset-01983)VUID-VkConditionalRenderingBeginInfoEXT-offset-01983  
  `offset` **must** be less than the size of `buffer` by at least 32 bits
- [](#VUID-VkConditionalRenderingBeginInfoEXT-offset-01984)VUID-VkConditionalRenderingBeginInfoEXT-offset-01984  
  `offset` **must** be a multiple of 4

Valid Usage (Implicit)

- [](#VUID-VkConditionalRenderingBeginInfoEXT-sType-sType)VUID-VkConditionalRenderingBeginInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CONDITIONAL_RENDERING_BEGIN_INFO_EXT`
- [](#VUID-VkConditionalRenderingBeginInfoEXT-pNext-pNext)VUID-VkConditionalRenderingBeginInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkConditionalRenderingBeginInfoEXT-buffer-parameter)VUID-VkConditionalRenderingBeginInfoEXT-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-VkConditionalRenderingBeginInfoEXT-flags-parameter)VUID-VkConditionalRenderingBeginInfoEXT-flags-parameter  
  `flags` **must** be a valid combination of [VkConditionalRenderingFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConditionalRenderingFlagBitsEXT.html) values

## [](#_see_also)See Also

[VK\_EXT\_conditional\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_conditional_rendering.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkConditionalRenderingFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConditionalRenderingFlagsEXT.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdBeginConditionalRenderingEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginConditionalRenderingEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkConditionalRenderingBeginInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0