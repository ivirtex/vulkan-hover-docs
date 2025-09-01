# VkImportMetalTextureInfoEXT(3) Manual Page

## Name

VkImportMetalTextureInfoEXT - Structure that identifies Metal MTLTexture objects to use when creating a VkImage.



## [](#_c_specification)C Specification

To import one or more existing Metal `MTLTexture` objects to underlie a [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) object, include one or more `VkImportMetalTextureInfoEXT` structures in the `pNext` chain of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) structure in a [vkCreateImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImage.html) command.

The `VkImportMetalTextureInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_metal_objects
typedef struct VkImportMetalTextureInfoEXT {
    VkStructureType          sType;
    const void*              pNext;
    VkImageAspectFlagBits    plane;
    MTLTexture_id            mtlTexture;
} VkImportMetalTextureInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `plane` specifies the plane of the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) that the `id<MTLTexture>` object should be attached to.
- `mtlTexture` is a the Metal `id<MTLTexture>` object that is to underlie the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) plane.

## [](#_description)Description

The `pNext` chain **must** include one `VkImportMetalTextureInfoEXT` structure for each plane in the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html). The application **must** ensure that the configuration of the Metal `id<MTLTexture>` objects are compatible with the configuration of the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html). Failure to do so results in undefined behavior.

Valid Usage (Implicit)

- [](#VUID-VkImportMetalTextureInfoEXT-sType-sType)VUID-VkImportMetalTextureInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_METAL_TEXTURE_INFO_EXT`
- [](#VUID-VkImportMetalTextureInfoEXT-plane-parameter)VUID-VkImportMetalTextureInfoEXT-plane-parameter  
  `plane` **must** be a valid [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html) value

## [](#_see_also)See Also

[VK\_EXT\_metal\_objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_objects.html), [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImportMetalTextureInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0