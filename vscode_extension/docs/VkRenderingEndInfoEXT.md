# VkRenderingEndInfoEXT(3) Manual Page

## Name

VkRenderingEndInfoEXT - Structure specifying render pass end information



## [](#_c_specification)C Specification

The `VkRenderingEndInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_fragment_density_map_offset
typedef struct VkRenderingEndInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
} VkRenderingEndInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkRenderingEndInfoEXT-sType-sType)VUID-VkRenderingEndInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDERING_END_INFO_EXT`
- [](#VUID-VkRenderingEndInfoEXT-pNext-pNext)VUID-VkRenderingEndInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkRenderPassFragmentDensityMapOffsetEndInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassFragmentDensityMapOffsetEndInfoEXT.html)
- [](#VUID-VkRenderingEndInfoEXT-sType-unique)VUID-VkRenderingEndInfoEXT-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_EXT\_fragment\_density\_map\_offset](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_fragment_density_map_offset.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdEndRendering2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndRendering2EXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderingEndInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0