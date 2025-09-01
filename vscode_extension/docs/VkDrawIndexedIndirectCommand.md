# VkDrawIndexedIndirectCommand(3) Manual Page

## Name

VkDrawIndexedIndirectCommand - Structure specifying an indexed indirect drawing command



## [](#_c_specification)C Specification

The `VkDrawIndexedIndirectCommand` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkDrawIndexedIndirectCommand {
    uint32_t    indexCount;
    uint32_t    instanceCount;
    uint32_t    firstIndex;
    int32_t     vertexOffset;
    uint32_t    firstInstance;
} VkDrawIndexedIndirectCommand;
```

## [](#_members)Members

- `indexCount` is the number of vertices to draw.
- `instanceCount` is the number of instances to draw.
- `firstIndex` is the base index within the index buffer.
- `vertexOffset` is the value added to the vertex index before indexing into the vertex buffer.
- `firstInstance` is the instance ID of the first instance to draw.

## [](#_description)Description

The members of `VkDrawIndexedIndirectCommand` have the same meaning as the similarly named parameters of [vkCmdDrawIndexed](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndexed.html).

Valid Usage

- [](#VUID-VkDrawIndexedIndirectCommand-pNext-09461)VUID-VkDrawIndexedIndirectCommand-pNext-09461  
  If the bound graphics pipeline state was created with [VkPipelineVertexInputDivisorStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputDivisorStateCreateInfo.html) in the `pNext` chain of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html)::`pVertexInputState`, any member of [VkPipelineVertexInputDivisorStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputDivisorStateCreateInfo.html)::`pVertexBindingDivisors` has a value other than `1` in `divisor`, and [VkPhysicalDeviceVertexAttributeDivisorProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVertexAttributeDivisorProperties.html)::`supportsNonZeroFirstInstance` is `VK_FALSE`, then `firstInstance` **must** be `0`
- [](#VUID-VkDrawIndexedIndirectCommand-None-09462)VUID-VkDrawIndexedIndirectCommand-None-09462  
  If [shader objects](#shaders-objects) are used for drawing or the bound graphics pipeline state was created with the `VK_DYNAMIC_STATE_VERTEX_INPUT_EXT` dynamic state enabled, any member of the `pVertexBindingDescriptions` parameter to the [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetVertexInputEXT.html) call that sets this dynamic state has a value other than `1` in `divisor`, and [VkPhysicalDeviceVertexAttributeDivisorProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVertexAttributeDivisorProperties.html)::`supportsNonZeroFirstInstance` is `VK_FALSE`, then `firstInstance` **must** be `0`

<!--THE END-->

- [](#VUID-VkDrawIndexedIndirectCommand-robustBufferAccess2-08798)VUID-VkDrawIndexedIndirectCommand-robustBufferAccess2-08798  
  If the [`robustBufferAccess2`](#features-robustBufferAccess2) feature is not enabled, (`indexSize` × (`firstIndex` + `indexCount`) + `offset`) **must** be less than or equal to the size of the bound index buffer, with `indexSize` being based on the type specified by `indexType`, where the index buffer, `indexType`, and `offset` are specified via `vkCmdBindIndexBuffer` or `vkCmdBindIndexBuffer2`. If `vkCmdBindIndexBuffer2` is used to bind the index buffer, the size of the bound index buffer is [vkCmdBindIndexBuffer2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindIndexBuffer2.html)::`size`
- [](#VUID-VkDrawIndexedIndirectCommand-firstInstance-00554)VUID-VkDrawIndexedIndirectCommand-firstInstance-00554  
  If the [`drawIndirectFirstInstance`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-drawIndirectFirstInstance) feature is not enabled, `firstInstance` **must** be `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [vkCmdDrawIndexedIndirect](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndexedIndirect.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDrawIndexedIndirectCommand).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0