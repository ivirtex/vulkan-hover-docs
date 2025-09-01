# VkPipelineVertexInputDivisorStateCreateInfo(3) Manual Page

## Name

VkPipelineVertexInputDivisorStateCreateInfo - Structure specifying vertex attributes assignment during instanced rendering



## [](#_c_specification)C Specification

If the [`vertexAttributeInstanceRateDivisor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-vertexAttributeInstanceRateDivisor) feature is enabled and the `pNext` chain of [VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputStateCreateInfo.html) includes a `VkPipelineVertexInputDivisorStateCreateInfo` structure, then that structure controls how vertex attributes are assigned to an instance when instanced rendering is enabled.

The `VkPipelineVertexInputDivisorStateCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPipelineVertexInputDivisorStateCreateInfo {
    VkStructureType                                  sType;
    const void*                                      pNext;
    uint32_t                                         vertexBindingDivisorCount;
    const VkVertexInputBindingDivisorDescription*    pVertexBindingDivisors;
} VkPipelineVertexInputDivisorStateCreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_vertex_attribute_divisor
typedef VkPipelineVertexInputDivisorStateCreateInfo VkPipelineVertexInputDivisorStateCreateInfoKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_vertex_attribute_divisor
typedef VkPipelineVertexInputDivisorStateCreateInfo VkPipelineVertexInputDivisorStateCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `vertexBindingDivisorCount` is the number of elements in the `pVertexBindingDivisors` array.
- `pVertexBindingDivisors` is a pointer to an array of [VkVertexInputBindingDivisorDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDivisorDescription.html) structures specifying the divisor value for each binding.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPipelineVertexInputDivisorStateCreateInfo-sType-sType)VUID-VkPipelineVertexInputDivisorStateCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_DIVISOR_STATE_CREATE_INFO`
- [](#VUID-VkPipelineVertexInputDivisorStateCreateInfo-pVertexBindingDivisors-parameter)VUID-VkPipelineVertexInputDivisorStateCreateInfo-pVertexBindingDivisors-parameter  
  `pVertexBindingDivisors` **must** be a valid pointer to an array of `vertexBindingDivisorCount` [VkVertexInputBindingDivisorDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDivisorDescription.html) structures
- [](#VUID-VkPipelineVertexInputDivisorStateCreateInfo-vertexBindingDivisorCount-arraylength)VUID-VkPipelineVertexInputDivisorStateCreateInfo-vertexBindingDivisorCount-arraylength  
  `vertexBindingDivisorCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_EXT\_vertex\_attribute\_divisor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_vertex_attribute_divisor.html), [VK\_KHR\_vertex\_attribute\_divisor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_vertex_attribute_divisor.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVertexInputBindingDivisorDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDivisorDescription.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineVertexInputDivisorStateCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0