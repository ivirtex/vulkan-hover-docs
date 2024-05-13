# VkVideoEncodeH264DpbSlotInfoKHR(3) Manual Page

## Name

VkVideoEncodeH264DpbSlotInfoKHR - Structure specifies H.264 encode DPB
picture information



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkVideoEncodeH264DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264DpbSlotInfoKHR.html)
structure is defined as:

``` c
// Provided by VK_KHR_video_encode_h264
typedef struct VkVideoEncodeH264DpbSlotInfoKHR {
    VkStructureType                           sType;
    const void*                               pNext;
    const StdVideoEncodeH264ReferenceInfo*    pStdReferenceInfo;
} VkVideoEncodeH264DpbSlotInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pStdReferenceInfo` is a pointer to a
  `StdVideoEncodeH264ReferenceInfo` structure specifying <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-reference-info"
  target="_blank" rel="noopener">H.264 reference information</a>.

## <a href="#_description" class="anchor"></a>Description

This structure is specified in the `pNext` chain of
[VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`pSetupReferenceSlot`,
if not `NULL`, and the `pNext` chain of the elements of
[VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`pReferenceSlots` to
specify the codec-specific reference picture information for an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264"
target="_blank" rel="noopener">H.264 encode operation</a>.

Active Reference Picture Information  
When this structure is specified in the `pNext` chain of the elements of
[VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`pReferenceSlots`,
one element is added to the list of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-active-reference-picture-info"
target="_blank" rel="noopener">active reference pictures</a> used by the
video encode operation for each element of
[VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`pReferenceSlots` as
follows:

- The image subregion used is determined according to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-picture-data-access"
  target="_blank" rel="noopener">H.264 Encode Picture Data Access</a>
  section.

- The reference picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slot</a> index specified in the
  `slotIndex` member of the corresponding element of
  [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`pReferenceSlots`.

- The reference picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-reference-info"
  target="_blank" rel="noopener">H.264 reference information</a>
  provided in `pStdReferenceInfo`.

<!-- -->

Reconstructed Picture Information  
When this structure is specified in the `pNext` chain of
[VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`pSetupReferenceSlot`,
the information related to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-reconstructed-picture-info"
target="_blank" rel="noopener">reconstructed picture</a> is defined as
follows:

- The image subregion used is determined according to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-picture-data-access"
  target="_blank" rel="noopener">H.264 Encode Picture Data Access</a>
  section.

- If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-ref-pic-setup"
  target="_blank" rel="noopener">reference picture setup</a> is
  requested, then the reconstructed picture is used to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
  target="_blank" rel="noopener">activate</a> the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slot</a> with the index specified
  in
  [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`pSetupReferenceSlot->slotIndex`.

- The reconstructed picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-reference-info"
  target="_blank" rel="noopener">H.264 reference information</a>
  provided in `pStdReferenceInfo`.

<!-- -->

Std Reference Information  
The members of the `StdVideoEncodeH264ReferenceInfo` structure pointed
to by `pStdReferenceInfo` are interpreted as follows:

- `flags.reserved` is used only for padding purposes and is otherwise
  ignored;

- `flags.used_for_long_term_reference` is used to indicate whether the
  picture is marked as “used for long-term reference” as defined in
  section 8.2.5.1 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

- `primary_pic_type` as defined in section 7.4.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

- `long_term_pic_num` and `long_term_frame_idx` as defined in section
  7.4.3 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

- `temporal_id` as defined in section G.7.4.1.1 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

- all other members are interpreted as defined in section 8.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>.

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeH264DpbSlotInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeH264DpbSlotInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeH264DpbSlotInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H264_DPB_SLOT_INFO_KHR`

- <a
  href="#VUID-VkVideoEncodeH264DpbSlotInfoKHR-pStdReferenceInfo-parameter"
  id="VUID-VkVideoEncodeH264DpbSlotInfoKHR-pStdReferenceInfo-parameter"></a>
  VUID-VkVideoEncodeH264DpbSlotInfoKHR-pStdReferenceInfo-parameter  
  `pStdReferenceInfo` **must** be a valid pointer to a valid
  `StdVideoEncodeH264ReferenceInfo` value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h264.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH264DpbSlotInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
