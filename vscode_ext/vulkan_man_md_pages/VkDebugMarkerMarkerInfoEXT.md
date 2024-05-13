# VkDebugMarkerMarkerInfoEXT(3) Manual Page

## Name

VkDebugMarkerMarkerInfoEXT - Specify parameters of a command buffer
marker region



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDebugMarkerMarkerInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_debug_marker
typedef struct VkDebugMarkerMarkerInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    const char*        pMarkerName;
    float              color[4];
} VkDebugMarkerMarkerInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pMarkerName` is a pointer to a null-terminated UTF-8 string
  containing the name of the marker.

- `color` is an **optional** RGBA color value that can be associated
  with the marker. A particular implementation **may** choose to ignore
  this color value. The values contain RGBA values in order, in the
  range 0.0 to 1.0. If all elements in `color` are set to 0.0 then it is
  ignored.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDebugMarkerMarkerInfoEXT-sType-sType"
  id="VUID-VkDebugMarkerMarkerInfoEXT-sType-sType"></a>
  VUID-VkDebugMarkerMarkerInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEBUG_MARKER_MARKER_INFO_EXT`

- <a href="#VUID-VkDebugMarkerMarkerInfoEXT-pNext-pNext"
  id="VUID-VkDebugMarkerMarkerInfoEXT-pNext-pNext"></a>
  VUID-VkDebugMarkerMarkerInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDebugMarkerMarkerInfoEXT-pMarkerName-parameter"
  id="VUID-VkDebugMarkerMarkerInfoEXT-pMarkerName-parameter"></a>
  VUID-VkDebugMarkerMarkerInfoEXT-pMarkerName-parameter  
  `pMarkerName` **must** be a null-terminated UTF-8 string

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_marker](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_marker.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdDebugMarkerBeginEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDebugMarkerBeginEXT.html),
[vkCmdDebugMarkerInsertEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDebugMarkerInsertEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDebugMarkerMarkerInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
