# vkGetPhysicalDeviceImageFormatProperties2(3) Manual Page

## Name

vkGetPhysicalDeviceImageFormatProperties2 - Lists physical device's image format capabilities



## [](#_c_specification)C Specification

To query additional capabilities specific to image types, call:

```c++
// Provided by VK_VERSION_1_1
VkResult vkGetPhysicalDeviceImageFormatProperties2(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceImageFormatInfo2*     pImageFormatInfo,
    VkImageFormatProperties2*                   pImageFormatProperties);
```

or the equivalent command

```c++
// Provided by VK_KHR_get_physical_device_properties2
VkResult vkGetPhysicalDeviceImageFormatProperties2KHR(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceImageFormatInfo2*     pImageFormatInfo,
    VkImageFormatProperties2*                   pImageFormatProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device from which to query the image capabilities.
- `pImageFormatInfo` is a pointer to a [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html) structure describing the parameters that would be consumed by [vkCreateImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImage.html).
- `pImageFormatProperties` is a pointer to a [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html) structure in which capabilities are returned.

## [](#_description)Description

`vkGetPhysicalDeviceImageFormatProperties2` behaves similarly to [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties.html), with the ability to return extended information in a `pNext` chain of output structures.

If the `pNext` chain of `pImageFormatInfo` includes a [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html) structure with a `profileCount` member greater than `0`, then this command returns format capabilities specific to image types used in conjunction with the specified [video profiles](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-profiles). In this case, this command will return one of the [video-profile-specific error codes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-profile-error-codes) if any of the profiles specified via [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html)::`pProfiles` are not supported. Furthermore, if [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`usage` includes any image usage flag not supported by the specified video profiles, then this command returns `VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR`.

If the [`hostImageCopy`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-hostImageCopy) feature is supported, and:

- `pImageFormatInfo->usage` includes `VK_IMAGE_USAGE_SAMPLED_BIT`
- `pImageFormatInfo->flags` does not include either of `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`, `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`, or `VK_IMAGE_CREATE_SPARSE_ALIASED_BIT`
- The `pNext` chain of `pImageFormatInfo` does not include a [VkPhysicalDeviceExternalImageFormatInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalImageFormatInfo.html) structure with non-zero `handleType`
- `pImageFormatInfo->tiling` is not `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`

Then the result of calls to `vkGetPhysicalDeviceImageFormatProperties2` with identical parameters except for the inclusion of `VK_IMAGE_USAGE_HOST_TRANSFER_BIT` in `pImageFormatInfo->usage` **must** be identical.

Valid Usage

- [](#VUID-vkGetPhysicalDeviceImageFormatProperties2-pNext-01868)VUID-vkGetPhysicalDeviceImageFormatProperties2-pNext-01868  
  If the `pNext` chain of `pImageFormatProperties` includes a [VkAndroidHardwareBufferUsageANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferUsageANDROID.html) structure, the `pNext` chain of `pImageFormatInfo` **must** include a [VkPhysicalDeviceExternalImageFormatInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalImageFormatInfo.html) structure with `handleType` set to `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`
- [](#VUID-vkGetPhysicalDeviceImageFormatProperties2-pNext-09004)VUID-vkGetPhysicalDeviceImageFormatProperties2-pNext-09004  
  If the `pNext` chain of `pImageFormatProperties` includes a [VkHostImageCopyDevicePerformanceQuery](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageCopyDevicePerformanceQuery.html) structure, `pImageFormatInfo->usage` **must** contain `VK_IMAGE_USAGE_HOST_TRANSFER_BIT`

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceImageFormatProperties2-physicalDevice-parameter)VUID-vkGetPhysicalDeviceImageFormatProperties2-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceImageFormatProperties2-pImageFormatInfo-parameter)VUID-vkGetPhysicalDeviceImageFormatProperties2-pImageFormatInfo-parameter  
  `pImageFormatInfo` **must** be a valid pointer to a valid [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html) structure
- [](#VUID-vkGetPhysicalDeviceImageFormatProperties2-pImageFormatProperties-parameter)VUID-vkGetPhysicalDeviceImageFormatProperties2-pImageFormatProperties-parameter  
  `pImageFormatProperties` **must** be a valid pointer to a [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_FORMAT_NOT_SUPPORTED`
- `VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`
- `VK_ERROR_VIDEO_PICTURE_LAYOUT_NOT_SUPPORTED_KHR`
- `VK_ERROR_VIDEO_PROFILE_CODEC_NOT_SUPPORTED_KHR`
- `VK_ERROR_VIDEO_PROFILE_FORMAT_NOT_SUPPORTED_KHR`
- `VK_ERROR_VIDEO_PROFILE_OPERATION_NOT_SUPPORTED_KHR`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceImageFormatProperties2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0