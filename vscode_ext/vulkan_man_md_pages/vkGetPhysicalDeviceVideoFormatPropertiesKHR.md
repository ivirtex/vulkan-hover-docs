# vkGetPhysicalDeviceVideoFormatPropertiesKHR(3) Manual Page

## Name

vkGetPhysicalDeviceVideoFormatPropertiesKHR - Query supported video
decode and encode image formats and capabilities



## <a href="#_c_specification" class="anchor"></a>C Specification

To enumerate the supported output, input and DPB image formats and
corresponding capabilities for a specific video profile, call:

``` c
// Provided by VK_KHR_video_queue
VkResult vkGetPhysicalDeviceVideoFormatPropertiesKHR(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceVideoFormatInfoKHR*   pVideoFormatInfo,
    uint32_t*                                   pVideoFormatPropertyCount,
    VkVideoFormatPropertiesKHR*                 pVideoFormatProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device from which to query the video
  format properties.

- `pVideoFormatInfo` is a pointer to a
  [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)
  structure specifying the usage and video profiles for which supported
  image formats and capabilities are returned.

- `pVideoFormatPropertyCount` is a pointer to an integer related to the
  number of video format properties available or queried, as described
  below.

- `pVideoFormatProperties` is a pointer to an array of
  [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html)
  structures in which supported image formats and capabilities are
  returned.

## <a href="#_description" class="anchor"></a>Description

If `pVideoFormatProperties` is `NULL`, then the number of video format
properties supported for the given `physicalDevice` is returned in
`pVideoFormatPropertyCount`. Otherwise, `pVideoFormatPropertyCount`
**must** point to a variable set by the user to the number of elements
in the `pVideoFormatProperties` array, and on return the variable is
overwritten with the number of values actually written to
`pVideoFormatProperties`. If the value of `pVideoFormatPropertyCount` is
less than the number of video format properties supported, at most
`pVideoFormatPropertyCount` values will be written to
`pVideoFormatProperties`, and `VK_INCOMPLETE` will be returned instead
of `VK_SUCCESS`, to indicate that not all the available values were
returned.

Video format properties are always queried with respect to a specific
set of video profiles. These are specified by chaining the
[VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html) structure to
`pVideoFormatInfo`.

For most use cases, the images are used by a single video session and a
single video profile is provided. For a use case such as video
transcoding, where a decode session output image **can** be used as
encode input in one or more encode sessions, multiple video profiles
corresponding to the video sessions that will share the image **must**
be provided.

If any of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profiles"
target="_blank" rel="noopener">video profiles</a> specified via
[VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)::`pProfiles`
are not supported, then this command returns one of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profile-error-codes"
target="_blank" rel="noopener">video-profile-specific error codes</a>.
Furthermore, if
[VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)::`imageUsage`
includes any image usage flags not supported by the specified video
profiles, then this command returns
`VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR`.

This command also returns `VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR` if
[VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)::`imageUsage`
does not include the appropriate flags as dictated by the decode
capability flags returned in
[VkVideoDecodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeCapabilitiesKHR.html)::`flags`
for any of the profiles specified in the
[VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html) structure
provided in the `pNext` chain of `pVideoFormatInfo`.

If the decode capability flags include
`VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_COINCIDE_BIT_KHR` but not
`VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_DISTINCT_BIT_KHR`, then in
order to query video format properties for decode DPB and output usage,
[VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)::`imageUsage`
**must** include both `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR` and
`VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`. Otherwise, the call will fail
with `VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR`.

If the decode capability flags include
`VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_DISTINCT_BIT_KHR` but not
`VK_VIDEO_DECODE_CAPABILITY_DPB_AND_OUTPUT_COINCIDE_BIT_KHR`, then in
order to query video format properties for decode DPB usage,
[VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)::`imageUsage`
**must** include `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`, but not
`VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`. Otherwise, the call will fail
with `VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR`. Similarly, to query video
format properties for decode output usage,
[VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)::`imageUsage`
**must** include `VK_IMAGE_USAGE_VIDEO_DECODE_DST_BIT_KHR`, but not
`VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`. Otherwise, the call will fail
with `VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR`.

The `imageUsage` member of the
[VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)
structure specifies the expected video usage flags that the returned
video formats **must** support. Correspondingly, the `imageUsageFlags`
member of each
[VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html) structure
returned will contain at least the same set of image usage flags.

If the implementation supports using video input, output, or DPB images
of a particular format in operations other than video decode/encode then
the `imageUsageFlags` member of the corresponding
[VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html) structure
returned will include additional image usage flags indicating that.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>For most use cases, only decode or encode related usage flags are
going to be specified. For a use case such as transcode, if the image
were to be shared between decode and encode session(s), then both decode
and encode related usage flags <strong>can</strong> be set.</p></td>
</tr>
</tbody>
</table>

Multiple `VkVideoFormatPropertiesKHR` entries **may** be returned with
the same `format` member with different `componentMapping`, `imageType`,
or `imageTiling` values, as described later.

In addition, a different set of `VkVideoFormatPropertiesKHR` entries
**may** be returned depending on the `imageUsage` member of the
`VkPhysicalDeviceVideoFormatInfoKHR` structure, even for the same set of
video profiles, for example, based on whether encode input, encode DPB,
decode output, and/or decode DPB usage is requested.

The application **can** select the parameters returned in the
`VkVideoFormatPropertiesKHR` entries and use compatible parameters when
creating the input, output, and DPB images. The implementation will
report all image creation and usage flags that are valid for images used
with the requested video profiles but applications **should** create
images only with those that are necessary for the particular use case.

Before creating an image, the application **can** obtain the complete
set of supported image format features by calling
[vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
using parameters derived from the members of one of the reported
`VkVideoFormatPropertiesKHR` entries and adding the same
[VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html) structure to
the `pNext` chain of
[VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html).

The following applies to all `VkVideoFormatPropertiesKHR` entries
returned by `vkGetPhysicalDeviceVideoFormatPropertiesKHR`:

- [vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2.html)
  **must** succeed when called with
  `VkVideoFormatPropertiesKHR`::`format`

- If `VkVideoFormatPropertiesKHR`::`imageTiling` is
  `VK_IMAGE_TILING_OPTIMAL`, then the `optimalTilingFeatures` returned
  by
  [vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2.html)
  **must** include all format features required by the image usage flags
  reported in `VkVideoFormatPropertiesKHR`::`imageUsageFlags` for the
  format, as indicated in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#format-feature-dependent-usage-flags"
  target="_blank" rel="noopener">Format Feature Dependent Usage Flags</a>
  section.

- If `VkVideoFormatPropertiesKHR`::`imageTiling` is
  `VK_IMAGE_TILING_LINEAR`, then the `linearTilingFeatures` returned by
  [vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2.html)
  **must** include all format features required by the image usage flags
  reported in `VkVideoFormatPropertiesKHR`::`imageUsageFlags` for the
  format, as indicated in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#format-feature-dependent-usage-flags"
  target="_blank" rel="noopener">Format Feature Dependent Usage Flags</a>
  section.

- [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
  **must** succeed when called with a
  [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)
  structure containing the following information:

  - The `pNext` chain including the same
    [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html)
    structure used to call
    `vkGetPhysicalDeviceVideoFormatPropertiesKHR`.

  - `format` set to the value of `VkVideoFormatPropertiesKHR`::`format`.

  - `type` set to the value of
    `VkVideoFormatPropertiesKHR`::`imageType`.

  - `tiling` set to the value of
    `VkVideoFormatPropertiesKHR`::`imageTiling`.

  - `usage` set to the value of
    `VkVideoFormatPropertiesKHR`::`imageUsageFlags`.

  - `flags` set to the value of
    `VkVideoFormatPropertiesKHR`::`imageCreateFlags`.

The `componentMapping` member of `VkVideoFormatPropertiesKHR` defines
the ordering of the Y′C<sub>B</sub>C<sub>R</sub> color channels from the
perspective of the video codec operations specified in
[VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html). For
example, if the implementation produces video decode output with the
format `VK_FORMAT_G8_B8R8_2PLANE_420_UNORM` where the blue and red
chrominance channels are swapped then the `componentMapping` member of
the corresponding `VkVideoFormatPropertiesKHR` structure will have the
following member values:

``` c
components.r = VK_COMPONENT_SWIZZLE_B;        // Cb component
components.g = VK_COMPONENT_SWIZZLE_IDENTITY; // Y component
components.b = VK_COMPONENT_SWIZZLE_R;        // Cr component
components.a = VK_COMPONENT_SWIZZLE_IDENTITY; // unused, defaults to 1.0
```

Valid Usage

- <a href="#VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pNext-06812"
  id="VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pNext-06812"></a>
  VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pNext-06812  
  The `pNext` chain of `pVideoFormatInfo` **must** include a
  [VkVideoProfileListInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoProfileListInfoKHR.html) structure
  with `profileCount` greater than `0`

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pVideoFormatInfo-parameter"
  id="VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pVideoFormatInfo-parameter"></a>
  VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pVideoFormatInfo-parameter  
  `pVideoFormatInfo` **must** be a valid pointer to a valid
  [VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pVideoFormatPropertyCount-parameter"
  id="VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pVideoFormatPropertyCount-parameter"></a>
  VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pVideoFormatPropertyCount-parameter  
  `pVideoFormatPropertyCount` **must** be a valid pointer to a
  `uint32_t` value

- <a
  href="#VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pVideoFormatProperties-parameter"
  id="VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pVideoFormatProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceVideoFormatPropertiesKHR-pVideoFormatProperties-parameter  
  If the value referenced by `pVideoFormatPropertyCount` is not `0`, and
  `pVideoFormatProperties` is not `NULL`, `pVideoFormatProperties`
  **must** be a valid pointer to an array of `pVideoFormatPropertyCount`
  [VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html)
  structures

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_IMAGE_USAGE_NOT_SUPPORTED_KHR`

- `VK_ERROR_VIDEO_PROFILE_OPERATION_NOT_SUPPORTED_KHR`

- `VK_ERROR_VIDEO_PROFILE_FORMAT_NOT_SUPPORTED_KHR`

- `VK_ERROR_VIDEO_PICTURE_LAYOUT_NOT_SUPPORTED_KHR`

- `VK_ERROR_VIDEO_PROFILE_CODEC_NOT_SUPPORTED_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkPhysicalDeviceVideoFormatInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoFormatInfoKHR.html),
[VkVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoFormatPropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceVideoFormatPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
