# VkProvokingVertexModeEXT(3) Manual Page

## Name

VkProvokingVertexModeEXT - Specify which vertex in a primitive is the provoking vertex



## [](#_c_specification)C Specification

Possible values of [VkPipelineRasterizationProvokingVertexStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationProvokingVertexStateCreateInfoEXT.html)::`provokingVertexMode` are:

```c++
// Provided by VK_EXT_provoking_vertex
typedef enum VkProvokingVertexModeEXT {
    VK_PROVOKING_VERTEX_MODE_FIRST_VERTEX_EXT = 0,
    VK_PROVOKING_VERTEX_MODE_LAST_VERTEX_EXT = 1,
} VkProvokingVertexModeEXT;
```

## [](#_description)Description

- `VK_PROVOKING_VERTEX_MODE_FIRST_VERTEX_EXT` specifies that the provoking vertex is the first non-adjacency vertex in the list of vertices used by a primitive.
- `VK_PROVOKING_VERTEX_MODE_LAST_VERTEX_EXT` specifies that the provoking vertex is the last non-adjacency vertex in the list of vertices used by a primitive.

These modes are described more precisely in [Primitive Topologies](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-primitive-topologies).

## [](#_see_also)See Also

[VK\_EXT\_provoking\_vertex](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_provoking_vertex.html), [VkPipelineRasterizationProvokingVertexStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationProvokingVertexStateCreateInfoEXT.html), [vkCmdSetProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetProvokingVertexModeEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkProvokingVertexModeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0