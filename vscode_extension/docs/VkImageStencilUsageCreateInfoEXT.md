# VkImageStencilUsageCreateInfo(3) Manual Page

## Name

VkImageStencilUsageCreateInfo - Specify separate usage flags for the stencil aspect of a depth-stencil image



## [](#_c_specification)C Specification

The `VkImageStencilUsageCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkImageStencilUsageCreateInfo {
    VkStructureType      sType;
    const void*          pNext;
    VkImageUsageFlags    stencilUsage;
} VkImageStencilUsageCreateInfo;
```

or the equivalent

```c++
// Provided by VK_EXT_separate_stencil_usage
typedef VkImageStencilUsageCreateInfo VkImageStencilUsageCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stencilUsage` is a bitmask of [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html) describing the intended usage of the stencil aspect of the image.

## [](#_description)Description

If the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) includes a `VkImageStencilUsageCreateInfo` structure, then that structure includes the usage flags specific to the stencil aspect of the image for an image with a depth-stencil format.

This structure specifies image usages which only apply to the stencil aspect of a depth/stencil format image. When this structure is included in the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), the stencil aspect of the image **must** only be used as specified by `stencilUsage`. When this structure is not included in the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), the stencil aspect of an image **must** only be used as specified by [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage`. Use of other aspects of an image are unaffected by this structure.

This structure **can** also be included in the `pNext` chain of [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html) to query additional capabilities specific to image creation parameter combinations including a separate set of usage flags for the stencil aspect of the image using [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html). When this structure is not included in the `pNext` chain of `VkPhysicalDeviceImageFormatInfo2` then the implicit value of `stencilUsage` matches that of `VkPhysicalDeviceImageFormatInfo2`::`usage`.

Valid Usage

- [](#VUID-VkImageStencilUsageCreateInfo-stencilUsage-02539)VUID-VkImageStencilUsageCreateInfo-stencilUsage-02539  
  If `stencilUsage` includes `VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT`, it **must** not include bits other than `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`

Valid Usage (Implicit)

- [](#VUID-VkImageStencilUsageCreateInfo-sType-sType)VUID-VkImageStencilUsageCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_STENCIL_USAGE_CREATE_INFO`
- [](#VUID-VkImageStencilUsageCreateInfo-stencilUsage-parameter)VUID-VkImageStencilUsageCreateInfo-stencilUsage-parameter  
  `stencilUsage` **must** be a valid combination of [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html) values
- [](#VUID-VkImageStencilUsageCreateInfo-stencilUsage-requiredbitmask)VUID-VkImageStencilUsageCreateInfo-stencilUsage-requiredbitmask  
  `stencilUsage` **must** not be `0`

## [](#_see_also)See Also

[VK\_EXT\_separate\_stencil\_usage](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_separate_stencil_usage.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageStencilUsageCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0