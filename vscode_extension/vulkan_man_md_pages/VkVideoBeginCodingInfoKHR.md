# VkVideoBeginCodingInfoKHR(3) Manual Page

## Name

VkVideoBeginCodingInfoKHR - Structure specifying video coding scope
begin information



## <a href="#_c_specification" class="anchor"></a>C Specification

The [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingInfoKHR.html)
structure is defined as:

``` c
// Provided by VK_KHR_video_queue
typedef struct VkVideoBeginCodingInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkVideoBeginCodingFlagsKHR            flags;
    VkVideoSessionKHR                     videoSession;
    VkVideoSessionParametersKHR           videoSessionParameters;
    uint32_t                              referenceSlotCount;
    const VkVideoReferenceSlotInfoKHR*    pReferenceSlots;
} VkVideoBeginCodingInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `videoSession` is the video session object to be bound for the
  processing of the video commands.

- `videoSessionParameters` is `VK_NULL_HANDLE` or a handle of a
  [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html) object
  to be used for the processing of the video commands. If
  `VK_NULL_HANDLE`, then no video session parameters object is bound for
  the duration of the video coding scope.

- `referenceSlotCount` is the number of elements in the
  `pReferenceSlots` array.

- `pReferenceSlots` is a pointer to an array of
  [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)
  structures specifying the information used to determine the set of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#bound-reference-picture-resources"
  target="_blank" rel="noopener">bound reference picture resources</a>
  for the video coding scope and their initial association with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slot</a> indices.

## <a href="#_description" class="anchor"></a>Description

Limiting values are defined below that are referenced by the relevant
valid usage statements of this structure.

- Let `VkOffset2D codedOffsetGranularity` be the minimum alignment
  requirement for the coded offset of video picture resources. Unless
  otherwise defined, the value of the `x` and `y` members of
  `codedOffsetGranularity` are `0`.

  - If `videoSession` was created with an <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-profile"
    target="_blank" rel="noopener">H.264 decode profile</a> with a
    [VkVideoDecodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264ProfileInfoKHR.html)::`pictureLayout`
    of
    `VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_SEPARATE_PLANES_BIT_KHR`,
    then `codedOffsetGranularity` is equal to
    [VkVideoDecodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264CapabilitiesKHR.html)::`fieldOffsetGranularity`,
    as returned by
    [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
    for that video profile.

Valid Usage

- <a href="#VUID-VkVideoBeginCodingInfoKHR-videoSession-07237"
  id="VUID-VkVideoBeginCodingInfoKHR-videoSession-07237"></a>
  VUID-VkVideoBeginCodingInfoKHR-videoSession-07237  
  `videoSession` **must** have memory bound to all of its memory
  bindings returned by
  [vkGetVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetVideoSessionMemoryRequirementsKHR.html)
  for `videoSession`

- <a href="#VUID-VkVideoBeginCodingInfoKHR-slotIndex-04856"
  id="VUID-VkVideoBeginCodingInfoKHR-slotIndex-04856"></a>
  VUID-VkVideoBeginCodingInfoKHR-slotIndex-04856  
  Each non-negative
  [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)::`slotIndex`
  specified in the elements of `pReferenceSlots` **must** be less than
  the
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`maxDpbSlots`
  specified when `videoSession` was created

- <a href="#VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07238"
  id="VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07238"></a>
  VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07238  
  Each video picture resource corresponding to any non-`NULL`
  `pPictureResource` member specified in the elements of
  `pReferenceSlots` **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resource-uniqueness"
  target="_blank" rel="noopener">unique</a> within `pReferenceSlots`

- <a href="#VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07240"
  id="VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07240"></a>
  VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07240  
  If the `pPictureResource` member of any element of `pReferenceSlots`
  is not `NULL`, then the image view specified in
  `pPictureResource->imageViewBinding` for that element **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profile-compatibility"
  target="_blank" rel="noopener">compatible</a> with the video profile
  `videoSession` was created with

- <a href="#VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07241"
  id="VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07241"></a>
  VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07241  
  If the `pPictureResource` member of any element of `pReferenceSlots`
  is not `NULL`, then the format of the image view specified in
  `pPictureResource->imageViewBinding` for that element **must** match
  the
  [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`referencePictureFormat`
  `videoSession` was created with

- <a href="#VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07242"
  id="VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07242"></a>
  VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07242  
  If the `pPictureResource` member of any element of `pReferenceSlots`
  is not `NULL`, then its `codedOffset` member **must** be an integer
  multiple of `codedOffsetGranularity`

- <a href="#VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07243"
  id="VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07243"></a>
  VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07243  
  If the `pPictureResource` member of any element of `pReferenceSlots`
  is not `NULL`, then its `codedExtent` member **must** be between
  `minCodedExtent` and `maxCodedExtent`, inclusive, `videoSession` was
  created with

- <a href="#VUID-VkVideoBeginCodingInfoKHR-flags-07244"
  id="VUID-VkVideoBeginCodingInfoKHR-flags-07244"></a>
  VUID-VkVideoBeginCodingInfoKHR-flags-07244  
  If [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`flags` does
  not include `VK_VIDEO_CAPABILITY_SEPARATE_REFERENCE_IMAGES_BIT_KHR`,
  as returned by
  [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
  for the video profile `videoSession` was created with, then
  `pPictureResource->imageViewBinding` of all elements of
  `pReferenceSlots` with a non-`NULL` `pPictureResource` member **must**
  specify image views created from the same image

- <a href="#VUID-VkVideoBeginCodingInfoKHR-slotIndex-07245"
  id="VUID-VkVideoBeginCodingInfoKHR-slotIndex-07245"></a>
  VUID-VkVideoBeginCodingInfoKHR-slotIndex-07245  
  If `videoSession` was created with a decode operation and the
  `slotIndex` member of any element of `pReferenceSlots` is not
  negative, then the image view specified in
  `pPictureResource->imageViewBinding` for that element **must** have
  been created with `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`

- <a href="#VUID-VkVideoBeginCodingInfoKHR-slotIndex-07246"
  id="VUID-VkVideoBeginCodingInfoKHR-slotIndex-07246"></a>
  VUID-VkVideoBeginCodingInfoKHR-slotIndex-07246  
  If `videoSession` was created with an encode operation and the
  `slotIndex` member of any element of `pReferenceSlots` is not
  negative, then the image view specified in
  `pPictureResource->imageViewBinding` for that element **must** have
  been created with `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`

- <a href="#VUID-VkVideoBeginCodingInfoKHR-videoSession-07247"
  id="VUID-VkVideoBeginCodingInfoKHR-videoSession-07247"></a>
  VUID-VkVideoBeginCodingInfoKHR-videoSession-07247  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then
  `videoSessionParameters` **must** not be `VK_NULL_HANDLE`

- <a href="#VUID-VkVideoBeginCodingInfoKHR-videoSession-07248"
  id="VUID-VkVideoBeginCodingInfoKHR-videoSession-07248"></a>
  VUID-VkVideoBeginCodingInfoKHR-videoSession-07248  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then
  `videoSessionParameters` **must** not be `VK_NULL_HANDLE`

- <a href="#VUID-VkVideoBeginCodingInfoKHR-videoSession-09261"
  id="VUID-VkVideoBeginCodingInfoKHR-videoSession-09261"></a>
  VUID-VkVideoBeginCodingInfoKHR-videoSession-09261  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then
  `videoSessionParameters` **must** not be `VK_NULL_HANDLE`

- <a href="#VUID-VkVideoBeginCodingInfoKHR-videoSession-07249"
  id="VUID-VkVideoBeginCodingInfoKHR-videoSession-07249"></a>
  VUID-VkVideoBeginCodingInfoKHR-videoSession-07249  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then
  `videoSessionParameters` **must** not be `VK_NULL_HANDLE`

- <a href="#VUID-VkVideoBeginCodingInfoKHR-videoSession-07250"
  id="VUID-VkVideoBeginCodingInfoKHR-videoSession-07250"></a>
  VUID-VkVideoBeginCodingInfoKHR-videoSession-07250  
  If `videoSession` was created with the video codec operation
  `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then
  `videoSessionParameters` **must** not be `VK_NULL_HANDLE`

- <a href="#VUID-VkVideoBeginCodingInfoKHR-videoSessionParameters-04857"
  id="VUID-VkVideoBeginCodingInfoKHR-videoSessionParameters-04857"></a>
  VUID-VkVideoBeginCodingInfoKHR-videoSessionParameters-04857  
  If `videoSessionParameters` is not `VK_NULL_HANDLE`, it **must** have
  been created with `videoSession` specified in
  [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateInfoKHR.html)::`videoSession`

Valid Usage (Implicit)

- <a href="#VUID-VkVideoBeginCodingInfoKHR-sType-sType"
  id="VUID-VkVideoBeginCodingInfoKHR-sType-sType"></a>
  VUID-VkVideoBeginCodingInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_BEGIN_CODING_INFO_KHR`

- <a href="#VUID-VkVideoBeginCodingInfoKHR-pNext-pNext"
  id="VUID-VkVideoBeginCodingInfoKHR-pNext-pNext"></a>
  VUID-VkVideoBeginCodingInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkVideoEncodeH264GopRemainingFrameInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264GopRemainingFrameInfoKHR.html),
  [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlInfoKHR.html),
  [VkVideoEncodeH265GopRemainingFrameInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265GopRemainingFrameInfoKHR.html),
  [VkVideoEncodeH265RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlInfoKHR.html),
  or
  [VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlInfoKHR.html)

- <a href="#VUID-VkVideoBeginCodingInfoKHR-sType-unique"
  id="VUID-VkVideoBeginCodingInfoKHR-sType-unique"></a>
  VUID-VkVideoBeginCodingInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkVideoBeginCodingInfoKHR-flags-zerobitmask"
  id="VUID-VkVideoBeginCodingInfoKHR-flags-zerobitmask"></a>
  VUID-VkVideoBeginCodingInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

- <a href="#VUID-VkVideoBeginCodingInfoKHR-videoSession-parameter"
  id="VUID-VkVideoBeginCodingInfoKHR-videoSession-parameter"></a>
  VUID-VkVideoBeginCodingInfoKHR-videoSession-parameter  
  `videoSession` **must** be a valid
  [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionKHR.html) handle

- <a
  href="#VUID-VkVideoBeginCodingInfoKHR-videoSessionParameters-parameter"
  id="VUID-VkVideoBeginCodingInfoKHR-videoSessionParameters-parameter"></a>
  VUID-VkVideoBeginCodingInfoKHR-videoSessionParameters-parameter  
  If `videoSessionParameters` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `videoSessionParameters`
  **must** be a valid
  [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html) handle

- <a href="#VUID-VkVideoBeginCodingInfoKHR-pReferenceSlots-parameter"
  id="VUID-VkVideoBeginCodingInfoKHR-pReferenceSlots-parameter"></a>
  VUID-VkVideoBeginCodingInfoKHR-pReferenceSlots-parameter  
  If `referenceSlotCount` is not `0`, `pReferenceSlots` **must** be a
  valid pointer to an array of `referenceSlotCount` valid
  [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)
  structures

- <a href="#VUID-VkVideoBeginCodingInfoKHR-videoSessionParameters-parent"
  id="VUID-VkVideoBeginCodingInfoKHR-videoSessionParameters-parent"></a>
  VUID-VkVideoBeginCodingInfoKHR-videoSessionParameters-parent  
  If `videoSessionParameters` is a valid handle, it **must** have been
  created, allocated, or retrieved from `videoSession`

- <a href="#VUID-VkVideoBeginCodingInfoKHR-commonparent"
  id="VUID-VkVideoBeginCodingInfoKHR-commonparent"></a>
  VUID-VkVideoBeginCodingInfoKHR-commonparent  
  Both of `videoSession`, and `videoSessionParameters` that are valid
  handles of non-ignored parameters **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoBeginCodingFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingFlagsKHR.html),
[VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html),
[VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionKHR.html),
[VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersKHR.html),
[vkCmdBeginVideoCodingKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginVideoCodingKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoBeginCodingInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
