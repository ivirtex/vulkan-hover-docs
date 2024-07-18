# VkDrawIndirectCommand(3) Manual Page

## Name

VkDrawIndirectCommand - Structure specifying an indirect drawing command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDrawIndirectCommand` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkDrawIndirectCommand {
    uint32_t    vertexCount;
    uint32_t    instanceCount;
    uint32_t    firstVertex;
    uint32_t    firstInstance;
} VkDrawIndirectCommand;
```

## <a href="#_members" class="anchor"></a>Members

- `vertexCount` is the number of vertices to draw.

- `instanceCount` is the number of instances to draw.

- `firstVertex` is the index of the first vertex to draw.

- `firstInstance` is the instance ID of the first instance to draw.

## <a href="#_description" class="anchor"></a>Description

The members of `VkDrawIndirectCommand` have the same meaning as the
similarly named parameters of [vkCmdDraw](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDraw.html).

Valid Usage

- <a href="#VUID-VkDrawIndirectCommand-pNext-09461"
  id="VUID-VkDrawIndirectCommand-pNext-09461"></a>
  VUID-VkDrawIndirectCommand-pNext-09461  
  If the bound graphics pipeline state was created with
  [VkPipelineVertexInputDivisorStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputDivisorStateCreateInfoKHR.html)
  in the `pNext` chain of
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)::`pVertexInputState`,
  any member of
  [VkPipelineVertexInputDivisorStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputDivisorStateCreateInfoKHR.html)::`pVertexBindingDivisors`
  has a value other than `1` in `divisor`, and
  [VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR.html)::`supportsNonZeroFirstInstance`
  is `VK_FALSE`, then `firstInstance` **must** be `0`

- <a href="#VUID-VkDrawIndirectCommand-None-09462"
  id="VUID-VkDrawIndirectCommand-None-09462"></a>
  VUID-VkDrawIndirectCommand-None-09462  
  If [shader objects](#shaders-objects) are used for drawing or the
  bound graphics pipeline state was created with the
  `VK_DYNAMIC_STATE_VERTEX_INPUT_EXT` dynamic state enabled, any member
  of the `pVertexBindingDescriptions` parameter to the
  [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetVertexInputEXT.html) call that sets
  this dynamic state has a value other than `1` in `divisor`, and
  [VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR.html)::`supportsNonZeroFirstInstance`
  is `VK_FALSE`, then `firstInstance` **must** be `0`

<!-- -->

- <a href="#VUID-VkDrawIndirectCommand-None-00500"
  id="VUID-VkDrawIndirectCommand-None-00500"></a>
  VUID-VkDrawIndirectCommand-None-00500  
  For a given vertex buffer binding, any attribute data fetched **must**
  be entirely contained within the corresponding vertex buffer binding,
  as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fxvertex-input"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fxvertex-input</a>

- <a href="#VUID-VkDrawIndirectCommand-firstInstance-00501"
  id="VUID-VkDrawIndirectCommand-firstInstance-00501"></a>
  VUID-VkDrawIndirectCommand-firstInstance-00501  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-drawIndirectFirstInstance"
  target="_blank"
  rel="noopener"><code>drawIndirectFirstInstance</code></a> feature is
  not enabled, `firstInstance` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[vkCmdDrawIndirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndirect.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDrawIndirectCommand"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
