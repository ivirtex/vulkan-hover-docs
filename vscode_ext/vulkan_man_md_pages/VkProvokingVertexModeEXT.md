# VkProvokingVertexModeEXT(3) Manual Page

## Name

VkProvokingVertexModeEXT - Specify which vertex in a primitive is the
provoking vertex



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkPipelineRasterizationProvokingVertexStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationProvokingVertexStateCreateInfoEXT.html)::`provokingVertexMode`
are:

``` c
// Provided by VK_EXT_provoking_vertex
typedef enum VkProvokingVertexModeEXT {
    VK_PROVOKING_VERTEX_MODE_FIRST_VERTEX_EXT = 0,
    VK_PROVOKING_VERTEX_MODE_LAST_VERTEX_EXT = 1,
} VkProvokingVertexModeEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PROVOKING_VERTEX_MODE_FIRST_VERTEX_EXT` specifies that the
  provoking vertex is the first non-adjacency vertex in the list of
  vertices used by a primitive.

- `VK_PROVOKING_VERTEX_MODE_LAST_VERTEX_EXT` specifies that the
  provoking vertex is the last non-adjacency vertex in the list of
  vertices used by a primitive.

These modes are described more precisely in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-primitive-topologies"
target="_blank" rel="noopener">Primitive Topologies</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_provoking_vertex](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_provoking_vertex.html),
[VkPipelineRasterizationProvokingVertexStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationProvokingVertexStateCreateInfoEXT.html),
[vkCmdSetProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetProvokingVertexModeEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkProvokingVertexModeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
