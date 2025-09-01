# vkGetPhysicalDeviceImageFormatProperties(3) Manual Page

## Name

vkGetPhysicalDeviceImageFormatProperties - Lists physical device's image format capabilities



## [](#_c_specification)C Specification

To query additional capabilities specific to image types, call:

Warning

This functionality is deprecated by [Vulkan Version 1.1](#versions-1.1). See [Deprecated Functionality](#deprecation-gpdp2) for more information.

```c++
// Provided by VK_VERSION_1_0
VkResult vkGetPhysicalDeviceImageFormatProperties(
    VkPhysicalDevice                            physicalDevice,
    VkFormat                                    format,
    VkImageType                                 type,
    VkImageTiling                               tiling,
    VkImageUsageFlags                           usage,
    VkImageCreateFlags                          flags,
    VkImageFormatProperties*                    pImageFormatProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device from which to query the image capabilities.
- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value specifying the image format, corresponding to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`format`.
- `type` is a [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html) value specifying the image type, corresponding to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`imageType`.
- `tiling` is a [VkImageTiling](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageTiling.html) value specifying the image tiling, corresponding to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`tiling`.
- `usage` is a bitmask of [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html) specifying the intended usage of the image, corresponding to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage`.
- `flags` is a bitmask of [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html) specifying additional parameters of the image, corresponding to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`flags`.
- `pImageFormatProperties` is a pointer to a [VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties.html) structure in which capabilities are returned.

## [](#_description)Description

The `format`, `type`, `tiling`, `usage`, and `flags` parameters correspond to parameters that would be consumed by [vkCreateImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImage.html) (as members of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)).

If `format` is not a supported image format, or if the combination of `format`, `type`, `tiling`, `usage`, and `flags` is not supported for images, then `vkGetPhysicalDeviceImageFormatProperties` returns `VK_ERROR_FORMAT_NOT_SUPPORTED`.

The limitations on an image format that are reported by `vkGetPhysicalDeviceImageFormatProperties` have the following property: if `usage1` and `usage2` of type [VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlags.html) are such that the bits set in `usage1` are a subset of the bits set in `usage2`, and `flags1` and `flags2` of type [VkImageCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlags.html) are such that the bits set in `flags1` are a subset of the bits set in `flags2`, then the limitations for `usage1` and `flags1` **must** be no more strict than the limitations for `usage2` and `flags2`, for all values of `format`, `type`, and `tiling`.

If the [`hostImageCopy`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-hostImageCopy) feature is supported, and:

- `usage` includes `VK_IMAGE_USAGE_SAMPLED_BIT`, and
- `flags` does not include any of `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`, `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`, or `VK_IMAGE_CREATE_SPARSE_ALIASED_BIT`

Then the result of calls to `vkGetPhysicalDeviceImageFormatProperties` with identical parameters except for the inclusion of `VK_IMAGE_USAGE_HOST_TRANSFER_BIT` in `usage` **must** be identical.

Valid Usage

- [](#VUID-vkGetPhysicalDeviceImageFormatProperties-tiling-02248)VUID-vkGetPhysicalDeviceImageFormatProperties-tiling-02248  
  `tiling` **must** not be `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`. (Use [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html) instead)

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceImageFormatProperties-physicalDevice-parameter)VUID-vkGetPhysicalDeviceImageFormatProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceImageFormatProperties-format-parameter)VUID-vkGetPhysicalDeviceImageFormatProperties-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-vkGetPhysicalDeviceImageFormatProperties-type-parameter)VUID-vkGetPhysicalDeviceImageFormatProperties-type-parameter  
  `type` **must** be a valid [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html) value
- [](#VUID-vkGetPhysicalDeviceImageFormatProperties-tiling-parameter)VUID-vkGetPhysicalDeviceImageFormatProperties-tiling-parameter  
  `tiling` **must** be a valid [VkImageTiling](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageTiling.html) value
- [](#VUID-vkGetPhysicalDeviceImageFormatProperties-usage-parameter)VUID-vkGetPhysicalDeviceImageFormatProperties-usage-parameter  
  `usage` **must** be a valid combination of [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html) values
- [](#VUID-vkGetPhysicalDeviceImageFormatProperties-usage-requiredbitmask)VUID-vkGetPhysicalDeviceImageFormatProperties-usage-requiredbitmask  
  `usage` **must** not be `0`
- [](#VUID-vkGetPhysicalDeviceImageFormatProperties-flags-parameter)VUID-vkGetPhysicalDeviceImageFormatProperties-flags-parameter  
  `flags` **must** be a valid combination of [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html) values
- [](#VUID-vkGetPhysicalDeviceImageFormatProperties-pImageFormatProperties-parameter)VUID-vkGetPhysicalDeviceImageFormatProperties-pImageFormatProperties-parameter  
  `pImageFormatProperties` **must** be a valid pointer to a [VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_FORMAT_NOT_SUPPORTED`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkImageCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlags.html), [VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties.html), [VkImageTiling](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageTiling.html), [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html), [VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlags.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceImageFormatProperties).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0