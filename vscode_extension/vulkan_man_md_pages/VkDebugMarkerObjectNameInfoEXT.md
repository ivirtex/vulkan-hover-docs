# VkDebugMarkerObjectNameInfoEXT(3) Manual Page

## Name

VkDebugMarkerObjectNameInfoEXT - Specify parameters of a name to give to
an object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDebugMarkerObjectNameInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_debug_marker
typedef struct VkDebugMarkerObjectNameInfoEXT {
    VkStructureType               sType;
    const void*                   pNext;
    VkDebugReportObjectTypeEXT    objectType;
    uint64_t                      object;
    const char*                   pObjectName;
} VkDebugMarkerObjectNameInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `objectType` is a
  [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html)
  specifying the type of the object to be named.

- `object` is the object to be named.

- `pObjectName` is a null-terminated UTF-8 string specifying the name to
  apply to `object`.

## <a href="#_description" class="anchor"></a>Description

Applications **may** change the name associated with an object simply by
calling `vkDebugMarkerSetObjectNameEXT` again with a new string. To
remove a previously set name, `pObjectName` **should** be set to an
empty string.

Valid Usage

- <a href="#VUID-VkDebugMarkerObjectNameInfoEXT-objectType-01490"
  id="VUID-VkDebugMarkerObjectNameInfoEXT-objectType-01490"></a>
  VUID-VkDebugMarkerObjectNameInfoEXT-objectType-01490  
  `objectType` **must** not be `VK_DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT`

- <a href="#VUID-VkDebugMarkerObjectNameInfoEXT-object-01491"
  id="VUID-VkDebugMarkerObjectNameInfoEXT-object-01491"></a>
  VUID-VkDebugMarkerObjectNameInfoEXT-object-01491  
  `object` **must** not be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkDebugMarkerObjectNameInfoEXT-object-01492"
  id="VUID-VkDebugMarkerObjectNameInfoEXT-object-01492"></a>
  VUID-VkDebugMarkerObjectNameInfoEXT-object-01492  
  `object` **must** be a Vulkan object of the type associated with
  `objectType` as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#debug-report-object-types"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#debug-report-object-types</a>

Valid Usage (Implicit)

- <a href="#VUID-VkDebugMarkerObjectNameInfoEXT-sType-sType"
  id="VUID-VkDebugMarkerObjectNameInfoEXT-sType-sType"></a>
  VUID-VkDebugMarkerObjectNameInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_NAME_INFO_EXT`

- <a href="#VUID-VkDebugMarkerObjectNameInfoEXT-pNext-pNext"
  id="VUID-VkDebugMarkerObjectNameInfoEXT-pNext-pNext"></a>
  VUID-VkDebugMarkerObjectNameInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDebugMarkerObjectNameInfoEXT-objectType-parameter"
  id="VUID-VkDebugMarkerObjectNameInfoEXT-objectType-parameter"></a>
  VUID-VkDebugMarkerObjectNameInfoEXT-objectType-parameter  
  `objectType` **must** be a valid
  [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html) value

- <a href="#VUID-VkDebugMarkerObjectNameInfoEXT-pObjectName-parameter"
  id="VUID-VkDebugMarkerObjectNameInfoEXT-pObjectName-parameter"></a>
  VUID-VkDebugMarkerObjectNameInfoEXT-pObjectName-parameter  
  `pObjectName` **must** be a null-terminated UTF-8 string

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_marker](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_marker.html),
[VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkDebugMarkerSetObjectNameEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDebugMarkerSetObjectNameEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDebugMarkerObjectNameInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
