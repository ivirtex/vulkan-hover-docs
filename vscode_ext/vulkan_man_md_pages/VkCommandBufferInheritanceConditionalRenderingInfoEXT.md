# VkCommandBufferInheritanceConditionalRenderingInfoEXT(3) Manual Page

## Name

VkCommandBufferInheritanceConditionalRenderingInfoEXT - Structure
specifying command buffer inheritance information



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of
[VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)
includes a `VkCommandBufferInheritanceConditionalRenderingInfoEXT`
structure, then that structure controls whether a command buffer **can**
be executed while conditional rendering is <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#active-conditional-rendering"
target="_blank" rel="noopener">active</a> in the primary command buffer.

The `VkCommandBufferInheritanceConditionalRenderingInfoEXT` structure is
defined as:

``` c
// Provided by VK_EXT_conditional_rendering
typedef struct VkCommandBufferInheritanceConditionalRenderingInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           conditionalRenderingEnable;
} VkCommandBufferInheritanceConditionalRenderingInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `conditionalRenderingEnable` specifies whether the command buffer
  **can** be executed while conditional rendering is active in the
  primary command buffer. If this is `VK_TRUE`, then this command buffer
  **can** be executed whether the primary command buffer has active
  conditional rendering or not. If this is `VK_FALSE`, then the primary
  command buffer **must** not have conditional rendering active.

## <a href="#_description" class="anchor"></a>Description

If this structure is not present, the behavior is as if
`conditionalRenderingEnable` is `VK_FALSE`.

Valid Usage

- <a
  href="#VUID-VkCommandBufferInheritanceConditionalRenderingInfoEXT-conditionalRenderingEnable-01977"
  id="VUID-VkCommandBufferInheritanceConditionalRenderingInfoEXT-conditionalRenderingEnable-01977"></a>
  VUID-VkCommandBufferInheritanceConditionalRenderingInfoEXT-conditionalRenderingEnable-01977  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-inheritedConditionalRendering"
  target="_blank"
  rel="noopener"><code>inheritedConditionalRendering</code></a> feature
  is not enabled, `conditionalRenderingEnable` **must** be `VK_FALSE`

Valid Usage (Implicit)

- <a
  href="#VUID-VkCommandBufferInheritanceConditionalRenderingInfoEXT-sType-sType"
  id="VUID-VkCommandBufferInheritanceConditionalRenderingInfoEXT-sType-sType"></a>
  VUID-VkCommandBufferInheritanceConditionalRenderingInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_CONDITIONAL_RENDERING_INFO_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_conditional_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_conditional_rendering.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCommandBufferInheritanceConditionalRenderingInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
