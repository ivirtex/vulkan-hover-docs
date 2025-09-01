# VkImportMetalIOSurfaceInfoEXT(3) Manual Page

## Name

VkImportMetalIOSurfaceInfoEXT - Structure that identifies a VkImage object and corresponding Metal IOSurfaceRef object to use.



## [](#_c_specification)C Specification

To import, or create, a Metal [IOSurfaceRef](https://registry.khronos.org/vulkan/specs/latest/man/html/IOSurfaceRef.html) object to underlie a [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) object, include a `VkImportMetalIOSurfaceInfoEXT` structure in the `pNext` chain of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) structure in a [vkCreateImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImage.html) command.

The `VkImportMetalIOSurfaceInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_metal_objects
typedef struct VkImportMetalIOSurfaceInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    IOSurfaceRef       ioSurface;
} VkImportMetalIOSurfaceInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `ioSurface` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or the Metal [IOSurfaceRef](https://registry.khronos.org/vulkan/specs/latest/man/html/IOSurfaceRef.html) object that is to underlie the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html).

## [](#_description)Description

If `ioSurface` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it will be used to underlie the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html). If `ioSurface` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the implementation will create a new `IOSurface` to underlie the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html).

If provided, the application **must** ensure that the configuration of the [IOSurfaceRef](https://registry.khronos.org/vulkan/specs/latest/man/html/IOSurfaceRef.html) object is compatible with the configuration of the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html). Failure to do so results in undefined behavior.

Valid Usage (Implicit)

- [](#VUID-VkImportMetalIOSurfaceInfoEXT-sType-sType)VUID-VkImportMetalIOSurfaceInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_METAL_IO_SURFACE_INFO_EXT`

## [](#_see_also)See Also

[VK\_EXT\_metal\_objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_objects.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImportMetalIOSurfaceInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0