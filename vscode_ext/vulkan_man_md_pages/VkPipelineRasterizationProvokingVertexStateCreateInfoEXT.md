# VkPipelineRasterizationProvokingVertexStateCreateInfoEXT(3) Manual Page

## Name

VkPipelineRasterizationProvokingVertexStateCreateInfoEXT - Structure
specifying provoking vertex mode used by a graphics pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

For a given primitive topology, the pipeline’s provoking vertex mode
determines which vertex is the provoking vertex. To specify the
provoking vertex mode, include a
`VkPipelineRasterizationProvokingVertexStateCreateInfoEXT` structure in
the
[VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)::`pNext`
chain when creating the pipeline.

The `VkPipelineRasterizationProvokingVertexStateCreateInfoEXT` structure
is defined as:

``` c
// Provided by VK_EXT_provoking_vertex
typedef struct VkPipelineRasterizationProvokingVertexStateCreateInfoEXT {
    VkStructureType             sType;
    const void*                 pNext;
    VkProvokingVertexModeEXT    provokingVertexMode;
} VkPipelineRasterizationProvokingVertexStateCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `provokingVertexMode` is a
  [VkProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkProvokingVertexModeEXT.html) value
  selecting the provoking vertex mode.

## <a href="#_description" class="anchor"></a>Description

If this struct is not provided when creating the pipeline, the pipeline
will use the `VK_PROVOKING_VERTEX_MODE_FIRST_VERTEX_EXT` mode.

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-provokingVertexModePerPipeline"
target="_blank"
rel="noopener"><code>provokingVertexModePerPipeline</code></a> limit is
`VK_FALSE`, then all pipelines bound within a render pass instance
**must** have the same `provokingVertexMode`.

Valid Usage

- <a
  href="#VUID-VkPipelineRasterizationProvokingVertexStateCreateInfoEXT-provokingVertexMode-04883"
  id="VUID-VkPipelineRasterizationProvokingVertexStateCreateInfoEXT-provokingVertexMode-04883"></a>
  VUID-VkPipelineRasterizationProvokingVertexStateCreateInfoEXT-provokingVertexMode-04883  
  If `provokingVertexMode` is
  `VK_PROVOKING_VERTEX_MODE_LAST_VERTEX_EXT`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-provokingVertexLast"
  target="_blank" rel="noopener"><code>provokingVertexLast</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineRasterizationProvokingVertexStateCreateInfoEXT-sType-sType"
  id="VUID-VkPipelineRasterizationProvokingVertexStateCreateInfoEXT-sType-sType"></a>
  VUID-VkPipelineRasterizationProvokingVertexStateCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_PROVOKING_VERTEX_STATE_CREATE_INFO_EXT`

- <a
  href="#VUID-VkPipelineRasterizationProvokingVertexStateCreateInfoEXT-provokingVertexMode-parameter"
  id="VUID-VkPipelineRasterizationProvokingVertexStateCreateInfoEXT-provokingVertexMode-parameter"></a>
  VUID-VkPipelineRasterizationProvokingVertexStateCreateInfoEXT-provokingVertexMode-parameter  
  `provokingVertexMode` **must** be a valid
  [VkProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkProvokingVertexModeEXT.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_provoking_vertex](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_provoking_vertex.html),
[VkProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkProvokingVertexModeEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineRasterizationProvokingVertexStateCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
