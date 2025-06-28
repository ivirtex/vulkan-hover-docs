# VkMultiDrawIndexedInfoEXT(3) Manual Page

## Name

VkMultiDrawIndexedInfoEXT - Structure specifying a multi-draw command



## [](#_c_specification)C Specification

The `VkMultiDrawIndexedInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_multi_draw
typedef struct VkMultiDrawIndexedInfoEXT {
    uint32_t    firstIndex;
    uint32_t    indexCount;
    int32_t     vertexOffset;
} VkMultiDrawIndexedInfoEXT;
```

## [](#_members)Members

- `firstIndex` is the first index to draw.
- `indexCount` is the number of vertices to draw.
- `vertexOffset` is the value added to the vertex index before indexing into the vertex buffer for indexed multidraws.

## [](#_description)Description

The `firstIndex`, `indexCount`, and `vertexOffset` members of `VkMultiDrawIndexedInfoEXT` have the same meaning as the `firstIndex`, `indexCount`, and `vertexOffset` parameters, respectively, of [vkCmdDrawIndexed](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawIndexed.html).

## [](#_see_also)See Also

[VK\_EXT\_multi\_draw](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_multi_draw.html), [vkCmdDrawMultiIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMultiIndexedEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMultiDrawIndexedInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0