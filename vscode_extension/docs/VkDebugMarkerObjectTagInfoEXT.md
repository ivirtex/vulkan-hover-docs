# VkDebugMarkerObjectTagInfoEXT(3) Manual Page

## Name

VkDebugMarkerObjectTagInfoEXT - Specify parameters of a tag to attach to an object



## [](#_c_specification)C Specification

The `VkDebugMarkerObjectTagInfoEXT` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `objectType` is a [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html) specifying the type of the object to be named.
- `object` is the object to be tagged.
- `tagName` is a numerical identifier of the tag.
- `tagSize` is the number of bytes of data to attach to the object.
- `pTag` is a pointer to an array of `tagSize` bytes containing the data to be associated with the object.

## [](#_description)Description

The `tagName` parameter gives a name or identifier to the type of data being tagged. This can be used by debugging layers to easily filter for only data that can be used by that implementation.

Valid Usage

- [](#VUID-VkDebugMarkerObjectTagInfoEXT-objectType-01493)VUID-VkDebugMarkerObjectTagInfoEXT-objectType-01493  
  `objectType` **must** not be `VK_DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT`
- [](#VUID-VkDebugMarkerObjectTagInfoEXT-object-01494)VUID-VkDebugMarkerObjectTagInfoEXT-object-01494  
  `object` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkDebugMarkerObjectTagInfoEXT-object-01495)VUID-VkDebugMarkerObjectTagInfoEXT-object-01495  
  `object` **must** be a Vulkan object of the type associated with `objectType` as defined in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#debug-report-object-types](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#debug-report-object-types)

Valid Usage (Implicit)

- [](#VUID-VkDebugMarkerObjectTagInfoEXT-sType-sType)VUID-VkDebugMarkerObjectTagInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_TAG_INFO_EXT`
- [](#VUID-VkDebugMarkerObjectTagInfoEXT-pNext-pNext)VUID-VkDebugMarkerObjectTagInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDebugMarkerObjectTagInfoEXT-objectType-parameter)VUID-VkDebugMarkerObjectTagInfoEXT-objectType-parameter  
  `objectType` **must** be a valid [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html) value
- [](#VUID-VkDebugMarkerObjectTagInfoEXT-pTag-parameter)VUID-VkDebugMarkerObjectTagInfoEXT-pTag-parameter  
  `pTag` **must** be a valid pointer to an array of `tagSize` bytes
- [](#VUID-VkDebugMarkerObjectTagInfoEXT-tagSize-arraylength)VUID-VkDebugMarkerObjectTagInfoEXT-tagSize-arraylength  
  `tagSize` **must** be greater than `0`

Host Synchronization

- Host access to `object` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_EXT\_debug\_marker](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_marker.html), [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkDebugMarkerSetObjectTagEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDebugMarkerSetObjectTagEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDebugMarkerObjectTagInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0