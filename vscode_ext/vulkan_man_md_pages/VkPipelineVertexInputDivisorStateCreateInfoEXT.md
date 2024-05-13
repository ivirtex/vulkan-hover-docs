# VkPipelineVertexInputDivisorStateCreateInfoKHR(3) Manual Page

## Name

VkPipelineVertexInputDivisorStateCreateInfoKHR - Structure specifying
vertex attributes assignment during instanced rendering



## <a href="#_c_specification" class="anchor"></a>C Specification

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-vertexAttributeInstanceRateDivisor"
target="_blank"
rel="noopener"><code>vertexAttributeInstanceRateDivisor</code></a>
feature is enabled and the `pNext` chain of
[VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputStateCreateInfo.html)
includes a `VkPipelineVertexInputDivisorStateCreateInfoKHR` structure,
then that structure controls how vertex attributes are assigned to an
instance when instanced rendering is enabled.

The `VkPipelineVertexInputDivisorStateCreateInfoKHR` structure is
defined as:

``` c
// Provided by VK_KHR_vertex_attribute_divisor
typedef struct VkPipelineVertexInputDivisorStateCreateInfoKHR {
    VkStructureType                                     sType;
    const void*                                         pNext;
    uint32_t                                            vertexBindingDivisorCount;
    const VkVertexInputBindingDivisorDescriptionKHR*    pVertexBindingDivisors;
} VkPipelineVertexInputDivisorStateCreateInfoKHR;
```

or the equivalent

``` c
// Provided by VK_EXT_vertex_attribute_divisor
typedef VkPipelineVertexInputDivisorStateCreateInfoKHR VkPipelineVertexInputDivisorStateCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `vertexBindingDivisorCount` is the number of elements in the
  `pVertexBindingDivisors` array.

- `pVertexBindingDivisors` is a pointer to an array of
  [VkVertexInputBindingDivisorDescriptionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDivisorDescriptionKHR.html)
  structures specifying the divisor value for each binding.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineVertexInputDivisorStateCreateInfoKHR-sType-sType"
  id="VUID-VkPipelineVertexInputDivisorStateCreateInfoKHR-sType-sType"></a>
  VUID-VkPipelineVertexInputDivisorStateCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_DIVISOR_STATE_CREATE_INFO_KHR`

- <a
  href="#VUID-VkPipelineVertexInputDivisorStateCreateInfoKHR-pVertexBindingDivisors-parameter"
  id="VUID-VkPipelineVertexInputDivisorStateCreateInfoKHR-pVertexBindingDivisors-parameter"></a>
  VUID-VkPipelineVertexInputDivisorStateCreateInfoKHR-pVertexBindingDivisors-parameter  
  `pVertexBindingDivisors` **must** be a valid pointer to an array of
  `vertexBindingDivisorCount`
  [VkVertexInputBindingDivisorDescriptionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDivisorDescriptionKHR.html)
  structures

- <a
  href="#VUID-VkPipelineVertexInputDivisorStateCreateInfoKHR-vertexBindingDivisorCount-arraylength"
  id="VUID-VkPipelineVertexInputDivisorStateCreateInfoKHR-vertexBindingDivisorCount-arraylength"></a>
  VUID-VkPipelineVertexInputDivisorStateCreateInfoKHR-vertexBindingDivisorCount-arraylength  
  `vertexBindingDivisorCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_vertex_attribute_divisor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_vertex_attribute_divisor.html),
[VK_KHR_vertex_attribute_divisor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_vertex_attribute_divisor.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVertexInputBindingDivisorDescriptionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDivisorDescriptionKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineVertexInputDivisorStateCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
