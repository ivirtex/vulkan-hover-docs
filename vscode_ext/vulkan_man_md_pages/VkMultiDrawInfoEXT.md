# VkMultiDrawInfoEXT(3) Manual Page

## Name

VkMultiDrawInfoEXT - Structure specifying a multi-draw command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMultiDrawInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_multi_draw
typedef struct VkMultiDrawInfoEXT {
    uint32_t    firstVertex;
    uint32_t    vertexCount;
} VkMultiDrawInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `firstVertex` is the first vertex to draw.

- `vertexCount` is the number of vertices to draw.

## <a href="#_description" class="anchor"></a>Description

The members of `VkMultiDrawInfoEXT` have the same meaning as the
`firstVertex` and `vertexCount` parameters in
[vkCmdDraw](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDraw.html).

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_multi_draw](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_multi_draw.html),
[vkCmdDrawMultiEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawMultiEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMultiDrawInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
