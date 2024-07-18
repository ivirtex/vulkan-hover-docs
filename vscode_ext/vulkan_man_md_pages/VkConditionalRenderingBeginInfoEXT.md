# VkConditionalRenderingBeginInfoEXT(3) Manual Page

## Name

VkConditionalRenderingBeginInfoEXT - Structure specifying conditional
rendering begin information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkConditionalRenderingBeginInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_conditional_rendering
typedef struct VkConditionalRenderingBeginInfoEXT {
    VkStructureType                   sType;
    const void*                       pNext;
    VkBuffer                          buffer;
    VkDeviceSize                      offset;
    VkConditionalRenderingFlagsEXT    flags;
} VkConditionalRenderingBeginInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `buffer` is a buffer containing the predicate for conditional
  rendering.

- `offset` is the byte offset into `buffer` where the predicate is
  located.

- `flags` is a bitmask of
  [VkConditionalRenderingFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConditionalRenderingFlagsEXT.html)
  specifying the behavior of conditional rendering.

## <a href="#_description" class="anchor"></a>Description

If the 32-bit value at `offset` in `buffer` memory is zero, then the
rendering commands are discarded, otherwise they are executed as normal.
If the value of the predicate in buffer memory changes while conditional
rendering is active, the rendering commands **may** be discarded in an
implementation-dependent way. Some implementations may latch the value
of the predicate upon beginning conditional rendering while others may
read it before every rendering command.

Valid Usage

- <a href="#VUID-VkConditionalRenderingBeginInfoEXT-buffer-01981"
  id="VUID-VkConditionalRenderingBeginInfoEXT-buffer-01981"></a>
  VUID-VkConditionalRenderingBeginInfoEXT-buffer-01981  
  If `buffer` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-VkConditionalRenderingBeginInfoEXT-buffer-01982"
  id="VUID-VkConditionalRenderingBeginInfoEXT-buffer-01982"></a>
  VUID-VkConditionalRenderingBeginInfoEXT-buffer-01982  
  `buffer` **must** have been created with the
  `VK_BUFFER_USAGE_CONDITIONAL_RENDERING_BIT_EXT` bit set

- <a href="#VUID-VkConditionalRenderingBeginInfoEXT-offset-01983"
  id="VUID-VkConditionalRenderingBeginInfoEXT-offset-01983"></a>
  VUID-VkConditionalRenderingBeginInfoEXT-offset-01983  
  `offset` **must** be less than the size of `buffer` by at least 32
  bits

- <a href="#VUID-VkConditionalRenderingBeginInfoEXT-offset-01984"
  id="VUID-VkConditionalRenderingBeginInfoEXT-offset-01984"></a>
  VUID-VkConditionalRenderingBeginInfoEXT-offset-01984  
  `offset` **must** be a multiple of 4

Valid Usage (Implicit)

- <a href="#VUID-VkConditionalRenderingBeginInfoEXT-sType-sType"
  id="VUID-VkConditionalRenderingBeginInfoEXT-sType-sType"></a>
  VUID-VkConditionalRenderingBeginInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_CONDITIONAL_RENDERING_BEGIN_INFO_EXT`

- <a href="#VUID-VkConditionalRenderingBeginInfoEXT-pNext-pNext"
  id="VUID-VkConditionalRenderingBeginInfoEXT-pNext-pNext"></a>
  VUID-VkConditionalRenderingBeginInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkConditionalRenderingBeginInfoEXT-buffer-parameter"
  id="VUID-VkConditionalRenderingBeginInfoEXT-buffer-parameter"></a>
  VUID-VkConditionalRenderingBeginInfoEXT-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkConditionalRenderingBeginInfoEXT-flags-parameter"
  id="VUID-VkConditionalRenderingBeginInfoEXT-flags-parameter"></a>
  VUID-VkConditionalRenderingBeginInfoEXT-flags-parameter  
  `flags` **must** be a valid combination of
  [VkConditionalRenderingFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConditionalRenderingFlagBitsEXT.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_conditional_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_conditional_rendering.html),
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkConditionalRenderingFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConditionalRenderingFlagsEXT.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdBeginConditionalRenderingEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginConditionalRenderingEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkConditionalRenderingBeginInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
