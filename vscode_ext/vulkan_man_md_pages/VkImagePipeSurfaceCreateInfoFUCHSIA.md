# VkImagePipeSurfaceCreateInfoFUCHSIA(3) Manual Page

## Name

VkImagePipeSurfaceCreateInfoFUCHSIA - Structure specifying parameters of
a newly created ImagePipe surface object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImagePipeSurfaceCreateInfoFUCHSIA` structure is defined as:

``` c
// Provided by VK_FUCHSIA_imagepipe_surface
typedef struct VkImagePipeSurfaceCreateInfoFUCHSIA {
    VkStructureType                         sType;
    const void*                             pNext;
    VkImagePipeSurfaceCreateFlagsFUCHSIA    flags;
    zx_handle_t                             imagePipeHandle;
} VkImagePipeSurfaceCreateInfoFUCHSIA;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `imagePipeHandle` is a `zx_handle_t` referring to the ImagePipe to
  associate with the surface.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-imagePipeHandle-04863"
  id="VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-imagePipeHandle-04863"></a>
  VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-imagePipeHandle-04863  
  `imagePipeHandle` **must** be a valid `zx_handle_t`

Valid Usage (Implicit)

- <a href="#VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-sType-sType"
  id="VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-sType-sType"></a>
  VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGEPIPE_SURFACE_CREATE_INFO_FUCHSIA`

- <a href="#VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-pNext-pNext"
  id="VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-pNext-pNext"></a>
  VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-flags-zerobitmask"
  id="VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-flags-zerobitmask"></a>
  VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_imagepipe_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_imagepipe_surface.html),
[VkImagePipeSurfaceCreateFlagsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePipeSurfaceCreateFlagsFUCHSIA.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateImagePipeSurfaceFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImagePipeSurfaceFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImagePipeSurfaceCreateInfoFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
