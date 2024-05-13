# VkPhysicalDeviceHostImageCopyPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceHostImageCopyPropertiesEXT - Structure enumerating image
layouts supported by an implementation for host memory copies



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceHostImageCopyPropertiesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_host_image_copy
typedef struct VkPhysicalDeviceHostImageCopyPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           copySrcLayoutCount;
    VkImageLayout*     pCopySrcLayouts;
    uint32_t           copyDstLayoutCount;
    VkImageLayout*     pCopyDstLayouts;
    uint8_t            optimalTilingLayoutUUID[VK_UUID_SIZE];
    VkBool32           identicalMemoryTypeRequirements;
} VkPhysicalDeviceHostImageCopyPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `copySrcLayoutCount` is an integer related to the number of image
  layouts for host copies from images available or queried, as described
  below.

- `pCopySrcLayouts` is a pointer to an array of
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) in which supported image layouts
  for use with host copy operations from images are returned.

- `copyDstLayoutCount` is an integer related to the number of image
  layouts for host copies to images available or queried, as described
  below.

- `pCopyDstLayouts` is a pointer to an array of
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) in which supported image layouts
  for use with host copy operations to images are returned.

- `optimalTilingLayoutUUID` is an array of `VK_UUID_SIZE` `uint8_t`
  values representing a universally unique identifier for the
  implementation’s swizzling layout of images created with
  `VK_IMAGE_TILING_OPTIMAL`.

- `identicalMemoryTypeRequirements` indicates that specifying the
  `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` flag in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage` does not affect
  the memory type requirements of the image.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceHostImageCopyPropertiesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

If `pCopyDstLayouts` is `NULL`, then the number of image layouts that
are supported in
[VkCopyMemoryToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToImageInfoEXT.html)::`dstImageLayout`
and
[VkCopyImageToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToImageInfoEXT.html)::`dstImageLayout`
is returned in `copyDstLayoutCount`. Otherwise, `copyDstLayoutCount`
**must** be set by the user to the number of elements in the
`pCopyDstLayouts` array, and on return the variable is overwritten with
the number of values actually written to `pCopyDstLayouts`. If the value
of `copyDstLayoutCount` is less than the number of image layouts that
are supported, at most `copyDstLayoutCount` values will be written to
`pCopyDstLayouts`. The implementation **must** include the
`VK_IMAGE_LAYOUT_GENERAL` image layout in `pCopyDstLayouts`.

If `pCopySrcLayouts` is `NULL`, then the number of image layouts that
are supported in
[VkCopyImageToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToMemoryInfoEXT.html)::`srcImageLayout`
and
[VkCopyImageToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToImageInfoEXT.html)::`srcImageLayout`
is returned in `copySrcLayoutCount`. Otherwise, `copySrcLayoutCount`
**must** be set by the user to the number of elements in the
`pCopySrcLayouts` array, and on return the variable is overwritten with
the number of values actually written to `pCopySrcLayouts`. If the value
of `copySrcLayoutCount` is less than the number of image layouts that
are supported, at most `copySrcLayoutCount` values will be written to
`pCopySrcLayouts`. The implementation **must** include the
`VK_IMAGE_LAYOUT_GENERAL` image layout in `pCopySrcLayouts`.

The `optimalTilingLayoutUUID` value can be used to ensure compatible
data layouts when using the `VK_HOST_IMAGE_COPY_MEMCPY_EXT` flag in
[vkCopyMemoryToImageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMemoryToImageEXT.html) and
[vkCopyImageToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyImageToMemoryEXT.html).

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceHostImageCopyPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceHostImageCopyPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceHostImageCopyPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_IMAGE_COPY_PROPERTIES_EXT`

- <a
  href="#VUID-VkPhysicalDeviceHostImageCopyPropertiesEXT-pCopySrcLayouts-parameter"
  id="VUID-VkPhysicalDeviceHostImageCopyPropertiesEXT-pCopySrcLayouts-parameter"></a>
  VUID-VkPhysicalDeviceHostImageCopyPropertiesEXT-pCopySrcLayouts-parameter  
  If `copySrcLayoutCount` is not `0`, and `pCopySrcLayouts` is not
  `NULL`, `pCopySrcLayouts` **must** be a valid pointer to an array of
  `copySrcLayoutCount` [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) values

- <a
  href="#VUID-VkPhysicalDeviceHostImageCopyPropertiesEXT-pCopyDstLayouts-parameter"
  id="VUID-VkPhysicalDeviceHostImageCopyPropertiesEXT-pCopyDstLayouts-parameter"></a>
  VUID-VkPhysicalDeviceHostImageCopyPropertiesEXT-pCopyDstLayouts-parameter  
  If `copyDstLayoutCount` is not `0`, and `pCopyDstLayouts` is not
  `NULL`, `pCopyDstLayouts` **must** be a valid pointer to an array of
  `copyDstLayoutCount` [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_host_image_copy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_host_image_copy.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceHostImageCopyPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
