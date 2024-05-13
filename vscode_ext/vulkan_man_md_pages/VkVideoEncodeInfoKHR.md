# VkVideoEncodeInfoKHR(3) Manual Page

## Name

VkVideoEncodeInfoKHR - Structure specifying video encode parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeInfoKHR` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is a pointer to a structure extending this structure.

- `flags` is reserved for future use.

- `dstBuffer` is the destination video bitstream buffer to write the
  encoded bitstream to.

- `dstBufferOffset` is the starting offset in bytes from the start of
  `dstBuffer` to write the encoded bitstream to.

- `dstBufferRange` is the maximum bitstream size in bytes that **can**
  be written to `dstBuffer`, starting from `dstBufferOffset`.

- `srcPictureResource` is the video picture resource to use as the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-input-picture"
  target="_blank" rel="noopener">encode input picture</a>.

- `pSetupReferenceSlot` is `NULL` or a pointer to a
  [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)
  structure specifying the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-reconstructed-picture-info"
  target="_blank" rel="noopener">reconstructed picture information</a>.

- `referenceSlotCount` is the number of elements in the
  `pReferenceSlots` array.

- `pReferenceSlots` is `NULL` or a pointer to an array of
  [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)
  structures describing the DPB slots and corresponding <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reference-picture"
  target="_blank" rel="noopener">reference picture</a> resources to use
  in this video encode operation (the set of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#active-reference-pictures"
  target="_blank" rel="noopener">active reference pictures</a>).

- `precedingExternallyEncodedBytes` is the number of bytes externally
  encoded by the application to the video bitstream and is used to
  update the internal state of the implementation’s <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-rate-control"
  target="_blank" rel="noopener">rate control</a> algorithm to account
  for the bitrate budget consumed by these externally encoded bytes.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkVideoEncodeInfoKHR-dstBuffer-08236"
  id="VUID-VkVideoEncodeInfoKHR-dstBuffer-08236"></a>
  VUID-VkVideoEncodeInfoKHR-dstBuffer-08236  
  `dstBuffer` **must** have been created with
  `VK_BUFFER_USAGE_VIDEO_ENCODE_DST_BIT_KHR` set

- <a href="#VUID-VkVideoEncodeInfoKHR-dstBufferOffset-08237"
  id="VUID-VkVideoEncodeInfoKHR-dstBufferOffset-08237"></a>
  VUID-VkVideoEncodeInfoKHR-dstBufferOffset-08237  
  `dstBufferOffset` **must** be less than the size of `dstBuffer`

- <a href="#VUID-VkVideoEncodeInfoKHR-dstBufferRange-08238"
  id="VUID-VkVideoEncodeInfoKHR-dstBufferRange-08238"></a>
  VUID-VkVideoEncodeInfoKHR-dstBufferRange-08238  
  `dstBufferRange` **must** be less than or equal to the size of
  `dstBuffer` minus `dstBufferOffset`

- <a href="#VUID-VkVideoEncodeInfoKHR-pSetupReferenceSlot-08239"
  id="VUID-VkVideoEncodeInfoKHR-pSetupReferenceSlot-08239"></a>
  VUID-VkVideoEncodeInfoKHR-pSetupReferenceSlot-08239  
  If `pSetupReferenceSlot` is not `NULL`, then its `slotIndex` member
  **must** not be negative

- <a href="#VUID-VkVideoEncodeInfoKHR-pSetupReferenceSlot-08240"
  id="VUID-VkVideoEncodeInfoKHR-pSetupReferenceSlot-08240"></a>
  VUID-VkVideoEncodeInfoKHR-pSetupReferenceSlot-08240  
  If `pSetupReferenceSlot` is not `NULL`, then its `pPictureResource`
  **must** not be `NULL`

- <a href="#VUID-VkVideoEncodeInfoKHR-slotIndex-08241"
  id="VUID-VkVideoEncodeInfoKHR-slotIndex-08241"></a>
  VUID-VkVideoEncodeInfoKHR-slotIndex-08241  
  The `slotIndex` member of each element of `pReferenceSlots` **must**
  not be negative

- <a href="#VUID-VkVideoEncodeInfoKHR-pPictureResource-08242"
  id="VUID-VkVideoEncodeInfoKHR-pPictureResource-08242"></a>
  VUID-VkVideoEncodeInfoKHR-pPictureResource-08242  
  The `pPictureResource` member of each element of `pReferenceSlots`
  **must** not be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_ENCODE_INFO_KHR`

- <a href="#VUID-VkVideoEncodeInfoKHR-pNext-pNext"
  id="VUID-VkVideoEncodeInfoKHR-pNext-pNext"></a>
  VUID-VkVideoEncodeInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkVideoEncodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264PictureInfoKHR.html),
  [VkVideoEncodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265PictureInfoKHR.html),
  or [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html)

- <a href="#VUID-VkVideoEncodeInfoKHR-sType-unique"
  id="VUID-VkVideoEncodeInfoKHR-sType-unique"></a>
  VUID-VkVideoEncodeInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkVideoEncodeInfoKHR-flags-zerobitmask"
  id="VUID-VkVideoEncodeInfoKHR-flags-zerobitmask"></a>
  VUID-VkVideoEncodeInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

- <a href="#VUID-VkVideoEncodeInfoKHR-dstBuffer-parameter"
  id="VUID-VkVideoEncodeInfoKHR-dstBuffer-parameter"></a>
  VUID-VkVideoEncodeInfoKHR-dstBuffer-parameter  
  `dstBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkVideoEncodeInfoKHR-srcPictureResource-parameter"
  id="VUID-VkVideoEncodeInfoKHR-srcPictureResource-parameter"></a>
  VUID-VkVideoEncodeInfoKHR-srcPictureResource-parameter  
  `srcPictureResource` **must** be a valid
  [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoPictureResourceInfoKHR.html)
  structure

- <a href="#VUID-VkVideoEncodeInfoKHR-pSetupReferenceSlot-parameter"
  id="VUID-VkVideoEncodeInfoKHR-pSetupReferenceSlot-parameter"></a>
  VUID-VkVideoEncodeInfoKHR-pSetupReferenceSlot-parameter  
  If `pSetupReferenceSlot` is not `NULL`, `pSetupReferenceSlot` **must**
  be a valid pointer to a valid
  [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)
  structure

- <a href="#VUID-VkVideoEncodeInfoKHR-pReferenceSlots-parameter"
  id="VUID-VkVideoEncodeInfoKHR-pReferenceSlots-parameter"></a>
  VUID-VkVideoEncodeInfoKHR-pReferenceSlots-parameter  
  If `referenceSlotCount` is not `0`, `pReferenceSlots` **must** be a
  valid pointer to an array of `referenceSlotCount` valid
  [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)
  structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoEncodeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeFlagsKHR.html),
[VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoPictureResourceInfoKHR.html),
[VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html),
[vkCmdEncodeVideoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdEncodeVideoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
