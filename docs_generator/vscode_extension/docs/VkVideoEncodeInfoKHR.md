# VkVideoEncodeInfoKHR(3) Manual Page

## Name

VkVideoEncodeInfoKHR - Structure specifying video encode parameters



## [](#_c_specification)C Specification

The `VkVideoEncodeInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_encode_queue
typedef struct VkVideoEncodeInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkVideoEncodeFlagsKHR                 flags;
    VkBuffer                              dstBuffer;
    VkDeviceSize                          dstBufferOffset;
    VkDeviceSize                          dstBufferRange;
    VkVideoPictureResourceInfoKHR         srcPictureResource;
    const VkVideoReferenceSlotInfoKHR*    pSetupReferenceSlot;
    uint32_t                              referenceSlotCount;
    const VkVideoReferenceSlotInfoKHR*    pReferenceSlots;
    uint32_t                              precedingExternallyEncodedBytes;
} VkVideoEncodeInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkVideoEncodeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeFlagBitsKHR.html) indicating video encode command flags.
- `dstBuffer` is the destination video bitstream buffer to write the encoded bitstream to.
- `dstBufferOffset` is the starting offset in bytes from the start of `dstBuffer` to write the encoded bitstream to.
- `dstBufferRange` is the maximum bitstream size in bytes that **can** be written to `dstBuffer`, starting from `dstBufferOffset`.
- `srcPictureResource` is the video picture resource to use as the [encode input picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-input-picture).
- `pSetupReferenceSlot` is `NULL` or a pointer to a [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html) structure specifying the [reconstructed picture information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-reconstructed-picture-info).
- `referenceSlotCount` is the number of elements in the `pReferenceSlots` array.
- `pReferenceSlots` is `NULL` or a pointer to an array of [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html) structures describing the DPB slots and corresponding [reference picture](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#reference-picture) resources to use in this video encode operation (the set of [active reference pictures](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#active-reference-pictures)).
- `precedingExternallyEncodedBytes` is the number of bytes externally encoded by the application to the video bitstream and is used to update the internal state of the implementation’s [rate control](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#encode-rate-control) algorithm to account for the bitrate budget consumed by these externally encoded bytes.

## [](#_description)Description

Valid Usage

- [](#VUID-VkVideoEncodeInfoKHR-dstBuffer-08236)VUID-VkVideoEncodeInfoKHR-dstBuffer-08236  
  `dstBuffer` **must** have been created with `VK_BUFFER_USAGE_VIDEO_ENCODE_DST_BIT_KHR` set
- [](#VUID-VkVideoEncodeInfoKHR-dstBufferOffset-08237)VUID-VkVideoEncodeInfoKHR-dstBufferOffset-08237  
  `dstBufferOffset` **must** be less than the size of `dstBuffer`
- [](#VUID-VkVideoEncodeInfoKHR-dstBufferRange-08238)VUID-VkVideoEncodeInfoKHR-dstBufferRange-08238  
  `dstBufferRange` **must** be less than or equal to the size of `dstBuffer` minus `dstBufferOffset`
- [](#VUID-VkVideoEncodeInfoKHR-pSetupReferenceSlot-08239)VUID-VkVideoEncodeInfoKHR-pSetupReferenceSlot-08239  
  If `pSetupReferenceSlot` is not `NULL`, then its `slotIndex` member **must** not be negative
- [](#VUID-VkVideoEncodeInfoKHR-pSetupReferenceSlot-08240)VUID-VkVideoEncodeInfoKHR-pSetupReferenceSlot-08240  
  If `pSetupReferenceSlot` is not `NULL`, then its `pPictureResource` **must** not be `NULL`
- [](#VUID-VkVideoEncodeInfoKHR-slotIndex-08241)VUID-VkVideoEncodeInfoKHR-slotIndex-08241  
  The `slotIndex` member of each element of `pReferenceSlots` **must** not be negative
- [](#VUID-VkVideoEncodeInfoKHR-pPictureResource-08242)VUID-VkVideoEncodeInfoKHR-pPictureResource-08242  
  The `pPictureResource` member of each element of `pReferenceSlots` **must** not be `NULL`

Valid Usage (Implicit)

- [](#VUID-VkVideoEncodeInfoKHR-sType-sType)VUID-VkVideoEncodeInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_INFO_KHR`
- [](#VUID-VkVideoEncodeInfoKHR-pNext-pNext)VUID-VkVideoEncodeInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkVideoEncodeAV1PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeAV1PictureInfoKHR.html), [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH264PictureInfoKHR.html), [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeH265PictureInfoKHR.html), [VkVideoEncodeQuantizationMapInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQuantizationMapInfoKHR.html), or [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoInlineQueryInfoKHR.html)
- [](#VUID-VkVideoEncodeInfoKHR-sType-unique)VUID-VkVideoEncodeInfoKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkVideoEncodeInfoKHR-flags-parameter)VUID-VkVideoEncodeInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of [VkVideoEncodeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeFlagBitsKHR.html) values
- [](#VUID-VkVideoEncodeInfoKHR-dstBuffer-parameter)VUID-VkVideoEncodeInfoKHR-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-VkVideoEncodeInfoKHR-srcPictureResource-parameter)VUID-VkVideoEncodeInfoKHR-srcPictureResource-parameter  
  `srcPictureResource` **must** be a valid [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoPictureResourceInfoKHR.html) structure
- [](#VUID-VkVideoEncodeInfoKHR-pSetupReferenceSlot-parameter)VUID-VkVideoEncodeInfoKHR-pSetupReferenceSlot-parameter  
  If `pSetupReferenceSlot` is not `NULL`, `pSetupReferenceSlot` **must** be a valid pointer to a valid [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html) structure
- [](#VUID-VkVideoEncodeInfoKHR-pReferenceSlots-parameter)VUID-VkVideoEncodeInfoKHR-pReferenceSlots-parameter  
  If `referenceSlotCount` is not `0`, `pReferenceSlots` **must** be a valid pointer to an array of `referenceSlotCount` valid [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html) structures

## [](#_see_also)See Also

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoEncodeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeFlagsKHR.html), [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoPictureResourceInfoKHR.html), [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html), [vkCmdEncodeVideoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEncodeVideoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoEncodeInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0