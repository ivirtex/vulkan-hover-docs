# VkVideoEncodeCapabilitiesKHR(3) Manual Page

## Name

VkVideoEncodeCapabilitiesKHR - Structure describing general video encode
capabilities for a video profile



## <a href="#_c_specification" class="anchor"></a>C Specification

When calling
[vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
with `pVideoProfile->videoCodecOperation` specifying an encode
operation, the
[VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilitiesKHR.html)
structure **must** be included in the `pNext` chain of the
[VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html) structure to
retrieve capabilities specific to video encoding.

The `VkVideoEncodeCapabilitiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_queue
typedef struct VkVideoEncodeCapabilitiesKHR {
    VkStructureType                         sType;
    void*                                   pNext;
    VkVideoEncodeCapabilityFlagsKHR         flags;
    VkVideoEncodeRateControlModeFlagsKHR    rateControlModes;
    uint32_t                                maxRateControlLayers;
    uint64_t                                maxBitrate;
    uint32_t                                maxQualityLevels;
    VkExtent2D                              encodeInputPictureGranularity;
    VkVideoEncodeFeedbackFlagsKHR           supportedEncodeFeedbackFlags;
} VkVideoEncodeCapabilitiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkVideoEncodeCapabilityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilityFlagBitsKHR.html)
  describing supported encoding features.

- `rateControlModes` is a bitmask of
  [VkVideoEncodeRateControlModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlModeFlagBitsKHR.html)
  indicating supported <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-modes"
  target="_blank" rel="noopener">rate control modes</a>.

- `maxRateControlLayers` indicates the maximum number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-layers"
  target="_blank" rel="noopener">rate control layers</a> supported.

- `maxBitrate` indicates the maximum supported bitrate.

- `maxQualityLevels` indicates the number of discrete <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-quality-level"
  target="_blank" rel="noopener">video encode quality levels</a>
  supported. Implementations **must** support at least one quality
  level.

- `encodeInputPictureGranularity` indicates the granularity at which <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-input-picture"
  target="_blank" rel="noopener">encode input picture</a> data is
  encoded and **may** indicate a texel granularity up to the size of the
  codec-specific coding block size. This capability does not impose any
  valid usage constraints on the application, however, depending on the
  contents of the encode input picture, it **may** have effects on the
  encoded bitstream, as described in more detail below.

- `supportedEncodeFeedbackFlags` is a bitmask of
  [VkVideoEncodeFeedbackFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeFeedbackFlagBitsKHR.html)
  values specifying the supported flags for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-video-encode-feedback"
  target="_blank" rel="noopener">video encode feedback queries</a>.

## <a href="#_description" class="anchor"></a>Description

Implementations **must** include support for at least
`VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_BUFFER_OFFSET_BIT_KHR` and
`VK_VIDEO_ENCODE_FEEDBACK_BITSTREAM_BYTES_WRITTEN_BIT_KHR` in
`supportedEncodeFeedbackFlags`.

`encodeInputPictureGranularity` provides information about the way <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-input-picture"
target="_blank" rel="noopener">encode input picture</a> data is used as
input to video encode operations. In particular, some implementations
**may** not be able to limit the set of texels used to encode the output
video bitstream to the image subregion specified in the
[VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoPictureResourceInfoKHR.html)
structure corresponding to the encode input picture (i.e. to the
resolution of the image data to encode specified in its `codedExtent`
member).

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>For example, the application requests the coded extent to be
1920x1080, but the implementation is only able to source the encode
input picture data at the granularity of the codec-specific coding block
size which is 16x16 pixels (or as otherwise indicated in
<code>encodeInputPictureGranularity</code>). In this example, the
content is horizontally aligned with the coding block size, but not
vertically aligned with it. Thus encoding of the last row of coding
blocks will be impacted by the contents of the input image at texel rows
1080 to 1087 (the latter being the next row which is vertically aligned
with the coding block size, assuming a zero-based texel row
index).</p></td>
</tr>
</tbody>
</table>

If `codedExtent` rounded up to the next integer multiple of
`encodeInputPictureGranularity` is greater than the extent of the image
subresource specified for the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-input-picture"
target="_blank" rel="noopener">encode input picture</a>, then the texel
values corresponding to texel coordinates outside of the bounds of the
image subresource **may** be undefined. However, implementations
**should** use well-defined default values for such texels in order to
maximize the encoding efficiency for the last coding block row/column,
and/or to ensure consistent encoding results across repeated encoding of
the same input content. Nonetheless, the values used for such texels
**must** not have an effect on whether the video encode operation
produces a compliant bitstream, and **must** not have any other effects
on the encoded picture data beyond what **may** otherwise result from
using these texel values as input to any compression algorithm, as
defined in the used video compression standard.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>While not required, it is generally a good practice for applications
to make sure that the image subresource used for the encode input
picture has an extent that is an integer multiple of the codec-specific
coding block size (or at least
<code>encodeInputPictureGranularity</code>) and that this padding area
is filled with known values in order to improve encoding efficiency,
portability, and reproducibility.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeCapabilitiesKHR-sType-sType"
  id="VUID-VkVideoEncodeCapabilitiesKHR-sType-sType"></a>
  VUID-VkVideoEncodeCapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_CAPABILITIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilityFlagsKHR.html),
[VkVideoEncodeFeedbackFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeFeedbackFlagsKHR.html),
[VkVideoEncodeRateControlModeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlModeFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeCapabilitiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
