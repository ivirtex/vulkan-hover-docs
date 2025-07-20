# vkGetPhysicalDeviceVideoFormatPropertiesKHR(3) Manual Page

## Name

vkGetPhysicalDeviceVideoFormatPropertiesKHR - Query supported video decode and encode image formats and capabilities



## [](#_c_specification)C Specification

To enumerate the supported video formats and corresponding capabilities for a specific video profile, call:

```c++
// Provided by VK_KHR_video_queue
VkResult vkGetPhysicalDeviceVideoFormatPropertiesKHR(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceVideoFormatInfoKHR*   pVideoFormatInfo,
    uint32_t*                                   pVideoFormatPropertyCount,
    VkVideoFormatPropertiesKHR*                 pVideoFormatProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device from which to query the video format properties.
- `pVideoFormatInfo` is a pointer to a [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html) structure specifying the usage and video profiles for which supported image formats and capabilities are returned.
- `pVideoFormatPropertyCount` is a pointer to an integer related to the number of video format properties available or queried, as described below.
- `pVideoFormatProperties` is a pointer to an array of [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html) structures in which supported image formats and capabilities are returned.

## [](#_description)Description

If `pVideoFormatProperties` is `NULL`, then the number of video format properties supported for the given `physicalDevice` is returned in `pVideoFormatPropertyCount`. Otherwise, `pVideoFormatPropertyCount` **must** point to a variable set by the application to the number of elements in the `pVideoFormatProperties` array, and on return the variable is overwritten with the number of values actually written to `pVideoFormatProperties`. If the value of `pVideoFormatPropertyCount` is less than the number of video format properties supported, at most `pVideoFormatPropertyCount` values will be written to `pVideoFormatProperties`, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the available values were returned.

Video format properties are always queried with respect to a specific set of video profiles. These are specified by chaining the [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html) structure to `pVideoFormatInfo`.

For most use cases, the images are used by a single video session and a single video profile is provided. For a use case such as video transcoding, where a decode session output image **can** be used as encode input in one or more encode sessions, multiple video profiles corresponding to the video sessions that will share the image **must** be provided.

If any of the [video profiles](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-profiles) specified via [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html)::`pProfiles` are not supported, then this command returns one of the [video-profile-specific error codes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-profile-error-codes). Furthermore, if [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)::`imageUsage` includes any image usage flags not supported by the specified video profiles, then this command returns `VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR`.

This command also returns `VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR` if [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)::`imageUsage` does not include the appropriate flags as dictated by the decode capability flags returned in [VkVideoDecodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeCapabilitiesKHR.html)::`flags` for any of the profiles specified in the [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html) structure provided in the `pNext` chain of `pVideoFormatInfo`.

If the decode capability flags include `VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_COINCIDE_BIT_KHR` but not `VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_DISTINCT_BIT_KHR`, then in order to query video format properties for decode DPB and output usage, [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)::`imageUsage` **must** include both `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR` and `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`. Otherwise, the call will fail with `VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR`.

If the decode capability flags include `VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_DISTINCT_BIT_KHR` but not `VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_COINCIDE_BIT_KHR`, then in order to query video format properties for decode DPB usage, [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)::`imageUsage` **must** include `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`, but not `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`. Otherwise, the call will fail with `VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR`. Similarly, to query video format properties for decode output usage, [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)::`imageUsage` **must** include `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`, but not `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`. Otherwise, the call will fail with `VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR`.

The `imageUsage` member of the [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html) structure specifies the expected video usage flags that the returned video formats **must** support. Correspondingly, the `imageUsageFlags` member of each [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html) structure returned will contain at least the same set of image usage flags.

If the implementation supports using images of a particular format in operations other than video decode/encode then the `imageUsageFlags` member of the corresponding [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html) structure returned will include additional image usage flags indicating that.

Note

For most use cases, only decode or encode related usage flags are going to be specified. For a use case such as transcode, if the image were to be shared between decode and encode session(s), then both decode and encode related usage flags **can** be set.

Multiple `VkVideoFormatPropertiesKHR` entries **may** be returned with the same `format` member with different `componentMapping`, `imageType`, or `imageTiling` values, as described later.

If [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)::`imageUsage` includes `VK_IMAGE_USAGE_VIDEO_ENCODE_QUANTIZATION_DELTA_MAP_BIT_KHR` or `VK_IMAGE_USAGE_VIDEO_ENCODE_EMPHASIS_MAP_BIT_KHR`, multiple `VkVideoFormatPropertiesKHR` entries **may** be returned with the same `format`, `componentMapping`, `imageType`, and `imageTiling` member values, but different `quantizationMapTexelSize` returned in the [VkVideoFormatQuantizationMapPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatQuantizationMapPropertiesKHR.html) structure, if one is included in the [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html)::`pNext` chain, when the queried [quantization map type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quantization-map) supports multiple distinct [quantization map texel sizes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-quantization-map-texel-size).

In addition, a different set of `VkVideoFormatPropertiesKHR` entries **may** be returned depending on the `imageUsage` member of the `VkPhysicalDeviceVideoFormatInfoKHR` structure, even for the same set of video profiles, for example, based on whether encode input, encode DPB, decode output, and/or decode DPB usage is requested.

The application **can** select the parameters returned in the `VkVideoFormatPropertiesKHR` entries and use compatible parameters when creating the input, output, and DPB images. The implementation will report all image creation and usage flags that are valid for images used with the requested video profiles but applications **should** create images only with those that are necessary for the particular use case.

Before creating an image, the application **can** obtain the complete set of supported image format features by calling [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html) using parameters derived from the members of one of the reported `VkVideoFormatPropertiesKHR` entries and adding the same [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html) structure to the `pNext` chain of [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html).

The following applies to all `VkVideoFormatPropertiesKHR` entries returned by `vkGetPhysicalDeviceVideoFormatPropertiesKHR`:

- [vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFormatProperties2.html) **must** succeed when called with `VkVideoFormatPropertiesKHR`::`format`
- If `VkVideoFormatPropertiesKHR`::`imageTiling` is `VK_IMAGE_TILING_OPTIMAL`, then the `optimalTilingFeatures` returned by [vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFormatProperties2.html) **must** include all format features required by the image usage flags reported in `VkVideoFormatPropertiesKHR`::`imageUsageFlags` for the format, as indicated in the [Format Feature Dependent Usage Flags](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#format-feature-dependent-usage-flags) section.
- If `VkVideoFormatPropertiesKHR`::`imageTiling` is `VK_IMAGE_TILING_LINEAR`, then the `linearTilingFeatures` returned by [vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFormatProperties2.html) **must** include all format features required by the image usage flags reported in `VkVideoFormatPropertiesKHR`::`imageUsageFlags` for the format, as indicated in the [Format Feature Dependent Usage Flags](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#format-feature-dependent-usage-flags) section.
- [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html) **must** succeed when called with a [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html) structure containing the following information:
  
  - The `pNext` chain including the same [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html) structure used to call `vkGetPhysicalDeviceVideoFormatPropertiesKHR`.
  - `format` set to the value of `VkVideoFormatPropertiesKHR`::`format`.
  - `type` set to the value of `VkVideoFormatPropertiesKHR`::`imageType`.
  - `tiling` set to the value of `VkVideoFormatPropertiesKHR`::`imageTiling`.
  - `usage` set to the value of `VkVideoFormatPropertiesKHR`::`imageUsageFlags`.
  - `flags` set to the value of `VkVideoFormatPropertiesKHR`::`imageCreateFlags`.

The `componentMapping` member of `VkVideoFormatPropertiesKHR` defines the ordering of the Y′CBCR color channels from the perspective of the video codec operations specified in [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html). For example, if the implementation produces video decode output with the format `VK_FORMAT_G8_B8R8_2PLANE_420_UNORM` where the blue and red chrominance channels are swapped then the `componentMapping` member of the corresponding `VkVideoFormatPropertiesKHR` structure will have the following member values:

```c++
components.r = VK_COMPONENT_SWIZZLE_B;        // Cb component
components.g = VK_COMPONENT_SWIZZLE_IDENTITY; // Y component
components.b = VK_COMPONENT_SWIZZLE_R;        // Cr component
components.a = VK_COMPONENT_SWIZZLE_IDENTITY; // unused, defaults to 1.0
```

Valid Usage

- [](#VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pNext-06812)VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pNext-06812  
  The `pNext` chain of `pVideoFormatInfo` **must** include a [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileListInfoKHR.html) structure with `profileCount` greater than `0`

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-physicalDevice-parameter)VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pVideoFormatInfo-parameter)VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pVideoFormatInfo-parameter  
  `pVideoFormatInfo` **must** be a valid pointer to a valid [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html) structure
- [](#VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pVideoFormatPropertyCount-parameter)VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pVideoFormatPropertyCount-parameter  
  `pVideoFormatPropertyCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pVideoFormatProperties-parameter)VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pVideoFormatProperties-parameter  
  If the value referenced by `pVideoFormatPropertyCount` is not `0`, and `pVideoFormatProperties` is not `NULL`, `pVideoFormatProperties` **must** be a valid pointer to an array of `pVideoFormatPropertyCount` [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html) structures

Return Codes

On success, this command returns

- `VK_INCOMPLETE`
- `VK_SUCCESS`

On failure, this command returns

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

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html), [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatPropertiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceVideoFormatPropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0