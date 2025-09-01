# VkImagePipeSurfaceCreateInfoFUCHSIA(3) Manual Page

## Name

VkImagePipeSurfaceCreateInfoFUCHSIA - Structure specifying parameters of a newly created ImagePipe surface object



## [](#_c_specification)C Specification

The `VkImagePipeSurfaceCreateInfoFUCHSIA` structure is defined as:

```c++
// Provided by VK_FUCHSIA_imagepipe_surface
typedef struct VkImagePipeSurfaceCreateInfoFUCHSIA {
    VkStructureType                         sType;
    const void*                             pNext;
    VkImagePipeSurfaceCreateFlagsFUCHSIA    flags;
    zx_handle_t                             imagePipeHandle;
} VkImagePipeSurfaceCreateInfoFUCHSIA;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `imagePipeHandle` is a `zx_handle_t` referring to the ImagePipe to associate with the surface.

## [](#_description)Description

Valid Usage

- [](#VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-imagePipeHandle-04863)VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-imagePipeHandle-04863  
  `imagePipeHandle` **must** be a valid `zx_handle_t`

Valid Usage (Implicit)

- [](#VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-sType-sType)VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGEPIPE_SURFACE_CREATE_INFO_FUCHSIA`
- [](#VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-pNext-pNext)VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-flags-zerobitmask)VUID-VkImagePipeSurfaceCreateInfoFUCHSIA-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_FUCHSIA\_imagepipe\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_imagepipe_surface.html), [VkImagePipeSurfaceCreateFlagsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImagePipeSurfaceCreateFlagsFUCHSIA.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateImagePipeSurfaceFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImagePipeSurfaceFUCHSIA.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImagePipeSurfaceCreateInfoFUCHSIA).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0