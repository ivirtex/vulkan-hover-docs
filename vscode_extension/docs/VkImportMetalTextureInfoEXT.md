# VkImportMetalTextureInfoEXT(3) Manual Page

## Name

VkImportMetalTextureInfoEXT - Structure that identifies Metal MTLTexture
objects to use when creating a VkImage.



## <a href="#_c_specification" class="anchor"></a>C Specification

To import one or more existing Metal `MTLTexture` objects to underlie a
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) object, include one or more
`VkImportMetalTextureInfoEXT` structures in the `pNext` chain of the
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure in a
[vkCreateImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html) command.

The `VkImportMetalTextureInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_metal_objects
typedef struct VkImportMetalTextureInfoEXT {
    VkStructureType          sType;
    const void*              pNext;
    VkImageAspectFlagBits    plane;
    MTLTexture_id            mtlTexture;
} VkImportMetalTextureInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `plane` indicates the plane of the [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) that the
  `id<MTLTexture>` object should be attached to.

- `mtlTexture` is a the Metal `id<MTLTexture>` object that is to
  underlie the [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) plane.

## <a href="#_description" class="anchor"></a>Description

The `pNext` chain **must** include one `VkImportMetalTextureInfoEXT`
structure for each plane in the [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html). The application
**must** ensure that the configuration of the Metal `id<MTLTexture>`
objects are compatible with the configuration of the
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html). Failure to do so results in undefined behavior.

Valid Usage (Implicit)

- <a href="#VUID-VkImportMetalTextureInfoEXT-sType-sType"
  id="VUID-VkImportMetalTextureInfoEXT-sType-sType"></a>
  VUID-VkImportMetalTextureInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_METAL_TEXTURE_INFO_EXT`

- <a href="#VUID-VkImportMetalTextureInfoEXT-plane-parameter"
  id="VUID-VkImportMetalTextureInfoEXT-plane-parameter"></a>
  VUID-VkImportMetalTextureInfoEXT-plane-parameter  
  `plane` **must** be a valid
  [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html),
[VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportMetalTextureInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
