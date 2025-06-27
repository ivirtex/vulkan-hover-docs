# VkVideoEncodeH264GopRemainingFrameInfoKHR(3) Manual Page

## Name

VkVideoEncodeH264GopRemainingFrameInfoKHR - Structure specifying H.264
encode rate control GOP remaining frame counts



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeH264GopRemainingFrameInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_h264
typedef struct VkVideoEncodeH264GopRemainingFrameInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           useGopRemainingFrames;
    uint32_t           gopRemainingI;
    uint32_t           gopRemainingP;
    uint32_t           gopRemainingB;
} VkVideoEncodeH264GopRemainingFrameInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `useGopRemainingFrames` indicates whether the implementation’s rate
  control algorithm **should** use the values specified in
  `gopRemainingI`, `gopRemainingP`, and `gopRemainingB`. If
  `useGopRemainingFrames` is `VK_FALSE`, then the values of
  `gopRemainingI`, `gopRemainingP`, and `gopRemainingB` are ignored.

- `gopRemainingI` specifies the number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-i-pic"
  target="_blank" rel="noopener">I frames</a> the implementation’s rate
  control algorithm **should** assume to be remaining in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-gop"
  target="_blank" rel="noopener">GOP</a> prior to executing the video
  encode operation.

- `gopRemainingP` specifies the number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-p-pic"
  target="_blank" rel="noopener">P frames</a> the implementation’s rate
  control algorithm **should** assume to be remaining in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-gop"
  target="_blank" rel="noopener">GOP</a> prior to executing the video
  encode operation.

- `gopRemainingB` specifies the number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-b-pic"
  target="_blank" rel="noopener">B frames</a> the implementation’s rate
  control algorithm **should** assume to be remaining in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-gop"
  target="_blank" rel="noopener">GOP</a> prior to executing the video
  encode operation.

## <a href="#_description" class="anchor"></a>Description

Setting `useGopRemainingFrames` to `VK_TRUE` and including this
structure in the `pNext` chain of
[VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingInfoKHR.html) is only
mandatory if the
[VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`requiresGopRemainingFrames`
reported for the used <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profiles"
target="_blank" rel="noopener">video profile</a> is `VK_TRUE`. However,
implementations **may** use these remaining frame counts, when
specified, even when it is not required. In particular, when the
application does not use a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-regular-gop"
target="_blank" rel="noopener">regular GOP structure</a>, these values
**may** provide additional guidance for the implementation’s rate
control algorithm.

The
[VkVideoEncodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilitiesKHR.html)::`prefersGopRemainingFrames`
capability is also used to indicate that the implementation’s rate
control algorithm **may** operate more accurately if the application
specifies the remaining frame counts using this structure.

As with other rate control guidance values, if the effective order and
number of frames encoded by the application are not in line with the
remaining frame counts specified in this structure at any given point,
then the behavior of the implementation’s rate control algorithm **may**
deviate from the one expected by the application.

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeH264GopRemainingFrameInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeH264GopRemainingFrameInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeH264GopRemainingFrameInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_GOP_REMAINING_FRAME_INFO_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h264.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH264GopRemainingFrameInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
