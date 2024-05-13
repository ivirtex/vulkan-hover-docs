# VkImageUsageFlags(3) Manual Page

## Name

VkImageUsageFlags - Bitmask of VkImageUsageFlagBits



## <a href="#_c_specification" class="anchor"></a>C Specification

``` c
// Provided by VK_VERSION_1_0
typedef VkFlags VkImageUsageFlags;
```

## <a href="#_description" class="anchor"></a>Description

`VkImageUsageFlags` is a bitmask type for setting a mask of zero or more
[VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html).

When creating a `VkImageView` one of the following
[VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) **must** be set:

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

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkFramebufferAttachmentImageInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentImageInfo.html),
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html),
[VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html),
[VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html),
[VkImageViewUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewUsageCreateInfo.html),
[VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV.html),
[VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html),
[VkPhysicalDeviceSparseImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSparseImageFormatInfo2.html),
[VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html),
[VkSharedPresentSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSharedPresentSurfaceCapabilitiesKHR.html),
[VkSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2EXT.html),
[VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesKHR.html),
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html),
[VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html),
[vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html),
[vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html),
[vkGetPhysicalDeviceSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageUsageFlags"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
