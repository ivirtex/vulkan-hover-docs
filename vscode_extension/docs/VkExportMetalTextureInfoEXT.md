# VkExportMetalTextureInfoEXT(3) Manual Page

## Name

VkExportMetalTextureInfoEXT - Structure that identifies a VkImage,
VkImageView, or VkBufferView object and corresponding Metal MTLTexture
object



## <a href="#_c_specification" class="anchor"></a>C Specification

To export a Metal `MTLTexture` object underlying a
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html), [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html), or
[VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) object, include a
`VkExportMetalTextureInfoEXT` structure in the `pNext` chain of the
`pMetalObjectsInfo` parameter of a
[vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkExportMetalObjectsEXT.html) call.

The `VkExportMetalTextureInfoEXT` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `image` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or a
  [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html).

- `imageView` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html).

- `bufferView` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or a
  [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html).

- `plane` indicates the plane of a multi-planar [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html)
  or [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html).

- `mtlTexture` is the Metal `id<MTLTexture>` object underlying the
  [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html), [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html), or
  [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) object in `image`, `imageView`, or
  `bufferView`, respectively, at the plane indicated in `aspectMask`.
  The implementation will return the `MTLTexture` in this member, or it
  will return `NULL` if no `MTLTexture` could be found underlying the
  [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html), [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html), or
  [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) object, at the plane indicated in
  `aspectMask`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkExportMetalTextureInfoEXT-sType-sType"
  id="VUID-VkExportMetalTextureInfoEXT-sType-sType"></a>
  VUID-VkExportMetalTextureInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_METAL_TEXTURE_INFO_EXT`

- <a href="#VUID-VkExportMetalTextureInfoEXT-image-parameter"
  id="VUID-VkExportMetalTextureInfoEXT-image-parameter"></a>
  VUID-VkExportMetalTextureInfoEXT-image-parameter  
  If `image` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `image`
  **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkExportMetalTextureInfoEXT-imageView-parameter"
  id="VUID-VkExportMetalTextureInfoEXT-imageView-parameter"></a>
  VUID-VkExportMetalTextureInfoEXT-imageView-parameter  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handle

- <a href="#VUID-VkExportMetalTextureInfoEXT-bufferView-parameter"
  id="VUID-VkExportMetalTextureInfoEXT-bufferView-parameter"></a>
  VUID-VkExportMetalTextureInfoEXT-bufferView-parameter  
  If `bufferView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `bufferView` **must** be a valid [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html)
  handle

- <a href="#VUID-VkExportMetalTextureInfoEXT-plane-parameter"
  id="VUID-VkExportMetalTextureInfoEXT-plane-parameter"></a>
  VUID-VkExportMetalTextureInfoEXT-plane-parameter  
  `plane` **must** be a valid
  [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html) value

- <a href="#VUID-VkExportMetalTextureInfoEXT-commonparent"
  id="VUID-VkExportMetalTextureInfoEXT-commonparent"></a>
  VUID-VkExportMetalTextureInfoEXT-commonparent  
  Each of `bufferView`, `image`, and `imageView` that are valid handles
  of non-ignored parameters **must** have been created, allocated, or
  retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html),
[VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html),
[VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExportMetalTextureInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
