# VkVideoDecodeH265DpbSlotInfoKHR(3) Manual Page

## Name

VkVideoDecodeH265DpbSlotInfoKHR - Structure specifies H.265 DPB
information when decoding a frame



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoDecodeH265DpbSlotInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_decode_h265
typedef struct VkVideoDecodeH265DpbSlotInfoKHR {
    VkStructureType                           sType;
    const void*                               pNext;
    const StdVideoDecodeH265ReferenceInfo*    pStdReferenceInfo;
} VkVideoDecodeH265DpbSlotInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pStdReferenceInfo` is a pointer to a
  `StdVideoDecodeH265ReferenceInfo` structure specifying reference
  picture information described in section 8.3 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>.

## <a href="#_description" class="anchor"></a>Description

This structure is specified in the `pNext` chain of
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pSetupReferenceSlot`,
if not `NULL`, and the `pNext` chain of the elements of
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pReferenceSlots` to
specify the codec-specific reference picture information for an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265"
target="_blank" rel="noopener">H.265 decode operation</a>.

Active Reference Picture Information  
When this structure is specified in the `pNext` chain of the elements of
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pReferenceSlots`,
one element is added to the list of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-active-reference-picture-info"
target="_blank" rel="noopener">active reference pictures</a> used by the
video decode operation for each element of
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pReferenceSlots` as
follows:

- The image subregion used is determined according to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-picture-data-access"
  target="_blank" rel="noopener">H.265 Decode Picture Data Access</a>
  section.

- The reference picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slot</a> index specified in the
  `slotIndex` member of the corresponding element of
  [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pReferenceSlots`.

- The reference picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-reference-info"
  target="_blank" rel="noopener">H.265 reference information</a>
  provided in `pStdReferenceInfo`.

<!-- -->

Reconstructed Picture Information  
When this structure is specified in the `pNext` chain of
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pSetupReferenceSlot`,
the information related to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-reconstructed-picture-info"
target="_blank" rel="noopener">reconstructed picture</a> is defined as
follows:

- The image subregion used is determined according to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-picture-data-access"
  target="_blank" rel="noopener">H.265 Decode Picture Data Access</a>
  section.

- If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-ref-pic-setup"
  target="_blank" rel="noopener">reference picture setup</a> is
  requested, then the reconstructed picture is used to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
  target="_blank" rel="noopener">activate</a> the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slot</a> with the index specified
  in
  [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pSetupReferenceSlot->slotIndex`.

- The reconstructed picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h265-reference-info"
  target="_blank" rel="noopener">H.265 reference information</a>
  provided in `pStdReferenceInfo`.

<!-- -->

Std Reference Information  
The members of the `StdVideoDecodeH265ReferenceInfo` structure pointed
to by `pStdReferenceInfo` are interpreted as follows:

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

- all other members are interpreted as defined in section 8.3 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h265"
  target="_blank" rel="noopener">ITU-T H.265 Specification</a>.

Valid Usage (Implicit)

- <a href="#VUID-VkVideoDecodeH265DpbSlotInfoKHR-sType-sType"
  id="VUID-VkVideoDecodeH265DpbSlotInfoKHR-sType-sType"></a>
  VUID-VkVideoDecodeH265DpbSlotInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_DECODE_H265_DPB_SLOT_INFO_KHR`

- <a
  href="#VUID-VkVideoDecodeH265DpbSlotInfoKHR-pStdReferenceInfo-parameter"
  id="VUID-VkVideoDecodeH265DpbSlotInfoKHR-pStdReferenceInfo-parameter"></a>
  VUID-VkVideoDecodeH265DpbSlotInfoKHR-pStdReferenceInfo-parameter  
  `pStdReferenceInfo` **must** be a valid pointer to a valid
  `StdVideoDecodeH265ReferenceInfo` value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_h265.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeH265DpbSlotInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
