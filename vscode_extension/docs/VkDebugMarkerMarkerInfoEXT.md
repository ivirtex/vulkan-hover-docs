# VkDebugMarkerMarkerInfoEXT(3) Manual Page

## Name

VkDebugMarkerMarkerInfoEXT - Specify parameters of a command buffer marker region



## [](#_c_specification)C Specification

The `VkDebugMarkerMarkerInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_debug_marker
typedef struct VkDebugMarkerMarkerInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    const char*        pMarkerName;
    float              color[4];
} VkDebugMarkerMarkerInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pMarkerName` is a pointer to a null-terminated UTF-8 string containing the name of the marker.
- `color` is an **optional** RGBA color value that can be associated with the marker. A particular implementation **may** choose to ignore this color value. The values contain RGBA values in order, in the range 0.0 to 1.0. If all elements in `color` are 0.0, then it is ignored.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDebugMarkerMarkerInfoEXT-sType-sType)VUID-VkDebugMarkerMarkerInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEBUG_MARKER_MARKER_INFO_EXT`
- [](#VUID-VkDebugMarkerMarkerInfoEXT-pNext-pNext)VUID-VkDebugMarkerMarkerInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDebugMarkerMarkerInfoEXT-pMarkerName-parameter)VUID-VkDebugMarkerMarkerInfoEXT-pMarkerName-parameter  
  `pMarkerName` **must** be a null-terminated UTF-8 string

## [](#_see_also)See Also

[VK\_EXT\_debug\_marker](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_marker.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdDebugMarkerBeginEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDebugMarkerBeginEXT.html), [vkCmdDebugMarkerInsertEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDebugMarkerInsertEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDebugMarkerMarkerInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0