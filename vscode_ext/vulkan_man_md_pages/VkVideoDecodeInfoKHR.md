# VkVideoDecodeInfoKHR(3) Manual Page

## Name

VkVideoDecodeInfoKHR - Structure specifying video decode parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoDecodeInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_decode_queue
typedef struct VkVideoDecodeInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkVideoDecodeFlagsKHR                 flags;
    VkBuffer                              srcBuffer;
    VkDeviceSize                          srcBufferOffset;
    VkDeviceSize                          srcBufferRange;
    VkVideoPictureResourceInfoKHR         dstPictureResource;
    const VkVideoReferenceSlotInfoKHR*    pSetupReferenceSlot;
    uint32_t                              referenceSlotCount;
    const VkVideoReferenceSlotInfoKHR*    pReferenceSlots;
} VkVideoDecodeInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `srcBuffer` is the source video bitstream buffer to read the encoded
  bitstream from.

- `srcBufferOffset` is the starting offset in bytes from the start of
  `srcBuffer` to read the encoded bitstream from.

- `srcBufferRange` is the size in bytes of the encoded bitstream to
  decode from `srcBuffer`, starting from `srcBufferOffset`.

- `dstPictureResource` is the video picture resource to use as the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-output-picture"
  target="_blank" rel="noopener">decode output picture</a>.

- `pSetupReferenceSlot` is `NULL` or a pointer to a
  [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)
  structure specifying the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-reconstructed-picture-info"
  target="_blank" rel="noopener">reconstructed picture information</a>.

- `referenceSlotCount` is the number of elements in the
  `pReferenceSlots` array.

- `pReferenceSlots` is `NULL` or a pointer to an array of
  [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)
  structures describing the DPB slots and corresponding <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#reference-picture"
  target="_blank" rel="noopener">reference picture</a> resources to use
  in this video decode operation (the set of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#active-reference-pictures"
  target="_blank" rel="noopener">active reference pictures</a>).

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkVideoDecodeInfoKHR-srcBuffer-07165"
  id="VUID-VkVideoDecodeInfoKHR-srcBuffer-07165"></a>
  VUID-VkVideoDecodeInfoKHR-srcBuffer-07165  
  `srcBuffer` **must** have been created with
  `VK_BUFFER_USAGE_VIDEO_DECODE_SRC_BIT_KHR` set

- <a href="#VUID-VkVideoDecodeInfoKHR-srcBufferOffset-07166"
  id="VUID-VkVideoDecodeInfoKHR-srcBufferOffset-07166"></a>
  VUID-VkVideoDecodeInfoKHR-srcBufferOffset-07166  
  `srcBufferOffset` **must** be less than the size of `srcBuffer`

- <a href="#VUID-VkVideoDecodeInfoKHR-srcBufferRange-07167"
  id="VUID-VkVideoDecodeInfoKHR-srcBufferRange-07167"></a>
  VUID-VkVideoDecodeInfoKHR-srcBufferRange-07167  
  `srcBufferRange` **must** be less than or equal to the size of
  `srcBuffer` minus `srcBufferOffset`

- <a href="#VUID-VkVideoDecodeInfoKHR-pSetupReferenceSlot-07168"
  id="VUID-VkVideoDecodeInfoKHR-pSetupReferenceSlot-07168"></a>
  VUID-VkVideoDecodeInfoKHR-pSetupReferenceSlot-07168  
  If `pSetupReferenceSlot` is not `NULL`, then its `slotIndex` member
  **must** not be negative

- <a href="#VUID-VkVideoDecodeInfoKHR-pSetupReferenceSlot-07169"
  id="VUID-VkVideoDecodeInfoKHR-pSetupReferenceSlot-07169"></a>
  VUID-VkVideoDecodeInfoKHR-pSetupReferenceSlot-07169  
  If `pSetupReferenceSlot` is not `NULL`, then its `pPictureResource`
  **must** not be `NULL`

- <a href="#VUID-VkVideoDecodeInfoKHR-slotIndex-07171"
  id="VUID-VkVideoDecodeInfoKHR-slotIndex-07171"></a>
  VUID-VkVideoDecodeInfoKHR-slotIndex-07171  
  The `slotIndex` member of each element of `pReferenceSlots` **must**
  not be negative

- <a href="#VUID-VkVideoDecodeInfoKHR-pPictureResource-07172"
  id="VUID-VkVideoDecodeInfoKHR-pPictureResource-07172"></a>
  VUID-VkVideoDecodeInfoKHR-pPictureResource-07172  
  The `pPictureResource` member of each element of `pReferenceSlots`
  **must** not be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-VkVideoDecodeInfoKHR-sType-sType"
  id="VUID-VkVideoDecodeInfoKHR-sType-sType"></a>
  VUID-VkVideoDecodeInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_DECODE_INFO_KHR`

- <a href="#VUID-VkVideoDecodeInfoKHR-pNext-pNext"
  id="VUID-VkVideoDecodeInfoKHR-pNext-pNext"></a>
  VUID-VkVideoDecodeInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkVideoDecodeAV1PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeAV1PictureInfoKHR.html),
  [VkVideoDecodeH264PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264PictureInfoKHR.html),
  [VkVideoDecodeH265PictureInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH265PictureInfoKHR.html),
  or [VkVideoInlineQueryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoInlineQueryInfoKHR.html)

- <a href="#VUID-VkVideoDecodeInfoKHR-sType-unique"
  id="VUID-VkVideoDecodeInfoKHR-sType-unique"></a>
  VUID-VkVideoDecodeInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkVideoDecodeInfoKHR-flags-zerobitmask"
  id="VUID-VkVideoDecodeInfoKHR-flags-zerobitmask"></a>
  VUID-VkVideoDecodeInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

- <a href="#VUID-VkVideoDecodeInfoKHR-srcBuffer-parameter"
  id="VUID-VkVideoDecodeInfoKHR-srcBuffer-parameter"></a>
  VUID-VkVideoDecodeInfoKHR-srcBuffer-parameter  
  `srcBuffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkVideoDecodeInfoKHR-dstPictureResource-parameter"
  id="VUID-VkVideoDecodeInfoKHR-dstPictureResource-parameter"></a>
  VUID-VkVideoDecodeInfoKHR-dstPictureResource-parameter  
  `dstPictureResource` **must** be a valid
  [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoPictureResourceInfoKHR.html)
  structure

- <a href="#VUID-VkVideoDecodeInfoKHR-pSetupReferenceSlot-parameter"
  id="VUID-VkVideoDecodeInfoKHR-pSetupReferenceSlot-parameter"></a>
  VUID-VkVideoDecodeInfoKHR-pSetupReferenceSlot-parameter  
  If `pSetupReferenceSlot` is not `NULL`, `pSetupReferenceSlot` **must**
  be a valid pointer to a valid
  [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)
  structure

- <a href="#VUID-VkVideoDecodeInfoKHR-pReferenceSlots-parameter"
  id="VUID-VkVideoDecodeInfoKHR-pReferenceSlots-parameter"></a>
  VUID-VkVideoDecodeInfoKHR-pReferenceSlots-parameter  
  If `referenceSlotCount` is not `0`, `pReferenceSlots` **must** be a
  valid pointer to an array of `referenceSlotCount` valid
  [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)
  structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_queue.html),
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoDecodeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeFlagsKHR.html),
[VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoPictureResourceInfoKHR.html),
[VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html),
[vkCmdDecodeVideoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDecodeVideoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
