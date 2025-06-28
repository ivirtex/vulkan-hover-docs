# VkMultiDrawInfoEXT(3) Manual Page

## Name

VkMultiDrawInfoEXT - Structure specifying a multi-draw command



## [](#_c_specification)C Specification

The `VkMultiDrawInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_multi_draw
typedef struct VkMultiDrawInfoEXT {
    uint32_t    firstVertex;
    uint32_t    vertexCount;
} VkMultiDrawInfoEXT;
```

## [](#_members)Members

- `firstVertex` is the first vertex to draw.
- `vertexCount` is the number of vertices to draw.

## [](#_description)Description

The members of `VkMultiDrawInfoEXT` have the same meaning as the `firstVertex` and `vertexCount` parameters in [vkCmdDraw](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDraw.html).

## [](#_see_also)See Also

[VK\_EXT\_multi\_draw](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_multi_draw.html), [vkCmdDrawMultiEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMultiEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMultiDrawInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0