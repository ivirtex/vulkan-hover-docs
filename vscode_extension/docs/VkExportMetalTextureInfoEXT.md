# VkExportMetalTextureInfoEXT(3) Manual Page

## Name

VkExportMetalTextureInfoEXT - Structure that identifies a VkImage, VkImageView, or VkBufferView object and corresponding Metal MTLTexture object



## [](#_c_specification)C Specification

To export a Metal `MTLTexture` object underlying a [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), or [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) object, include a `VkExportMetalTextureInfoEXT` structure in the `pNext` chain of the `pMetalObjectsInfo` parameter of a [vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkExportMetalObjectsEXT.html) call.

The `VkExportMetalTextureInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_metal_objects
typedef struct VkExportMetalTextureInfoEXT {
    VkStructureType          sType;
    const void*              pNext;
    VkImage                  image;
    VkImageView              imageView;
    VkBufferView             bufferView;
    VkImageAspectFlagBits    plane;
    MTLTexture_id            mtlTexture;
} VkExportMetalTextureInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `image` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html).
- `imageView` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html).
- `bufferView` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html).
- `plane` specifies the plane of a multi-planar [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) or [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html).
- `mtlTexture` is the Metal `id<MTLTexture>` object underlying the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), or [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) object in `image`, `imageView`, or `bufferView`, respectively, at the plane indicated in `aspectMask`. The implementation will return the `MTLTexture` in this member, or it will return `NULL` if no `MTLTexture` could be found underlying the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), or [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) object, at the plane indicated in `aspectMask`.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkExportMetalTextureInfoEXT-sType-sType)VUID-VkExportMetalTextureInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_METAL_TEXTURE_INFO_EXT`
- [](#VUID-VkExportMetalTextureInfoEXT-image-parameter)VUID-VkExportMetalTextureInfoEXT-image-parameter  
  If `image` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkExportMetalTextureInfoEXT-imageView-parameter)VUID-VkExportMetalTextureInfoEXT-imageView-parameter  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle
- [](#VUID-VkExportMetalTextureInfoEXT-bufferView-parameter)VUID-VkExportMetalTextureInfoEXT-bufferView-parameter  
  If `bufferView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `bufferView` **must** be a valid [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) handle
- [](#VUID-VkExportMetalTextureInfoEXT-plane-parameter)VUID-VkExportMetalTextureInfoEXT-plane-parameter  
  `plane` **must** be a valid [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html) value
- [](#VUID-VkExportMetalTextureInfoEXT-commonparent)VUID-VkExportMetalTextureInfoEXT-commonparent  
  Each of `bufferView`, `image`, and `imageView` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_EXT\_metal\_objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_objects.html), [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExportMetalTextureInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0