# VkVideoEncodeH265DpbSlotInfoKHR(3) Manual Page

## Name

VkVideoEncodeH265DpbSlotInfoKHR - Structure specifies H.265 encode DPB
picture information



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkVideoEncodeH265DpbSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265DpbSlotInfoKHR.html)
structure is defined as:

``` c
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265DpbSlotInfoKHR {
    VkStructureType                           sType;
    const void*                               pNext;
    const StdVideoEncodeH265ReferenceInfo*    pStdReferenceInfo;
} VkVideoEncodeH265DpbSlotInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pStdReferenceInfo` is a pointer to a
  `StdVideoEncodeH265ReferenceInfo` structure specifying <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-reference-info"
  target="_blank" rel="noopener">H.265 reference information</a>.

## <a href="#_description" class="anchor"></a>Description

This structure is specified in the `pNext` chain of
[VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`pSetupReferenceSlot`,
if not `NULL`, and the `pNext` chain of the elements of
[VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`pReferenceSlots` to
specify the codec-specific reference picture information for an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265"
target="_blank" rel="noopener">H.265 encode operation</a>.

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
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-picture-data-access"
  target="_blank" rel="noopener">H.265 Encode Picture Data Access</a>
  section.

- The reference picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slot</a> index specified in the
  `slotIndex` member of the corresponding element of
  [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`pReferenceSlots`.

- The reference picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-reference-info"
  target="_blank" rel="noopener">H.265 reference information</a>
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
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-picture-data-access"
  target="_blank" rel="noopener">H.265 Encode Picture Data Access</a>
  section.

- If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-ref-pic-setup"
  target="_blank" rel="noopener">reference picture setup</a> is
  requested, then the reconstructed picture is used to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
  target="_blank" rel="noopener">activate</a> the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slot</a> with the index specified
  in
  [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html)::`pSetupReferenceSlot->slotIndex`.

- The reconstructed picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-reference-info"
  target="_blank" rel="noopener">H.265 reference information</a>
  provided in `pStdReferenceInfo`.

<!-- -->

Std Reference Information  
The members of the `StdVideoEncodeH265ReferenceInfo` structure pointed
to by `pStdReferenceInfo` are interpreted as follows:

- `flags.reserved` is used only for padding purposes and is otherwise
  ignored;

- `flags.used_for_long_term_reference` is used to indicate whether the
  picture is marked as “used for long-term reference” as defined in
  section 8.3.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- `flags.unused_for_reference` is used to indicate whether the picture
  is marked as “unused for reference” as defined in section 8.3.2 of the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- `pic_type` as defined in section 7.4.3.5 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- `PicOrderCntVal` as defined in section 8.3.1 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>;

- `TemporalId` as defined in section 7.4.2.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>.

Valid Usage (Implicit)

- <a href="#VUID-VkVideoEncodeH265DpbSlotInfoKHR-sType-sType"
  id="VUID-VkVideoEncodeH265DpbSlotInfoKHR-sType-sType"></a>
  VUID-VkVideoEncodeH265DpbSlotInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_ENCODE_H265_DPB_SLOT_INFO_KHR`

- <a
  href="#VUID-VkVideoEncodeH265DpbSlotInfoKHR-pStdReferenceInfo-parameter"
  id="VUID-VkVideoEncodeH265DpbSlotInfoKHR-pStdReferenceInfo-parameter"></a>
  VUID-VkVideoEncodeH265DpbSlotInfoKHR-pStdReferenceInfo-parameter  
  `pStdReferenceInfo` **must** be a valid pointer to a valid
  `StdVideoEncodeH265ReferenceInfo` value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h265.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH265DpbSlotInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
