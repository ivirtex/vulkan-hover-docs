# VkExportMetalIOSurfaceInfoEXT(3) Manual Page

## Name

VkExportMetalIOSurfaceInfoEXT - Structure that identifies a VkImage
object and corresponding Metal IOSurfaceRef object



## <a href="#_c_specification" class="anchor"></a>C Specification

To export the Metal [IOSurfaceRef](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/IOSurfaceRef.html) object underlying
a [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) object, include a
`VkExportMetalIOSurfaceInfoEXT` structure in the `pNext` chain of the
`pMetalObjectsInfo` parameter of a
[vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkExportMetalObjectsEXT.html) call.

The `VkExportMetalIOSurfaceInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_metal_objects
typedef struct VkExportMetalIOSurfaceInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkImage            image;
    IOSurfaceRef       ioSurface;
} VkExportMetalIOSurfaceInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `image` is a [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html).

- `ioSurface` is the Metal [IOSurfaceRef](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/IOSurfaceRef.html) object
  underlying the [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) object in `image`. The
  implementation will return the [IOSurfaceRef](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/IOSurfaceRef.html) in
  this member, or it will return `NULL` if no
  [IOSurfaceRef](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/IOSurfaceRef.html) could be found underlying the
  [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) object.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkExportMetalIOSurfaceInfoEXT-sType-sType"
  id="VUID-VkExportMetalIOSurfaceInfoEXT-sType-sType"></a>
  VUID-VkExportMetalIOSurfaceInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_EXPORT_METAL_IO_SURFACE_INFO_EXT`

- <a href="#VUID-VkExportMetalIOSurfaceInfoEXT-image-parameter"
  id="VUID-VkExportMetalIOSurfaceInfoEXT-image-parameter"></a>
  VUID-VkExportMetalIOSurfaceInfoEXT-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html),
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExportMetalIOSurfaceInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
