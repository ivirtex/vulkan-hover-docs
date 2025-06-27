# VkVideoEncodeRateControlInfoKHR(3) Manual Page

## Name

VkVideoEncodeRateControlInfoKHR - Structure to set encode stream rate
control parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeRateControlInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_queue
typedef struct VkVideoEncodeRateControlInfoKHR {
    VkStructureType                                sType;
    const void*                                    pNext;
    VkVideoEncodeRateControlFlagsKHR               flags;
    VkVideoEncodeRateControlModeFlagBitsKHR        rateControlMode;
    uint32_t                                       layerCount;
    const VkVideoEncodeRateControlLayerInfoKHR*    pLayers;
    uint32_t                                       virtualBufferSizeInMs;
    uint32_t                                       initialVirtualBufferSizeInMs;
} VkVideoEncodeRateControlInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `rateControlMode` is a
  [VkVideoEncodeRateControlModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlModeFlagBitsKHR.html)
  value specifying the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-modes"
  target="_blank" rel="noopener">rate control mode</a>.

- `layerCount` specifies the number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-layers"
  target="_blank" rel="noopener">rate control layers</a> to use.

- `pLayers` is a pointer to an array of `layerCount`
  [VkVideoEncodeRateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlLayerInfoKHR.html)
  structures, each specifying the rate control configuration of the
  corresponding rate control layer.

- `virtualBufferSizeInMs` is the size in milliseconds of the virtual
  buffer used by the implementation’s rate control algorithm for the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-leaky-bucket-model"
  target="_blank" rel="noopener">leaky bucket model</a>, with respect to
  the average bitrate of the stream calculated by summing the values of
  the `averageBitrate` members of the elements of the `pLayers` array.

- `initialVirtualBufferSizeInMs` is the initial occupancy in
  milliseconds of the virtual buffer used by the implementation’s rate
  control algorithm for the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-leaky-bucket-model"
  target="_blank" rel="noopener">leaky bucket model</a>.

## <a href="#_description" class="anchor"></a>Description

If `layerCount` is zero then the values of `virtualBufferSizeInMs` and
`initialVirtualBufferSizeInMs` are ignored.

This structure **can** be specified in the following places:

- In the `pNext` chain of
  [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingInfoKHR.html) to specify
  the current rate control state expected to be configured when
  beginning a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-coding-scope"
  target="_blank" rel="noopener">video coding scope</a>.

- In the `pNext` chain of
  [VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlInfoKHR.html) to
  change the rate control configuration of the bound video session.

Including this structure in the `pNext` chain of
[VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlInfoKHR.html) and
including `VK_VIDEO_CODING_CONTROL_ENCODE_RATE_CONTROL_BIT_KHR` in
[VkVideoCodingControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlInfoKHR.html)::`flags`
enables updating the rate control configuration of the bound video
session. This replaces the entire rate control configuration of the
bound video session and **may** reset the state of all enabled rate
control layers to an initial state according to the codec-specific rate
control semantics defined in the corresponding sections listed below.

When `layerCount` is greater than one, multiple <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control-layers"
target="_blank" rel="noopener">rate control layers</a> are configured,
and each rate control layer is applied to the corresponding video coding
layer identified by the index of the corresponding element of `pLayer`.

- If the video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then this index
  specifies the H.264 temporal layer ID of the video coding layer the
  rate control layer is applied to.

- If the video session was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then this index
  specifies the H.265 temporal ID of the video coding layer the rate
  control layer is applied to.

Additional structures providing codec-specific rate control parameters
**can** be included in the `pNext` chain of
`VkVideoCodingControlInfoKHR` depending on the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profiles"
target="_blank" rel="noopener">video profile</a> the bound video session
was created. For further details see:

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-coding-control"
  target="_blank" rel="noopener">Video Coding Control</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-rate-control"
  target="_blank" rel="noopener">H.264 Encode Rate Control</a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-rate-control"
  target="_blank" rel="noopener">H.265 Encode Rate Control</a>

The new rate control configuration takes effect when the corresponding
[vkCmdControlVideoCodingKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdControlVideoCodingKHR.html) is
executed on the device, and only impacts video encode operations that
follow in execution order.

Valid Usage

- <a href="#VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-08248"
  id="VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-08248"></a>
  VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-08248  
  If `rateControlMode` is
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DEFAULT_KHR` or
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DISABLED_BIT_KHR`, then
  `layerCount` **must** be `0`

- <a href="#VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-08275"
  id="VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-08275"></a>
  VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-08275  
  If `rateControlMode` is
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_CBR_BIT_KHR` or
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_VBR_BIT_KHR`, then `layerCount`
  **must** be greater than `0`

- <a href="#VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-08244"
  id="VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-08244"></a>
  VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-08244  
  If `rateControlMode` is not
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_DEFAULT_KHR`, then it **must**
  specify one of the bits included in
  [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilitiesKHR.html)::`rateControlModes`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile

- <a href="#VUID-VkVideoEncodeRateControlInfoKHR-layerCount-08245"
  id="VUID-VkVideoEncodeRateControlInfoKHR-layerCount-08245"></a>
  VUID-VkVideoEncodeRateControlInfoKHR-layerCount-08245  
  `layerCount` member **must** be less than or equal to
  [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilitiesKHR.html)::`maxRateControlLayers`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile

- <a href="#VUID-VkVideoEncodeRateControlInfoKHR-pLayers-08276"
  id="VUID-VkVideoEncodeRateControlInfoKHR-pLayers-08276"></a>
  VUID-VkVideoEncodeRateControlInfoKHR-pLayers-08276  
  For each element of `pLayers`, its `averageBitrate` member **must** be
  between `1` and
  [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilitiesKHR.html)::`maxBitrate`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile

- <a href="#VUID-VkVideoEncodeRateControlInfoKHR-pLayers-08277"
  id="VUID-VkVideoEncodeRateControlInfoKHR-pLayers-08277"></a>
  VUID-VkVideoEncodeRateControlInfoKHR-pLayers-08277  
  For each element of `pLayers`, its `maxBitrate` member **must** be
  between `1` and
  [VkVideoEncodeCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilitiesKHR.html)::`maxBitrate`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the used video profile

- <a href="#VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-08356"
  id="VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-08356"></a>
  VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-08356  
  If `rateControlMode` is
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_CBR_BIT_KHR`, then for each element
  of `pLayers`, its `averageBitrate` member **must** equal its
  `maxBitrate` member

- <a href="#VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-08278"
  id="VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-08278"></a>
  VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-08278  
  If `rateControlMode` is
  `VK_VIDEO_ENCODE_RATE_CONTROL_MODE_VBR_BIT_KHR`, then for each element
  of `pLayers`, its `averageBitrate` member **must** be less than or
  equal to its `maxBitrate` member

- <a href="#VUID-VkVideoEncodeRateControlInfoKHR-layerCount-08357"
  id="VUID-VkVideoEncodeRateControlInfoKHR-layerCount-08357"></a>
  VUID-VkVideoEncodeRateControlInfoKHR-layerCount-08357  
  If `layerCount` is not zero, then `virtualBufferSizeInMs` **must** be
  greater than zero

- <a href="#VUID-VkVideoEncodeRateControlInfoKHR-layerCount-08358"
  id="VUID-VkVideoEncodeRateControlInfoKHR-layerCount-08358"></a>
  VUID-VkVideoEncodeRateControlInfoKHR-layerCount-08358  
  If `layerCount` is not zero, then `initialVirtualBufferSizeInMs`
  **must** be less than `virtualBufferSizeInMs`

- <a
  href="#VUID-VkVideoEncodeRateControlInfoKHR-videoCodecOperation-07022"
  id="VUID-VkVideoEncodeRateControlInfoKHR-videoCodecOperation-07022"></a>
  VUID-VkVideoEncodeRateControlInfoKHR-videoCodecOperation-07022  
  If the `videoCodecOperation` of the used video profile is
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, the `pNext` chain this
  structure is included in also includes an instance of the
  [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlInfoKHR.html)
  structure, and `layerCount` is greater than `1`, then `layerCount`
  **must** equal
  [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlInfoKHR.html)::`temporalLayerCount`

- <a
  href="#VUID-VkVideoEncodeRateControlInfoKHR-videoCodecOperation-07025"
  id="VUID-VkVideoEncodeRateControlInfoKHR-videoCodecOperation-07025"></a>
  VUID-VkVideoEncodeRateControlInfoKHR-videoCodecOperation-07025  
  If the `videoCodecOperation` of the used video profile is
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, the `pNext` chain this
  structure is included in also includes an instance of the
  [VkVideoEncodeH265RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlInfoKHR.html)
  structure, and `layerCount` is greater than `1`, then `layerCount`
  **must** equal
  [VkVideoEncodeH265RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlInfoKHR.html)::`subLayerCount`

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeRateControlInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeRateControlInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeRateControlInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_RATE_CONTROL_INFO_KHR`

- <a href="#VUID-VkVideoEncodeRateControlInfoKHR-flags-zerobitmask"
  id="VUID-VkVideoEncodeRateControlInfoKHR-flags-zerobitmask"></a>
  VUID-VkVideoEncodeRateControlInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-parameter"
  id="VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-parameter"></a>
  VUID-VkVideoEncodeRateControlInfoKHR-rateControlMode-parameter  
  If `rateControlMode` is not `0`, `rateControlMode` **must** be a valid
  [VkVideoEncodeRateControlModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlModeFlagBitsKHR.html)
  value

- <a href="#VUID-VkVideoEncodeRateControlInfoKHR-pLayers-parameter"
  id="VUID-VkVideoEncodeRateControlInfoKHR-pLayers-parameter"></a>
  VUID-VkVideoEncodeRateControlInfoKHR-pLayers-parameter  
  If `layerCount` is not `0`, `pLayers` **must** be a valid pointer to
  an array of `layerCount` valid
  [VkVideoEncodeRateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlLayerInfoKHR.html)
  structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeRateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlFlagsKHR.html),
[VkVideoEncodeRateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlLayerInfoKHR.html),
[VkVideoEncodeRateControlModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlModeFlagBitsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeRateControlInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
