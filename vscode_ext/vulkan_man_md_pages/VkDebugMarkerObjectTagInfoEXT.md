# VkDebugMarkerObjectTagInfoEXT(3) Manual Page

## Name

VkDebugMarkerObjectTagInfoEXT - Specify parameters of a tag to attach to
an object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDebugMarkerObjectTagInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_debug_marker
typedef struct VkDebugMarkerObjectTagInfoEXT {
    VkStructureType               sType;
    const void*                   pNext;
    VkDebugReportObjectTypeEXT    objectType;
    uint64_t                      object;
    uint64_t                      tagName;
    size_t                        tagSize;
    const void*                   pTag;
} VkDebugMarkerObjectTagInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `objectType` is a
  [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html)
  specifying the type of the object to be named.

- `object` is the object to be tagged.

- `tagName` is a numerical identifier of the tag.

- `tagSize` is the number of bytes of data to attach to the object.

- `pTag` is a pointer to an array of `tagSize` bytes containing the data
  to be associated with the object.

## <a href="#_description" class="anchor"></a>Description

The `tagName` parameter gives a name or identifier to the type of data
being tagged. This can be used by debugging layers to easily filter for
only data that can be used by that implementation.

Valid Usage

- <a href="#VUID-VkDebugMarkerObjectTagInfoEXT-objectType-01493"
  id="VUID-VkDebugMarkerObjectTagInfoEXT-objectType-01493"></a>
  VUID-VkDebugMarkerObjectTagInfoEXT-objectType-01493  
  `objectType` **must** not be `VK_DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT`

- <a href="#VUID-VkDebugMarkerObjectTagInfoEXT-object-01494"
  id="VUID-VkDebugMarkerObjectTagInfoEXT-object-01494"></a>
  VUID-VkDebugMarkerObjectTagInfoEXT-object-01494  
  `object` **must** not be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkDebugMarkerObjectTagInfoEXT-object-01495"
  id="VUID-VkDebugMarkerObjectTagInfoEXT-object-01495"></a>
  VUID-VkDebugMarkerObjectTagInfoEXT-object-01495  
  `object` **must** be a Vulkan object of the type associated with
  `objectType` as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#debug-report-object-types"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#debug-report-object-types</a>

Valid Usage (Implicit)

- <a href="#VUID-VkDebugMarkerObjectTagInfoEXT-sType-sType"
  id="VUID-VkDebugMarkerObjectTagInfoEXT-sType-sType"></a>
  VUID-VkDebugMarkerObjectTagInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_TAG_INFO_EXT`

- <a href="#VUID-VkDebugMarkerObjectTagInfoEXT-pNext-pNext"
  id="VUID-VkDebugMarkerObjectTagInfoEXT-pNext-pNext"></a>
  VUID-VkDebugMarkerObjectTagInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDebugMarkerObjectTagInfoEXT-objectType-parameter"
  id="VUID-VkDebugMarkerObjectTagInfoEXT-objectType-parameter"></a>
  VUID-VkDebugMarkerObjectTagInfoEXT-objectType-parameter  
  `objectType` **must** be a valid
  [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html) value

- <a href="#VUID-VkDebugMarkerObjectTagInfoEXT-pTag-parameter"
  id="VUID-VkDebugMarkerObjectTagInfoEXT-pTag-parameter"></a>
  VUID-VkDebugMarkerObjectTagInfoEXT-pTag-parameter  
  `pTag` **must** be a valid pointer to an array of `tagSize` bytes

- <a href="#VUID-VkDebugMarkerObjectTagInfoEXT-tagSize-arraylength"
  id="VUID-VkDebugMarkerObjectTagInfoEXT-tagSize-arraylength"></a>
  VUID-VkDebugMarkerObjectTagInfoEXT-tagSize-arraylength  
  `tagSize` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_marker](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_marker.html),
[VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkDebugMarkerSetObjectTagEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDebugMarkerSetObjectTagEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDebugMarkerObjectTagInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
