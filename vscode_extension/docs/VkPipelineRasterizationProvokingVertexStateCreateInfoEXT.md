# VkPipelineRasterizationProvokingVertexStateCreateInfoEXT(3) Manual Page

## Name

VkPipelineRasterizationProvokingVertexStateCreateInfoEXT - Structure specifying provoking vertex mode used by a graphics pipeline



## [](#_c_specification)C Specification

For a given primitive topology, the pipeline’s provoking vertex mode determines which vertex is the provoking vertex. To specify the provoking vertex mode, include a `VkPipelineRasterizationProvokingVertexStateCreateInfoEXT` structure in the [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html)::`pNext` chain when creating the pipeline.

The `VkPipelineRasterizationProvokingVertexStateCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_provoking_vertex
typedef struct VkPipelineRasterizationProvokingVertexStateCreateInfoEXT {
    VkStructureType             sType;
    const void*                 pNext;
    VkProvokingVertexModeEXT    provokingVertexMode;
} VkPipelineRasterizationProvokingVertexStateCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `provokingVertexMode` is a [VkProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkProvokingVertexModeEXT.html) value selecting the provoking vertex mode.

## [](#_description)Description

If this structure is not provided when creating the pipeline, the pipeline will use the `VK_PROVOKING_VERTEX_MODE_FIRST_VERTEX_EXT` mode.

If the [`provokingVertexModePerPipeline`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-provokingVertexModePerPipeline) limit is `VK_FALSE`, then all pipelines bound within a render pass instance **must** have the same `provokingVertexMode`.

Valid Usage

- [](#VUID-VkPipelineRasterizationProvokingVertexStateCreateInfoEXT-provokingVertexMode-04883)VUID-VkPipelineRasterizationProvokingVertexStateCreateInfoEXT-provokingVertexMode-04883  
  If `provokingVertexMode` is `VK_PROVOKING_VERTEX_MODE_LAST_VERTEX_EXT`, then the [`provokingVertexLast`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-provokingVertexLast) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-VkPipelineRasterizationProvokingVertexStateCreateInfoEXT-sType-sType)VUID-VkPipelineRasterizationProvokingVertexStateCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_PROVOKING_VERTEX_STATE_CREATE_INFO_EXT`
- [](#VUID-VkPipelineRasterizationProvokingVertexStateCreateInfoEXT-provokingVertexMode-parameter)VUID-VkPipelineRasterizationProvokingVertexStateCreateInfoEXT-provokingVertexMode-parameter  
  `provokingVertexMode` **must** be a valid [VkProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkProvokingVertexModeEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_provoking\_vertex](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_provoking_vertex.html), [VkProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkProvokingVertexModeEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineRasterizationProvokingVertexStateCreateInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0