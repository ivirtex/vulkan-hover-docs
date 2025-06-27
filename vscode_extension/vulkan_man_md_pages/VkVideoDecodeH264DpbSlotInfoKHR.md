# VkVideoDecodeH264DpbSlotInfoKHR(3) Manual Page

## Name

VkVideoDecodeH264DpbSlotInfoKHR - Structure specifies H.264 decode DPB
picture information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoDecodeH264DpbSlotInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_decode_h264
typedef struct VkVideoDecodeH264DpbSlotInfoKHR {
    VkStructureType                           sType;
    const void*                               pNext;
    const StdVideoDecodeH264ReferenceInfo*    pStdReferenceInfo;
} VkVideoDecodeH264DpbSlotInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pStdReferenceInfo` is a pointer to a
  `StdVideoDecodeH264ReferenceInfo` structure specifying <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-reference-info"
  target="_blank" rel="noopener">H.264 reference information</a>.

## <a href="#_description" class="anchor"></a>Description

This structure is specified in the `pNext` chain of
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pSetupReferenceSlot`,
if not `NULL`, and the `pNext` chain of the elements of
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pReferenceSlots` to
specify the codec-specific reference picture information for an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264"
target="_blank" rel="noopener">H.264 decode operation</a>.

Active Reference Picture Information  
When this structure is specified in the `pNext` chain of the elements of
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pReferenceSlots`,
one or two elements are added to the list of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-active-reference-picture-info"
target="_blank" rel="noopener">active reference pictures</a> used by the
video decode operation for each element of
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pReferenceSlots` as
follows:

- If neither `pStdReferenceInfo->flags.top_field_flag` nor
  `pStdReferenceInfo->flags.bottom_field_flag` is set, then the picture
  is added as a frame reference to the list of active reference
  pictures.

- If `pStdReferenceInfo->flags.top_field_flag` is set, then the picture
  is added as a top field reference to the list of active reference
  pictures.

- If `pStdReferenceInfo->flags.bottom_field_flag` is set, then the
  picture is added as a bottom field reference to the list of active
  reference pictures.

- For each added reference picture, the corresponding image subregion
  used is determined according to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-picture-data-access"
  target="_blank" rel="noopener">H.264 Decode Picture Data Access</a>
  section.

- Each added reference picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slot</a> index specified in the
  `slotIndex` member of the corresponding element of
  [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pReferenceSlots`.

- Each added reference picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-reference-info"
  target="_blank" rel="noopener">H.264 reference information</a>
  provided in `pStdReferenceInfo`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>When both the top and bottom field of an interlaced frame currently
associated with a DPB slot is intended to be used as an active reference
picture and both fields are stored in the same image subregion (which is
the case when using
<code>VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_INTERLEAVED_LINES_BIT_KHR</code>
which stores the two fields at even and odd scanlines of the same image
subregion), both references have to be provided through a single <a
href="VkVideoReferenceSlotInfoKHR.html">VkVideoReferenceSlotInfoKHR</a>
structure that has both <code>flags.top_field_flag</code> and
<code>flags.bottom_field_flag</code> set in the
<code>StdVideoDecodeH264ReferenceInfo</code> structure pointed to by the
<code>pStdReferenceInfo</code> member of the <a
href="VkVideoDecodeH264DpbSlotInfoKHR.html">VkVideoDecodeH264DpbSlotInfoKHR</a>
structure included in the corresponding <a
href="VkVideoReferenceSlotInfoKHR.html">VkVideoReferenceSlotInfoKHR</a>
structure’s <code>pNext</code> chain. However, this approach can only be
used when both fields are stored in the same image subregion. If that is
not the case (e.g. when using
<code>VK_VIDEO_DECODE_H264_PICTURE_LAYOUT_INTERLACED_SEPARATE_PLANES_BIT_KHR</code>
which requires separate <code>codedOffset</code> values for the two
fields and also allows storing the two fields of a frame in separate
image layers or entirely separate images), then a separate <a
href="VkVideoReferenceSlotInfoKHR.html">VkVideoReferenceSlotInfoKHR</a>
structure needs to be provided for referencing the two fields, each only
setting one of <code>flags.top_field_flag</code> or
<code>flags.bottom_field_flag</code>, and providing the appropriate
video picture resource information in <a
href="VkVideoReferenceSlotInfoKHR.html">VkVideoReferenceSlotInfoKHR</a>::<code>pPictureResource</code>.</p></td>
</tr>
</tbody>
</table>

Reconstructed Picture Information  
When this structure is specified in the `pNext` chain of
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pSetupReferenceSlot`,
the information related to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-reconstructed-picture-info"
target="_blank" rel="noopener">reconstructed picture</a> is defined as
follows:

- If neither `pStdReferenceInfo->flags.top_field_flag` nor
  `pStdReferenceInfo->flags.bottom_field_flag` is set, then the picture
  represents a frame.

- If `pStdReferenceInfo->flags.top_field_flag` is set, then the picture
  represents a field, specifically, the top field of the frame.

- If `pStdReferenceInfo->flags.bottom_field_flag` is set, then the
  picture represents a field, specifically, the bottom field of the
  frame.

- The image subregion used is determined according to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-picture-data-access"
  target="_blank" rel="noopener">H.264 Decode Picture Data Access</a>
  section.

- If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-ref-pic-setup"
  target="_blank" rel="noopener">reference picture setup</a> is
  requested, then the reconstructed picture is used to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot-states"
  target="_blank" rel="noopener">activate</a> the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slot</a> with the index specified
  in
  [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html)::`pSetupReferenceSlot->slotIndex`.

- The reconstructed picture is associated with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#decode-h264-reference-info"
  target="_blank" rel="noopener">H.264 reference information</a>
  provided in `pStdReferenceInfo`.

<!-- -->

Std Reference Information  
The members of the `StdVideoDecodeH264ReferenceInfo` structure pointed
to by `pStdReferenceInfo` are interpreted as follows:

- `flags.top_field_flag` is used to indicate whether the reference is
  used as top field reference;

- `flags.bottom_field_flag` is used to indicate whether the reference is
  used as bottom field reference;

- `flags.used_for_long_term_reference` is used to indicate whether the
  picture is marked as “used for long-term reference” as defined in
  section 8.2.5.1 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

- `flags.is_non_existing` is used to indicate whether the picture is
  marked as “non-existing” as defined in section 8.2.5.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>;

- all other members are interpreted as defined in section 8.2 of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#itu-t-h264"
  target="_blank" rel="noopener">ITU-T H.264 Specification</a>.

Valid Usage (Implicit)

- <a href="#VUID-VkVideoDecodeH264DpbSlotInfoKHR-sType-sType"
  id="VUID-VkVideoDecodeH264DpbSlotInfoKHR-sType-sType"></a>
  VUID-VkVideoDecodeH264DpbSlotInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_DECODE_H264_DPB_SLOT_INFO_KHR`

- <a
  href="#VUID-VkVideoDecodeH264DpbSlotInfoKHR-pStdReferenceInfo-parameter"
  id="VUID-VkVideoDecodeH264DpbSlotInfoKHR-pStdReferenceInfo-parameter"></a>
  VUID-VkVideoDecodeH264DpbSlotInfoKHR-pStdReferenceInfo-parameter  
  `pStdReferenceInfo` **must** be a valid pointer to a valid
  `StdVideoDecodeH264ReferenceInfo` value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_decode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_decode_h264.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoDecodeH264DpbSlotInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
