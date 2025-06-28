# vkGetPhysicalDeviceVideoCapabilitiesKHR(3) Manual Page

## Name

vkGetPhysicalDeviceVideoCapabilitiesKHR - Query video coding capabilities



## [](#_c_specification)C Specification

To query video coding capabilities for a specific video profile, call:

```c++
// Provided by VK_KHR_video_queue
VkResult vkGetPhysicalDeviceVideoCapabilitiesKHR(
    VkPhysicalDevice                            physicalDevice,
    const VkVideoProfileInfoKHR*                pVideoProfile,
    VkVideoCapabilitiesKHR*                     pCapabilities);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device from which to query the video decode or encode capabilities.
- `pVideoProfile` is a pointer to a [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure.
- `pCapabilities` is a pointer to a [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html) structure in which the capabilities are returned.

## [](#_description)Description

If the [video profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-profiles) described by `pVideoProfile` is supported by the implementation, then this command returns `VK_SUCCESS` and `pCapabilities` is filled with the capabilities supported with the specified video profile. Otherwise, one of the [video-profile-specific error codes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-profile-error-codes) are returned.

Valid Usage

- [](#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07183)VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07183  
  If `pVideoProfile->videoCodecOperation` specifies a decode operation, then the `pNext` chain of `pCapabilities` **must** include a [VkVideoDecodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeCapabilitiesKHR.html) structure
- [](#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07184)VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07184  
  If `pVideoProfile->videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the `pNext` chain of `pCapabilities` **must** include a [VkVideoDecodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264CapabilitiesKHR.html) structure
- [](#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07185)VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07185  
  If `pVideoProfile->videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the `pNext` chain of `pCapabilities` **must** include a [VkVideoDecodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265CapabilitiesKHR.html) structure
- [](#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-10792)VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-10792  
  If `pVideoProfile->videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_DECODE_VP9_BIT_KHR`, then the `pNext` chain of `pCapabilities` **must** include a [VkVideoDecodeVP9CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeVP9CapabilitiesKHR.html) structure
- [](#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-09257)VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-09257  
  If `pVideoProfile->videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then the `pNext` chain of `pCapabilities` **must** include a [VkVideoDecodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeAV1CapabilitiesKHR.html) structure
- [](#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07186)VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07186  
  If `pVideoProfile->videoCodecOperation` specifies an encode operation, then the `pNext` chain of `pCapabilities` **must** include a [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeCapabilitiesKHR.html) structure
- [](#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07187)VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07187  
  If `pVideoProfile->videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the `pNext` chain of `pCapabilities` **must** include a [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264CapabilitiesKHR.html) structure
- [](#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07188)VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07188  
  If `pVideoProfile->videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the `pNext` chain of `pCapabilities` **must** include a [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265CapabilitiesKHR.html) structure
- [](#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-10263)VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-10263  
  If `pVideoProfile->videoCodecOperation` is `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR`, then the `pNext` chain of `pCapabilities` **must** include a [VkVideoEncodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1CapabilitiesKHR.html) structure

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-physicalDevice-parameter)VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-parameter)VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-parameter  
  `pVideoProfile` **must** be a valid pointer to a valid [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html) structure
- [](#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pCapabilities-parameter)VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pCapabilities-parameter  
  `pCapabilities` **must** be a valid pointer to a [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_VIDEO_PROFILE_OPERATION_NOT_SUPPORTED_KHR`
- `VK_ERROR_VIDEO_PROFILE_FORMAT_NOT_SUPPORTED_KHR`
- `VK_ERROR_VIDEO_PICTURE_LAYOUT_NOT_SUPPORTED_KHR`
- `VK_ERROR_VIDEO_PROFILE_CODEC_NOT_SUPPORTED_KHR`

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html), [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceVideoCapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0