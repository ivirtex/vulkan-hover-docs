# VkCommandBufferInheritanceConditionalRenderingInfoEXT(3) Manual Page

## Name

VkCommandBufferInheritanceConditionalRenderingInfoEXT - Structure specifying command buffer inheritance information



## [](#_c_specification)C Specification

If the `pNext` chain of [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html) includes a `VkCommandBufferInheritanceConditionalRenderingInfoEXT` structure, then that structure controls whether a command buffer **can** be executed while conditional rendering is [active](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#active-conditional-rendering) in the primary command buffer.

The `VkCommandBufferInheritanceConditionalRenderingInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_conditional_rendering
typedef struct VkCommandBufferInheritanceConditionalRenderingInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           conditionalRenderingEnable;
} VkCommandBufferInheritanceConditionalRenderingInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `conditionalRenderingEnable` specifies whether the command buffer **can** be executed while conditional rendering is active in the primary command buffer. If this is `VK_TRUE`, then this command buffer **can** be executed whether the primary command buffer has active conditional rendering or not. If this is `VK_FALSE`, then the primary command buffer **must** not have conditional rendering active.

## [](#_description)Description

If this structure is not present, the behavior is as if `conditionalRenderingEnable` is `VK_FALSE`.

Valid Usage

- [](#VUID-VkCommandBufferInheritanceConditionalRenderingInfoEXT-conditionalRenderingEnable-01977)VUID-VkCommandBufferInheritanceConditionalRenderingInfoEXT-conditionalRenderingEnable-01977  
  If the [`inheritedConditionalRendering`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-inheritedConditionalRendering) feature is not enabled, `conditionalRenderingEnable` **must** be `VK_FALSE`

Valid Usage (Implicit)

- [](#VUID-VkCommandBufferInheritanceConditionalRenderingInfoEXT-sType-sType)VUID-VkCommandBufferInheritanceConditionalRenderingInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_CONDITIONAL_RENDERING_INFO_EXT`

## [](#_see_also)See Also

[VK\_EXT\_conditional\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_conditional_rendering.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCommandBufferInheritanceConditionalRenderingInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0