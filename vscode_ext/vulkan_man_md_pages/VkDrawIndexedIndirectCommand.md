# VkDrawIndexedIndirectCommand(3) Manual Page

## Name

VkDrawIndexedIndirectCommand - Structure specifying an indexed indirect
drawing command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDrawIndexedIndirectCommand` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkDrawIndexedIndirectCommand {
    uint32_t    indexCount;
    uint32_t    instanceCount;
    uint32_t    firstIndex;
    int32_t     vertexOffset;
    uint32_t    firstInstance;
} VkDrawIndexedIndirectCommand;
```

## <a href="#_members" class="anchor"></a>Members

- `indexCount` is the number of vertices to draw.

- `instanceCount` is the number of instances to draw.

- `firstIndex` is the base index within the index buffer.

- `vertexOffset` is the value added to the vertex index before indexing
  into the vertex buffer.

- `firstInstance` is the instance ID of the first instance to draw.

## <a href="#_description" class="anchor"></a>Description

The members of `VkDrawIndexedIndirectCommand` have the same meaning as
the similarly named parameters of
[vkCmdDrawIndexed](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexed.html).

Valid Usage

- <a href="#VUID-VkDrawIndexedIndirectCommand-pNext-09461"
  id="VUID-VkDrawIndexedIndirectCommand-pNext-09461"></a>
  VUID-VkDrawIndexedIndirectCommand-pNext-09461  
  If the bound graphics pipeline state was created with
  [VkPipelineVertexInputDivisorStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputDivisorStateCreateInfoKHR.html)
  in the `pNext` chain of
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)::`pVertexInputState`,
  any member of
  [VkPipelineVertexInputDivisorStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputDivisorStateCreateInfoKHR.html)::`pVertexBindingDivisors`
  has a value other than `1` in `divisor`, and
  [VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR.html)::`supportsNonZeroFirstInstance`
  is `VK_FALSE`, then `firstInstance` **must** be `0`

- <a href="#VUID-VkDrawIndexedIndirectCommand-None-09462"
  id="VUID-VkDrawIndexedIndirectCommand-None-09462"></a>
  VUID-VkDrawIndexedIndirectCommand-None-09462  
  If [shader objects](#shaders-objects) are used for drawing or the
  bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_VERTEX_INPUT_EXT` dynamic state enabled, any member
  of the `pVertexBindingDescriptions` parameter to the
  [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetVertexInputEXT.html) call that sets
  this dynamic state has a value other than `1` in `divisor`, and
  [VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR.html)::`supportsNonZeroFirstInstance`
  is `VK_FALSE`, then `firstInstance` **must** be `0`

<!-- -->

- <a href="#VUID-VkDrawIndexedIndirectCommand-robustBufferAccess2-08798"
  id="VUID-VkDrawIndexedIndirectCommand-robustBufferAccess2-08798"></a>
  VUID-VkDrawIndexedIndirectCommand-robustBufferAccess2-08798  
  If [`robustBufferAccess2`](#features-robustBufferAccess2) is not
  enabled, (`indexSize` × (`firstIndex` + `indexCount`) + `offset`)
  **must** be less than or equal to the size of the bound index buffer,
  with `indexSize` being based on the type specified by `indexType`,
  where the index buffer, `indexType`, and `offset` are specified via
  `vkCmdBindIndexBuffer` or `vkCmdBindIndexBuffer2KHR`. If
  `vkCmdBindIndexBuffer2KHR` is used to bind the index buffer, the size
  of the bound index buffer is
  [vkCmdBindIndexBuffer2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer2KHR.html)::`size`

- <a href="#VUID-VkDrawIndexedIndirectCommand-None-00552"
  id="VUID-VkDrawIndexedIndirectCommand-None-00552"></a>
  VUID-VkDrawIndexedIndirectCommand-None-00552  
  For a given vertex buffer binding, any attribute data fetched **must**
  be entirely contained within the corresponding vertex buffer binding,
  as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fxvertex-input"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fxvertex-input</a>

- <a href="#VUID-VkDrawIndexedIndirectCommand-firstInstance-00554"
  id="VUID-VkDrawIndexedIndirectCommand-firstInstance-00554"></a>
  VUID-VkDrawIndexedIndirectCommand-firstInstance-00554  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-drawIndirectFirstInstance"
  target="_blank"
  rel="noopener"><code>drawIndirectFirstInstance</code></a> feature is
  not enabled, `firstInstance` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[vkCmdDrawIndexedIndirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexedIndirect.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDrawIndexedIndirectCommand"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
