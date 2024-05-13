# VkMultiDrawIndexedInfoEXT(3) Manual Page

## Name

VkMultiDrawIndexedInfoEXT - Structure specifying a multi-draw command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMultiDrawIndexedInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_multi_draw
typedef struct VkMultiDrawIndexedInfoEXT {
    uint32_t    firstIndex;
    uint32_t    indexCount;
    int32_t     vertexOffset;
} VkMultiDrawIndexedInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `firstIndex` is the first index to draw.

- `indexCount` is the number of vertices to draw.

- `vertexOffset` is the value added to the vertex index before indexing
  into the vertex buffer for indexed multidraws.

## <a href="#_description" class="anchor"></a>Description

The `firstIndex`, `indexCount`, and `vertexOffset` members of
`VkMultiDrawIndexedInfoEXT` have the same meaning as the `firstIndex`,
`indexCount`, and `vertexOffset` parameters, respectively, of
[vkCmdDrawIndexed](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexed.html).

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_multi_draw](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_multi_draw.html),
[vkCmdDrawMultiIndexedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMultiIndexedEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMultiDrawIndexedInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
