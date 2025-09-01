# VkPhysicalDeviceHostImageCopyProperties(3) Manual Page

## Name

VkPhysicalDeviceHostImageCopyProperties - Structure enumerating image layouts supported by an implementation for host memory copies



## [](#_c_specification)C Specification

The `VkPhysicalDeviceHostImageCopyProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPhysicalDeviceHostImageCopyProperties {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           copySrcLayoutCount;
    VkImageLayout*     pCopySrcLayouts;
    uint32_t           copyDstLayoutCount;
    VkImageLayout*     pCopyDstLayouts;
    uint8_t            optimalTilingLayoutUUID[VK_UUID_SIZE];
    VkBool32           identicalMemoryTypeRequirements;
} VkPhysicalDeviceHostImageCopyProperties;
```

or the equivalent

```c++
// Provided by VK_EXT_host_image_copy
typedef VkPhysicalDeviceHostImageCopyProperties VkPhysicalDeviceHostImageCopyPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- `copySrcLayoutCount` is an integer related to the number of image layouts for host copies from images available or queried, as described below.
- `pCopySrcLayouts` is a pointer to an array of [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) in which supported image layouts for use with host copy operations from images are returned.
- `copyDstLayoutCount` is an integer related to the number of image layouts for host copies to images available or queried, as described below.
- `pCopyDstLayouts` is a pointer to an array of [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) in which supported image layouts for use with host copy operations to images are returned.
- `optimalTilingLayoutUUID` is an array of `VK_UUID_SIZE` `uint8_t` values representing a universally unique identifier for the implementation’s swizzling layout of images created with `VK_IMAGE_TILING_OPTIMAL`.
- `identicalMemoryTypeRequirements` indicates that specifying the `VK_IMAGE_USAGE_HOST_TRANSFER_BIT` flag in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage` does not affect the memory type requirements of the image.

If the `VkPhysicalDeviceHostImageCopyProperties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

If `pCopyDstLayouts` is `NULL`, then the number of image layouts that are supported in [VkCopyMemoryToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageInfo.html)::`dstImageLayout` and [VkCopyImageToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToImageInfo.html)::`dstImageLayout` is returned in `copyDstLayoutCount`. Otherwise, `copyDstLayoutCount` **must** be set by the application to the number of elements in the `pCopyDstLayouts` array, and on return the variable is overwritten with the number of values actually written to `pCopyDstLayouts`. If the value of `copyDstLayoutCount` is less than the number of image layouts that are supported, at most `copyDstLayoutCount` values will be written to `pCopyDstLayouts`. The implementation **must** include the `VK_IMAGE_LAYOUT_GENERAL` image layout in `pCopyDstLayouts`. If the [`unifiedImageLayouts`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-unifiedImageLayouts) feature is enabled, the implementation **must** include all the image layouts that are interchangeable with `VK_IMAGE_LAYOUT_GENERAL` in `pCopyDstLayouts`.

If `pCopySrcLayouts` is `NULL`, then the number of image layouts that are supported in [VkCopyImageToMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToMemoryInfo.html)::`srcImageLayout` and [VkCopyImageToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToImageInfo.html)::`srcImageLayout` is returned in `copySrcLayoutCount`. Otherwise, `copySrcLayoutCount` **must** be set by the application to the number of elements in the `pCopySrcLayouts` array, and on return the variable is overwritten with the number of values actually written to `pCopySrcLayouts`. If the value of `copySrcLayoutCount` is less than the number of image layouts that are supported, at most `copySrcLayoutCount` values will be written to `pCopySrcLayouts`. The implementation **must** include the `VK_IMAGE_LAYOUT_GENERAL` image layout in `pCopySrcLayouts`. If the [`unifiedImageLayouts`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-unifiedImageLayouts) feature is enabled, the implementation **must** include all the image layouts that are interchangeable with `VK_IMAGE_LAYOUT_GENERAL` in `pCopySrcLayouts`.

The `optimalTilingLayoutUUID` value can be used to ensure compatible data layouts when using the `VK_HOST_IMAGE_COPY_MEMCPY_BIT` flag in [vkCopyMemoryToImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMemoryToImage.html) and [vkCopyImageToMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyImageToMemory.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceHostImageCopyProperties-sType-sType)VUID-VkPhysicalDeviceHostImageCopyProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_IMAGE_COPY_PROPERTIES`
- [](#VUID-VkPhysicalDeviceHostImageCopyProperties-pCopySrcLayouts-parameter)VUID-VkPhysicalDeviceHostImageCopyProperties-pCopySrcLayouts-parameter  
  If `copySrcLayoutCount` is not `0`, and `pCopySrcLayouts` is not `NULL`, `pCopySrcLayouts` **must** be a valid pointer to an array of `copySrcLayoutCount` [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) values
- [](#VUID-VkPhysicalDeviceHostImageCopyProperties-pCopyDstLayouts-parameter)VUID-VkPhysicalDeviceHostImageCopyProperties-pCopyDstLayouts-parameter  
  If `copyDstLayoutCount` is not `0`, and `pCopyDstLayouts` is not `NULL`, `pCopyDstLayouts` **must** be a valid pointer to an array of `copyDstLayoutCount` [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) values

## [](#_see_also)See Also

[VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceHostImageCopyProperties).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0