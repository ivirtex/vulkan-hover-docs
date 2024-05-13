# VkVideoReferenceSlotInfoKHR(3) Manual Page

## Name

VkVideoReferenceSlotInfoKHR - Structure specifying information about a
reference picture slot



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoReferenceSlotInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_queue
typedef struct VkVideoReferenceSlotInfoKHR {
    VkStructureType                         sType;
    const void*                             pNext;
    int32_t                                 slotIndex;
    const VkVideoPictureResourceInfoKHR*    pPictureResource;
} VkVideoReferenceSlotInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `slotIndex` is the index of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slot</a> or a negative integer
  value.

- `pPictureResource` is `NULL` or a pointer to a
  [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoPictureResourceInfoKHR.html)
  structure describing the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resources"
  target="_blank" rel="noopener">video picture resource</a> associated
  with the DPB slot index specified by `slotIndex`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoReferenceSlotInfoKHR-sType-sType"
  id="VUID-VkVideoReferenceSlotInfoKHR-sType-sType"></a>
  VUID-VkVideoReferenceSlotInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_REFERENCE_SLOT_INFO_KHR`

- <a href="#VUID-VkVideoReferenceSlotInfoKHR-pNext-pNext"
  id="VUID-VkVideoReferenceSlotInfoKHR-pNext-pNext"></a>
  VUID-VkVideoReferenceSlotInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkVideoDecodeAV1DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1DpbSlotInfoKHR.html),
  [VkVideoDecodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264DpbSlotInfoKHR.html),
  [VkVideoDecodeH265DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265DpbSlotInfoKHR.html),
  [VkVideoEncodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264DpbSlotInfoKHR.html),
  or
  [VkVideoEncodeH265DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265DpbSlotInfoKHR.html)

- <a href="#VUID-VkVideoReferenceSlotInfoKHR-sType-unique"
  id="VUID-VkVideoReferenceSlotInfoKHR-sType-unique"></a>
  VUID-VkVideoReferenceSlotInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkVideoReferenceSlotInfoKHR-pPictureResource-parameter"
  id="VUID-VkVideoReferenceSlotInfoKHR-pPictureResource-parameter"></a>
  VUID-VkVideoReferenceSlotInfoKHR-pPictureResource-parameter  
  If `pPictureResource` is not `NULL`, `pPictureResource` **must** be a
  valid pointer to a valid
  [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoPictureResourceInfoKHR.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoBeginCodingInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingInfoKHR.html),
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html),
[VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html),
[VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoPictureResourceInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoReferenceSlotInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
