# VkImportMetalIOSurfaceInfoEXT(3) Manual Page

## Name

VkImportMetalIOSurfaceInfoEXT - Structure that identifies a VkImage
object and corresponding Metal IOSurfaceRef object to use.



## <a href="#_c_specification" class="anchor"></a>C Specification

To import, or create, a Metal [IOSurfaceRef](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/IOSurfaceRef.html) object
to underlie a [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) object, include a
`VkImportMetalIOSurfaceInfoEXT` structure in the `pNext` chain of the
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure in a
[vkCreateImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html) command.

The `VkImportMetalIOSurfaceInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_metal_objects
typedef struct VkImportMetalIOSurfaceInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    IOSurfaceRef       ioSurface;
} VkImportMetalIOSurfaceInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `ioSurface` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or the Metal
  [IOSurfaceRef](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/IOSurfaceRef.html) object that is to underlie the
  [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html).

## <a href="#_description" class="anchor"></a>Description

If `ioSurface` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it will be
used to underlie the [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html). If `ioSurface` is
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the implementation will create a
new `IOSurface` to underlie the [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html).

If provided, the application **must** ensure that the configuration of
the [IOSurfaceRef](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/IOSurfaceRef.html) object is compatible with the
configuration of the [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html). Failure to do so results
in undefined behavior.

Valid Usage (Implicit)

- <a href="#VUID-VkImportMetalIOSurfaceInfoEXT-sType-sType"
  id="VUID-VkImportMetalIOSurfaceInfoEXT-sType-sType"></a>
  VUID-VkImportMetalIOSurfaceInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMPORT_METAL_IO_SURFACE_INFO_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportMetalIOSurfaceInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
