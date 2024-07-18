# vkGetPhysicalDeviceVideoCapabilitiesKHR(3) Manual Page

## Name

vkGetPhysicalDeviceVideoCapabilitiesKHR - Query video coding
capabilities



## <a href="#_c_specification" class="anchor"></a>C Specification

To query video coding capabilities for a specific video profile, call:

``` c
// Provided by VK_KHR_video_queue
VkResult vkGetPhysicalDeviceVideoCapabilitiesKHR(
    VkPhysicalDevice                            physicalDevice,
    const VkVideoProfileInfoKHR*                pVideoProfile,
    VkVideoCapabilitiesKHR*                     pCapabilities);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device from which to query the video
  decode or encode capabilities.

- `pVideoProfile` is a pointer to a
  [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html) structure.

- `pCapabilities` is a pointer to a
  [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html) structure in
  which the capabilities are returned.

## <a href="#_description" class="anchor"></a>Description

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profiles"
target="_blank" rel="noopener">video profile</a> described by
`pVideoProfile` is supported by the implementation, then this command
returns `VK_SUCCESS` and `pCapabilities` is filled with the capabilities
supported with the specified video profile. Otherwise, one of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profile-error-codes"
target="_blank" rel="noopener">video-profile-specific error codes</a>
are returned.

Valid Usage

- <a
  href="#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07183"
  id="VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07183"></a>
  VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07183  
  If `pVideoProfile->videoCodecOperation` specifies a decode operation,
  then the `pNext` chain of `pCapabilities` **must** include a
  [VkVideoDecodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeCapabilitiesKHR.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07184"
  id="VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07184"></a>
  VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07184  
  If `pVideoProfile->videoCodecOperation` is
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then the `pNext` chain
  of `pCapabilities` **must** include a
  [VkVideoDecodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264CapabilitiesKHR.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07185"
  id="VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07185"></a>
  VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07185  
  If `pVideoProfile->videoCodecOperation` is
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then the `pNext` chain
  of `pCapabilities` **must** include a
  [VkVideoDecodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265CapabilitiesKHR.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-09257"
  id="VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-09257"></a>
  VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-09257  
  If `pVideoProfile->videoCodecOperation` is
  `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then the `pNext` chain
  of `pCapabilities` **must** include a
  [VkVideoDecodeAV1CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1CapabilitiesKHR.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07186"
  id="VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07186"></a>
  VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07186  
  If `pVideoProfile->videoCodecOperation` specifies an encode operation,
  then the `pNext` chain of `pCapabilities` **must** include a
  [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilitiesKHR.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07187"
  id="VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07187"></a>
  VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07187  
  If `pVideoProfile->videoCodecOperation` is
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then the `pNext` chain
  of `pCapabilities` **must** include a
  [VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07188"
  id="VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07188"></a>
  VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-07188  
  If `pVideoProfile->videoCodecOperation` is
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then the `pNext` chain
  of `pCapabilities` **must** include a
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)
  structure

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-parameter"
  id="VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-parameter"></a>
  VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pVideoProfile-parameter  
  `pVideoProfile` **must** be a valid pointer to a valid
  [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html) structure

- <a
  href="#VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pCapabilities-parameter"
  id="VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pCapabilities-parameter"></a>
  VUID-vkGetPhysicalDeviceVideoCapabilitiesKHR-pCapabilities-parameter  
  `pCapabilities` **must** be a valid pointer to a
  [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html) structure

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

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html),
[VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceVideoCapabilitiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
