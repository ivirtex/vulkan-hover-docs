# VkVideoBeginCodingInfoKHR(3) Manual Page

## Name

VkVideoBeginCodingInfoKHR - Structure specifying video coding scope begin information



## [](#_c_specification)C Specification

The [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoBeginCodingInfoKHR.html) structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `videoSession` is the video session object to be bound for the processing of the video commands.
- `videoSessionParameters` is `VK_NULL_HANDLE` or a handle of a [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersKHR.html) object to be used for the processing of the video commands. If `VK_NULL_HANDLE`, then no video session parameters object is bound for the duration of the video coding scope.
- `referenceSlotCount` is the number of elements in the `pReferenceSlots` array.
- `pReferenceSlots` is a pointer to an array of [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html) structures specifying the information used to determine the set of [bound reference picture resources](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#bound-reference-picture-resources) for the video coding scope and their initial association with [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) indices.

## [](#_description)Description

Limiting values are defined below that are referenced by the relevant valid usage statements of this structure.

- Let `VkOffset2D codedOffsetGranularity` be the minimum alignment requirement for the coded offset of video picture resources. Unless otherwise defined, the value of the `x` and `y` members of `codedOffsetGranularity` are `0`.
  
  - If `videoSession` was created with an [H.264 decode profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#decode-h264-profile) with a [VkVideoDecodeH264ProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264ProfileInfoKHR.html)::`pictureLayout` of `VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_SEPARATE_PLANES_BIT_KHR`, then `codedOffsetGranularity` is equal to [VkVideoDecodeH264CapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264CapabilitiesKHR.html)::`fieldOffsetGranularity`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for that video profile.

Valid Usage

- [](#VUID-VkVideoBeginCodingInfoKHR-videoSession-07237)VUID-VkVideoBeginCodingInfoKHR-videoSession-07237  
  `videoSession` **must** have memory bound to all of its memory bindings returned by [vkGetVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetVideoSessionMemoryRequirementsKHR.html) for `videoSession`
- [](#VUID-VkVideoBeginCodingInfoKHR-slotIndex-04856)VUID-VkVideoBeginCodingInfoKHR-slotIndex-04856  
  Each non-negative [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html)::`slotIndex` specified in the elements of `pReferenceSlots` **must** be less than the [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateInfoKHR.html)::`maxDpbSlots` specified when `videoSession` was created
- [](#VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07238)VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07238  
  Each video picture resource corresponding to any non-`NULL` `pPictureResource` member specified in the elements of `pReferenceSlots` **must** be [unique](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-picture-resource-uniqueness) within `pReferenceSlots`
- [](#VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07240)VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07240  
  If the `pPictureResource` member of any element of `pReferenceSlots` is not `NULL`, then the image view specified in `pPictureResource->imageViewBinding` for that element **must** be [compatible](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-profile-compatibility) with the video profile `videoSession` was created with
- [](#VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07241)VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07241  
  If the `pPictureResource` member of any element of `pReferenceSlots` is not `NULL`, then the format of the image view specified in `pPictureResource->imageViewBinding` for that element **must** match the [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateInfoKHR.html)::`referencePictureFormat` `videoSession` was created with
- [](#VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07242)VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07242  
  If the `pPictureResource` member of any element of `pReferenceSlots` is not `NULL`, then its `codedOffset` member **must** be an integer multiple of `codedOffsetGranularity`
- [](#VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07243)VUID-VkVideoBeginCodingInfoKHR-pPictureResource-07243  
  If the `pPictureResource` member of any element of `pReferenceSlots` is not `NULL`, then its `codedExtent` member **must** be between `minCodedExtent` and `maxCodedExtent`, inclusive, `videoSession` was created with
- [](#VUID-VkVideoBeginCodingInfoKHR-flags-07244)VUID-VkVideoBeginCodingInfoKHR-flags-07244  
  If [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)::`flags` does not include `VK_VIDEO_CAPABILITY_SEPARATE_REFERENCE_IMAGES_BIT_KHR`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the video profile `videoSession` was created with, then `pPictureResource->imageViewBinding` of all elements of `pReferenceSlots` with a non-`NULL` `pPictureResource` member **must** specify image views created from the same image
- [](#VUID-VkVideoBeginCodingInfoKHR-slotIndex-07245)VUID-VkVideoBeginCodingInfoKHR-slotIndex-07245  
  If `videoSession` was created with a decode operation and the `slotIndex` member of any element of `pReferenceSlots` is not negative, then the image view specified in `pPictureResource->imageViewBinding` for that element **must** have been created with `VK_IMAGE_USAGE_VIDEO_DECODE_DPB_BIT_KHR`
- [](#VUID-VkVideoBeginCodingInfoKHR-slotIndex-07246)VUID-VkVideoBeginCodingInfoKHR-slotIndex-07246  
  If `videoSession` was created with an encode operation and the `slotIndex` member of any element of `pReferenceSlots` is not negative, then the image view specified in `pPictureResource->imageViewBinding` for that element **must** have been created with `VK_IMAGE_USAGE_VIDEO_ENCODE_DPB_BIT_KHR`
- [](#VUID-VkVideoBeginCodingInfoKHR-videoSession-07247)VUID-VkVideoBeginCodingInfoKHR-videoSession-07247  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H264_BIT_KHR`, then `videoSessionParameters` **must** not be `VK_NULL_HANDLE` , unless the [`videoMaintenance2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-videoMaintenance2) feature is enabled
- [](#VUID-VkVideoBeginCodingInfoKHR-videoSession-07248)VUID-VkVideoBeginCodingInfoKHR-videoSession-07248  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_H265_BIT_KHR`, then `videoSessionParameters` **must** not be `VK_NULL_HANDLE` , unless the [`videoMaintenance2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-videoMaintenance2) feature is enabled
- [](#VUID-VkVideoBeginCodingInfoKHR-videoSession-09261)VUID-VkVideoBeginCodingInfoKHR-videoSession-09261  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_DECODE_AV1_BIT_KHR`, then `videoSessionParameters` **must** not be `VK_NULL_HANDLE` , unless the [`videoMaintenance2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-videoMaintenance2) feature is enabled
- [](#VUID-VkVideoBeginCodingInfoKHR-videoSession-07249)VUID-VkVideoBeginCodingInfoKHR-videoSession-07249  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H264_BIT_KHR`, then `videoSessionParameters` **must** not be `VK_NULL_HANDLE`
- [](#VUID-VkVideoBeginCodingInfoKHR-videoSession-07250)VUID-VkVideoBeginCodingInfoKHR-videoSession-07250  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_H265_BIT_KHR`, then `videoSessionParameters` **must** not be `VK_NULL_HANDLE`
- [](#VUID-VkVideoBeginCodingInfoKHR-videoSession-10283)VUID-VkVideoBeginCodingInfoKHR-videoSession-10283  
  If `videoSession` was created with the video codec operation `VK_VIDEO_CODEC_OPERATION_ENCODE_AV1_BIT_KHR`, then `videoSessionParameters` **must** not be `VK_NULL_HANDLE`
- [](#VUID-VkVideoBeginCodingInfoKHR-videoSessionParameters-04857)VUID-VkVideoBeginCodingInfoKHR-videoSessionParameters-04857  
  If `videoSessionParameters` is not `VK_NULL_HANDLE`, it **must** have been created with `videoSession` specified in [VkVideoSessionParametersCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersCreateInfoKHR.html)::`videoSession`

Valid Usage (Implicit)

- [](#VUID-VkVideoBeginCodingInfoKHR-sType-sType)VUID-VkVideoBeginCodingInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_BEGIN_CODING_INFO_KHR`
- [](#VUID-VkVideoBeginCodingInfoKHR-pNext-pNext)VUID-VkVideoBeginCodingInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkVideoEncodeAV1GopRemainingFrameInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1GopRemainingFrameInfoKHR.html), [VkVideoEncodeAV1RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1RateControlInfoKHR.html), [VkVideoEncodeH264GopRemainingFrameInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264GopRemainingFrameInfoKHR.html), [VkVideoEncodeH264RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264RateControlInfoKHR.html), [VkVideoEncodeH265GopRemainingFrameInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265GopRemainingFrameInfoKHR.html), [VkVideoEncodeH265RateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265RateControlInfoKHR.html), or [VkVideoEncodeRateControlInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRateControlInfoKHR.html)
- [](#VUID-VkVideoBeginCodingInfoKHR-sType-unique)VUID-VkVideoBeginCodingInfoKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkVideoBeginCodingInfoKHR-flags-zerobitmask)VUID-VkVideoBeginCodingInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkVideoBeginCodingInfoKHR-videoSession-parameter)VUID-VkVideoBeginCodingInfoKHR-videoSession-parameter  
  `videoSession` **must** be a valid [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionKHR.html) handle
- [](#VUID-VkVideoBeginCodingInfoKHR-videoSessionParameters-parameter)VUID-VkVideoBeginCodingInfoKHR-videoSessionParameters-parameter  
  If `videoSessionParameters` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `videoSessionParameters` **must** be a valid [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersKHR.html) handle
- [](#VUID-VkVideoBeginCodingInfoKHR-pReferenceSlots-parameter)VUID-VkVideoBeginCodingInfoKHR-pReferenceSlots-parameter  
  If `referenceSlotCount` is not `0`, `pReferenceSlots` **must** be a valid pointer to an array of `referenceSlotCount` valid [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html) structures
- [](#VUID-VkVideoBeginCodingInfoKHR-videoSessionParameters-parent)VUID-VkVideoBeginCodingInfoKHR-videoSessionParameters-parent  
  If `videoSessionParameters` is a valid handle, it **must** have been created, allocated, or retrieved from `videoSession`
- [](#VUID-VkVideoBeginCodingInfoKHR-commonparent)VUID-VkVideoBeginCodingInfoKHR-commonparent  
  Both of `videoSession`, and `videoSessionParameters` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoBeginCodingFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoBeginCodingFlagsKHR.html), [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html), [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionKHR.html), [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersKHR.html), [vkCmdBeginVideoCodingKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginVideoCodingKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoBeginCodingInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0