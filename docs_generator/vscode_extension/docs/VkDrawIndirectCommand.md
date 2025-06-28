# VkDrawIndirectCommand(3) Manual Page

## Name

VkDrawIndirectCommand - Structure specifying an indirect drawing command



## [](#_c_specification)C Specification

The `VkDrawIndirectCommand` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkDrawIndirectCommand {
    uint32_t    vertexCount;
    uint32_t    instanceCount;
    uint32_t    firstVertex;
    uint32_t    firstInstance;
} VkDrawIndirectCommand;
```

## [](#_members)Members

- `vertexCount` is the number of vertices to draw.
- `instanceCount` is the number of instances to draw.
- `firstVertex` is the index of the first vertex to draw.
- `firstInstance` is the instance ID of the first instance to draw.

## [](#_description)Description

The members of `VkDrawIndirectCommand` have the same meaning as the similarly named parameters of [vkCmdDraw](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDraw.html).

Valid Usage

- [](#VUID-VkDrawIndirectCommand-pNext-09461)VUID-VkDrawIndirectCommand-pNext-09461  
  If the bound graphics pipeline state was created with [VkPipelineVertexInputDivisorStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputDivisorStateCreateInfo.html) in the `pNext` chain of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html)::`pVertexInputState`, any member of [VkPipelineVertexInputDivisorStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputDivisorStateCreateInfo.html)::`pVertexBindingDivisors` has a value other than `1` in `divisor`, and [VkPhysicalDeviceVertexAttributeDivisorProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVertexAttributeDivisorProperties.html)::`supportsNonZeroFirstInstance` is `VK_FALSE`, then `firstInstance` **must** be `0`
- [](#VUID-VkDrawIndirectCommand-None-09462)VUID-VkDrawIndirectCommand-None-09462  
  If [shader objects](#shaders-objects) are used for drawing or the bound graphics pipeline state was created with the `VK_DYNAMIC_STATE_VERTEX_INPUT_EXT` dynamic state enabled, any member of the `pVertexBindingDescriptions` parameter to the [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetVertexInputEXT.html) call that sets this dynamic state has a value other than `1` in `divisor`, and [VkPhysicalDeviceVertexAttributeDivisorProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVertexAttributeDivisorProperties.html)::`supportsNonZeroFirstInstance` is `VK_FALSE`, then `firstInstance` **must** be `0`

<!--THE END-->

- [](#VUID-VkDrawIndirectCommand-firstInstance-00501)VUID-VkDrawIndirectCommand-firstInstance-00501  
  If the [`drawIndirectFirstInstance`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-drawIndirectFirstInstance) feature is not enabled, `firstInstance` **must** be `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [vkCmdDrawIndirect](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndirect.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDrawIndirectCommand)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0