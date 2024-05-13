# VkVideoEncodeH265RateControlLayerInfoKHR(3) Manual Page

## Name

VkVideoEncodeH265RateControlLayerInfoKHR - Structure describing H.265
per-layer rate control parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeH265RateControlLayerInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265RateControlLayerInfoKHR {
    VkStructureType                  sType;
    const void*                      pNext;
    VkBool32                         useMinQp;
    VkVideoEncodeH265QpKHR           minQp;
    VkBool32                         useMaxQp;
    VkVideoEncodeH265QpKHR           maxQp;
    VkBool32                         useMaxFrameSize;
    VkVideoEncodeH265FrameSizeKHR    maxFrameSize;
} VkVideoEncodeH265RateControlLayerInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `useMinQp` indicates whether the QP values determined by rate control
  will be clamped to the lower bounds on the QP values specified in
  `minQp`.

- `minQp` specifies the lower bounds on the QP values, for each picture
  type, that the implementation’s rate control algorithm will use when
  `useMinQp` is set to `VK_TRUE`.

- `useMaxQp` indicates whether the QP values determined by rate control
  will be clamped to the upper bounds on the QP values specified in
  `maxQp`.

- `maxQp` specifies the upper bounds on the QP values, for each picture
  type, that the implementation’s rate control algorithm will use when
  `useMaxQp` is set to `VK_TRUE`.

- `useMaxFrameSize` indicates whether the implementation’s rate control
  algorithm **should** use the values specified in `maxFrameSize` as the
  upper bounds on the encoded frame size for each picture type.

- `maxFrameSize` specifies the upper bounds on the encoded frame size,
  for each picture type, when `useMaxFrameSize` is set to `VK_TRUE`.

## <a href="#_description" class="anchor"></a>Description

When used, the values in `minQp` and `maxQp` guarantee that the
effective QP values used by the implementation will respect those lower
and upper bounds, respectively. However, limiting the range of QP values
that the implementation is able to use will also limit the capabilities
of the implementation’s rate control algorithm to comply to other
constraints. In particular, the implementation **may** not be able to
comply to the following:

- The average and/or peak <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-bitrate"
  target="_blank" rel="noopener">bitrate</a> values to be used for the
  encoded bitstream specified in the `averageBitrate` and `maxBitrate`
  members of the
  [VkVideoEncodeRateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlLayerInfoKHR.html)
  structure.

- The upper bounds on the encoded frame size, for each picture type,
  specified in the `maxFrameSize` member of
  `VkVideoEncodeH265RateControlLayerInfoKHR`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>In general, applications need to configure rate control parameters
appropriately in order to be able to get the desired rate control
behavior, as described in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control"
target="_blank" rel="noopener">Video Encode Rate Control</a>
section.</p></td>
</tr>
</tbody>
</table>

When an instance of this structure is included in the `pNext` chain of a
[VkVideoEncodeRateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlLayerInfoKHR.html)
structure specified in one of the elements of the `pLayers` array member
of the
[VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlInfoKHR.html)
structure passed to the
[vkCmdControlVideoCodingKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdControlVideoCodingKHR.html) command,
[VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlInfoKHR.html)::`flags`
includes `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_BIT_KHR`, and the
bound video session was created with the video codec operation
`VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, it specifies the
H.265-specific rate control parameters of the rate control layer
corresponding to that element of `pLayers`.

Valid Usage

- <a href="#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMinQp-08297"
  id="VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMinQp-08297"></a>
  VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMinQp-08297  
  If `useMinQp` is `VK_TRUE`, then the `qpI`, `qpP`, and `qpB` members
  of `minQp` **must** all be between
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`minQp`
  and
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxQp`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile

- <a href="#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMaxQp-08298"
  id="VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMaxQp-08298"></a>
  VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMaxQp-08298  
  If `useMaxQp` is `VK_TRUE`, then the `qpI`, `qpP`, and `qpB` members
  of `maxQp` **must** all be between
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`minQp`
  and
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`maxQp`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile

- <a href="#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMinQp-08299"
  id="VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMinQp-08299"></a>
  VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMinQp-08299  
  If `useMinQp` is `VK_TRUE` and
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile, does not include
  `VK_VIDEO_ENCODE_H265_CAPABILITY_PER_PICTURE_TYPE_MIN_MAX_QP_BIT_KHR`,
  then the `qpI`, `qpP`, and `qpB` members of `minQp` **must** all
  specify the same value

- <a href="#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMaxQp-08300"
  id="VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMaxQp-08300"></a>
  VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMaxQp-08300  
  If `useMaxQp` is `VK_TRUE` and
  [VkVideoEncodeH265CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilitiesKHR.html)::`flags`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile, does not include
  `VK_VIDEO_ENCODE_H265_CAPABILITY_PER_PICTURE_TYPE_MIN_MAX_QP_BIT_KHR`,
  then the `qpI`, `qpP`, and `qpB` members of `maxQp` **must** all
  specify the same value

- <a href="#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMinQp-08375"
  id="VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMinQp-08375"></a>
  VUID-VkVideoEncodeH265RateControlLayerInfoKHR-useMinQp-08375  
  If `useMinQp` and `useMaxQp` are both `VK_TRUE`, then the `qpI`,
  `qpP`, and `qpB` members of `minQp` **must** all be less than or equal
  to the respective members of `maxQp`

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeH265RateControlLayerInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeH265RateControlLayerInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_RATE_CONTROL_LAYER_INFO_KHR`

- <a href="#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-minQp-parameter"
  id="VUID-VkVideoEncodeH265RateControlLayerInfoKHR-minQp-parameter"></a>
  VUID-VkVideoEncodeH265RateControlLayerInfoKHR-minQp-parameter  
  `minQp` **must** be a valid
  [VkVideoEncodeH265QpKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265QpKHR.html) structure

- <a href="#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-maxQp-parameter"
  id="VUID-VkVideoEncodeH265RateControlLayerInfoKHR-maxQp-parameter"></a>
  VUID-VkVideoEncodeH265RateControlLayerInfoKHR-maxQp-parameter  
  `maxQp` **must** be a valid
  [VkVideoEncodeH265QpKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265QpKHR.html) structure

- <a
  href="#VUID-VkVideoEncodeH265RateControlLayerInfoKHR-maxFrameSize-parameter"
  id="VUID-VkVideoEncodeH265RateControlLayerInfoKHR-maxFrameSize-parameter"></a>
  VUID-VkVideoEncodeH265RateControlLayerInfoKHR-maxFrameSize-parameter  
  `maxFrameSize` **must** be a valid
  [VkVideoEncodeH265FrameSizeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265FrameSizeKHR.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h265.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeH265FrameSizeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265FrameSizeKHR.html),
[VkVideoEncodeH265QpKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265QpKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH265RateControlLayerInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
