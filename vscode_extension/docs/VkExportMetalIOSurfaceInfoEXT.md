# VkExportMetalIOSurfaceInfoEXT(3) Manual Page

## Name

VkExportMetalIOSurfaceInfoEXT - Structure that identifies a VkImage object and corresponding Metal IOSurfaceRef object



## [](#_c_specification)C Specification

To export the Metal [IOSurfaceRef](https://registry.khronos.org/vulkan/specs/latest/man/html/IOSurfaceRef.html) object underlying a [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) object, include a `VkExportMetalIOSurfaceInfoEXT` structure in the `pNext` chain of the `pMetalObjectsInfo` parameter of a [vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkExportMetalObjectsEXT.html) call.

The `VkExportMetalIOSurfaceInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_metal_objects
typedef struct VkExportMetalIOSurfaceInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkImage            image;
    IOSurfaceRef       ioSurface;
} VkExportMetalIOSurfaceInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `image` is a [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html).
- `ioSurface` is the Metal [IOSurfaceRef](https://registry.khronos.org/vulkan/specs/latest/man/html/IOSurfaceRef.html) object underlying the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) object in `image`. The implementation will return the [IOSurfaceRef](https://registry.khronos.org/vulkan/specs/latest/man/html/IOSurfaceRef.html) in this member, or it will return `NULL` if no [IOSurfaceRef](https://registry.khronos.org/vulkan/specs/latest/man/html/IOSurfaceRef.html) could be found underlying the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) object.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkExportMetalIOSurfaceInfoEXT-sType-sType)VUID-VkExportMetalIOSurfaceInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_METAL_IO_SURFACE_INFO_EXT`
- [](#VUID-VkExportMetalIOSurfaceInfoEXT-image-parameter)VUID-VkExportMetalIOSurfaceInfoEXT-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle

## [](#_see_also)See Also

[VK\_EXT\_metal\_objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_objects.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExportMetalIOSurfaceInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0