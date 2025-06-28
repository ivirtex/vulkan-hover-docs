# VkDebugMarkerObjectNameInfoEXT(3) Manual Page

## Name

VkDebugMarkerObjectNameInfoEXT - Specify parameters of a name to give to an object



## [](#_c_specification)C Specification

The `VkDebugMarkerObjectNameInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_debug_marker
typedef struct VkDebugMarkerObjectNameInfoEXT {
    VkStructureType               sType;
    const void*                   pNext;
    VkDebugReportObjectTypeEXT    objectType;
    uint64_t                      object;
    const char*                   pObjectName;
} VkDebugMarkerObjectNameInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `objectType` is a [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html) specifying the type of the object to be named.
- `object` is the object to be named.
- `pObjectName` is a null-terminated UTF-8 string specifying the name to apply to `object`.

## [](#_description)Description

Applications **may** change the name associated with an object simply by calling `vkDebugMarkerSetObjectNameEXT` again with a new string. To remove a previously set name, `pObjectName` **should** be an empty string.

Valid Usage

- [](#VUID-VkDebugMarkerObjectNameInfoEXT-objectType-01490)VUID-VkDebugMarkerObjectNameInfoEXT-objectType-01490  
  `objectType` **must** not be `VK_DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT`
- [](#VUID-VkDebugMarkerObjectNameInfoEXT-object-01491)VUID-VkDebugMarkerObjectNameInfoEXT-object-01491  
  `object` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkDebugMarkerObjectNameInfoEXT-object-01492)VUID-VkDebugMarkerObjectNameInfoEXT-object-01492  
  `object` **must** be a Vulkan object of the type associated with `objectType` as defined in [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#debug-report-object-types](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#debug-report-object-types)

Valid Usage (Implicit)

- [](#VUID-VkDebugMarkerObjectNameInfoEXT-sType-sType)VUID-VkDebugMarkerObjectNameInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEBUG_MARKER_OBJECT_NAME_INFO_EXT`
- [](#VUID-VkDebugMarkerObjectNameInfoEXT-pNext-pNext)VUID-VkDebugMarkerObjectNameInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDebugMarkerObjectNameInfoEXT-objectType-parameter)VUID-VkDebugMarkerObjectNameInfoEXT-objectType-parameter  
  `objectType` **must** be a valid [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html) value
- [](#VUID-VkDebugMarkerObjectNameInfoEXT-pObjectName-parameter)VUID-VkDebugMarkerObjectNameInfoEXT-pObjectName-parameter  
  `pObjectName` **must** be a null-terminated UTF-8 string

Host Synchronization

- Host access to `object` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_EXT\_debug\_marker](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_marker.html), [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkDebugMarkerSetObjectNameEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDebugMarkerSetObjectNameEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDebugMarkerObjectNameInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0