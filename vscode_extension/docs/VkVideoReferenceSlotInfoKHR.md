# VkVideoReferenceSlotInfoKHR(3) Manual Page

## Name

VkVideoReferenceSlotInfoKHR - Structure specifying information about a reference picture slot



## [](#_c_specification)C Specification

The `VkVideoReferenceSlotInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_queue
typedef struct VkVideoReferenceSlotInfoKHR {
    VkStructureType                         sType;
    const void*                             pNext;
    int32_t                                 slotIndex;
    const VkVideoPictureResourceInfoKHR*    pPictureResource;
} VkVideoReferenceSlotInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `slotIndex` is the index of the [DPB slot](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) or a negative integer value.
- `pPictureResource` is `NULL` or a pointer to a [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoPictureResourceInfoKHR.html) structure describing the [video picture resource](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-picture-resources) associated with the DPB slot index specified by `slotIndex`.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoReferenceSlotInfoKHR-sType-sType)VUID-VkVideoReferenceSlotInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_REFERENCE_SLOT_INFO_KHR`
- [](#VUID-VkVideoReferenceSlotInfoKHR-pNext-pNext)VUID-VkVideoReferenceSlotInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkVideoDecodeAV1DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeAV1DpbSlotInfoKHR.html), [VkVideoDecodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH264DpbSlotInfoKHR.html), [VkVideoDecodeH265DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeH265DpbSlotInfoKHR.html), [VkVideoEncodeAV1DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1DpbSlotInfoKHR.html), [VkVideoEncodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264DpbSlotInfoKHR.html), [VkVideoEncodeH265DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265DpbSlotInfoKHR.html), or [VkVideoReferenceIntraRefreshInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceIntraRefreshInfoKHR.html)
- [](#VUID-VkVideoReferenceSlotInfoKHR-sType-unique)VUID-VkVideoReferenceSlotInfoKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkVideoReferenceSlotInfoKHR-pPictureResource-parameter)VUID-VkVideoReferenceSlotInfoKHR-pPictureResource-parameter  
  If `pPictureResource` is not `NULL`, `pPictureResource` **must** be a valid pointer to a valid [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoPictureResourceInfoKHR.html) structure

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoBeginCodingInfoKHR.html), [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html), [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html), [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoPictureResourceInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoReferenceSlotInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0