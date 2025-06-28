# VkImageUsageFlags(3) Manual Page

## Name

VkImageUsageFlags - Bitmask of VkImageUsageFlagBits



## [](#_c_specification)C Specification

```c++
// Provided by VK_VERSION_1_0
typedef VkFlags VkImageUsageFlags;
```

## [](#_description)Description

`VkImageUsageFlags` is a bitmask type for setting a mask of zero or more [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html).

When creating a `VkImageView` one of the following [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html) **must** be set:

- `VK_IMAGE_USAGE_SAMPLED_BIT`
- `VK_IMAGE_USAGE_STORAGE_BIT`
- `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`
- `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`
- `VK_IMAGE_USAGE_TRANSIENT_ATTACHMENT_BIT`
- `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- `VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT`
- `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`
- `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`
- `VK_IMAGE_USAGE_VIDEO_ENCODE_SRC_BIT_KHR`
- `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`
- `VK_IMAGE_USAGE_SAMPLE_WEIGHT_BIT_QCOM`
- `VK_IMAGE_USAGE_SAMPLE_BLOCK_MATCH_BIT_QCOM`
- `VK_IMAGE_USAGE_VIDEO_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR`
- `VK_IMAGE_USAGE_VIDEO_ENCODE_EMPHASIS_MAP_BIT_KHR`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFlags.html), [VkFramebufferAttachmentImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentImageInfo.html), [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html), [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html), [VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewUsageCreateInfo.html), [VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV.html), [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html), [VkPhysicalDeviceSparseImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSparseImageFormatInfo2.html), [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html), [VkSharedPresentSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSharedPresentSurfaceCapabilitiesKHR.html), [VkSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2EXT.html), [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html), [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html), [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html), [vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html), [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties.html), [vkGetPhysicalDeviceSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSparseImageFormatProperties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageUsageFlags)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0