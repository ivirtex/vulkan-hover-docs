# vkGetPhysicalDeviceImageFormatProperties2(3) Manual Page

## Name

vkGetPhysicalDeviceImageFormatProperties2 - Lists physical device's
image format capabilities



## <a href="#_c_specification" class="anchor"></a>C Specification

To query additional capabilities specific to image types, call:

``` c
// Provided by VK_VERSION_1_1
VkResult vkGetPhysicalDeviceImageFormatProperties2(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceImageFormatInfo2*     pImageFormatInfo,
    VkImageFormatProperties2*                   pImageFormatProperties);
```

or the equivalent command

``` c
// Provided by VK_KHR_get_physical_device_properties2
VkResult vkGetPhysicalDeviceImageFormatProperties2KHR(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceImageFormatInfo2*     pImageFormatInfo,
    VkImageFormatProperties2*                   pImageFormatProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device from which to query the image
  capabilities.

- `pImageFormatInfo` is a pointer to a
  [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)
  structure describing the parameters that would be consumed by
  [vkCreateImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html).

- `pImageFormatProperties` is a pointer to a
  [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html) structure in
  which capabilities are returned.

## <a href="#_description" class="anchor"></a>Description

`vkGetPhysicalDeviceImageFormatProperties2` behaves similarly to
[vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html),
with the ability to return extended information in a `pNext` chain of
output structures.

If the `pNext` chain of `pImageFormatInfo` includes a
[VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html) structure
with a `profileCount` member greater than `0`, then this command returns
format capabilities specific to image types used in conjunction with the
specified <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profiles"
target="_blank" rel="noopener">video profiles</a>. In this case, this
command will return one of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profile-error-codes"
target="_blank" rel="noopener">video-profile-specific error codes</a> if
any of the profiles specified via
[VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)::`pProfiles`
are not supported. Furthermore, if
[VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)::`usage`
includes any image usage flag not supported by the specified video
profiles, then this command returns
`VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR`.

Valid Usage

- <a href="#VUID-vkGetPhysicalDeviceImageFormatProperties2-pNext-01868"
  id="VUID-vkGetPhysicalDeviceImageFormatProperties2-pNext-01868"></a>
  VUID-vkGetPhysicalDeviceImageFormatProperties2-pNext-01868  
  If the `pNext` chain of `pImageFormatProperties` includes a
  [VkAndroidHardwareBufferUsageANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferUsageANDROID.html)
  structure, the `pNext` chain of `pImageFormatInfo` **must** include a
  [VkPhysicalDeviceExternalImageFormatInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalImageFormatInfo.html)
  structure with `handleType` set to
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`

- <a href="#VUID-vkGetPhysicalDeviceImageFormatProperties2-pNext-09004"
  id="VUID-vkGetPhysicalDeviceImageFormatProperties2-pNext-09004"></a>
  VUID-vkGetPhysicalDeviceImageFormatProperties2-pNext-09004  
  If the `pNext` chain of `pImageFormatProperties` includes a
  [VkHostImageCopyDevicePerformanceQueryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageCopyDevicePerformanceQueryEXT.html)
  structure, `pImageFormatInfo->usage` **must** contain
  `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT`

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceImageFormatProperties2-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceImageFormatProperties2-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceImageFormatProperties2-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceImageFormatProperties2-pImageFormatInfo-parameter"
  id="VUID-vkGetPhysicalDeviceImageFormatProperties2-pImageFormatInfo-parameter"></a>
  VUID-vkGetPhysicalDeviceImageFormatProperties2-pImageFormatInfo-parameter  
  `pImageFormatInfo` **must** be a valid pointer to a valid
  [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceImageFormatProperties2-pImageFormatProperties-parameter"
  id="VUID-vkGetPhysicalDeviceImageFormatProperties2-pImageFormatProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceImageFormatProperties2-pImageFormatProperties-parameter  
  `pImageFormatProperties` **must** be a valid pointer to a
  [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html) structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_FORMAT_NOT_SUPPORTED`

- `VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR`

- `VK_ERROR_VIDEO_PROFILE_OPERATION_NOT_SUPPORTED_KHR`

- `VK_ERROR_VIDEO_PROFILE_FORMAT_NOT_SUPPORTED_KHR`

- `VK_ERROR_VIDEO_PICTURE_LAYOUT_NOT_SUPPORTED_KHR`

- `VK_ERROR_VIDEO_PROFILE_CODEC_NOT_SUPPORTED_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceImageFormatProperties2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
